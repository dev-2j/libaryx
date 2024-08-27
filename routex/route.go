package routex

import (
	"example.com/m/action"
	"example.com/m/service/connx"
)

func Route() {
	var GO_DEV = connx.New("example.com", "module", "service", 4000)
	GO_DEV.Connect(func(g *connx.GX) {
		g.GinEngine.GET("action/new", action.NewFunction)
		g.GinEngine.POST("action/new2", action.NewFunction2)
	})
}
