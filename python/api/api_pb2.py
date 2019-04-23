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


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='api.proto',
  package='api',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\tapi.proto\x12\x03\x61pi\x1a\x1cgoogle/api/annotations.proto\"\xd6\x01\n\x08UserInfo\x12\x0c\n\x04name\x18\x06 \x01(\t\x12\x12\n\ngiven_name\x18\x07 \x01(\t\x12\x13\n\x0b\x66\x61mily_name\x18\x08 \x01(\t\x12\x0e\n\x06gender\x18\t \x01(\t\x12\x11\n\tbirthdate\x18\n \x01(\t\x12\r\n\x05\x65mail\x18\x0b \x01(\t\x12\x0f\n\x07picture\x18\x0c \x01(\t\x12(\n\ruser_metadata\x18\r \x01(\x0b\x32\x11.api.UserMetadata\x12&\n\x0c\x61pp_metadata\x18\x0e \x01(\x0b\x32\x10.api.AppMetadata\"8\n\x0cUserMetadata\x12\r\n\x05phone\x18\x01 \x01(\t\x12\x19\n\x11preferred_contact\x18\x02 \x01(\t\".\n\x0b\x41ppMetadata\x12\x0c\n\x04plan\x18\x01 \x01(\t\x12\x11\n\tpay_token\x18\x02 \x01(\t\"\x1c\n\x0b\x41\x63\x63\x65ssToken\x12\r\n\x05token\x18\x01 \x01(\t\"\x18\n\x07IDToken\x12\r\n\x05token\x18\x01 \x01(\t\"\x1d\n\x0cRefreshToken\x12\r\n\x05token\x18\x01 \x01(\t\"h\n\x06Tokens\x12\x18\n\x02id\x18\x01 \x01(\x0b\x32\x0c.api.IDToken\x12 \n\x06\x61\x63\x63\x65ss\x18\x02 \x01(\x0b\x32\x10.api.AccessToken\x12\"\n\x07refresh\x18\x03 \x01(\x0b\x32\x11.api.RefreshToken\"\xee\x01\n\x05Paths\x12\x0c\n\x04home\x18\x01 \x01(\t\x12\x11\n\tdashboard\x18\x02 \x01(\t\x12\x10\n\x08settings\x18\x03 \x01(\t\x12\x0e\n\x06logout\x18\x04 \x01(\t\x12\x10\n\x08\x63\x61llback\x18\x05 \x01(\t\x12\r\n\x05login\x18\x06 \x01(\t\x12\x11\n\tsubscribe\x18\x07 \x01(\t\x12\x13\n\x0bunsubscribe\x18\x08 \x01(\t\x12\x0b\n\x03\x66\x61q\x18\t \x01(\t\x12\x0f\n\x07support\x18\n \x01(\t\x12\r\n\x05terms\x18\x0b \x01(\t\x12\x0f\n\x07privacy\x18\x0c \x01(\t\x12\r\n\x05\x64\x65\x62ug\x18\r \x01(\t\x12\x0c\n\x04\x62log\x18\x0e \x01(\t\"9\n\x15\x43reateCustomerRequest\x12 \n\tuser_info\x18\x01 \x01(\x0b\x32\r.api.UserInfo\"$\n\x16\x43reateCustomerResponse\x12\n\n\x02id\x18\x01 \x01(\t2v\n\x0f\x43ustomerService\x12\x63\n\x0e\x43reateCustomer\x12\x1a.api.CreateCustomerRequest\x1a\x1b.api.CreateCustomerResponse\"\x18\x82\xd3\xe4\x93\x02\x12\"\r/v1/customers:\x01*b\x06proto3')
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,])




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
  serialized_start=49,
  serialized_end=263,
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
  serialized_start=265,
  serialized_end=321,
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
  serialized_start=323,
  serialized_end=369,
)


_ACCESSTOKEN = _descriptor.Descriptor(
  name='AccessToken',
  full_name='api.AccessToken',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='token', full_name='api.AccessToken.token', index=0,
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
  serialized_start=371,
  serialized_end=399,
)


_IDTOKEN = _descriptor.Descriptor(
  name='IDToken',
  full_name='api.IDToken',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='token', full_name='api.IDToken.token', index=0,
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
  serialized_start=401,
  serialized_end=425,
)


_REFRESHTOKEN = _descriptor.Descriptor(
  name='RefreshToken',
  full_name='api.RefreshToken',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='token', full_name='api.RefreshToken.token', index=0,
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
  serialized_start=427,
  serialized_end=456,
)


_TOKENS = _descriptor.Descriptor(
  name='Tokens',
  full_name='api.Tokens',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='api.Tokens.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='access', full_name='api.Tokens.access', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='refresh', full_name='api.Tokens.refresh', index=2,
      number=3, type=11, cpp_type=10, label=1,
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
  serialized_start=458,
  serialized_end=562,
)


