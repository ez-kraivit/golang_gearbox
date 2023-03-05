package middleware

import (
	"log"
	"strconv"

	plugins "kgearboxs/plugins"

	"github.com/gogearbox/gearbox"
)

func LogMiddleware(ctx gearbox.Context) {
	color := plugins.NewColor()
	method := string(ctx.Context().Method())
	host := string(ctx.Context().Host())
	path := string(ctx.Context().Path())
	status_code := strconv.Itoa(ctx.Context().Response.StatusCode())
	log.SetPrefix(color.Purple + "[KRAIVIT]" + color.Reset + " | ")
	log.Printf("| %s %s | %s %s", color.Green+"["+status_code+"]"+color.Reset, host, method, path)
	ctx.Next()
}
