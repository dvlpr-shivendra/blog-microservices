import * as amqp from "amqplib";
import { POST_LIKED_EVENT, DELIVERY_MODE_PERSISTENT } from "./amqp";
import { LikeData } from "./types";
import * as validator from "./validator";
import logger from "./logger";
import { Repository } from "./repository.interface";

export class LikeService {
  private channel: amqp.Channel;
  private repository: Repository;

  constructor(channel: amqp.Channel, repository: Repository) {
    this.channel = channel;
    this.repository = repository;
  }

  async createLike(data: LikeData) {
    // Validate input data
    if (!validator.isValidPostId(data.postId)) {
      throw new Error("Invalid postId format");
    }

    if (!validator.isValidUserId(data.userId)) {
      throw new Error("Invalid userId format");
    }

    const like = await this.repository.create(data);

    const published = this.channel.publish(
      POST_LIKED_EVENT,
      "",
      Buffer.from(JSON.stringify(data)),
      {
        contentType: "application/json",
        deliveryMode: DELIVERY_MODE_PERSISTENT,
      }
    );

    logger.info({ published });

    return like;
  }
}
