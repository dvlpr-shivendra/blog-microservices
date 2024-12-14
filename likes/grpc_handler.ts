import * as grpc from '@grpc/grpc-js';
import * as service from './service';

// gRPC Handler for the LikeService
export const createLike: grpc.handleUnaryCall<{ PostId: number }, { success: boolean }> = async (
    call,
    callback
) => {
    try {
        const { PostId } = call.request;
        const UserId = 1; // Simulated user ID

        service.createLike({ postId: PostId, userId: UserId })

        console.log(`Like created: PostId=${PostId}, UserId=${UserId}`);
        callback(null, { success: true });
    } catch (error) {
        console.error('Error creating like:', error);
        callback({
            code: grpc.status.INTERNAL,
            message: 'Internal error occurred',
        });
    }
};
