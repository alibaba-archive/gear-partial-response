package partial_test

import (
	"github.com/teambition/gear"
	partial "github.com/teambition/gear-partial-response"
)

func Example() {
	app := gear.New()
	app.Set(gear.SetSender, partial.New(partial.Options{Query: "fields"}))
}
