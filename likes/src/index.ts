import * as grpc from "@grpc/grpc-js";
import * as protoLoader from "@grpc/proto-loader";
import * as path from "path";
import { ConsulRegistry } from "./discovery/consul";
import { env } from "./helpers/app";
import GRPCHandler from "./grpc_handler";
import { connect } from "./amqp";
import { db } from "./db"; // Import the database manager

import "dotenv/config";
import { LikeService } from "./service";

const PROTO_PATH = path.resolve(__dirname, "../../common/proto/blog.proto");

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

const protoDescriptor = grpc.loadPackageDefinition(packageDefinition) as any;
const { LikeService: UnimplementedLikeService } = protoDescriptor.proto as any;

async function main() {
  try {
    // Connect to MongoDB with connection pooling
    await db.connect();

    const { channel, close } = await connect(
      "guest",
      "guest",
      "localhost",
      "5672"
    );
    const service = new LikeService(channel);
    const grpcHandler = new GRPCHandler(service);
    const server = new grpc.Server();

    server.addService(UnimplementedLikeService.service, {
      createLike: grpcHandler.createLike,
    });

    const grpcHost = env("GRPC_HOST", "localhost");
    const grpcPort = parseInt(env("GRPC_PORT", "2002"), 10);

    const serverAddr = `${grpcHost}:${grpcPort}`;

    server.bindAsync(
      serverAddr,
      grpc.ServerCredentials.createInsecure(),
      async (error) => {
        if (error) {
          console.error("Failed to bind server:", error);
          return;
        }

        console.log(`GRPC Server running at ${serverAddr}`);

        const registry = new ConsulRegistry(grpcHost, grpcPort, "likes");
        await registry.register();

        function shutdown() {
          console.log("Shutting down...");
          registry.deregister();
          server.forceShutdown();
          db.disconnect().catch(console.error); // Close DB connection properly
          process.exit(0);
        }

        process.on("SIGTERM", shutdown);
        process.on("SIGINT", shutdown);
        process.on("SIGQUIT", shutdown);
        process.on("SIGHUP", shutdown);
      }
    );
  } catch (error) {
    console.error("Failed to start server:", error);
    process.exit(1);
  }
}

main().catch((error) => {
  console.error("Failed to start server:", error);
  process.exit(1);
});
