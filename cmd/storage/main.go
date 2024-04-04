package main

import (
	"fmt"
	"github.com/mabduqayum/storage/v2/internal/storage"
	"log"
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
