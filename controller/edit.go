package controller

import (
	"fmt"
	"golanta/data"
	"golanta/temps"
	"net/http"
	"strconv"
)

var err error
var Id int

func Edit(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test", r.URL.Query().Get("id"))
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Println("erreur")
	}
	for _, t := range data.LstAventurier {
		if t.Id == id {
			data.Section = t
			break
		}
	}
	Id = id
	temps.Temp.ExecuteTemplate(w, "Edit", data.Section)

}

func EditArticle(w http.ResponseWriter, r *http.Request) {
	data.LstAventurier, err = ReadJSON()
	if err != nil {
		fmt.Println("erreur")
	}
	var c int
	for _, t := range data.LstAventurier {
		if t.Id == Id {
			data.Section.Nom = r.FormValue("Nom")
			data.Section.Prenom = r.FormValue("Prenom")
			data.Section.Equipe = r.FormValue("Equipe")
			break
		}
		c++

	}
	data.LstAventurier[c] = data.Section
	EditJSON(data.LstAventurier)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
