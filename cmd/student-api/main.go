package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/proj/internal/config"
	"github.com/proj/internal/http/handlers/student"
	"github.com/proj/internal/storage/sqlite"
)

func main() {
	 
	cfg := config.MustLoad()

	//db setup
	storage,err:=sqlite.New(cfg)

	if err!=nil{
		log.Fatal(err)
	}
	slog.Info("storage  initialized", slog.String("env",cfg.Env))

	router := http.NewServeMux()

	router.HandleFunc("POST /api/students",student.New(storage))

	router.HandleFunc("GET /api/students/{id}",student.GetById(storage))
	router.HandleFunc("GET /api/students",student.GetList(storage))

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	log.Println("Server starting on", cfg.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
	fmt.Println("Server is up")
}
