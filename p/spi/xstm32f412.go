// DO NOT EDIT THIS FILE. GENERATED BY xgen.

// +build stm32f412

package spi

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	CR1     RCR1
	CR2     RCR2
	SR      RSR
	DR      RDR
	CRCPR   RCRCPR
	RXCRCR  RRXCRCR
	TXCRCR  RTXCRCR
	I2SCFGR RI2SCFGR
	I2SPR   RI2SPR
}

func I2S2ext() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.I2S2ext_BASE))) }
func I2S3ext() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.I2S3ext_BASE))) }
func SPI1() *Periph    { return (*Periph)(unsafe.Pointer(uintptr(mmap.SPI1_BASE))) }
func SPI2() *Periph    { return (*Periph)(unsafe.Pointer(uintptr(mmap.SPI2_BASE))) }
func SPI3() *Periph    { return (*Periph)(unsafe.Pointer(uintptr(mmap.SPI3_BASE))) }
func SPI4() *Periph    { return (*Periph)(unsafe.Pointer(uintptr(mmap.SPI4_BASE))) }
func SPI5() *Periph    { return (*Periph)(unsafe.Pointer(uintptr(mmap.SPI5_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

func (p *Periph) Bus() bus.Bus {
	switch p.BaseAddr() {
	default:
		return bus.APB2
	case mmap.SPI2_BASE, mmap.SPI3_BASE:
		return bus.APB1
	}
}

type CR1 uint32

type RCR1 struct{ mmio.U32 }

func (r *RCR1) LoadBits(mask CR1) CR1 { return CR1(r.U32.LoadBits(uint32(mask))) }
func (r *RCR1) StoreBits(mask, b CR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR1) SetBits(mask CR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR1) ClearBits(mask CR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR1) Load() CR1             { return CR1(r.U32.Load()) }
func (r *RCR1) Store(b CR1)           { r.U32.Store(uint32(b)) }

type RMCR1 struct{ mmio.UM32 }

func (rm RMCR1) Load() CR1   { return CR1(rm.UM32.Load()) }
func (rm RMCR1) Store(b CR1) { rm.UM32.Store(uint32(b)) }

func (p *Periph) CPHA() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(CPHA)}}
}

func (p *Periph) CPOL() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(CPOL)}}
}

func (p *Periph) MSTR() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(MSTR)}}
}

func (p *Periph) BR() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(BR)}}
}

func (p *Periph) SPE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(SPE)}}
}

func (p *Periph) LSBFIRST() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(LSBFIRST)}}
}

func (p *Periph) SSI() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(SSI)}}
}

func (p *Periph) SSM() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(SSM)}}
}

func (p *Periph) RXONLY() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(RXONLY)}}
}

func (p *Periph) DFF() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(DFF)}}
}

func (p *Periph) CRCNEXT() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(CRCNEXT)}}
}

func (p *Periph) CRCEN() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(CRCEN)}}
}

func (p *Periph) BIDIOE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(BIDIOE)}}
}

func (p *Periph) BIDIMODE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(BIDIMODE)}}
}

type CR2 uint32

type RCR2 struct{ mmio.U32 }

func (r *RCR2) LoadBits(mask CR2) CR2 { return CR2(r.U32.LoadBits(uint32(mask))) }
func (r *RCR2) StoreBits(mask, b CR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR2) SetBits(mask CR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR2) ClearBits(mask CR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR2) Load() CR2             { return CR2(r.U32.Load()) }
func (r *RCR2) Store(b CR2)           { r.U32.Store(uint32(b)) }

type RMCR2 struct{ mmio.UM32 }

func (rm RMCR2) Load() CR2   { return CR2(rm.UM32.Load()) }
func (rm RMCR2) Store(b CR2) { rm.UM32.Store(uint32(b)) }

func (p *Periph) RXDMAEN() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(RXDMAEN)}}
}

func (p *Periph) TXDMAEN() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(TXDMAEN)}}
}

func (p *Periph) SSOE() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(SSOE)}}
}

func (p *Periph) FRF() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(FRF)}}
}

func (p *Periph) ERRIE() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(ERRIE)}}
}

func (p *Periph) RXNEIE() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(RXNEIE)}}
}

func (p *Periph) TXEIE() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(TXEIE)}}
}

type SR uint32

type RSR struct{ mmio.U32 }

func (r *RSR) LoadBits(mask SR) SR  { return SR(r.U32.LoadBits(uint32(mask))) }
func (r *RSR) StoreBits(mask, b SR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSR) SetBits(mask SR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSR) ClearBits(mask SR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSR) Load() SR             { return SR(r.U32.Load()) }
func (r *RSR) Store(b SR)           { r.U32.Store(uint32(b)) }

