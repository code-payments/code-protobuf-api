<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chat/v2/chat_service.proto

namespace Code\Chat\V2;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Encrypted piece of content using NaCl box encryption
 *
 * Generated from protobuf message <code>code.chat.v2.NaclBoxEncryptedContent</code>
 */
class NaclBoxEncryptedContent extends \Google\Protobuf\Internal\Message
{
    /**
     * The sender's public key that is used to derive the shared private key for
     * decryption for message content.
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId peer_public_key = 1 [(.validate.rules) = {</code>
     */
    protected $peer_public_key = null;
    /**
     * Globally random nonce that is unique to this encrypted piece of content
     *
     * Generated from protobuf field <code>bytes nonce = 2 [(.validate.rules) = {</code>
     */
    protected $nonce = '';
    /**
     * The encrypted piece of message content
     *
     * Generated from protobuf field <code>bytes encrypted_payload = 3 [(.validate.rules) = {</code>
     */
    protected $encrypted_payload = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Code\Common\V1\SolanaAccountId $peer_public_key
     *           The sender's public key that is used to derive the shared private key for
     *           decryption for message content.
     *     @type string $nonce
     *           Globally random nonce that is unique to this encrypted piece of content
     *     @type string $encrypted_payload
     *           The encrypted piece of message content
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Chat\V2\ChatService::initOnce();
        parent::__construct($data);
    }

    /**
     * The sender's public key that is used to derive the shared private key for
     * decryption for message content.
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId peer_public_key = 1 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\SolanaAccountId|null
     */
    public function getPeerPublicKey()
    {
        return $this->peer_public_key;
    }

    public function hasPeerPublicKey()
    {
        return isset($this->peer_public_key);
    }

    public function clearPeerPublicKey()
    {
        unset($this->peer_public_key);
    }

    /**
     * The sender's public key that is used to derive the shared private key for
     * decryption for message content.
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId peer_public_key = 1 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\SolanaAccountId $var
     * @return $this
     */
    public function setPeerPublicKey($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\SolanaAccountId::class);
        $this->peer_public_key = $var;

        return $this;
    }

    /**
     * Globally random nonce that is unique to this encrypted piece of content
     *
     * Generated from protobuf field <code>bytes nonce = 2 [(.validate.rules) = {</code>
     * @return string
     */
    public function getNonce()
    {
        return $this->nonce;
    }

    /**
     * Globally random nonce that is unique to this encrypted piece of content
     *
     * Generated from protobuf field <code>bytes nonce = 2 [(.validate.rules) = {</code>
     * @param string $var
     * @return $this
     */
    public function setNonce($var)
    {
        GPBUtil::checkString($var, False);
        $this->nonce = $var;

        return $this;
    }

    /**
     * The encrypted piece of message content
     *
     * Generated from protobuf field <code>bytes encrypted_payload = 3 [(.validate.rules) = {</code>
     * @return string
     */
    public function getEncryptedPayload()
    {
        return $this->encrypted_payload;
    }

    /**
     * The encrypted piece of message content
     *
     * Generated from protobuf field <code>bytes encrypted_payload = 3 [(.validate.rules) = {</code>
     * @param string $var
     * @return $this
     */
    public function setEncryptedPayload($var)
    {
        GPBUtil::checkString($var, False);
        $this->encrypted_payload = $var;

        return $this;
    }

}

