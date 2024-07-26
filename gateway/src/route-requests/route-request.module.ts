import { Module } from '@nestjs/common';
import { RouteRequestsController } from './route-requests.controller';
import { MessageQueueService } from 'src/message-queue/message-queue.service';

@Module({
  controllers: [RouteRequestsController],
  providers: [MessageQueueService],
})
export class RouteRequestModule {}
