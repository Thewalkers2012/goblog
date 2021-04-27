package route

import (
	"log"

	"github.com/gorilla/mux"
)

var Router = mux.NewRouter() // 使用gorilla/mux 更改的一行

// Initialize 初始化路由
func Initialize() {
	Router = mux.NewRouter()
}

// RouteName2URL 通过路由名来获取 URL
func RouteName2URL(routeName string, pairs ...string) string {
	url, err := Router.Get(routeName).URL(pairs...)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return url.String()
}
