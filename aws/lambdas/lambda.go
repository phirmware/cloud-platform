package lambdas

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type LambdaServiceInterface interface {
	ListFunctions(input *lambda.ListFunctionsInput) (*lambda.ListFunctionsOutput, error)
	GetFunction(input *lambda.GetFunctionInput) (*lambda.GetFunctionOutput, error)
	CreateFunction(input *lambda.CreateFunctionInput) (*lambda.FunctionConfiguration, error)
	DeleteFunction(input *lambda.DeleteFunctionInput) (*lambda.DeleteFunctionOutput, error)
}

type Lambdas struct {
	region string
	svc    LambdaServiceInterface
}

func NewLambdas(config FunctionsConfig) *Lambdas {
	if config.UseLocalConfig {
		os.Setenv("AWS_SDK_LOAD_CONFIG", "true")
	}

	awsConfig := aws.NewConfig()
	if config.Region != "" {
		awsConfig.WithRegion(config.Region)
	}

	svc := lambda.New(session.New(), awsConfig)

	return &Lambdas{
		region: config.Region,
		svc:    svc,
	}
}

func (l *Lambdas) ListFunctions() ([]*lambda.FunctionConfiguration, error) {
	input := &lambda.ListFunctionsInput{}
	result, err := l.svc.ListFunctions(input)
	if err != nil {
		return nil, err
	}

	return result.Functions, nil
}

func (l *Lambdas) GetFunction(functionName string) (*lambda.GetFunctionOutput, error) {
	input := &lambda.GetFunctionInput{
		FunctionName: aws.String(functionName),
	}

	result, err := l.svc.GetFunction(input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (l *Lambdas) CreateFunction(data CreateFunctionData) (*lambda.FunctionConfiguration, error) {
	input := &lambda.CreateFunctionInput{
		Code: &lambda.FunctionCode{
			S3Bucket: aws.String(data.S3Bucket),
			S3Key:    aws.String(data.S3Key),
		},
		Description:  aws.String(data.Description),
		FunctionName: aws.String(data.FunctionName),
		Handler:      aws.String(data.Handler),
		MemorySize:   aws.Int64(256),
		Role:         aws.String(data.Role),
		Runtime:      aws.String(data.Runtime),
	}

	return l.svc.CreateFunction(input)
}

func (l *Lambdas) DeleteFunction(functionName string) (*lambda.DeleteFunctionOutput, error) {
	input := lambda.DeleteFunctionInput{
		FunctionName: aws.String(functionName),
	}

	result, err := l.svc.DeleteFunction(&input)
	if err != nil {
		return nil, err
	}

	return result, nil
}
