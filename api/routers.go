package api

import "github.com/Yazzyk/skillfulPenCMS/app/user"

func routers() {
	v1 := app.Group("/v1")
	v1.Mount("/user", user.User())
}
