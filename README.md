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

go run main.go analyze -c config.json
go run main.go analyze -c config.json -o rapport.json
go run main.go --help
go run main.go analyze --help

go build -o loganalyzer.exe
./loganalyzer.exe analyze -c config.json
./loganalyzer.exe analyze -c config.json -o rapport.json

go run main.go --help
go run main.go analyze --help
go run main.go analyze -c config.json
go run main.go analyze -c config.json -o rapport.json
./loganalyzer.exe analyze -c config.json
./loganalyzer.exe analyze -c config.json -o rapport.json

