<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Code\Micropayment\V1;

/**
 * todo: Migrate this to a generic "request" service
 */
class MicroPaymentClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * GetStatus gets basic request status
     * @param \Code\Micropayment\V1\GetStatusRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetStatus(\Code\Micropayment\V1\GetStatusRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.micropayment.v1.MicroPayment/GetStatus',
        $argument,
        ['\Code\Micropayment\V1\GetStatusResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * RegisterWebhook registers a webhook for a request
     *
     * todo: Once Kik codes can encode the entire payment request details, we can
     *       remove the messaging service component and have a Create RPC that
     *       reserves the intent ID with payment details, plus registers the webhook
     *       at the same time. Until that's possible, we're stuck with two RPC calls.
     * @param \Code\Micropayment\V1\RegisterWebhookRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function RegisterWebhook(\Code\Micropayment\V1\RegisterWebhookRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.micropayment.v1.MicroPayment/RegisterWebhook',
        $argument,
        ['\Code\Micropayment\V1\RegisterWebhookResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * Codify adds a trial micro paywall to any URL
     * @param \Code\Micropayment\V1\CodifyRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function Codify(\Code\Micropayment\V1\CodifyRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.micropayment.v1.MicroPayment/Codify',
        $argument,
        ['\Code\Micropayment\V1\CodifyResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * GetPathMetadata gets codified website metadata for a given path
     *
     * Important Note: This RPC's current implementation is insecure and
     * it's sole design is to enable PoC and trials.
     * @param \Code\Micropayment\V1\GetPathMetadataRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetPathMetadata(\Code\Micropayment\V1\GetPathMetadataRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.micropayment.v1.MicroPayment/GetPathMetadata',
        $argument,
        ['\Code\Micropayment\V1\GetPathMetadataResponse', 'decode'],
        $metadata, $options);
    }

}
