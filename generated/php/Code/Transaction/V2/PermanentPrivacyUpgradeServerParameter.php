<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: transaction/v2/transaction_service.proto

namespace Code\Transaction\V2;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.transaction.v2.PermanentPrivacyUpgradeServerParameter</code>
 */
class PermanentPrivacyUpgradeServerParameter extends \Google\Protobuf\Internal\Message
{
    /**
     * The new commitment that is being paid
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId new_commitment = 1 [(.validate.rules) = {</code>
     */
    protected $new_commitment = null;
    /**
     * The new commitment account's transcript. This is purely needed by client
     * to validate merkle_root with commitment PDA logic.
     *
     * Generated from protobuf field <code>.code.common.v1.Hash new_commitment_transcript = 2 [(.validate.rules) = {</code>
     */
    protected $new_commitment_transcript = null;
    /**
     * The new commitment account's destination. This is purely needed by client
     * to validate merkle_root with commitment PDA logic.
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId new_commitment_destination = 3 [(.validate.rules) = {</code>
     */
    protected $new_commitment_destination = null;
    /**
     * The new commitment account's payment amount. This is purely needed by client
     * to validate merkle_root with commitment PDA logic.
     *
     * Generated from protobuf field <code>uint64 new_commitment_amount = 4 [(.validate.rules) = {</code>
     */
    protected $new_commitment_amount = 0;
    /**
     * The merkle root, which was the recent root used in the new commitment account
     *
     * Generated from protobuf field <code>.code.common.v1.Hash merkle_root = 5 [(.validate.rules) = {</code>
     */
    protected $merkle_root = null;
    /**
     * The merkle proof that validates the original commitment occurred prior to
     * the new commitment server is asking client to pay
     *
     * Generated from protobuf field <code>repeated .code.common.v1.Hash merkle_proof = 6 [(.validate.rules) = {</code>
     */
    private $merkle_proof;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Code\Common\V1\SolanaAccountId $new_commitment
     *           The new commitment that is being paid
     *     @type \Code\Common\V1\Hash $new_commitment_transcript
     *           The new commitment account's transcript. This is purely needed by client
     *           to validate merkle_root with commitment PDA logic.
     *     @type \Code\Common\V1\SolanaAccountId $new_commitment_destination
     *           The new commitment account's destination. This is purely needed by client
     *           to validate merkle_root with commitment PDA logic.
     *     @type int|string $new_commitment_amount
     *           The new commitment account's payment amount. This is purely needed by client
     *           to validate merkle_root with commitment PDA logic.
     *     @type \Code\Common\V1\Hash $merkle_root
     *           The merkle root, which was the recent root used in the new commitment account
     *     @type array<\Code\Common\V1\Hash>|\Google\Protobuf\Internal\RepeatedField $merkle_proof
     *           The merkle proof that validates the original commitment occurred prior to
     *           the new commitment server is asking client to pay
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Transaction\V2\TransactionService::initOnce();
        parent::__construct($data);
    }

    /**
     * The new commitment that is being paid
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId new_commitment = 1 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\SolanaAccountId|null
     */
    public function getNewCommitment()
    {
        return $this->new_commitment;
    }

    public function hasNewCommitment()
    {
        return isset($this->new_commitment);
    }

    public function clearNewCommitment()
    {
        unset($this->new_commitment);
    }

    /**
     * The new commitment that is being paid
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId new_commitment = 1 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\SolanaAccountId $var
     * @return $this
     */
    public function setNewCommitment($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\SolanaAccountId::class);
        $this->new_commitment = $var;

        return $this;
    }

    /**
     * The new commitment account's transcript. This is purely needed by client
     * to validate merkle_root with commitment PDA logic.
     *
     * Generated from protobuf field <code>.code.common.v1.Hash new_commitment_transcript = 2 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\Hash|null
     */
    public function getNewCommitmentTranscript()
    {
        return $this->new_commitment_transcript;
    }

    public function hasNewCommitmentTranscript()
    {
        return isset($this->new_commitment_transcript);
    }

    public function clearNewCommitmentTranscript()
    {
        unset($this->new_commitment_transcript);
    }

    /**
     * The new commitment account's transcript. This is purely needed by client
     * to validate merkle_root with commitment PDA logic.
     *
     * Generated from protobuf field <code>.code.common.v1.Hash new_commitment_transcript = 2 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\Hash $var
     * @return $this
     */
    public function setNewCommitmentTranscript($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\Hash::class);
        $this->new_commitment_transcript = $var;

        return $this;
    }

    /**
     * The new commitment account's destination. This is purely needed by client
     * to validate merkle_root with commitment PDA logic.
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId new_commitment_destination = 3 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\SolanaAccountId|null
     */
    public function getNewCommitmentDestination()
    {
        return $this->new_commitment_destination;
    }

    public function hasNewCommitmentDestination()
    {
        return isset($this->new_commitment_destination);
    }

    public function clearNewCommitmentDestination()
    {
        unset($this->new_commitment_destination);
    }

    /**
     * The new commitment account's destination. This is purely needed by client
     * to validate merkle_root with commitment PDA logic.
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId new_commitment_destination = 3 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\SolanaAccountId $var
     * @return $this
     */
    public function setNewCommitmentDestination($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\SolanaAccountId::class);
        $this->new_commitment_destination = $var;

        return $this;
    }

    /**
     * The new commitment account's payment amount. This is purely needed by client
     * to validate merkle_root with commitment PDA logic.
     *
     * Generated from protobuf field <code>uint64 new_commitment_amount = 4 [(.validate.rules) = {</code>
     * @return int|string
     */
    public function getNewCommitmentAmount()
    {
        return $this->new_commitment_amount;
    }

    /**
     * The new commitment account's payment amount. This is purely needed by client
     * to validate merkle_root with commitment PDA logic.
     *
     * Generated from protobuf field <code>uint64 new_commitment_amount = 4 [(.validate.rules) = {</code>
     * @param int|string $var
     * @return $this
     */
    public function setNewCommitmentAmount($var)
    {
        GPBUtil::checkUint64($var);
        $this->new_commitment_amount = $var;

        return $this;
    }

    /**
     * The merkle root, which was the recent root used in the new commitment account
     *
     * Generated from protobuf field <code>.code.common.v1.Hash merkle_root = 5 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\Hash|null
     */
    public function getMerkleRoot()
    {
        return $this->merkle_root;
    }

    public function hasMerkleRoot()
    {
        return isset($this->merkle_root);
    }

    public function clearMerkleRoot()
    {
        unset($this->merkle_root);
    }

    /**
     * The merkle root, which was the recent root used in the new commitment account
     *
     * Generated from protobuf field <code>.code.common.v1.Hash merkle_root = 5 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\Hash $var
     * @return $this
     */
    public function setMerkleRoot($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\Hash::class);
        $this->merkle_root = $var;

        return $this;
    }

    /**
     * The merkle proof that validates the original commitment occurred prior to
     * the new commitment server is asking client to pay
     *
     * Generated from protobuf field <code>repeated .code.common.v1.Hash merkle_proof = 6 [(.validate.rules) = {</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getMerkleProof()
    {
        return $this->merkle_proof;
    }

    /**
     * The merkle proof that validates the original commitment occurred prior to
     * the new commitment server is asking client to pay
     *
     * Generated from protobuf field <code>repeated .code.common.v1.Hash merkle_proof = 6 [(.validate.rules) = {</code>
     * @param array<\Code\Common\V1\Hash>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setMerkleProof($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Code\Common\V1\Hash::class);
        $this->merkle_proof = $arr;

        return $this;
    }

}

