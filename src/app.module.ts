import { Module } from '@nestjs/common';
import { MessageQueueModule } from './message-queue/message-queue.module';

@Module({
  imports: [MessageQueueModule],
})
export class AppModule {}
