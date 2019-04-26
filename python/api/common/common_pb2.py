# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: common/common.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='common/common.proto',
  package='common',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x13\x63ommon/common.proto\x12\x06\x63ommon\"\x16\n\x06String\x12\x0c\n\x04text\x18\x01 \x01(\t\"\x15\n\x05\x42ytes\x12\x0c\n\x04\x62its\x18\x01 \x01(\x0c\"\x16\n\x04\x42ool\x12\x0e\n\x06\x61nswer\x18\x01 \x01(\x08\"\x1e\n\x0bStringArray\x12\x0f\n\x07strings\x18\x01 \x03(\t\"s\n\tStringMap\x12\x34\n\nstring_map\x18\x01 \x03(\x0b\x32 .common.StringMap.StringMapEntry\x1a\x30\n\x0eStringMapEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\x07\n\x05\x45mpty\"\x18\n\nIdentifier\x12\n\n\x02id\x18\x01 \x01(\t\"\x18\n\x07Message\x12\r\n\x05value\x18\x01 \x01(\t\"\x16\n\x06Secret\x12\x0c\n\x04text\x18\x01 \x01(\t\"\x17\n\x05Query\x12\x0e\n\x06lucene\x18\x01 \x01(\t\"\x18\n\x08Password\x12\x0c\n\x04text\x18\x01 \x01(\t\"\x14\n\x05Int64\x12\x0b\n\x03num\x18\x01 \x01(\x03\"\x14\n\x05Int32\x12\x0b\n\x03num\x18\x01 \x01(\x05\"\x16\n\x07\x46loat64\x12\x0b\n\x03num\x18\x01 \x01(\x01\"\x8c\x01\n\x0bHTTPRequest\x12\"\n\x06method\x18\x01 \x01(\x0e\x32\x12.common.HTTPMethod\x12\x1b\n\x03url\x18\x02 \x01(\x0b\x32\x0e.common.String\x12\x1f\n\x04\x66orm\x18\x03 \x01(\x0b\x32\x11.common.StringMap\x12\x1b\n\x04\x62ody\x18\x04 \x01(\x0b\x32\r.common.Bytes\"F\n\x08Template\x12\x1c\n\x04name\x18\x01 \x01(\x0b\x32\x0e.common.String\x12\x1c\n\x04text\x18\x02 \x01(\x0b\x32\x0e.common.String\"\xba\x01\n\x05Token\x12$\n\x0c\x61\x63\x63\x65ss_token\x18\x01 \x01(\x0b\x32\x0e.common.String\x12\"\n\ntoken_type\x18\x02 \x01(\x0b\x32\x0e.common.String\x12%\n\rrefresh_token\x18\x03 \x01(\x0b\x32\x0e.common.String\x12\x1e\n\x06\x65xpiry\x18\x04 \x01(\x0b\x32\x0e.common.String\x12 \n\x08id_token\x18\x05 \x01(\x0b\x32\x0e.common.String**\n\nHTTPMethod\x12\x07\n\x03GET\x10\x00\x12\x08\n\x04POST\x10\x01\x12\t\n\x05PATCH\x10\x02\x62\x06proto3')
)

_HTTPMETHOD = _descriptor.EnumDescriptor(
  name='HTTPMethod',
  full_name='common.HTTPMethod',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='GET', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='POST', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PATCH', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=859,
  serialized_end=901,
)
_sym_db.RegisterEnumDescriptor(_HTTPMETHOD)

HTTPMethod = enum_type_wrapper.EnumTypeWrapper(_HTTPMETHOD)
GET = 0
POST = 1
PATCH = 2



_STRING = _descriptor.Descriptor(
  name='String',
  full_name='common.String',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='text', full_name='common.String.text', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=31,
  serialized_end=53,
)


_BYTES = _descriptor.Descriptor(
  name='Bytes',
  full_name='common.Bytes',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='bits', full_name='common.Bytes.bits', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=55,
  serialized_end=76,
)


