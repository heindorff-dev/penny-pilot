package logout

import (
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	logoutUrl, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/v2/logout")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + c.Request.Host)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	logoutUrl.RawQuery = parameters.Encode()

	c.Redirect(http.StatusTemporaryRedirect, logoutUrl.String())
}
