package main

import (
	"fmt"
	"os"

	"github.com/rerost/bq-table-validator/cmd"
)

func main() {
	fmt.Println("RUNCMD")
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
