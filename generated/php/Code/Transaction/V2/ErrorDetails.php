<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: transaction/v2/transaction_service.proto

namespace Code\Transaction\V2;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.transaction.v2.ErrorDetails</code>
 */
class ErrorDetails extends \Google\Protobuf\Internal\Message
{
    protected $type;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Code\Transaction\V2\ReasonStringErrorDetails $reason_string
     *     @type \Code\Transaction\V2\InvalidSignatureErrorDetails $invalid_signature
     *     @type \Code\Transaction\V2\DeniedErrorDetails $denied
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Transaction\V2\TransactionService::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.code.transaction.v2.ReasonStringErrorDetails reason_string = 1;</code>
     * @return \Code\Transaction\V2\ReasonStringErrorDetails|null
     */
    public function getReasonString()
    {
        return $this->readOneof(1);
    }

    public function hasReasonString()
    {
        return $this->hasOneof(1);
    }

    /**
     * Generated from protobuf field <code>.code.transaction.v2.ReasonStringErrorDetails reason_string = 1;</code>
     * @param \Code\Transaction\V2\ReasonStringErrorDetails $var
     * @return $this
     */
    public function setReasonString($var)
    {
        GPBUtil::checkMessage($var, \Code\Transaction\V2\ReasonStringErrorDetails::class);
        $this->writeOneof(1, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.code.transaction.v2.InvalidSignatureErrorDetails invalid_signature = 2;</code>
     * @return \Code\Transaction\V2\InvalidSignatureErrorDetails|null
     */
    public function getInvalidSignature()
    {
        return $this->readOneof(2);
    }

    public function hasInvalidSignature()
    {
        return $this->hasOneof(2);
    }

    /**
     * Generated from protobuf field <code>.code.transaction.v2.InvalidSignatureErrorDetails invalid_signature = 2;</code>
     * @param \Code\Transaction\V2\InvalidSignatureErrorDetails $var
     * @return $this
     */
    public function setInvalidSignature($var)
    {
        GPBUtil::checkMessage($var, \Code\Transaction\V2\InvalidSignatureErrorDetails::class);
        $this->writeOneof(2, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.code.transaction.v2.DeniedErrorDetails denied = 3;</code>
     * @return \Code\Transaction\V2\DeniedErrorDetails|null
     */
    public function getDenied()
    {
        return $this->readOneof(3);
    }

    public function hasDenied()
    {
        return $this->hasOneof(3);
    }

    /**
     * Generated from protobuf field <code>.code.transaction.v2.DeniedErrorDetails denied = 3;</code>
     * @param \Code\Transaction\V2\DeniedErrorDetails $var
     * @return $this
     */
    public function setDenied($var)
    {
        GPBUtil::checkMessage($var, \Code\Transaction\V2\DeniedErrorDetails::class);
        $this->writeOneof(3, $var);

        return $this;
    }

    /**
     * @return string
     */
    public function getType()
    {
        return $this->whichOneof("type");
    }

}

