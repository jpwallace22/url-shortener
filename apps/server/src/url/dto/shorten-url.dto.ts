import {
  IsNotEmpty,
  IsOptional,
  IsString,
  IsUrl,
  Length,
  MaxLength,
} from 'class-validator';

export class ShortenUrlDto {
  @IsString()
  @IsNotEmpty()
  @IsUrl()
  @MaxLength(2048)
  url: string;

  @IsString()
  @IsOptional()
  @Length(8, 256)
  password?: string;
}
