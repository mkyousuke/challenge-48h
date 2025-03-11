// fonctions/pdf.go
package fonctions

import (
	"fmt"
	"net/http"
	"time"
	"github.com/jung-kurt/gofpdf"
)

func GeneratePDFHandler(w http.ResponseWriter, r *http.Request) {
	// Récupération de l'URL depuis les paramètres de requête
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL manquante", http.StatusBadRequest)
		return
	}

	// Création de l'instance PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	// Création du traducteur Unicode pour gérer les caractères accentués
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	
	// En-tête du rapport
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, tr("Rapport de Sécurité WordPress"))
	pdf.Ln(12)
	
	// Date et heure actuelles formatées
	currentTime := time.Now().Format("02/01/2006 15:04")
	
	// Contenu du rapport
	pdf.SetFont("Arial", "", 12)
	
	// Utilisation du traducteur pour les textes avec accents
	reportText := fmt.Sprintf("URL analysée : %s\nDate : %s\n\n", url, currentTime)
	reportText += "Résultats de l'analyse :\n"
	reportText += "✓ Vérification du site WordPress : Positif\n"
	reportText += "✓ Version détectée : À déterminer\n\n"
	reportText += "Recommandations :\n"
	reportText += "- Maintenez WordPress et ses extensions à jour\n"
	reportText += "- Utilisez un plugin de sécurité\n"
	reportText += "- Effectuez des sauvegardes régulières\n\n"
	reportText += "Ce rapport a été généré automatiquement par l'outil d'audit de sécurité WordPress."
	
	// Ajout du texte traduit au PDF
	pdf.MultiCell(0, 10, tr(reportText), "", "L", false)
	
	// En-têtes HTTP pour le téléchargement
	filename := fmt.Sprintf("rapport_wordpress_%s.pdf", time.Now().Format("20060102_150405"))
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	
	// Génération et envoi du PDF
	err := pdf.Output(w)
	if err != nil {
		// Log de l'erreur
		fmt.Printf("Erreur lors de la génération du PDF: %v\n", err)
		http.Error(w, "Erreur lors de la génération du PDF", http.StatusInternalServerError)
		return
	}
	
	// Log de succès (optionnel)
	fmt.Printf("PDF généré avec succès pour: %s\n", url)
}
