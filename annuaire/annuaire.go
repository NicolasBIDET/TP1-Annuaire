package annuaire

import (
	"fmt"
)

var contacts = map[string]string{
	"Charlie": "0811223344",
	"Dave":    "0811333445",
	"Eve":     "0811445566",
}

func AjouterContact(nom string, tel string) (string, error) {
	if nom == "" || tel == "" {
		return "", fmt.Errorf("nom et numéro de téléphone requis")
	}
	if _, exists := contacts[nom]; exists {
		return "", fmt.Errorf("le contact %s existe déjà", nom)
	}
	contacts[nom] = tel
	return fmt.Sprintf("Contact %s ajouté avec succès", nom), nil
}

func RechercherContact(nom string) (string, error) {
	if nom == "" {
		return "", fmt.Errorf("nom requis pour rechercher un contact")
	}
	if tel, exists := contacts[nom]; exists {
		return fmt.Sprintf("Contact %s trouvé : %s", nom, tel), nil
	}
	return "", fmt.Errorf("contact %s non trouvé", nom)
}

func ListerContacts() string {
	if len(contacts) == 0 {
		return "Aucun contact trouvé"
	}
	result := "Liste des contacts :\n"
	for nom, tel := range contacts {
		result += fmt.Sprintf("%s : %s\n", nom, tel)
	}
	return result
}

func MettreAJourContact(nom string, tel string) (string, error) {
	if nom == "" || tel == "" {
		return "", fmt.Errorf("nom et numéro de téléphone requis pour mettre à jour un contact")
	}
	if _, exists := contacts[nom]; !exists {
		return "", fmt.Errorf("le contact %s n'existe pas", nom)
	}
	contacts[nom] = tel
	return fmt.Sprintf("Contact %s mis à jour avec succès", nom), nil
}

func SupprimerContact(nom string) (string, error) {
	if nom == "" {
		return "", fmt.Errorf("nom requis pour supprimer un contact")
	}
	if _, exists := contacts[nom]; !exists {
		return "", fmt.Errorf("le contact %s n'existe pas", nom)
	}
	delete(contacts, nom)
	return fmt.Sprintf("Contact %s supprimé avec succès", nom), nil
}

// ResetContacts permet de réinitialiser la liste des contacts pour les tests
func ResetContacts() {
	contacts = map[string]string{
		"Charlie": "0811223344",
		"Dave":    "0811333445",
		"Eve":     "0811445566",
	}
}
func EmptyContacts() {
	contacts = make(map[string]string)
}
