# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: analyzer.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import plugin_pb2 as plugin__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='analyzer.proto',
  package='pulumirpc',
  syntax='proto3',
  serialized_pb=_b('\n\x0e\x61nalyzer.proto\x12\tpulumirpc\x1a\x0cplugin.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\"K\n\x0e\x41nalyzeRequest\x12\x0c\n\x04type\x18\x01 \x01(\t\x12+\n\nproperties\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\">\n\x0f\x41nalyzeResponse\x12+\n\x08\x66\x61ilures\x18\x01 \x03(\x0b\x32\x19.pulumirpc.AnalyzeFailure\"2\n\x0e\x41nalyzeFailure\x12\x10\n\x08property\x18\x01 \x01(\t\x12\x0e\n\x06reason\x18\x02 \x01(\t2\x90\x01\n\x08\x41nalyzer\x12\x42\n\x07\x41nalyze\x12\x19.pulumirpc.AnalyzeRequest\x1a\x1a.pulumirpc.AnalyzeResponse\"\x00\x12@\n\rGetPluginInfo\x12\x16.google.protobuf.Empty\x1a\x15.pulumirpc.PluginInfo\"\x00\x62\x06proto3')
  ,
  dependencies=[plugin__pb2.DESCRIPTOR,google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,google_dot_protobuf_dot_struct__pb2.DESCRIPTOR,])




_ANALYZEREQUEST = _descriptor.Descriptor(
  name='AnalyzeRequest',
  full_name='pulumirpc.AnalyzeRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='pulumirpc.AnalyzeRequest.type', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='properties', full_name='pulumirpc.AnalyzeRequest.properties', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=102,
  serialized_end=177,
)


_ANALYZERESPONSE = _descriptor.Descriptor(
  name='AnalyzeResponse',
  full_name='pulumirpc.AnalyzeResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='failures', full_name='pulumirpc.AnalyzeResponse.failures', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=179,
  serialized_end=241,
)


_ANALYZEFAILURE = _descriptor.Descriptor(
  name='AnalyzeFailure',
  full_name='pulumirpc.AnalyzeFailure',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='property', full_name='pulumirpc.AnalyzeFailure.property', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='reason', full_name='pulumirpc.AnalyzeFailure.reason', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=243,
  serialized_end=293,
)

_ANALYZEREQUEST.fields_by_name['properties'].message_type = google_dot_protobuf_dot_struct__pb2._STRUCT
_ANALYZERESPONSE.fields_by_name['failures'].message_type = _ANALYZEFAILURE
DESCRIPTOR.message_types_by_name['AnalyzeRequest'] = _ANALYZEREQUEST
DESCRIPTOR.message_types_by_name['AnalyzeResponse'] = _ANALYZERESPONSE
DESCRIPTOR.message_types_by_name['AnalyzeFailure'] = _ANALYZEFAILURE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

AnalyzeRequest = _reflection.GeneratedProtocolMessageType('AnalyzeRequest', (_message.Message,), dict(
  DESCRIPTOR = _ANALYZEREQUEST,
  __module__ = 'analyzer_pb2'
  # @@protoc_insertion_point(class_scope:pulumirpc.AnalyzeRequest)
  ))
_sym_db.RegisterMessage(AnalyzeRequest)

AnalyzeResponse = _reflection.GeneratedProtocolMessageType('AnalyzeResponse', (_message.Message,), dict(
  DESCRIPTOR = _ANALYZERESPONSE,
  __module__ = 'analyzer_pb2'
  # @@protoc_insertion_point(class_scope:pulumirpc.AnalyzeResponse)
  ))
_sym_db.RegisterMessage(AnalyzeResponse)

AnalyzeFailure = _reflection.GeneratedProtocolMessageType('AnalyzeFailure', (_message.Message,), dict(
  DESCRIPTOR = _ANALYZEFAILURE,
  __module__ = 'analyzer_pb2'
  # @@protoc_insertion_point(class_scope:pulumirpc.AnalyzeFailure)
  ))
_sym_db.RegisterMessage(AnalyzeFailure)



_ANALYZER = _descriptor.ServiceDescriptor(
  name='Analyzer',
  full_name='pulumirpc.Analyzer',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=296,
  serialized_end=440,
  methods=[
  _descriptor.MethodDescriptor(
    name='Analyze',
    full_name='pulumirpc.Analyzer.Analyze',
    index=0,
    containing_service=None,
    input_type=_ANALYZEREQUEST,
    output_type=_ANALYZERESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetPluginInfo',
    full_name='pulumirpc.Analyzer.GetPluginInfo',
    index=1,
    containing_service=None,
    input_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    output_type=plugin__pb2._PLUGININFO,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_ANALYZER)

DESCRIPTOR.services_by_name['Analyzer'] = _ANALYZER

# @@protoc_insertion_point(module_scope)