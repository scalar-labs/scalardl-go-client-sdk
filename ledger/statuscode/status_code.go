package statuscode

// StatusCode hints at what happened after a gRPC request is handled.
type StatusCode int

const (
	// OK indicates that the registration/execution request has succeeded.
	OK = 200

	// InvalidHash indicates that the existing hash value of an asset record is different from the expected value.
	InvalidHash = 300

	// InvalidPrevHash indicates that the existing prev_hash value of an asset record is different from the expected value.
	InvalidPrevHash = 301

	// InvalidContract indicates that some previously executed contract produced an asset record which could not be validated.
	InvalidContract = 302

	// InvalidOutput indicates that the existing data value of an asset record is different from the expected value.
	InvalidOutput = 303

	// InvalidNonce indicates that the same nonce value has been used more than once.
	InvalidNonce = 304

	// InconsistentStates indicates that the ledger states between multiple organizations are inconsistent.
	InconsistentStates = 305

	// InconsistentRequest indicates that a request is inconsistent and could be maliciously tampered.
	InconsistentRequest = 306

	// InvalidSignature indicates that the given signature is invalid or a signature can not be created for some reason.
	InvalidSignature = 400

	// UnloadableKey indicates that the given key could not be loaded for some reason, e.g. it is invalid.
	UnloadableKey = 401

	// UnloadableContract indicates that the given contract could not be loaded for some reason, e.g., instantiation failure.
	UnloadableContract = 402

	// CertificateNotFound indicates that the given certificate is not found.
	CertificateNotFound = 403

	// ContractNotFound indicates that the given contract is not found.
	ContractNotFound = 404

	// CertificateAlreadyRegistered indicates that the given certificate is already registered.
	CertificateAlreadyRegistered = 405

	// ContractAlreadyRegistered indicates that the given contract is already registered.
	ContractAlreadyRegistered = 406

	// InvalidRequest indicates that the request is invalid.
	InvalidRequest = 407

	// ContractContextualError indicates that the contract has a contextual error that is not recoverable by the ledger.
	ContractContextualError = 408

	// AssetNotFound indicates that the specified asset is not found.
	AssetNotFound = 409

	// FunctionNotFound indicates that the given function is not found.
	FunctionNotFound = 410

	// UnloadableFunction indicates that the given function could not be loaded for some reason.
	UnloadableFunction = 411

	// InvalidFunction indicates that the given function is invalid.
	InvalidFunction = 412

	// DatabaseError indicates that the system encountered a database error such as IO error.
	DatabaseError = 500

	// UnknownTransactionStatus indicates that the system encountered a unknown transaction status.
	UnknownTransactionStatus = 501

	// RuntimeError indicates that the system encountered a runtime error.
	RuntimeError = 502

	// Unavailable indicates that the system is temporarily unavailable.
	Unavailable = 503

	// Conflict indicates that the system encountered conflicting transactions.
	Conflict = 504
)
