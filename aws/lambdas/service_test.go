package lambdas

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/lambda"
)

func TestIsSameConfig(t *testing.T) {
	Description := "test lambda deployed with CLI"
	FunctionName := "test-deploy"
	Handler := "index.handler"
	Runtime := "nodejs12.x"
	Role := "xxxx:37105331058:role/test-lambda"

	t.Run("isSameConfig() should return true when configs are same", func(t *testing.T) {
		resourceSpec := ResourceSpec{
			Description,
			FunctionName,
			Handler,
			Runtime,
			Role,
		}

		functionConfig := &lambda.FunctionConfiguration{
			Description:  &Description,
			FunctionName: &FunctionName,
			Handler:      &Handler,
			Runtime:      &Runtime,
			Role:         &Role,
		}

		got := isSameConfig(functionConfig, resourceSpec)
		want := true

		if got != want {
			t.Errorf("isSameConfig(): configuration check failed: want %t got %t", want, got)
		}
	})
	t.Run("isSameConfig() should return false when configs are different", func(t *testing.T) {
		resourceSpec := ResourceSpec{
			Description:  "desc",
			FunctionName: FunctionName,
			Handler:      Handler,
			Runtime:      Runtime,
			Role:         Role,
		}

		functionConfig := &lambda.FunctionConfiguration{
			Description:  &Description,
			FunctionName: &FunctionName,
			Handler:      &Handler,
			Runtime:      &Runtime,
			Role:         &Role,
		}

		got := isSameConfig(functionConfig, resourceSpec)
		want := false

		if got != want {
			t.Errorf("isSameConfig(): configuration check failed: want %t got %t", want, got)
		}
	})
}