_PATHS = _descriptor.Descriptor(
  name='Paths',
  full_name='api.Paths',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='home', full_name='api.Paths.home', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dashboard', full_name='api.Paths.dashboard', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='settings', full_name='api.Paths.settings', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='logout', full_name='api.Paths.logout', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='callback', full_name='api.Paths.callback', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='login', full_name='api.Paths.login', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='subscribe', full_name='api.Paths.subscribe', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='unsubscribe', full_name='api.Paths.unsubscribe', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='faq', full_name='api.Paths.faq', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='support', full_name='api.Paths.support', index=9,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='terms', full_name='api.Paths.terms', index=10,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='privacy', full_name='api.Paths.privacy', index=11,
      number=12, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='debug', full_name='api.Paths.debug', index=12,
      number=13, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='blog', full_name='api.Paths.blog', index=13,
      number=14, type=9, cpp_type=9, label=1,
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
  serialized_start=565,
  serialized_end=803,
)


_CREATECUSTOMERREQUEST = _descriptor.Descriptor(
  name='CreateCustomerRequest',
  full_name='api.CreateCustomerRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_info', full_name='api.CreateCustomerRequest.user_info', index=0,
      number=1, type=11, cpp_type=10, label=1,
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
  serialized_start=805,
  serialized_end=862,
)


_CREATECUSTOMERRESPONSE = _descriptor.Descriptor(
  name='CreateCustomerResponse',
  full_name='api.CreateCustomerResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='api.CreateCustomerResponse.id', index=0,
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
  serialized_start=864,
  serialized_end=900,
)

_USERINFO.fields_by_name['user_metadata'].message_type = _USERMETADATA
_USERINFO.fields_by_name['app_metadata'].message_type = _APPMETADATA
_TOKENS.fields_by_name['id'].message_type = _IDTOKEN
_TOKENS.fields_by_name['access'].message_type = _ACCESSTOKEN
_TOKENS.fields_by_name['refresh'].message_type = _REFRESHTOKEN
_CREATECUSTOMERREQUEST.fields_by_name['user_info'].message_type = _USERINFO
DESCRIPTOR.message_types_by_name['UserInfo'] = _USERINFO
DESCRIPTOR.message_types_by_name['UserMetadata'] = _USERMETADATA
DESCRIPTOR.message_types_by_name['AppMetadata'] = _APPMETADATA
DESCRIPTOR.message_types_by_name['AccessToken'] = _ACCESSTOKEN
DESCRIPTOR.message_types_by_name['IDToken'] = _IDTOKEN
DESCRIPTOR.message_types_by_name['RefreshToken'] = _REFRESHTOKEN
DESCRIPTOR.message_types_by_name['Tokens'] = _TOKENS
DESCRIPTOR.message_types_by_name['Paths'] = _PATHS
DESCRIPTOR.message_types_by_name['CreateCustomerRequest'] = _CREATECUSTOMERREQUEST
DESCRIPTOR.message_types_by_name['CreateCustomerResponse'] = _CREATECUSTOMERRESPONSE
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

AccessToken = _reflection.GeneratedProtocolMessageType('AccessToken', (_message.Message,), dict(
  DESCRIPTOR = _ACCESSTOKEN,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.AccessToken)
  ))
_sym_db.RegisterMessage(AccessToken)

IDToken = _reflection.GeneratedProtocolMessageType('IDToken', (_message.Message,), dict(
  DESCRIPTOR = _IDTOKEN,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.IDToken)
  ))
_sym_db.RegisterMessage(IDToken)

RefreshToken = _reflection.GeneratedProtocolMessageType('RefreshToken', (_message.Message,), dict(
  DESCRIPTOR = _REFRESHTOKEN,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.RefreshToken)
  ))
_sym_db.RegisterMessage(RefreshToken)

Tokens = _reflection.GeneratedProtocolMessageType('Tokens', (_message.Message,), dict(
  DESCRIPTOR = _TOKENS,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.Tokens)
  ))
_sym_db.RegisterMessage(Tokens)

Paths = _reflection.GeneratedProtocolMessageType('Paths', (_message.Message,), dict(
  DESCRIPTOR = _PATHS,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.Paths)
  ))
_sym_db.RegisterMessage(Paths)

CreateCustomerRequest = _reflection.GeneratedProtocolMessageType('CreateCustomerRequest', (_message.Message,), dict(
  DESCRIPTOR = _CREATECUSTOMERREQUEST,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.CreateCustomerRequest)
  ))
_sym_db.RegisterMessage(CreateCustomerRequest)

CreateCustomerResponse = _reflection.GeneratedProtocolMessageType('CreateCustomerResponse', (_message.Message,), dict(
  DESCRIPTOR = _CREATECUSTOMERRESPONSE,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.CreateCustomerResponse)
  ))
_sym_db.RegisterMessage(CreateCustomerResponse)



_CUSTOMERSERVICE = _descriptor.ServiceDescriptor(
  name='CustomerService',
  full_name='api.CustomerService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=902,
  serialized_end=1020,
  methods=[
  _descriptor.MethodDescriptor(
    name='CreateCustomer',
    full_name='api.CustomerService.CreateCustomer',
    index=0,
    containing_service=None,
    input_type=_CREATECUSTOMERREQUEST,
    output_type=_CREATECUSTOMERRESPONSE,
    serialized_options=_b('\202\323\344\223\002\022\"\r/v1/customers:\001*'),
  ),
])
_sym_db.RegisterServiceDescriptor(_CUSTOMERSERVICE)

DESCRIPTOR.services_by_name['CustomerService'] = _CUSTOMERSERVICE

# @@protoc_insertion_point(module_scope)
