<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: transaction/v2/transaction_service.proto

namespace Code\Transaction\V2\SubmitIntentResponse;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.transaction.v2.SubmitIntentResponse.Success</code>
 */
class Success extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.code.transaction.v2.SubmitIntentResponse.Success.Code code = 1;</code>
     */
    protected $code = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $code
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Transaction\V2\TransactionService::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.code.transaction.v2.SubmitIntentResponse.Success.Code code = 1;</code>
     * @return int
     */
    public function getCode()
    {
        return $this->code;
    }

    /**
     * Generated from protobuf field <code>.code.transaction.v2.SubmitIntentResponse.Success.Code code = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setCode($var)
    {
        GPBUtil::checkEnum($var, \Code\Transaction\V2\SubmitIntentResponse\Success\Code::class);
        $this->code = $var;

        return $this;
    }

}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(Success::class, \Code\Transaction\V2\SubmitIntentResponse_Success::class);

