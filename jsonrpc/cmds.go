// Copyright (c) 2014-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// NOTE: This file is intended to house the RPC commands that are supported by
// a chain server.

package jsonrpc

import (
	"github.com/uxuycom/indexer/model"
)

// EmptyCmd defines the empty JSON-RPC command.
type EmptyCmd struct{}

// FindAllInscriptionsCmd defines the inscription JSON-RPC command.
type FindAllInscriptionsCmd struct {
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
	Chain    string `json:"chain"`
	Protocol string `json:"protocol"`
	Tick     string `json:"tick"`
	DeployBy string `json:"deploy_by"`
	Sort     int    `json:"sort"`
	//SortMode int    `json:"sort_mode"`
}

type IndsGetTicksCmd struct {
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
	Chain    string `json:"chain"`
	Protocol string `json:"protocol"`
	Tick     string `json:"tick"`
	DeployBy string `json:"deploy_by"`
	Sort     int    `json:"sort"`
	SortMode int    `json:"sort_mode"`
}

type IndsGetTickCmd struct {
	Chain      string
	Protocol   string
	Tick       string
	DeployHash string
}

type FindAllInscriptionsResponse struct {
	Inscriptions interface{} `json:"inscriptions"`
	Total        int64       `json:"total"`
	Limit        int         `json:"limit"`
	Offset       int         `json:"offset"`
}

type InscriptionInfo struct {
	Chain        string `json:"chain"`
	Protocol     string `json:"protocol"`
	Tick         string `json:"tick"`
	Name         string `json:"name"`
	LimitPerMint string `json:"limit_per_mint"`
	DeployBy     string `json:"deploy_by"`
	TotalSupply  string `json:"total_supply"`
	DeployHash   string `json:"deploy_hash"`
	DeployTime   uint32 `json:"deploy_time"`
	TransferType int8   `json:"transfer_type"`
	CreatedAt    uint32 `json:"created_at"`
	UpdatedAt    uint32 `json:"updated_at"`
	Decimals     int8   `json:"decimals"`
	Minted       string `json:"minted"`
	Holders      uint64 `json:"holders"`
	TxCnt        uint64 `json:"tx_cnt"`
	Progress     string `json:"progress"`
}

// FindInscriptionTickCmd defines the inscription JSON-RPC command.
type FindInscriptionTickCmd struct {
	Chain    string
	Protocol string
	Tick     string
}

type FindInscriptionTickResponse struct {
	Tick interface{} `json:"tick"`
}

// FindUserTransactionsCmd defines the inscription JSON-RPC command.
type FindUserTransactionsCmd struct {
	Limit    int
	Offset   int
	Address  string
	Chain    string
	Protocol string
	Tick     string
	Event    int8
}

type AddressTransaction struct {
	Chain     string `json:"chain"`
	Protocol  string `json:"protocol"`
	Tick      string `json:"tick"`
	Address   string `json:"address"`
	From      string `json:"from"`
	To        string `json:"to"`
	TxHash    string `json:"tx_hash"`
	Amount    string `json:"amount"`
	Event     int8   `json:"event"`
	Operate   string `json:"operate"`
	Status    int8   `json:"status"`
	CreatedAt uint32 `json:"created_at"`
	UpdatedAt uint32 `json:"updated_at"`
}

type FindUserTransactionsResponse struct {
	Transactions interface{} `json:"transactions"`
	Total        int64       `json:"total"`
	Limit        int         `json:"limit"`
	Offset       int         `json:"offset"`
}

// FindUserBalancesCmd defines the inscription JSON-RPC command.
type FindUserBalancesCmd struct {
	Limit    int
	Offset   int
	Address  string
	Chain    string
	Protocol string
	Tick     string
}

type IndsGetBalanceByAddressCmd struct {
	Limit    int
	Offset   int
	Address  string
	Chain    string
	Protocol string
	Tick     string
	Sort     int
}

type FindUserBalanceCmd struct {
	Address  string
	Chain    string
	Protocol string
	Tick     string
}

type BalanceInfo struct {
	Chain        string `json:"chain"`
	Protocol     string `json:"protocol"`
	Tick         string `json:"tick"`
	Address      string `json:"address"`
	Balance      string `json:"balance"`
	DeployHash   string `json:"deploy_hash"`
	TransferType int8   `json:"transfer_type"`
}

type TickHolder struct {
	Chain       string `json:"chain"`
	Protocol    string `json:"protocol"`
	Tick        string `json:"tick"`
	DeployHash  string `json:"deploy_hash"`
	Address     string `json:"address"`
	Balance     string `json:"balance"`
	TotalSupply string `json:"total_supply"`
}

