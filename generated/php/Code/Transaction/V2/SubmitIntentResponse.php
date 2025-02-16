<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: transaction/v2/transaction_service.proto

namespace Code\Transaction\V2;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.transaction.v2.SubmitIntentResponse</code>
 */
class SubmitIntentResponse extends \Google\Protobuf\Internal\Message
{
    protected $response;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Code\Transaction\V2\SubmitIntentResponse\ServerParameters $server_parameters
     *     @type \Code\Transaction\V2\SubmitIntentResponse\Success $success
     *     @type \Code\Transaction\V2\SubmitIntentResponse\Error $error
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Transaction\V2\TransactionService::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.code.transaction.v2.SubmitIntentResponse.ServerParameters server_parameters = 1;</code>
     * @return \Code\Transaction\V2\SubmitIntentResponse\ServerParameters|null
     */
    public function getServerParameters()
    {
        return $this->readOneof(1);
    }

    public function hasServerParameters()
    {
        return $this->hasOneof(1);
    }

    /**
     * Generated from protobuf field <code>.code.transaction.v2.SubmitIntentResponse.ServerParameters server_parameters = 1;</code>
     * @param \Code\Transaction\V2\SubmitIntentResponse\ServerParameters $var
     * @return $this
     */
    public function setServerParameters($var)
    {
        GPBUtil::checkMessage($var, \Code\Transaction\V2\SubmitIntentResponse\ServerParameters::class);
        $this->writeOneof(1, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.code.transaction.v2.SubmitIntentResponse.Success success = 2;</code>
     * @return \Code\Transaction\V2\SubmitIntentResponse\Success|null
     */
    public function getSuccess()
    {
        return $this->readOneof(2);
    }

    public function hasSuccess()
    {
        return $this->hasOneof(2);
    }

    /**
     * Generated from protobuf field <code>.code.transaction.v2.SubmitIntentResponse.Success success = 2;</code>
     * @param \Code\Transaction\V2\SubmitIntentResponse\Success $var
     * @return $this
     */
    public function setSuccess($var)
    {
        GPBUtil::checkMessage($var, \Code\Transaction\V2\SubmitIntentResponse\Success::class);
        $this->writeOneof(2, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.code.transaction.v2.SubmitIntentResponse.Error error = 3;</code>
     * @return \Code\Transaction\V2\SubmitIntentResponse\Error|null
     */
    public function getError()
    {
        return $this->readOneof(3);
    }

    public function hasError()
    {
        return $this->hasOneof(3);
    }

    /**
     * Generated from protobuf field <code>.code.transaction.v2.SubmitIntentResponse.Error error = 3;</code>
     * @param \Code\Transaction\V2\SubmitIntentResponse\Error $var
     * @return $this
     */
    public function setError($var)
    {
        GPBUtil::checkMessage($var, \Code\Transaction\V2\SubmitIntentResponse\Error::class);
        $this->writeOneof(3, $var);

        return $this;
    }

    /**
     * @return string
     */
    public function getResponse()
    {
        return $this->whichOneof("response");
    }

}

