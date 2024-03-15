//
//  Generated code. Do not modify.
//  source: example/v2/examplev2.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'package:protobuf/protobuf.dart' as $pb;

import 'examplev2.pb.dart' as $1;

export 'examplev2.pb.dart';

@$pb.GrpcServiceName('api.example.v2.ExampleService')
class ExampleServiceClient extends $grpc.Client {
  static final _$example = $grpc.ClientMethod<$1.ExampleRequest, $1.ExampleResponse>(
      '/api.example.v2.ExampleService/Example',
      ($1.ExampleRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.ExampleResponse.fromBuffer(value));

  ExampleServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseFuture<$1.ExampleResponse> example($1.ExampleRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$example, request, options: options);
  }
}

@$pb.GrpcServiceName('api.example.v2.ExampleService')
abstract class ExampleServiceBase extends $grpc.Service {
  $core.String get $name => 'api.example.v2.ExampleService';

  ExampleServiceBase() {
    $addMethod($grpc.ServiceMethod<$1.ExampleRequest, $1.ExampleResponse>(
        'Example',
        example_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.ExampleRequest.fromBuffer(value),
        ($1.ExampleResponse value) => value.writeToBuffer()));
  }

  $async.Future<$1.ExampleResponse> example_Pre($grpc.ServiceCall call, $async.Future<$1.ExampleRequest> request) async {
    return example(call, await request);
  }

  $async.Future<$1.ExampleResponse> example($grpc.ServiceCall call, $1.ExampleRequest request);
}
