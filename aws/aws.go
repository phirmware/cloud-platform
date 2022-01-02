package aws

import (
	"io/ioutil"
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

func Execute(file string, cmd string) {
	switch cmd {
	case "apply":
		Apply(file, os.Open)
	}
}

func Apply(file string, open func(string) (*os.File, error)) {
	var def Definition
	f, _ := open(file)
	defer f.Close()
	yaml.ParseYamlDefinition(f, &def, ioutil.ReadAll)

	switch def.Spec.Resource.Type {
	case "lambda":
		executeLambda(file)
	}
}
