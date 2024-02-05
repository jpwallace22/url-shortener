import { IsNotEmpty, IsString, Length } from 'class-validator';

export class VerifyPasswordDto {
  @IsString()
  @IsNotEmpty()
  url_id: string;

  @IsString()
  @IsNotEmpty()
  @Length(8, 50)
  password: string;
}
