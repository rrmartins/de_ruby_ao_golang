package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/rrmartins/de_ruby_ao_golang/listback/db"
)

func handlerDefault(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Route Invalid!")
}

func handlerListAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	lists := db.ListAll()
	lists_json, err := json.Marshal(lists)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(lists_json)
}

func handlerCreate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	title := r.FormValue("title")
	description := r.FormValue("description")
	go db.Insert(title, description)
	result, err := json.Marshal("result: true")

	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func handlerUpdate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	go db.Update(id, r.FormValue("title"), r.FormValue("description"))
	result, _ := json.Marshal("result: true")
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func handlerShow(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("idshow"))
	list := db.Show(id)
	result, _ := json.Marshal(list)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func handlerDelete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	go db.Delete(id)
	result, _ := json.Marshal("result: true")
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func main() {
	route := httprouter.New()
	route.GET("/", handlerDefault)

	route.GET("/list/all", handlerListAll)
	route.POST("/list/new", handlerCreate)
	route.PUT("/list/:id", handlerUpdate)
	route.DELETE("/list/:id", handlerDelete)
	route.GET("/lists/:idshow/show", handlerShow)

	log.Println("server start in port :3002")
	http.ListenAndServe(":3002", route)
}