_BOOL = _descriptor.Descriptor(
  name='Bool',
  full_name='common.Bool',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='answer', full_name='common.Bool.answer', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=78,
  serialized_end=100,
)


_STRINGARRAY = _descriptor.Descriptor(
  name='StringArray',
  full_name='common.StringArray',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='strings', full_name='common.StringArray.strings', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=102,
  serialized_end=132,
)


_STRINGMAP_STRINGMAPENTRY = _descriptor.Descriptor(
  name='StringMapEntry',
  full_name='common.StringMap.StringMapEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='common.StringMap.StringMapEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='common.StringMap.StringMapEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=201,
  serialized_end=249,
)

_STRINGMAP = _descriptor.Descriptor(
  name='StringMap',
  full_name='common.StringMap',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='string_map', full_name='common.StringMap.string_map', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_STRINGMAP_STRINGMAPENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=134,
  serialized_end=249,
)


_EMPTY = _descriptor.Descriptor(
  name='Empty',
  full_name='common.Empty',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=251,
  serialized_end=258,
)


_IDENTIFIER = _descriptor.Descriptor(
  name='Identifier',
  full_name='common.Identifier',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='common.Identifier.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=260,
  serialized_end=284,
)


_MESSAGE = _descriptor.Descriptor(
  name='Message',
  full_name='common.Message',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='common.Message.value', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=286,
  serialized_end=310,
)


_SECRET = _descriptor.Descriptor(
  name='Secret',
  full_name='common.Secret',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='text', full_name='common.Secret.text', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=312,
  serialized_end=334,
)


_QUERY = _descriptor.Descriptor(
  name='Query',
  full_name='common.Query',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='lucene', full_name='common.Query.lucene', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=336,
  serialized_end=359,
)


_PASSWORD = _descriptor.Descriptor(
  name='Password',
  full_name='common.Password',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='text', full_name='common.Password.text', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=361,
  serialized_end=385,
)


_INT64 = _descriptor.Descriptor(
  name='Int64',
  full_name='common.Int64',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='num', full_name='common.Int64.num', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=387,
  serialized_end=407,
)


_INT32 = _descriptor.Descriptor(
  name='Int32',
  full_name='common.Int32',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='num', full_name='common.Int32.num', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=409,
  serialized_end=429,
)


_FLOAT64 = _descriptor.Descriptor(
  name='Float64',
  full_name='common.Float64',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='num', full_name='common.Float64.num', index=0,
      number=1, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=431,
  serialized_end=453,
)


_HTTPREQUEST = _descriptor.Descriptor(
  name='HTTPRequest',
  full_name='common.HTTPRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='method', full_name='common.HTTPRequest.method', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='url', full_name='common.HTTPRequest.url', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='form', full_name='common.HTTPRequest.form', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='body', full_name='common.HTTPRequest.body', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=456,
  serialized_end=596,
)


_TEMPLATE = _descriptor.Descriptor(
  name='Template',
  full_name='common.Template',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='common.Template.name', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='text', full_name='common.Template.text', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=598,
  serialized_end=668,
)


_TOKEN = _descriptor.Descriptor(
  name='Token',
  full_name='common.Token',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='access_token', full_name='common.Token.access_token', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='token_type', full_name='common.Token.token_type', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='refresh_token', full_name='common.Token.refresh_token', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='expiry', full_name='common.Token.expiry', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='id_token', full_name='common.Token.id_token', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=671,
  serialized_end=857,
)

