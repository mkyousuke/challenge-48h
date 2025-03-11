// fonctions/pdf.go
package fonctions

import (
	"fmt"
	"net/http"
	"time"
	"github.com/jung-kurt/gofpdf"
)

func GeneratePDFHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL manquante", http.StatusBadRequest)
		return
	}

	// Génération du nom de fichier unique
	filename := fmt.Sprintf("rapport_%s.pdf", time.Now().Format("20060102_150405"))
	
	// Création du PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	
	// Contenu du rapport (à personnaliser)
	pdf.Cell(40, 10, "Rapport de Sécurité WordPress")
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(0, 10, fmt.Sprintf("Rapport généré pour: %s\n\nDate: %s\n\n[CONTENU DU RAPPORT À AJOUTER]", url, time.Now().Format("02/01/2006 15:04")), "", "L", false)

	// Envoi du PDF au client
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", filename))
	
	if err := pdf.Output(w); err != nil {
		http.Error(w, "Erreur de génération PDF", http.StatusInternalServerError)
	}
}
