package tests

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2/utils"
	"github.com/joho/godotenv"
	"github.com/laraviet/email-service-fiber/routes"
)

func ApiRoute(method string, body io.Reader) *http.Request {
	r1 := httptest.NewRequest(method, "/email/send", body)
	r1.Header.Set("Content-Type", "application/json")
	return r1
}

func AddBasicAuth(req *http.Request) {
	req.SetBasicAuth(os.Getenv("USER_BASIC_AUTH"), os.Getenv("PASSWORD_BASIC_AUTH"))
}

func Test_Email_Send_Enpoint(t *testing.T) {
	t.Parallel()
	godotenv.Load(".env.testing")
	app := routes.SetRoutes()

	t.Run("GET 401", func(t *testing.T) {
		resp, _ := app.Test(ApiRoute("GET", nil))
		utils.AssertEqual(t, 401, resp.StatusCode, "Return 401 Forbidden")
	})

	t.Run("GET 405", func(t *testing.T) {
		r1 := ApiRoute("GET", nil)
		AddBasicAuth(r1)
		resp, _ := app.Test(r1)
		utils.AssertEqual(t, 405, resp.StatusCode, "Return 405 Method Not Allow")
	})

	t.Run("POST 200", func(t *testing.T) {
		bodyReader := strings.NewReader(`{"subject": "test",
			"from":    {"email": "thanhcttsp@gmail.com"},
			"to":      {"email": "eric.n@liv3ly.io"},
			"content": {
				"type": "text/html",
				"value": "<strong>this is test email</strong>"
			}
		}`)
		r1 := ApiRoute("POST", bodyReader)
		AddBasicAuth(r1)
		resp, _ := app.Test(r1)
		defer resp.Body.Close()
		bodyR, _ := ioutil.ReadAll(resp.Body)
		utils.AssertEqual(t, 200, resp.StatusCode, "Return 200 Successful")
		utils.AssertEqual(t, `{"message":"parse successful","status":"ok"}`, string(bodyR), "Return Successful Message")
	})
}
