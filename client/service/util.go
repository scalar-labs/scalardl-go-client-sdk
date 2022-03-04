package service

import (
	clientError "github.com/scalar-labs/scalardl-go-client-sdk/v3/client/error"
	"github.com/scalar-labs/scalardl-go-client-sdk/v3/ledger/statuscode"
	"github.com/scalar-labs/scalardl-go-client-sdk/v3/rpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

func getClientErrorFromTrailer(trailer metadata.MD) (err clientError.ClientError) {
	err = clientError.NewClientError(statuscode.UnknownTransactionStatus, "")

	var statusInTailer []string = trailer.Get("rpc.status-bin")
	if len(statusInTailer) == 0 {
		return
	}

	var (
		status        rpc.Status
		statusInBytes = []byte(statusInTailer[0])
	)
	if e := proto.Unmarshal(statusInBytes, &status); e == nil {
		err = clientError.NewClientError(
			statuscode.StatusCode(status.GetCode()),
			status.GetMessage(),
		)
	}

	return
}
