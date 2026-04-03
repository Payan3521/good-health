# API Gateway - Good Health

Este directorio contiene el API Gateway, que actúa como punto de entrada centralizado conectando el frontend con los microservicios del proyecto **Good Health**.

## Stack Tecnológico
- **Tecnologías:** TypeScript 5.7.3 + NestJS 11.0.1
- **Gestor de Dependencias:** npm/npx 11.9.0
- **Integración:** Traefik

## Ejecución
- **Puerto expuesto:** `8000`

### Comandos disponibles

| Acción | Comando |
| :--- | :--- |
| **Run** | `npm run start:dev` |
| **Clean** | `rm -rf dist` |
| **Build** | `npm run build` |
| **Clean & Run** | `rm -rf dist && npm run start:dev` |