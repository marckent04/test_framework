package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

func main() {
	files := []string{
		"form",
		"table",
		"visual",
		"details",
	}

	_, filename, _, _ := runtime.Caller(0)
	currDir := filepath.Dir(filename)

	for _, file := range files {
		http.HandleFunc(fmt.Sprintf("/%s", file), func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL.Path, "requested")
			filePath := path.Join(currDir, fmt.Sprintf("%s.html", file))
			http.ServeFile(w, r, filePath)
		})
	}

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		const code = 200
		w.WriteHeader(code)
	})

	const port = 3000
	const timeout = 3
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		ReadHeaderTimeout: timeout * time.Second,
	}

	log.Printf("tests e2e server launched on port %d\n", port)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
