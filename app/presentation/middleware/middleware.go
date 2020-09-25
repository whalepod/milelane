package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORSHeaders enable gin to check CORS request.
// This prohibit unauthorized access.
// TODO: disable access from unpermitted ip/hosts.
func CORSHeaders() gin.HandlerFunc {
	allowOrigins := []string{
		"http://app.milelane.co",
		"https://app.milelane.co",
		"http://localhost:8080",
		"ws://127.0.0.1:5858",
	}

	return cors.New(cors.Config{
		AllowMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"X-Device-UUID",
			"Authorization",
		},
		AllowOriginFunc: func(origin string) bool {
			for _, o := range allowOrigins {
				if origin == o {
					return true
				}
			}
			return false
		},
		MaxAge: 24 * time.Hour,
	})
}
