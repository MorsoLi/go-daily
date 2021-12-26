package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world! go web")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("form.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm", err)
		}
		form := "username :" + r.FormValue("username")
		//fmt.Println("form name: ", r.Form["username"])
		//fmt.Println("form passwd: ", r.Form["password"])
		fmt.Fprintf(w, form)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		current_time := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(current_time, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8080", nil)
}
