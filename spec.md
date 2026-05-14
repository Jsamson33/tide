# Prompt: Project Tide - Go CLI

**Objectif :** Créer un outil CLI nommé `tide` qui génère des analyses basées sur une logique binaire de 6 états (hexagrammes).

**Mécanique système :**
- L'outil doit simuler 6 tirages binaires équiprobables (50/50) pour construire une pile de 6 lignes (pleines ou brisées). Il n'y a pas de gestion des lignes mutantes.
- Le tirage du jour doit être déterministe sur la journée.
- Le tirage "j'ai de la chance" est aléatoire à chaque lancement de l'application.
- À la fin de l'affichage du tirage, on quitte l'app.

**Interface CLI :**
- Commande principale : `tide` (déclenche immédiatement le tirage du jour, déterministe).
- Commande interactive : `tide ask` (déclenche immédiatement un tirage aléatoire "j'ai de la chance").
- Sortie : Représentation ASCII minimaliste des lignes (pas de connotation asiatique ou chinoise).
- Flag `--json` : Pour exporter le résultat du tirage (états des 6 lignes) vers un autre programme.

**Livrables :** Code Go modulaire, propre et factuel.
j'aimerai que le code respecte une archi hexagonale.

J'aimerai avoir un répertoire (à nommer) contenant les 8 trigrammes et 64 hexagrammes et autres fichiers si besoin, au format markdown.
Ces fichiers seront "lus" par l'app au moment du tirage.

Il faudra générer tous ces fichiers, sans faire de plagiat et en respectant au maximum le droit international sur le yijing. Le contenu devra être généré par IA de manière neutre, factuelle et inspirante pour éviter tout problème de droits d'auteur (libre de droits).

Il faudra un script pour la création du binaire et son installation dans le PATH (macOS / Linux).