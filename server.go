package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/sfreiberg/gotwilio"
)

var (
	port string = os.Getenv("PORT")
)

func main() {
	e := echo.New()

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

	twilio := gotwilio.NewTwilioClient(
		request.AccountSid,
		request.AuthToken,
	)

	callbackParams := gotwilio.NewCallbackParameters(request.CallbackUrl)

	response, exeption, err := twilio.CallWithUrlCallbacks(
		request.From,
		request.To,
		callbackParams,
	)

	return c.JSON(http.StatusOK, Response{Response: response, Exception: exeption, Error: err})
}
