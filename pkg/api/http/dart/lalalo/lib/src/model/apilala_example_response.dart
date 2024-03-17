//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

// ignore_for_file: unused_element
import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'apilala_example_response.g.dart';

/// ApilalaExampleResponse
///
/// Properties:
/// * [title] 
/// * [description] 
/// * [body] 
@BuiltValue()
abstract class ApilalaExampleResponse implements Built<ApilalaExampleResponse, ApilalaExampleResponseBuilder> {
  @BuiltValueField(wireName: r'title')
  String? get title;

  @BuiltValueField(wireName: r'description')
  String? get description;

  @BuiltValueField(wireName: r'body')
  String? get body;

  ApilalaExampleResponse._();

  factory ApilalaExampleResponse([void updates(ApilalaExampleResponseBuilder b)]) = _$ApilalaExampleResponse;

  @BuiltValueHook(initializeBuilder: true)
  static void _defaults(ApilalaExampleResponseBuilder b) => b;

  @BuiltValueSerializer(custom: true)
  static Serializer<ApilalaExampleResponse> get serializer => _$ApilalaExampleResponseSerializer();
}

class _$ApilalaExampleResponseSerializer implements PrimitiveSerializer<ApilalaExampleResponse> {
  @override
  final Iterable<Type> types = const [ApilalaExampleResponse, _$ApilalaExampleResponse];

  @override
  final String wireName = r'ApilalaExampleResponse';

  Iterable<Object?> _serializeProperties(
    Serializers serializers,
    ApilalaExampleResponse object, {
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
    ApilalaExampleResponse object, {
    FullType specifiedType = FullType.unspecified,
  }) {
    return _serializeProperties(serializers, object, specifiedType: specifiedType).toList();
  }

  void _deserializeProperties(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
    required List<Object?> serializedList,
    required ApilalaExampleResponseBuilder result,
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
  ApilalaExampleResponse deserialize(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
  }) {
    final result = ApilalaExampleResponseBuilder();
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

