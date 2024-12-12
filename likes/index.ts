import * as grpc from '@grpc/grpc-js';
import * as protoLoader from '@grpc/proto-loader';
import * as path from 'path';

// Path to the .proto file in the sibling directory 'common'
const PROTO_PATH = path.resolve(__dirname, '../common/proto/blog.proto');

// Load the protobuf definition
const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

const grpcObject = grpc.loadPackageDefinition(packageDefinition) as any;
const LikeService = grpcObject.proto.LikeService;

// In-memory data to simulate a database
const likes: Array<{ PostId: number; UserId: number }> = [];

// Implement the CreateLike method
const createLike = (
  call: grpc.ServerUnaryCall<{ PostId: number }, { PostId: number; UserId: number }>,
  callback: grpc.sendUnaryData<{ PostId: number; UserId: number }>
) => {
  const { PostId } = call.request;
  const UserId = Math.floor(Math.random() * 1000); // Simulated user ID
  const like = { PostId, UserId };
  likes.push(like);
  console.log(`Like created: ${JSON.stringify(like)}`);
  callback(null, like);
};

// Create the gRPC server
const server = new grpc.Server();
server.addService(LikeService.service, { CreateLike: createLike });

// Start the server
const PORT = 50051;
server.bindAsync(`0.0.0.0:${PORT}`, grpc.ServerCredentials.createInsecure(), () => {
  console.log(`gRPC server is running on http://0.0.0.0:${PORT}`);
  server.start();
});
