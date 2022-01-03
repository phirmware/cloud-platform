package aws

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/phirmware/cloud-platform/aws/lambdas"
)

func createTempFile(t *testing.T, pattern string, initialValue string) (*os.File, func()) {
	t.Helper()
	f, err := ioutil.TempFile("", pattern)
	if err != nil {
		t.Fatalf("Could not create file: %s", err)
	}

	if _, err := f.Write([]byte(initialValue)); err != nil {
		t.Fatalf("Could not write initial data to file: %s", err)
	}

	removeFile := func() {
		f.Close()
		os.Remove(f.Name())
	}

	return f, removeFile
}
func TestApply(t *testing.T) {
	file := "testfile.yaml"
	content := `apiversion: v1
platform: aws
spec:
 resource:
  type: lambda`

	f, close := createTempFile(t, file, content)
	defer close()

	called := false
	executeLambda = func(file string) {
		called = true
	}
	// reset executeLambda to original value after test runs
	defer func() { executeLambda = lambdas.Execute }()

	open := func(s string) (*os.File, error) {
		if s != file {
			t.Fatalf("Apply(): invalid file name passed into apply, got: %s, want: %s", s, file)
		}
		return f, nil
	}

	reader := func(r io.Reader) ([]byte, error) {
		return os.ReadFile(r.(*os.File).Name())
	}

	Apply(file, open, reader)

	if !called {
		t.Errorf("Apply(): Invalid case run for resource type: expected executeLambda to have been called")
	}
}
