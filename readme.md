💻 Install
===============
Dans le répertoire ```/2eme-Partie```, ouvrez un terminal et exécuter ces 2 commandes :

```bash
go build

./project-particles
```

💡Fonctionalités
===============
Initialisation d’un système de particules configurable.

Gestion d’un spawn aléatoire ou fixe.

Déplacement individuel de chaque particule.

Rotation initiale aléatoire.

Gestion d’un accumulateur de spawn permettant des vitesses d’apparition fractionnaires.

Génération de vitesses aléatoires (dans une plage configurable).

Tests unitaires complets (positions, vitesses, spawn, accumulateur…).

🚀Usage 
========

1. Modifier le config.json

Vous pouvez ajuster les paramètres principaux :

```
"ParticleImage"    : chemin vers le sprite

"InitNumParticles" : nombre initial de particules

"RandomSpawn"      : true / false

"SpawnRate"        : nombre de particules par update (peut être décimal)

"SpeedFactor"      : vitesse maximale aléatoire

"SpawnX" / "SpawnY": position fixe si RandomSpawn = false

"ScaleX" / "ScaleY": échelles
```
2. Vous pouvez changer ces paramètres dynamiquement en utilisant les touches de votre clavier.

🛠️Dev
=========

```NewSystem()```

Fonction d’initialisation du système de particules.

Rôle :

```

    Initialise une liste de particules (container/list).

    Calcule la zone de spawn (min/max) en fonction de la marge.

    Gère le spawn aléatoire ou fixe.

    Initialise pour chaque particule :

        -Position

        -Rotation aléatoire

        -Échelle

        -Couleur / Opacité

        -Vitesse aléatoire dans la plage [-SpeedFactor, +SpeedFactor].

    Renvoie un System contenant :

        -Content → liste de particules

        -SpawnAccumulator → compteur pour les spawns fractionnaires

```

```Update()```

Fonction principale appelée à chaque frame / tick.


Rôle :

```

1. Gestion du spawn

    Ajoute SpawnRate à SpawnAccumulator.

    Tant qu’il dépasse 1 → spawn d’une particule.

    Supporte SpawnRate fractionnaire (0.2, 0.7…) grâce à l’accumulateur.

2. Création d'une particule lors du spawn

    Même logique que NewSystem() :

        Position aléatoire (ou fixe)

        Rotation aléatoire

        Vitesse aléatoire

    Ajout en tête de liste.

3. Mise à jour des particules existantes
	
	Ajout de la vitesse aà chaque particules

    Note : Aucune vérification de sortie d'écran n'est encore effectuée, ce qui peut être ajouté dans un futur développement.
```

ℹ️ Extensions
=========

```Gravity()```

Gravité: désactivée/activée Force : float ('G' ON/OFF 'T' + 'B' -)

```Fashion()```

Color : désactivée/activée -> Fire/Océan/Forest/Chromatic ('H' ON/OFF 'Y' Changement de couleur)

```Lifespan()```

Durée de vie : désactivée/activée et int ('F' ON/OFF 'R' + 'V' -)

```Outofscreen()```

Sortie d'écran : désactivée/activée ('J' ON/OFF)

```Collision()```

Collision : désactivée/activée ('A' ON/OFF)

```Modesouris()```

Mode Souris : désactivée/activée ('S' ON/OFF)
