//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

// ignore_for_file: unused_element
import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'apilala_example_request.g.dart';

/// ApilalaExampleRequest
///
/// Properties:
/// * [title] 
/// * [description] 
/// * [body] 
@BuiltValue()
abstract class ApilalaExampleRequest implements Built<ApilalaExampleRequest, ApilalaExampleRequestBuilder> {
  @BuiltValueField(wireName: r'title')
  String? get title;

  @BuiltValueField(wireName: r'description')
  String? get description;

  @BuiltValueField(wireName: r'body')
  String? get body;

  ApilalaExampleRequest._();

  factory ApilalaExampleRequest([void updates(ApilalaExampleRequestBuilder b)]) = _$ApilalaExampleRequest;

  @BuiltValueHook(initializeBuilder: true)
  static void _defaults(ApilalaExampleRequestBuilder b) => b;

  @BuiltValueSerializer(custom: true)
  static Serializer<ApilalaExampleRequest> get serializer => _$ApilalaExampleRequestSerializer();
}

class _$ApilalaExampleRequestSerializer implements PrimitiveSerializer<ApilalaExampleRequest> {
  @override
  final Iterable<Type> types = const [ApilalaExampleRequest, _$ApilalaExampleRequest];

  @override
  final String wireName = r'ApilalaExampleRequest';

  Iterable<Object?> _serializeProperties(
    Serializers serializers,
    ApilalaExampleRequest object, {
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
    ApilalaExampleRequest object, {
    FullType specifiedType = FullType.unspecified,
  }) {
    return _serializeProperties(serializers, object, specifiedType: specifiedType).toList();
  }

  void _deserializeProperties(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
    required List<Object?> serializedList,
    required ApilalaExampleRequestBuilder result,
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
  ApilalaExampleRequest deserialize(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
  }) {
    final result = ApilalaExampleRequestBuilder();
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

