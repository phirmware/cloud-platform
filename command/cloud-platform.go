package command

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/phirmware/cloud-platform/aws"
)

func init() {
	log.SetPrefix(fmt.Sprintf("[%5d] ", os.Getpid()))
}

func Execute() {
	// parse yaml definition
	file := flag.String("file", "test.yaml", "File path to definiiton")
	flag.Parse()

	log.Println(flag.Args())
	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	platform := flag.Arg(0)
	cmd := flag.Arg(1)

	switch platform {
	case "aws":
		aws.Execute(*file, cmd)
	}
}
