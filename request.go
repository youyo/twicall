package main

// request format
type (
	Call struct {
		AccountSid  string   `json:"account-sid"`
		AuthToken   string   `json:"auth-token"`
		From        string   `json:"from"`
		To          []string `json:"to"`
		CallbackUrl string   `json:"callback-url"`
	}
)
