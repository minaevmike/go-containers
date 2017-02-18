package main

import (
	"flag"
	"text/template"
	"strings"
	"github.com/minaevmike/go-containers/consts"
	"log"
	"bytes"
	"io/ioutil"
	"os"
)

var (
	simpleTypes []string = []string{
		"bool",
		"byte",
		"complex128",
		"comples64",
		"error",
		"float32",
		"float64",
		"int",
		"int16",
		"int32",
		"int64",
		"int8",
		"rune",
		"string",
		"uint",
		"uint16",
		"uint32",
		"uint64",
		"uint8",
		"uintptr",
	}
)

type TemplateParams struct {
	ShowType    string
	PackageName string
	Type        string
}

func main() {
	c := flag.String("c", "list", "type of generated container, supported: [list]")
	t := flag.String("t", "int64", "specialization type of this container, e.g int64 or float64")
	o := flag.String("o", "./", "path to store generated file")
	flag.Parse()
	switch *c {
	case "list":
		tp := &TemplateParams{}
		tp.Type = *t
		tp.ShowType= strings.Title(*t)
		tp.PackageName = "main"
		tt, err := template.New("list").Parse(consts.ListTemplate)
		if err != nil {
			log.Printf("Error at parsing list template: %v", err)
			return
		}
		b := bytes.Buffer{}
		err = tt.Execute(&b, tp)
		if err != nil {
			log.Printf("Error at executing: %v", err)
			return
		}
		f := *o + "/" + *t + "_" + *c + ".go"
		ioutil.WriteFile(f, b.Bytes(), os.ModePerm)
		log.Printf("file saved to %s", f)
	default:
		log.Fatalf("Unkonw type: %s", *c)

	}
}
