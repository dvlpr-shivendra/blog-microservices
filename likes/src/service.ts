import * as amqp from "amqplib";
import * as store from "./store";
import { POST_LIKED_EVENT, DELIVERY_MODE_PERSISTENT } from "./amqp";
import { LikeData } from "./types";
import * as validator from "./validator";

export class LikeService {
  private channel: amqp.Channel;

  constructor(channel: amqp.Channel) {
    this.channel = channel;
  }

  async createLike(data: LikeData) {
    // Validate input data
    if (!validator.isValidPostId(data.postId)) {
      throw new Error("Invalid postId format");
    }

    if (!validator.isValidUserId(data.userId)) {
      throw new Error("Invalid userId format");
    }

    const like = await store.save(data);

    const published = this.channel.publish(
      POST_LIKED_EVENT,
      "",
      Buffer.from(JSON.stringify(data)),
      {
        contentType: "application/json",
        deliveryMode: DELIVERY_MODE_PERSISTENT,
      }
    );

    console.log({ published });

    return like;
  }
}
