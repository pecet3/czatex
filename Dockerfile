# Etap 1: Zbuduj aplikację
FROM golang:latest AS builder

# Ustaw ścieżkę roboczą w kontenerze
WORKDIR /app

# Skopiuj pliki go.mod i go.sum i pobierz zależności
COPY go.mod go.sum ./
RUN go mod download

# Skopiuj cały kod źródłowy do kontenera
COPY . .

# Skompiluj aplikację do binarnego pliku wykonywalnego
RUN CGO_ENABLED=0 GOOS=linux go build -o /czatex

# Etap 2: Przygotuj lekki obraz do uruchomienia
FROM alpine:latest

# Utwórz katalog na pliki 'view' i skopiuj je
WORKDIR /app
COPY --from=builder /app/view ./view

# Skopiuj skompilowany plik wykonywalny z etapu budowania
COPY --from=builder /czatex .

# Ustaw port, na którym będzie nasłuchiwać aplikacja
EXPOSE 8080

# Uruchom aplikację
CMD ["./czatex"]