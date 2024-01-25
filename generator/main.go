//go:build generator

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

const (
	generatedHeader = "Code generated by Aiven. DO NOT EDIT."
	versionIDParam  = `/{version_id:latest|\d+}`
	configPrefix    = "GEN"
)

type config struct {
	Module      string `envconfig:"MODULE" default:"github.com/aiven/aiven-go-client-v3"`
	Package     string `envconfig:"PACKAGE" default:"aiven"`
	HandlerDir  string `envconfig:"HANDLER_DIR" default:"handler"`
	ConfigFile  string `envconfig:"CONFIG_FILE" default:"config.yaml"`
	ClientFile  string `envconfig:"CLIENT_FILE" default:"client_generated.go"`
	OpenAPIFile string `envconfig:"OPENAPI_FILE" default:"openapi.json"`
}

var pathClean = regexp.MustCompile(`\{[^{]+}`)

func main() {
	err := exec()
	if err != nil {
		log.Fatal(err)
	}
}

func exec() error {
	cfg := new(config)
	err := envconfig.Process(configPrefix, cfg)
	if err != nil {
		return err
	}

	configBytes, err := os.ReadFile(cfg.ConfigFile)
	if err != nil {
		return err
	}

	config := make(map[string][]string)
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		return err
	}

	docBytes, err := os.ReadFile(cfg.OpenAPIFile)
	if err != nil {
		return err
	}

	doc := new(Doc)
	err = json.Unmarshal(docBytes, doc)
	if err != nil {
		return err
	}

	pkgs := make(map[string][]*Path)
	for path, v := range doc.Paths {
		for meth, p := range v {
			if p.Deprecated {
				continue
			}

			p.Path = path
			p.Method = strings.ToUpper(meth)
			p.ID = p.OperationID

			var pkg string
		outer:
			for k, idList := range config {
				for _, id := range idList {
					if p.ID == id {
						pkg = k
						break outer
					}
				}
			}

			if pkg == "" {
				log.Printf("%q id not found in config. Skipping", p.ID)
				continue
			}

			pkgs[pkg] = append(pkgs[pkg], p)
			params := make([]*Parameter, 0, len(p.Parameters))
			for _, ref := range p.Parameters {
				parts := strings.Split(ref.Ref, "/")
				name := parts[len(parts)-1]
				param, ok := doc.Components.Parameters[name]
				if !ok {
					return fmt.Errorf("param %q not found", ref.Ref)
				}

				if param.In != ParameterInPath {
					log.Printf("%q param %s in %q", p.OperationID, param.Name, param.In)
				}

				param.Ref = ref.Ref
				params = append(params, param)
			}

			if strings.HasSuffix(p.Path, versionIDParam) {
				params = append(params, &Parameter{
					Name:   "version_id",
					Schema: &Schema{Type: SchemaTypeInteger},
				})
			}

			p.Parameters = params
		}
	}

	const doerName = "doer"

	ctx := jen.Id("ctx").Qual("context", "Context")
	doer := jen.Type().Id(doerName).Interface(
		jen.Id("Do").Params(
			ctx,
			jen.List(jen.Id("operationID"), jen.Id("method"), jen.Id("path")).String(),
			jen.Id("v").Any(),
		).Parens(jen.List(jen.Index().Byte(), jen.Error())),
	).Line()
	clientFields := make([]jen.Code, 0, len(pkgs))
	clientValues := jen.Dict{}
	clientTypeValues := make([]jen.Code, 0, len(pkgs))

	for _, pkg := range sortedKeys(pkgs) {
		paths := pkgs[pkg]
		fileName := strings.ToLower(pkg)
		handlerName := pkg + "Handler"
		handlerType := "Handler"
		newHandler := "New" + handlerType
		scope := make(map[string]*Schema)
		for _, p := range paths {
			p.FuncName = p.OperationID
		}

		sort.SliceStable(paths, func(i, j int) bool {
			return paths[i].FuncName < paths[j].FuncName
		})

		file := jen.NewFile(fileName)
		file.HeaderComment(generatedHeader)
		handler := file.Type().Id(handlerType)
		file.Func().Id(newHandler).Params(jen.Id(doerName).Id(doerName)).Id(handlerName).Block(
			jen.Return(jen.Id(handlerName).Values(jen.Id(doerName))),
		)
		file.Add(doer)
		file.Type().Id(handlerName).Struct(jen.Id(doerName).Id(doerName))
		typeMethods := make([]jen.Code, len(paths))
		for _, path := range paths {
			schemas := make([]*Schema, 0)
			params := make([]jen.Code, 0, len(path.Parameters))
			params = append(params, ctx)
			for _, p := range path.Parameters {
				p.Schema.required = true
				p.Schema.init(doc, scope, pkg, p.Name)
				schemas = append(schemas, p.Schema)
				param := jen.Id(strcase.ToLowerCamel(p.Schema.camelName)).Add(getType(p.Schema))
				params = append(params, param)
			}

			// todo: support 204
			out := path.Out.OK.Content["application/json"]
			if out == nil && path.Out.NoContent.Content == nil {
				log.Printf("%q has no json response. Skipping", path.OperationID)
				continue
			}

			in := path.In.Content["application/json"]
			if in != nil {
				schemaIn, err := doc.getSchema(in.Schema.Ref)
				if err != nil {
					return err
				}
				schemaIn.in = true
				schemaIn.init(doc, scope, "", path.FuncName+"In")
				schemas = append(schemas, schemaIn)
				params = append(params, jen.Id("in").Id("*"+schemaIn.camelName))
			}

			typeMeth := jen.Id(path.FuncName).Params(params...)
			structMeth := jen.Func().Params(jen.Id("h").Id("*" + handlerName)).Id(path.FuncName).Params(params...)

			var rsp, schemaOut *Schema
			if out != nil {
				schemaOut, err = doc.getSchema(out.Schema.Ref)
				if err != nil {
					return err
				}
				schemaOut.out = true
				schemaOut.init(doc, scope, "", path.FuncName+"Out")
				rsp = getResponse(schemaOut)
			}

			if rsp != nil {
				ret := jen.List(getType(rsp), jen.Error())
				typeMeth.Parens(ret)
				structMeth.Parens(ret)
			} else {
				typeMeth.Error()
				structMeth.Error()
			}

			typeMethods = append(typeMethods, path.Comment(), typeMeth.Line())

			paramIndex := -1
			url := pathClean.ReplaceAllStringFunc(path.Path, func(_ string) string {
				paramIndex++
				switch t := path.Parameters[paramIndex].Schema.Type; t {
				case SchemaTypeInteger:
					return "%d"
				case SchemaTypeString:
					return "%s"
				default:
					panic(fmt.Sprintf("%s unexpected parameter type %s", path.OperationID, t))
				}
			})
			urlParams := make([]jen.Code, 0, len(params))
			urlParams = append(urlParams, jen.Lit(url))
			inObj := jen.Nil()
			for _, s := range schemas {
				if s.isObject() {
					inObj = jen.Id("in")
					continue
				}
				urlParams = append(urlParams, jen.Id(strcase.ToLowerCamel(s.camelName)))
			}

			outObj := jen.Id("_")
			returnErr := jen.Return(jen.Err())
			if rsp != nil {
				outObj = jen.Id("b")

				// In most cases, "nil" is for error return
				// But for required scalars should be zero values
				returnErr = jen.Return(jen.Nil(), jen.Err())
				if rsp.required || rsp.Type == SchemaTypeString {
					switch rsp.Type {
					case SchemaTypeString:
						returnErr = jen.Return(jen.Lit(""), jen.Err())
					case SchemaTypeInteger, SchemaTypeNumber:
						returnErr = jen.Return(jen.Lit(0), jen.Err())
					case SchemaTypeBoolean:
						returnErr = jen.Return(jen.False(), jen.Err())
					}
				}
			}

			block := []jen.Code{
				jen.Id("path").Op(":=").Qual("fmt", "Sprintf").Call(urlParams...),
				jen.List(outObj, jen.Err()).Op(":=").Id("h.doer.Do").Call(
					jen.Id("ctx"),
					jen.Lit(path.OperationID),
					jen.Lit(path.Method),
					jen.Id("path"),
					inObj,
				),
			}

			if rsp == nil {
				block = append(block, jen.Return(jen.Err()))
			} else {
				outReturn := jen.Id("out")
				if rsp.camelName != schemaOut.camelName {
					// Takes original name and turns to camel.
					// "camelName" field might have been modified because of name collisions
					outReturn.Dot(strcase.ToCamel(rsp.name))
				}

				block = append(
					block,
					jen.Id("out").Op(":=").New(jen.Id(schemaOut.camelName)),
					jen.Err().Op("=").Qual("encoding/json", "Unmarshal").Call(jen.Id("b"), jen.Id("out")),
					jen.If(jen.Err().Op("!=").Nil()).Block(returnErr),
					jen.Return(outReturn, jen.Nil()),
				)
			}

			file.Add(structMeth.Block(block...))
		}

		for _, k := range sortedKeys(scope) {
			err = writeStruct(file, scope[k])
			if err != nil {
				return err
			}
		}

		dirPath := filepath.Join(cfg.HandlerDir, fileName)
		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}

		handler.Interface(typeMethods...)
		err = file.Save(filepath.Join(dirPath, fileName+".go"))
		if err != nil {
			return err
		}

		pkgName := filepath.Join(cfg.Module, cfg.HandlerDir, fileName)
		clientFields = append(clientFields, jen.Qual(pkgName, handlerName))
		clientValues[jen.Id(handlerName)] = jen.Qual(pkgName, newHandler).Call(jen.Id(doerName))
		clientTypeValues = append(clientTypeValues, jen.Qual(pkgName, handlerType))
	}

	client := jen.NewFile(cfg.Package)
	client.HeaderComment(generatedHeader)
	client.Add(doer)
	client.Func().Id("newClient").Params(jen.Id(doerName).Id(doerName)).Id("Client").Block(
		jen.Return(jen.Id("&client").Values(clientValues)),
	)
	client.Type().Id("client").Struct(clientFields...)
	client.Type().Id("Client").Interface(clientTypeValues...)
	return client.Save(cfg.ClientFile)
}

