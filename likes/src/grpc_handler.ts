import * as grpc from "@grpc/grpc-js";
import { LikeService } from "./service";
import logger from "./logger";

export default class GRPCHandler {
  private service: LikeService;
  constructor(service: LikeService) {
    this.service = service;
  }

  createLike: grpc.handleUnaryCall<{ PostId: string }, { success: boolean }> =
    async (call, callback) => {
      try {
        const postId = call.request.PostId;

        const userId = "1";

        await this.service.createLike({ postId, userId });
        callback(null, { success: true });
      } catch (error) {
        logger.error("Error creating like:", error);

        callback({
          code: grpc.status.INTERNAL,
          message: "Internal error occurred",
        });
      }
    };
}
