package main

import (
	"heartbit/articles"
	"heartbit/utils"
	"log"
	"net/http"
)

var (
	address   = ":5000"
	urlPrefix = "/api/1.0"
)

func main() {
	mux := http.NewServeMux()

	articles.InitializeDatabase()
	articleGroupName := "/article"
	mux.Handle(urlPrefix+articleGroupName+"/", ArticlesEndpoints(mux, articleGroupName))

	log.Printf("Server starting on: http://localhost%s\n", address)
	http.ListenAndServe(address, mux)
}

func ArticlesEndpoints(router *http.ServeMux, groupName string) http.Handler {
	articlePrefix := urlPrefix + groupName

	router.Handle(articlePrefix+"/list", utils.Get(http.HandlerFunc(articles.ListArticle)))
	router.Handle(articlePrefix+"/get", utils.Get(http.HandlerFunc(articles.GetArticle)))
	router.Handle(articlePrefix+"/create", utils.Post(http.HandlerFunc(articles.CreateArticle)))
	router.Handle(articlePrefix+"/update", utils.Put(http.HandlerFunc(articles.UpdateArticle)))
	router.Handle(articlePrefix+"/delete", utils.Delete(http.HandlerFunc(articles.DeleteArticle)))

	return http.StripPrefix(articlePrefix, router)
}
