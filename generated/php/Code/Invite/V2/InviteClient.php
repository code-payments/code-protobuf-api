<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Code\Invite\V2;

/**
 */
class InviteClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * GetInviteCount gets the number of invites that a user can send out.
     * @param \Code\Invite\V2\GetInviteCountRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetInviteCount(\Code\Invite\V2\GetInviteCountRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.invite.v2.Invite/GetInviteCount',
        $argument,
        ['\Code\Invite\V2\GetInviteCountResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * InvitePhoneNumber invites someone to join via their phone number. A phone number
     * can only be invited once by a unique user or invite code. This is to avoid having
     * a phone number consuming more than one invite count globally.
     * @param \Code\Invite\V2\InvitePhoneNumberRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function InvitePhoneNumber(\Code\Invite\V2\InvitePhoneNumberRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.invite.v2.Invite/InvitePhoneNumber',
        $argument,
        ['\Code\Invite\V2\InvitePhoneNumberResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * GetInvitationStatus gets a phone number's invitation status.
     * @param \Code\Invite\V2\GetInvitationStatusRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetInvitationStatus(\Code\Invite\V2\GetInvitationStatusRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.invite.v2.Invite/GetInvitationStatus',
        $argument,
        ['\Code\Invite\V2\GetInvitationStatusResponse', 'decode'],
        $metadata, $options);
    }

}
