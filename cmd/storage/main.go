package main

import (
	"fmt"
	"log"
	"storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("tests.txt", []byte("Hello, World!"))
	if err != nil {
		log.Fatal(err)
	}

	file, err = st.GetById(file.Id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Worked", file)
}
