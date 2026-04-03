# Doctor Schedule Microservice - Good Health

Este directorio contiene el microservicio encargado de la gestión de horarios y agendas médicas para el proyecto **Good Health**.

## Stack Tecnológico
- **Tecnologías:** Go 1.26.1 + GIN 1.12.0
- **Base de Datos:** PostgreSQL (SQL)

## Ejecución
- **Puerto expuesto:** `8081`

### Comandos disponibles

| Acción | Comando |
| :--- | :--- |
| **Run** | `go run main.go` o `go run .` |
| **Clean** | `go clean -cache` |
| **Build** | `go build -o app main.go` |
| **Clean & Run** | `go clean -cache && go run main.go` |