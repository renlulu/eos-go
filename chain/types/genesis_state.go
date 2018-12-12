package types

import (
	"github.com/eosspark/eos-go/common"
	"github.com/eosspark/eos-go/crypto"
	"github.com/eosspark/eos-go/crypto/ecc"
	"github.com/eosspark/eos-go/crypto/rlp"
	"github.com/eosspark/eos-go/log"
)

/*var isActiveGenesis bool = false
var instance = &GenesisState{}*/
type GenesisState struct {
	InitialTimestamp common.TimePoint `json:"initial_timestamp"`
	InitialKey       ecc.PublicKey    `json:"initial_key"`
}

const EosioRootKey = "EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV"

func NewGenesisState() *GenesisState {
	g := &GenesisState{}
	its, err := common.FromIsoString("2018-06-01T12:00:00")
	if err != nil {
		log.Error("NewGenesisState is error detail:", err)
	}
	//g.EosioRootKey = "EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV"
	g.InitialTimestamp = its
	key, err := ecc.NewPublicKey(EosioRootKey)
	if err != nil {
		log.Error("", err)
	}
	g.InitialKey = key
	g.Initial()
	return g
}

func (g *GenesisState) ComputeChainID() common.ChainIdType {

	b, err := rlp.EncodeToBytes(g)
	if err != nil {
		log.Error("ComputeChainID EncodeToBytes is error:", err)
	}
	return common.ChainIdType(crypto.Hash256(b))
}

func (g *GenesisState) Initial() ChainConfig {
	InitialConfiguration := ChainConfig{
		MaxBlockNetUsage:               common.DefaultConfig.MaxBlockNetUsage,
		TargetBlockNetUsagePct:         common.DefaultConfig.TargetBlockNetUsagePct,
		MaxTransactionNetUsage:         common.DefaultConfig.MaxTransactionNetUsage,
		BasePerTransactionNetUsage:     common.DefaultConfig.BasePerTransactionNetUsage,
		NetUsageLeeway:                 common.DefaultConfig.NetUsageLeeway,
		ContextFreeDiscountNetUsageNum: common.DefaultConfig.ContextFreeDiscountNetUsageNum,
		ContextFreeDiscountNetUsageDen: common.DefaultConfig.ContextFreeDiscountNetUsageDen,

		MaxBlockCpuUsage:       common.DefaultConfig.MaxBlockCpuUsage,
		TargetBlockCpuUsagePct: common.DefaultConfig.TargetBlockCpuUsagePct,
		MaxTransactionCpuUsage: common.DefaultConfig.MaxTransactionCpuUsage,
		MinTransactionCpuUsage: common.DefaultConfig.MinTransactionCpuUsage,

		MaxTrxLifetime:              common.DefaultConfig.MaxTrxLifetime,
		DeferredTrxExpirationWindow: common.DefaultConfig.DeferredTrxExpirationWindow,
		MaxTrxDelay:                 common.DefaultConfig.MaxTrxDelay,
		MaxInlineActionSize:         common.DefaultConfig.MaxInlineActionSize,
		MaxInlineActionDepth:        common.DefaultConfig.MaxInlineActionDepth,
		MaxAuthorityDepth:           common.DefaultConfig.MaxAuthorityDepth,
	}

	return InitialConfiguration
}
