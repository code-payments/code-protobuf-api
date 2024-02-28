# Identity Service Protobuf Schema Documentation

This document provides detailed explanations and comments for the Identity Service Protobuf schema, including service operations, request and response messages, and data structures.

## Service: IdentityService

The `IdentityService` manages user identities, account linking, and preferences.

### RPC Operations

- **LinkAccount**: Links an owner account to the user identified and authenticated by a one-time use token. This operation automatically creates a new user if one does not exist and creates a new data container for at least every unique owner account linked to the user.

- **UnlinkAccount**: Removes links from an owner account but does not remove existing associations between users, owner accounts, and identifying features. This operation is idempotent and will not fail if the link has already been removed.

- **GetUser**: Retrieves user information given a user identifier and an owner account.

- **UpdatePreferences**: Updates user preferences, including server-side localization settings.

- **LoginToThirdPartyApp**: Logs a user into a third-party app for a given intent ID. If the request requires payment, `SubmitIntent` must be called.

- **GetLoginForThirdPartyApp**: Retrieves a login for a third-party app from an existing request, supporting all paths where login is possible.

### Messages

#### LinkAccountRequest and Response

- `LinkAccountRequest` contains the public key of the owner account to be linked, a signature to validate ownership, and a one-time use token for user identification and authentication.
- `LinkAccountResponse` includes the result of the operation, user information, and a data container ID for storing user data. It may also contain metadata related to the phone-based identifying feature.

#### UnlinkAccountRequest and Response

- `UnlinkAccountRequest` includes the public key of the owner account to be unlinked and a signature for authentication.
- `UnlinkAccountResponse` provides the result of the unlink operation.

#### GetUserRequest and Response

- `GetUserRequest` requires the public key of the owner account and a signature for authentication.
- `GetUserResponse` returns user information, including the data container ID and any applicable metadata.

#### UpdatePreferencesRequest and Response

- `UpdatePreferencesRequest` is used to update user preferences, requiring the public key of the owner account, a data container ID for the contact list, and a signature. It also includes the user's locale for localization purposes.
- `UpdatePreferencesResponse` indicates the result of the update operation.

#### LoginToThirdPartyAppRequest and Response

- `LoginToThirdPartyAppRequest` and `GetLoginForThirdPartyAppRequest` are used to manage third-party app logins, requiring intent IDs, user or verifier account information, and signatures for authentication.
- The corresponding responses indicate the outcome of the login attempt and provide relevant user information.

### Data Structures

- **User**: Represents the highest order of identity within Code, including a unique user ID and a view containing identifying features.
- **View**: A set of identifying features for a user, currently limited to a phone number.
- **PhoneMetadata**: Provides state information about whether a phone number is linked to an owner account.

This documentation should accompany the refactored Protobuf schema to provide a comprehensive understanding of its structure and functionality.