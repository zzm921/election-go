package cmd

import (
	"context"

	"election/internal/consts"
	"election/internal/controller"
	"election/internal/lib/validator"
	"election/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server of election",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(ghttp.MiddlewareHandlerResponse)
			validator.ValiidatorInit()
			s.Group("/crm", func(group *ghttp.RouterGroup) {
				// Group middlewares.
				group.Middleware(
					ghttp.MiddlewareCORS,
				)
				// Register route handlers.
				group.Bind(
					controller.Account,
					controller.Electoin,
					controller.Candidate,
				)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().AccountAuth)
					group.Bind(
						controller.Electoin,
						controller.Candidate,
					)
				})
			})

			s.Group("/user", func(group *ghttp.RouterGroup) {
				// Group middlewares.
				group.Middleware(
					ghttp.MiddlewareCORS,
				)
				group.Bind(
					controller.User,
				)
			})

			// Custom enhance API document.
			enhanceOpenAPIDoc(s)
			// Just run the server.
			s.Run()
			return nil
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Contact: &goai.Contact{
			Name: "GoFrame",
			URL:  "https://goframe.org",
		},
	}
}
