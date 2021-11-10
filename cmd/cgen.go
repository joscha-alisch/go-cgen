package main

import (
	"github.com/joscha-alisch/go-cgen"
)

func main() {
	gen := cgen.New()

	err := gen.Run()
	if err != nil {
		panic(err)
	}
}
