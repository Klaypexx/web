package main

import (
	"html/template"
	"log"
	"net/http" //служит в Go основным средством для разработки HTTP-клиентов и серверов
)

type indexPage struct {
	Title         string
	FeaturedPosts []featuredPostData
	MostRecent    []mostPostData
}

type postPage struct {
	Title string
}

type featuredPostData struct {
	Title       string
	Subtitle    string
	ImgModifier string
	Author      string
	AuthorImg   string
	PublishDate string
}

type mostPostData struct {
	Title       string
	Subtitle    string
	ImgModifier string
	Author      string
	AuthorImg   string
	PublishDate string
}

func index(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := indexPage{
		Title:         "Escape.",
		FeaturedPosts: featuredPosts(),
		MostRecent:    mostPosts(),
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
			ImgModifier: "static/img/featured_posts/road.png",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/author/mat.png",
			PublishDate: "September 25,2015",
		},
		{
			Title:       "From Top Down",
			Subtitle:    "Once a year, go someplace you're never been before.",
			ImgModifier: "static/img/featured_posts/top_down.png",
			Author:      "William Wong",
			AuthorImg:   "static/img/author/william.png",
			PublishDate: "September 25,2015",
		},
	}
}

func mostPosts() []mostPostData {
	return []mostPostData{
		{
			Title:       "Still Standing Tall",
			Subtitle:    "Life begins at the end of your comfort zone.",
			ImgModifier: "static/img/most_recent/standing.png",
			Author:      "William Wong",
			AuthorImg:   "static/img/author/william.png",
			PublishDate: "9/25/2015",
		},
		{
			Title:       "Sunny Side Up",
			Subtitle:    "No place is ever as bad as they tell you it's going to be.",
			ImgModifier: "static/img/most_recent/sunny.png",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/author/mat.png",
			PublishDate: "9/25/2015",
		},
		{
			Title:       "Water Falls",
			Subtitle:    "We travel not to espace life, but for life not to espace us.",
			ImgModifier: "static/img/most_recent/water.png",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/author/mat.png",
			PublishDate: "9/25/2015",
		},
		{
			Title:       "Through the Mist",
			Subtitle:    "Travel makes you see what a tiny place you occupy in the world.",
			ImgModifier: "static/img/most_recent/through.png",
			Author:      "William Wong",
			AuthorImg:   "static/img/author/william.png",
			PublishDate: "9/25/2015",
		},
		{
			Title:       "Awaken Early",
			Subtitle:    "Not all those who wander are lost",
			ImgModifier: "static/img/most_recent/awaken.png",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/author/mat.png",
			PublishDate: "9/25/2015",
		},
		{
			Title:       "Try it Always",
			Subtitle:    "The world is a book, and those who do not travel read only one page.",
			ImgModifier: "static/img/most_recent/try_it.png",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/author/mat.png",
			PublishDate: "9/25/2015",
		},
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/the_road_ahead.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := postPage{
		Title: "Escape.",
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}
