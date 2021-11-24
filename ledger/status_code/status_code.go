package status_code

// StatusCode hints at what happened after a gRPC request is handled.
type StatusCode int

const (
	// OK indicates that the registration/execution request has succeeded.
	OK = 200

	// INVALID_HASH indicates that the existing hash value of an asset record is different from the expected value.
	INVALID_HASH = 300

	// INVALID_PREV_HASH indicates that the existing prev_hash value of an asset record is different from the expected value.
	INVALID_PREV_HASH = 301

	// INVALID_CONTRACT indicates that some previously executed contract produced an asset record which could not be validated.
	INVALID_CONTRACT = 302

	// INVALID_OUTPUT indicates that the existing data value of an asset record is different from the expected value.
	INVALID_OUTPUT = 303

	// INVALID_NONCE indicates that the same nonce value has been used more than once.
	INVALID_NONCE = 304

	// INCONSISTENT_STATES indicates that the ledger states between multiple organizations are inconsistent.
	INCONSISTENT_STATES = 305

	// INCONSISTENT_REQUEST indicates that a request is inconsistent and could be maliciously tampered.
	INCONSISTENT_REQUEST = 306

	// INVALID_SIGNATURE indicates that the given signature is invalid or a signature can not be created for some reason.
	INVALID_SIGNATURE = 400

	// UNLOADABLE_KEY indicates that the given key could not be loaded for some reason, e.g. it is invalid.
	UNLOADABLE_KEY = 401

	// UNLOADABLE_CONTRACT indicates that the given contract could not be loaded for some reason, e.g., instantiation failure.
	UNLOADABLE_CONTRACT = 402

	// CERTIFICATE_NOT_FOUND indicates that the given certificate is not found.
	CERTIFICATE_NOT_FOUND = 403

	// CONTRACT_NOT_FOUND indicates that the given contract is not found.
	CONTRACT_NOT_FOUND = 404

	// CERTIFICATE_ALREADY_REGISTERED indicates that the given certificate is already registered.
	CERTIFICATE_ALREADY_REGISTERED = 405

	// CONTRACT_ALREADY_REGISTERED indicates that the given contract is already registered.
	CONTRACT_ALREADY_REGISTERED = 406

	// INVALID_REQUEST indicates that the request is invalid.
	INVALID_REQUEST = 407

	// CONTRACT_CONTEXTUAL_ERROR indicates that the contract has a contextual error that is not recoverable by the ledger.
	CONTRACT_CONTEXTUAL_ERROR = 408

	// ASSET_NOT_FOUND indicates that the specified asset is not found.
	ASSET_NOT_FOUND = 409

	// FUNCTION_NOT_FOUND indicates that the given function is not found.
	FUNCTION_NOT_FOUND = 410

	// UNLOADABLE_FUNCTION indicates that the given function could not be loaded for some reason.
	UNLOADABLE_FUNCTION = 411

	// INVALID_FUNCTION indicates that the given function is invalid.
	INVALID_FUNCTION = 412

	// DATABASE_ERROR indicates that the system encountered a database error such as IO error.
	DATABASE_ERROR = 500

	// UNKNOWN_TRANSACTION_STATUS indicates that the system encountered a unknown transaction status.
	UNKNOWN_TRANSACTION_STATUS = 501

	// RUNTIME_ERROR indicates that the system encountered a runtime error.
	RUNTIME_ERROR = 502

	// UNAVAILABLE indicates that the system is temporarily unavailable.
	UNAVAILABLE = 503

	// CONFLICT indicates that the system encountered conflicting transactions.
	CONFLICT = 504
)
