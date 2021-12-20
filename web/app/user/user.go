package user

import (
	"net/http"
	"log"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Handler for our logged-in user page.
func Handler(ctx *gin.Context) {
	log.Print("LOGGED IN USER HANDLER: starting")
	session := sessions.Default(ctx)
	profile := session.Get("profile")
	ctx.HTML(http.StatusOK, "user.html", profile)
}
