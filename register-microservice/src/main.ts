import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';

async function bootstrap() {  
  const app = await NestFactory.create(AppModule);
  await app.listen(process.env.PORT ?? 8087);
  console.log(`Register Microservice is running on port ${process.env.PORT ?? 8087}`);
}
bootstrap();
