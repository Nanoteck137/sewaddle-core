package main

import (
	"log"

	"github.com/spf13/cobra"
)

var AppName = "helper"

var rootCmd = &cobra.Command{
	Use: "helper",
}

func init() {
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
