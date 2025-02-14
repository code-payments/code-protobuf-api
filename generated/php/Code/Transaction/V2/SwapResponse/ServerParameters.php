<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: transaction/v2/transaction_service.proto

namespace Code\Transaction\V2\SwapResponse;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.transaction.v2.SwapResponse.ServerParameters</code>
 */
class ServerParameters extends \Google\Protobuf\Internal\Message
{
    /**
     * Subisdizer account that will be paying for the swap
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId payer = 1 [(.validate.rules) = {</code>
     */
    protected $payer = null;
    /**
     * Recent blockhash
     *
     * Generated from protobuf field <code>.code.common.v1.Blockhash recent_blockhash = 2 [(.validate.rules) = {</code>
     */
    protected $recent_blockhash = null;
    /**
     * Compute unit limit provided to the ComputeBudget::SetComputeUnitLimit
     * instruction. If the value is 0, then the instruction can be omitted.
     *
     * Generated from protobuf field <code>uint32 compute_unit_limit = 3;</code>
     */
    protected $compute_unit_limit = 0;
    /**
     * Compute unit price provided in the ComputeBudget::SetComputeUnitPrice
     * instruction. If the value is 0, then the instruction can be omitted.
     *
     * Generated from protobuf field <code>uint64 compute_unit_price = 4;</code>
     */
    protected $compute_unit_price = 0;
    /**
     * On-chain program that will be performing the swap
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId swap_program = 5 [(.validate.rules) = {</code>
     */
    protected $swap_program = null;
    /**
     * Accounts provided to the swap instruction
     *
     * Generated from protobuf field <code>repeated .code.common.v1.InstructionAccount swap_ixn_accounts = 6 [(.validate.rules) = {</code>
     */
    private $swap_ixn_accounts;
    /**
     * Instruction data for the swap instruction
     *
     * Generated from protobuf field <code>bytes swap_ixn_data = 7 [(.validate.rules) = {</code>
     */
    protected $swap_ixn_data = '';
    /**
     * Maximum quarks that will be sent out of the source account after
     * executing the swap. If not, the validation instruction will cause
     * the transaction to fail.
     *
     * Generated from protobuf field <code>uint64 max_to_send = 8 [(.validate.rules) = {</code>
     */
    protected $max_to_send = 0;
    /**
     * Minimum quarks that will be received into the destination account
     * after executing the swap. If not, the validation instruction will
     * cause the transaction to fail.
     *
     * Generated from protobuf field <code>uint64 min_to_receive = 9 [(.validate.rules) = {</code>
     */
    protected $min_to_receive = 0;
    /**
     * Nonce to use in swap validator state account PDA
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId nonce = 10 [(.validate.rules) = {</code>
     */
    protected $nonce = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Code\Common\V1\SolanaAccountId $payer
     *           Subisdizer account that will be paying for the swap
     *     @type \Code\Common\V1\Blockhash $recent_blockhash
     *           Recent blockhash
     *     @type int $compute_unit_limit
     *           Compute unit limit provided to the ComputeBudget::SetComputeUnitLimit
     *           instruction. If the value is 0, then the instruction can be omitted.
     *     @type int|string $compute_unit_price
     *           Compute unit price provided in the ComputeBudget::SetComputeUnitPrice
     *           instruction. If the value is 0, then the instruction can be omitted.
     *     @type \Code\Common\V1\SolanaAccountId $swap_program
     *           On-chain program that will be performing the swap
     *     @type array<\Code\Common\V1\InstructionAccount>|\Google\Protobuf\Internal\RepeatedField $swap_ixn_accounts
     *           Accounts provided to the swap instruction
     *     @type string $swap_ixn_data
     *           Instruction data for the swap instruction
     *     @type int|string $max_to_send
     *           Maximum quarks that will be sent out of the source account after
     *           executing the swap. If not, the validation instruction will cause
     *           the transaction to fail.
     *     @type int|string $min_to_receive
     *           Minimum quarks that will be received into the destination account
     *           after executing the swap. If not, the validation instruction will
     *           cause the transaction to fail.
     *     @type \Code\Common\V1\SolanaAccountId $nonce
     *           Nonce to use in swap validator state account PDA
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Transaction\V2\TransactionService::initOnce();
        parent::__construct($data);
    }

    /**
     * Subisdizer account that will be paying for the swap
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId payer = 1 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\SolanaAccountId|null
     */
    public function getPayer()
    {
        return $this->payer;
    }

    public function hasPayer()
    {
        return isset($this->payer);
    }

    public function clearPayer()
    {
        unset($this->payer);
    }

    /**
     * Subisdizer account that will be paying for the swap
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId payer = 1 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\SolanaAccountId $var
     * @return $this
     */
    public function setPayer($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\SolanaAccountId::class);
        $this->payer = $var;

        return $this;
    }

    /**
     * Recent blockhash
     *
     * Generated from protobuf field <code>.code.common.v1.Blockhash recent_blockhash = 2 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\Blockhash|null
     */
    public function getRecentBlockhash()
    {
        return $this->recent_blockhash;
    }

    public function hasRecentBlockhash()
    {
        return isset($this->recent_blockhash);
    }

    public function clearRecentBlockhash()
    {
        unset($this->recent_blockhash);
    }

    /**
     * Recent blockhash
     *
     * Generated from protobuf field <code>.code.common.v1.Blockhash recent_blockhash = 2 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\Blockhash $var
     * @return $this
     */
    public function setRecentBlockhash($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\Blockhash::class);
        $this->recent_blockhash = $var;

        return $this;
    }

    /**
     * Compute unit limit provided to the ComputeBudget::SetComputeUnitLimit
     * instruction. If the value is 0, then the instruction can be omitted.
     *
     * Generated from protobuf field <code>uint32 compute_unit_limit = 3;</code>
     * @return int
     */
    public function getComputeUnitLimit()
    {
        return $this->compute_unit_limit;
    }

    /**
     * Compute unit limit provided to the ComputeBudget::SetComputeUnitLimit
     * instruction. If the value is 0, then the instruction can be omitted.
     *
     * Generated from protobuf field <code>uint32 compute_unit_limit = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setComputeUnitLimit($var)
    {
        GPBUtil::checkUint32($var);
        $this->compute_unit_limit = $var;

        return $this;
    }

    /**
     * Compute unit price provided in the ComputeBudget::SetComputeUnitPrice
     * instruction. If the value is 0, then the instruction can be omitted.
     *
     * Generated from protobuf field <code>uint64 compute_unit_price = 4;</code>
     * @return int|string
     */
    public function getComputeUnitPrice()
    {
        return $this->compute_unit_price;
    }

    /**
     * Compute unit price provided in the ComputeBudget::SetComputeUnitPrice
     * instruction. If the value is 0, then the instruction can be omitted.
     *
     * Generated from protobuf field <code>uint64 compute_unit_price = 4;</code>
     * @param int|string $var
     * @return $this
     */
    public function setComputeUnitPrice($var)
    {
        GPBUtil::checkUint64($var);
        $this->compute_unit_price = $var;

        return $this;
    }

    /**
     * On-chain program that will be performing the swap
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId swap_program = 5 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\SolanaAccountId|null
     */
    public function getSwapProgram()
    {
        return $this->swap_program;
    }

    public function hasSwapProgram()
    {
        return isset($this->swap_program);
    }

    public function clearSwapProgram()
    {
        unset($this->swap_program);
    }

    /**
     * On-chain program that will be performing the swap
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId swap_program = 5 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\SolanaAccountId $var
     * @return $this
     */
    public function setSwapProgram($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\SolanaAccountId::class);
        $this->swap_program = $var;

        return $this;
    }

    /**
     * Accounts provided to the swap instruction
     *
     * Generated from protobuf field <code>repeated .code.common.v1.InstructionAccount swap_ixn_accounts = 6 [(.validate.rules) = {</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getSwapIxnAccounts()
    {
        return $this->swap_ixn_accounts;
    }

    /**
     * Accounts provided to the swap instruction
     *
     * Generated from protobuf field <code>repeated .code.common.v1.InstructionAccount swap_ixn_accounts = 6 [(.validate.rules) = {</code>
     * @param array<\Code\Common\V1\InstructionAccount>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setSwapIxnAccounts($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Code\Common\V1\InstructionAccount::class);
        $this->swap_ixn_accounts = $arr;

        return $this;
    }

    /**
     * Instruction data for the swap instruction
     *
     * Generated from protobuf field <code>bytes swap_ixn_data = 7 [(.validate.rules) = {</code>
     * @return string
     */
    public function getSwapIxnData()
    {
        return $this->swap_ixn_data;
    }

    /**
     * Instruction data for the swap instruction
     *
     * Generated from protobuf field <code>bytes swap_ixn_data = 7 [(.validate.rules) = {</code>
     * @param string $var
     * @return $this
     */
    public function setSwapIxnData($var)
    {
        GPBUtil::checkString($var, False);
        $this->swap_ixn_data = $var;

        return $this;
    }

    /**
     * Maximum quarks that will be sent out of the source account after
     * executing the swap. If not, the validation instruction will cause
     * the transaction to fail.
     *
     * Generated from protobuf field <code>uint64 max_to_send = 8 [(.validate.rules) = {</code>
     * @return int|string
     */
    public function getMaxToSend()
    {
        return $this->max_to_send;
    }

    /**
     * Maximum quarks that will be sent out of the source account after
     * executing the swap. If not, the validation instruction will cause
     * the transaction to fail.
     *
     * Generated from protobuf field <code>uint64 max_to_send = 8 [(.validate.rules) = {</code>
     * @param int|string $var
     * @return $this
     */
    public function setMaxToSend($var)
    {
        GPBUtil::checkUint64($var);
        $this->max_to_send = $var;

        return $this;
    }

    /**
     * Minimum quarks that will be received into the destination account
     * after executing the swap. If not, the validation instruction will
     * cause the transaction to fail.
     *
     * Generated from protobuf field <code>uint64 min_to_receive = 9 [(.validate.rules) = {</code>
     * @return int|string
     */
    public function getMinToReceive()
    {
        return $this->min_to_receive;
    }

    /**
     * Minimum quarks that will be received into the destination account
     * after executing the swap. If not, the validation instruction will
     * cause the transaction to fail.
     *
     * Generated from protobuf field <code>uint64 min_to_receive = 9 [(.validate.rules) = {</code>
     * @param int|string $var
     * @return $this
     */
    public function setMinToReceive($var)
    {
        GPBUtil::checkUint64($var);
        $this->min_to_receive = $var;

        return $this;
    }

    /**
     * Nonce to use in swap validator state account PDA
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId nonce = 10 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\SolanaAccountId|null
     */
    public function getNonce()
    {
        return $this->nonce;
    }

    public function hasNonce()
    {
        return isset($this->nonce);
    }

    public function clearNonce()
    {
        unset($this->nonce);
    }

    /**
     * Nonce to use in swap validator state account PDA
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId nonce = 10 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\SolanaAccountId $var
     * @return $this
     */
    public function setNonce($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\SolanaAccountId::class);
        $this->nonce = $var;

        return $this;
    }

}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(ServerParameters::class, \Code\Transaction\V2\SwapResponse_ServerParameters::class);

