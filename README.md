# Twicall

Twilio api wrapper.  
Demo site is opened by Heroku.  
https://twicall-0.herokuapp.com/api/call


## Example

```
$ cat data.json
{
  "account-sid": "xxxx",
  "auth-token": "xxxx",
  "from": "+81xxxx",
  "to": ["+81xxxx","+81yyyy"],
  "callback-url": "https://xxxx"
}

$ curl -v \
	-X POST \
	-H "Content-Type: application/json" \
	-d @data.json \
	https://twicall-0.herokuapp.com/api/call
```
