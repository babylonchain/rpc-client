package client

import (
	btcctypes "github.com/babylonchain/babylon/x/btccheckpoint/types"
	btclctypes "github.com/babylonchain/babylon/x/btclightclient/types"
	checkpointingtypes "github.com/babylonchain/babylon/x/checkpointing/types"
	epochingtypes "github.com/babylonchain/babylon/x/epoching/types"
	"github.com/babylonchain/rpc-client/config"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type BabylonClient interface {
	Stop()
	GetConfig() *config.BabylonConfig
	GetTagIdx() uint8
	GetAddr() (sdk.AccAddress, error)
	MustGetAddr() sdk.AccAddress
	InsertBTCSpvProof(msg *btcctypes.MsgInsertBTCSpvProof) (*sdk.TxResponse, error)
	InsertHeader(msg *btclctypes.MsgInsertHeader) (*sdk.TxResponse, error)
	InsertHeaders(msgs []*btclctypes.MsgInsertHeader) (*sdk.TxResponse, error)
	MustInsertBTCSpvProof(msg *btcctypes.MsgInsertBTCSpvProof) *sdk.TxResponse

	// staking module related queries
	QueryStakingParams() (*stakingtypes.Params, error)

	// epoch module related queries
	QueryEpochingParams() (*epochingtypes.Params, error)
	QueryCurrentEpoch() (uint64, error)

	// btclightclient module related queries
	QueryBTCLightclientParams() (*btclctypes.Params, error)
	QueryHeaderChainTip() (*chainhash.Hash, uint64, error)
	QueryBaseHeader() (*wire.BlockHeader, uint64, error)
	QueryContainsBlock(blockHash *chainhash.Hash) (bool, error)

	// btccheckpoint module related queries
	QueryBTCCheckpointParams() (*btcctypes.Params, error)
	MustQueryBTCCheckpointParams() *btcctypes.Params

	// checkpointing module related queries
	QueryRawCheckpoint(epochNumber uint64) (*checkpointingtypes.RawCheckpointWithMeta, error)
	QueryRawCheckpointList(status checkpointingtypes.CheckpointStatus) ([]*checkpointingtypes.RawCheckpointWithMeta, error)
	BlsPublicKeyList(epochNumber uint64) ([]*checkpointingtypes.ValidatorWithBlsKey, error)

	// monitor module related queries
	QueryFinishedEpochBtcHeight(epochNum uint64) (uint64, error)
	QueryReportedCheckpointBtcHeight(hash []byte) (uint64, error)
}
