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
      console.log('Failed to fix the products file', error);
      return {
        message: 'Failed',
      };
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
      console.log('Failed to send data to the gateway', error);
      return {
        message: 'Failed',
      };
    }
  }
}
