# Install and run in local
1. `cp .env.example .env`
2. Update .env for `SENDGRID_API_KEY` and `USER_BASIC_AUTH`, `PASSWORD_BASIC_AUTH`
3. `docker-compose up -d --build`

- Note: Server will auto reload once you have updated source code. No need to re-build docker after updating code.
- Warning: if you meet the issue when building docker `curl: (6) Could not resolve host: raw.githubusercontent.com` => then you need to update docker version to fix that

# Endpoint

POST `http://localhost:8010/email/send`

Header:
- Accept: application/json
- Content-Type: application/json
- Authorization: Basic {token} 

Body
```
{
	"subject": "test",
	"from": {"email": "thanhcttsp@gmail.com"},
	"to": {"email": "eric.n@liv3ly.io"},
	"content": {
		"type": "text/html",
		"value": "<strong>this is test email</strong>"
	}
}
```

Note:
- You can update `SENDGRID_API_KEY` in `.env` to use your own sendgrid_api_key
- Header.Authorization is the basic authen for `USER_BASIC_AUTH`, `PASSWORD_BASIC_AUTH` in `.env`

# UnitTest
- `cp .env.example tests/.env.testing`
- Update `tests/.env.testing` according to your test environment
- `go test -v ./tests`