package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
	//"encoding/binary"
)

func main() {
	file, err := ioutil.ReadFile("\\example.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	data := yaml.MapSlice{}
	err = yaml.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- example.yaml:\n%v\n\n", data)
	/*
		var bytes []byte
		bytes = append(bytes, []byte("a")...)
		bytes = append(bytes, []byte("s")...)
		bytes = append(bytes, []byte("m")...)
		bytes = append(bytes, []byte("l")...)

		for i := 0; i < len(bytes); i++ {
			fmt.Println(data[i].Value.(string))
		}
		fmt.Println(bytes)
	*/
}
