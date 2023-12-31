package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
		gozshprompt.WithInitHandler(func(ctx context.Context) string {
			return `export PS1="init > "`
		}),
		gozshprompt.WithHandler(func(ctx context.Context, input string) string {
			fmt.Printf("Solve %v\n", input)
			return `export PS1="` + input + ` > "`
		}),
		gozshprompt.WithSuggestHandler(func(ctx context.Context, input string, cursor int) string {
			return strings.Join([]string{input[:cursor], ".", input[cursor:]}, "")
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
