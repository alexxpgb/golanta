package data

type Aventurier struct {
	Id     int    `json:"id"`
	Prenom string `json:"Prenom"`
	Nom    string `json:"nom"`
	Equipe string `json:"Equipe"`
	Image  string `json:"Image"`
}

var Section Aventurier
var LstAventurier []Aventurier
