package main

import (
	"fmt"
	"os"

	"github.com/rerost/bq-table-validator/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
