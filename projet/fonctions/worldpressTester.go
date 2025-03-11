package fonctions

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jung-kurt/gofpdf"
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

	// Vérifie la présence de balises spécifiques à WordPress
	if doc.Find("meta[name='generator'][content^='WordPress']").Length() > 0 {
		return nil
	}

	return errors.New("le site n'est pas un site WordPress")
}

func CheckVulnerabilities(url string) map[string]string {
	vulnerabilities := make(map[string]string)

	// Example check for SQL injection vulnerability
	resp, err := http.Get(url + "'")
	if err == nil && resp.StatusCode == http.StatusInternalServerError {
		vulnerabilities["SQL Injection"] = "Possible SQL injection vulnerability detected."
	} else {
		vulnerabilities["SQL Injection"] = "No SQL injection vulnerability detected."
	}
	// Add more vulnerability checks as needed

	return vulnerabilities
}

func CheckWordPressHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	url := r.FormValue("site-url")
	err := IsWordPressSite(url)
	response := make(map[string]string)

	if err != nil {
		response["result"] = "Le site n'est pas un site WordPress"
	} else {
		response["result"] = "Le site est un site WordPress"
	}

	vulnerabilities := CheckVulnerabilities(url)
	for key, value := range vulnerabilities {
		response[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DownloadPDFHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("site-url")
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Erreur lors de la requête HTTP", http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Erreur lors de la récupération du contenu du site", http.StatusBadRequest)
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Erreur lors de l'analyse du document", http.StatusInternalServerError)
		return
	}

	if doc.Find("meta[name='generator'][content^='WordPress']").Length() == 0 {
		http.Error(w, "Le site n'est pas un site WordPress", http.StatusBadRequest)
		return
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Rapport de Sécurité WordPress")
	pdf.Ln(20)
	pdf.SetFont("Arial", "", 12)

	// Extract text and images from the site
	doc.Find("body").Each(func(i int, s *goquery.Selection) {
		s.Find("p, h1, h2, h3, h4, h5, h6").Each(func(i int, s *goquery.Selection) {
			text := s.Text()
			pdf.MultiCell(0, 10, text, "", "", false)
		})
		s.Find("img").Each(func(i int, s *goquery.Selection) {
			imgSrc, _ := s.Attr("src")
			if imgSrc != "" {
				if !strings.HasPrefix(imgSrc, "http") {
					imgSrc = url + imgSrc
				}
				resp, err := http.Get(imgSrc)
				if err == nil {
					defer resp.Body.Close()
					imgData, err := ioutil.ReadAll(resp.Body)
					if err == nil {
						tmpFile, err := ioutil.TempFile("", "img-*.jpg")
						if err == nil {
							defer os.Remove(tmpFile.Name())
							tmpFile.Write(imgData)
							tmpFile.Close()
							pdf.Ln(10)
							pdf.Image(tmpFile.Name(), 10, pdf.GetY(), 0, 0, false, "", 0, "")
							pdf.Ln(10)
						}
					}
				}
			}
		})
	})

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=rapport.pdf")
	err = pdf.Output(w)
	if err != nil {
		http.Error(w, "Erreur lors de la génération du PDF", http.StatusInternalServerError)
	}
}
