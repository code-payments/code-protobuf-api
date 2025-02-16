<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: transaction/v2/transaction_service.proto

namespace Code\Transaction\V2;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.transaction.v2.GetLimitsRequest</code>
 */
class GetLimitsRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * The owner account whose limits will be calculated. Any other owner accounts
     * linked with the same identity of the owner will also be applied.
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId owner = 1 [(.validate.rules) = {</code>
     */
    protected $owner = null;
    /**
     * The signature is of serialize(GetLimitsRequest) without this field set
     * using the private key of the owner account. This provides an authentication
     * mechanism to the RPC.
     *
     * Generated from protobuf field <code>.code.common.v1.Signature signature = 2 [(.validate.rules) = {</code>
     */
    protected $signature = null;
    /**
     * All transactions starting at this time will be incorporated into the consumed
     * limit calculation. Clients should set this to the start of the current day in
     * the client's current time zone (because server has no knowledge of this atm).
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp consumed_since = 3 [(.validate.rules) = {</code>
     */
    protected $consumed_since = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Code\Common\V1\SolanaAccountId $owner
     *           The owner account whose limits will be calculated. Any other owner accounts
     *           linked with the same identity of the owner will also be applied.
     *     @type \Code\Common\V1\Signature $signature
     *           The signature is of serialize(GetLimitsRequest) without this field set
     *           using the private key of the owner account. This provides an authentication
     *           mechanism to the RPC.
     *     @type \Google\Protobuf\Timestamp $consumed_since
     *           All transactions starting at this time will be incorporated into the consumed
     *           limit calculation. Clients should set this to the start of the current day in
     *           the client's current time zone (because server has no knowledge of this atm).
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Transaction\V2\TransactionService::initOnce();
        parent::__construct($data);
    }

    /**
     * The owner account whose limits will be calculated. Any other owner accounts
     * linked with the same identity of the owner will also be applied.
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId owner = 1 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\SolanaAccountId|null
     */
    public function getOwner()
    {
        return $this->owner;
    }

    public function hasOwner()
    {
        return isset($this->owner);
    }

    public function clearOwner()
    {
        unset($this->owner);
    }

    /**
     * The owner account whose limits will be calculated. Any other owner accounts
     * linked with the same identity of the owner will also be applied.
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId owner = 1 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\SolanaAccountId $var
     * @return $this
     */
    public function setOwner($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\SolanaAccountId::class);
        $this->owner = $var;

        return $this;
    }

    /**
     * The signature is of serialize(GetLimitsRequest) without this field set
     * using the private key of the owner account. This provides an authentication
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
     * The signature is of serialize(GetLimitsRequest) without this field set
     * using the private key of the owner account. This provides an authentication
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
     * All transactions starting at this time will be incorporated into the consumed
     * limit calculation. Clients should set this to the start of the current day in
     * the client's current time zone (because server has no knowledge of this atm).
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp consumed_since = 3 [(.validate.rules) = {</code>
     * @return \Google\Protobuf\Timestamp|null
     */
    public function getConsumedSince()
    {
        return $this->consumed_since;
    }

    public function hasConsumedSince()
    {
        return isset($this->consumed_since);
    }

    public function clearConsumedSince()
    {
        unset($this->consumed_since);
    }

    /**
     * All transactions starting at this time will be incorporated into the consumed
     * limit calculation. Clients should set this to the start of the current day in
     * the client's current time zone (because server has no knowledge of this atm).
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp consumed_since = 3 [(.validate.rules) = {</code>
     * @param \Google\Protobuf\Timestamp $var
     * @return $this
     */
    public function setConsumedSince($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\Timestamp::class);
        $this->consumed_since = $var;

        return $this;
    }

}

