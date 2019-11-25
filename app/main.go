package main

import (
	"core_gin/app/core"
	"fmt"
)

func main() {
	router := core.NewEngine()
	router.GET("/test", func(ctx *core.Context) {
		fmt.Println("get request")

		// todo parse query param
		pm := ctx.Request.URL.Query()
		if v, ok := pm["id"]; ok {
			fmt.Println("request url", ctx.Request.URL.String(), " parameter id value =", v)
		}
		ctx.ResponseWriter.WriteHeader(200)
		ctx.ResponseWriter.Write([]byte("hello core gin."))
		//todo json render
		//r := render.JSON{Data:"success"}
		//r.WriteContentType(ctx.ResponseWriter)
		//
		//if err := r.Render(ctx.ResponseWriter); err != nil{
		//	panic(err)
		//}
	})
	router.Run(":2222")
}
