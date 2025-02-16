<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: contact/v1/contact_list_service.proto

namespace Code\Contact\V1\RemoveContactsResponse;

use UnexpectedValueException;

/**
 * Protobuf type <code>code.contact.v1.RemoveContactsResponse.Result</code>
 */
class Result
{
    /**
     * Generated from protobuf enum <code>OK = 0;</code>
     */
    const OK = 0;

    private static $valueToName = [
        self::OK => 'OK',
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
class_alias(Result::class, \Code\Contact\V1\RemoveContactsResponse_Result::class);

