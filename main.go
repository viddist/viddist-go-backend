package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	link := os.Args[1]
	fmt.Println(link)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}
