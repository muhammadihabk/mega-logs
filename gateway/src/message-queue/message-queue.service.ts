import { Injectable, OnModuleDestroy, OnModuleInit } from '@nestjs/common';
import * as amqp from 'amqplib';
import { Queues } from './enums';
import { ConfigService } from '@nestjs/config';

@Injectable()
export class MessageQueueService implements OnModuleInit, OnModuleDestroy {
  private connection: any;

  private channel: any;

  constructor(private configService: ConfigService) {}

  async onModuleInit() {
    const amqpUrl = this.configService.get<string>('AMQP_URL');
    this.connection = await amqp.connect(amqpUrl);
    this.channel = await this.connection.createChannel();

    const options = {
      arguments: {
        'x-queue-type': 'quorum',
        'x-delivery-limit': 5,
        'x-dead-letter-exchange': 'dlx_exchange',
        'x-dead-letter-routing-key': 'dlx_routing_key',
      },
      durable: true,
    };
    await this.channel.assertQueue(Queues.CUSTOMERS, options);
    await this.channel.assertQueue(Queues.PRODUCTS, options);
    await this.channel.assertQueue(Queues.ORDERS, options);
    await this.channel.assertQueue(Queues.ORDER_ITEMS, options);
    await this.channel.assertQueue(Queues.SELLERS, options);

    const dlxExchange = 'dlx_exchange';
    const dlxQueue = 'dlx_queue';
    const dlxQueueOptions = {
      arguments: { 'x-queue-type': 'quorum' },
      durable: true,
    };
    await this.channel.assertExchange(dlxExchange, 'direct');
    await this.channel.assertQueue(dlxQueue, dlxQueueOptions);
    await this.channel.bindQueue(dlxQueue, dlxExchange, 'dlx_routing_key');
    console.log('amqb connection has started.');
  }

  getConnection() {
    return this.channel;
  }

  async onModuleDestroy() {
    await this.connection.close();
    console.log('amqb connection has been closed.');
  }
}
