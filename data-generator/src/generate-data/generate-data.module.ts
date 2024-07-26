import { Module } from '@nestjs/common';
import { GenerateDataService } from './generate-data.service';
import { GenerateDataController } from './generate-data.controller';

@Module({
  controllers: [GenerateDataController],
  providers: [GenerateDataService],
})
export class GenerateDataModule {}