type RMSR struct{ mmio.UM32 }

func (rm RMSR) Load() SR   { return SR(rm.UM32.Load()) }
func (rm RMSR) Store(b SR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) RXNE() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(RXNE)}}
}

func (p *Periph) TXE() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(TXE)}}
}

func (p *Periph) CHSIDE() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(CHSIDE)}}
}

func (p *Periph) UDR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(UDR)}}
}

func (p *Periph) CRCERR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(CRCERR)}}
}

func (p *Periph) MODF() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(MODF)}}
}

func (p *Periph) OVR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(OVR)}}
}

func (p *Periph) BSY() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(BSY)}}
}

func (p *Periph) TIFRFE() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(TIFRFE)}}
}

type DR uint32

type RDR struct{ mmio.U32 }

func (r *RDR) LoadBits(mask DR) DR  { return DR(r.U32.LoadBits(uint32(mask))) }
func (r *RDR) StoreBits(mask, b DR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDR) SetBits(mask DR)      { r.U32.SetBits(uint32(mask)) }
func (r *RDR) ClearBits(mask DR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDR) Load() DR             { return DR(r.U32.Load()) }
func (r *RDR) Store(b DR)           { r.U32.Store(uint32(b)) }

type RMDR struct{ mmio.UM32 }

func (rm RMDR) Load() DR   { return DR(rm.UM32.Load()) }
func (rm RMDR) Store(b DR) { rm.UM32.Store(uint32(b)) }

type CRCPR uint32

type RCRCPR struct{ mmio.U32 }

func (r *RCRCPR) LoadBits(mask CRCPR) CRCPR { return CRCPR(r.U32.LoadBits(uint32(mask))) }
func (r *RCRCPR) StoreBits(mask, b CRCPR)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCRCPR) SetBits(mask CRCPR)        { r.U32.SetBits(uint32(mask)) }
func (r *RCRCPR) ClearBits(mask CRCPR)      { r.U32.ClearBits(uint32(mask)) }
func (r *RCRCPR) Load() CRCPR               { return CRCPR(r.U32.Load()) }
func (r *RCRCPR) Store(b CRCPR)             { r.U32.Store(uint32(b)) }

type RMCRCPR struct{ mmio.UM32 }

func (rm RMCRCPR) Load() CRCPR   { return CRCPR(rm.UM32.Load()) }
func (rm RMCRCPR) Store(b CRCPR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) CRCPOLY() RMCRCPR {
	return RMCRCPR{mmio.UM32{&p.CRCPR.U32, uint32(CRCPOLY)}}
}

type RXCRCR uint32

type RRXCRCR struct{ mmio.U32 }

func (r *RRXCRCR) LoadBits(mask RXCRCR) RXCRCR { return RXCRCR(r.U32.LoadBits(uint32(mask))) }
func (r *RRXCRCR) StoreBits(mask, b RXCRCR)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRXCRCR) SetBits(mask RXCRCR)         { r.U32.SetBits(uint32(mask)) }
func (r *RRXCRCR) ClearBits(mask RXCRCR)       { r.U32.ClearBits(uint32(mask)) }
func (r *RRXCRCR) Load() RXCRCR                { return RXCRCR(r.U32.Load()) }
func (r *RRXCRCR) Store(b RXCRCR)              { r.U32.Store(uint32(b)) }

type RMRXCRCR struct{ mmio.UM32 }

func (rm RMRXCRCR) Load() RXCRCR   { return RXCRCR(rm.UM32.Load()) }
func (rm RMRXCRCR) Store(b RXCRCR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) RxCRC() RMRXCRCR {
	return RMRXCRCR{mmio.UM32{&p.RXCRCR.U32, uint32(RxCRC)}}
}

type TXCRCR uint32

type RTXCRCR struct{ mmio.U32 }

func (r *RTXCRCR) LoadBits(mask TXCRCR) TXCRCR { return TXCRCR(r.U32.LoadBits(uint32(mask))) }
func (r *RTXCRCR) StoreBits(mask, b TXCRCR)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTXCRCR) SetBits(mask TXCRCR)         { r.U32.SetBits(uint32(mask)) }
func (r *RTXCRCR) ClearBits(mask TXCRCR)       { r.U32.ClearBits(uint32(mask)) }
func (r *RTXCRCR) Load() TXCRCR                { return TXCRCR(r.U32.Load()) }
func (r *RTXCRCR) Store(b TXCRCR)              { r.U32.Store(uint32(b)) }

