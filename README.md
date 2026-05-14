# Tide

Tide est un outil CLI en Go conçu pour générer des analyses basées sur une logique binaire de 6 états (hexagrammes). Il simule 6 tirages binaires équiprobables pour construire une pile de 6 lignes, sans gestion de lignes mutantes.

## Architecture

Le projet suit une **Architecture Hexagonale** (Ports et Adapteurs) pour garantir la modularité et la séparation des préoccupations :
- **Domaine** : Logique de tirage et structures de données.
- **Adaptateurs** : 
  - `file` : Gestion du contenu via des fichiers Markdown embarqués.
  - `cli` : Interface utilisateur en ligne de commande.

## Installation

### Prérequis
- Go 1.25+
- Make (optionnel)

### Installation rapide
```bash
make build
sudo make install
```

Ou manuellement :
```bash
go build -o tide cmd/tide/main.go
sudo mv tide /usr/local/bin/
```

## Utilisation

### Tirage du jour
Le tirage du jour est déterministe. Pour une date donnée, le résultat sera toujours le même.
```bash
tide
```

### Mode interactif (Lucky Draw)
Pour effectuer un tirage aléatoire immédiat :
```bash
tide ask
```

### Export JSON
Pour intégrer Tide avec d'autres outils :
```bash
tide --json
```

## Développement

### Générer les données
Pour régénérer les fichiers de données (64 hexagrammes et 8 trigrammes) :
```bash
make generate
```

### Tests
```bash
make test
```

## Licence
LICENSE libre de droits.
