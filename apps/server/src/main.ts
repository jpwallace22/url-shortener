import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { getConfig } from '~/config';

const { serverPort, baseUrl } = getConfig();

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.setGlobalPrefix('/api');
  console.log(`Listening on ${baseUrl}${serverPort}`);
  await app.listen(serverPort);
}
bootstrap();
