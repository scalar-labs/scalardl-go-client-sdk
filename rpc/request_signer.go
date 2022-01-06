package rpc

import (
	"encoding/binary"
	"unsafe"

	"github.com/scalar-labs/dl/v3/crypto"
)

// SignWith signs ContractRegistrationRequest with the given signer and fill the signature.
func (r *ContractRegistrationRequest) SignWith(signer crypto.Signer) (err error) {
	var certVersionBytes []byte = make([]byte, unsafe.Sizeof(r.GetCertVersion()))
	binary.BigEndian.PutUint32(certVersionBytes, r.GetCertVersion())

	var signing []byte
	signing = append(signing, []byte(r.GetContractId())...)
	signing = append(signing, []byte(r.GetContractBinaryName())...)
	signing = append(signing, r.GetContractByteCode()...)
	signing = append(signing, []byte(r.GetContractProperties())...)
	signing = append(signing, []byte(r.GetCertHolderId())...)
	signing = append(signing, certVersionBytes...)

	r.Signature, err = signer.Sign(signing)

	return
}

// SignWith signs ContractsListingRequest with the given signer and fill the signature.
func (r *ContractsListingRequest) SignWith(signer crypto.Signer) (err error) {
	var certVersionBytes []byte = make([]byte, unsafe.Sizeof(r.GetCertVersion()))
	binary.BigEndian.PutUint32(certVersionBytes, r.GetCertVersion())

	var signing []byte
	signing = append(signing, []byte(r.GetContractId())...)
	signing = append(signing, []byte(r.GetCertHolderId())...)
	signing = append(signing, certVersionBytes...)

	r.Signature, err = signer.Sign(signing)

	return
}

// SignWith signs ContractExecutionRequest with the given signer and fill the signature.
func (r *ContractExecutionRequest) SignWith(signer crypto.Signer) (err error) {
	var certVersionBytes []byte = make([]byte, unsafe.Sizeof(r.GetCertVersion()))
	binary.BigEndian.PutUint32(certVersionBytes, r.GetCertVersion())

	var signing []byte
	signing = append(signing, []byte(r.GetContractId())...)
	signing = append(signing, []byte(r.GetContractArgument())...)
	signing = append(signing, []byte(r.GetCertHolderId())...)
	signing = append(signing, certVersionBytes...)

	r.Signature, err = signer.Sign(signing)

	return
}

// SignWith signs LedgerValidationRequest with the given signer and fill the signature
func (r *LedgerValidationRequest) SignWith(signer crypto.Signer) (err error) {
	var (
		startAgeBytes    []byte = make([]byte, unsafe.Sizeof(r.GetStartAge()))
		endAgeBytes      []byte = make([]byte, unsafe.Sizeof(r.GetEndAge()))
		certVersionBytes []byte = make([]byte, unsafe.Sizeof(r.GetCertVersion()))
	)

	binary.BigEndian.PutUint32(startAgeBytes, r.GetStartAge())
	binary.BigEndian.PutUint32(endAgeBytes, r.GetEndAge())
	binary.BigEndian.PutUint32(certVersionBytes, r.GetCertVersion())

	var signing []byte
	signing = append(signing, []byte(r.GetAssetId())...)
	signing = append(signing, startAgeBytes...)
	signing = append(signing, endAgeBytes...)
	signing = append(signing, []byte(r.GetCertHolderId())...)
	signing = append(signing, certVersionBytes...)

	r.Signature, err = signer.Sign(signing)

	return
}

// SignWith signs LedgersValidationRequest with the given signer and fill the signature.
func (r *LedgersValidationRequest) SignWith(signer crypto.Signer) (err error) {
	var certVersionBytes []byte = make([]byte, unsafe.Sizeof(r.GetCertVersion()))
	binary.BigEndian.PutUint32(certVersionBytes, r.GetCertVersion())

	var signing []byte
	signing = append(signing, []byte(r.GetAssetId())...)
	signing = append(signing, []byte(r.GetCertHolderId())...)
	signing = append(signing, certVersionBytes...)

	r.Signature, err = signer.Sign(signing)

	return
}

// SignWith signs AssetProofRetrievalRequest with the given signer and fill the signature.
func (r *AssetProofRetrievalRequest) SignWith(signer crypto.Signer) (err error) {
	var (
		ageBytes         []byte = make([]byte, unsafe.Sizeof(r.GetAge()))
		certVersionBytes []byte = make([]byte, unsafe.Sizeof(r.GetCertVersion()))
	)

	binary.BigEndian.PutUint32(ageBytes, uint32(r.GetAge()))
	binary.BigEndian.PutUint32(certVersionBytes, r.GetCertVersion())

	var signing []byte
	signing = append(signing, []byte(r.GetAssetId())...)
	signing = append(signing, ageBytes...)
	signing = append(signing, []byte(r.GetCertHolderId())...)
	signing = append(signing, certVersionBytes...)

	r.Signature, err = signer.Sign(signing)

	return
}

// SignWith signs ExecutionAbortRequest with the given signer and fill the signature.
func (r *ExecutionAbortRequest) SignWith(signer crypto.Signer) (err error) {
	var certVersionBytes []byte = make([]byte, unsafe.Sizeof(r.GetCertVersion()))
	binary.BigEndian.PutUint32(certVersionBytes, r.GetCertVersion())

	var signing []byte
	signing = append(signing, []byte(r.GetNonce())...)
	signing = append(signing, []byte(r.GetCertHolderId())...)
	signing = append(signing, certVersionBytes...)

	r.Signature, err = signer.Sign(signing)

	return
}
