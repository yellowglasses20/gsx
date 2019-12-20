package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yellowGlasses20/gsx"
)

func main() {
	log.SetFlags(0)
	err := gsx.Run(os.Args[1:], os.Stdout, os.Stderr)
	if err != nil && err != flag.ErrHelp {
		log.Println(err)
		exitCode := 1
		if ecoder, ok := err.(interface{ ExitCode() int }); ok {
			exitCode = ecoder.ExitCode()
		}
		os.Exit(exitCode)
	}

}
