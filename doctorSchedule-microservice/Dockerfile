FROM golang:1.26-alpine AS build

WORKDIR /app

# Copiar archivos de dependencias y descargarlas
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación (CGO_ENABLED=0 para asegurar binario estático compatible con Alpine)
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

# Imagen final ligera
FROM alpine:latest

WORKDIR /app

# Copiar el binario compilado desde la etapa de construcción
COPY --from=build /app/app app

# Exponer el puerto del microservicio
EXPOSE 8081

# Comando de entrada
ENTRYPOINT ["./app"]
