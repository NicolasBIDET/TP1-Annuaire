package main

import (
	"flag"
	"fmt"
	"log"

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
		str, err := annuaire.AjouterContact(nom, tel)
		if err != nil {
			log.Fatalf("Erreur : %v\n", err)
		} else {
			fmt.Println(str)
		}
	case "modifier":
		str, err := annuaire.MettreAJourContact(nom, tel)
		if err != nil {
			log.Fatalf("Erreur : %v\n", err)
		} else {
			fmt.Println(str)
		}
	case "supprimer":
		str, err := annuaire.SupprimerContact(nom)
		if err != nil {
			log.Fatalf("Erreur : %v\n", err)
		} else {
			fmt.Println(str)
		}
	case "lister":
		str := annuaire.ListerContacts()
		fmt.Println(str)
	case "rechercher":
		str, err := annuaire.RechercherContact(nom)
		if err != nil {
			log.Fatalf("Erreur : %v\n", err)
		} else {
			fmt.Println(str)
		}
	default:
		fmt.Println("Action non reconnue")
	}
}
