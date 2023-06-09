# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: nlp.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\tnlp.proto\"\x13\n\x03\x44oc\x12\x0c\n\x04text\x18\x01 \x01(\t\"*\n\x06Intent\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x12\n\nconfidence\x18\x02 \x01(\x02\"W\n\x06\x45ntity\x12\r\n\x05start\x18\x01 \x01(\r\x12\x0b\n\x03\x65nd\x18\x02 \x01(\r\x12\x0c\n\x04type\x18\x03 \x01(\t\x12\r\n\x05value\x18\x04 \x01(\t\x12\x14\n\x0cnormal_value\x18\x05 \x01(\t\"Z\n\x06Result\x12\x0c\n\x04text\x18\x01 \x01(\t\x12\x1c\n\x06intent\x18\x02 \x01(\x0b\x32\x07.IntentH\x00\x88\x01\x01\x12\x19\n\x08\x65ntities\x18\x03 \x03(\x0b\x32\x07.EntityB\t\n\x07_intent2\x1d\n\x03NLP\x12\x16\n\x05Parse\x12\x04.Doc\x1a\x07.ResultB\x13Z\x11internal/grpc/nlpb\x06proto3')



_DOC = DESCRIPTOR.message_types_by_name['Doc']
_INTENT = DESCRIPTOR.message_types_by_name['Intent']
_ENTITY = DESCRIPTOR.message_types_by_name['Entity']
_RESULT = DESCRIPTOR.message_types_by_name['Result']
Doc = _reflection.GeneratedProtocolMessageType('Doc', (_message.Message,), {
  'DESCRIPTOR' : _DOC,
  '__module__' : 'nlp_pb2'
  # @@protoc_insertion_point(class_scope:Doc)
  })
_sym_db.RegisterMessage(Doc)

Intent = _reflection.GeneratedProtocolMessageType('Intent', (_message.Message,), {
  'DESCRIPTOR' : _INTENT,
  '__module__' : 'nlp_pb2'
  # @@protoc_insertion_point(class_scope:Intent)
  })
_sym_db.RegisterMessage(Intent)

Entity = _reflection.GeneratedProtocolMessageType('Entity', (_message.Message,), {
  'DESCRIPTOR' : _ENTITY,
  '__module__' : 'nlp_pb2'
  # @@protoc_insertion_point(class_scope:Entity)
  })
_sym_db.RegisterMessage(Entity)

Result = _reflection.GeneratedProtocolMessageType('Result', (_message.Message,), {
  'DESCRIPTOR' : _RESULT,
  '__module__' : 'nlp_pb2'
  # @@protoc_insertion_point(class_scope:Result)
  })
_sym_db.RegisterMessage(Result)

_NLP = DESCRIPTOR.services_by_name['NLP']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\021internal/grpc/nlp'
  _DOC._serialized_start=13
  _DOC._serialized_end=32
  _INTENT._serialized_start=34
  _INTENT._serialized_end=76
  _ENTITY._serialized_start=78
  _ENTITY._serialized_end=165
  _RESULT._serialized_start=167
  _RESULT._serialized_end=257
  _NLP._serialized_start=259
  _NLP._serialized_end=288
# @@protoc_insertion_point(module_scope)
