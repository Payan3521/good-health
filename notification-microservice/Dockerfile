FROM mcr.microsoft.com/dotnet/sdk:10.0 AS build

WORKDIR /app

# Copiar el archivo de proyecto y restaurar dependencias
COPY *.csproj ./
RUN dotnet restore

# Copiar el código fuente
COPY . .

# Compilar y publicar el proyecto
RUN dotnet publish -c Release -o out

# Imagen final ligera
FROM mcr.microsoft.com/dotnet/aspnet:10.0

WORKDIR /app

# Copiar el build publicado desde la etapa de construcción
COPY --from=build /app/out .

# Exponer el puerto del microservicio
EXPOSE 8085

# Comando de entrada
ENTRYPOINT ["dotnet", "NotificationMicroservice.dll"]