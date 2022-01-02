package lambdas

type (
	LambdaService struct {
		lambda *Lambdas
		Region string
	}

	FunctionsConfig struct {
		// CLients aws region
		Region string

		// use config and credentials file at (~/.aws/config) and (~/.aws/credentials)
		UseLocalConfig bool
	}

	CreateFunctionData struct {
		S3Bucket     string `json:"s_3_bucket,omitempty"`
		S3Key        string `json:"s_3_key,omitempty"`
		Description  string `json:"description,omitempty"`
		FunctionName string `json:"function_name,omitempty"`
		Handler      string `json:"handler,omitempty"`
		Role         string `json:"role,omitempty"`
		Runtime      string `json:"runtime,omitempty"`
	}
)
