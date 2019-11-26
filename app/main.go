package main

import (
	"core_gin/app/core"
	"encoding/json"
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

	v1 := router.Group("/v1")
	v1.Use(middleWareHandler)
	v1.POST("/post", func(ctx *core.Context) {
		pm := ctx.Request.URL.Query()
		if v, ok := pm["id"]; ok {
			fmt.Println("request url", ctx.Request.URL.String(), " parameter id value =", v)
		}
		data := map[string]interface{}{}
		data["data"] = fmt.Sprintf("post id is %v", pm["id"][0])
		str, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		ctx.ResponseWriter.Header().Add("Content-Type", "application/json")
		//ctx.ResponseWriter.WriteHeader(200)
		ctx.ResponseWriter.Write(str)
	})

	router.Run(":2222")
}

func middleWareHandler(ctx *core.Context) {
	//todo ctx.Abort() and ctx.Next()
	fmt.Println("exec the middleware ")
}
