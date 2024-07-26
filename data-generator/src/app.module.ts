import { Module } from '@nestjs/common';
import { GenerateDataModule } from './generate-data/generate-data.module';

@Module({
  imports: [GenerateDataModule],
})
export class AppModule {}
