<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: messaging/v1/messaging_service.proto

namespace Code\Messaging\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.messaging.v1.Message</code>
 */
class Message extends \Google\Protobuf\Internal\Message
{
    /**
     * MessageId is the Id of the message. This ID is generated by the
     * server, and will _always_ be set when receiving a message.
     * Server generates the message to:
     *     1. Reserve the ability for any future ID changes
     *     2. Prevent clients attempting to collide message IDs.
     *
     * Generated from protobuf field <code>.code.messaging.v1.MessageId id = 1 [(.validate.rules) = {</code>
     */
    protected $id = null;
    /**
     * The signature sent from SendMessageRequest, which will be injected by server.
     * This enables clients to ensure no MITM attacks were performed to hijack contents
     * of the typed message. This is only applicable for messages not generated by server.
     *
     * Generated from protobuf field <code>.code.common.v1.Signature send_message_request_signature = 3 [(.validate.rules) = {</code>
     */
    protected $send_message_request_signature = null;
    protected $kind;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Code\Messaging\V1\MessageId $id
     *           MessageId is the Id of the message. This ID is generated by the
     *           server, and will _always_ be set when receiving a message.
     *           Server generates the message to:
     *               1. Reserve the ability for any future ID changes
     *               2. Prevent clients attempting to collide message IDs.
     *     @type \Code\Common\V1\Signature $send_message_request_signature
     *           The signature sent from SendMessageRequest, which will be injected by server.
     *           This enables clients to ensure no MITM attacks were performed to hijack contents
     *           of the typed message. This is only applicable for messages not generated by server.
     *     @type \Code\Messaging\V1\RequestToGrabBill $request_to_grab_bill
     *     @type \Code\Messaging\V1\RequestToReceiveBill $request_to_receive_bill
     *     @type \Code\Messaging\V1\CodeScanned $code_scanned
     *     @type \Code\Messaging\V1\ClientRejectedPayment $client_rejected_payment
     *     @type \Code\Messaging\V1\IntentSubmitted $intent_submitted
     *     @type \Code\Messaging\V1\WebhookCalled $webhook_called
     *     @type \Code\Messaging\V1\RequestToLogin $request_to_login
     *     @type \Code\Messaging\V1\ClientRejectedLogin $client_rejected_login
     *     @type \Code\Messaging\V1\AirdropReceived $airdrop_received
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Messaging\V1\MessagingService::initOnce();
        parent::__construct($data);
    }

    /**
     * MessageId is the Id of the message. This ID is generated by the
     * server, and will _always_ be set when receiving a message.
     * Server generates the message to:
     *     1. Reserve the ability for any future ID changes
     *     2. Prevent clients attempting to collide message IDs.
     *
     * Generated from protobuf field <code>.code.messaging.v1.MessageId id = 1 [(.validate.rules) = {</code>
     * @return \Code\Messaging\V1\MessageId|null
     */
    public function getId()
    {
        return $this->id;
    }

    public function hasId()
    {
        return isset($this->id);
    }

    public function clearId()
    {
        unset($this->id);
    }

    /**
     * MessageId is the Id of the message. This ID is generated by the
     * server, and will _always_ be set when receiving a message.
     * Server generates the message to:
     *     1. Reserve the ability for any future ID changes
     *     2. Prevent clients attempting to collide message IDs.
     *
     * Generated from protobuf field <code>.code.messaging.v1.MessageId id = 1 [(.validate.rules) = {</code>
     * @param \Code\Messaging\V1\MessageId $var
     * @return $this
     */
    public function setId($var)
    {
        GPBUtil::checkMessage($var, \Code\Messaging\V1\MessageId::class);
        $this->id = $var;

        return $this;
    }

    /**
     * The signature sent from SendMessageRequest, which will be injected by server.
     * This enables clients to ensure no MITM attacks were performed to hijack contents
     * of the typed message. This is only applicable for messages not generated by server.
     *
     * Generated from protobuf field <code>.code.common.v1.Signature send_message_request_signature = 3 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\Signature|null
     */
    public function getSendMessageRequestSignature()
    {
        return $this->send_message_request_signature;
    }

    public function hasSendMessageRequestSignature()
    {
        return isset($this->send_message_request_signature);
    }

    public function clearSendMessageRequestSignature()
    {
        unset($this->send_message_request_signature);
    }

    /**
     * The signature sent from SendMessageRequest, which will be injected by server.
     * This enables clients to ensure no MITM attacks were performed to hijack contents
     * of the typed message. This is only applicable for messages not generated by server.
     *
     * Generated from protobuf field <code>.code.common.v1.Signature send_message_request_signature = 3 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\Signature $var
     * @return $this
     */
    public function setSendMessageRequestSignature($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\Signature::class);
        $this->send_message_request_signature = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.code.messaging.v1.RequestToGrabBill request_to_grab_bill = 2;</code>
     * @return \Code\Messaging\V1\RequestToGrabBill|null
     */
    public function getRequestToGrabBill()
    {
        return $this->readOneof(2);
    }

    public function hasRequestToGrabBill()
    {
        return $this->hasOneof(2);
    }

    /**
     * Generated from protobuf field <code>.code.messaging.v1.RequestToGrabBill request_to_grab_bill = 2;</code>
     * @param \Code\Messaging\V1\RequestToGrabBill $var
     * @return $this
     */
    public function setRequestToGrabBill($var)
    {
        GPBUtil::checkMessage($var, \Code\Messaging\V1\RequestToGrabBill::class);
        $this->writeOneof(2, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.code.messaging.v1.RequestToReceiveBill request_to_receive_bill = 5;</code>
     * @return \Code\Messaging\V1\RequestToReceiveBill|null
     */
    public function getRequestToReceiveBill()
    {
        return $this->readOneof(5);
    }

    public function hasRequestToReceiveBill()
    {
        return $this->hasOneof(5);
    }

    /**
     * Generated from protobuf field <code>.code.messaging.v1.RequestToReceiveBill request_to_receive_bill = 5;</code>
     * @param \Code\Messaging\V1\RequestToReceiveBill $var
     * @return $this
     */
    public function setRequestToReceiveBill($var)
    {
        GPBUtil::checkMessage($var, \Code\Messaging\V1\RequestToReceiveBill::class);
        $this->writeOneof(5, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.code.messaging.v1.CodeScanned code_scanned = 6;</code>
     * @return \Code\Messaging\V1\CodeScanned|null
     */
    public function getCodeScanned()
    {
        return $this->readOneof(6);
    }

    public function hasCodeScanned()
    {
        return $this->hasOneof(6);
    }

    /**
     * Generated from protobuf field <code>.code.messaging.v1.CodeScanned code_scanned = 6;</code>
     * @param \Code\Messaging\V1\CodeScanned $var
     * @return $this
     */
    public function setCodeScanned($var)
    {
        GPBUtil::checkMessage($var, \Code\Messaging\V1\CodeScanned::class);
        $this->writeOneof(6, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.code.messaging.v1.ClientRejectedPayment client_rejected_payment = 7;</code>
     * @return \Code\Messaging\V1\ClientRejectedPayment|null
     */
    public function getClientRejectedPayment()
    {
        return $this->readOneof(7);
    }

    public function hasClientRejectedPayment()
    {
        return $this->hasOneof(7);
    }

    /**
     * Generated from protobuf field <code>.code.messaging.v1.ClientRejectedPayment client_rejected_payment = 7;</code>
     * @param \Code\Messaging\V1\ClientRejectedPayment $var
     * @return $this
     */
    public function setClientRejectedPayment($var)
    {
        GPBUtil::checkMessage($var, \Code\Messaging\V1\ClientRejectedPayment::class);
        $this->writeOneof(7, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.code.messaging.v1.IntentSubmitted intent_submitted = 8;</code>
     * @return \Code\Messaging\V1\IntentSubmitted|null
     */
    public function getIntentSubmitted()
    {
        return $this->readOneof(8);
    }

    public function hasIntentSubmitted()
    {
        return $this->hasOneof(8);
    }

    /**
     * Generated from protobuf field <code>.code.messaging.v1.IntentSubmitted intent_submitted = 8;</code>
     * @param \Code\Messaging\V1\IntentSubmitted $var
     * @return $this
     */
    public function setIntentSubmitted($var)
    {
        GPBUtil::checkMessage($var, \Code\Messaging\V1\IntentSubmitted::class);
        $this->writeOneof(8, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.code.messaging.v1.WebhookCalled webhook_called = 9;</code>
     * @return \Code\Messaging\V1\WebhookCalled|null
     */
    public function getWebhookCalled()
    {
        return $this->readOneof(9);
    }

    public function hasWebhookCalled()
    {
        return $this->hasOneof(9);
    }

    /**
     * Generated from protobuf field <code>.code.messaging.v1.WebhookCalled webhook_called = 9;</code>
     * @param \Code\Messaging\V1\WebhookCalled $var
     * @return $this
     */
    public function setWebhookCalled($var)
    {
        GPBUtil::checkMessage($var, \Code\Messaging\V1\WebhookCalled::class);
        $this->writeOneof(9, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.code.messaging.v1.RequestToLogin request_to_login = 10;</code>
     * @return \Code\Messaging\V1\RequestToLogin|null
     */
    public function getRequestToLogin()
    {
        return $this->readOneof(10);
    }

    public function hasRequestToLogin()
    {
        return $this->hasOneof(10);
    }

    /**
     * Generated from protobuf field <code>.code.messaging.v1.RequestToLogin request_to_login = 10;</code>
     * @param \Code\Messaging\V1\RequestToLogin $var
     * @return $this
     */
    public function setRequestToLogin($var)
    {
        GPBUtil::checkMessage($var, \Code\Messaging\V1\RequestToLogin::class);
        $this->writeOneof(10, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.code.messaging.v1.ClientRejectedLogin client_rejected_login = 12;</code>
     * @return \Code\Messaging\V1\ClientRejectedLogin|null
     */
    public function getClientRejectedLogin()
    {
        return $this->readOneof(12);
    }

    public function hasClientRejectedLogin()
    {
        return $this->hasOneof(12);
    }

    /**
     * Generated from protobuf field <code>.code.messaging.v1.ClientRejectedLogin client_rejected_login = 12;</code>
     * @param \Code\Messaging\V1\ClientRejectedLogin $var
     * @return $this
     */
    public function setClientRejectedLogin($var)
    {
        GPBUtil::checkMessage($var, \Code\Messaging\V1\ClientRejectedLogin::class);
        $this->writeOneof(12, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.code.messaging.v1.AirdropReceived airdrop_received = 4;</code>
     * @return \Code\Messaging\V1\AirdropReceived|null
     */
    public function getAirdropReceived()
    {
        return $this->readOneof(4);
    }

    public function hasAirdropReceived()
    {
        return $this->hasOneof(4);
    }

    /**
     * Generated from protobuf field <code>.code.messaging.v1.AirdropReceived airdrop_received = 4;</code>
     * @param \Code\Messaging\V1\AirdropReceived $var
     * @return $this
     */
    public function setAirdropReceived($var)
    {
        GPBUtil::checkMessage($var, \Code\Messaging\V1\AirdropReceived::class);
        $this->writeOneof(4, $var);

        return $this;
    }

    /**
     * @return string
     */
    public function getKind()
    {
        return $this->whichOneof("kind");
    }

}

