package fonctions

import (
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// func WordPressPageChecker(url string) bool {
// 	if len(os.Args) > 1 {
// 		url = os.Args[1]
// 	} else {
// 		fmt.Print("Entrez l'URL du site WordPress: ")
// 		fmt.Scanln(&url)
// 	}

// 	if isWordPressSite(url) {
// 		fmt.Printf("Le site %s est un site WordPress.\n", url)
// 	} else {
// 		fmt.Printf("Le site %s n'est pas un site WordPress.\n", url)
// 	}
// }

func IsWordPressSite(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return errors.New("erreur lors de la requête HTTP")
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return errors.New("erreur lors de l'analyse du document")
	}

	// Vérifie la présence de balises spécifiques à WordPress
	if doc.Find("meta[name='generator'][content^='WordPress']").Length() > 0 {
		return nil
	}

	return errors.New("le site n'est pas un site WordPress")
}

