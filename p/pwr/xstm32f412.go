// DO NOT EDIT THIS FILE. GENERATED BY xgen.

// +build stm32f412

package pwr

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	CR  RCR
	CSR RCSR
}

func PWR() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.PWR_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

func (p *Periph) Bus() bus.Bus {
	return bus.APB1
}

type CR uint32

type RCR struct{ mmio.U32 }

func (r *RCR) LoadBits(mask CR) CR  { return CR(r.U32.LoadBits(uint32(mask))) }
func (r *RCR) StoreBits(mask, b CR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR) SetBits(mask CR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR) ClearBits(mask CR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR) Load() CR             { return CR(r.U32.Load()) }
func (r *RCR) Store(b CR)           { r.U32.Store(uint32(b)) }

type RMCR struct{ mmio.UM32 }

func (rm RMCR) Load() CR   { return CR(rm.UM32.Load()) }
func (rm RMCR) Store(b CR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) LPDS() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(LPDS)}}
}

func (p *Periph) PDDS() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PDDS)}}
}

func (p *Periph) CWUF() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(CWUF)}}
}

func (p *Periph) CSBF() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(CSBF)}}
}

func (p *Periph) PVDE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PVDE)}}
}

func (p *Periph) PLS() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PLS)}}
}

func (p *Periph) DBP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBP)}}
}

func (p *Periph) FPDS() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(FPDS)}}
}

func (p *Periph) ADCDC1() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ADCDC1)}}
}

func (p *Periph) VOS() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(VOS)}}
}

type CSR uint32

type RCSR struct{ mmio.U32 }

func (r *RCSR) LoadBits(mask CSR) CSR { return CSR(r.U32.LoadBits(uint32(mask))) }
func (r *RCSR) StoreBits(mask, b CSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCSR) SetBits(mask CSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCSR) ClearBits(mask CSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCSR) Load() CSR             { return CSR(r.U32.Load()) }
func (r *RCSR) Store(b CSR)           { r.U32.Store(uint32(b)) }

type RMCSR struct{ mmio.UM32 }

func (rm RMCSR) Load() CSR   { return CSR(rm.UM32.Load()) }
func (rm RMCSR) Store(b CSR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) WUF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(WUF)}}
}

func (p *Periph) SBF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(SBF)}}
}

func (p *Periph) PVDO() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(PVDO)}}
}

func (p *Periph) BRR() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(BRR)}}
}

func (p *Periph) EWUP() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(EWUP)}}
}

func (p *Periph) BRE() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(BRE)}}
}

func (p *Periph) VOSRDY() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(VOSRDY)}}
}
