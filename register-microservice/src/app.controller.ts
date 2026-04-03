import { Controller, Get } from '@nestjs/common';

@Controller()
export class AppController {
  @Get('/')
  getTest(): string {
    return 'Register Microservice is running with NestJS and TypeScript';
  }
}
