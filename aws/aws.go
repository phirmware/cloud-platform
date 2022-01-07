package aws

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/phirmware/cloud-platform/aws/lambdas"
	"github.com/phirmware/cloud-platform/yaml"
)

type Definition struct {
	Spec struct {
		Resource struct {
			Type string
		}
	}
}

var executeLambda = lambdas.Execute
var deleteLambda = lambdas.Delete

func Execute(file string, cmd string) {
	switch cmd {
	case "apply":
		Apply(file, os.Open, ioutil.ReadAll)
	case "delete":
		Delete(file, os.Open, ioutil.ReadAll)
	}
}

func Apply(file string, open func(string) (*os.File, error), reader func(io.Reader) ([]byte, error)) {
	resourceType := getDefinitionResourceType(file, open, reader)

	switch resourceType {
	case "lambda":
		executeLambda(file)
	}
}

func Delete(file string, open func(string) (*os.File, error), reader func(io.Reader) ([]byte, error)) {
	resourceType := getDefinitionResourceType(file, open, reader)

	switch resourceType {
	case "lambda":
		deleteLambda(file)
	}
}

func getDefinitionResourceType(file string, open func(string) (*os.File, error), reader func(io.Reader) ([]byte, error)) string {
	var def Definition
	f, err := open(file)
	if err != nil {
		log.Fatalf("FATAL: failed to load file: %v", err)
	}

	defer f.Close()
	yaml.ParseYamlDefinition(f, &def, reader)

	return def.Spec.Resource.Type
}
