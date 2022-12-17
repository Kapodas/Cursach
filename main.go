package main

import (
	"html/template"
	"log"
	"net/http"
	//"golang.org/x/tools/go/analysis/passes/nilfunc"
	//"github.com/julienschmidt/httprouter"
	//"github.com/julienschmidt/httprouter"
)

type ContactDetails struct {
	Login         string
	Password      string
	Button        string
	Success       bool
	StorageAccess string
}

var (
	tmpl  = template.Must(template.ParseFiles("login.html"))
	tmpl2 = template.Must(template.ParseFiles("index.html"))
)

func LoginPage(w http.ResponseWriter, r *http.Request) { //p httprouter.Params) {
	data := ContactDetails{
		Login:    r.FormValue("login"),
		Password: r.FormValue("password"),
	}
	tmpl.Execute(w, data)
	if data.Login != "" && data.Password != "" {
		data.Success = true
		http.RedirectHandler("./index.html", 1)
		//	fmt.Fprintf(w, "Имя: %s Возраст: %s", data.Login, data.Password)
	}

	data.StorageAccess = "index.html"

}

func main() {
	//создаем и запускаем в работу роутер для обслуживания запросов
	//	r := httprouter.New()
	//	routes(r)
	//прикрепляемся к хосту и свободному порту для приема и обслуживания входящих запросов
	//вторым параметром передается роутер, который будет работать с запросами
	//	err := http.ListenAndServe("localhost:4444", r)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	http.HandleFunc("/", LoginPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*func routes(r *httprouter.Router) {
	//путь к папке со внешними файлами: html, js, css, изображения и т.д.
	r.ServeFiles("D:/Cursach", http.Dir("public"))
	//что следует выполнять при входящих запросах указанного типа и по указанному адресу
	r.GET("/", LoginPage)

}
*/
