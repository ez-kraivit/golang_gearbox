package middleware

import (
	"fmt"
	"strings"

	"github.com/gogearbox/gearbox"
)

func AuthorizedMiddleware(ctx gearbox.Context) {
	header := ctx.Status(gearbox.StatusUnauthorized).Get("Authorization")
	ctx.Context().Value(ctx.Status(gearbox.StatusUnauthorized).Get("Authorization"))
	splitHeader := strings.Split(header, "Bearer ")
	if len(splitHeader) == 2 {
		ctx.Status(200)
		token := splitHeader[1]
		fmt.Println(token)
		ctx.Next()
	} else {
		ctx.Status(401)
		ctx.Next()
	}
}
