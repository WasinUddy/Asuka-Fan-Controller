package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/WasinUddy/Ayanami-Fan-Controller/internal/controller"
	"github.com/WasinUddy/Ayanami-Fan-Controller/internal/middleware"
)

func main() {
	mux := http.NewServeMux()

	// Serve Static Files
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/", fs)

	// Serve API Endpoints
	mux.HandleFunc("/fan/status", controller.GetFanStatus)
	mux.HandleFunc("/fan/mode", controller.SetFanMode)
	mux.HandleFunc("/fan/speed", controller.SetFanSpeed)

	// Wrap with logging middleware
	loggedMux := middleware.LoggingMiddleware(mux)
	corsMux := middleware.CORS(loggedMux)

	// Start the server
	fmt.Println("🚀 Starting Asuka Fan Controller Web Server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsMux))
}
