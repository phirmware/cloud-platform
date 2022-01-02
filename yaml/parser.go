package yaml

import (
	"io"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func ParseYamlDefinition(
	file *os.File,
	definition interface{},
	reader func(io.Reader) ([]byte, error),
) error {
	data, err := reader(file)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, definition)
}
