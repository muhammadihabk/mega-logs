import { Module } from '@nestjs/common';
import { MessageQueueService } from './message-queue.service';

@Module({
  providers: [MessageQueueService],
})
export class MessageQueueModule {}
