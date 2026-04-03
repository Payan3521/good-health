*MEDICAL RECORD - MICROSERVICE*

TECNOLOGIAS: Python **3.13.5** + FastAPI **0.135.3**
GESTOR DE DEPENDENCIAS: Pip **25.1.1**
BASE DE DATOS: NoSQL MongoDB **version**

RUN: python main.py
DEBUG: ...
CLEAN: find . -type d -name "__pycache__" -exec rm -rf {} +
BUILD: pip freeze > requirements.txt

CLEAN AND RUN: find . -type d -name "__pycache__" -exec rm -rf {} + && python main.py

PUERTO: 8084