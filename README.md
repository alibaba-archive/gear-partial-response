# gear-partial-response
[![Build Status](https://travis-ci.org/teambition/gear-partial-response.svg?branch=master)](https://travis-ci.org/teambition/gear-partial-response)
[![Coverage Status](https://coveralls.io/repos/github/teambition/gear-partial-response/badge.svg?branch=master)](https://coveralls.io/github/teambition/gear-partial-response?branch=master)

Partial response middleware for Gear.

## Installation

```
go get -u github.com/teambition/gear-partial-response
```

## Documentation

API documentation can be found here: https://godoc.org/github.com/teambition/gear-partial-response

## Usage

```go
import (
	"github.com/teambition/gear"
	partial "github.com/teambition/gear-partial-response"
)
```

```go
app := gear.New()
app.Set(gear.SetSender, partial.New(partial.Options{Query: "fields"}))
```
