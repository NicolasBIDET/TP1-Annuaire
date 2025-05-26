# TP1-Annuaire

Un projet Go simple pour gérer un annuaire (address book).

## Utilisation

Pour lancer le projet, utilisez la commande suivante :

```bash
go run main.go [flags]
```

### Flags disponibles

- `--action [ajouter|modifier|supprimer|lister|rechercher` : choisir l'action a effectuer parmis les suivantes.
- `--nom <nom>` : Nom du contact. Requis pour les actions : ajouter, modifier, supprimer et  rechercher
- `--tel <tel>` : Numéro de telephone du contact. Requis pour les actions : ajouter et modifier

Exemple :

```bash
go run main.go --action ajouter --nom "John" --tel "0123456789"
go run main.go --action modifier --nom "John" --tel "0123456780"
go run main.go --action supprimer --nom "John"
go run main.go --action rechercher --nom "John"
go run main.go --action lister
```

## Tests

Pour lancer les tests, utilisez la commande suivante :

```bash
go test annuaire
```

## Membres du groupe

- Nicolas BIDET
- Timothé BONIFACIO

