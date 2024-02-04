import { Injectable, NotFoundException } from '@nestjs/common';
import { DbService } from 'src/db/db.service';

@Injectable()
export class AppService {
  constructor(private db: DbService) {}

  async redirect(id: string) {
    try {
      const data = await this.db.url.findFirst({
        where: {
          url_id: id,
        },
      });
      if (!data) throw new NotFoundException();
      return data.url;
    } catch (e) {
      throw new NotFoundException();
    }
  }
}
