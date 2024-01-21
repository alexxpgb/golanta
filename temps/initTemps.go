package temps

// ce code va init tous les templates
import (
	"fmt"
	"html/template"
	"os"
)

// variabl globale qui permet de executert les templates
var Temp *template.Template

func InitTemplate() {
	temp, errTemp := template.ParseGlob("./temps/*.html")
	if errTemp != nil {
		fmt.Printf("Oupss une erreur lié aux Templates : %v", errTemp.Error())
		os.Exit(1)
	}
	Temp = temp
}
