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
  serialized_pb=_b('\n\tapi.proto\x12\x03\x61pi\x1a\x1cgoogle/api/annotations.proto\"\x18\n\nIdentifier\x12\n\n\x02id\x18\x01 \x01(\t\"\\\n\tSMSStatus\x12\x1b\n\x02id\x18\x01 \x01(\x0b\x32\x0f.api.Identifier\x12\x15\n\x03sms\x18\x02 \x01(\x0b\x32\x08.api.SMS\x12\x0e\n\x06status\x18\x03 \x01(\t\x12\x0b\n\x03uri\x18\x04 \x01(\t\"r\n\x03SMS\x12\x0f\n\x07service\x18\x01 \x01(\t\x12\n\n\x02to\x18\x02 \x01(\t\x12\x1d\n\x07message\x18\x03 \x01(\x0b\x32\x0c.api.Message\x12\x10\n\x08mediaURL\x18\x04 \x01(\t\x12\x10\n\x08\x63\x61llback\x18\x05 \x01(\t\x12\x0b\n\x03\x61pp\x18\x06 \x01(\t\"P\n\x0c\x45mailRequest\x12\x11\n\tfrom_name\x18\x01 \x01(\t\x12\x12\n\nfrom_email\x18\x02 \x01(\t\x12\x19\n\x05\x65mail\x18\x03 \x01(\x0b\x32\n.api.Email\"T\n\x05\x45mail\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0f\n\x07\x61\x64\x64ress\x18\x02 \x01(\t\x12\x0f\n\x07subject\x18\x03 \x01(\t\x12\r\n\x05plain\x18\x04 \x01(\t\x12\x0c\n\x04html\x18\x05 \x01(\t\"-\n\x04\x43\x61ll\x12\x0c\n\x04\x66rom\x18\x01 \x01(\t\x12\n\n\x02to\x18\x02 \x01(\t\x12\x0b\n\x03\x61pp\x18\x03 \x01(\t\"\x18\n\x07Message\x12\r\n\x05value\x18\x01 \x01(\t\"\xd6\x01\n\x08UserInfo\x12\x0c\n\x04name\x18\x06 \x01(\t\x12\x12\n\ngiven_name\x18\x07 \x01(\t\x12\x13\n\x0b\x66\x61mily_name\x18\x08 \x01(\t\x12\x0e\n\x06gender\x18\t \x01(\t\x12\x11\n\tbirthdate\x18\n \x01(\t\x12\r\n\x05\x65mail\x18\x0b \x01(\t\x12\x0f\n\x07picture\x18\x0c \x01(\t\x12(\n\ruser_metadata\x18\r \x01(\x0b\x32\x11.api.UserMetadata\x12&\n\x0c\x61pp_metadata\x18\x0e \x01(\x0b\x32\x10.api.AppMetadata\"V\n\x0cUserMetadata\x12\r\n\x05phone\x18\x01 \x01(\t\x12\x19\n\x11preferred_contact\x18\x02 \x01(\t\x12\x0e\n\x06status\x18\x03 \x01(\t\x12\x0c\n\x04tags\x18\x04 \x03(\t\"P\n\x0b\x41ppMetadata\x12\x0c\n\x04plan\x18\x01 \x01(\t\x12\x11\n\tpay_token\x18\x02 \x01(\t\x12\x12\n\ndelinquent\x18\x03 \x01(\t\x12\x0c\n\x04tags\x18\x04 \x03(\t\"t\n\x04\x41uth\x12\x0e\n\x06\x64omain\x18\x01 \x01(\t\x12\x11\n\tclient_id\x18\x02 \x01(\t\x12\x15\n\rclient_secret\x18\x03 \x01(\t\x12\x10\n\x08redirect\x18\x04 \x01(\t\x12\x10\n\x08\x61udience\x18\x05 \x01(\t\x12\x0e\n\x06scopes\x18\x06 \x03(\t\"\x15\n\x05\x42ytes\x12\x0c\n\x04\x62its\x18\x01 \x01(\x0c\"&\n\x08Template\x12\x0c\n\x04text\x18\x01 \x01(\t\x12\x0c\n\x04\x64\x61ta\x18\x02 \x01(\x0c\x32\xe8\x05\n\x0eUtilityService\x12<\n\x04\x45\x63ho\x12\x0c.api.Message\x1a\x0c.api.Message\"\x18\x82\xd3\xe4\x93\x02\x12\"\r/utility/echo:\x01*\x12K\n\x0b\x45\x63hoSpanish\x12\x0c.api.Message\x1a\x0c.api.Message\" \x82\xd3\xe4\x93\x02\x1a\"\x15/utility/echo/spanish:\x01*\x12K\n\x0b\x45\x63hoChinese\x12\x0c.api.Message\x1a\x0c.api.Message\" \x82\xd3\xe4\x93\x02\x1a\"\x15/utility/echo/chinese:\x01*\x12K\n\x0b\x45\x63hoEnglish\x12\x0c.api.Message\x1a\x0c.api.Message\" \x82\xd3\xe4\x93\x02\x1a\"\x15/utility/echo/english:\x01*\x12G\n\tEchoHindi\x12\x0c.api.Message\x1a\x0c.api.Message\"\x1e\x82\xd3\xe4\x93\x02\x18\"\x13/utility/echo/hindi:\x01*\x12I\n\nEchoArabic\x12\x0c.api.Message\x1a\x0c.api.Message\"\x1f\x82\xd3\xe4\x93\x02\x19\"\x14/utility/echo/arabic:\x01*\x12G\n\x0bMarshalJSON\x12\n.api.Bytes\x1a\n.api.Bytes\" \x82\xd3\xe4\x93\x02\x1a\"\x15/utility/marshal/json:\x01*\x12G\n\x0bMarshalYAML\x12\n.api.Bytes\x1a\n.api.Bytes\" \x82\xd3\xe4\x93\x02\x1a\"\x15/utility/marshal/yaml:\x01*\x12\x45\n\nMarshalXML\x12\n.api.Bytes\x1a\n.api.Bytes\"\x1f\x82\xd3\xe4\x93\x02\x19\"\x14/utility/marshal/xml:\x01*\x12\x44\n\x06Render\x12\r.api.Template\x1a\n.api.Bytes\"\x1f\x82\xd3\xe4\x93\x02\x19\"\x14/utility/marshal/xml:\x01*2\xbd\x02\n\x0e\x43ontactService\x12\x42\n\x07SendSMS\x12\x08.api.SMS\x1a\x0f.api.Identifier\"\x1c\x82\xd3\xe4\x93\x02\x16\"\x11/contact/sms/{to}:\x01*\x12G\n\x06GetSMS\x12\x0f.api.Identifier\x1a\x0e.api.SMSStatus\"\x1c\x82\xd3\xe4\x93\x02\x16\"\x11/contact/sms/{id}:\x01*\x12W\n\tSendEmail\x12\x11.api.EmailRequest\x1a\x0c.api.Message\")\x82\xd3\xe4\x93\x02#\"\x1e/contact/email/{email.address}:\x01*\x12\x45\n\x08SendCall\x12\t.api.Call\x1a\x0f.api.Identifier\"\x1d\x82\xd3\xe4\x93\x02\x17\"\x12/contact/call/{to}:\x01*b\x06proto3')
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,])




