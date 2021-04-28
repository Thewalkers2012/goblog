package route

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var Router = mux.NewRouter() // 使用gorilla/mux 更改的一行

// RouteName2URL 通过路由名来获取 URL
func RouteName2URL(routeName string, pairs ...string) string {
	url, err := Router.Get(routeName).URL(pairs...)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return url.String()
}

// GetRouteVariable 获取 URL 路由参数
func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}
