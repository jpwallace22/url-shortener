import { Controller, Get, Param, Redirect } from '@nestjs/common';
import { AppService } from 'src/app/app.service';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  // GET /:id
  @Get(':id')
  @Redirect()
  async redirect(@Param('id') id: string) {
    const url = await this.appService.redirect(id);
    return { url, statusCode: 302 };
  }
}
