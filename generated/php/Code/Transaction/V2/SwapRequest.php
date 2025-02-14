<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: transaction/v2/transaction_service.proto

namespace Code\Transaction\V2;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.transaction.v2.SwapRequest</code>
 */
class SwapRequest extends \Google\Protobuf\Internal\Message
{
    protected $request;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Code\Transaction\V2\SwapRequest\Initiate $initiate
     *     @type \Code\Transaction\V2\SwapRequest\SubmitSignature $submit_signature
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Transaction\V2\TransactionService::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.code.transaction.v2.SwapRequest.Initiate initiate = 1;</code>
     * @return \Code\Transaction\V2\SwapRequest\Initiate|null
     */
    public function getInitiate()
    {
        return $this->readOneof(1);
    }

    public function hasInitiate()
    {
        return $this->hasOneof(1);
    }

    /**
     * Generated from protobuf field <code>.code.transaction.v2.SwapRequest.Initiate initiate = 1;</code>
     * @param \Code\Transaction\V2\SwapRequest\Initiate $var
     * @return $this
     */
    public function setInitiate($var)
    {
        GPBUtil::checkMessage($var, \Code\Transaction\V2\SwapRequest\Initiate::class);
        $this->writeOneof(1, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.code.transaction.v2.SwapRequest.SubmitSignature submit_signature = 2;</code>
     * @return \Code\Transaction\V2\SwapRequest\SubmitSignature|null
     */
    public function getSubmitSignature()
    {
        return $this->readOneof(2);
    }

    public function hasSubmitSignature()
    {
        return $this->hasOneof(2);
    }

    /**
     * Generated from protobuf field <code>.code.transaction.v2.SwapRequest.SubmitSignature submit_signature = 2;</code>
     * @param \Code\Transaction\V2\SwapRequest\SubmitSignature $var
     * @return $this
     */
    public function setSubmitSignature($var)
    {
        GPBUtil::checkMessage($var, \Code\Transaction\V2\SwapRequest\SubmitSignature::class);
        $this->writeOneof(2, $var);

        return $this;
    }

    /**
     * @return string
     */
    public function getRequest()
    {
        return $this->whichOneof("request");
    }

}

