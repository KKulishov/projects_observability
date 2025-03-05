package main

import (
	"log"
	"net/http"
	"os"
	"pushsimple/internal/bad"

	"github.com/grafana/pyroscope-go"
)

func main() {
	serverAddress := os.Getenv("PYROSCOPE_SERVER_ADDRESS")
	if serverAddress == "" {
		serverAddress = "http://localhost:4040"
	}

	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: "simple.golang.app",
		ServerAddress:   serverAddress,
		Logger:          pyroscope.StandardLogger,
	})
	if err != nil {
		log.Fatalf("error starting pyroscope profiler: %v", err)
	}

	http.HandleFunc("/fastapp", bad.FastHandler)
	http.HandleFunc("/slowapp", bad.SlowHandler)
	http.HandleFunc("/mem-leak", bad.MemLeakHandler) // Утечка памяти
	http.HandleFunc("/gorout", bad.GoroutineHandler) // Утечка памяти

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
