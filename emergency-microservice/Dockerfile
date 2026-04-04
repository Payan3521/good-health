FROM eclipse-temurin:21-jdk-alpine AS build

WORKDIR /app

# Copiar el wrapper de gradle y los archivos de configuración
COPY gradlew .
COPY gradle gradle
COPY build.gradle settings.gradle ./

# Dar permisos de ejecución al wrapper
RUN chmod +x gradlew

# Descargar las dependencias
RUN ./gradlew dependencies --no-daemon

# Copiar el código fuente
COPY src src

# Compilar el proyecto omitiendo los tests
RUN ./gradlew build -x test --no-daemon

# Imagen final ligera con JRE
FROM eclipse-temurin:21-jre-alpine

WORKDIR /app

# Copiar el jar compilado (utilizando comodín por si tiene versión en el nombre)
COPY --from=build /app/build/libs/*.jar app.jar

# Exponer el puerto del microservicio
EXPOSE 8082

# Entorno activo
ENV SPRING_PROFILES_ACTIVE=docker

# Comando de entrada
ENTRYPOINT ["java", "-jar", "app.jar"]