_IDENTIFIER = _descriptor.Descriptor(
  name='Identifier',
  full_name='api.Identifier',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='api.Identifier.id', index=0,
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
  serialized_end=72,
)


_SMSSTATUS = _descriptor.Descriptor(
  name='SMSStatus',
  full_name='api.SMSStatus',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='api.SMSStatus.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sms', full_name='api.SMSStatus.sms', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='status', full_name='api.SMSStatus.status', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='uri', full_name='api.SMSStatus.uri', index=3,
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
  serialized_start=74,
  serialized_end=166,
)


_SMS = _descriptor.Descriptor(
  name='SMS',
  full_name='api.SMS',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='service', full_name='api.SMS.service', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='to', full_name='api.SMS.to', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='message', full_name='api.SMS.message', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='mediaURL', full_name='api.SMS.mediaURL', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='callback', full_name='api.SMS.callback', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='app', full_name='api.SMS.app', index=5,
      number=6, type=9, cpp_type=9, label=1,
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
  serialized_start=168,
  serialized_end=282,
)


_EMAILREQUEST = _descriptor.Descriptor(
  name='EmailRequest',
  full_name='api.EmailRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='from_name', full_name='api.EmailRequest.from_name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='from_email', full_name='api.EmailRequest.from_email', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='email', full_name='api.EmailRequest.email', index=2,
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
  serialized_start=284,
  serialized_end=364,
)


