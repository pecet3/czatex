FROM golang:latest AS builder

# Set destination for COPY
WORKDIR /app

# Copy only necessary files for module download
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

# Stage 2: Runtime stage
FROM scratch

# Set destination for COPY
WORKDIR /app

# Copy only the built binary from the previous stage
COPY --from=builder /app/main .

# Copy the SSL certificate and key
COPY /etc/letsencrypt/live/czatex.pecet.it/fullchain.pem /fullchain.pem
COPY /etc/letsencrypt/live/czatex.pecet.it/privkey.pem /privkey.pem

# Set environment variables
ENV PORT=8080
ENV CERT_FILE=/app/fullchain.pem
ENV KEY_FILE=/app/privkey.pem

# Expose the port the app runs on
EXPOSE $PORT

# Run the application
CMD ["./main"]