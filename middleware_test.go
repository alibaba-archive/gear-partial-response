package partial

import (
	"net/http"
	"testing"

	"github.com/DavidCai1993/request"

	"github.com/stretchr/testify/suite"
	"github.com/teambition/gear"
)

const (
	Addr = ":28080"
)

type MiddlewareSuite struct {
	suite.Suite
	app *gear.App
}

type S1 struct {
	A *string  `json:"a"`
	B *int     `json:"b"`
	C *float64 `json:"c"`
}

func (s *MiddlewareSuite) SetupSuite() {
	s.app = gear.New()
	s.app.Set(gear.SetSender, New(Options{Query: "c"}))
	s.app.Use(func(ctx *gear.Context) error {
		a := "aaa"
		b := 3
		c := float64(2.2)

		return ctx.Send(http.StatusOK, S1{A: &a, B: &b, C: &c})
	})
	go func() {
		s.Require().Nil(s.app.Listen(Addr))
	}()
}

func (s *MiddlewareSuite) TestVaildQuery() {
	json, err := request.Get("http://127.0.0.1" + Addr + "?c=a,c").JSON(new(S1))

	s.Nil(err)

	s.Equal(*(json.(*S1).A), "aaa")
	s.Nil(json.(*S1).B)
	s.Equal(*(json.(*S1).C), float64(2.2))
}

func TestMiddleware(t *testing.T) {
	suite.Run(t, new(MiddlewareSuite))
}
