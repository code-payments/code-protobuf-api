<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: push/v1/push_service.proto

namespace Code\Push\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.push.v1.RemoveTokenRequest</code>
 */
class RemoveTokenRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * The public key of the owner account that signed this request message.
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId owner_account_id = 1 [(.validate.rules) = {</code>
     */
    protected $owner_account_id = null;
    /**
     * The signature is of serialize(AddTokenRequest) without this field set
     * using the private key of owner_account_id. This provides an authentication
     * mechanism to the RPC.
     *
     * Generated from protobuf field <code>.code.common.v1.Signature signature = 2 [(.validate.rules) = {</code>
     */
    protected $signature = null;
    /**
     * The data container where the push token was stored.
     *
     * Generated from protobuf field <code>.code.common.v1.DataContainerId container_id = 3 [(.validate.rules) = {</code>
     */
    protected $container_id = null;
    /**
     * The push token to remove.
     *
     * Generated from protobuf field <code>string push_token = 4 [(.validate.rules) = {</code>
     */
    protected $push_token = '';
    /**
     * The type of push token to remove.
     *
     * Generated from protobuf field <code>.code.push.v1.TokenType token_type = 5 [(.validate.rules) = {</code>
     */
    protected $token_type = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Code\Common\V1\SolanaAccountId $owner_account_id
     *           The public key of the owner account that signed this request message.
     *     @type \Code\Common\V1\Signature $signature
     *           The signature is of serialize(AddTokenRequest) without this field set
     *           using the private key of owner_account_id. This provides an authentication
     *           mechanism to the RPC.
     *     @type \Code\Common\V1\DataContainerId $container_id
     *           The data container where the push token was stored.
     *     @type string $push_token
     *           The push token to remove.
     *     @type int $token_type
     *           The type of push token to remove.
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Push\V1\PushService::initOnce();
        parent::__construct($data);
    }

    /**
     * The public key of the owner account that signed this request message.
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId owner_account_id = 1 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\SolanaAccountId|null
     */
    public function getOwnerAccountId()
    {
        return $this->owner_account_id;
    }

    public function hasOwnerAccountId()
    {
        return isset($this->owner_account_id);
    }

    public function clearOwnerAccountId()
    {
        unset($this->owner_account_id);
    }

    /**
     * The public key of the owner account that signed this request message.
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId owner_account_id = 1 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\SolanaAccountId $var
     * @return $this
     */
    public function setOwnerAccountId($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\SolanaAccountId::class);
        $this->owner_account_id = $var;

        return $this;
    }

    /**
     * The signature is of serialize(AddTokenRequest) without this field set
     * using the private key of owner_account_id. This provides an authentication
     * mechanism to the RPC.
     *
     * Generated from protobuf field <code>.code.common.v1.Signature signature = 2 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\Signature|null
     */
    public function getSignature()
    {
        return $this->signature;
    }

    public function hasSignature()
    {
        return isset($this->signature);
    }

    public function clearSignature()
    {
        unset($this->signature);
    }

    /**
     * The signature is of serialize(AddTokenRequest) without this field set
     * using the private key of owner_account_id. This provides an authentication
     * mechanism to the RPC.
     *
     * Generated from protobuf field <code>.code.common.v1.Signature signature = 2 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\Signature $var
     * @return $this
     */
    public function setSignature($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\Signature::class);
        $this->signature = $var;

        return $this;
    }

    /**
     * The data container where the push token was stored.
     *
     * Generated from protobuf field <code>.code.common.v1.DataContainerId container_id = 3 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\DataContainerId|null
     */
    public function getContainerId()
    {
        return $this->container_id;
    }

    public function hasContainerId()
    {
        return isset($this->container_id);
    }

    public function clearContainerId()
    {
        unset($this->container_id);
    }

    /**
     * The data container where the push token was stored.
     *
     * Generated from protobuf field <code>.code.common.v1.DataContainerId container_id = 3 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\DataContainerId $var
     * @return $this
     */
    public function setContainerId($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\DataContainerId::class);
        $this->container_id = $var;

        return $this;
    }

    /**
     * The push token to remove.
     *
     * Generated from protobuf field <code>string push_token = 4 [(.validate.rules) = {</code>
     * @return string
     */
    public function getPushToken()
    {
        return $this->push_token;
    }

    /**
     * The push token to remove.
     *
     * Generated from protobuf field <code>string push_token = 4 [(.validate.rules) = {</code>
     * @param string $var
     * @return $this
     */
    public function setPushToken($var)
    {
        GPBUtil::checkString($var, True);
        $this->push_token = $var;

        return $this;
    }

    /**
     * The type of push token to remove.
     *
     * Generated from protobuf field <code>.code.push.v1.TokenType token_type = 5 [(.validate.rules) = {</code>
     * @return int
     */
    public function getTokenType()
    {
        return $this->token_type;
    }

    /**
     * The type of push token to remove.
     *
     * Generated from protobuf field <code>.code.push.v1.TokenType token_type = 5 [(.validate.rules) = {</code>
     * @param int $var
     * @return $this
     */
    public function setTokenType($var)
    {
        GPBUtil::checkEnum($var, \Code\Push\V1\TokenType::class);
        $this->token_type = $var;

        return $this;
    }

}

