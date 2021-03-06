package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sfreiberg/gotwilio"
)

var (
	port string = os.Getenv("PORT")
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error { return inputHttp(c) })

	// routing to api
	api := e.Group("/api")
	api.POST("/call", func(c echo.Context) error { return apiCall(c) })

	e.Logger.Fatal(e.Start(":" + port))
}

func inputHttp(c echo.Context) error {
	req := c.Request()
	ip := strings.Split(req.RemoteAddr, ":")[0]
	return c.String(http.StatusOK, ip)
}

func apiCall(c echo.Context) error {
	request := new(Call)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Response: nil, Exception: nil, Error: err})
	}

	exeption := []*gotwilio.Exception{}
	err := new(error)

	for _, number := range request.To {
		_, exep, err := callTwilio(request.AccountSid, request.AuthToken, request.From, number, request.CallbackUrl, request.Method)
		exeption = append(exeption, exep)
		if err != nil {
			return c.JSON(http.StatusOK, Response{Response: "success", Exception: exeption, Error: err})
		}
	}

	return c.JSON(http.StatusOK, Response{Response: "success", Exception: exeption, Error: *err})
}

func callTwilio(accountSid, authToken, from, to, callbackUrl, method string) (*gotwilio.VoiceResponse, *gotwilio.Exception, error) {
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)
	callbackParams := gotwilio.NewCallbackParameters(callbackUrl)
	switch method {
	case "POST":
		callbackParams.Method = "POST"
	default:
		callbackParams.Method = "GET"
	}
	return twilio.CallWithUrlCallbacks(from, to, callbackParams)
}
