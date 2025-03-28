package gozshprompt

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

type Handler func(ctx context.Context, input string) string
type InitHandler func(ctx context.Context) string
type SuggestHandler func(ctx context.Context, input string, cursor int) []string

type handlers struct {
	p *Prompt

	init    InitHandler
	handler Handler
	suggest SuggestHandler

	handlerType   string
	suggestCursor int
}

func (h *handlers) handle(ctx context.Context, input string) string {
	if h.handler != nil {
		return h.handler(ctx, input)
	}
	fmt.Fprintln(h.p.GetStderr(), "Do nothing.")
	return ""
}

func (h *handlers) handleInit(ctx context.Context) string {
	if h.init != nil {
		return h.init(ctx)
	}
	return ""
}

func (h *handlers) handleSuggest(ctx context.Context, input string, cursor int) []string {
	if h.suggest != nil {
		return h.suggest(ctx, input, cursor)
	}
	return []string{}
}

func (h *handlers) solveInput(ctx context.Context, input string) (string, error) {
	if h.handlerType == "" {
		h.handlerType = input
		h.suggestCursor = -1

		switch h.handlerType {
		case "init":
			newSource := h.handleInit(ctx)

			h.handlerType = ""
			return newSource + "\n", nil
		}
	} else {
		switch h.handlerType {
		case "handle":
			newSource := h.handle(ctx, input)

			h.handlerType = ""
			return newSource + "\n", nil

		case "suggest":
			if h.suggestCursor == -1 {
				h.suggestCursor, _ = strconv.Atoi(input)
				return "", nil
			} else {
				suggests := h.handleSuggest(ctx, input, h.suggestCursor)

				h.handlerType = ""
				return strings.Join(suggests, "\n") + "\n\n", nil
			}
		}
	}
	return "", nil
}
