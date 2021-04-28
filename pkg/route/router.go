package route

import (
	"goblog/pkg/logger"
	"net/http"

	"github.com/gorilla/mux"
)

var route *mux.Router // 使用gorilla/mux 更改的一行

// SetRoute 设置路由实例，以供 Name2URL 等函数的使用
func SetRoute(r *mux.Router) {
	route = r
}

// RouteName2URL 通过路由名来获取 URL
func RouteName2URL(routeName string, pairs ...string) string {
	url, err := route.Get(routeName).URL(pairs...)
	if err != nil {
		logger.LogError(err)
		return ""
	}

	return url.String()
}

// GetRouteVariable 获取 URL 路由参数
func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}
