package main

import (
	"net/http"
	"html/template"
	"fmt"
)

type M map[string]interface{}

func main()  {
	var tmpl, err = template.ParseGlob("views/*")

	if  err != nil {
		panic(err.Error())
		return
	}


	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request){
		var data = M {"name": "indo"}
		err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Batman"}
		err = tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println("Serve started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}