// Copyright 2017 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// gen-accessors generates accessor methods for structs with pointer fields.
//
// It is meant to be used by go-github contributors in conjunction with the
// go generate tool before sending a PR to GitHub.
// Please see the CONTRIBUTING.md file for more information.
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
)

func main() {
	src := `
package main

import "fmt"
import "strings"

func main() {
    hello := "Hello"
    world := "World"
    words := []string{hello, world}
    SayHello(words)
}

// SayHello says Hello
func SayHello(words []string) {
    fmt.Println(joinStrings(words))
}

// joinStrings joins strings
func joinStrings(words []string) string {
    return strings.Join(words, ", ")
}

type Engine struct {

}

type Car struct {
   muStr string
   myPointerStr *string
   enginePointer *Engine
   engine Engine
   myMap map[string]string
   myMapOfPointers map[string]*Engine
   myArray []string
   myArrayOfStructs []Engine
   myArrayOfPointers []*Engine
}
`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	for _, f := range f.Decls {
		dcl, ok := f.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, s := range dcl.Specs {
			ts, ok := s.(*ast.TypeSpec)
			if !ok {
				continue
			}
			fmt.Println(ts.Name.Name)

			// Skip unexported identifiers.
			if !ts.Name.IsExported() {
				fmt.Printf("Struct %v is unexported; skipping.\n", ts.Name)
				continue
			}
			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				continue
			}
			for _, field := range st.Fields.List {
				fmt.Println("---")
				fmt.Printf("Field name: %s\n", field.Names[0].Name)
				// fmt.Printf("Field type: %s\n", reflect.TypeOf(field.Type))

				switch v := field.Type.(type) {
				case *ast.Ident:
					fmt.Println("Ident:", v)
				case *ast.StarExpr:
					//fmt.Println("StarExpr:", v)
					fmt.Println("Pointer of:", v.X.(*ast.Ident).Name)
				case *ast.MapType:
					//fmt.Println("MapType:", v)
					fmt.Println("Map key", v.Key.(*ast.Ident).Name)

					switch mV := v.Value.(type) {
					case *ast.StarExpr:
						fmt.Println("Map pointer value", mV.X.(*ast.Ident).Name)
					case *ast.Ident:
						fmt.Println("Map ident", mV.Name)
					default:
						fmt.Println("Unsupported type??", reflect.TypeOf(mV))
					}

				case *ast.ArrayType:
					switch aV := v.Elt.(type) {
					case *ast.StarExpr:
						fmt.Println("Array pointer value", aV.X.(*ast.Ident).Name)
					case *ast.Ident:
						fmt.Println("Array ident", aV.Name)
					default:
						fmt.Println("Unsupported type??", reflect.TypeOf(aV))
					}
				default:
					fmt.Println("Unsupported type??", reflect.TypeOf(v))
				}
			}
		}
	}
}

// x = Parent Object we are unmarshalling
// v = fastJson "Value" that was parsed

type UnmarshalStringTemplateData struct {
	FieldName    string
	APIFieldName string
}

const UnmarshalStringTemplate = `
	x.{{ .FieldName }} = string(v.GetStringBytes("{{ .APIFieldName }}"))
`

type UnmarshalPointerObjTemplateData struct {
	FieldType    string
	APIFieldName string
}

const UnmarshalPointerObjTemplate = `
	{
		xField := &{{ .FieldType }}{}
		err := xField.UnmarshalFromObj(v.GetObject("{{ .APIFieldName }}"))
		if err != nil {
			return err
		}
	}
`

type UnmarshalArrayTemplateData struct {
	FieldType    string
	APIFieldName string
}

const UnmarshalArrayTemplate = `
	{
		for _, jV := range v.GetArray("{{.APIFieldName}}") {
			{{ .UnmarshalElementTemplate }}
			x.{{.FieldName}} = append(x.{{.FieldName}}, y)
		}
	}
`

type UnmarshalMapTemplateData struct {
	ValueType    string
	APIFieldName string
}

const UnmarshalMapTemplate = `
	{
		oMap := make(map[string]{{.ValueType}})
		jO := v.GetObject("{{.APIFieldName}}")
		if jO != nil {
			jO.Visit(func(key []byte, val *fastjson.Value) {
				y[(key)] = {{TBD}}
			})
		}
	}
`

type UnmarshalPointerObjFunctionTemplateData struct {
	ModelType string
}

const UnmarshalPointerObjFunctionTemplate = `
func (i *{{ .ModelType }}) Unmarshal(j string) error {
	{{- $lowerModelType := lowerCaseFirst .ModelType }}

	p := {{$lowerModelType}}JsonParserPool.Get()
	defer {{$lowerModelType}}JsonParserPool.Put(p)
	v, err := p.Parse(j)

	{{- range .Fields -}}
	{{ renderUnmarshalField . }}
	{{end -}}

	return nil
}
`
