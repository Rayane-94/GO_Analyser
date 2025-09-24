# 🚀 GoLog Analyzer# 🚀 GoLog Analyzer# TP : GoLog Analyzer - Analyse de Logs Distribuée



Outil CLI en Go pour analyser des fichiers de logs en parallèle avec goroutines.



## 📁 ArchitectureUn outil CLI en Go pour analyser des fichiers de logs en parallèle.### Contexte



```

loganizer/

├── main.go              # Point d'entrée## 📋 DescriptionVotre équipe est chargée de développer un outil en ligne de commande (CLI) en Go, nommé `loganalyzer`. Son but est d'aider les administrateurs système à analyser des fichiers de logs (journaux) provenant de diverses sources (serveurs, applications). L'objectif est de pouvoir centraliser l'analyse de multiples logs en parallèle et d'en extraire des informations clés, tout en gérant les erreurs potentielles de manière robuste.

├── config.json          # Configuration d'exemple

├── cmd/                 # Commandes CLI avec Cobra

├── internal/config/     # Lecture JSON

├── internal/analyzer/   # Logique d'analyse + erreurs personnaliséesGoLog Analyzer est un programme en ligne de commande qui permet d'analyser plusieurs fichiers de logs simultanément en utilisant les goroutines de Go. Il lit une configuration JSON, traite chaque log en parallèle, et génère un rapport détaillé.### Objectifs d'apprentissage

├── internal/reporter/   # Export JSON

└── test_logs/          # Fichiers de test

```

## 🏗️ Structure du ProjetCe TP vous permettra de renforcer vos compétences sur les concepts suivants :

## 🚀 Utilisation



```bash

# Cloner et installer```- **Concurrence :** Utiliser les **goroutines** et les **WaitGroups** pour traiter plusieurs tâches en parallèle.

git clone https://github.com/Rayane-94/GO_Analyser.git

cd GO_Analyser/loganizerloganizer/- **Gestion des Erreurs :** Implémenter des **erreurs personnalisées** et les gérer proprement avec `errors.Is` et `errors.As`.

go mod tidy

├── main.go                    # Point d'entrée du programme- **Outil CLI avec Cobra :** Structurer une application en ligne de commande avec des **sous-commandes** et des **drapeaux (flags)**.

# Analyser (console)

go run main.go analyze -c config.json├── config.json               # Exemple de configuration- **Import/Export JSON :** Manipuler des données au format JSON pour la configuration d'entrée et le rapport de sortie.



# Analyser avec export JSON├── go.mod / go.sum           # Gestion des dépendances Go- **Modularité :** Organiser le code en **packages** logiques (`internal/`).

go run main.go analyze -c config.json -o rapport.json

├── cmd/                      # Commandes CLI (Cobra)

# Aide

go run main.go --help│   ├── root.go              # Commande racine---

go run main.go analyze --help

│   └── analyze.go           # Commande d'analyse principale

# Compiler (optionnel)

go build -o loganalyzer.exe├── internal/                 # Code privé de l'application### Cahier des charges

./loganalyzer.exe analyze -c config.json

