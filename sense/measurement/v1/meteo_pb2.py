# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: measurement/v1/meteo.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1ameasurement/v1/meteo.proto\x12\x0emeasurement.v1\x1a\x1fgoogle/protobuf/timestamp.proto\"C\n\x0bMeasurement\x12\n\n\x02id\x18\x01 \x01(\t\x12(\n\x04time\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\x9b\x01\n\x07\x43limate\x12\x30\n\x0bmeasurement\x18\x01 \x01(\x0b\x32\x1b.measurement.v1.Measurement\x12\x10\n\x08humidity\x18\x02 \x01(\x02\x12\x10\n\x08pressure\x18\x03 \x01(\x02\x12\x1c\n\x14temperature_humidity\x18\x04 \x01(\x02\x12\x1c\n\x14temperature_pressure\x18\x05 \x01(\x02\"\xd1\x03\n\x06System\x12\x30\n\x0bmeasurement\x18\x01 \x01(\x0b\x32\x1b.measurement.v1.Measurement\x12\x17\n\x0f\x63pu_temperature\x18\x02 \x01(\x02\x12%\n\x02la\x18\x03 \x01(\x0b\x32\x19.measurement.v1.System.LA\x12-\n\x06memory\x18\x04 \x01(\x0b\x32\x1d.measurement.v1.System.Memory\x12)\n\x04\x64isk\x18\x05 \x01(\x0b\x32\x1b.measurement.v1.System.Disk\x1a/\n\x02LA\x12\x0c\n\x04min1\x18\x01 \x01(\x02\x12\x0c\n\x04min5\x18\x02 \x01(\x02\x12\r\n\x05min15\x18\x03 \x01(\x02\x1aw\n\x06Memory\x12\x10\n\x08total_kb\x18\x01 \x01(\x03\x12\x0f\n\x07used_kb\x18\x02 \x01(\x03\x12\x0f\n\x07\x66ree_kb\x18\x03 \x01(\x03\x12\x11\n\tshared_kb\x18\x04 \x01(\x03\x12\x10\n\x08\x63\x61\x63he_kb\x18\x05 \x01(\x03\x12\x14\n\x0c\x61vailable_kb\x18\x06 \x01(\x03\x1aQ\n\x04\x44isk\x12\x10\n\x08total_kb\x18\x01 \x01(\x03\x12\x0f\n\x07used_kb\x18\x02 \x01(\x03\x12\x14\n\x0c\x61vailable_kb\x18\x03 \x01(\x03\x12\x10\n\x08use_prct\x18\x04 \x01(\x05\x42\x10Z\x0emeasurement/v1b\x06proto3')



_MEASUREMENT = DESCRIPTOR.message_types_by_name['Measurement']
_CLIMATE = DESCRIPTOR.message_types_by_name['Climate']
_SYSTEM = DESCRIPTOR.message_types_by_name['System']
_SYSTEM_LA = _SYSTEM.nested_types_by_name['LA']
_SYSTEM_MEMORY = _SYSTEM.nested_types_by_name['Memory']
_SYSTEM_DISK = _SYSTEM.nested_types_by_name['Disk']
Measurement = _reflection.GeneratedProtocolMessageType('Measurement', (_message.Message,), {
  'DESCRIPTOR' : _MEASUREMENT,
  '__module__' : 'measurement.v1.meteo_pb2'
  # @@protoc_insertion_point(class_scope:measurement.v1.Measurement)
  })
_sym_db.RegisterMessage(Measurement)

Climate = _reflection.GeneratedProtocolMessageType('Climate', (_message.Message,), {
  'DESCRIPTOR' : _CLIMATE,
  '__module__' : 'measurement.v1.meteo_pb2'
  # @@protoc_insertion_point(class_scope:measurement.v1.Climate)
  })
_sym_db.RegisterMessage(Climate)

System = _reflection.GeneratedProtocolMessageType('System', (_message.Message,), {

  'LA' : _reflection.GeneratedProtocolMessageType('LA', (_message.Message,), {
    'DESCRIPTOR' : _SYSTEM_LA,
    '__module__' : 'measurement.v1.meteo_pb2'
    # @@protoc_insertion_point(class_scope:measurement.v1.System.LA)
    })
  ,

  'Memory' : _reflection.GeneratedProtocolMessageType('Memory', (_message.Message,), {
    'DESCRIPTOR' : _SYSTEM_MEMORY,
    '__module__' : 'measurement.v1.meteo_pb2'
    # @@protoc_insertion_point(class_scope:measurement.v1.System.Memory)
    })
  ,

  'Disk' : _reflection.GeneratedProtocolMessageType('Disk', (_message.Message,), {
    'DESCRIPTOR' : _SYSTEM_DISK,
    '__module__' : 'measurement.v1.meteo_pb2'
    # @@protoc_insertion_point(class_scope:measurement.v1.System.Disk)
    })
  ,
  'DESCRIPTOR' : _SYSTEM,
  '__module__' : 'measurement.v1.meteo_pb2'
  # @@protoc_insertion_point(class_scope:measurement.v1.System)
  })
_sym_db.RegisterMessage(System)
_sym_db.RegisterMessage(System.LA)
_sym_db.RegisterMessage(System.Memory)
_sym_db.RegisterMessage(System.Disk)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\016measurement/v1'
  _MEASUREMENT._serialized_start=79
  _MEASUREMENT._serialized_end=146
  _CLIMATE._serialized_start=149
  _CLIMATE._serialized_end=304
  _SYSTEM._serialized_start=307
  _SYSTEM._serialized_end=772
  _SYSTEM_LA._serialized_start=521
  _SYSTEM_LA._serialized_end=568
  _SYSTEM_MEMORY._serialized_start=570
  _SYSTEM_MEMORY._serialized_end=689
  _SYSTEM_DISK._serialized_start=691
  _SYSTEM_DISK._serialized_end=772
# @@protoc_insertion_point(module_scope)
