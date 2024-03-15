//
//  Generated code. Do not modify.
//  source: example/v2/examplev2.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use exampleRequestDescriptor instead')
const ExampleRequest$json = {
  '1': 'ExampleRequest',
  '2': [
    {'1': 'title', '3': 1, '4': 1, '5': 9, '10': 'title'},
    {'1': 'description', '3': 2, '4': 1, '5': 9, '10': 'description'},
    {'1': 'body', '3': 3, '4': 1, '5': 9, '10': 'body'},
  ],
};

/// Descriptor for `ExampleRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List exampleRequestDescriptor = $convert.base64Decode(
    'Cg5FeGFtcGxlUmVxdWVzdBIUCgV0aXRsZRgBIAEoCVIFdGl0bGUSIAoLZGVzY3JpcHRpb24YAi'
    'ABKAlSC2Rlc2NyaXB0aW9uEhIKBGJvZHkYAyABKAlSBGJvZHk=');

@$core.Deprecated('Use exampleResponseDescriptor instead')
const ExampleResponse$json = {
  '1': 'ExampleResponse',
  '2': [
    {'1': 'title', '3': 1, '4': 1, '5': 9, '10': 'title'},
    {'1': 'description', '3': 2, '4': 1, '5': 9, '10': 'description'},
    {'1': 'body', '3': 3, '4': 1, '5': 9, '10': 'body'},
  ],
};

/// Descriptor for `ExampleResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List exampleResponseDescriptor = $convert.base64Decode(
    'Cg9FeGFtcGxlUmVzcG9uc2USFAoFdGl0bGUYASABKAlSBXRpdGxlEiAKC2Rlc2NyaXB0aW9uGA'
    'IgASgJUgtkZXNjcmlwdGlvbhISCgRib2R5GAMgASgJUgRib2R5');

