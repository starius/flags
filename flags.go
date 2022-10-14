package flags

import (
	"log"
	"os"

	goflags "github.com/jessevdk/go-flags"
)

func Parse(config interface{}) {
	errWrongCommand := 2

	_, err := goflags.Parse(config)
	if err != nil {
		if err, ok := err.(*goflags.Error); ok && err.Type == goflags.ErrHelp {
			os.Exit(errWrongCommand)
		}
		log.Fatalf("Error during flags parsing: %v.", err)
	}
}
