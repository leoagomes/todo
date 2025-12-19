package main

import (
	"fmt"
	"leoagomes/todo/internal/config"
	"log"
)

func main() {
	cfg, err := config.LoadFileOrDefault(".todo.yml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	fmt.Printf("Config: %+v\n", cfg)
}
