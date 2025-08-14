package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var request struct {
		Message string `json:"message"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	response := fmt.Sprintf("收到訊息: %s", request.Message)
	fmt.Fprintln(w, response)
}

func main() {
	// 靜態文件服務
	webDir := "./web/"
	fs := http.FileServer(http.Dir(webDir))
	http.Handle("/", fs)
	
	// API 路由
	http.HandleFunc("/api/status", statusHandler)
	http.HandleFunc("/api/hello", helloHandler)
	
	fmt.Println("Server running at http://localhost:8080")
	if absPath, err := filepath.Abs(webDir); err == nil {
		fmt.Println("Static files served from:", absPath)
	}
	http.ListenAndServe(":8080", nil)
}
