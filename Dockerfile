# Użyj oficjalnego obrazu Go jako bazowego obrazu
FROM golang:latest

# Ustaw zmienne środowiskowe
ENV GO111MODULE=on

# Skopiuj pliki z projektu do kontenera
COPY . /app

# Przejdź do katalogu z kodem źródłowym
WORKDIR /app/cmd

# Zbuduj aplikację
RUN go build -o app

# Udostępnij port, na którym będzie działać aplikacja
EXPOSE 8080

# Ustaw punkt wejścia dla kontenera
ENTRYPOINT ["/app/cmd/app"]
