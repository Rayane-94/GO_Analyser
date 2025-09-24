# 🚀 GoLog Analyzer

Analyseur de logs en parallèle avec Go et goroutines.  
Outil CLI qui lit un fichier de configuration JSON, lance l’analyse de plusieurs logs simultanément et peut exporter les résultats en JSON.

---

## ⚙️ Installation

```bash
git clone https://github.com/Rayane-94/GO_Analyser.git
cd GO_Analyser/loganizer
go mod tidy

```
## Commande disponible

<br> go run main.go analyze -c config.json 
<br> go run main.go analyze -c config.json -o rapport.json
<br> go run main.go --help
<br> go run main.go analyze --help

<br>go build -o loganalyzer.exe
<br>./loganalyzer.exe analyze -c config.json
<br>./loganalyzer.exe analyze -c config.json -o rapport.json

<br>go run main.go --help
<br>go run main.go analyze --help
<br>go run main.go analyze -c config.json
<br>go run main.go analyze -c config.json -o rapport.json
<br>./loganalyzer.exe analyze -c config.json
<br>./loganalyzer.exe analyze -c config.json -o rapport.json

