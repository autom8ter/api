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
  serialized_pb=_b('\n\tapi.proto\x12\x03\x61pi\x1a\x1cgoogle/api/annotations.proto\"\x1b\n\nGetByEmail\x12\r\n\x05\x65mail\x18\x01 \x01(\t\"\xc4\x01\n\x07IDToken\x12\x0b\n\x03iss\x18\x01 \x01(\t\x12\x0b\n\x03sub\x18\x02 \x01(\t\x12\x0b\n\x03\x61ud\x18\x03 \x01(\t\x12\x0b\n\x03\x65xp\x18\x04 \x01(\x03\x12\x0b\n\x03iat\x18\x05 \x01(\x03\x12\x0c\n\x04name\x18\x06 \x01(\t\x12\x12\n\ngiven_name\x18\x07 \x01(\t\x12\x13\n\x0b\x66\x61mily_name\x18\x08 \x01(\t\x12\x0e\n\x06gender\x18\t \x01(\t\x12\x11\n\tbirthdate\x18\n \x01(\t\x12\r\n\x05\x65mail\x18\x0b \x01(\t\x12\x0f\n\x07picture\x18\x0c \x01(\t\"T\n\x0cUserMetadata\x12\r\n\x05phone\x18\x01 \x01(\t\x12\x0c\n\x04plan\x18\x02 \x01(\t\x12\x11\n\tpay_token\x18\x03 \x01(\t\x12\x14\n\x0clast_contact\x18\x04 \x01(\t\"j\n\x0b\x41\x63\x63\x65ssToken\x12\x0b\n\x03iss\x18\x01 \x01(\t\x12\x0b\n\x03sub\x18\x02 \x01(\t\x12\x0b\n\x03\x61ud\x18\x03 \x03(\t\x12\x0b\n\x03\x61zp\x18\x04 \x01(\t\x12\x0b\n\x03\x65xp\x18\x05 \x01(\x03\x12\x0b\n\x03iat\x18\x06 \x01(\x03\x12\r\n\x05scope\x18\x07 \x01(\t2T\n\tIDService\x12G\n\nGetIDToken\x12\x0f.api.GetByEmail\x1a\x0c.api.IDToken\"\x1a\x82\xd3\xe4\x93\x02\x14\"\x0f/api/id/{email}:\x01*2`\n\rAccessService\x12O\n\nGetIDToken\x12\x0f.api.GetByEmail\x1a\x10.api.AccessToken\"\x1e\x82\xd3\xe4\x93\x02\x18\"\x13/api/access/{email}:\x01*b\x06proto3')
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,])




_GETBYEMAIL = _descriptor.Descriptor(
  name='GetByEmail',
  full_name='api.GetByEmail',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='email', full_name='api.GetByEmail.email', index=0,
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
  serialized_start=48,
  serialized_end=75,
)


_IDTOKEN = _descriptor.Descriptor(
  name='IDToken',
  full_name='api.IDToken',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='iss', full_name='api.IDToken.iss', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sub', full_name='api.IDToken.sub', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='aud', full_name='api.IDToken.aud', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='exp', full_name='api.IDToken.exp', index=3,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='iat', full_name='api.IDToken.iat', index=4,
      number=5, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='api.IDToken.name', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='given_name', full_name='api.IDToken.given_name', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='family_name', full_name='api.IDToken.family_name', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='gender', full_name='api.IDToken.gender', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='birthdate', full_name='api.IDToken.birthdate', index=9,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='email', full_name='api.IDToken.email', index=10,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='picture', full_name='api.IDToken.picture', index=11,
      number=12, type=9, cpp_type=9, label=1,
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
  serialized_start=78,
  serialized_end=274,
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
      name='plan', full_name='api.UserMetadata.plan', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='pay_token', full_name='api.UserMetadata.pay_token', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='last_contact', full_name='api.UserMetadata.last_contact', index=3,
      number=4, type=9, cpp_type=9, label=1,
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
  serialized_start=276,
  serialized_end=360,
)


_ACCESSTOKEN = _descriptor.Descriptor(
  name='AccessToken',
  full_name='api.AccessToken',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='iss', full_name='api.AccessToken.iss', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sub', full_name='api.AccessToken.sub', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='aud', full_name='api.AccessToken.aud', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='azp', full_name='api.AccessToken.azp', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='exp', full_name='api.AccessToken.exp', index=4,
      number=5, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='iat', full_name='api.AccessToken.iat', index=5,
      number=6, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='scope', full_name='api.AccessToken.scope', index=6,
      number=7, type=9, cpp_type=9, label=1,
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
  serialized_start=362,
  serialized_end=468,
)

DESCRIPTOR.message_types_by_name['GetByEmail'] = _GETBYEMAIL
DESCRIPTOR.message_types_by_name['IDToken'] = _IDTOKEN
DESCRIPTOR.message_types_by_name['UserMetadata'] = _USERMETADATA
DESCRIPTOR.message_types_by_name['AccessToken'] = _ACCESSTOKEN
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

GetByEmail = _reflection.GeneratedProtocolMessageType('GetByEmail', (_message.Message,), dict(
  DESCRIPTOR = _GETBYEMAIL,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.GetByEmail)
  ))
_sym_db.RegisterMessage(GetByEmail)

IDToken = _reflection.GeneratedProtocolMessageType('IDToken', (_message.Message,), dict(
  DESCRIPTOR = _IDTOKEN,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.IDToken)
  ))
_sym_db.RegisterMessage(IDToken)

UserMetadata = _reflection.GeneratedProtocolMessageType('UserMetadata', (_message.Message,), dict(
  DESCRIPTOR = _USERMETADATA,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.UserMetadata)
  ))
_sym_db.RegisterMessage(UserMetadata)

AccessToken = _reflection.GeneratedProtocolMessageType('AccessToken', (_message.Message,), dict(
  DESCRIPTOR = _ACCESSTOKEN,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.AccessToken)
  ))
_sym_db.RegisterMessage(AccessToken)



_IDSERVICE = _descriptor.ServiceDescriptor(
  name='IDService',
  full_name='api.IDService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=470,
  serialized_end=554,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetIDToken',
    full_name='api.IDService.GetIDToken',
    index=0,
    containing_service=None,
    input_type=_GETBYEMAIL,
    output_type=_IDTOKEN,
    serialized_options=_b('\202\323\344\223\002\024\"\017/api/id/{email}:\001*'),
  ),
])
_sym_db.RegisterServiceDescriptor(_IDSERVICE)

DESCRIPTOR.services_by_name['IDService'] = _IDSERVICE


_ACCESSSERVICE = _descriptor.ServiceDescriptor(
  name='AccessService',
  full_name='api.AccessService',
  file=DESCRIPTOR,
  index=1,
  serialized_options=None,
  serialized_start=556,
  serialized_end=652,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetIDToken',
    full_name='api.AccessService.GetIDToken',
    index=0,
    containing_service=None,
    input_type=_GETBYEMAIL,
    output_type=_ACCESSTOKEN,
    serialized_options=_b('\202\323\344\223\002\030\"\023/api/access/{email}:\001*'),
  ),
])
_sym_db.RegisterServiceDescriptor(_ACCESSSERVICE)

DESCRIPTOR.services_by_name['AccessService'] = _ACCESSSERVICE

# @@protoc_insertion_point(module_scope)
