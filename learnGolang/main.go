package main

import "fmt"
import "net/http"
import "html/template"
import "path"






func handlerIndex(w http.ResponseWriter, r *http.Request) {
    var message = "Welcome"
    w.Write([]byte(message))
}

func handlerHellow(w http.ResponseWriter, r *http.Request){
	var message = "Hello world"
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		var filepath = path.Join("views","index.html")
		var tmpl, err = template.ParseFiles(filepath)

		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data = map[string]interface{}{
			"title": "Learning Golang Web",
			"name": "Batman",
		}

		err = tmpl.Execute(w, data)
		if err  != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		
	})
	http.Handle("/static/", 
	http.StripPrefix("/static/", 
		http.FileServer(http.Dir("assets"))))

fmt.Println("server started at localhost:9000")
http.ListenAndServe(":9000", nil)
	var address = "localhost:9000"
	fmt.Printf("server started at %s\n", address)
	err :=http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Printf(err.Error())	
	}
}