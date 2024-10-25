package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arizdn234/Marlin-Express/config"
	"github.com/arizdn234/Marlin-Express/kafka"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()
	fmt.Println("Database connected successfully!")

	kafkaWriter := kafka.NewKafkaWriter("kafka:9092", "package-updates")
	defer kafkaWriter.Close()
	fmt.Println("Kafka connected successfully!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Package Delivery Backend API is running!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
