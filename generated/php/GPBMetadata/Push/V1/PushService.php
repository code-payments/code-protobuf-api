<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: push/v1/push_service.proto

namespace GPBMetadata\Push\V1;

class PushService
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
push/v1/push_service.protocode.push.v1validate/validate.proto"�
AddTokenRequestC
owner_account_id (2.code.common.v1.SolanaAccountIdB�B�6
	signature (2.code.common.v1.SignatureB�B�?
container_id (2.code.common.v1.DataContainerIdB�B�

push_token (	B
�Br� 7

token_type (2.code.push.v1.TokenTypeB
�B�1
app_install (2.code.common.v1.AppInstallId"s
AddTokenResponse5
result (2%.code.push.v1.AddTokenResponse.Result"(
Result
OK 
INVALID_PUSH_TOKEN"�
RemoveTokenRequestC
owner_account_id (2.code.common.v1.SolanaAccountIdB�B�6
	signature (2.code.common.v1.SignatureB�B�?
container_id (2.code.common.v1.DataContainerIdB�B�

push_token (	B
�Br� 7

token_type (2.code.push.v1.TokenTypeB
�B�"a
RemoveTokenResponse8
result (2(.code.push.v1.RemoveTokenResponse.Result"
Result
OK *7
	TokenType
UNKNOWN 
FCM_ANDROID
FCM_APNS2�
PushI
AddToken.code.push.v1.AddTokenRequest.code.push.v1.AddTokenResponseR
RemoveToken .code.push.v1.RemoveTokenRequest!.code.push.v1.RemoveTokenResponseBk
com.codeinc.gen.push.v1ZDgithub.com/code-payments/code-protobuf-api/generated/go/push/v1;push�	APBPushV1bproto3'
        , true);

        static::$is_initialized = true;
    }
}