_EMAIL = _descriptor.Descriptor(
  name='Email',
  full_name='api.Email',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='api.Email.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='address', full_name='api.Email.address', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='subject', full_name='api.Email.subject', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='plain', full_name='api.Email.plain', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='html', full_name='api.Email.html', index=4,
      number=5, type=9, cpp_type=9, label=1,
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
  serialized_start=366,
  serialized_end=450,
)


_CALL = _descriptor.Descriptor(
  name='Call',
  full_name='api.Call',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='from', full_name='api.Call.from', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='to', full_name='api.Call.to', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='app', full_name='api.Call.app', index=2,
      number=3, type=9, cpp_type=9, label=1,
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
  serialized_start=452,
  serialized_end=497,
)


_MESSAGE = _descriptor.Descriptor(
  name='Message',
  full_name='api.Message',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='api.Message.value', index=0,
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
  serialized_start=499,
  serialized_end=523,
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
  serialized_start=526,
  serialized_end=740,
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
    _descriptor.FieldDescriptor(
      name='status', full_name='api.UserMetadata.status', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tags', full_name='api.UserMetadata.tags', index=3,
      number=4, type=9, cpp_type=9, label=3,
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
  serialized_start=742,
  serialized_end=828,
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
    _descriptor.FieldDescriptor(
      name='delinquent', full_name='api.AppMetadata.delinquent', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tags', full_name='api.AppMetadata.tags', index=3,
      number=4, type=9, cpp_type=9, label=3,
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
  serialized_start=830,
  serialized_end=910,
)


_AUTH = _descriptor.Descriptor(
  name='Auth',
  full_name='api.Auth',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='domain', full_name='api.Auth.domain', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='client_id', full_name='api.Auth.client_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='client_secret', full_name='api.Auth.client_secret', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='redirect', full_name='api.Auth.redirect', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='audience', full_name='api.Auth.audience', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='scopes', full_name='api.Auth.scopes', index=5,
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
  serialized_start=912,
  serialized_end=1028,
)


_BYTES = _descriptor.Descriptor(
  name='Bytes',
  full_name='api.Bytes',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='bits', full_name='api.Bytes.bits', index=0,
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
  serialized_start=1030,
  serialized_end=1051,
)


_TEMPLATE = _descriptor.Descriptor(
  name='Template',
  full_name='api.Template',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='text', full_name='api.Template.text', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='data', full_name='api.Template.data', index=1,
      number=2, type=12, cpp_type=9, label=1,
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
  serialized_start=1053,
  serialized_end=1091,
)

_SMSSTATUS.fields_by_name['id'].message_type = _IDENTIFIER
_SMSSTATUS.fields_by_name['sms'].message_type = _SMS
_SMS.fields_by_name['message'].message_type = _MESSAGE
_EMAILREQUEST.fields_by_name['email'].message_type = _EMAIL
_USERINFO.fields_by_name['user_metadata'].message_type = _USERMETADATA
_USERINFO.fields_by_name['app_metadata'].message_type = _APPMETADATA
DESCRIPTOR.message_types_by_name['Identifier'] = _IDENTIFIER
DESCRIPTOR.message_types_by_name['SMSStatus'] = _SMSSTATUS
DESCRIPTOR.message_types_by_name['SMS'] = _SMS
DESCRIPTOR.message_types_by_name['EmailRequest'] = _EMAILREQUEST
DESCRIPTOR.message_types_by_name['Email'] = _EMAIL
DESCRIPTOR.message_types_by_name['Call'] = _CALL
DESCRIPTOR.message_types_by_name['Message'] = _MESSAGE
DESCRIPTOR.message_types_by_name['UserInfo'] = _USERINFO
DESCRIPTOR.message_types_by_name['UserMetadata'] = _USERMETADATA
DESCRIPTOR.message_types_by_name['AppMetadata'] = _APPMETADATA
DESCRIPTOR.message_types_by_name['Auth'] = _AUTH
DESCRIPTOR.message_types_by_name['Bytes'] = _BYTES
DESCRIPTOR.message_types_by_name['Template'] = _TEMPLATE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Identifier = _reflection.GeneratedProtocolMessageType('Identifier', (_message.Message,), dict(
  DESCRIPTOR = _IDENTIFIER,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.Identifier)
  ))
_sym_db.RegisterMessage(Identifier)

SMSStatus = _reflection.GeneratedProtocolMessageType('SMSStatus', (_message.Message,), dict(
  DESCRIPTOR = _SMSSTATUS,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.SMSStatus)
  ))
