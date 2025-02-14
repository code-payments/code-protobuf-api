<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chat/v2/chat_service.proto

namespace Code\Chat\V2\NotifyIsTypingResponse;

use UnexpectedValueException;

/**
 * Protobuf type <code>code.chat.v2.NotifyIsTypingResponse.Result</code>
 */
class Result
{
    /**
     * Generated from protobuf enum <code>OK = 0;</code>
     */
    const OK = 0;
    /**
     * Generated from protobuf enum <code>DENIED = 1;</code>
     */
    const DENIED = 1;
    /**
     * Generated from protobuf enum <code>CHAT_NOT_FOUND = 2;</code>
     */
    const CHAT_NOT_FOUND = 2;

    private static $valueToName = [
        self::OK => 'OK',
        self::DENIED => 'DENIED',
        self::CHAT_NOT_FOUND => 'CHAT_NOT_FOUND',
    ];

    public static function name($value)
    {
        if (!isset(self::$valueToName[$value])) {
            throw new UnexpectedValueException(sprintf(
                    'Enum %s has no name defined for value %s', __CLASS__, $value));
        }
        return self::$valueToName[$value];
    }


    public static function value($name)
    {
        $const = __CLASS__ . '::' . strtoupper($name);
        if (!defined($const)) {
            throw new UnexpectedValueException(sprintf(
                    'Enum %s has no value defined for name %s', __CLASS__, $name));
        }
        return constant($const);
    }
}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(Result::class, \Code\Chat\V2\NotifyIsTypingResponse_Result::class);

