package http

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// RunHTTPServer setups and runs the http server
func RunHTTPServer() {
	handler := getHandler()
	
	s := &http.Server{
		Addr:           "0.0.0.0:8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	clearScreen()
	fmt.Println("Listen on ", s.Addr)
	log.Fatal(s.ListenAndServe())
}
