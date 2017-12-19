package partial

import (
	"net/http"
	"testing"

	"github.com/DavidCai1993/request"
	"github.com/stretchr/testify/suite"
	"github.com/teambition/gear"
)

const (
	Addr  = ":28080"
	Addr2 = ":28081"
)

type MiddlewareSuite struct {
	suite.Suite
	app             *gear.App
	defualtQueryApp *gear.App
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

	s.defualtQueryApp = gear.New()
	s.defualtQueryApp.Set(gear.SetSender, New(Options{}))
	s.defualtQueryApp.Use(func(ctx *gear.Context) error {
		a := "aaa"
		b := 3
		c := float64(2.2)

		return ctx.Send(http.StatusOK, S1{A: &a, B: &b, C: &c})
	})
	go func() {
		s.Require().Nil(s.defualtQueryApp.Listen(Addr2))
	}()
}

func (s *MiddlewareSuite) TestVaildQuery() {
	json, err := request.Get("http://127.0.0.1" + Addr + "?c=a,c").JSON(new(S1))

	s.Nil(err)
	s.Equal(*(json.(*S1).A), "aaa")
	s.Nil(json.(*S1).B)
	s.Equal(*(json.(*S1).C), float64(2.2))
}

func (s *MiddlewareSuite) TestOtherQuery() {
	json, err := request.Get("http://127.0.0.1" + Addr + "?c=e,f").JSON(new(S1))

	s.Nil(err)
	s.Nil(json.(*S1).A)
	s.Nil(json.(*S1).B)
	s.Nil(json.(*S1).C)
}

func (s *MiddlewareSuite) TestEmptyQuery() {
	json, err := request.Get("http://127.0.0.1" + Addr + "?c=").JSON(new(S1))

	s.Nil(err)
	s.Equal(*(json.(*S1).A), "aaa")
	s.Equal(*(json.(*S1).B), 3)
	s.Equal(*(json.(*S1).C), float64(2.2))
}

func (s *MiddlewareSuite) TestWithoutQuery() {
	json, err := request.Get("http://127.0.0.1" + Addr + "?d=d,v&&e=f").JSON(new(S1))

	s.Nil(err)
	s.Equal(*(json.(*S1).A), "aaa")
	s.Equal(*(json.(*S1).B), 3)
	s.Equal(*(json.(*S1).C), float64(2.2))
}

func (s *MiddlewareSuite) TestDefaultQueryApp() {
	json, err := request.Get("http://127.0.0.1" + Addr2 + "?fields=a,c").JSON(new(S1))

	s.Nil(err)
	s.Equal(*(json.(*S1).A), "aaa")
	s.Nil(json.(*S1).B)
	s.Equal(*(json.(*S1).C), float64(2.2))
}

func TestMiddleware(t *testing.T) {
	suite.Run(t, new(MiddlewareSuite))
}
