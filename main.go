package main

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("error while reading config file:\n", err)
	}

	src := viper.GetString("src")
	if src == "" {
		log.Fatalf("configuration error: src must not be empty")
	}

	dest := viper.GetString("dest")
	if src == "" {
		log.Fatalf("configuration error: dest must not be empty")
	}

	_, err = os.ReadDir(src)
	if errors.Is(err, os.ErrNotExist) {
		log.Fatalf("source directory does not exist: %s", src)
	} else if err != nil {
		log.Fatalf("error while reading source directory: %w", err)
	}

	_, err = os.ReadDir(dest)
	if errors.Is(err, os.ErrNotExist) {
		log.Fatalf("destination directory does not exist: %s", dest)
	} else if err != nil {
		log.Fatalf("error while reading destination directory: %w", err)
	}
}