_sym_db.RegisterMessage(SMSStatus)

SMS = _reflection.GeneratedProtocolMessageType('SMS', (_message.Message,), dict(
  DESCRIPTOR = _SMS,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.SMS)
  ))
_sym_db.RegisterMessage(SMS)

EmailRequest = _reflection.GeneratedProtocolMessageType('EmailRequest', (_message.Message,), dict(
  DESCRIPTOR = _EMAILREQUEST,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.EmailRequest)
  ))
_sym_db.RegisterMessage(EmailRequest)

Email = _reflection.GeneratedProtocolMessageType('Email', (_message.Message,), dict(
  DESCRIPTOR = _EMAIL,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.Email)
  ))
_sym_db.RegisterMessage(Email)

Call = _reflection.GeneratedProtocolMessageType('Call', (_message.Message,), dict(
  DESCRIPTOR = _CALL,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.Call)
  ))
_sym_db.RegisterMessage(Call)

Message = _reflection.GeneratedProtocolMessageType('Message', (_message.Message,), dict(
  DESCRIPTOR = _MESSAGE,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.Message)
  ))
_sym_db.RegisterMessage(Message)

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

Auth = _reflection.GeneratedProtocolMessageType('Auth', (_message.Message,), dict(
  DESCRIPTOR = _AUTH,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.Auth)
  ))
_sym_db.RegisterMessage(Auth)

Bytes = _reflection.GeneratedProtocolMessageType('Bytes', (_message.Message,), dict(
  DESCRIPTOR = _BYTES,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.Bytes)
  ))
_sym_db.RegisterMessage(Bytes)

Template = _reflection.GeneratedProtocolMessageType('Template', (_message.Message,), dict(
  DESCRIPTOR = _TEMPLATE,
  __module__ = 'api_pb2'
  # @@protoc_insertion_point(class_scope:api.Template)
  ))
_sym_db.RegisterMessage(Template)



