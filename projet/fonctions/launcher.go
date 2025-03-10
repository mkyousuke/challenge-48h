package fonctions

import (
	"fmt"
)

// Lance l'application web
func LaunchApp() {
	fmt.Println("Lancement de l'application web")
	err := IsWordPressSite("https://www.starwarsblog.net")
	if err != nil {
		fmt.Println("Erreur lors de la vérification du site WordPress: ", err)
	} else {
		fmt.Println("Le site EST un site WordPress")
	}
}
