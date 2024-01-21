package main

import (
	"golanta/routeur"
	"golanta/temps"
)

func main() {
	temps.InitTemplate()
	routeur.InitServe()
}
