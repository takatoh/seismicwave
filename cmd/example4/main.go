package main

import (
	"fmt"
	"os"

	"github.com/takatoh/seismicwave"
	"github.com/BurntSushi/toml"
)

func main() {
	inputfile := os.Args[1]
	var input seismicwave.InputWave

	_, err := toml.DecodeFile(inputfile, &input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, w := range input.Waves {
		fmt.Printf("%#v\n", w)
	}
}
