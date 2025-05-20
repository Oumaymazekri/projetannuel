# Utiliser l'image officielle de Go
FROM golang:1.23.4 AS build

# Définir le répertoire de travail
WORKDIR /app

# Copier les fichiers Go et go.mod
COPY go.mod go.sum ./

# Exécuter go mod tidy pour résoudre les dépendances
RUN go mod tidy

# Copier le reste du code source
COPY . .

# Compiler le service
RUN go build -o product-service

# Exposer le port du service
EXPOSE 3001

# Démarrer le service
CMD ["go", "run", "main.go"]

