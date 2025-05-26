package main

import (
	"flag"
	"fmt"

	"github.com/NicolasBIDET/TP1-Annuaire/annuaire"
)

func main() {
	var action string
	var nom string
	var tel string

	flag.StringVar(&action, "action", "ajouter", "Action à effectuer (ajouter, modifier, supprimer, lister, rechercher)")
	flag.StringVar(&nom, "nom", "", "Nom du contact")
	flag.StringVar(&tel, "tel", "", "Numéro de téléphone du contact")
	flag.Parse()

	switch action {
	case "ajouter":
		annuaire.AjouterContact(nom, tel)
	case "modifier":
		annuaire.MettreAJourContact(nom, tel)
	case "supprimer":
		annuaire.SupprimerContact(nom)
	case "lister":
		annuaire.ListerContacts()
	case "rechercher":
		annuaire.RechercherContact(nom)
	default:
		fmt.Println("Action non reconnue")
	}
}
