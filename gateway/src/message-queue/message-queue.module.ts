import { Module } from '@nestjs/common';
import { MessageQueueService } from './message-queue.service';

@Module({
  exports: [MessageQueueService],
  providers: [MessageQueueService],
})
export class MessageQueueModule {}
