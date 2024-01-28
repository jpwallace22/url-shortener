import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';

const port = parseInt(process.env.PORT || '3000', 10);

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.setGlobalPrefix('/api');
  console.log(`Listening on http://localhost:${port}`);
  await app.listen(port);
}
bootstrap();
