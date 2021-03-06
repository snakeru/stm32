// DO NOT EDIT THIS FILE. GENERATED BY xgen.

// +build stm32g471xx

package pwr

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	CR1  RCR1
	CR2  RCR2
	CR3  RCR3
	CR4  RCR4
	SR1  RSR1
	SR2  RSR2
	SCR  RSCR
	_    uint32
	PUDC [7]RPUDC
	_    [10]uint32
	CR5  RCR5
}

func PWR() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.PWR_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

func (p *Periph) Bus() bus.Bus {
	return bus.APB1
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

func (p *Periph) LPMS() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(LPMS)}}
}

func (p *Periph) DBP() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(DBP)}}
}

func (p *Periph) VOS() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(VOS)}}
}

func (p *Periph) LPR() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(LPR)}}
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

func (p *Periph) PVDE() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(PVDE)}}
}

func (p *Periph) PLS() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(PLS)}}
}

func (p *Periph) PVMEN1() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(PVMEN1)}}
}

func (p *Periph) PVMEN2() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(PVMEN2)}}
}

func (p *Periph) PVMEN3() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(PVMEN3)}}
}

func (p *Periph) PVMEN4() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(PVMEN4)}}
}

type CR3 uint32

type RCR3 struct{ mmio.U32 }

func (r *RCR3) LoadBits(mask CR3) CR3 { return CR3(r.U32.LoadBits(uint32(mask))) }
func (r *RCR3) StoreBits(mask, b CR3) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR3) SetBits(mask CR3)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR3) ClearBits(mask CR3)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR3) Load() CR3             { return CR3(r.U32.Load()) }
func (r *RCR3) Store(b CR3)           { r.U32.Store(uint32(b)) }

type RMCR3 struct{ mmio.UM32 }

func (rm RMCR3) Load() CR3   { return CR3(rm.UM32.Load()) }
func (rm RMCR3) Store(b CR3) { rm.UM32.Store(uint32(b)) }

func (p *Periph) EWUP1() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(EWUP1)}}
}

func (p *Periph) EWUP2() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(EWUP2)}}
}

func (p *Periph) EWUP3() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(EWUP3)}}
}

func (p *Periph) EWUP4() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(EWUP4)}}
}

func (p *Periph) EWUP5() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(EWUP5)}}
}

func (p *Periph) RRS() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(RRS)}}
}

func (p *Periph) APC() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(APC)}}
}

func (p *Periph) UCPD1_STDBY() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(UCPD1_STDBY)}}
}

func (p *Periph) UCPD1_DBDIS() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(UCPD1_DBDIS)}}
}

func (p *Periph) EIWUL() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(EIWUL)}}
}

type CR4 uint32

type RCR4 struct{ mmio.U32 }

func (r *RCR4) LoadBits(mask CR4) CR4 { return CR4(r.U32.LoadBits(uint32(mask))) }
func (r *RCR4) StoreBits(mask, b CR4) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR4) SetBits(mask CR4)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR4) ClearBits(mask CR4)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR4) Load() CR4             { return CR4(r.U32.Load()) }
func (r *RCR4) Store(b CR4)           { r.U32.Store(uint32(b)) }

type RMCR4 struct{ mmio.UM32 }

func (rm RMCR4) Load() CR4   { return CR4(rm.UM32.Load()) }
func (rm RMCR4) Store(b CR4) { rm.UM32.Store(uint32(b)) }

func (p *Periph) WP1() RMCR4 {
	return RMCR4{mmio.UM32{&p.CR4.U32, uint32(WP1)}}
}

func (p *Periph) WP2() RMCR4 {
	return RMCR4{mmio.UM32{&p.CR4.U32, uint32(WP2)}}
}

func (p *Periph) WP3() RMCR4 {
	return RMCR4{mmio.UM32{&p.CR4.U32, uint32(WP3)}}
}

func (p *Periph) WP4() RMCR4 {
	return RMCR4{mmio.UM32{&p.CR4.U32, uint32(WP4)}}
}

func (p *Periph) WP5() RMCR4 {
	return RMCR4{mmio.UM32{&p.CR4.U32, uint32(WP5)}}
}

func (p *Periph) VBE() RMCR4 {
	return RMCR4{mmio.UM32{&p.CR4.U32, uint32(VBE)}}
}

