package dcache

import (
	"fmt"
	"github.com/shopspring/decimal"
	"open-indexer/xylog"
	"strings"
	"sync"
)

// InscriptionStats
/*****************************************************
 * Build cache for all inscriptions stats data
 * Mainly used for statics data query
 ****************************************************/
type InscriptionStats struct {
	sid   uint32
	ticks *sync.Map
}

type InsStats struct {
	SID     uint32
	Minted  decimal.Decimal
	Holders int64
	TxCnt   uint64
}

func NewInscriptionStats() *InscriptionStats {
	return &InscriptionStats{
		ticks: &sync.Map{},
	}
}

/***************************************
 * idx define protocol tick unique id
 ***************************************/
func (d *InscriptionStats) idx(protocol, tick string) string {
	return fmt.Sprintf("%s_%s", strings.ToLower(protocol), strings.ToLower(tick))
}

// Update
/***************************************
 * update ticks
 ***************************************/
func (d *InscriptionStats) Update(protocol, tick string, stats *InsStats) *InsStats {
	ok, insStats := d.Get(protocol, tick)
	if !ok {
		return nil
	}

	if stats.Minted.GreaterThan(decimal.Zero) {
		insStats.Minted = stats.Minted
	}

	if stats.Holders > 0 {
		insStats.Holders = stats.Holders
	}

	if stats.TxCnt > 0 {
		insStats.TxCnt = stats.TxCnt
	}
	return insStats
}

// Create
/***************************************
 * init tick's id
 ***************************************/
func (d *InscriptionStats) Create(protocol, tick string, stats *InsStats) *InsStats {
	// Add auto_increment ID
	if stats.SID <= 0 {
		d.sid++
		stats.SID = d.sid
	}

	idx := d.idx(protocol, tick)
	d.ticks.Store(idx, stats)
	return stats
}

func (d *InscriptionStats) Mint(protocol, tick string, amount decimal.Decimal) *InsStats {
	ok, insStats := d.Get(protocol, tick)
	if !ok {
		return nil
	}

	if amount.LessThanOrEqual(decimal.Zero) {
		return insStats
	}

	insStats.Minted = insStats.Minted.Add(amount)
	return insStats
}

func (d *InscriptionStats) Holders(protocol, tick string, incr int64) *InsStats {
	ok, insStats := d.Get(protocol, tick)
	if !ok {
		return nil
	}

	insStats.Holders = insStats.Holders + incr

	if insStats.Holders < 0 {
		xylog.Logger.Fatalf("protocol:%s, tick:%s holders < 0", protocol, tick)
	}
	return insStats
}

func (d *InscriptionStats) TxCnt(protocol, tick string, incr uint64) *InsStats {
	ok, insStats := d.Get(protocol, tick)
	if !ok {
		return nil
	}

	insStats.TxCnt = insStats.TxCnt + incr
	return insStats
}

// SetSid set auto_increment id
func (d *InscriptionStats) SetSid(sid uint32) {
	if sid > d.sid {
		d.sid = sid
	}
}

// Get
/***************************************
 * get tick meta data contains filed (id, transfer_type)
 ***************************************/
func (d *InscriptionStats) Get(protocol, tick string) (bool, *InsStats) {
	idx := d.idx(protocol, tick)
	t, ok := d.ticks.Load(idx)
	if !ok {
		return false, nil
	}
	return true, t.(*InsStats)
}
