package lambdas

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/phirmware/cloud-platform/yaml"
)

type Definition struct {
	APIVersion string         `yaml:"apiVersion"`
	Platform   string         `yaml:"platform"`
	Action     string         `yaml:"action"`
	Spec       DefinitionSpec `yaml:"spec"`
}

type DefinitionSpec struct {
	Resource Resource `yaml:"resource"`
	Config   Config   `yaml:"config"`
}

type Resource struct {
	Type   string       `yaml:"type"`
	Spec   ResourceSpec `yaml:"spec"`
	Source Source       `yaml:"source"`
}

type Config struct {
	UseLocalConfig bool   `yaml:"useLocalConfig"`
	Region         string `yaml:"region,omitempty"`
}

type ResourceSpec struct {
	Description  string `yaml:"description,omitempty"`
	FunctionName string `yaml:"functionName"`
	Handler      string `yaml:"handler"`
	Runtime      string `yaml:"runtime"`
	Role         string `yaml:"role"`
}

type Source struct {
	Type     string   `yaml:"type"`
	Metadata Metadata `yaml:"metadata"`
}

type Metadata struct {
	S3Bucket string `yaml:"s3Bucket"`
	S3Key    string `yaml:"s3Key"`
}

func Execute(file string) {
	var def Definition

	f, _ := os.Open(file)
	defer f.Close()

	yaml.ParseYamlDefinition(f, &def, ioutil.ReadAll)

	config := def.Spec.Config
	l := NewLambdas(FunctionsConfig{UseLocalConfig: config.UseLocalConfig, Region: config.Region})

	switch def.Action {
	case "create":
		createLambdaFunction(l, def, log.Println)
	case "manage":
		manageLambda(l, def)
	default:
		log.Fatal("Invalid 'action' value in YAML definition")
	}
}

func manageLambda(l *Lambdas, def Definition) {
	functionName := def.Spec.Resource.Spec.FunctionName
	// TODO: after testing the lambda package, write tests for this function manageLambda()
	function, err := l.GetFunction(functionName)
	if err != nil {
		if awserr, ok := err.(awserr.Error); ok {
			switch awserr.Code() {
			case lambda.ErrCodeResourceNotFoundException:
				createLambdaFunction(l, def, log.Println)
			default:
				log.Fatalf(awserr.Error())
			}
		} else {
			log.Fatalf("Error getting lambda info: Err: %s", err)
		}
		return
	}

	functionConfig := function.Configuration
	if isSameConfig(functionConfig, def.Spec.Resource.Spec) {
		log.Println("No config change")
	} else {
		fmt.Println("Config change, Updating resource")
		// update resource on aws
	}
}

func isSameConfig(config *lambda.FunctionConfiguration, def ResourceSpec) bool {
	// we support some keys in the spec, so we check for changes in the values of those keys
	// If new keys are added, we need to add them to this check if they are to be managed
	type ManageConfig struct {
		Description  string
		FunctionName string
		Handler      string
		Runtime      string
		Role         string
	}

	incomingConfig := ManageConfig(def)
	currentConfig := ManageConfig{
		Description:  *config.Description,
		FunctionName: *config.FunctionName,
		Handler:      *config.Handler,
		Runtime:      *config.Runtime,
		Role:         *config.Role,
	}

	return reflect.DeepEqual(incomingConfig, currentConfig)
}

func createLambdaFunction(l *Lambdas, def Definition, logger func(v ...interface{})) {
	result, err := l.CreateFunction(CreateFunctionData{
		S3Bucket:     def.Spec.Resource.Source.Metadata.S3Bucket,
		S3Key:        def.Spec.Resource.Source.Metadata.S3Key,
		Description:  def.Spec.Resource.Spec.Description,
		FunctionName: def.Spec.Resource.Spec.FunctionName,
		Handler:      def.Spec.Resource.Spec.Handler,
		Runtime:      def.Spec.Resource.Spec.Runtime,
		Role:         def.Spec.Resource.Spec.Role,
	})
	if err != nil {
		log.Fatalf("ERROR: %+v", err)
	}
	log.Println(result)
}

func Delete(file string) {
	var def Definition

	f, _ := os.Open(file)
	defer f.Close()

	yaml.ParseYamlDefinition(f, &def, ioutil.ReadAll)

	config := def.Spec.Config
	l := NewLambdas(FunctionsConfig{UseLocalConfig: config.UseLocalConfig, Region: config.Region})

	_, err := l.DeleteFunction(def.Spec.Resource.Spec.FunctionName)
	if err != nil {
		log.Fatalf("Could not delete lambda funciton: Error: %s", err)
	}

	log.Println("Done")
}
