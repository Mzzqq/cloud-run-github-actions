# Gunakan base image untuk Go
FROM golang:1.22

# Set working directory
WORKDIR /app

# Copy semua file ke container
COPY . .

# Build aplikasi
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]