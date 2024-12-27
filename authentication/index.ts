import * as grpc from '@grpc/grpc-js';
import * as protoLoader from '@grpc/proto-loader';
import * as path from 'path';
import { ConsulRegistry } from './discovery/consul';
import { env } from './helpers/app';
import { signup } from './grpc_handler';


import 'dotenv/config'
import mongoose from 'mongoose';

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
const { AuthService } = (protoDescriptor.proto as any);

mongoose
    .connect(env("MONGODB_URI", ""))
    .then(() => console.log("Connected to DB"))
    .catch((err) => console.error("Failed to connect to DB:", err));

async function main() {
    const server = new grpc.Server();

    server.addService(AuthService.service, {
        signup
    });

    const grpcHost = env('GRPC_HOST', 'localhost');
    const grpcPort = parseInt(env('GRPC_PORT', '2004'), 10);

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

            const registry = new ConsulRegistry(grpcHost, grpcPort, 'authentication');
            await registry.register();

            function shutdown() {
                console.log('Shutting down...');
                registry.deregister();
                server.forceShutdown();
                process.exit(0);
            }

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
