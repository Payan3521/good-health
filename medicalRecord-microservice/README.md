# Medical Record Microservice - Good Health

Este directorio contiene el microservicio encargado del historial y registros médicos de los pacientes para el proyecto **Good Health**.

## Stack Tecnológico
- **Tecnologías:** Python 3.13.5 + FastAPI 0.135.3
- **Gestor de Dependencias:** Pip 25.1.1
- **Base de Datos:** MongoDB (NoSQL)

## Ejecución
- **Puerto expuesto:** `8084`

### Comandos disponibles

| Acción | Comando |
| :--- | :--- |
| **Run** | `python main.py` |
| **Clean** | `find . -type d -name "__pycache__" -exec rm -rf {} +` |
| **Build** | `pip freeze > requirements.txt` |
| **Clean & Run** | `find . -type d -name "__pycache__" -exec rm -rf {} + && python main.py` |