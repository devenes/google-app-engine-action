package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	port := "80"

	// Check if the PORT environment variable is set
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	// Serve static files from the "./static" directory
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Start the HTTP server in a goroutine
	log.Printf("Server starting on port %s", port)
	go func() {
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatal(err)
		}
	}()

	// Open the browser with the URL
	url := "http://localhost:" + port
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Println(err)
	}

	// Stay alive
	select {}
}
