# syntax=docker/dockerfile:1

FROM golang

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download
COPY . .



# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go 

ENV PORT=8080
ENV CERT_FILE=""
ENV KEY_FILE=""

# Skopiuj certyfikat i klucz do kontenera
COPY /etc/letsencrypt/live/czatex.pecet.it/fullchain.pem /fullchain.pem
COPY /priv_path /privkey.pem

# Ustaw zmienne środowiskowe dla konfiguracji aplikacji
ENV PORT=8080
ENV CERT_FILE=/fullchain.pem
ENV KEY_FILE=/privkey.pem

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
CMD ["/app/main"]