package main

import (
	"fmt"
	"log"

	"github.com/bdbrwr/bootdev_blogagg/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Cannot read config: %v", err)
	}
	fmt.Printf("Read config: %v\n", cfg)

	err = cfg.SetUser("bdbrwr")
	if err != nil {
		log.Fatalf("Cannot set current user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Cannot read config: %v", err)
	}
	fmt.Printf("Read config, second time: %v\n", cfg)
}
