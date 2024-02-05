import {
  Injectable,
  NotFoundException,
  UnauthorizedException,
} from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { DbService } from 'src/db/db.service';
import * as argon from 'argon2';
import { Url } from '@prisma/client';

@Injectable()
export class AppService {
  constructor(
    private db: DbService,
    private config: ConfigService,
  ) {}

  async redirect(id: string) {
    try {
      const data = await this.db.url.findFirst({
        where: {
          url_id: id,
        },
      });

      if (!data) throw new NotFoundException();

      await this.addClickToUrl(data);

      if (data.hash)
        return new URL(
          `${this.config.get('PUBLIC_APP_URL')}/verify?id=${id}`,
        ).toString();

      return new URL(data.url).toString();
    } catch (e) {
      throw new NotFoundException();
    }
  }

  private async addClickToUrl({ url_id, clicks }: Url) {
    await this.db.url.update({
      where: {
        url_id,
      },
      data: {
        clicks: clicks + 1,
      },
    });
  }
}
