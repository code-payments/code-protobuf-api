<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: transaction/v2/transaction_service.proto

namespace Code\Transaction\V2;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Sends a payment to a destination account with initial temporary privacy. Clients
 * should also reorganize their bucket accounts and rotate their temporary outgoing
 * account.
 * Action Spec (In Person Cash Payment or Withdrawal or Tip):
 * actions = [
 *   // Section 1: Transfer ExchangeData.Quarks from BUCKET_X_KIN accounts to TEMPORARY_OUTGOING account with reogranizations
 *   TemporaryPrivacyExchangeAction(BUCKET_X_KIN, BUCKET_X_KIN, multiple * bucketSize),
 *   TemporaryPrivacyTransferAction(BUCKET_X_KIN, TEMPORARY_OUTGOING[index], multiple * bucketSize),
 *   ...,
 *   TemporaryPrivacyExchangeAction(BUCKET_X_KIN, BUCKET_X_KIN, multiple * bucketSize),
 *   TemporaryPrivacyTransferAction(BUCKET_X_KIN, TEMPORARY_OUTGOING[index], multiple * bucketSize),
 *   // Section 2: Rotate TEMPORARY_OUTGOING account
 *   // Below must appear last in this exact order
 *   NoPrivacyWithdrawAction(TEMPORARY_OUTGOING[index], destination, ExchangeData.Quarks),
 *   OpenAccountAction(TEMPORARY_OUTGOING[index + 1]),
 *   CloseDormantAccount(TEMPORARY_OUTGOING[index + 1]),
 * ]
 * Action Spec (Remote Send):
 * actions = [
 *   // Section 1: Open REMOTE_SEND_GIFT_CARD account
 *   OpenAccountAction(REMOTE_SEND_GIFT_CARD),
 *   // Section 2: Transfer ExchangeData.Quarks from BUCKET_X_KIN accounts to TEMPORARY_OUTGOING account with reogranizations
 *   TemporaryPrivacyExchangeAction(BUCKET_X_KIN, BUCKET_X_KIN, multiple * bucketSize),
 *   TemporaryPrivacyTransferAction(BUCKET_X_KIN, TEMPORARY_OUTGOING[index], multiple * bucketSize),
 *   ...,
 *   TemporaryPrivacyExchangeAction(BUCKET_X_KIN, BUCKET_X_KIN, multiple * bucketSize),
 *   TemporaryPrivacyTransferAction(BUCKET_X_KIN, TEMPORARY_OUTGOING[index], multiple * bucketSize),
 *   // Section 3: Rotate TEMPORARY_OUTGOING account
 *   // Below must appear last in this exact order
 *   NoPrivacyWithdrawAction(TEMPORARY_OUTGOING[index], REMOTE_SEND_GIFT_CARD, ExchangeData.Quarks),
 *   OpenAccountAction(TEMPORARY_OUTGOING[index + 1]),
 *   CloseDormantAccount(TEMPORARY_OUTGOING[index + 1]),
 *   // Section 4: Close REMOTE_SEND_GIFT_CARD if not redeemed after period of time
 *   CloseDormantAccount(REMOTE_SEND_GIFT_CARD),
 * Action Spec (Micro Payment):
 * actions = [
 *   // Section 1: Transfer ExchangeData.Quarks from BUCKET_X_KIN accounts to TEMPORARY_OUTGOING account with reogranizations
 *   TemporaryPrivacyExchangeAction(BUCKET_X_KIN, BUCKET_X_KIN, multiple * bucketSize),
 *   TemporaryPrivacyTransferAction(BUCKET_X_KIN, TEMPORARY_OUTGOING[index], multiple * bucketSize),
 *   ...,
 *   TemporaryPrivacyExchangeAction(BUCKET_X_KIN, BUCKET_X_KIN, multiple * bucketSize),
 *   TemporaryPrivacyTransferAction(BUCKET_X_KIN, TEMPORARY_OUTGOING[index], multiple * bucketSize),
 *   // Section 2: Fee payments
 *   // Hard-coded Code $0.01 USD fee to a dynamic fee account
 *   FeePayment(TEMPORARY_OUTGOING[index], codeFeeAccount, $0.01 USD of Kin),
 *   // Additional fees, exactly as specified in the original payment request
 *   FeePayment(TEMPORARY_OUTGOING[index], additionalFeeAccount0, additionalFeeQuarks0),
 *   ...
 *   FeePayment(TEMPORARY_OUTGOING[index], additionalFeeAccountN, additionalFeeQuarksN),
 *   // Section 3: Rotate TEMPORARY_OUTGOING account
 *   // Below must appear last in this exact order
 *   NoPrivacyWithdrawAction(TEMPORARY_OUTGOING[index], destination, ExchangeData.Quarks - $0.01 USD of Kin - additionalFeeQuarks0 - ... - additionalFeeQuarksN),
 *   OpenAccountAction(TEMPORARY_OUTGOING[index + 1]),
 *   CloseDormantAccount(TEMPORARY_OUTGOING[index + 1]),
 * ]
 *
 * Generated from protobuf message <code>code.transaction.v2.SendPrivatePaymentMetadata</code>
 */
class SendPrivatePaymentMetadata extends \Google\Protobuf\Internal\Message
{
    /**
     * The destination token account to send funds to
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId destination = 1 [(.validate.rules) = {</code>
     */
    protected $destination = null;
    /**
     * The exchange data of total funds being sent to the destination
     *
     * Generated from protobuf field <code>.code.transaction.v2.ExchangeData exchange_data = 2 [(.validate.rules) = {</code>
     */
    protected $exchange_data = null;
    /**
     * Is the payment a withdrawal? For destinations that are not Code temporary
     * accounts, this must be set to true.
     *
     * Generated from protobuf field <code>bool is_withdrawal = 3;</code>
     */
    protected $is_withdrawal = false;
    /**
     * Is the payment for a remote send?
     *
     * Generated from protobuf field <code>bool is_remote_send = 4;</code>
     */
    protected $is_remote_send = false;
    /**
     * Is the payment for a tip?
     *
     * Generated from protobuf field <code>bool is_tip = 5;</code>
     */
    protected $is_tip = false;
    /**
     * If is_tip is true, the user being tipped
     *
     * Generated from protobuf field <code>.code.transaction.v2.TippedUser tipped_user = 6;</code>
     */
    protected $tipped_user = null;
    /**
     * Is the payment for a chat?
     *
     * Generated from protobuf field <code>bool is_chat = 7;</code>
     */
    protected $is_chat = false;
    /**
     * If is_chat is true, the chat being paid for.
     *
     * Generated from protobuf field <code>.code.common.v1.ChatId chat_id = 8;</code>
     */
    protected $chat_id = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Code\Common\V1\SolanaAccountId $destination
     *           The destination token account to send funds to
     *     @type \Code\Transaction\V2\ExchangeData $exchange_data
     *           The exchange data of total funds being sent to the destination
     *     @type bool $is_withdrawal
     *           Is the payment a withdrawal? For destinations that are not Code temporary
     *           accounts, this must be set to true.
     *     @type bool $is_remote_send
     *           Is the payment for a remote send?
     *     @type bool $is_tip
     *           Is the payment for a tip?
     *     @type \Code\Transaction\V2\TippedUser $tipped_user
     *           If is_tip is true, the user being tipped
     *     @type bool $is_chat
     *           Is the payment for a chat?
     *     @type \Code\Common\V1\ChatId $chat_id
     *           If is_chat is true, the chat being paid for.
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Transaction\V2\TransactionService::initOnce();
        parent::__construct($data);
    }

    /**
     * The destination token account to send funds to
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId destination = 1 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\SolanaAccountId|null
     */
    public function getDestination()
    {
        return $this->destination;
    }

    public function hasDestination()
    {
        return isset($this->destination);
    }

    public function clearDestination()
    {
        unset($this->destination);
    }

    /**
     * The destination token account to send funds to
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId destination = 1 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\SolanaAccountId $var
     * @return $this
     */
    public function setDestination($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\SolanaAccountId::class);
        $this->destination = $var;

        return $this;
    }

    /**
     * The exchange data of total funds being sent to the destination
     *
     * Generated from protobuf field <code>.code.transaction.v2.ExchangeData exchange_data = 2 [(.validate.rules) = {</code>
     * @return \Code\Transaction\V2\ExchangeData|null
     */
    public function getExchangeData()
    {
        return $this->exchange_data;
    }

    public function hasExchangeData()
    {
        return isset($this->exchange_data);
    }

    public function clearExchangeData()
    {
        unset($this->exchange_data);
    }

    /**
     * The exchange data of total funds being sent to the destination
     *
     * Generated from protobuf field <code>.code.transaction.v2.ExchangeData exchange_data = 2 [(.validate.rules) = {</code>
     * @param \Code\Transaction\V2\ExchangeData $var
     * @return $this
     */
    public function setExchangeData($var)
    {
        GPBUtil::checkMessage($var, \Code\Transaction\V2\ExchangeData::class);
        $this->exchange_data = $var;

        return $this;
    }

    /**
     * Is the payment a withdrawal? For destinations that are not Code temporary
     * accounts, this must be set to true.
     *
     * Generated from protobuf field <code>bool is_withdrawal = 3;</code>
     * @return bool
     */
    public function getIsWithdrawal()
    {
        return $this->is_withdrawal;
    }

    /**
     * Is the payment a withdrawal? For destinations that are not Code temporary
     * accounts, this must be set to true.
     *
     * Generated from protobuf field <code>bool is_withdrawal = 3;</code>
     * @param bool $var
     * @return $this
     */
    public function setIsWithdrawal($var)
    {
        GPBUtil::checkBool($var);
        $this->is_withdrawal = $var;

        return $this;
    }

    /**
     * Is the payment for a remote send?
     *
     * Generated from protobuf field <code>bool is_remote_send = 4;</code>
     * @return bool
     */
    public function getIsRemoteSend()
    {
        return $this->is_remote_send;
    }

    /**
     * Is the payment for a remote send?
     *
     * Generated from protobuf field <code>bool is_remote_send = 4;</code>
     * @param bool $var
     * @return $this
     */
    public function setIsRemoteSend($var)
    {
        GPBUtil::checkBool($var);
        $this->is_remote_send = $var;

        return $this;
    }

    /**
     * Is the payment for a tip?
     *
     * Generated from protobuf field <code>bool is_tip = 5;</code>
     * @return bool
     */
    public function getIsTip()
    {
        return $this->is_tip;
    }

    /**
     * Is the payment for a tip?
     *
     * Generated from protobuf field <code>bool is_tip = 5;</code>
     * @param bool $var
     * @return $this
     */
    public function setIsTip($var)
    {
        GPBUtil::checkBool($var);
        $this->is_tip = $var;

        return $this;
    }

    /**
     * If is_tip is true, the user being tipped
     *
     * Generated from protobuf field <code>.code.transaction.v2.TippedUser tipped_user = 6;</code>
     * @return \Code\Transaction\V2\TippedUser|null
     */
    public function getTippedUser()
    {
        return $this->tipped_user;
    }

    public function hasTippedUser()
    {
        return isset($this->tipped_user);
    }

    public function clearTippedUser()
    {
        unset($this->tipped_user);
    }

    /**
     * If is_tip is true, the user being tipped
     *
     * Generated from protobuf field <code>.code.transaction.v2.TippedUser tipped_user = 6;</code>
     * @param \Code\Transaction\V2\TippedUser $var
     * @return $this
     */
    public function setTippedUser($var)
    {
        GPBUtil::checkMessage($var, \Code\Transaction\V2\TippedUser::class);
        $this->tipped_user = $var;

        return $this;
    }

    /**
     * Is the payment for a chat?
     *
     * Generated from protobuf field <code>bool is_chat = 7;</code>
     * @return bool
     */
    public function getIsChat()
    {
        return $this->is_chat;
    }

    /**
     * Is the payment for a chat?
     *
     * Generated from protobuf field <code>bool is_chat = 7;</code>
     * @param bool $var
     * @return $this
     */
    public function setIsChat($var)
    {
        GPBUtil::checkBool($var);
        $this->is_chat = $var;

        return $this;
    }

    /**
     * If is_chat is true, the chat being paid for.
     *
     * Generated from protobuf field <code>.code.common.v1.ChatId chat_id = 8;</code>
     * @return \Code\Common\V1\ChatId|null
     */
    public function getChatId()
    {
        return $this->chat_id;
    }

    public function hasChatId()
    {
        return isset($this->chat_id);
    }

    public function clearChatId()
    {
        unset($this->chat_id);
    }

    /**
     * If is_chat is true, the chat being paid for.
     *
     * Generated from protobuf field <code>.code.common.v1.ChatId chat_id = 8;</code>
     * @param \Code\Common\V1\ChatId $var
     * @return $this
     */
    public function setChatId($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\ChatId::class);
        $this->chat_id = $var;

        return $this;
    }

}

