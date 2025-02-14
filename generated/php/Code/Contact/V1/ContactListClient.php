<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Code\Contact\V1;

/**
 */
class ContactListClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * AddContacts adds a batch of contacts to a user's contact list
     * @param \Code\Contact\V1\AddContactsRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function AddContacts(\Code\Contact\V1\AddContactsRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.contact.v1.ContactList/AddContacts',
        $argument,
        ['\Code\Contact\V1\AddContactsResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * RemoveContacts removes a batch of contacts from a user's contact list
     * @param \Code\Contact\V1\RemoveContactsRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function RemoveContacts(\Code\Contact\V1\RemoveContactsRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.contact.v1.ContactList/RemoveContacts',
        $argument,
        ['\Code\Contact\V1\RemoveContactsResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * GetContacts gets a subset of contacts from a user's contact list
     * @param \Code\Contact\V1\GetContactsRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetContacts(\Code\Contact\V1\GetContactsRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/code.contact.v1.ContactList/GetContacts',
        $argument,
        ['\Code\Contact\V1\GetContactsResponse', 'decode'],
        $metadata, $options);
    }

}
