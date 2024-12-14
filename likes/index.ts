import * as grpc from '@grpc/grpc-js';
import * as protoLoader from '@grpc/proto-loader';
import * as path from 'path';
import { ConsulRegistry } from './discovery/consul';
import { env } from './helpers/app';

import 'dotenv/config'

// Path to the proto file
const PROTO_PATH = path.resolve(__dirname, '../common/proto/blog.proto');

// Load protobuf
const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
});

const protoDescriptor = grpc.loadPackageDefinition(packageDefinition) as any;
const { LikeService } = (protoDescriptor.proto as any);

// In-memory storage for likes
const likes: Array<{ PostId: number; UserId: number }> = [];

// Implement the CreateLike method
const createLike = (
    call: grpc.ServerUnaryCall<{ PostId: number }, any>,
    callback: grpc.sendUnaryData<{ success: boolean }>
) => {
    try {
        const { PostId } = call.request;
        const UserId = Math.floor(Math.random() * 1000); // Simulated user ID
        
        likes.push({ PostId, UserId });
        
        console.log(`Like created: PostId=${PostId}, UserId=${UserId}`);
        callback(null, { success: true });
    } catch (error) {
        console.error('Error creating like:', error);
        callback({
            code: grpc.status.INTERNAL,
            message: 'Internal error occurred'
        });
    }
};

async function main() {
    // Create gRPC server
    const server = new grpc.Server();
    
    // Add the LikeService service implementation to the server
    server.addService(LikeService.service, {
        createLike: createLike
    });

    // Get configuration from environment
    const grpcHost = env('GRPC_HOST', 'localhost');
    const grpcPort = parseInt(env('GRPC_PORT', '2003'), 10);

    // Start the server
    const serverAddr = `${grpcHost}:${grpcPort}`;
    server.bindAsync(
        serverAddr,
        grpc.ServerCredentials.createInsecure(),
        async (error) => {
            if (error) {
                console.error('Failed to bind server:', error);
                return;
            }

            console.log(`GRPC Server running at ${serverAddr}`);

            // Register with Consul
            const registry = new ConsulRegistry(grpcHost, grpcPort, 'likes');
            await registry.register();

            function shutdown() {
                console.log('Shutting down...');
                registry.deregister();
                server.forceShutdown();
                process.exit(0);
            }

            // Handle graceful shutdown
            process.on('SIGTERM', shutdown);
            process.on('SIGINT', shutdown);
            process.on('SIGQUIT', shutdown);
            process.on('SIGHUP', shutdown);
        }
    );
}

main().catch(error => {
    console.error('Failed to start server:', error);
    process.exit(1);
});
