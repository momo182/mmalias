package mmalias

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/gookit/goutil/strutil"
)

type Fdata map[string]interface{}

func Fstr(format string, data Fdata) (string, error) {
	t, err := template.New("fstring").Parse(format)
	FailOnErr(err, "creating template")
	output := new(bytes.Buffer)
	err = t.Execute(output, data)
	FailOnErr(err, "executing template")
	return output.String(), nil
}

func Cmd(s string) []string {
	result := strutil.Split(s, " ")
	return result
}

func FailOnErr(e error, msg string) {
	if e != nil {
		fmt.Println("error: "+msg+":", e)
		os.Exit(500)
	}
}
