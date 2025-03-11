package fonctions

import (
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)


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

	if doc.Find("meta[name='generator'][content^='WordPress']").Length() > 0 {
		return nil
	}

	return errors.New("le site n'est pas un site WordPress")
}

