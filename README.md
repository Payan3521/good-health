# Good Health - Proyecto Final Electiva 3

Bienvenido al repositorio principal de **Good Health**, una plataforma orientada a microservicios para la gestión e historial médico.

Este proyecto utiliza una arquitectura de microservicios distribuida con múltiples tecnologías especializadas tanto para el Backend como para el Frontend.

## Estructura del Proyecto

Cada directorio del proyecto aloja una capa de la aplicación o un microservicio de manera independiente:

- **[`/authentication-microservice`](./authentication-microservice/README.md)**: Autenticación y gestión de usuarios (Spring Boot).
- **[`/doctorSchedule-microservice`](./doctorSchedule-microservice/README.MD)**: Gestión de agendas o citas médicas (Go).
- **[`/emergency-microservice`](./emergency-microservice/README.md)**: Gestión de emergencias médicas (Spring Boot).
- **[`/frontend-client`](./frontend-client/README.md)**: Aplicación web orientada a la experiencia de usuario (Angular).
- **[`/gateway-microservice`](./gateway-microservice/README.md)**: API Gateway que rutea las llamadas entre los múltiples servicios (NestJS).
- **[`/IAIntegration-microservice`](./IAIntegration-microservice/README.md)**: Integración con agentes de Inteligencia Artificial (FastAPI).
- **[`/infra`](./infra/README.md)**: Archivos compartidos de infraestructura y orquestación (Traefik, Kafka, Docker Compose).
- **[`/medicalRecord-microservice`](./medicalRecord-microservice/README.md)**: Historial clínico electrónico de los pacientes (FastAPI).
- **[`/notification-microservice`](./notification-microservice/README.MD)**: Envío de alertas, emails y comunicación (.NET).
- **[`/pharmacy-microservice`](./pharmacy-microservice/README.md)**: Gestión de pedidos e inventario de la farmacia (Flask).
- **[`/register-microservice`](./register-microservice/README.md)**: Registro e incorporación de nuevos usuarios (NestJS).

> **Nota:** Por favor consulta el `README.md` ubicado en el interior de cada servicio para obtener información más detallada del stack tecnológico, variables de entorno, puertos y sus respectivos comandos de ejecución y construcción.
