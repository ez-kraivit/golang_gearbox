package middleware

import (
	"github.com/gogearbox/gearbox"
)

func CorsMiddleware(ctx gearbox.Context) {
	// ctx.Context().Request.Header.Set("Access-Control-Allow-Origin", "kraivit.com")
	ctx.Context().Request.Header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	header := ctx.Status(gearbox.StatusUnauthorized).Get("Access-Control-Allow-Origin")
	if header != "kraivit.com" {
		ctx.Status(401)
		return
	} else {
		ctx.Next()
	}

}
