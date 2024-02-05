import { Body, Controller, Post, UnauthorizedException } from '@nestjs/common';
import { UrlService } from './url.service';
import { ShortenUrlDto, VerifyPasswordDto } from '../url/dto';

@Controller('urls')
export class UrlController {
  constructor(private readonly urlService: UrlService) {}

  // POST /urls/shorten
  @Post('shorten')
  createNewUrl(@Body() createUrlDto: ShortenUrlDto) {
    return this.urlService.createNewUrl(createUrlDto);
  }

  // POST /urls/verify
  @Post('verify')
  async verify(@Body() { url_id, password }: VerifyPasswordDto) {
    const url = await this.urlService.verify(url_id, password);

    return { url, statusCode: 200 };
  }
}
