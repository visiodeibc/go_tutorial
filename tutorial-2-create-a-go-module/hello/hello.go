package main

import (
	"fmt"
	//log 라이브러리를 로드해준다.
	"greetings"
	"log"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	//여기 이름 추가
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	for _, messages := range message {
		fmt.Println(messages)
	}
}
