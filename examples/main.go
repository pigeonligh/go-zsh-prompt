package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	gozshprompt "github.com/pigeonligh/go-zsh-prompt"
)

func main() {
	home, err := filepath.Abs(".home")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	prompt, err := gozshprompt.NewFromPath(
		gozshprompt.WithHome(home),
		gozshprompt.WithHandler(func(ctx context.Context, input string) error {
			fmt.Printf("Solve %v\n", input)
			return nil
		}),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = prompt.Run(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
