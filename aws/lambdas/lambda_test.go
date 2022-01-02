package lambdas

// import (
// 	"fmt"

// 	"github.com/aws/aws-sdk-go/service/lambda"
// )

// type svctest struct{}

// func (s svctest) ListFunctions(input *lambda.ListFunctionsInput) (*lambda.ListFunctionsOutput, error) {
// 	return nil, nil
// }

// func (s svctest) GetFunction(input *lambda.GetFunctionInput) (*lambda.GetFunctionOutput, error) {
// 	return nil, nil
// }

// func (s svctest) CreateFunction(input *lambda.CreateFunctionInput) (*lambda.FunctionConfiguration, error) {
// 	return nil, nil
// }

// func (s svctest) DeleteFunction(input *lambda.DeleteFunctionInput) (*lambda.DeleteFunctionOutput, error) {
// 	return &lambda.DeleteFunctionOutput{}, nil
// }

// func lambdaService() *Lambdas {
// 	var test svctest
// 	return &Lambdas{
// 		svc: test,
// 	}
// }

// func ExampleLambda_DeleteFunction() {
// 	lambda := lambdaService()

// 	result, err := lambda.DeleteFunction("simple-deploy")
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(result)
// }

// func createFunction() {
// 	options := FunctionsConfig{
// 		UseLocalConfig: true,
// 		Region:         "us-east-2",
// 	}
// 	lambda := NewLambdas(options)

// 	data := CreateFunctionData{
// 		S3Bucket:     "phirmware-lambda-test",
// 		S3Key:        "function.zip",
// 		Description:  "simple API deployed lambda function",
// 		FunctionName: "simple-deploy",
// 		Handler:      "index.handler",
// 		Runtime:      "nodejs12.x",
// 		Role:         "arn:aws:iam::937105331058:role/phirmware-lambda",
// 	}
// 	result, err := lambda.CreateFunction(data)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(result)
// }
