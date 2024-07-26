import { Injectable, OnModuleDestroy, OnModuleInit } from '@nestjs/common';
import * as amqp from 'amqplib';

@Injectable()
export class MessageQueueService implements OnModuleInit, OnModuleDestroy {
  private connection: any;

  async onModuleInit() {
    const queue = 'messages';
    this.connection = await amqp.connect('amqp://guest:guest@172.17.0.2:5672');
    const channel1 = await this.connection.createChannel();
    await channel1.asserQueue(queue);
    console.log('amqb connection is started');
  }

  async onModuleDestroy() {
    await this.connection.close();
    console.log('amqb connection is closed');

  }
}
