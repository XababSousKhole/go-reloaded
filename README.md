# go-reloaded

Petit outil en Go qui complète, édite et corrige un texte.

## Usage

```
go run . sample.txt result.txt
```

Le premier argument est le fichier d'entrée, le second le fichier de sortie.

## Règles

- `(hex)` : le mot précédent (hexadécimal) devient décimal
- `(bin)` : le mot précédent (binaire) devient décimal
- `(up)`, `(low)`, `(cap)` : majuscules, minuscules, majuscule initiale
- `(up, N)`, `(low, N)`, `(cap, N)` : pareil sur les N mots précédents
- ponctuation `. , ! ? : ;` collée au mot précédent, espace après
- groupes `...` et `!?` gardés tels quels
- `' '` collés au texte qu'ils entourent
- `a` devient `an` devant une voyelle ou un `h`

## Fichiers

- `main.go` : lecture et écriture des fichiers
- `process.go` : enchaîne les transformations
- `tags.go`, `case.go` : balises (hex), (bin), (up), (low), (cap)
- `punct.go` : ponctuation
- `quotes.go` : apostrophes
- `articles.go` : a / an
- `process_test.go` : tests (`go test ./...`)
