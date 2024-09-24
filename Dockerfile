# Gunakan image Go untuk build
FROM golang:alpine AS builder

# Set direktori kerja di container
WORKDIR /app

# Copy semua file dari direktori lokal ke container
COPY . .

# Build aplikasi
# RUN go mod init cpuload && go mod tidy
RUN go build -o cpuload .

# Stage kedua: menggunakan base image yang lebih ringan
FROM alpine:latest

WORKDIR /root/

# Copy binary yang sudah dibuild dari stage builder
COPY --from=builder /app/cpuload .

# Jalankan aplikasi
CMD ["./cpuload"]
