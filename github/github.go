package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/wrzolekLukasz")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: %s", resp.Status)
	}
	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatalf("error: cant copy - %s", err)
	}

}
