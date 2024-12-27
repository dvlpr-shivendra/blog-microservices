import * as grpc from '@grpc/grpc-js';
import * as service from './service';

export const signup: grpc.handleUnaryCall<UserPayload, { success: boolean }> = async (
    call,
    callback
) => {
    try {
        const {
            name,
            username,
            email,
            password
        } = call.request;

        service.createUser({ name, username, email, password })

        callback(null, { success: true });
    } catch (error) {
        console.error('Error creating like:', error);
        callback({
            code: grpc.status.INTERNAL,
            message: 'Internal error occurred',
        });
    }
};
