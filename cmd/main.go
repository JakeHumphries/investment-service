// Package main is the entry point for the service
package main

import (
	"context"
	"log"

	"github.com/JakeHumphries/investment-service/service"
)

func main() {
	log.Println("Starting the investment service...")
	ctx := context.Background()
	s, err := service.NewService(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close(ctx)

	if err = s.Start(); err != nil {
		log.Fatal(err)
	}

	log.Println("Service started, waiting for requests...")

	s.Wait()
}
