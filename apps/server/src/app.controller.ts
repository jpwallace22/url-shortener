import { Controller, Get, Post } from '@nestjs/common';
import { DbService } from '~/db/db.service';

@Controller()
export class AppController {
  constructor(private db: DbService) {}
  @Get()
  sayHello() {
    return { message: 'Hello from nest!' };
  }

  @Post()
  addToDb() {
    const res = this.db.user.create({
      data: {
        url: 'https://www.justinwallace.dev',
        hash: 'blah',
      },
    });
    return res;
  }
}
