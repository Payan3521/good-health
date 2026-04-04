FROM node:24.14.0-alpine AS build

WORKDIR /app

# Copiar archivos de dependencias e instalarlas
COPY package*.json ./
RUN npm ci

# Copiar el código fuente
COPY . .

# Compilar el proyecto
RUN npm run build

# Imagen final ligera
FROM node:24.14.0-alpine

WORKDIR /app

# Copiar dependencias de producción únicamente
COPY package*.json ./
RUN npm ci --omit=dev

# Copiar el build compilado desde la etapa de construcción
COPY --from=build /app/dist ./dist

# Exponer el puerto del gateway
EXPOSE 8000

# Comando de entrada
ENTRYPOINT ["node", "dist/main"]