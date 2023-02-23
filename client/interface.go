package client

import (
	btcctypes "github.com/babylonchain/babylon/x/btccheckpoint/types"
	btclctypes "github.com/babylonchain/babylon/x/btclightclient/types"
	"github.com/babylonchain/rpc-client/config"
	bbnquery "github.com/babylonchain/rpc-client/query"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type BabylonClient interface {
	bbnquery.BabylonQueryClient

	GetConfig() *config.BabylonConfig
	GetTagIdx() uint8
	GetAddr() (sdk.AccAddress, error)
	MustGetAddr() sdk.AccAddress

	InsertBTCSpvProof(msg *btcctypes.MsgInsertBTCSpvProof) (*sdk.TxResponse, error)
	InsertHeader(msg *btclctypes.MsgInsertHeader) (*sdk.TxResponse, error)
	InsertHeaders(msgs []*btclctypes.MsgInsertHeader) (*sdk.TxResponse, error)
}