// reMakesSense sometimes there are invalid enums, for instance just a comma ","
var reMakesSense = regexp.MustCompile(`\w`)

func writeStruct(f *jen.File, s *Schema) error {
	if s.isMap() || s.isArray() || s.isScalar() && !s.isEnum() {
		return nil
	}

	if s.isEnum() {
		kind := getScalarType(s)
		o := f.Type().Id(s.camelName)
		o.Add(kind)
		enums := make([]jen.Code, len(s.Enum))
		values := make([]jen.Code, len(s.Enum))
		for _, e := range s.Enum {
			literal := fmt.Sprint(e)
			if !reMakesSense.MatchString(literal) {
				continue
			}

			constant := s.camelName + strcase.ToCamel(literal)

			// KafkaMirror ReplicationPolicyClassType makes bad generated name
			if strings.HasPrefix(literal, "org.apache.kafka.connect.mirror.") {
				constant = s.camelName + literal[32:len(literal)-17]
			}

			// OpenSearch HealthType has value "red*"
			if strings.HasSuffix(literal, "*") {
				constant += "Asterisk"
			}
			enums = append(enums, jen.Id(constant).Op(s.camelName).Op("=").Lit(literal))
			values = append(values, jen.Lit(literal))
		}

		if len(enums) == 0 {
			return nil
		}

		o.Line().Const().Defs(enums...)

		if !s.isOut() {
			o.Line().Func().Id(s.camelName + "Choices").Params().Index().Add(kind).Block(
				jen.Return(jen.Index().Add(kind).Values(values...)),
			)
		}
		return nil
	}

	fields := make([]jen.Code, 0, len(s.Properties))
	for _, k := range s.propertyNames {
		p := s.Properties[k]
		field := jen.Id(strcase.ToCamel(k)).Add(getType(p))
		tag := k
		if !p.required {
			if !p.isArray() {
				tag += ",omitempty"
			} else {
				if p.MinItems > 0 || p.Default != nil {
					// Slices never omitted.
					// But we must make sure they don't have default values or min items constraint.
					// So no default value will be overriden, and validation is OK with an empty list.
					return fmt.Errorf("%q has minItems=%d and default=%v", k, p.MinItems, p.Default)
				}
			}
		}

		field = field.Tag(map[string]string{"json": strings.ReplaceAll(tag, `\`, "")})
		fields = append(fields, field)
	}

	f.Type().Id(s.camelName).Struct(fields...)
	return nil
}

func getResponse(s *Schema) *Schema {
	switch len(s.Properties) {
	case 1:
		// If the schema has just one field, then uses it as out dto.
		// That makes code simpler.
		// Then we don't need the original response object to be exposed.
		s.camelName = lowerFirst(s.camelName)
		return s.Properties[s.propertyNames[0]]
	case 0:
		return nil
	}
	return s
}

func toSingle(src string) string {
	s := strings.TrimSuffix(src, "ies")
	if s != src {
		return s + "y"
	}
	return strings.TrimSuffix(src, "s")
}
