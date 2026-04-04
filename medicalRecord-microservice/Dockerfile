FROM python:3.13-alpine AS build

WORKDIR /app

# Copiar archivos de dependencias e instalarlas
COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

# Copiar el código fuente
COPY . .

# Imagen final ligera
FROM python:3.13-alpine

WORKDIR /app

# Copiar dependencias instaladas desde la etapa de construcción
COPY --from=build /usr/local/lib/python3.13/site-packages /usr/local/lib/python3.13/site-packages
COPY --from=build /usr/local/bin /usr/local/bin

# Copiar el código fuente
COPY --from=build /app .

# Exponer el puerto del microservicio
EXPOSE 8084

# Comando de entrada
ENTRYPOINT ["python", "main.py"]