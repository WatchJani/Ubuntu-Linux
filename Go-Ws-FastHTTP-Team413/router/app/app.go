package app

import (
	"crypto/tls"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
	"time"
	"websocket-client/router/handlers"
)

// Here we can put also drivers from dbs, or something else that is configurable
type App struct {
	Router *router.Router
}

func (a *App) CreateRouterAndSetRoutes(hd *handlers.HD) *router.Router {
	r := router.New()
	r.Handle("GET", "/login", hd.WsStreamHCLogin2)
	return r
}

// Run the app on it's router
func (a *App) Run() {

	//config := &conf.GlobalConfig{}
	//// dir := util.GetWorkDirectory()
	//util.ReadInstanceYaml(yamlFilePath, config)

	// https://go.dev/src/crypto/tls/example_test.go
	// https://gist.github.com/jim3ma/00523f865b8801390475c4e2049fe8c3
	c1, err := tls.LoadX509KeyPair("ssl/localhost.crt", "ssl/localhost.key")
	if err != nil {
		log.Fatal(err)
	}
	//==================================================================================================================
	//cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	//listener, err := tls.Listen("tcp", ":2000", cfg)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//_ = listener
	//==================================================================================================================

	tc := &tls.Config{
		Certificates: []tls.Certificate{c1},
		ClientCAs:    nil,
		RootCAs:      nil,

		InsecureSkipVerify: false, // test server certificate is not trusted.
	}

	srv := &fasthttp.Server{
		Handler:          a.Router.Handler,
		TLSConfig:        tc,
		DisableKeepalive: false,
		MaxConnsPerIP:    8,
		ReadBufferSize:   4096,
		WriteBufferSize:  4096,
		WriteTimeout:     15 * time.Second,
		ReadTimeout:      15 * time.Second,
		//Addr:             "localhost:8800",
		////WriteTimeout:      15 * time.Second,
		////ReadTimeout:       15 * time.Second,
		//ReadHeaderTimeout: 100 * time.Millisecond,
		//MaxHeaderBytes:    2048,
	}
	////http.ListenAndServe(":8080",r) - r is router!
	//
	log.Fatal(srv.ListenAndServe("localhost:8800"))
	//a.Router.ANY("/1", Logging(Index))
	//a.Router.Handle("GET", "/2", Logging(Index))  // Method not allowed
	//a.Router.Handle("GET", "/3", la1.WebService1) // Method not allowed
	//a.Router.GET("/hello/{name}", Logging(Hello))
	//log.Fatal(fasthttp.ListenAndServe("localhost:8800", a.Router.Handler))
	//log.Fatal(fasthttp.ListenAndServeTLS(":8082", "ssl/localhost.crt", "ssl/localhost.key", r.Handler))
}
