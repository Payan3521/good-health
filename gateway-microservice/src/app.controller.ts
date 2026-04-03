import { Controller, Get } from '@nestjs/common';

@Controller()
export class AppController {
  @Get('/')
  getTest(): string {
    return 'Gateway Microservice is running with NestJS and TypeScript';
  }
}
