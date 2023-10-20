package main

import (
	"context"

	gozshprompt "github.com/pigeonligh/go-zsh-prompt"
)

func main() {
	prompt, err := gozshprompt.NewFromPath()
	if err != nil {
		panic(err)
	}

	prompt.Run(context.Background())
}
