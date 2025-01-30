import * as amqp from "amqplib";
import * as store from "./store";
import { POST_LIKED_EVENT, DELIVERY_MODE_PERRSISTENT } from "./amqp";

export class LikeService {
  private channel: amqp.Channel;

  constructor(channel: amqp.Channel) {
    this.channel = channel;
  }

  async createLike(data: LikeData) {
    const like = await store.save(data);

    const published = this.channel.publish(
      POST_LIKED_EVENT,
      "",
      Buffer.from(JSON.stringify(data)),
      {
        contentType: "application/json",
        deliveryMode: DELIVERY_MODE_PERRSISTENT,
      }
    );

    console.log({ published });

    return like;
  }
}
