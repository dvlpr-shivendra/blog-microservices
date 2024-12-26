package proto;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.69.0)",
    comments = "Source: blog.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class LikeServiceGrpc {

  private LikeServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "proto.LikeService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<proto.Blog.CreateLikeRequest,
      proto.Blog.CreateLikeResponse> getCreateLikeMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateLike",
      requestType = proto.Blog.CreateLikeRequest.class,
      responseType = proto.Blog.CreateLikeResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<proto.Blog.CreateLikeRequest,
      proto.Blog.CreateLikeResponse> getCreateLikeMethod() {
    io.grpc.MethodDescriptor<proto.Blog.CreateLikeRequest, proto.Blog.CreateLikeResponse> getCreateLikeMethod;
    if ((getCreateLikeMethod = LikeServiceGrpc.getCreateLikeMethod) == null) {
      synchronized (LikeServiceGrpc.class) {
        if ((getCreateLikeMethod = LikeServiceGrpc.getCreateLikeMethod) == null) {
          LikeServiceGrpc.getCreateLikeMethod = getCreateLikeMethod =
              io.grpc.MethodDescriptor.<proto.Blog.CreateLikeRequest, proto.Blog.CreateLikeResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateLike"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.Blog.CreateLikeRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.Blog.CreateLikeResponse.getDefaultInstance()))
              .setSchemaDescriptor(new LikeServiceMethodDescriptorSupplier("CreateLike"))
              .build();
        }
      }
    }
    return getCreateLikeMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static LikeServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<LikeServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<LikeServiceStub>() {
        @java.lang.Override
        public LikeServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new LikeServiceStub(channel, callOptions);
        }
      };
    return LikeServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static LikeServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<LikeServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<LikeServiceBlockingStub>() {
        @java.lang.Override
        public LikeServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new LikeServiceBlockingStub(channel, callOptions);
        }
      };
    return LikeServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static LikeServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<LikeServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<LikeServiceFutureStub>() {
        @java.lang.Override
        public LikeServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new LikeServiceFutureStub(channel, callOptions);
        }
      };
    return LikeServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     */
    default void createLike(proto.Blog.CreateLikeRequest request,
        io.grpc.stub.StreamObserver<proto.Blog.CreateLikeResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateLikeMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service LikeService.
   */
  public static abstract class LikeServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return LikeServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service LikeService.
   */
  public static final class LikeServiceStub
      extends io.grpc.stub.AbstractAsyncStub<LikeServiceStub> {
    private LikeServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected LikeServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new LikeServiceStub(channel, callOptions);
    }

    /**
     */
    public void createLike(proto.Blog.CreateLikeRequest request,
        io.grpc.stub.StreamObserver<proto.Blog.CreateLikeResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateLikeMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service LikeService.
   */
  public static final class LikeServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<LikeServiceBlockingStub> {
    private LikeServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected LikeServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new LikeServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public proto.Blog.CreateLikeResponse createLike(proto.Blog.CreateLikeRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateLikeMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service LikeService.
   */
  public static final class LikeServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<LikeServiceFutureStub> {
    private LikeServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected LikeServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new LikeServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<proto.Blog.CreateLikeResponse> createLike(
        proto.Blog.CreateLikeRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateLikeMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_CREATE_LIKE = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final AsyncService serviceImpl;
    private final int methodId;

    MethodHandlers(AsyncService serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_CREATE_LIKE:
          serviceImpl.createLike((proto.Blog.CreateLikeRequest) request,
              (io.grpc.stub.StreamObserver<proto.Blog.CreateLikeResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  public static final io.grpc.ServerServiceDefinition bindService(AsyncService service) {
    return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
        .addMethod(
          getCreateLikeMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              proto.Blog.CreateLikeRequest,
              proto.Blog.CreateLikeResponse>(
                service, METHODID_CREATE_LIKE)))
        .build();
  }

  private static abstract class LikeServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    LikeServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return proto.Blog.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("LikeService");
    }
  }

  private static final class LikeServiceFileDescriptorSupplier
      extends LikeServiceBaseDescriptorSupplier {
    LikeServiceFileDescriptorSupplier() {}
  }

  private static final class LikeServiceMethodDescriptorSupplier
      extends LikeServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    LikeServiceMethodDescriptorSupplier(java.lang.String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (LikeServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new LikeServiceFileDescriptorSupplier())
              .addMethod(getCreateLikeMethod())
              .build();
        }
      }
    }
    return result;
  }
}