```│   ├── config/              # Lecture des configurations JSON



**Développé par Rayane-94**│   │   └── config.goVotre outil `loganalyzer` devra implémenter les fonctionnalités suivantes :

│   ├── analyzer/            # Logique d'analyse et erreurs personnalisées

│   │   └── analyzer.go#### 1. Commande `analyze`

│   └── reporter/            # Export des rapports JSON

│       └── reporter.go- **Entrée JSON :** La commande prendra un chemin vers un **fichier de configuration JSON** via un drapeau `--config <path>` (raccourci `-c`). Ce fichier contiendra la liste des logs à analyser.

└── test_logs/               # Fichiers de logs d'exemple

    ├── access.log  **Exemple de fichier `config.json` :**

    ├── errors.log    ```json

    ├── corrupted.log    [

    └── empty.log      {

```        "id": "web-server-1",

        "path": "/var/log/nginx/access.log",

## 🚀 Installation et Utilisation        "type": "nginx-access"

      },

### Prérequis      {

- Go 1.24+ installé        "id": "app-backend-2",

        "path": "/var/log/my_app/errors.log",

### Installation        "type": "custom-app"

```bash      }

# Cloner le projet    ]

git clone https://github.com/Rayane-94/GO_Analyser.git    ```

cd GO_Analyser/loganizer  - `id` : Un identifiant unique pour le log.

  - `path` : Le chemin (absolu ou relatif) vers le fichier de log.

# Installer les dépendances  - `type` : Le type de log (peut être ignoré mais doit être présent).

go mod tidy

- **Traitement concurrentiel :** Une **goroutine** sera lancée pour chaque log :

# Compiler (optionnel)  - Vérifier si le fichier existe et est lisible.

go build -o loganalyzer.exe  - Simuler l'analyse avec un `time.Sleep` aléatoire (50 à 200 ms).

```



### Utilisation- **Collecte et Exportation des résultats :**

  - Résultats collectés via un **canal sécurisé**.

#### 1. Analyser avec affichage console uniquement  - Export possible via `--output <path>` (raccourci `-o`) dans un fichier JSON.

```bash

go run main.go analyze -c config.json    **Exemple de fichier `report.json` :**

```    ```json

    [

#### 2. Analyser avec export JSON      {

```bash        "log_id": "web-server-1",

go run main.go analyze -c config.json -o rapport.json        "file_path": "/var/log/nginx/access.log",

```        "status": "OK",

        "message": "Analyse terminée avec succès.",

#### 3. Avec le binaire compilé        "error_details": ""

```bash      },

./loganalyzer.exe analyze --config config.json --output rapport.json      {

```        "log_id": "invalid-path",

        "file_path": "/non/existent/log.log",

#### 4. Voir l'aide        "status": "FAILED",

```bash        "message": "Fichier introuvable.",

go run main.go --help        "error_details": "open /non/existent/log.log: no such file or directory"

go run main.go analyze --help      }

```    ]

    ```

## 📄 Format de Configuration

- **Affichage sur console :** Un résumé doit être affiché pour chaque log : ID, chemin, statut, message, erreur (si applicable).

Le fichier de configuration doit être un JSON contenant un tableau d'objets :

#### 2. Gestion des Erreurs Personnalisées

```json

[- Implémenter au moins **deux types d'erreurs personnalisées** :

  {  - Fichier introuvable/inaccessible.

    "id": "web-server-1",  - Erreur de parsing.

    "path": "test_logs/access.log",- Utiliser `errors.Is()` et/ou `errors.As()` pour les gérer proprement.

    "type": "nginx-access"

  },---

  {

    "id": "app-backend-2",### Architecture suggérée (packages `internal/`)

    "path": "test_logs/errors.log", 

    "type": "custom-app"Organisez le projet comme suit :

  }

]- `internal/config` : Lecture des configurations JSON.

```- `internal/analyzer` : Analyse, erreurs personnalisées, rapport.

- `internal/reporter` : Export JSON des résultats.

## 📊 Format du Rapport- `cmd/` :

  - `root.go` : Commande racine.

Le rapport généré contient :  - `analyze.go` : Commande `analyze`.



```json---

[

  {### Critères d'évaluation

    "log_id": "web-server-1",

    "file_path": "test_logs/access.log",L’évaluation portera sur :

    "status": "OK",

    "message": "Analyse terminée avec succès",- **Fonctionnalité :** La commande `analyze` fonctionne-t-elle comme spécifié ?

    "error_details": ""- **Concurrence :** Traitement en parallèle via `goroutines` et `WaitGroup` ? Résultats collectés via `channel` ?

  },- **Gestion des Erreurs :** Utilisation et gestion correcte des erreurs personnalisées ? Messages d’erreur clairs ?

  {- **CLI :** Interface Cobra fonctionnelle, avec drapeaux et descriptions ?

    "log_id": "invalid-log",- **JSON :** Import/export respectant les structures attendues ?

    "file_path": "/inexistant.log",- **Modularité :** Code organisé proprement en packages ?

    "status": "FAILED",- **Documentation :** Je veux voir **un beau readme** qui explique le fonctionnement de votre programme, vos commandes, et j'en passe ET **la documentation de votre code** et **les membres de votre team**.

    "message": "Fichier introuvable",

    "error_details": "fichier introuvable: /inexistant.log"### Type de rendu

  }

]- Un lien github

```



## ⚙️ Fonctionnalités Principales### 🎁 BONUS



- ✅ **Traitement parallèle** avec goroutinesVous avez l'âme d'un.e développeur.euse courageux.euse ? Je vous laisse ici quelques bonus si vous voulez vous amuser un peu et avoir un programme plus complet.

- ✅ **Interface CLI** professionnelle avec Cobra

- ✅ **Gestion d'erreurs** robuste avec erreurs personnalisées  **1. Gestion des dossiers d'exportation **

- ✅ **Export JSON** formaté* Si le chemin de sortie JSON (`--output`) inclut des répertoires qui n'existent pas (ex: `rapports/2024/mon_rapport.json`), faire en sorte que le programme crée automatiquement ces répertoires avant d'écrire le fichier.

- ✅ **Architecture modulaire** avec packages séparés* **Indice** : `os.MkdirAll(filepath.Dir(path), 0755)`

- ✅ **Configuration flexible** via JSON* **Intérêt** : Rend l'outil plus robuste et convivial.



## 🎯 Exemples d'Utilisation**2. Horodatage des Exports JSON**

* Nommer les fichiers de sortie JSON avec une date :

```bash  * **Modifier la commande `analyze`** pour que, si le drapeau `--output` est fourni, le nom du fichier de sortie JSON inclue la date du jour au format AAMMJJ (AnnéeMoisJour).

# Test avec les fichiers d'exemple  * **Exemple** : au lieu de `report.json`, le fichier serait nommé `240524_report.json` (pour le 24 mai 2024).

go run main.go analyze -c config.json  * **Indice** : Utiliser le package `time` de Go (`time.Now()`, `time.Format()`).

  * **Intérêt** : Ajoute une fonctionnalité pratique pour l'organisation des rapports, et force à manipuler les dates en Go.

# Export du rapport

go run main.go analyze -c config.json -o mon_rapport.json**2. Commande `add-log`**

* **Ajouter une nouvelle sous-commande add-log** qui permettrait d'ajouter manuellement une configuration de log au fichier config.json existant.

# Test de gestion d'erreur avec fichier malformé* **Drapeaux nécessaires** : `--id`, `--path`, `--type`, `--file` (chemin du fichier `config.json`).

go run main.go analyze -c bad_config.json

```**3. Filtrage des résultats d'analyse**

* **Ajouter un drapeau `--status <status>`** (ex: `--status FAILED` ou `--status OK`) à la commande analyze pour n'afficher et/ou n'exporter que les logs ayant un certain statut.

## 👥 Équipe* **Intérêt** : Ajoute une fonctionnalité utile et demande de la logique de filtrage avant l'affichage/l'export.



- **Développeur** : Rayane-94

- **Repository** : [GO_Analyser](https://github.com/Rayane-94/GO_Analyser)---



---### Pour démarrer (Prérequis)



*Développé en Go avec ❤️*1. Créer un module : `go mod init`
2. Installer Cobra : `go get github.com/spf13/cobra@latest`
3. Avoir bien lu le readme ;)

---

Bon courage !
