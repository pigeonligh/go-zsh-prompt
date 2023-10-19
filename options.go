package gozshprompt

import "io"

type PromptOption func(*Prompt)

func WithHandler(h Handler) PromptOption {
	return func(p *Prompt) {
		p.handler = h
	}
}

func WithStdin(stdin io.Reader) PromptOption {
	return func(p *Prompt) {
		p.stdin = stdin
	}
}

func WithStdout(stdout io.Writer) PromptOption {
	return func(p *Prompt) {
		p.stdout = stdout
	}
}

func WithStderr(stderr io.Writer) PromptOption {
	return func(p *Prompt) {
		p.stderr = stderr
	}
}