_UTILITYSERVICE = _descriptor.ServiceDescriptor(
  name='UtilityService',
  full_name='api.UtilityService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=1094,
  serialized_end=1838,
  methods=[
  _descriptor.MethodDescriptor(
    name='Echo',
    full_name='api.UtilityService.Echo',
    index=0,
    containing_service=None,
    input_type=_MESSAGE,
    output_type=_MESSAGE,
    serialized_options=_b('\202\323\344\223\002\022\"\r/utility/echo:\001*'),
  ),
  _descriptor.MethodDescriptor(
    name='EchoSpanish',
    full_name='api.UtilityService.EchoSpanish',
    index=1,
    containing_service=None,
    input_type=_MESSAGE,
    output_type=_MESSAGE,
    serialized_options=_b('\202\323\344\223\002\032\"\025/utility/echo/spanish:\001*'),
  ),
  _descriptor.MethodDescriptor(
    name='EchoChinese',
    full_name='api.UtilityService.EchoChinese',
    index=2,
    containing_service=None,
    input_type=_MESSAGE,
    output_type=_MESSAGE,
    serialized_options=_b('\202\323\344\223\002\032\"\025/utility/echo/chinese:\001*'),
  ),
  _descriptor.MethodDescriptor(
    name='EchoEnglish',
    full_name='api.UtilityService.EchoEnglish',
    index=3,
    containing_service=None,
    input_type=_MESSAGE,
    output_type=_MESSAGE,
    serialized_options=_b('\202\323\344\223\002\032\"\025/utility/echo/english:\001*'),
  ),
  _descriptor.MethodDescriptor(
    name='EchoHindi',
    full_name='api.UtilityService.EchoHindi',
    index=4,
    containing_service=None,
    input_type=_MESSAGE,
    output_type=_MESSAGE,
    serialized_options=_b('\202\323\344\223\002\030\"\023/utility/echo/hindi:\001*'),
  ),
  _descriptor.MethodDescriptor(
    name='EchoArabic',
    full_name='api.UtilityService.EchoArabic',
    index=5,
    containing_service=None,
    input_type=_MESSAGE,
    output_type=_MESSAGE,
    serialized_options=_b('\202\323\344\223\002\031\"\024/utility/echo/arabic:\001*'),
  ),
  _descriptor.MethodDescriptor(
    name='MarshalJSON',
    full_name='api.UtilityService.MarshalJSON',
    index=6,
    containing_service=None,
    input_type=_BYTES,
    output_type=_BYTES,
    serialized_options=_b('\202\323\344\223\002\032\"\025/utility/marshal/json:\001*'),
  ),
  _descriptor.MethodDescriptor(
    name='MarshalYAML',
    full_name='api.UtilityService.MarshalYAML',
    index=7,
    containing_service=None,
    input_type=_BYTES,
    output_type=_BYTES,
    serialized_options=_b('\202\323\344\223\002\032\"\025/utility/marshal/yaml:\001*'),
  ),
  _descriptor.MethodDescriptor(
    name='MarshalXML',
    full_name='api.UtilityService.MarshalXML',
    index=8,
    containing_service=None,
    input_type=_BYTES,
    output_type=_BYTES,
    serialized_options=_b('\202\323\344\223\002\031\"\024/utility/marshal/xml:\001*'),
  ),
  _descriptor.MethodDescriptor(
    name='Render',
    full_name='api.UtilityService.Render',
    index=9,
    containing_service=None,
    input_type=_TEMPLATE,
    output_type=_BYTES,
    serialized_options=_b('\202\323\344\223\002\031\"\024/utility/marshal/xml:\001*'),
  ),
])
_sym_db.RegisterServiceDescriptor(_UTILITYSERVICE)

DESCRIPTOR.services_by_name['UtilityService'] = _UTILITYSERVICE


_CONTACTSERVICE = _descriptor.ServiceDescriptor(
  name='ContactService',
  full_name='api.ContactService',
  file=DESCRIPTOR,
  index=1,
  serialized_options=None,
  serialized_start=1841,
  serialized_end=2158,
  methods=[
  _descriptor.MethodDescriptor(
    name='SendSMS',
    full_name='api.ContactService.SendSMS',
    index=0,
    containing_service=None,
    input_type=_SMS,
    output_type=_IDENTIFIER,
    serialized_options=_b('\202\323\344\223\002\026\"\021/contact/sms/{to}:\001*'),
  ),
  _descriptor.MethodDescriptor(
    name='GetSMS',
    full_name='api.ContactService.GetSMS',
    index=1,
    containing_service=None,
    input_type=_IDENTIFIER,
    output_type=_SMSSTATUS,
    serialized_options=_b('\202\323\344\223\002\026\"\021/contact/sms/{id}:\001*'),
  ),
  _descriptor.MethodDescriptor(
    name='SendEmail',
    full_name='api.ContactService.SendEmail',
    index=2,
    containing_service=None,
    input_type=_EMAILREQUEST,
    output_type=_MESSAGE,
    serialized_options=_b('\202\323\344\223\002#\"\036/contact/email/{email.address}:\001*'),
  ),
  _descriptor.MethodDescriptor(
    name='SendCall',
    full_name='api.ContactService.SendCall',
    index=3,
    containing_service=None,
    input_type=_CALL,
    output_type=_IDENTIFIER,
    serialized_options=_b('\202\323\344\223\002\027\"\022/contact/call/{to}:\001*'),
  ),
])
_sym_db.RegisterServiceDescriptor(_CONTACTSERVICE)

DESCRIPTOR.services_by_name['ContactService'] = _CONTACTSERVICE

# @@protoc_insertion_point(module_scope)
