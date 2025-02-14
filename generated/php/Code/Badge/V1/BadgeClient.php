<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Code\Badge\V1;

/**
 */
class BadgeClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * ResetBadgeCount resets an owner account's app icon badge count back to zero
     * @param \Code\Badge\V1\ResetBadgeCountRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function ResetBadgeCount(\Code\Badge\V1\ResetBadgeCountRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.badge.v1.Badge/ResetBadgeCount',
        $argument,
        ['\Code\Badge\V1\ResetBadgeCountResponse', 'decode'],
        $metadata, $options);
    }

}
