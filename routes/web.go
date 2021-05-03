package routes

import (
	"goblog/app/http/controllers"
	"goblog/app/http/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterWebRoutes 注册网页相关的路由
func RegisterWebRoutes(r *mux.Router) {

	// 静态页面
	pc := new(controllers.PagesController)
	// r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.AboutHandler).Methods("GET").Name("about")
	r.NotFoundHandler = http.HandlerFunc(pc.NotFoundHandler)
	ac := new(controllers.ArticlesController)
	r.HandleFunc("/", ac.Index).Methods("GET").Name("home")
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")
	r.HandleFunc("/articles/create", ac.Create).Methods("GET").Name("articles.create")
	r.HandleFunc("/articles", ac.Store).Methods("POST").Name("articles.store")
	r.HandleFunc("/articles/{id:[0-9]+}/edit", ac.Edit).Methods("GET").Name("articles.edit")
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Update).Methods("POST").Name("articles.update")
	r.HandleFunc("/articles/{id:[0-9]+}/delete", ac.Delete).Methods("POST").Name("articles.delete")
	// 静态资源
	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))

	// 中间件：强制内容类型为 HTML
	//	r.Use(middlewares.ForceHTML)
	r.Use(middlewares.StartSession)

	// 用户认证
	auc := new(controllers.AuthController)
	r.HandleFunc("/auth/register", auc.Register).Methods("GET").Name("auth.register")
	r.HandleFunc("/auth/do-register", auc.DoRegister).Methods("POST").Name("auth.doregister")
	r.HandleFunc("/auth/login", auc.Login).Methods("GET").Name("auth.login")
	r.HandleFunc("/auth/dologin", auc.DoLogin).Methods("POST").Name("auth.dologin")
	r.HandleFunc("/auth/logout", auc.Logout).Methods("POST").Name("auth.logout")

}
