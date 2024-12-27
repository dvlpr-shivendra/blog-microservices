import * as grpc from '@grpc/grpc-js';
import * as service from './service';

interface Response {
    token: string,
    message: string
}

export const signup: grpc.handleUnaryCall<UserPayload, Response> = async (
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

        const token = await service.createUser({ name, username, email, password })

        callback(null, { token, message: "success" });
    } catch (error) {
        console.error('Error creating like:', error);
        callback({
            code: grpc.status.INTERNAL,
            message: 'Internal error occurred',
        });
    }
};
