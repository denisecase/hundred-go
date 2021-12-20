package logout

import (
	"net/http"
	"net/url"
	"os"
	"log"

	"github.com/gin-gonic/gin"
)

// Handler for our logout.
func Handler(ctx *gin.Context) {
	logoutUrl, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/v2/logout")
	log.Print("LOGOUT HANDLER: logoutUrl="+logoutUrl.String())
	if err != nil {
		log.Print("LOGOUT HANDLER: err="+err.Error())
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}
	log.Print("LOGOUT HANDLER: scheme="+scheme)


	returnTo, err := url.Parse(scheme + "://" + ctx.Request.Host)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	log.Print("LOGOUT HANDLER: returnTo="+returnTo.String())

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	logoutUrl.RawQuery = parameters.Encode()

	//ctx.Redirect(http.StatusTemporaryRedirect, logoutUrl.String())
	ctx.Redirect(http.StatusTemporaryRedirect, returnTo.String())

}
