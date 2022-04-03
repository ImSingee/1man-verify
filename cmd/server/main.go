package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var app = &cobra.Command{
	Use:           "1man-verify",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func main() {
	err := app.Execute()
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}
