import { NestFactory } from '@nestjs/core';
import { AppModule } from './app/app.module';
import { getConfig } from './config';
import { ValidationPipe } from '@nestjs/common';

const { SERVER_PORT, BASE_URL } = getConfig();

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  app.useGlobalPipes(
    new ValidationPipe({
      // whitelist: true,
      transform: true,
    }),
  );

  app.enableCors();

  console.log(`Listening on ${BASE_URL}${SERVER_PORT}`);
  await app.listen(SERVER_PORT);
}

bootstrap();
