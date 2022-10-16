package middleware

import (
	"election/internal/service"
	"fmt"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	sMiddleware struct{}
)

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}

// CORS allows Cross-origin resource sharing.
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// CORS allows Cross-origin resource sharing.
func (s *sMiddleware) AccountAuth(r *ghttp.Request) {
	header := r.Request.Header
	// var accountToken string = header['accountToken']
	// if accountToken == nil {

	// }]
	fmt.Println(header)

	r.Middleware.Next()
}
