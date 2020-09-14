package http

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// RunHTTPServer setups and runs the http server
func RunHTTPServer() {
	port := ":8080"
	setupRouter()

	clearScreen()
	fmt.Println("Listen on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
