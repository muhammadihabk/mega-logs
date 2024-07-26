import { Body, Controller, Post } from '@nestjs/common';
import { Queues } from 'src/message-queue/enums';
import { MessageQueueService } from 'src/message-queue/message-queue.service';

@Controller('route-requests')
export class RouteRequestsController {
  constructor(private messageQueueService: MessageQueueService) {}

  @Post('customers')
  sendCreateCustomer(@Body() body) {
    this.messageQueueService
      .getConnection()
      .sendToQueue(Queues.CUSTOMERS, Buffer.from(JSON.stringify(body)));
    return {
      message: 'Customer to be created.',
    };
  }

  @Post('products')
  sendCreateProduct(@Body() body) {
    this.messageQueueService
      .getConnection()
      .sendToQueue(Queues.PRODUCTS, Buffer.from(JSON.stringify(body)));
    return {
      message: 'Product to be created.',
    };
  }

  @Post('orders')
  sendCreateOrders(@Body() body) {
    this.messageQueueService
      .getConnection()
      .sendToQueue(Queues.ORDERS, Buffer.from(JSON.stringify(body)));
    return {
      message: 'Order to be created.',
    };
  }

  @Post('ordersItems')
  sendCreateOrdersItems(@Body() body) {
    this.messageQueueService
      .getConnection()
      .sendToQueue(Queues.ORDER_ITEMS, Buffer.from(JSON.stringify(body)));
    return {
      message: 'Order Item to be created.',
    };
  }
}
