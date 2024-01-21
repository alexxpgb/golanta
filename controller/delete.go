package controller

import (
	"fmt"
	"golanta/data"
	"net/http"
	"strconv"
)

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
	data.LstAventurier, err = ReadJSON()
	if err != nil {
		fmt.Fprintf(w, "Error reading JSON file : %s", err)
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Println("erreur")
	}

	for i, v := range data.LstAventurier {
		if v.Id == id {
			data.LstAventurier = append(data.LstAventurier[:i], data.LstAventurier[i+1:]...)
			//Supprime de la liste des personnage
		}

	}
	EditJSON(data.LstAventurier)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
