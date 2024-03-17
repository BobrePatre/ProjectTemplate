//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

// ignore_for_file: unused_element
import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'apiexample_example_request.g.dart';

/// ApiexampleExampleRequest
///
/// Properties:
/// * [title] 
/// * [description] 
/// * [body] 
@BuiltValue()
abstract class ApiexampleExampleRequest implements Built<ApiexampleExampleRequest, ApiexampleExampleRequestBuilder> {
  @BuiltValueField(wireName: r'title')
  String? get title;

  @BuiltValueField(wireName: r'description')
  String? get description;

  @BuiltValueField(wireName: r'body')
  String? get body;

  ApiexampleExampleRequest._();

  factory ApiexampleExampleRequest([void updates(ApiexampleExampleRequestBuilder b)]) = _$ApiexampleExampleRequest;

  @BuiltValueHook(initializeBuilder: true)
  static void _defaults(ApiexampleExampleRequestBuilder b) => b;

  @BuiltValueSerializer(custom: true)
  static Serializer<ApiexampleExampleRequest> get serializer => _$ApiexampleExampleRequestSerializer();
}

class _$ApiexampleExampleRequestSerializer implements PrimitiveSerializer<ApiexampleExampleRequest> {
  @override
  final Iterable<Type> types = const [ApiexampleExampleRequest, _$ApiexampleExampleRequest];

  @override
  final String wireName = r'ApiexampleExampleRequest';

  Iterable<Object?> _serializeProperties(
    Serializers serializers,
    ApiexampleExampleRequest object, {
    FullType specifiedType = FullType.unspecified,
  }) sync* {
    if (object.title != null) {
      yield r'title';
      yield serializers.serialize(
        object.title,
        specifiedType: const FullType(String),
      );
    }
    if (object.description != null) {
      yield r'description';
      yield serializers.serialize(
        object.description,
        specifiedType: const FullType(String),
      );
    }
    if (object.body != null) {
      yield r'body';
      yield serializers.serialize(
        object.body,
        specifiedType: const FullType(String),
      );
    }
  }

  @override
  Object serialize(
    Serializers serializers,
    ApiexampleExampleRequest object, {
    FullType specifiedType = FullType.unspecified,
  }) {
    return _serializeProperties(serializers, object, specifiedType: specifiedType).toList();
  }

  void _deserializeProperties(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
    required List<Object?> serializedList,
    required ApiexampleExampleRequestBuilder result,
    required List<Object?> unhandled,
  }) {
    for (var i = 0; i < serializedList.length; i += 2) {
      final key = serializedList[i] as String;
      final value = serializedList[i + 1];
      switch (key) {
        case r'title':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.title = valueDes;
          break;
        case r'description':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.description = valueDes;
          break;
        case r'body':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.body = valueDes;
          break;
        default:
          unhandled.add(key);
          unhandled.add(value);
          break;
      }
    }
  }

  @override
  ApiexampleExampleRequest deserialize(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
  }) {
    final result = ApiexampleExampleRequestBuilder();
    final serializedList = (serialized as Iterable<Object?>).toList();
    final unhandled = <Object?>[];
    _deserializeProperties(
      serializers,
      serialized,
      specifiedType: specifiedType,
      serializedList: serializedList,
      unhandled: unhandled,
      result: result,
    );
    return result.build();
  }
}

