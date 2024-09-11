package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("config/frontend.yml")
	if err != nil {
		return
	}
	log.Println(string(file))

	var a map[string]interface{}

	err = yaml.Unmarshal(file, &a)
	if err != nil {
		log.Fatal(err)
		return
	}

}