func (p *Periph) VBRS() RMCR4 {
	return RMCR4{mmio.UM32{&p.CR4.U32, uint32(VBRS)}}
}

type SR1 uint32

type RSR1 struct{ mmio.U32 }

func (r *RSR1) LoadBits(mask SR1) SR1 { return SR1(r.U32.LoadBits(uint32(mask))) }
func (r *RSR1) StoreBits(mask, b SR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSR1) SetBits(mask SR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RSR1) ClearBits(mask SR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSR1) Load() SR1             { return SR1(r.U32.Load()) }
func (r *RSR1) Store(b SR1)           { r.U32.Store(uint32(b)) }

type RMSR1 struct{ mmio.UM32 }

func (rm RMSR1) Load() SR1   { return SR1(rm.UM32.Load()) }
func (rm RMSR1) Store(b SR1) { rm.UM32.Store(uint32(b)) }

func (p *Periph) WUF1() RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(WUF1)}}
}

func (p *Periph) WUF2() RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(WUF2)}}
}

func (p *Periph) WUF3() RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(WUF3)}}
}

func (p *Periph) WUF4() RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(WUF4)}}
}

func (p *Periph) WUF5() RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(WUF5)}}
}

func (p *Periph) SBF() RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(SBF)}}
}

func (p *Periph) WUFI() RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(WUFI)}}
}

type SR2 uint32

type RSR2 struct{ mmio.U32 }

func (r *RSR2) LoadBits(mask SR2) SR2 { return SR2(r.U32.LoadBits(uint32(mask))) }
func (r *RSR2) StoreBits(mask, b SR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSR2) SetBits(mask SR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RSR2) ClearBits(mask SR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSR2) Load() SR2             { return SR2(r.U32.Load()) }
func (r *RSR2) Store(b SR2)           { r.U32.Store(uint32(b)) }

type RMSR2 struct{ mmio.UM32 }

func (rm RMSR2) Load() SR2   { return SR2(rm.UM32.Load()) }
func (rm RMSR2) Store(b SR2) { rm.UM32.Store(uint32(b)) }

func (p *Periph) REGLPS() RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(REGLPS)}}
}

func (p *Periph) REGLPF() RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(REGLPF)}}
}

func (p *Periph) VOSF() RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(VOSF)}}
}

func (p *Periph) PVDO() RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(PVDO)}}
}

func (p *Periph) PVMO1() RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(PVMO1)}}
}

func (p *Periph) PVMO2() RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(PVMO2)}}
}

func (p *Periph) PVMO3() RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(PVMO3)}}
}

func (p *Periph) PVMO4() RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(PVMO4)}}
}

type SCR uint32

type RSCR struct{ mmio.U32 }

func (r *RSCR) LoadBits(mask SCR) SCR { return SCR(r.U32.LoadBits(uint32(mask))) }
func (r *RSCR) StoreBits(mask, b SCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSCR) SetBits(mask SCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSCR) ClearBits(mask SCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSCR) Load() SCR             { return SCR(r.U32.Load()) }
func (r *RSCR) Store(b SCR)           { r.U32.Store(uint32(b)) }

type RMSCR struct{ mmio.UM32 }

func (rm RMSCR) Load() SCR   { return SCR(rm.UM32.Load()) }
func (rm RMSCR) Store(b SCR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) CWUF1() RMSCR {
	return RMSCR{mmio.UM32{&p.SCR.U32, uint32(CWUF1)}}
}

func (p *Periph) CWUF2() RMSCR {
	return RMSCR{mmio.UM32{&p.SCR.U32, uint32(CWUF2)}}
}

func (p *Periph) CWUF3() RMSCR {
	return RMSCR{mmio.UM32{&p.SCR.U32, uint32(CWUF3)}}
}

func (p *Periph) CWUF4() RMSCR {
	return RMSCR{mmio.UM32{&p.SCR.U32, uint32(CWUF4)}}
}

func (p *Periph) CWUF5() RMSCR {
	return RMSCR{mmio.UM32{&p.SCR.U32, uint32(CWUF5)}}
}

func (p *Periph) CSBF() RMSCR {
	return RMSCR{mmio.UM32{&p.SCR.U32, uint32(CSBF)}}
}

