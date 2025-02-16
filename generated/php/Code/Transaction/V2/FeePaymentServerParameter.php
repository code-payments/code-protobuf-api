<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: transaction/v2/transaction_service.proto

namespace Code\Transaction\V2;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.transaction.v2.FeePaymentServerParameter</code>
 */
class FeePaymentServerParameter extends \Google\Protobuf\Internal\Message
{
    /**
     * The destination account where Code fee payments should be sent. This will
     * only be set when the corresponding FeePaymentAction Type is CODE.
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId code_destination = 1;</code>
     */
    protected $code_destination = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Code\Common\V1\SolanaAccountId $code_destination
     *           The destination account where Code fee payments should be sent. This will
     *           only be set when the corresponding FeePaymentAction Type is CODE.
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Transaction\V2\TransactionService::initOnce();
        parent::__construct($data);
    }

    /**
     * The destination account where Code fee payments should be sent. This will
     * only be set when the corresponding FeePaymentAction Type is CODE.
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId code_destination = 1;</code>
     * @return \Code\Common\V1\SolanaAccountId|null
     */
    public function getCodeDestination()
    {
        return $this->code_destination;
    }

    public function hasCodeDestination()
    {
        return isset($this->code_destination);
    }

    public function clearCodeDestination()
    {
        unset($this->code_destination);
    }

    /**
     * The destination account where Code fee payments should be sent. This will
     * only be set when the corresponding FeePaymentAction Type is CODE.
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId code_destination = 1;</code>
     * @param \Code\Common\V1\SolanaAccountId $var
     * @return $this
     */
    public function setCodeDestination($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\SolanaAccountId::class);
        $this->code_destination = $var;

        return $this;
    }

}

