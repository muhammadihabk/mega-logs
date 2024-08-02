import { Injectable, OnModuleDestroy, OnModuleInit } from '@nestjs/common';
import * as amqp from 'amqplib';
import { Queues } from './enums';

@Injectable()
export class MessageQueueService implements OnModuleInit, OnModuleDestroy {
  private connection: any;
  private channel: any;

  async onModuleInit() {
    this.connection = await amqp.connect('amqp://guest:guest@messageQueue:5672');
    this.channel = await this.connection.createChannel();
    await this.channel.assertQueue(Queues.CUSTOMERS);
    await this.channel.assertQueue(Queues.PRODUCTS);
    await this.channel.assertQueue(Queues.ORDERS);
    await this.channel.assertQueue(Queues.ORDER_ITEMS);
    await this.channel.assertQueue(Queues.SELLERS);
    console.log('amqb connection is started');
  }

  getConnection() {
    return this.channel;
  }

  async onModuleDestroy() {
    await this.connection.close();
    console.log('amqb connection is closed');

  }
}
