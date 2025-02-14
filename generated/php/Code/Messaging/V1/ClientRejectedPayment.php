<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: messaging/v1/messaging_service.proto

namespace Code\Messaging\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Payment is rejected by the client
 * This message type is only initiated by clients
 *
 * Generated from protobuf message <code>code.messaging.v1.ClientRejectedPayment</code>
 */
class ClientRejectedPayment extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.code.common.v1.IntentId intent_id = 1 [(.validate.rules) = {</code>
     */
    protected $intent_id = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Code\Common\V1\IntentId $intent_id
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Messaging\V1\MessagingService::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.code.common.v1.IntentId intent_id = 1 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\IntentId|null
     */
    public function getIntentId()
    {
        return $this->intent_id;
    }

    public function hasIntentId()
    {
        return isset($this->intent_id);
    }

    public function clearIntentId()
    {
        unset($this->intent_id);
    }

    /**
     * Generated from protobuf field <code>.code.common.v1.IntentId intent_id = 1 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\IntentId $var
     * @return $this
     */
    public function setIntentId($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\IntentId::class);
        $this->intent_id = $var;

        return $this;
    }

}

