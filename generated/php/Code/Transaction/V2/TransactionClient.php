<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Code\Transaction\V2;

/**
 */
class TransactionClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * SubmitIntent is the mechanism for client and server to agree upon a set of
     * client actions to execute on the blockchain using the Code sequencer for
     * fulfillment.
     *
     * Transactions are never exchanged between client and server. Instead, the
     * required accounts and arguments for instructions known to each actor are
     * exchanged to allow independent and local transaction construction.
     *
     * Client and server are expected to fully validate the intent. Proofs will
     * be provided for any parameter requiring one. Signatures should only be
     * generated after approval of all transactions.
     *
     * This RPC is not a traditional streaming endpoint. It bundles two unary calls
     * to enable DB-level transaction semantics.
     *
     * The high-level happy path flow for the RPC is as follows:
     *   1.  Client initiates a stream and sends SubmitIntentRequest.SubmitActions
     *   2.  Server validates the intent, its actions and metadata
     *   3a. If there are transactions requiring the user's signature, then server
     *       returns SubmitIntentResponse.ServerParameters
     *   3b. Otherwise, server returns SubmitIntentResponse.Success and closes the
     *       stream
     *   4.  For each transaction requiring the user's signature, the client locally
     *       constructs it, performs validation and collects the signature
     *   5.  Client sends SubmitIntentRequest.SubmitSignatures with the signature
     *       list generated from 4
     *   6.  Server validates all signatures are submitted and are the expected values
     *       using locally constructed transactions.
     *   7.  Server returns SubmitIntentResponse.Success and closes the stream
     * In the error case:
     *   * Server will return SubmitIntentResponse.Error and close the stream
     *   * Client will close the stream
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\BidiStreamingCall
     */
    public function SubmitIntent($metadata = [], $options = []) {
        return $this->_bidiRequest('/code.transaction.v2.Transaction/SubmitIntent',
        ['\Code\Transaction\V2\SubmitIntentResponse','decode'],
        $metadata, $options);
    }

    /**
     * GetIntentMetadata gets basic metadata on an intent. It can also be used
     * to fetch the status of submitted intents. Metadata exists only for intents
     * that have been successfully submitted.
     * @param \Code\Transaction\V2\GetIntentMetadataRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetIntentMetadata(\Code\Transaction\V2\GetIntentMetadataRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.transaction.v2.Transaction/GetIntentMetadata',
        $argument,
        ['\Code\Transaction\V2\GetIntentMetadataResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * GetPrivacyUpgradeStatus gets the status of a private transaction and the
     * ability to upgrade it to permanent privacy.
     * @param \Code\Transaction\V2\GetPrivacyUpgradeStatusRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetPrivacyUpgradeStatus(\Code\Transaction\V2\GetPrivacyUpgradeStatusRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.transaction.v2.Transaction/GetPrivacyUpgradeStatus',
        $argument,
        ['\Code\Transaction\V2\GetPrivacyUpgradeStatusResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * GetPrioritizedIntentsForPrivacyUpgrade allows clients to get private
     * intent actions that can be upgraded in a secure and verifiable manner.
     * @param \Code\Transaction\V2\GetPrioritizedIntentsForPrivacyUpgradeRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetPrioritizedIntentsForPrivacyUpgrade(\Code\Transaction\V2\GetPrioritizedIntentsForPrivacyUpgradeRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.transaction.v2.Transaction/GetPrioritizedIntentsForPrivacyUpgrade',
        $argument,
        ['\Code\Transaction\V2\GetPrioritizedIntentsForPrivacyUpgradeResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * GetLimits gets limits for money moving intents for an owner account in an
     * identity-aware manner
     * @param \Code\Transaction\V2\GetLimitsRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetLimits(\Code\Transaction\V2\GetLimitsRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.transaction.v2.Transaction/GetLimits',
        $argument,
        ['\Code\Transaction\V2\GetLimitsResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * GetPaymentHistory gets an owner account's payment history inferred from intents
     *
     * Deprecated: Payment history has migrated to chats
     * @param \Code\Transaction\V2\GetPaymentHistoryRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetPaymentHistory(\Code\Transaction\V2\GetPaymentHistoryRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.transaction.v2.Transaction/GetPaymentHistory',
        $argument,
        ['\Code\Transaction\V2\GetPaymentHistoryResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * CanWithdrawToAccount provides hints to clients for submitting withdraw intents.
     * The RPC indicates if a withdrawal is possible, and how it should be performed.
     * @param \Code\Transaction\V2\CanWithdrawToAccountRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function CanWithdrawToAccount(\Code\Transaction\V2\CanWithdrawToAccountRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.transaction.v2.Transaction/CanWithdrawToAccount',
        $argument,
        ['\Code\Transaction\V2\CanWithdrawToAccountResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * Airdrop airdrops Kin to the requesting account
     * @param \Code\Transaction\V2\AirdropRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function Airdrop(\Code\Transaction\V2\AirdropRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.transaction.v2.Transaction/Airdrop',
        $argument,
        ['\Code\Transaction\V2\AirdropResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * Swap performs an on-chain swap. The high-level flow mirrors SubmitIntent
     * closely. However, due to the time-sensitive nature and unreliability of
     * swaps, they do not fit within the broader intent system. This results in
     * a few key differences:
     *  * Transactions are submitted on a best-effort basis outside of the Code
     *    Sequencer within the RPC handler
     *  * Balance changes are applied after the transaction has finalized
     *  * Transactions use recent blockhashes over a nonce
     *
     * The transaction will have the following instruction format:
     *   1. ComputeBudget::SetComputeUnitLimit
     *   2. ComputeBudget::SetComputeUnitPrice
     *   3. SwapValidator::PreSwap
     *   4. Dynamic swap instruction
     *   5. SwapValidator::PostSwap
     *
     * Note: Currently limited to swapping USDC to Kin.
     * Note: Kin is deposited into the primary account.
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\BidiStreamingCall
     */
    public function Swap($metadata = [], $options = []) {
        return $this->_bidiRequest('/code.transaction.v2.Transaction/Swap',
        ['\Code\Transaction\V2\SwapResponse','decode'],
        $metadata, $options);
    }

    /**
     * DeclareFiatOnrampPurchaseAttempt is called whenever a user attempts to use a fiat
     * onramp to purchase crypto for use in Code.
     * @param \Code\Transaction\V2\DeclareFiatOnrampPurchaseAttemptRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function DeclareFiatOnrampPurchaseAttempt(\Code\Transaction\V2\DeclareFiatOnrampPurchaseAttemptRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.transaction.v2.Transaction/DeclareFiatOnrampPurchaseAttempt',
        $argument,
        ['\Code\Transaction\V2\DeclareFiatOnrampPurchaseAttemptResponse', 'decode'],
        $metadata, $options);
    }

}
