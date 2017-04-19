package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
        "path/filepath"

	"github.com/ghodss/yaml"
)

var errEmpty = errors.New("stdin is empty")
var prog = filepath.Base(os.Args[0])
const usage = `< [FILENAME]
Convert the json contents of FILENAME to yaml and print it on stdout.
Program needs FILENAME in stdin with '<' sign.
The input is read completely before re-encoding begins.`

// read json from stdin
// print as yaml
func main() {
        //prog := filepath.Base(os.Args[0])

	if err := checkStdin(); err != nil {
                fmt.Fprintf(os.Stderr, "Usage: %s %s\n", prog, usage)
		os.Exit(1)
	}
	fmt.Println(readStdinConvertToYaml())
}

func checkStdin() (noError error) {
	fi, err := os.Stdin.Stat()

	if err != nil {
		return err
	}
	if fi.Size() == 0 {
		return errEmpty
	}

	return noError
}

func readStdinConvertToYaml() (s string) {
	if json, err := ioutil.ReadAll(os.Stdin); err != nil {
		s = err.Error()
	} else {
		s = json2Yaml(json)
	}
	return
}

func json2Yaml(json []byte) (s string) {
	if y, err := yaml.JSONToYAML(json); err != nil {
		s = err.Error()
	} else {
		s = string(y)
	}
	return
}
