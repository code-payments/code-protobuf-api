<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: device/v1/device_service.proto

namespace GPBMetadata\Device\V1;

class DeviceService
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\Common\V1\Model::initOnce();
        \GPBMetadata\Validate\Validate::initOnce();
        $pool->internalAddGeneratedFile(
            '
�	
device/v1/device_service.protocode.device.v1validate/validate.proto"�
RegisterLoggedInAccountsRequest;
app_install (2.code.common.v1.AppInstallIdB�B�;
owners (2.code.common.v1.SolanaAccountIdB
�B� 9

signatures (2.code.common.v1.SignatureB
�B� "�
 RegisterLoggedInAccountsResponseG
result (27.code.device.v1.RegisterLoggedInAccountsResponse.ResultC
invalid_owners (2.code.common.v1.SolanaAccountIdB
�B� "#
Result
OK 
INVALID_OWNER"Y
GetLoggedInAccountsRequest;
app_install (2.code.common.v1.AppInstallIdB�B�"�
GetLoggedInAccountsResponseB
result (22.code.device.v1.GetLoggedInAccountsResponse.Result;
owners (2.code.common.v1.SolanaAccountIdB
�B� "
Result
OK 2�
Device}
RegisterLoggedInAccounts/.code.device.v1.RegisterLoggedInAccountsRequest0.code.device.v1.RegisterLoggedInAccountsResponsen
GetLoggedInAccounts*.code.device.v1.GetLoggedInAccountsRequest+.code.device.v1.GetLoggedInAccountsResponseBt
com.codeinc.gen.device.v1ZHgithub.com/code-payments/code-protobuf-api/generated/go/device/v1;device�CPBDevicetV1bproto3'
        , true);

        static::$is_initialized = true;
    }
}