type RMTXCRCR struct{ mmio.UM32 }

func (rm RMTXCRCR) Load() TXCRCR   { return TXCRCR(rm.UM32.Load()) }
func (rm RMTXCRCR) Store(b TXCRCR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) TxCRC() RMTXCRCR {
	return RMTXCRCR{mmio.UM32{&p.TXCRCR.U32, uint32(TxCRC)}}
}

type I2SCFGR uint32

type RI2SCFGR struct{ mmio.U32 }

func (r *RI2SCFGR) LoadBits(mask I2SCFGR) I2SCFGR { return I2SCFGR(r.U32.LoadBits(uint32(mask))) }
func (r *RI2SCFGR) StoreBits(mask, b I2SCFGR)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RI2SCFGR) SetBits(mask I2SCFGR)          { r.U32.SetBits(uint32(mask)) }
func (r *RI2SCFGR) ClearBits(mask I2SCFGR)        { r.U32.ClearBits(uint32(mask)) }
func (r *RI2SCFGR) Load() I2SCFGR                 { return I2SCFGR(r.U32.Load()) }
func (r *RI2SCFGR) Store(b I2SCFGR)               { r.U32.Store(uint32(b)) }

type RMI2SCFGR struct{ mmio.UM32 }

func (rm RMI2SCFGR) Load() I2SCFGR   { return I2SCFGR(rm.UM32.Load()) }
func (rm RMI2SCFGR) Store(b I2SCFGR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) CHLEN() RMI2SCFGR {
	return RMI2SCFGR{mmio.UM32{&p.I2SCFGR.U32, uint32(CHLEN)}}
}

func (p *Periph) DATLEN() RMI2SCFGR {
	return RMI2SCFGR{mmio.UM32{&p.I2SCFGR.U32, uint32(DATLEN)}}
}

func (p *Periph) CKPOL() RMI2SCFGR {
	return RMI2SCFGR{mmio.UM32{&p.I2SCFGR.U32, uint32(CKPOL)}}
}

func (p *Periph) I2SSTD() RMI2SCFGR {
	return RMI2SCFGR{mmio.UM32{&p.I2SCFGR.U32, uint32(I2SSTD)}}
}

func (p *Periph) PCMSYNC() RMI2SCFGR {
	return RMI2SCFGR{mmio.UM32{&p.I2SCFGR.U32, uint32(PCMSYNC)}}
}

func (p *Periph) I2SCFG() RMI2SCFGR {
	return RMI2SCFGR{mmio.UM32{&p.I2SCFGR.U32, uint32(I2SCFG)}}
}

func (p *Periph) I2SE() RMI2SCFGR {
	return RMI2SCFGR{mmio.UM32{&p.I2SCFGR.U32, uint32(I2SE)}}
}

func (p *Periph) I2SMOD() RMI2SCFGR {
	return RMI2SCFGR{mmio.UM32{&p.I2SCFGR.U32, uint32(I2SMOD)}}
}

type I2SPR uint32

type RI2SPR struct{ mmio.U32 }

func (r *RI2SPR) LoadBits(mask I2SPR) I2SPR { return I2SPR(r.U32.LoadBits(uint32(mask))) }
func (r *RI2SPR) StoreBits(mask, b I2SPR)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RI2SPR) SetBits(mask I2SPR)        { r.U32.SetBits(uint32(mask)) }
func (r *RI2SPR) ClearBits(mask I2SPR)      { r.U32.ClearBits(uint32(mask)) }
func (r *RI2SPR) Load() I2SPR               { return I2SPR(r.U32.Load()) }
func (r *RI2SPR) Store(b I2SPR)             { r.U32.Store(uint32(b)) }

type RMI2SPR struct{ mmio.UM32 }

func (rm RMI2SPR) Load() I2SPR   { return I2SPR(rm.UM32.Load()) }
func (rm RMI2SPR) Store(b I2SPR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) I2SDIV() RMI2SPR {
	return RMI2SPR{mmio.UM32{&p.I2SPR.U32, uint32(I2SDIV)}}
}

func (p *Periph) ODD() RMI2SPR {
	return RMI2SPR{mmio.UM32{&p.I2SPR.U32, uint32(ODD)}}
}

func (p *Periph) MCKOE() RMI2SPR {
	return RMI2SPR{mmio.UM32{&p.I2SPR.U32, uint32(MCKOE)}}
}
