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

	i := 0

	for i < len(servers) {

		for _, server := range servers {
			go serverReview(server, canal)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-canal)
		i++

	}

	pastTime := time.Since(inicio)
	fmt.Printf("Tiempo de ejecución %s\n", pastTime)
}

func serverReview(server string, canal chan string) {
	_, err := http.Get(server)
	if err != nil {
		canal <- server + "- server is not available"
	} else {
		canal <- server + "- server is working"
	}
}