_STRINGMAP_STRINGMAPENTRY.containing_type = _STRINGMAP
_STRINGMAP.fields_by_name['string_map'].message_type = _STRINGMAP_STRINGMAPENTRY
_HTTPREQUEST.fields_by_name['method'].enum_type = _HTTPMETHOD
_HTTPREQUEST.fields_by_name['url'].message_type = _STRING
_HTTPREQUEST.fields_by_name['form'].message_type = _STRINGMAP
_HTTPREQUEST.fields_by_name['body'].message_type = _BYTES
_TEMPLATE.fields_by_name['name'].message_type = _STRING
_TEMPLATE.fields_by_name['text'].message_type = _STRING
_TOKEN.fields_by_name['access_token'].message_type = _STRING
_TOKEN.fields_by_name['token_type'].message_type = _STRING
_TOKEN.fields_by_name['refresh_token'].message_type = _STRING
_TOKEN.fields_by_name['expiry'].message_type = _STRING
_TOKEN.fields_by_name['id_token'].message_type = _STRING
DESCRIPTOR.message_types_by_name['String'] = _STRING
DESCRIPTOR.message_types_by_name['Bytes'] = _BYTES
DESCRIPTOR.message_types_by_name['Bool'] = _BOOL
DESCRIPTOR.message_types_by_name['StringArray'] = _STRINGARRAY
DESCRIPTOR.message_types_by_name['StringMap'] = _STRINGMAP
DESCRIPTOR.message_types_by_name['Empty'] = _EMPTY
DESCRIPTOR.message_types_by_name['Identifier'] = _IDENTIFIER
DESCRIPTOR.message_types_by_name['Message'] = _MESSAGE
DESCRIPTOR.message_types_by_name['Secret'] = _SECRET
DESCRIPTOR.message_types_by_name['Query'] = _QUERY
DESCRIPTOR.message_types_by_name['Password'] = _PASSWORD
DESCRIPTOR.message_types_by_name['Int64'] = _INT64
DESCRIPTOR.message_types_by_name['Int32'] = _INT32
DESCRIPTOR.message_types_by_name['Float64'] = _FLOAT64
DESCRIPTOR.message_types_by_name['HTTPRequest'] = _HTTPREQUEST
DESCRIPTOR.message_types_by_name['Template'] = _TEMPLATE
DESCRIPTOR.message_types_by_name['Token'] = _TOKEN
DESCRIPTOR.enum_types_by_name['HTTPMethod'] = _HTTPMETHOD
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

String = _reflection.GeneratedProtocolMessageType('String', (_message.Message,), dict(
  DESCRIPTOR = _STRING,
  __module__ = 'common.common_pb2'
  # @@protoc_insertion_point(class_scope:common.String)
  ))
_sym_db.RegisterMessage(String)

Bytes = _reflection.GeneratedProtocolMessageType('Bytes', (_message.Message,), dict(
  DESCRIPTOR = _BYTES,
  __module__ = 'common.common_pb2'
  # @@protoc_insertion_point(class_scope:common.Bytes)
  ))
_sym_db.RegisterMessage(Bytes)

Bool = _reflection.GeneratedProtocolMessageType('Bool', (_message.Message,), dict(
  DESCRIPTOR = _BOOL,
  __module__ = 'common.common_pb2'
  # @@protoc_insertion_point(class_scope:common.Bool)
  ))
_sym_db.RegisterMessage(Bool)

StringArray = _reflection.GeneratedProtocolMessageType('StringArray', (_message.Message,), dict(
  DESCRIPTOR = _STRINGARRAY,
  __module__ = 'common.common_pb2'
  # @@protoc_insertion_point(class_scope:common.StringArray)
  ))
_sym_db.RegisterMessage(StringArray)

StringMap = _reflection.GeneratedProtocolMessageType('StringMap', (_message.Message,), dict(

  StringMapEntry = _reflection.GeneratedProtocolMessageType('StringMapEntry', (_message.Message,), dict(
    DESCRIPTOR = _STRINGMAP_STRINGMAPENTRY,
    __module__ = 'common.common_pb2'
    # @@protoc_insertion_point(class_scope:common.StringMap.StringMapEntry)
    ))
  ,
  DESCRIPTOR = _STRINGMAP,
  __module__ = 'common.common_pb2'
  # @@protoc_insertion_point(class_scope:common.StringMap)
  ))
_sym_db.RegisterMessage(StringMap)
_sym_db.RegisterMessage(StringMap.StringMapEntry)

