import { Module } from '@nestjs/common';
import { MessageQueueModule } from './message-queue/message-queue.module';
import { RouteRequestModule } from './route-requests/route-request.module';

@Module({
  imports: [MessageQueueModule, RouteRequestModule],
})
export class AppModule {}
