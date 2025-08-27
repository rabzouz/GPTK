# GPTK - Assistant IA pour Kali Linux
## Description : Bot basé sur OpenAI pour codage et outils Kali.
# README pour GPTK
# README pour GPTK

Bienvenue dans le dépôt GitHub de **GPTK** ! Ce projet est un assistant IA personnalisé, basé sur l'API OpenAI, conçu spécifiquement pour Kali Linux. Il aide au codage, répond à des questions bien formulées, simule des recherches approfondies (via prompts intelligents), génère des images (via DALL·E), et est optimisé pour les outils de pentesting Kali (comme Nmap, Metasploit, etc.). Développé en Go pour une portabilité maximale entre Windows et Linux, il inclut une persistance des sessions, une interface web optionnelle, un export CSV, et des modules extensibles.

Ce README est détaillé pour vous guider de A à Z : installation, configuration, usage, et contribution. Si vous avez des questions, ouvrez une issue !

## Description
GPTK (GPT for Kali) est un bot CLI (ligne de commande) qui utilise l'API OpenAI pour :
- Répondre en français (forcé via un prompt système).
- Assister au codage (ex. : générer des scripts Go, Kotlin, ou Python).
- Expliquer et générer des commandes pour outils Kali (via modules dédiés).
- Sauvegarder l'historique des sessions dans un fichier JSON persistant avec timestamps.
- Générer des images visuelles (ex. : diagrammes de réseaux pour Nmap).
- Exporter l'historique en CSV pour analyse.
- Offrir une interface web simple (optionnelle) pour un usage via navigateur.

Le projet est open-source, facile à étendre, et adapté à vos intérêts en développement Android, pentesting Kali, et IA.

## Fonctionnalités Principales
- **Réponses IA** : Basées sur GPT-4o, toujours en français, claires et concises, avec exemples de commandes pour Kali.
- **Historique Persistant** : Sauvegarde automatique des questions/réponses dans `history.json` avec timestamp (format ISO).
- **Flags CLI** :
  - `--show-history` : Affiche l'historique formaté (numéroté, avec timestamp).
  - `--clear-history` : Efface le fichier d'historique.
  - `--export-csv` : Exporte l'historique en CSV (history.csv).
  - `--image "prompt"` : Génère et sauvegarde une image via DALL·E (ex. : generated_image.png).
- **Modules pour Kali** : Explications et génération de commandes pour Nmap, Metasploit, etc. (extensible via dossier `modules/`).
- **Interface Web Optionnelle** : Lancez avec `--web` pour un chat via navigateur (localhost:8080).
- **Portabilité** : Compile sur Windows et Kali ; persistance locale.
- **Sécurité** : Clé API via variable d'environnement (jamais stockée dans le code).

## Prérequis
- **Clé API OpenAI** : Gratuite pour un usage basique. Générez-la sur [platform.openai.com](https://platform.openai.com/account/api-keys) avec votre email (ex. : rabzouz6481@gmail.com).
- **Go** : Version 1.23+ (installé sur Windows et Kali).
- **Git** : Pour cloner le dépôt.
- **Accès Internet** : Pour les appels API et génération d'images.

## Installation
### Sur Windows (PowerShell)
1. Installez Go : Téléchargez depuis [go.dev/dl](https://go.dev/dl) (ex. : go1.23.0.windows-amd64.msi) et installez.
2. Clonez le dépôt :
   ```
   git clone https://github.com/rabzouz/GPTK.git
   cd GPTK
   ```
3. Résolvez les dépendances :
   ```
   go mod tidy
   ```
4. Compilez :
   ```
   go build -o gptk.exe
   ```
5. Configurez la clé API :
   ```
   setx OPENAI_API_KEY "sk-proj-votre-clé"
   ```
   Rouvrez PowerShell.

### Sur Kali Linux (Terminal)
1. Installez Go et Git (si pas déjà fait) :
   ```
   sudo apt update
   sudo apt install golang-go git -y
   ```
2. Clonez le dépôt :
   ```
   mkdir ~/projects
   cd ~/projects
   git clone https://github.com/rabzouz/GPTK.git
   cd GPTK
   ```
3. Résolvez les dépendances :
   ```
   go mod tidy
   ```
4. Compilez :
   ```
   go build -o gptk
   ```
5. Configurez la clé API (persistante) :
   ```
   nano ~/.bashrc
   ```
   Ajoutez : `export OPENAI_API_KEY="sk-proj-votre-clé"`. Sauvegardez et :
   ```
   source ~/.bashrc
   ```
6. Rendez global (optionnel) :
   ```
   sudo mv gptk /usr/local/bin/
   ```

## Usage
Lancez l'exécutable :
- Windows : `.\gptk.exe`
- Kali : `./gptk` (ou `gptk` si global)

Exemple de session :
```
Historique chargé : 0 entrées
GPTK prêt – posez vos questions (Ctrl+C pour quitter).
Explique Nmap -sS
--- Réponse GPTK ---
[Réponse en français sur Nmap -sS avec exemples]
>>>
```

- **Flags** :
  - `--show-history` : Liste l'historique (ex. : "1) [2025-08-27TXX:XX:XXZ] Question ⇒ Réponse").
  - `--clear-history` : Efface `history.json`.
  - `--export-csv` : Exporte en history.csv.
  - `--image "prompt"` : Génère une image (ex. : generated_image.png).

Pour des recherches : Posez comme "Recherche approfondie sur Metasploit exploits" – l'IA simule via son entraînement. Pour images : `--image "Diagramme Nmap"`.

## Configuration Avancée
- **Clé API** : Si expirée, régénérez sur OpenAI et mettez à jour la variable.
- **Étendre les Modules** : Ajoutez des fichiers dans `modules/` (ex. : `kali.go` pour Nmap) et importez-les dans `main.go`.
- **Historique** : Stocké dans `history.json` (ajoutez à .gitignore pour ne pas le pusher).
- **Limites** : Ajoutez une limite (ex. : 100 entrées) pour éviter un fichier trop gros.

## Contribution
1. Forkez le dépôt.
2. Créez une branche : `git checkout -b feature/nouveau-module`.
3. Committez : `git commit -m "Ajout module Nmap"`.
4. Poussez : `git push origin feature/nouveau-module`.
5. Ouvrez une Pull Request.

Idées : Ajoutez des modules pour Wireshark, Burp Suite, ou une UI web plus avancée.

## Licence
MIT License – Voir le fichier [LICENSE](LICENSE) pour détails. Utilisez librement, mais citez la source.

Merci d'utiliser GPTK ! Si bugs ou idées, ouvrez une issue sur GitHub. 😊