Empty = _reflection.GeneratedProtocolMessageType('Empty', (_message.Message,), dict(
  DESCRIPTOR = _EMPTY,
  __module__ = 'common.common_pb2'
  # @@protoc_insertion_point(class_scope:common.Empty)
  ))
_sym_db.RegisterMessage(Empty)

Identifier = _reflection.GeneratedProtocolMessageType('Identifier', (_message.Message,), dict(
  DESCRIPTOR = _IDENTIFIER,
  __module__ = 'common.common_pb2'
  # @@protoc_insertion_point(class_scope:common.Identifier)
  ))
_sym_db.RegisterMessage(Identifier)

Message = _reflection.GeneratedProtocolMessageType('Message', (_message.Message,), dict(
  DESCRIPTOR = _MESSAGE,
  __module__ = 'common.common_pb2'
  # @@protoc_insertion_point(class_scope:common.Message)
  ))
_sym_db.RegisterMessage(Message)

Secret = _reflection.GeneratedProtocolMessageType('Secret', (_message.Message,), dict(
  DESCRIPTOR = _SECRET,
  __module__ = 'common.common_pb2'
  # @@protoc_insertion_point(class_scope:common.Secret)
  ))
_sym_db.RegisterMessage(Secret)

Query = _reflection.GeneratedProtocolMessageType('Query', (_message.Message,), dict(
  DESCRIPTOR = _QUERY,
  __module__ = 'common.common_pb2'
  # @@protoc_insertion_point(class_scope:common.Query)
  ))
_sym_db.RegisterMessage(Query)

Password = _reflection.GeneratedProtocolMessageType('Password', (_message.Message,), dict(
  DESCRIPTOR = _PASSWORD,
  __module__ = 'common.common_pb2'
  # @@protoc_insertion_point(class_scope:common.Password)
  ))
_sym_db.RegisterMessage(Password)

Int64 = _reflection.GeneratedProtocolMessageType('Int64', (_message.Message,), dict(
  DESCRIPTOR = _INT64,
  __module__ = 'common.common_pb2'
  # @@protoc_insertion_point(class_scope:common.Int64)
  ))
_sym_db.RegisterMessage(Int64)

Int32 = _reflection.GeneratedProtocolMessageType('Int32', (_message.Message,), dict(
  DESCRIPTOR = _INT32,
  __module__ = 'common.common_pb2'
  # @@protoc_insertion_point(class_scope:common.Int32)
  ))
_sym_db.RegisterMessage(Int32)

Float64 = _reflection.GeneratedProtocolMessageType('Float64', (_message.Message,), dict(
  DESCRIPTOR = _FLOAT64,
  __module__ = 'common.common_pb2'
  # @@protoc_insertion_point(class_scope:common.Float64)
  ))
_sym_db.RegisterMessage(Float64)

HTTPRequest = _reflection.GeneratedProtocolMessageType('HTTPRequest', (_message.Message,), dict(
  DESCRIPTOR = _HTTPREQUEST,
  __module__ = 'common.common_pb2'
  # @@protoc_insertion_point(class_scope:common.HTTPRequest)
  ))
_sym_db.RegisterMessage(HTTPRequest)

Template = _reflection.GeneratedProtocolMessageType('Template', (_message.Message,), dict(
  DESCRIPTOR = _TEMPLATE,
  __module__ = 'common.common_pb2'
  # @@protoc_insertion_point(class_scope:common.Template)
  ))
_sym_db.RegisterMessage(Template)

Token = _reflection.GeneratedProtocolMessageType('Token', (_message.Message,), dict(
  DESCRIPTOR = _TOKEN,
  __module__ = 'common.common_pb2'
  # @@protoc_insertion_point(class_scope:common.Token)
  ))
_sym_db.RegisterMessage(Token)


_STRINGMAP_STRINGMAPENTRY._options = None
# @@protoc_insertion_point(module_scope)
