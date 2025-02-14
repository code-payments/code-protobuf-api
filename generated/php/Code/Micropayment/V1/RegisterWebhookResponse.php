<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: micropayment/v1/micro_payment_service.proto

namespace Code\Micropayment\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.micropayment.v1.RegisterWebhookResponse</code>
 */
class RegisterWebhookResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.code.micropayment.v1.RegisterWebhookResponse.Result result = 1;</code>
     */
    protected $result = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $result
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Micropayment\V1\MicroPaymentService::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.code.micropayment.v1.RegisterWebhookResponse.Result result = 1;</code>
     * @return int
     */
    public function getResult()
    {
        return $this->result;
    }

    /**
     * Generated from protobuf field <code>.code.micropayment.v1.RegisterWebhookResponse.Result result = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setResult($var)
    {
        GPBUtil::checkEnum($var, \Code\Micropayment\V1\RegisterWebhookResponse\Result::class);
        $this->result = $var;

        return $this;
    }

}

