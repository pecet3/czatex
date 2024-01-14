# Użyj oficjalnego obrazu Go jako bazowego obrazu
FROM golang:latest

# Ustaw zmienną środowiskową GOPATH
RUN export GO111MODULE=on

# Przejdź do katalogu z kodem źródłowym
RUN mkdir /build
WORKDIR /build

RUN cd /build && go git clone github.com/pecet3/czatex

RUN cd /build/czatex/cmd && go build

# Skompiluj aplikację
RUN go build -o czatex .

# Expose the port on which the application will run
EXPOSE 3000

ENTRYPOINT [ "/build/czatex/cmd/main" ]
