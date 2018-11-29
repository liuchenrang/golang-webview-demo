package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var pageMap=make(map[string]string)
var window *Window
var server *http.Server
var HttpPort ="8080"
var RPCPort="3333"
func startHttpHandler(){

	server=startHttpServer()
}


func startHttpServer() *http.Server {
	srv := &http.Server{Addr: ":"+HttpPort}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path:=r.URL.Path
		if strings.HasPrefix(path,"/static/"){
			s := strings.TrimLeft(path,"/")
			if b,e:=ioutil.ReadFile(s);e==nil{
				w.WriteHeader(200)
				w.Write(b)
			}
			return
		}
		fmt.Println(path)
		if pagePath:=pageMap[path];pagePath!=""{
			if b,e:=ioutil.ReadFile(pagePath);e==nil{
				w.WriteHeader(200)
				w.Write(b)
			}else{
				log.Printf("get page error ", e)
			}

		}
	})
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()

	return srv
}
