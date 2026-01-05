# ===============================
# 1️⃣ BUILD STAGE
# ===============================
FROM golang:1.23-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

# Copier les fichiers go.mod pour le cache
COPY go.mod go.sum ./
RUN go mod download

# Copier le reste du code
COPY . .

# Build optimisé (binaire léger)
RUN go build -ldflags="-s -w" -o product-service main.go


# ===============================
# 2️⃣ RUNTIME STAGE (ULTRA LÉGER)
# ===============================
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/product-service .

EXPOSE 3001

USER nonroot:nonroot
CMD ["/app/product-service"]
