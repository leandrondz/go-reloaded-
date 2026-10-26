go-reloaded

Ce projet consiste à créer un programme en Go qui lit un fichier texte le modifie selon certaines règles et enregistre le résultat dans un autre fichier.

Fonctions principales

- main() : lance le programme.
- os.Args : récupère les arguments donnés dans la commande.
- os.ReadFile() : lit le fichier d'entrée.
- Transform() : transforme le texte selon les règles demandées.
- os.WriteFile() : écrit le résultat dans le fichier de sortie.
- fmt.Println() : affiche des messages dans le terminal.

Transformations

- (up) : met le mot précédent en majuscules.
- (low) : met le mot précédent en minuscules.
- (cap) : met une majuscule au début du mot précédent.
- (up N) : met les N mots précédents en majuscules.
- (low N) : met les N mots précédents en minuscules.
- (cap N) : met une majuscule au début des N mots précédents.
- (hex) : convertit un nombre hexadécimal en décimal.
- (bin) : convertit un nombre binaire en décimal.

Le programme corrige aussi les espaces autour de la ponctuation gère certaines apostrophes et remplace a par an devant une voyelle.

Utilisation

Pour lancer le programme utiliser la commande suivante :

go run . input.txt output.txt

