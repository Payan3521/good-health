FROM node:24.14.0-alpine AS build

WORKDIR /app

# Copiar archivos de dependencias e instalarlas
COPY package*.json ./
RUN npm ci

# Copiar el código fuente
COPY . .

# Compilar el proyecto
RUN npx ng build --configuration=production

# Imagen final ligera con Nginx
FROM nginx:alpine

# Copiar el build compilado desde la etapa de construcción
COPY --from=build /app/dist/*/browser /usr/share/nginx/html

# Copiar configuración personalizada de Nginx
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Exponer el puerto del frontend
EXPOSE 4200

# Nginx corre por defecto
ENTRYPOINT ["nginx", "-g", "daemon off;"]