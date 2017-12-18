package partial

import (
	mask "github.com/DavidCai1993/json-mask-go"
	"github.com/teambition/gear"
)

// Options is the partial response middleware options.
type Options struct {
	// Query specifies the querystring to use. By defaults it is "fields".
	Query string
}

// Sender is to implement gear.Sender interface.
type Sender struct {
	query string
}

// Send is to implement gear.Sender interface.
func (s *Sender) Send(ctx *gear.Context, code int, data interface{}) error {
	if s.query == "" {
		return ctx.JSON(code, data)
	}

	maskedData, err := mask.Mask(data, s.query)
	if err != nil {
		return ctx.JSON(code, data)
	}

	return ctx.JSON(code, maskedData)
}

// New returns a new partial response middleware for your gear app.
func New(opts Options) *Sender {
	if opts.Query == "" {
		opts.Query = "fields"
	}

	return &Sender{query: opts.Query}
}
