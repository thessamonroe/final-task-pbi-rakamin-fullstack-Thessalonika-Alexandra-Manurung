package main

import (
	"fmt"
	"log"
	"fullStack/app/database"
	"fullStack/app/router"

	"github.com/gin-gonic/gin"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatalf("Database gagal terkoneksi: %v", err)
	}

	r := router.SetupRouter()

	port := ":8080"
	fmt.Printf("Server berjalan di port %s\n", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Gagal untuk menjalankan server: %v", err)
	}
}

