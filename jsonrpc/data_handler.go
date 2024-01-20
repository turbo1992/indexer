package jsonrpc

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/uxuycom/indexer/model"
	"strings"
)

func findAddressBalances(s *RpcServer, limit, offset int, address, chain, protocol, tick string, sort int) (interface{}, error) {
	protocol = strings.ToLower(protocol)
	tick = strings.ToLower(tick)
	cacheKey := fmt.Sprintf("addr_balances_%d_%d_%s_%s_%s_%s_%d", limit, offset, address, chain, protocol, tick, sort)
	if ins, ok := s.cacheStore.Get(cacheKey); ok {
		if allIns, ok := ins.(*FindUserBalancesResponse); ok {
			return allIns, nil
		}
	}

	balances, total, err := s.dbc.GetAddressInscriptions(limit, offset, address, chain, protocol, tick, sort)
	if err != nil {
		return ErrRPCInternal, err
	}

	list := make([]*BalanceInfo, 0, len(balances))
	for _, b := range balances {
		balance := &BalanceInfo{
			Chain:        b.Chain,
			Protocol:     b.Protocol,
			Tick:         b.Tick,
			Address:      b.Address,
			Balance:      b.Balance.String(),
			DeployHash:   b.DeployHash,
			TransferType: b.TransferType,
		}
		list = append(list, balance)
	}

	resp := &FindUserBalancesResponse{
		Inscriptions: list,
		Total:        total,
		Limit:        limit,
		Offset:       offset,
	}
	s.cacheStore.Set(cacheKey, resp)
	return resp, nil
}

func findInsciptions(s *RpcServer, limit, offset int, chain, protocol, tick, deployBy string, sort, sortMode int) (interface{}, error) {
	protocol = strings.ToLower(protocol)
	tick = strings.ToLower(tick)
	cacheKey := fmt.Sprintf("all_ins_%d_%d_%s_%s_%s_%s_%d_%d", limit, offset, chain, protocol, tick, deployBy, sort, sortMode)
	if ins, ok := s.cacheStore.Get(cacheKey); ok {
		if allIns, ok := ins.(*FindAllInscriptionsResponse); ok {
			return allIns, nil
		}
	}
	inscriptions, total, err := s.dbc.GetInscriptions(limit, offset, chain, protocol, tick, deployBy, sort, sortMode)
	if err != nil {
		return ErrRPCInternal, err
	}

	result := make([]*model.InscriptionBrief, 0, len(inscriptions))

	for _, ins := range inscriptions {
		brief := &model.InscriptionBrief{
			Chain:        ins.Chain,
			Protocol:     ins.Protocol,
			Tick:         ins.Name,
			DeployBy:     ins.DeployBy,
			DeployHash:   ins.DeployHash,
			TotalSupply:  ins.TotalSupply.String(),
			Holders:      ins.Holders,
			Minted:       ins.Minted.String(),
			LimitPerMint: ins.LimitPerMint.String(),
			TransferType: ins.TransferType,
			Status:       model.MintStatusProcessing,
			TxCnt:        ins.TxCnt,
			CreatedAt:    uint32(ins.CreatedAt.Unix()),
		}

		minted := ins.Minted
		totalSupply := ins.TotalSupply

		if totalSupply != decimal.Zero && minted != decimal.Zero {
			percentage, _ := minted.Div(totalSupply).Float64()
			if percentage >= 1 {
				percentage = 1
			}
			brief.MintedPercent = fmt.Sprintf("%.4f", percentage)
		}

		if ins.Minted.Cmp(ins.TotalSupply) >= 0 {
			brief.Status = model.MintStatusAllMinted
		}

		result = append(result, brief)
	}

	resp := &FindAllInscriptionsResponse{
		Inscriptions: result,
		Total:        total,
		Limit:        limit,
		Offset:       offset,
	}
	s.cacheStore.Set(cacheKey, resp)
	return resp, nil
}

func findInsciption(s *RpcServer, chain, protocol, tick, deployHash string) (interface{}, error) {
	protocol = strings.ToLower(protocol)
	tick = strings.ToLower(tick)

	cacheKey := fmt.Sprintf("tick_%s_%s_%s_%s", chain, protocol, tick, deployHash)
	if ins, ok := s.cacheStore.Get(cacheKey); ok {
		if ticks, ok := ins.(*InscriptionInfo); ok {
			return ticks, nil
		}
	}

	inscription, err := s.dbc.FindInscriptionInfo(chain, protocol, tick, deployHash)
	if err != nil {
		return ErrRPCInternal, err
	}
	if inscription == nil {
		return ErrRPCRecordNotFound, err
	}

	resp := &InscriptionInfo{
		Chain:        inscription.Chain,
		Protocol:     inscription.Protocol,
		Tick:         inscription.Tick,
		Name:         inscription.Name,
		LimitPerMint: inscription.LimitPerMint.String(),
		DeployBy:     inscription.DeployBy,
		TotalSupply:  inscription.TotalSupply.String(),
		DeployHash:   inscription.DeployHash,
		TransferType: inscription.TransferType,
		Decimals:     inscription.Decimals,
		Minted:       inscription.Minted.String(),
		Holders:      inscription.Holders,
		TxCnt:        inscription.TxCnt,
		Progress:     inscription.Progress.String(),
		DeployTime:   uint32(inscription.DeployTime.Unix()),
		CreatedAt:    uint32(inscription.CreatedAt.Unix()),
		UpdatedAt:    uint32(inscription.UpdatedAt.Unix()),
	}

	s.cacheStore.Set(cacheKey, resp)
	return resp, nil
}

func findTickHolders(s *RpcServer, limit int, offset int, chain, protocol, tick string, sortMode int) (interface{}, error) {
	protocol = strings.ToLower(protocol)
	tick = strings.ToLower(tick)
	cacheKey := fmt.Sprintf("all_ins_%d_%d_%s_%s_%s_%d", limit, offset, chain, protocol, tick, sortMode)
	if ins, ok := s.cacheStore.Get(cacheKey); ok {
		if allIns, ok := ins.(*FindTickHoldersResponse); ok {
			return allIns, nil
		}
	}

	// get inscription info
	inscription, err := s.dbc.FindInscriptionByTick(chain, protocol, tick)
	if err != nil {
		return ErrRPCInternal, err
	}
	if inscription == nil {
		return nil, errors.New("Record not found")
	}

	// get holders
	holders, total, err := s.dbc.GetHoldersByTick(limit, offset, chain, protocol, tick, sortMode)
	if err != nil {
		return ErrRPCInternal, err
	}

	list := make([]*TickHolder, 0, len(holders))
	for _, holder := range holders {
		balance := &TickHolder{
			Chain:       holder.Chain,
			Protocol:    holder.Protocol,
			Tick:        holder.Tick,
			DeployHash:  inscription.DeployHash,
			Address:     holder.Address,
			Balance:     holder.Balance.String(),
			TotalSupply: inscription.TotalSupply.String(),
		}
		list = append(list, balance)
	}

	resp := &FindTickHoldersResponse{
		Holders: list,
		Total:   total,
		Limit:   limit,
		Offset:  offset,
	}

	s.cacheStore.Set(cacheKey, resp)
	return resp, nil
}
