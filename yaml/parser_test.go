package yaml

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

type testDef struct {
	Name string
}

func TestParseYamlDefinition(t *testing.T) {
	var def testDef
	name := "test"
	f, err := ioutil.TempFile("", "test.yaml")
	if err != nil {
		t.Fatalf("ParseYamlDefinition(), Could not create temporary file: Error: %s", err)
	}

	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()

	f.Write([]byte(fmt.Sprintf(`name: %s`, name)))

	reader := func(r io.Reader) ([]byte, error) {
		return os.ReadFile(r.(*os.File).Name())
	}

	ParseYamlDefinition(f, &def, reader)
	
	if def.Name != name {
		t.Errorf("ParseYamlDefinition(): Expected %s got %s", name, def.Name)
	}
}
