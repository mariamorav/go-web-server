package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()
	canal := make(chan string)
	servers := []string{
		"https://platzi.com",
		"https://google.com",
		"https://facebook.com",
		"https://instagram.com",
	}

	for _, server := range servers {
		go serverReview(server, canal)
	}

	for i := 0; i < len(servers); i++ {
		fmt.Println(<-canal)
	}

	pastTime := time.Since(inicio)
	fmt.Printf("Tiempo de ejecución %s\n", pastTime)
}

func serverReview(server string, canal chan string) {
	_, err := http.Get(server)
	if err != nil {
		fmt.Println(server, "server is not available")
		canal <- server + " server is not available"
	} else {
		fmt.Println(server, "is working")
		canal <- server + "server is working"
	}
}
