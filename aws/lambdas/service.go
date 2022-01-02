package lambdas

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/phirmware/cloud-platform/yaml"
)

type Definition struct {
	APIVersion string `yaml:"apiVersion"`
	Platform   string `yaml:"platform"`
	Spec       struct {
		Resource struct {
			Type string `yaml:"type"`
		} `yaml:"resource"`
		Config struct {
			UseLocalConfig bool   `yaml:"useLocalConfig"`
			Region         string `yaml:"region,omitempty"`
		} `yaml:"config"`
		Action string `yaml:"action"`
		Spec   struct {
			S3Bucket     string `yaml:"s3Bucket"`
			S3Key        string `yaml:"s3Key"`
			Description  string `yaml:"description,omitempty"`
			FunctionName string `yaml:"functionName"`
			Handler      string `yaml:"handler"`
			Runtime      string `yaml:"runtime"`
			Role         string `yaml:"role"`
		} `yaml:"spec"`
	} `yaml:"spec"`
}

func Execute(file string) {
	var def Definition

	f, _ := os.Open(file)
	defer f.Close()

	yaml.ParseYamlDefinition(f, &def, ioutil.ReadAll)

	spec := def.Spec.Spec
	config := def.Spec.Config

	l := NewLambdas(FunctionsConfig{UseLocalConfig: config.UseLocalConfig, Region: config.Region})

	switch def.Spec.Action {
	case "create":
		result, err := l.CreateFunction(CreateFunctionData{
			S3Bucket: spec.S3Bucket,
			S3Key: spec.S3Key,
			Description: spec.Description,
			FunctionName: spec.FunctionName,
			Handler: spec.Handler,
			Runtime: spec.Runtime,
			Role: spec.Role,
		})
		if err != nil {
			log.Printf("ERROR: %+v", err)
		}
		log.Println(result)
	default:
		log.Println("Nothing")
	}
}
