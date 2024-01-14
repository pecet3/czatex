FROM golang:latest as builder

# Ustaw zmienne środowiskowe
ENV GO111MODULE=on

# Skopiuj pliki z projektu do kontenera
COPY . /app

# Przejdź do katalogu z kodem źródłowym
WORKDIR /app/cmd

# Zainstaluj dodatkowe narzędzia potrzebne do budowy
RUN apt-get update && apt-get install -y --no-install-recommends \
    git \
    && rm -rf /var/lib/apt/lists/*

# Zbuduj aplikację
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Drugi etap, aby utworzyć obraz Alpine
FROM alpine:latest

# Instaluj pakiet iproute2, który jest potrzebny dla niektórych funkcji (jeśli są używane)
RUN apk add --no-cache iproute2

# Skopiuj skompilowany plik binarny z pierwszego etapu
COPY --from=builder /app/cmd/app /app/app

# Udostępnij port, na którym będzie działać aplikacja
EXPOSE 8080

# Ustaw punkt wejścia dla kontenera
ENTRYPOINT ["/app/app"]

