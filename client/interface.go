package client

import (
	btcctypes "github.com/babylonchain/babylon/x/btccheckpoint/types"
	btclctypes "github.com/babylonchain/babylon/x/btclightclient/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/babylonchain/rpc-client/config"
	bbnquery "github.com/babylonchain/rpc-client/query"
)

type BabylonClient interface {
	bbnquery.BabylonQueryClient

	GetConfig() *config.BabylonConfig
	GetAddr() (sdk.AccAddress, error)
	MustGetAddr() sdk.AccAddress

	InsertBTCSpvProof(msg *btcctypes.MsgInsertBTCSpvProof) (*sdk.TxResponse, error)
	InsertHeader(msg *btclctypes.MsgInsertHeader) (*sdk.TxResponse, error)
	InsertHeaders(msgs []*btclctypes.MsgInsertHeader) (*sdk.TxResponse, error)
}
