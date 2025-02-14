<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Code\Push\V1;

/**
 */
class PushClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * AddToken stores a push token in a data container. The call is idempotent
     * and adding an existing valid token will not fail. Token types will be
     * validated against the user agent and any mismatches will result in an
     * INVALID_ARGUMENT status error.
     *
     * The token will be unlinked from any and all other accounts that it was
     * previously bound to.
     * @param \Code\Push\V1\AddTokenRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function AddToken(\Code\Push\V1\AddTokenRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.push.v1.Push/AddToken',
        $argument,
        ['\Code\Push\V1\AddTokenResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * RemoveToken removes the provided push token from the account.
     *
     * The provided token must be bound to the current account.
     * Otherwise, the RPC will succeed with without removal.
     * @param \Code\Push\V1\RemoveTokenRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function RemoveToken(\Code\Push\V1\RemoveTokenRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.push.v1.Push/RemoveToken',
        $argument,
        ['\Code\Push\V1\RemoveTokenResponse', 'decode'],
        $metadata, $options);
    }

}
