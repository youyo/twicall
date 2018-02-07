# Twicall

Twilio api wrapper.  
Demo site is opened by Heroku.  
https://twicall-0.herokuapp.com/api/call


## Example

```
$ curl -v \
	-X POST \
	-H "Content-Type: application/json" \
	-d '{"account-sid":"xxxx","auth-token":"xxxx","from":"+81xxxx","to":"+81xxxx","callback-url":"https://xxxx"}' \
	https://twicall-0.herokuapp.com/api/call
```
