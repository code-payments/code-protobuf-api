# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: micropayment/v1/micro_payment_service.proto
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
    'micropayment/v1/micro_payment_service.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from common.v1 import model_pb2 as common_dot_v1_dot_model__pb2
from validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n+micropayment/v1/micro_payment_service.proto\x12\x14\x63ode.micropayment.v1\x1a\x15\x63ommon/v1/model.proto\x1a\x17validate/validate.proto\"I\n\x10GetStatusRequest\x12\x35\n\tintent_id\x18\x01 \x01(\x0b\x32\x18.code.common.v1.IntentIdB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\"S\n\x11GetStatusResponse\x12\x0e\n\x06\x65xists\x18\x01 \x01(\x08\x12\x14\n\x0c\x63ode_scanned\x18\x02 \x01(\x08\x12\x18\n\x10intent_submitted\x18\x03 \x01(\x08\"q\n\x16RegisterWebhookRequest\x12\x35\n\tintent_id\x18\x01 \x01(\x0b\x32\x18.code.common.v1.IntentIdB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12 \n\x03url\x18\x02 \x01(\tB\x13\xfa\x42\x10r\x0e\x10\x01\x18\x80\x08:\x04http\x88\x01\x01\"\xc4\x01\n\x17RegisterWebhookResponse\x12\x44\n\x06result\x18\x01 \x01(\x0e\x32\x34.code.micropayment.v1.RegisterWebhookResponse.Result\"c\n\x06Result\x12\x06\n\x02OK\x10\x00\x12\x16\n\x12\x41LREADY_REGISTERED\x10\x01\x12\x15\n\x11REQUEST_NOT_FOUND\x10\x02\x12\x11\n\rINTENT_EXISTS\x10\x03\x12\x0f\n\x0bINVALID_URL\x10\x04\"\xbb\x02\n\rCodifyRequest\x12 \n\x03url\x18\x01 \x01(\tB\x13\xfa\x42\x10r\x0e\x10\x01\x18\x80\x04:\x04http\x88\x01\x01\x12#\n\x08\x63urrency\x18\x02 \x01(\tB\x11\xfa\x42\x0er\x0c\x32\n^[a-z]{3}$\x12%\n\rnative_amount\x18\x03 \x01(\x01\x42\x0e\xfa\x42\x0b\x12\t!\x00\x00\x00\x00\x00\x00\x00\x00\x12@\n\rowner_account\x18\x04 \x01(\x0b\x32\x1f.code.common.v1.SolanaAccountIdB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12\x42\n\x0fprimary_account\x18\x05 \x01(\x0b\x32\x1f.code.common.v1.SolanaAccountIdB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12\x36\n\tsignature\x18\x06 \x01(\x0b\x32\x19.code.common.v1.SignatureB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\"\xdf\x01\n\x0e\x43odifyResponse\x12;\n\x06result\x18\x01 \x01(\x0e\x32+.code.micropayment.v1.CodifyResponse.Result\x12\x1d\n\x0c\x63odified_url\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x18@\"q\n\x06Result\x12\x06\n\x02OK\x10\x00\x12\x0f\n\x0bINVALID_URL\x10\x01\x12\x13\n\x0fINVALID_ACCOUNT\x10\x02\x12\x18\n\x14UNSUPPORTED_CURRENCY\x10\x03\x12\x1f\n\x1bNATIVE_AMOUNT_EXCEEDS_LIMIT\x10\x04\"A\n\x16GetPathMetadataRequest\x12\'\n\x04path\x18\x01 \x01(\tB\x19\xfa\x42\x16r\x14\x32\x12^[0-9a-zA-Z]{1,8}$\"\xf4\x01\n\x17GetPathMetadataResponse\x12\x44\n\x06result\x18\x01 \x01(\x0e\x32\x34.code.micropayment.v1.GetPathMetadataResponse.Result\x12\x34\n\x0b\x64\x65stination\x18\x02 \x01(\x0b\x32\x1f.code.common.v1.SolanaAccountId\x12\x10\n\x08\x63urrency\x18\x03 \x01(\t\x12\x15\n\rnative_amount\x18\x04 \x01(\x01\x12\x13\n\x0bredirct_url\x18\x05 \x01(\t\"\x1f\n\x06Result\x12\x06\n\x02OK\x10\x00\x12\r\n\tNOT_FOUND\x10\x01\x32\xa1\x03\n\x0cMicroPayment\x12\\\n\tGetStatus\x12&.code.micropayment.v1.GetStatusRequest\x1a\'.code.micropayment.v1.GetStatusResponse\x12n\n\x0fRegisterWebhook\x12,.code.micropayment.v1.RegisterWebhookRequest\x1a-.code.micropayment.v1.RegisterWebhookResponse\x12S\n\x06\x43odify\x12#.code.micropayment.v1.CodifyRequest\x1a$.code.micropayment.v1.CodifyResponse\x12n\n\x0fGetPathMetadata\x12,.code.micropayment.v1.GetPathMetadataRequest\x1a-.code.micropayment.v1.GetPathMetadataResponseB\x8b\x01\n\x1f\x63om.codeinc.gen.micropayment.v1ZTgithub.com/code-payments/code-protobuf-api/generated/go/micropayment/v1;micropayment\xa2\x02\x11\x41PBMicroPaymentV1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'micropayment.v1.micro_payment_service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\037com.codeinc.gen.micropayment.v1ZTgithub.com/code-payments/code-protobuf-api/generated/go/micropayment/v1;micropayment\242\002\021APBMicroPaymentV1'
  _globals['_GETSTATUSREQUEST'].fields_by_name['intent_id']._loaded_options = None
  _globals['_GETSTATUSREQUEST'].fields_by_name['intent_id']._serialized_options = b'\372B\005\212\001\002\020\001'
  _globals['_REGISTERWEBHOOKREQUEST'].fields_by_name['intent_id']._loaded_options = None
  _globals['_REGISTERWEBHOOKREQUEST'].fields_by_name['intent_id']._serialized_options = b'\372B\005\212\001\002\020\001'
  _globals['_REGISTERWEBHOOKREQUEST'].fields_by_name['url']._loaded_options = None
  _globals['_REGISTERWEBHOOKREQUEST'].fields_by_name['url']._serialized_options = b'\372B\020r\016\020\001\030\200\010:\004http\210\001\001'
  _globals['_CODIFYREQUEST'].fields_by_name['url']._loaded_options = None
  _globals['_CODIFYREQUEST'].fields_by_name['url']._serialized_options = b'\372B\020r\016\020\001\030\200\004:\004http\210\001\001'
  _globals['_CODIFYREQUEST'].fields_by_name['currency']._loaded_options = None
  _globals['_CODIFYREQUEST'].fields_by_name['currency']._serialized_options = b'\372B\016r\0142\n^[a-z]{3}$'
  _globals['_CODIFYREQUEST'].fields_by_name['native_amount']._loaded_options = None
  _globals['_CODIFYREQUEST'].fields_by_name['native_amount']._serialized_options = b'\372B\013\022\t!\000\000\000\000\000\000\000\000'
  _globals['_CODIFYREQUEST'].fields_by_name['owner_account']._loaded_options = None
  _globals['_CODIFYREQUEST'].fields_by_name['owner_account']._serialized_options = b'\372B\005\212\001\002\020\001'
  _globals['_CODIFYREQUEST'].fields_by_name['primary_account']._loaded_options = None
  _globals['_CODIFYREQUEST'].fields_by_name['primary_account']._serialized_options = b'\372B\005\212\001\002\020\001'
  _globals['_CODIFYREQUEST'].fields_by_name['signature']._loaded_options = None
  _globals['_CODIFYREQUEST'].fields_by_name['signature']._serialized_options = b'\372B\005\212\001\002\020\001'
  _globals['_CODIFYRESPONSE'].fields_by_name['codified_url']._loaded_options = None
  _globals['_CODIFYRESPONSE'].fields_by_name['codified_url']._serialized_options = b'\372B\004r\002\030@'
  _globals['_GETPATHMETADATAREQUEST'].fields_by_name['path']._loaded_options = None
  _globals['_GETPATHMETADATAREQUEST'].fields_by_name['path']._serialized_options = b'\372B\026r\0242\022^[0-9a-zA-Z]{1,8}$'
  _globals['_GETSTATUSREQUEST']._serialized_start=117
  _globals['_GETSTATUSREQUEST']._serialized_end=190
  _globals['_GETSTATUSRESPONSE']._serialized_start=192
  _globals['_GETSTATUSRESPONSE']._serialized_end=275
  _globals['_REGISTERWEBHOOKREQUEST']._serialized_start=277
  _globals['_REGISTERWEBHOOKREQUEST']._serialized_end=390
  _globals['_REGISTERWEBHOOKRESPONSE']._serialized_start=393
  _globals['_REGISTERWEBHOOKRESPONSE']._serialized_end=589
  _globals['_REGISTERWEBHOOKRESPONSE_RESULT']._serialized_start=490
  _globals['_REGISTERWEBHOOKRESPONSE_RESULT']._serialized_end=589
  _globals['_CODIFYREQUEST']._serialized_start=592
  _globals['_CODIFYREQUEST']._serialized_end=907
  _globals['_CODIFYRESPONSE']._serialized_start=910
  _globals['_CODIFYRESPONSE']._serialized_end=1133
  _globals['_CODIFYRESPONSE_RESULT']._serialized_start=1020
  _globals['_CODIFYRESPONSE_RESULT']._serialized_end=1133
  _globals['_GETPATHMETADATAREQUEST']._serialized_start=1135
  _globals['_GETPATHMETADATAREQUEST']._serialized_end=1200
  _globals['_GETPATHMETADATARESPONSE']._serialized_start=1203
  _globals['_GETPATHMETADATARESPONSE']._serialized_end=1447
  _globals['_GETPATHMETADATARESPONSE_RESULT']._serialized_start=1416
  _globals['_GETPATHMETADATARESPONSE_RESULT']._serialized_end=1447
  _globals['_MICROPAYMENT']._serialized_start=1450
  _globals['_MICROPAYMENT']._serialized_end=1867
# @@protoc_insertion_point(module_scope)
