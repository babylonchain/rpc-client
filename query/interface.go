package query

import (
	"context"

	btcctypes "github.com/babylonchain/babylon/x/btccheckpoint/types"
	btclctypes "github.com/babylonchain/babylon/x/btclightclient/types"
	checkpointingtypes "github.com/babylonchain/babylon/x/checkpointing/types"
	epochingtypes "github.com/babylonchain/babylon/x/epoching/types"
	monitortypes "github.com/babylonchain/babylon/x/monitor/types"
	zctypes "github.com/babylonchain/babylon/x/zoneconcierge/types"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	sdkquerytypes "github.com/cosmos/cosmos-sdk/types/query"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

type BabylonQueryClient interface {
	Stop()

	// APIs for BTCCheckpoint
	QueryBTCCheckpoint(f func(ctx context.Context, queryClient btcctypes.QueryClient) error) error
	BTCCheckpointParams() (*btcctypes.QueryParamsResponse, error)
	BTCPositionAtEpoch(epochNumber uint64) (*btcctypes.QueryBtcCheckpointInfoResponse, error)
	BTCPositionForEpochRange(startEpoch uint64, endEpoch uint64, pagination *sdkquerytypes.PageRequest) (*btcctypes.QueryBtcCheckpointsInfoResponse, error)

	// APIs for BTCLightclient
	QueryBTCLightclient(f func(ctx context.Context, queryClient btclctypes.QueryClient) error) error
	BTCLightClientParams() (*btclctypes.QueryParamsResponse, error)
	BTCHeaderChainTip() (*btclctypes.QueryTipResponse, error)
	BTCBaseHeader() (*btclctypes.QueryBaseHeaderResponse, error)
	ContainsBTCBlock(blockHash *chainhash.Hash) (*btclctypes.QueryContainsBytesResponse, error)
	BTCMainChain(pagination *sdkquerytypes.PageRequest) (*btclctypes.QueryMainChainResponse, error)

	// APIs for Checkpointing
	QueryCheckpointing(f func(ctx context.Context, queryClient checkpointingtypes.QueryClient) error) error
	RawCheckpoint(epochNumber uint64) (*checkpointingtypes.QueryRawCheckpointResponse, error)
	RawCheckpointList(status checkpointingtypes.CheckpointStatus, pagination *sdkquerytypes.PageRequest) (*checkpointingtypes.QueryRawCheckpointListResponse, error)
	BlsPublicKeyList(epochNumber uint64, pagination *sdkquerytypes.PageRequest) (*checkpointingtypes.QueryBlsPublicKeyListResponse, error)
	EpochStatusCount(epochCount uint64) (*checkpointingtypes.QueryRecentEpochStatusCountResponse, error)
	LatestEpochFromStatus(status checkpointingtypes.CheckpointStatus) (*checkpointingtypes.QueryLastCheckpointWithStatusResponse, error)

	// APIs for Epoching
	QueryEpoching(f func(ctx context.Context, queryClient epochingtypes.QueryClient) error) error
	EpochingParams() (*epochingtypes.QueryParamsResponse, error)
	CurrentEpoch() (*epochingtypes.QueryCurrentEpochResponse, error)
	EpochsInfoForEpochRange(startEpoch uint64, endEpoch uint64) (*epochingtypes.QueryEpochsInfoResponse, error)
	EpochsInfo(pagination *sdkquerytypes.PageRequest) (*epochingtypes.QueryEpochsInfoResponse, error)
	LatestEpochMsgs(endEpoch uint64, epochCount uint64, pagination *sdkquerytypes.PageRequest) (*epochingtypes.QueryLatestEpochMsgsResponse, error)
	DelegationLifecycle(delegator string) (*epochingtypes.QueryDelegationLifecycleResponse, error)

	// APIs for Monitor
	QueryMonitor(f func(ctx context.Context, queryClient monitortypes.QueryClient) error) error
	EndedEpochBTCHeight(epochNum uint64) (*monitortypes.QueryEndedEpochBtcHeightResponse, error)
	ReportedCheckpointBTCHeight(hashStr string) (*monitortypes.QueryReportedCheckpointBtcHeightResponse, error)

	// APIs for ZoneConcierge
	QueryZoneConcierge(f func(ctx context.Context, queryClient zctypes.QueryClient) error) error
	FinalizedConnectedChainInfo(chainID string) (*zctypes.QueryFinalizedChainInfoResponse, error)
	ConnectedChainInfo(chainID string) (*zctypes.QueryChainInfoResponse, error)
	ConnectedChainList() (*zctypes.QueryChainListResponse, error)
	ConnectedChainHeaders(chainID string, pagination *sdkquerytypes.PageRequest) (*zctypes.QueryListHeadersResponse, error)
	ConnectedChainEpochInfo(chainID string, epochNum uint64) (*zctypes.QueryEpochChainInfoResponse, error)

	// APIs for Staking
	QueryStaking(f func(ctx context.Context, queryClient stakingtypes.QueryClient) error) error
	StakingParams() (*stakingtypes.QueryParamsResponse, error)

	// APIs for Tendermint
	GetBlock(height int64) (*coretypes.ResultBlock, error)
	TxSearch(events []string, prove bool, page *int, perPage *int, orderBy string) (*coretypes.ResultTxSearch, error)
	GetTx(hash []byte) (*coretypes.ResultTx, error)
}
