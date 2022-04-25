package main

import (
	"log"

	"github.com/phlashdev/file-importer/pkg/diff"
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
	dest := viper.GetString("dest")

	differ, err := diff.NewDirectoryDiffer(src, dest)
	if err != nil {
		log.Fatalf("error creating differ: %s", err)
	}

	err = differ.Diff()
	if err != nil {
		log.Fatalf("error while executing diff: %s", err)
	}
}
