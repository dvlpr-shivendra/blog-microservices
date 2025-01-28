import * as grpc from "@grpc/grpc-js";
import { LikeService } from "./service";

export default class GRPCHandler {
  private service: LikeService;
  constructor(service: LikeService) {
    this.service = service;
  }

  createLike: grpc.handleUnaryCall<{ PostId: number }, { success: boolean }> =
    async (call, callback) => {
      try {
        const { PostId } = call.request;

        const UserId = 1;

        this.service.createLike({ postId: PostId, userId: UserId });

        console.log(`Like created: PostId=${PostId}, UserId=${UserId}`);

        callback(null, { success: true });
      } catch (error) {
        console.error("Error creating like:", error);

        callback({
          code: grpc.status.INTERNAL,
          message: "Internal error occurred",
        });
      }
    };
}
