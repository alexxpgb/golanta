package controller

import (
	"fmt"
	"golanta/data"
	"golanta/temps"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
	temps.Temp.ExecuteTemplate(w, "create", nil)
}

func AddArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add")
	var err error
	data.LstAventurier, err = ReadJSON()
	if err != nil {
		fmt.Println("error encodage", err.Error())
		return
	}
	data.Section.Nom = r.FormValue("Nom")
	data.Section.Prenom = r.FormValue("Prenom")
	data.Section.Equipe = r.FormValue("Equipe")
	data.Section.Id = CallId()

	data.LstAventurier = append(data.LstAventurier, data.Section)
	EditJSON(data.LstAventurier)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}
