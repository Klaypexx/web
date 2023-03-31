package main

import (
	"html/template"
	"log"
	"mime"
	"net/http" // служит в Go основным средством для разработки HTTP-клиентов и серверов
	"path/filepath"
)

const (
	port = ":3000"
)

type indexPage struct {
	Title         string
	Subtitle      string
	AuthorImg     string
	FeaturedPosts []featuredPostData
}

type featuredPostData struct {
	Title       string
	Subtitle    string
	Author      string
	PublishDate string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", index) //если url будет "/home", то он вызывает функцию index

	// Реализуем отдачу статики
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/static"))))

	cssContentType := mime.TypeByExtension(filepath.Ext("static/css/styles.css"))
	http.HandleFunc("/static/css/styles.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", cssContentType)
		http.ServeFile(w, r, "static/css/styles.css")
	})

	log.Println("Start server " + port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := indexPage{
		Title:         "Blog for traveling",
		Subtitle:      "My best blog for adventures and burgers",
		AuthorImg:     "static/img/background.png",
		FeaturedPosts: featuredPosts(),
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func featuredPosts() []featuredPostData {
	return []featuredPostData{
		{
			Title:       "The Road Ahead",
			Subtitle:    "The road ahead might be paved - it might not be.",
			Author:      "Mat Vogels",
			PublishDate: "54 September, 1488a",
		},
	}
}
