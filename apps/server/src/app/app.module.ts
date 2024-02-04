import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { getConfig } from '../config';
import { DbService } from '../db/db.service';
import { UrlModule } from '../url/url.module';
import { DbModule } from 'src/db/db.module';
import { AppController } from 'src/app/app.controller';
import { AppService } from './app.service';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      load: [getConfig],
    }),
    UrlModule,
    DbModule,
  ],
  controllers: [AppController],
  providers: [DbService, AppService],
})
export class AppModule {}
