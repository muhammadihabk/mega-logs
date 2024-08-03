import { Module } from '@nestjs/common';
import { MessageQueueModule } from './message-queue/message-queue.module';
import { RouteRequestModule } from './route-requests/route-request.module';
import { ConfigModule } from '@nestjs/config';

@Module({
  imports: [
    MessageQueueModule,
    RouteRequestModule,
    ConfigModule.forRoot({ isGlobal: true }),
  ],
})
export class AppModule {}
