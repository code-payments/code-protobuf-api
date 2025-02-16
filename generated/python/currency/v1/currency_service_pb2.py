# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: currency/v1/currency_service.proto
# Protobuf Python Version: 5.29.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    29,
    0,
    '',
    'currency/v1/currency_service.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from validate import validate_pb2 as validate_dot_validate__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\"currency/v1/currency_service.proto\x12\x10\x63ode.currency.v1\x1a\x17validate/validate.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"C\n\x12GetAllRatesRequest\x12-\n\ttimestamp\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\xb5\x02\n\x13GetAllRatesResponse\x12<\n\x06result\x18\x01 \x01(\x0e\x32,.code.currency.v1.GetAllRatesResponse.Result\x12\x33\n\x05\x61s_of\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x08\xfa\x42\x05\xb2\x01\x02\x08\x01\x12Y\n\x05rates\x18\x03 \x03(\x0b\x32\x30.code.currency.v1.GetAllRatesResponse.RatesEntryB\x18\xfa\x42\x15\x9a\x01\x12\"\x10r\x0e\x32\x0c^[a-z]{3,4}$\x1a,\n\nRatesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\x01:\x02\x38\x01\"\"\n\x06Result\x12\x06\n\x02OK\x10\x00\x12\x10\n\x0cMISSING_DATA\x10\x01\x32\x66\n\x08\x43urrency\x12Z\n\x0bGetAllRates\x12$.code.currency.v1.GetAllRatesRequest\x1a%.code.currency.v1.GetAllRatesResponseB{\n\x1b\x63om.codeinc.gen.currency.v1ZLgithub.com/code-payments/code-protobuf-api/generated/go/currency/v1;currency\xa2\x02\rCPBCurrencyV1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'currency.v1.currency_service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\033com.codeinc.gen.currency.v1ZLgithub.com/code-payments/code-protobuf-api/generated/go/currency/v1;currency\242\002\rCPBCurrencyV1'
  _globals['_GETALLRATESRESPONSE_RATESENTRY']._loaded_options = None
  _globals['_GETALLRATESRESPONSE_RATESENTRY']._serialized_options = b'8\001'
  _globals['_GETALLRATESRESPONSE'].fields_by_name['as_of']._loaded_options = None
  _globals['_GETALLRATESRESPONSE'].fields_by_name['as_of']._serialized_options = b'\372B\005\262\001\002\010\001'
  _globals['_GETALLRATESRESPONSE'].fields_by_name['rates']._loaded_options = None
  _globals['_GETALLRATESRESPONSE'].fields_by_name['rates']._serialized_options = b'\372B\025\232\001\022\"\020r\0162\014^[a-z]{3,4}$'
  _globals['_GETALLRATESREQUEST']._serialized_start=114
  _globals['_GETALLRATESREQUEST']._serialized_end=181
  _globals['_GETALLRATESRESPONSE']._serialized_start=184
  _globals['_GETALLRATESRESPONSE']._serialized_end=493
  _globals['_GETALLRATESRESPONSE_RATESENTRY']._serialized_start=413
  _globals['_GETALLRATESRESPONSE_RATESENTRY']._serialized_end=457
  _globals['_GETALLRATESRESPONSE_RESULT']._serialized_start=459
  _globals['_GETALLRATESRESPONSE_RESULT']._serialized_end=493
  _globals['_CURRENCY']._serialized_start=495
  _globals['_CURRENCY']._serialized_end=597
# @@protoc_insertion_point(module_scope)
