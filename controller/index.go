package controller

import (
	"encoding/json"
	"fmt"
	"golanta/data"
	"golanta/temps"
	"net/http"
	"os"
)

var Aventurier []data.Aventurier

func Index(w http.ResponseWriter, r *http.Request) {
	GetDataFromJson()
	temps.Temp.ExecuteTemplate(w, "index", &Aventurier)
}

func GetDataFromJson() {
	data, err := os.ReadFile("data/data.json") //ouverture et lecture du json
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	json.Unmarshal(data, &Aventurier) //passage en json vers la struct
}
