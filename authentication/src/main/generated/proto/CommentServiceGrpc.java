package proto;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.69.0)",
    comments = "Source: blog.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class CommentServiceGrpc {

  private CommentServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "proto.CommentService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<proto.Blog.CreateCommentRequest,
      proto.Blog.Comment> getCreateCommentMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateComment",
      requestType = proto.Blog.CreateCommentRequest.class,
      responseType = proto.Blog.Comment.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<proto.Blog.CreateCommentRequest,
      proto.Blog.Comment> getCreateCommentMethod() {
    io.grpc.MethodDescriptor<proto.Blog.CreateCommentRequest, proto.Blog.Comment> getCreateCommentMethod;
    if ((getCreateCommentMethod = CommentServiceGrpc.getCreateCommentMethod) == null) {
      synchronized (CommentServiceGrpc.class) {
        if ((getCreateCommentMethod = CommentServiceGrpc.getCreateCommentMethod) == null) {
          CommentServiceGrpc.getCreateCommentMethod = getCreateCommentMethod =
              io.grpc.MethodDescriptor.<proto.Blog.CreateCommentRequest, proto.Blog.Comment>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateComment"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.Blog.CreateCommentRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.Blog.Comment.getDefaultInstance()))
              .setSchemaDescriptor(new CommentServiceMethodDescriptorSupplier("CreateComment"))
              .build();
        }
      }
    }
    return getCreateCommentMethod;
  }

  private static volatile io.grpc.MethodDescriptor<proto.Blog.GetCommentsRequest,
      proto.Blog.GetCommentsResponse> getGetCommentsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetComments",
      requestType = proto.Blog.GetCommentsRequest.class,
      responseType = proto.Blog.GetCommentsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<proto.Blog.GetCommentsRequest,
      proto.Blog.GetCommentsResponse> getGetCommentsMethod() {
    io.grpc.MethodDescriptor<proto.Blog.GetCommentsRequest, proto.Blog.GetCommentsResponse> getGetCommentsMethod;
    if ((getGetCommentsMethod = CommentServiceGrpc.getGetCommentsMethod) == null) {
      synchronized (CommentServiceGrpc.class) {
        if ((getGetCommentsMethod = CommentServiceGrpc.getGetCommentsMethod) == null) {
          CommentServiceGrpc.getGetCommentsMethod = getGetCommentsMethod =
              io.grpc.MethodDescriptor.<proto.Blog.GetCommentsRequest, proto.Blog.GetCommentsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetComments"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.Blog.GetCommentsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.Blog.GetCommentsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommentServiceMethodDescriptorSupplier("GetComments"))
              .build();
        }
      }
    }
    return getGetCommentsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<proto.Blog.UpdateCommentRequest,
      proto.Blog.Comment> getUpdateCommentMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdateComment",
      requestType = proto.Blog.UpdateCommentRequest.class,
      responseType = proto.Blog.Comment.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<proto.Blog.UpdateCommentRequest,
      proto.Blog.Comment> getUpdateCommentMethod() {
    io.grpc.MethodDescriptor<proto.Blog.UpdateCommentRequest, proto.Blog.Comment> getUpdateCommentMethod;
    if ((getUpdateCommentMethod = CommentServiceGrpc.getUpdateCommentMethod) == null) {
      synchronized (CommentServiceGrpc.class) {
        if ((getUpdateCommentMethod = CommentServiceGrpc.getUpdateCommentMethod) == null) {
          CommentServiceGrpc.getUpdateCommentMethod = getUpdateCommentMethod =
              io.grpc.MethodDescriptor.<proto.Blog.UpdateCommentRequest, proto.Blog.Comment>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdateComment"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.Blog.UpdateCommentRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.Blog.Comment.getDefaultInstance()))
              .setSchemaDescriptor(new CommentServiceMethodDescriptorSupplier("UpdateComment"))
              .build();
        }
      }
    }
    return getUpdateCommentMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static CommentServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<CommentServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<CommentServiceStub>() {
        @java.lang.Override
        public CommentServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new CommentServiceStub(channel, callOptions);
        }
      };
    return CommentServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static CommentServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<CommentServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<CommentServiceBlockingStub>() {
        @java.lang.Override
        public CommentServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new CommentServiceBlockingStub(channel, callOptions);
        }
      };
    return CommentServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static CommentServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<CommentServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<CommentServiceFutureStub>() {
        @java.lang.Override
        public CommentServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new CommentServiceFutureStub(channel, callOptions);
        }
      };
    return CommentServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     */
    default void createComment(proto.Blog.CreateCommentRequest request,
        io.grpc.stub.StreamObserver<proto.Blog.Comment> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateCommentMethod(), responseObserver);
    }

    /**
     */
    default void getComments(proto.Blog.GetCommentsRequest request,
        io.grpc.stub.StreamObserver<proto.Blog.GetCommentsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetCommentsMethod(), responseObserver);
    }

    /**
     */
    default void updateComment(proto.Blog.UpdateCommentRequest request,
        io.grpc.stub.StreamObserver<proto.Blog.Comment> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateCommentMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service CommentService.
   */
  public static abstract class CommentServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return CommentServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service CommentService.
   */
  public static final class CommentServiceStub
      extends io.grpc.stub.AbstractAsyncStub<CommentServiceStub> {
    private CommentServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CommentServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new CommentServiceStub(channel, callOptions);
    }

    /**
     */
    public void createComment(proto.Blog.CreateCommentRequest request,
        io.grpc.stub.StreamObserver<proto.Blog.Comment> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateCommentMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getComments(proto.Blog.GetCommentsRequest request,
        io.grpc.stub.StreamObserver<proto.Blog.GetCommentsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetCommentsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateComment(proto.Blog.UpdateCommentRequest request,
        io.grpc.stub.StreamObserver<proto.Blog.Comment> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateCommentMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service CommentService.
   */
  public static final class CommentServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<CommentServiceBlockingStub> {
    private CommentServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CommentServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new CommentServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public proto.Blog.Comment createComment(proto.Blog.CreateCommentRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateCommentMethod(), getCallOptions(), request);
    }

    /**
     */
    public proto.Blog.GetCommentsResponse getComments(proto.Blog.GetCommentsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetCommentsMethod(), getCallOptions(), request);
    }

    /**
     */
    public proto.Blog.Comment updateComment(proto.Blog.UpdateCommentRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateCommentMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service CommentService.
   */
  public static final class CommentServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<CommentServiceFutureStub> {
    private CommentServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CommentServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new CommentServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<proto.Blog.Comment> createComment(
        proto.Blog.CreateCommentRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateCommentMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<proto.Blog.GetCommentsResponse> getComments(
        proto.Blog.GetCommentsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetCommentsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<proto.Blog.Comment> updateComment(
        proto.Blog.UpdateCommentRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateCommentMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_CREATE_COMMENT = 0;
  private static final int METHODID_GET_COMMENTS = 1;
  private static final int METHODID_UPDATE_COMMENT = 2;

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
        case METHODID_CREATE_COMMENT:
          serviceImpl.createComment((proto.Blog.CreateCommentRequest) request,
              (io.grpc.stub.StreamObserver<proto.Blog.Comment>) responseObserver);
          break;
        case METHODID_GET_COMMENTS:
          serviceImpl.getComments((proto.Blog.GetCommentsRequest) request,
              (io.grpc.stub.StreamObserver<proto.Blog.GetCommentsResponse>) responseObserver);
          break;
        case METHODID_UPDATE_COMMENT:
          serviceImpl.updateComment((proto.Blog.UpdateCommentRequest) request,
              (io.grpc.stub.StreamObserver<proto.Blog.Comment>) responseObserver);
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
          getCreateCommentMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              proto.Blog.CreateCommentRequest,
              proto.Blog.Comment>(
                service, METHODID_CREATE_COMMENT)))
        .addMethod(
          getGetCommentsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              proto.Blog.GetCommentsRequest,
              proto.Blog.GetCommentsResponse>(
                service, METHODID_GET_COMMENTS)))
        .addMethod(
          getUpdateCommentMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              proto.Blog.UpdateCommentRequest,
              proto.Blog.Comment>(
                service, METHODID_UPDATE_COMMENT)))
        .build();
  }

  private static abstract class CommentServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    CommentServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return proto.Blog.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("CommentService");
    }
  }

  private static final class CommentServiceFileDescriptorSupplier
      extends CommentServiceBaseDescriptorSupplier {
    CommentServiceFileDescriptorSupplier() {}
  }

  private static final class CommentServiceMethodDescriptorSupplier
      extends CommentServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    CommentServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (CommentServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new CommentServiceFileDescriptorSupplier())
              .addMethod(getCreateCommentMethod())
              .addMethod(getGetCommentsMethod())
              .addMethod(getUpdateCommentMethod())
              .build();
        }
      }
    }
    return result;
  }
}
