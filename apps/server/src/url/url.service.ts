import {
  Injectable,
  NotFoundException,
  UnauthorizedException,
} from '@nestjs/common';
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

  async createNewUrl({ url, password }: ShortenUrlDto) {
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

  async verify(id: string, password: string) {
    try {
      const data = await this.db.url.findFirst({
        where: {
          url_id: id,
        },
      });

      if (!data || !data.hash) throw new NotFoundException();

      const isValid = await argon2.verify(data.hash, password);

      if (!isValid) throw new UnauthorizedException();

      return data.url;
    } catch (e) {
      if (e instanceof UnauthorizedException)
        throw new UnauthorizedException(
          'Password and/or the provided URL do not match',
        );

      throw new NotFoundException();
    }
  }
}
