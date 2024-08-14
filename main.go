package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var input = make(map[string]string)

	all, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("could not read json input")
		return
	}

	err = json.Unmarshal(all, &input)

	if err != nil {
		log.Fatal("could not unmarshal json input")
	}

	for k, v := range input {
		decodeString, err := base64.StdEncoding.DecodeString(v)
		if err != nil {
			log.Fatal("unable to decode value for " + k)
			return
		}

		fmt.Printf("%s: %s\n", k, decodeString)
	}
}
