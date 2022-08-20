package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()
	servers := []string{
		"https://platzi.com",
		"https://google.com",
		"https://facebook.com",
		"https://instagram.com",
	}

	for _, server := range servers {
		serverReview(server)
	}

	pastTime := time.Since(inicio)
	fmt.Printf("Tiempo de ejecución %s\n", pastTime)
}

func serverReview(server string) {
	_, err := http.Get(server)
	if err != nil {
		fmt.Println(server, "server is not available")
	} else {
		fmt.Println(server, "is working")
	}
}
