import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { getConfig } from '~/config';

const { SERVER_PORT, BASE_URL } = getConfig();

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.setGlobalPrefix('/api');
  console.log(`Listening on ${BASE_URL}${SERVER_PORT}`);
  await app.listen(SERVER_PORT);
}

bootstrap();
