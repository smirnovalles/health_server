package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os/exec"
)

type Response struct {
	Msg string `json:"msg"`
}

var (
	port       string
	scriptPath string
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command(scriptPath, "-v")
	output, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, "Error executing script", http.StatusInternalServerError)
		return
	}
	resp := Response{Msg: string(output)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	flag.StringVar(&port, "port", "8080", "Port for the server")
	flag.StringVar(&scriptPath, "script", "./health_check.sh", "Path to health_check.sh")
	flag.Parse()
	http.HandleFunc("/health_check", healthCheckHandler)
	fmt.Printf("Server running on port %sn", port)
	http.ListenAndServe(":"+port, nil)
}
