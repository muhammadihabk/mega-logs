import { Controller, InternalServerErrorException, Post } from '@nestjs/common';
import { GenerateDataService } from './generate-data.service';

@Controller('generate-data')
export class GenerateDataController {
  constructor(private readonly generateDataService: GenerateDataService) {}

  @Post('fix-products-file')
  async fixProductsFile() {
    try {
      await this.generateDataService.fixProductsFile();
      return {
        message: 'Succeeded',
      };
    } catch (error) {
      throw new InternalServerErrorException();
    }
  }

  @Post('send-data-to-gateway')
  async sendDataToGateway() {
    try {
      await this.generateDataService.sendDataToGateway();
      return {
        message: 'Succeeded',
      };
    } catch (error) {
      throw new InternalServerErrorException();
    }
  }
}