type RPUDC struct {
	PUCR RPUCR
	PDCR RPDCR
}

type PUCR uint32

type RPUCR struct{ mmio.U32 }

func (r *RPUCR) LoadBits(mask PUCR) PUCR { return PUCR(r.U32.LoadBits(uint32(mask))) }
func (r *RPUCR) StoreBits(mask, b PUCR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPUCR) SetBits(mask PUCR)       { r.U32.SetBits(uint32(mask)) }
func (r *RPUCR) ClearBits(mask PUCR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RPUCR) Load() PUCR              { return PUCR(r.U32.Load()) }
func (r *RPUCR) Store(b PUCR)            { r.U32.Store(uint32(b)) }

type RMPUCR struct{ mmio.UM32 }

func (rm RMPUCR) Load() PUCR   { return PUCR(rm.UM32.Load()) }
func (rm RMPUCR) Store(b PUCR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) PU0(n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU0)}}
}

func (p *Periph) PU1(n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU1)}}
}

func (p *Periph) PU2(n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU2)}}
}

func (p *Periph) PU3(n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU3)}}
}

func (p *Periph) PU4(n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU4)}}
}

func (p *Periph) PU5(n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU5)}}
}

func (p *Periph) PU6(n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU6)}}
}

func (p *Periph) PU7(n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU7)}}
}

func (p *Periph) PU8(n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU8)}}
}

func (p *Periph) PU9(n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU9)}}
}

func (p *Periph) PU10(n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU10)}}
}

func (p *Periph) PU11(n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU11)}}
}

func (p *Periph) PU12(n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU12)}}
}

func (p *Periph) PU13(n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU13)}}
}

func (p *Periph) PU15(n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU15)}}
}

type PDCR uint32

type RPDCR struct{ mmio.U32 }

func (r *RPDCR) LoadBits(mask PDCR) PDCR { return PDCR(r.U32.LoadBits(uint32(mask))) }
func (r *RPDCR) StoreBits(mask, b PDCR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPDCR) SetBits(mask PDCR)       { r.U32.SetBits(uint32(mask)) }
func (r *RPDCR) ClearBits(mask PDCR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RPDCR) Load() PDCR              { return PDCR(r.U32.Load()) }
func (r *RPDCR) Store(b PDCR)            { r.U32.Store(uint32(b)) }

type RMPDCR struct{ mmio.UM32 }

func (rm RMPDCR) Load() PDCR   { return PDCR(rm.UM32.Load()) }
func (rm RMPDCR) Store(b PDCR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) PD0(n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD0)}}
}

func (p *Periph) PD1(n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD1)}}
}

func (p *Periph) PD2(n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD2)}}
}

func (p *Periph) PD3(n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD3)}}
}

func (p *Periph) PD4(n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD4)}}
}

func (p *Periph) PD5(n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD5)}}
}

func (p *Periph) PD6(n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD6)}}
}

func (p *Periph) PD7(n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD7)}}
}

func (p *Periph) PD8(n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD8)}}
}

func (p *Periph) PD9(n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD9)}}
}

func (p *Periph) PD10(n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD10)}}
}

func (p *Periph) PD11(n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD11)}}
}

func (p *Periph) PD12(n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD12)}}
}

func (p *Periph) PD14(n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD14)}}
}

type CR5 uint32

type RCR5 struct{ mmio.U32 }

func (r *RCR5) LoadBits(mask CR5) CR5 { return CR5(r.U32.LoadBits(uint32(mask))) }
func (r *RCR5) StoreBits(mask, b CR5) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR5) SetBits(mask CR5)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR5) ClearBits(mask CR5)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR5) Load() CR5             { return CR5(r.U32.Load()) }
func (r *RCR5) Store(b CR5)           { r.U32.Store(uint32(b)) }

type RMCR5 struct{ mmio.UM32 }

func (rm RMCR5) Load() CR5   { return CR5(rm.UM32.Load()) }
func (rm RMCR5) Store(b CR5) { rm.UM32.Store(uint32(b)) }

func (p *Periph) R1MODE() RMCR5 {
	return RMCR5{mmio.UM32{&p.CR5.U32, uint32(R1MODE)}}
}
