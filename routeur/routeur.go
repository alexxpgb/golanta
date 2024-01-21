package routeur

import (
	"fmt"
	"golanta/controller"
	"net/http"
	"os"
)

func InitServe() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/create", controller.Create)
	http.HandleFunc("/create/addarcticle", controller.AddArticle)
	http.HandleFunc("/edit", controller.Edit)
	http.HandleFunc("/edit/editarticle", controller.EditArticle)
	http.HandleFunc("/delete/deletearticle", controller.DeleteArticle)
	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	fmt.Println("(http://localhost:8080/) - Server started on port:8080")
	// permet de control click
	http.ListenAndServe("localhost:8080", nil)
	// demarre le serveur
	fmt.Println("Server closed")
}
