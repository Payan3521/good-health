# Pharmacy Microservice - Good Health

Este directorio contiene el microservicio encargado de la gestión de farmacia y medicamentos para el proyecto **Good Health**.

## Stack Tecnológico
- **Tecnologías:** Python 3.13.5 + Flask 3.1.3
- **Gestor de Dependencias:** Pip 25.1.1
- **Base de Datos:** PostgreSQL (SQL)

## Ejecución
- **Puerto expuesto:** `8086`

### Comandos disponibles

| Acción | Comando |
| :--- | :--- |
| **Run** | `python main.py` |
| **Clean** | `find . -type d -name "__pycache__" -exec rm -rf {} +` |
| **Build** | `pip freeze > requirements.txt` |
| **Clean & Run** | `find . -type d -name "__pycache__" -exec rm -rf {} + && python main.py` |