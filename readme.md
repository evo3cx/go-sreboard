# Simple API webservice with Golang

Simmple Service run By Golang,
Service ini akan semua perintah pada terminal melalui param URL

Contoh:
  127.0.0.1:8080/ls_-la

  output: akan menghasilkan list file/directory pada folder server


# Install

Makesure you have govendor installed, and run

`govendor sync`

Run Service with command

`go run server.go`

# Docker

Run with docker

`docker build -t sregolang .`

`docker run sregolang`
