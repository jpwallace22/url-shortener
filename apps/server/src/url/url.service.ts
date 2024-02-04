import { Injectable } from '@nestjs/common';
import { DbService } from '../db/db.service';
import { ShortenUrlDto } from 'src/url/dto';
import { ConfigService } from '@nestjs/config';
import { nanoid } from 'nanoid';
import * as argon2 from 'argon2';

@Injectable()
export class UrlService {
  constructor(
    private db: DbService,
    private config: ConfigService,
  ) {}

  private ID_SIZE = 10;

  async createUrl({ url, password }: ShortenUrlDto) {
    const hash = password ? await argon2.hash(password) : null;
    const id = nanoid(this.ID_SIZE);
    const short_url = new URL(id, this.config.get('PUBLIC_API_URL')).toString();

    const { hash: _removed, ...response } = await this.db.url.create({
      data: {
        url_id: id,
        url,
        short_url,
        hash,
      },
    });

    return response;
  }
}