type BalanceBrief struct {
	Tick         string       `json:"tick"`
	Balance      string       `json:"balance"`
	TransferType int8         `json:"transfer_type"`
	Utxos        []*UTXOBrief `json:"utxos,omitempty"`
	DeployHash   string       `json:"deploy_hash"`
	Available    string       `json:"available"`
}

type UTXOBrief struct {
	Tick     string `json:"tick"`
	Amount   string `json:"amount"`
	RootHash string `json:"root_hash"`
}

type FindUserBalancesResponse struct {
	Inscriptions interface{} `json:"inscriptions"`
	Total        int64       `json:"total"`
	Limit        int         `json:"limit"`
	Offset       int         `json:"offset"`
}

type FindUserBalanceResponse struct {
	Balance interface{} `json:"balance"`
}

type FindTickHoldersCmd struct {
	Limit    int
	Offset   int
	Chain    string
	Protocol string
	Tick     string
}

type IndsGetHoldersByTickCmd struct {
	Limit    int
	Offset   int
	Chain    string
	Protocol string
	Tick     string
	SortMode int
}

type GetTickBriefsCmd struct {
	Addresses []*TickAddress `json:"addresses"`
}

type TickAddress struct {
	Chain      string `json:"chain"`
	DeployHash string `json:"deploy_hash"`
}

type GetTickBriefsResp struct {
	Inscriptions []*model.InscriptionOverView `json:"inscriptions"`
}

type FindTickHoldersResponse struct {
	Holders interface{} `json:"holders"`
	Total   int64       `json:"total"`
	Limit   int         `json:"limit"`
	Offset  int         `json:"offset"`
}

type BlockInfo struct {
	Chain       string `json:"chain"`
	BlockNumber string `json:"block_number"`
	BlockTime   string `json:"block_time"`
	TimeStamp   uint32 `json:"timestamp"`
}

type LastBlockNumberCmd struct {
	Chains []string
}

type TxOperateCmd struct {
	Chain     string
	InputData string
}

type TxOperateResponse struct {
	Operate    string `json:"operate"`
	Protocol   string `json:"protocol"`
	Tick       string `json:"tick"`
	DeployHash string `json:"deploy_hash"`
}

type GetTxByHashCmd struct {
	Chain  string
	TxHash string
}

type TransactionInfo struct {
	Protocol   string `json:"protocol"`
	Tick       string `json:"tick"`
	DeployHash string `json:"deploy_hash"`
	From       string `json:"from"`
	To         string `json:"to"`
	Amount     string `json:"amount"`
	Op         string `json:"op"`
}

type GetTxByHashResponse struct {
	IsInscription bool             `json:"is_inscription"`
	Transaction   *TransactionInfo `json:"transaction,omitempty"`
}

func init() {
	// No special flags for commands in this file.
	flags := UsageFlag(0)

	MustRegisterCmd("inscription.All", (*FindAllInscriptionsCmd)(nil), flags)
	MustRegisterCmd("inscription.Tick", (*FindInscriptionTickCmd)(nil), flags)
	MustRegisterCmd("address.Transactions", (*FindUserTransactionsCmd)(nil), flags)
	MustRegisterCmd("address.Balances", (*FindUserBalancesCmd)(nil), flags)
	MustRegisterCmd("address.Balance", (*FindUserBalanceCmd)(nil), flags)
	MustRegisterCmd("tick.Holders", (*FindTickHoldersCmd)(nil), flags)
	MustRegisterCmd("block.LastNumber", (*LastBlockNumberCmd)(nil), flags)
	MustRegisterCmd("tool.InscriptionTxOperate", (*TxOperateCmd)(nil), flags)
	MustRegisterCmd("transaction.Info", (*GetTxByHashCmd)(nil), flags)
	MustRegisterCmd("tick.GetBriefs", (*GetTickBriefsCmd)(nil), flags)

	//v2
	MustRegisterCmd("inds_getTicks", (*IndsGetTicksCmd)(nil), flags)
	MustRegisterCmd("inds_getTick", (*IndsGetTickCmd)(nil), flags)
	MustRegisterCmd("inds_getTransactionByAddress", (*FindUserTransactionsCmd)(nil), flags)
	MustRegisterCmd("inds_getBalanceByAddress", (*IndsGetBalanceByAddressCmd)(nil), flags)
	MustRegisterCmd("inds_getHoldersByTick", (*IndsGetHoldersByTickCmd)(nil), flags)
	MustRegisterCmd("inds_getLastBlockNumberIndexed", (*LastBlockNumberCmd)(nil), flags)
	MustRegisterCmd("inds_getTickByCallData", (*TxOperateCmd)(nil), flags)
	MustRegisterCmd("inds_getTransactionByHash", (*GetTxByHashCmd)(nil), flags)
}
