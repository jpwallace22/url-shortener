import { Body, Controller, Post } from '@nestjs/common';
import { UrlService } from './url.service';
import { ShortenUrlDto } from '../url/dto';

@Controller('urls')
export class UrlController {
  constructor(private readonly urlService: UrlService) {}

  // POST api/urls/shorten
  @Post('shorten')
  createUrl(@Body() createUrlDto: ShortenUrlDto) {
    return this.urlService.createUrl(createUrlDto);
  }
}
