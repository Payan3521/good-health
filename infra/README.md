# Infraestructura - Good Health

Este directorio contiene la configuración de infraestructura necesaria para el despliegue y funcionamiento de los microservicios del proyecto **Good Health**.

La infraestructura está dividida en diferentes componentes, cada uno en su respectivo directorio:

## Componentes

### 1. Enrutamiento (Traefik) - `/routing`
Contiene la configuración de **Traefik**, que actúa como API Gateway, proxy inverso y balanceador de carga. Se encarga de recibir las peticiones externas y enrutarlas a los microservicios correspondientes.

### 2. Broker de Mensajería (Kafka) - `/broker`
Contiene la configuración del broker de **Apache Kafka**. Kafka se utiliza como la columna vertebral para la comunicación asíncrona y la arquitectura orientada a eventos entre los diferentes microservicios, asegurando la escalabilidad y la entrega de mensajes de manera confiable.

### 3. Orquestación Local (Docker Compose) - `/compose`
En este directorio se encuentra el archivo `docker-compose.yml` (y otros recursos relacionados con Compose), utilizado para centralizar y levantar de forma local toda la pila de infraestructura (Traefik, Kafka, bases de datos y herramientas de monitoreo) de manera rápida y consistente durante el desarrollo.

## Cómo empezar
Para poder iniciar la infraestructura en un entorno de desarrollo local, comúnmente debes dirigirte al directorio respectivo (o utilizar el `docker-compose.yml` principal) y ejecutar:

```bash
cd compose
docker-compose up -d
```

> **Nota:** Revise la documentación en cada subdirectorio para detalles específicos de configuración, variables de entorno u otros pre-requisitos antes de iniciar los servicios.