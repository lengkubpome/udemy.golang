package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"os"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// กรณีที่ 1 ข้อมูลจากเว็บ
	// resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	// todoDecoder := json.NewDecoder(resp.Body)
	// todos := []Todo{}
	// todoDecoder.Decode(&todos)
	// resp.Body.Close()

	// กรณีที่ 2 ข้อมูลจากไฟล์
	resp, err := ioutil.ReadFile("D:/Code/Golang/src/udemy.golang/htmlTemplate/todos.json") // ดึงจากไฟล์
	if err != nil {
		return
	}
	todoDecoder := json.NewDecoder(bytes.NewReader(resp))
	todos := []Todo{}
	todoDecoder.Decode(&todos)

	indexTemplateBytes, err := ioutil.ReadFile("D:/Code/Golang/src/udemy.golang/htmlTemplate/index.html")
	if err != nil {
		return
	}
	indexTemplate := template.Must(template.New("index").Parse(string(indexTemplateBytes)))
	indexTemplate.Execute(os.Stdout, todos)
}
