package controller

import (
	"encoding/json"
	"fmt"
	"golanta/data"
	"math/rand"
	"os"
)

func ReadJSON() ([]data.Aventurier, error) {
	jsonFile, err := os.ReadFile("data/data.json")
	if err != nil {
		fmt.Println("error reading", err.Error())
	}
	var jsonData []data.Aventurier
	err = json.Unmarshal(jsonFile, &jsonData)
	return jsonData, err
}

func EditJSON(ModifiedAventurier []data.Aventurier) {

	modifiedJSON, errMarshal := json.Marshal(ModifiedAventurier)
	if errMarshal != nil {
		fmt.Println("error encoding", errMarshal.Error())
		return
	}
	err := os.WriteFile("data/data.json", modifiedJSON, 0644)
	if err != nil {
		fmt.Println("erreur lors de l'ecriture du json edited", err)
		return
	}
}

func CallId() int {
	var Id = rand.Intn(50)

	for _, c := range data.LstAventurier {
		if c.Id == Id {
			return CallId()
		}
	}
	return Id
}
