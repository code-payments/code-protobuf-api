<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user/v1/identity_service.proto

namespace Code\User\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.user.v1.UnlinkAccountRequest</code>
 */
class UnlinkAccountRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * The public key of the owner account that will be unliked.
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId owner_account_id = 1 [(.validate.rules) = {</code>
     */
    protected $owner_account_id = null;
    /**
     * The signature is of serialize(UnlinkAccountRequest) without this field set
     * using the private key of owner_account_id. This provides an authentication
     * mechanism to the RPC.
     *
     * Generated from protobuf field <code>.code.common.v1.Signature signature = 2 [(.validate.rules) = {</code>
     */
    protected $signature = null;
    protected $identifying_feature;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Code\Common\V1\SolanaAccountId $owner_account_id
     *           The public key of the owner account that will be unliked.
     *     @type \Code\Common\V1\Signature $signature
     *           The signature is of serialize(UnlinkAccountRequest) without this field set
     *           using the private key of owner_account_id. This provides an authentication
     *           mechanism to the RPC.
     *     @type \Code\Common\V1\PhoneNumber $phone_number
     *           The phone number associated with the owner account.
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\User\V1\IdentityService::initOnce();
        parent::__construct($data);
    }

    /**
     * The public key of the owner account that will be unliked.
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
     * The public key of the owner account that will be unliked.
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
     * The signature is of serialize(UnlinkAccountRequest) without this field set
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
     * The signature is of serialize(UnlinkAccountRequest) without this field set
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
     * The phone number associated with the owner account.
     *
     * Generated from protobuf field <code>.code.common.v1.PhoneNumber phone_number = 4 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\PhoneNumber|null
     */
    public function getPhoneNumber()
    {
        return $this->readOneof(4);
    }

    public function hasPhoneNumber()
    {
        return $this->hasOneof(4);
    }

    /**
     * The phone number associated with the owner account.
     *
     * Generated from protobuf field <code>.code.common.v1.PhoneNumber phone_number = 4 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\PhoneNumber $var
     * @return $this
     */
    public function setPhoneNumber($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\PhoneNumber::class);
        $this->writeOneof(4, $var);

        return $this;
    }

    /**
     * @return string
     */
    public function getIdentifyingFeature()
    {
        return $this->whichOneof("identifying_feature");
    }

}

