package login

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"log"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"hundred-go/platform/authenticator"
)

// Handler for our login.
func Handler(auth *authenticator.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		state, err := generateRandomState()
		if err != nil {
			log.Print("LOGIN HANDLER: error generating state="+err.Error())
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Save the state inside the session.
		session := sessions.Default(ctx)
		session.Set("state", state)

		if err := session.Save(); err != nil {
			log.Print("LOGIN HANDLER: error saving state="+err.Error())
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		log.Print("LOGIN HANDLER: redirecting ")
		ctx.Redirect(http.StatusTemporaryRedirect, auth.AuthCodeURL(state))
	}
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	state := base64.StdEncoding.EncodeToString(b)
	return state, nil
}
