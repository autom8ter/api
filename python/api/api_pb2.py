# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='api.proto',
  package='api',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\tapi.proto\x12\x03\x61pi\"\xd6\x01\n\x08UserInfo\x12\x0c\n\x04name\x18\x06 \x01(\t\x12\x12\n\ngiven_name\x18\x07 \x01(\t\x12\x13\n\x0b\x66\x61mily_name\x18\x08 \x01(\t\x12\x0e\n\x06gender\x18\t \x01(\t\x12\x11\n\tbirthdate\x18\n \x01(\t\x12\r\n\x05\x65mail\x18\x0b \x01(\t\x12\x0f\n\x07picture\x18\x0c \x01(\t\x12(\n\ruser_metadata\x18\r \x01(\x0b\x32\x11.api.UserMetadata\x12&\n\x0c\x61pp_metadata\x18\x0e \x01(\x0b\x32\x10.api.AppMetadata\"8\n\x0cUserMetadata\x12\r\n\x05phone\x18\x01 \x01(\t\x12\x19\n\x11preferred_contact\x18\x02 \x01(\t\".\n\x0b\x41ppMetadata\x12\x0c\n\x04plan\x18\x01 \x01(\t\x12\x11\n\tpay_token\x18\x02 \x01(\t\"v\n\x06\x43onfig\x12\x0e\n\x06\x64omain\x18\x01 \x01(\t\x12\x11\n\tclient_id\x18\x02 \x01(\t\x12\x15\n\rclient_secret\x18\x03 \x01(\t\x12\x10\n\x08redirect\x18\x04 \x01(\t\x12\x10\n\x08\x61udience\x18\x05 \x01(\t\x12\x0e\n\x06scopes\x18\x06 \x03(\tb\x06proto3')
)




_USERINFO = _descriptor.Descriptor(
  name='UserInfo',
  full_name='api.UserInfo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='api.UserInfo.name', index=0,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='given_name', full_name='api.UserInfo.given_name', index=1,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='family_name', full_name='api.UserInfo.family_name', index=2,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='gender', full_name='api.UserInfo.gender', index=3,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='birthdate', full_name='api.UserInfo.birthdate', index=4,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='email', full_name='api.UserInfo.email', index=5,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='picture', full_name='api.UserInfo.picture', index=6,
      number=12, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='user_metadata', full_name='api.UserInfo.user_metadata', index=7,
      number=13, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='app_metadata', full_name='api.UserInfo.app_metadata', index=8,
      number=14, type=11, cpp_type=10, label=1,
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
  serialized_start=19,
  serialized_end=233,
)


_USERMETADATA = _descriptor.Descriptor(
  name='UserMetadata',
  full_name='api.UserMetadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='phone', full_name='api.UserMetadata.phone', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='preferred_contact', full_name='api.UserMetadata.preferred_contact', index=1,
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
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=235,
  serialized_end=291,
)


_APPMETADATA = _descriptor.Descriptor(
  name='AppMetadata',
  full_name='api.AppMetadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='plan', full_name='api.AppMetadata.plan', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='pay_token', full_name='api.AppMetadata.pay_token', index=1,
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
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=293,
  serialized_end=339,
)


_CONFIG = _descriptor.Descriptor(
  name='Config',
  full_name='api.Config',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='domain', full_name='api.Config.domain', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='client_id', full_name='api.Config.client_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='client_secret', full_name='api.Config.client_secret', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='redirect', full_name='api.Config.redirect', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='audience', full_name='api.Config.audience', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='scopes', full_name='api.Config.scopes', index=5,
      number=6, type=9, cpp_type=9, label=3,
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
  serialized_start=341,
  serialized_end=459,
)

_USERINFO.fields_by_name['user_metadata'].message_type = _USERMETADATA
_USERINFO.fields_by_name['app_metadata'].message_type = _APPMETADATA
DESCRIPTOR.message_types_by_name['UserInfo'] = _USERINFO
DESCRIPTOR.message_types_by_name['UserMetadata'] = _USERMETADATA
DESCRIPTOR.message_types_by_name['AppMetadata'] = _APPMETADATA
DESCRIPTOR.message_types_by_name['Config'] = _CONFIG
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

UserInfo = _reflection.GeneratedProtocolMessageType('UserInfo', (_message.Message,), dict(
  DESCRIPTOR = _USERINFO,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.UserInfo)
  ))
_sym_db.RegisterMessage(UserInfo)

UserMetadata = _reflection.GeneratedProtocolMessageType('UserMetadata', (_message.Message,), dict(
  DESCRIPTOR = _USERMETADATA,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.UserMetadata)
  ))
_sym_db.RegisterMessage(UserMetadata)

AppMetadata = _reflection.GeneratedProtocolMessageType('AppMetadata', (_message.Message,), dict(
  DESCRIPTOR = _APPMETADATA,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.AppMetadata)
  ))
_sym_db.RegisterMessage(AppMetadata)

Config = _reflection.GeneratedProtocolMessageType('Config', (_message.Message,), dict(
  DESCRIPTOR = _CONFIG,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.Config)
  ))
_sym_db.RegisterMessage(Config)


# @@protoc_insertion_point(module_scope)
