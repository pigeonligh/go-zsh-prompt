package gozshprompt

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type Handler func(ctx context.Context, input string) error

type Prompt struct {
	executable string
	home       string

	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer

	handler Handler
}

func NewFromPath(options ...PromptOption) (*Prompt, error) {
	executable, err := exec.LookPath("zsh")
	if err != nil {
		return nil, err
	}

	return New(executable, options...), nil
}

func New(executable string, options ...PromptOption) *Prompt {
	p := &Prompt{
		executable: executable,
	}

	for _, option := range options {
		option(p)
	}

	return p
}

func (p *Prompt) SetHandler(h Handler) {
	p.handler = h
}

func (p *Prompt) SetStdin(stdin io.Reader) {
	p.stdin = stdin
}

func (p *Prompt) SetStdout(stdout io.Writer) {
	p.stdout = stdout
}

func (p *Prompt) SetStderr(stderr io.Writer) {
	p.stderr = stderr
}

func (p *Prompt) GetStdin() io.Reader {
	if p.stdin != nil {
		return p.stdin
	}
	return os.Stdin
}

func (p *Prompt) GetStdout() io.Writer {
	if p.stdout != nil {
		return p.stdout
	}
	return os.Stdout
}

func (p *Prompt) GetStderr() io.Writer {
	if p.stderr != nil {
		return p.stderr
	}
	return os.Stderr
}

func (p *Prompt) handle(ctx context.Context, input string) error {
	if p.handler != nil {
		return p.handler(ctx, input)
	}
	fmt.Fprintln(p.stderr, "Do nothing.")
	return nil
}

func (p *Prompt) Run(ctx context.Context) error {
	parentRead, parentWrite, childRead, childWrite, err := pipe()
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, p.executable)
	cmd.Env = []string{
		"HOME=" + p.home,
	}
	cmd.ExtraFiles = []*os.File{
		childRead,
		childWrite,
	}
	cmd.Stdin = p.GetStdin()
	cmd.Stdout = p.GetStdout()
	cmd.Stderr = p.GetStderr()

	routineCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func(ctx context.Context) {
		buffer := make([]byte, 0)
		for {
			block := make([]byte, 4096)
			n, err := parentRead.Read(block)
			if err != nil {
				// error
				return
			}
			buffer = append(buffer, block[:n]...)

			if i := bytes.IndexByte(buffer, 0); i != -1 {
				err := p.handle(ctx, string(buffer[:i]))
				if err != nil {
					// error
					return
				}
				_, _ = parentWrite.Write([]byte{0})

				buffer = buffer[i+1:]
			}
		}
	}(routineCtx)

	return cmd.Run()
}
