import { Body, Controller, Post } from '@nestjs/common';
import { Queues } from 'src/message-queue/enums';
import { MessageQueueService } from 'src/message-queue/message-queue.service';

@Controller('route-requests')
export class RouteRequestsController {
  constructor(private messageQueueService: MessageQueueService) {}

  @Post('customers')
  sendCreateCustomer(@Body() body) {
    try {
      this.messageQueueService
        .getConnection()
        .sendToQueue(Queues.CUSTOMERS, Buffer.from(JSON.stringify(body)));
      return {
        message: 'Customer to be created.',
      };
    } catch (error) {
      console.log('Failed to send customers message to the queue', error);
      return {
        message: 'Failed',
      };
    }
  }

  @Post('products')
  sendCreateProduct(@Body() body) {
    try {
      this.messageQueueService
        .getConnection()
        .sendToQueue(Queues.PRODUCTS, Buffer.from(JSON.stringify(body)));
      return {
        message: 'Product to be created.',
      };
    } catch (error) {
      console.log('Failed to send products message to the queue', error);
      return {
        message: 'Failed',
      };
    }
  }

  @Post('orders')
  sendCreateOrders(@Body() body) {
    try {
      this.messageQueueService
        .getConnection()
        .sendToQueue(Queues.ORDERS, Buffer.from(JSON.stringify(body)));
      return {
        message: 'Order to be created.',
      };
    } catch (error) {
      console.log('Failed to send orders message to the queue', error);
      return {
        message: 'Failed',
      };
    }
  }

  @Post('orderItems')
  sendCreateOrderItems(@Body() body) {
    try {
      this.messageQueueService
        .getConnection()
        .sendToQueue(Queues.ORDER_ITEMS, Buffer.from(JSON.stringify(body)));
      return {
        message: 'Order Item to be created.',
      };
    } catch (error) {
      console.log('Failed to send orders item message to the queue', error);
      return {
        message: 'Failed',
      };
    }
  }

  @Post('sellers')
  sendCreateOrderSellers(@Body() body) {
    try {
      this.messageQueueService
        .getConnection()
        .sendToQueue(Queues.SELLERS, Buffer.from(JSON.stringify(body)));
      return {
        message: 'Seller to be created.',
      };
    } catch (error) {
      console.log('Failed to send seller message to the queue', error);
      return {
        message: 'Failed',
      };
    }
  }
}
