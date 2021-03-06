// DO NOT EDIT THIS FILE. GENERATED BY xgen.

// +build stm32h7x3

package pwr

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	CR1     RCR1
	CSR1    RCSR1
	CR2     RCR2
	CR3     RCR3
	CPUCR   RCPUCR
	_       uint32
	D3CR    RD3CR
	_       uint32
	WKUPCR  RWKUPCR
	WKUPFR  RWKUPFR
	WKUPEPR RWKUPEPR
}

func PWR() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.PWR_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
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

func (p *Periph) LPDS() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(LPDS)}}
}

func (p *Periph) PVDE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(PVDE)}}
}

func (p *Periph) PLS() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(PLS)}}
}

func (p *Periph) DBP() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(DBP)}}
}

func (p *Periph) FLPS() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(FLPS)}}
}

func (p *Periph) SVOS() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(SVOS)}}
}

func (p *Periph) AVDEN() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(AVDEN)}}
}

func (p *Periph) ALS() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(ALS)}}
}

type CSR1 uint32

type RCSR1 struct{ mmio.U32 }

func (r *RCSR1) LoadBits(mask CSR1) CSR1 { return CSR1(r.U32.LoadBits(uint32(mask))) }
func (r *RCSR1) StoreBits(mask, b CSR1)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCSR1) SetBits(mask CSR1)       { r.U32.SetBits(uint32(mask)) }
func (r *RCSR1) ClearBits(mask CSR1)     { r.U32.ClearBits(uint32(mask)) }
func (r *RCSR1) Load() CSR1              { return CSR1(r.U32.Load()) }
func (r *RCSR1) Store(b CSR1)            { r.U32.Store(uint32(b)) }

type RMCSR1 struct{ mmio.UM32 }

func (rm RMCSR1) Load() CSR1   { return CSR1(rm.UM32.Load()) }
func (rm RMCSR1) Store(b CSR1) { rm.UM32.Store(uint32(b)) }

func (p *Periph) PVDO() RMCSR1 {
	return RMCSR1{mmio.UM32{&p.CSR1.U32, uint32(PVDO)}}
}

func (p *Periph) ACTVOSRDY() RMCSR1 {
	return RMCSR1{mmio.UM32{&p.CSR1.U32, uint32(ACTVOSRDY)}}
}

func (p *Periph) ACTVOS() RMCSR1 {
	return RMCSR1{mmio.UM32{&p.CSR1.U32, uint32(ACTVOS)}}
}

func (p *Periph) AVDO() RMCSR1 {
	return RMCSR1{mmio.UM32{&p.CSR1.U32, uint32(AVDO)}}
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

func (p *Periph) BREN() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(BREN)}}
}

func (p *Periph) MONEN() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(MONEN)}}
}

func (p *Periph) BRRDY() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(BRRDY)}}
}

func (p *Periph) VBATL() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(VBATL)}}
}

func (p *Periph) VBATH() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(VBATH)}}
}

func (p *Periph) TEMPL() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(TEMPL)}}
}

func (p *Periph) TEMPH() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(TEMPH)}}
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

func (p *Periph) BYPASS() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(BYPASS)}}
}

func (p *Periph) LDOEN() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(LDOEN)}}
}

func (p *Periph) SCUEN() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(SCUEN)}}
}

func (p *Periph) VBE() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(VBE)}}
}

func (p *Periph) VBRS() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(VBRS)}}
}

func (p *Periph) USB33DEN() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(USB33DEN)}}
}

func (p *Periph) USBREGEN() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(USBREGEN)}}
}

func (p *Periph) USB33RDY() RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(USB33RDY)}}
}

type CPUCR uint32

type RCPUCR struct{ mmio.U32 }

func (r *RCPUCR) LoadBits(mask CPUCR) CPUCR { return CPUCR(r.U32.LoadBits(uint32(mask))) }
func (r *RCPUCR) StoreBits(mask, b CPUCR)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCPUCR) SetBits(mask CPUCR)        { r.U32.SetBits(uint32(mask)) }
func (r *RCPUCR) ClearBits(mask CPUCR)      { r.U32.ClearBits(uint32(mask)) }
func (r *RCPUCR) Load() CPUCR               { return CPUCR(r.U32.Load()) }
func (r *RCPUCR) Store(b CPUCR)             { r.U32.Store(uint32(b)) }

type RMCPUCR struct{ mmio.UM32 }

func (rm RMCPUCR) Load() CPUCR   { return CPUCR(rm.UM32.Load()) }
func (rm RMCPUCR) Store(b CPUCR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) PDDS_D1() RMCPUCR {
	return RMCPUCR{mmio.UM32{&p.CPUCR.U32, uint32(PDDS_D1)}}
}

func (p *Periph) PDDS_D2() RMCPUCR {
	return RMCPUCR{mmio.UM32{&p.CPUCR.U32, uint32(PDDS_D2)}}
}

func (p *Periph) PDDS_D3() RMCPUCR {
	return RMCPUCR{mmio.UM32{&p.CPUCR.U32, uint32(PDDS_D3)}}
}

func (p *Periph) STOPF() RMCPUCR {
	return RMCPUCR{mmio.UM32{&p.CPUCR.U32, uint32(STOPF)}}
}

func (p *Periph) SBF() RMCPUCR {
	return RMCPUCR{mmio.UM32{&p.CPUCR.U32, uint32(SBF)}}
}

func (p *Periph) SBF_D1() RMCPUCR {
	return RMCPUCR{mmio.UM32{&p.CPUCR.U32, uint32(SBF_D1)}}
}

func (p *Periph) SBF_D2() RMCPUCR {
	return RMCPUCR{mmio.UM32{&p.CPUCR.U32, uint32(SBF_D2)}}
}

func (p *Periph) CSSF() RMCPUCR {
	return RMCPUCR{mmio.UM32{&p.CPUCR.U32, uint32(CSSF)}}
}

func (p *Periph) RUN_D3() RMCPUCR {
	return RMCPUCR{mmio.UM32{&p.CPUCR.U32, uint32(RUN_D3)}}
}

type D3CR uint32

type RD3CR struct{ mmio.U32 }

func (r *RD3CR) LoadBits(mask D3CR) D3CR { return D3CR(r.U32.LoadBits(uint32(mask))) }
func (r *RD3CR) StoreBits(mask, b D3CR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RD3CR) SetBits(mask D3CR)       { r.U32.SetBits(uint32(mask)) }
func (r *RD3CR) ClearBits(mask D3CR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RD3CR) Load() D3CR              { return D3CR(r.U32.Load()) }
func (r *RD3CR) Store(b D3CR)            { r.U32.Store(uint32(b)) }

type RMD3CR struct{ mmio.UM32 }

func (rm RMD3CR) Load() D3CR   { return D3CR(rm.UM32.Load()) }
func (rm RMD3CR) Store(b D3CR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) VOSRDY() RMD3CR {
	return RMD3CR{mmio.UM32{&p.D3CR.U32, uint32(VOSRDY)}}
}

func (p *Periph) VOS() RMD3CR {
	return RMD3CR{mmio.UM32{&p.D3CR.U32, uint32(VOS)}}
}

type WKUPCR uint32

type RWKUPCR struct{ mmio.U32 }

func (r *RWKUPCR) LoadBits(mask WKUPCR) WKUPCR { return WKUPCR(r.U32.LoadBits(uint32(mask))) }
func (r *RWKUPCR) StoreBits(mask, b WKUPCR)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RWKUPCR) SetBits(mask WKUPCR)         { r.U32.SetBits(uint32(mask)) }
func (r *RWKUPCR) ClearBits(mask WKUPCR)       { r.U32.ClearBits(uint32(mask)) }
func (r *RWKUPCR) Load() WKUPCR                { return WKUPCR(r.U32.Load()) }
func (r *RWKUPCR) Store(b WKUPCR)              { r.U32.Store(uint32(b)) }

type RMWKUPCR struct{ mmio.UM32 }

func (rm RMWKUPCR) Load() WKUPCR   { return WKUPCR(rm.UM32.Load()) }
func (rm RMWKUPCR) Store(b WKUPCR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) WKUPC() RMWKUPCR {
	return RMWKUPCR{mmio.UM32{&p.WKUPCR.U32, uint32(WKUPC)}}
}

type WKUPFR uint32

type RWKUPFR struct{ mmio.U32 }

func (r *RWKUPFR) LoadBits(mask WKUPFR) WKUPFR { return WKUPFR(r.U32.LoadBits(uint32(mask))) }
func (r *RWKUPFR) StoreBits(mask, b WKUPFR)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RWKUPFR) SetBits(mask WKUPFR)         { r.U32.SetBits(uint32(mask)) }
func (r *RWKUPFR) ClearBits(mask WKUPFR)       { r.U32.ClearBits(uint32(mask)) }
func (r *RWKUPFR) Load() WKUPFR                { return WKUPFR(r.U32.Load()) }
func (r *RWKUPFR) Store(b WKUPFR)              { r.U32.Store(uint32(b)) }

type RMWKUPFR struct{ mmio.UM32 }

func (rm RMWKUPFR) Load() WKUPFR   { return WKUPFR(rm.UM32.Load()) }
func (rm RMWKUPFR) Store(b WKUPFR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) WKUPF1() RMWKUPFR {
	return RMWKUPFR{mmio.UM32{&p.WKUPFR.U32, uint32(WKUPF1)}}
}

func (p *Periph) WKUPF2() RMWKUPFR {
	return RMWKUPFR{mmio.UM32{&p.WKUPFR.U32, uint32(WKUPF2)}}
}

func (p *Periph) WKUPF3() RMWKUPFR {
	return RMWKUPFR{mmio.UM32{&p.WKUPFR.U32, uint32(WKUPF3)}}
}

func (p *Periph) WKUPF4() RMWKUPFR {
	return RMWKUPFR{mmio.UM32{&p.WKUPFR.U32, uint32(WKUPF4)}}
}

func (p *Periph) WKUPF5() RMWKUPFR {
	return RMWKUPFR{mmio.UM32{&p.WKUPFR.U32, uint32(WKUPF5)}}
}

func (p *Periph) WKUPF6() RMWKUPFR {
	return RMWKUPFR{mmio.UM32{&p.WKUPFR.U32, uint32(WKUPF6)}}
}

type WKUPEPR uint32

type RWKUPEPR struct{ mmio.U32 }

func (r *RWKUPEPR) LoadBits(mask WKUPEPR) WKUPEPR { return WKUPEPR(r.U32.LoadBits(uint32(mask))) }
func (r *RWKUPEPR) StoreBits(mask, b WKUPEPR)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RWKUPEPR) SetBits(mask WKUPEPR)          { r.U32.SetBits(uint32(mask)) }
func (r *RWKUPEPR) ClearBits(mask WKUPEPR)        { r.U32.ClearBits(uint32(mask)) }
func (r *RWKUPEPR) Load() WKUPEPR                 { return WKUPEPR(r.U32.Load()) }
func (r *RWKUPEPR) Store(b WKUPEPR)               { r.U32.Store(uint32(b)) }

type RMWKUPEPR struct{ mmio.UM32 }

func (rm RMWKUPEPR) Load() WKUPEPR   { return WKUPEPR(rm.UM32.Load()) }
func (rm RMWKUPEPR) Store(b WKUPEPR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) WKUPEN1() RMWKUPEPR {
	return RMWKUPEPR{mmio.UM32{&p.WKUPEPR.U32, uint32(WKUPEN1)}}
}

func (p *Periph) WKUPEN2() RMWKUPEPR {
	return RMWKUPEPR{mmio.UM32{&p.WKUPEPR.U32, uint32(WKUPEN2)}}
}

func (p *Periph) WKUPEN3() RMWKUPEPR {
	return RMWKUPEPR{mmio.UM32{&p.WKUPEPR.U32, uint32(WKUPEN3)}}
}

func (p *Periph) WKUPEN4() RMWKUPEPR {
	return RMWKUPEPR{mmio.UM32{&p.WKUPEPR.U32, uint32(WKUPEN4)}}
}

func (p *Periph) WKUPEN5() RMWKUPEPR {
	return RMWKUPEPR{mmio.UM32{&p.WKUPEPR.U32, uint32(WKUPEN5)}}
}

func (p *Periph) WKUPEN6() RMWKUPEPR {
	return RMWKUPEPR{mmio.UM32{&p.WKUPEPR.U32, uint32(WKUPEN6)}}
}

func (p *Periph) WKUPP1() RMWKUPEPR {
	return RMWKUPEPR{mmio.UM32{&p.WKUPEPR.U32, uint32(WKUPP1)}}
}

func (p *Periph) WKUPP2() RMWKUPEPR {
	return RMWKUPEPR{mmio.UM32{&p.WKUPEPR.U32, uint32(WKUPP2)}}
}

func (p *Periph) WKUPP3() RMWKUPEPR {
	return RMWKUPEPR{mmio.UM32{&p.WKUPEPR.U32, uint32(WKUPP3)}}
}

func (p *Periph) WKUPP4() RMWKUPEPR {
	return RMWKUPEPR{mmio.UM32{&p.WKUPEPR.U32, uint32(WKUPP4)}}
}

func (p *Periph) WKUPP5() RMWKUPEPR {
	return RMWKUPEPR{mmio.UM32{&p.WKUPEPR.U32, uint32(WKUPP5)}}
}

func (p *Periph) WKUPP6() RMWKUPEPR {
	return RMWKUPEPR{mmio.UM32{&p.WKUPEPR.U32, uint32(WKUPP6)}}
}

func (p *Periph) WKUPPUPD1() RMWKUPEPR {
	return RMWKUPEPR{mmio.UM32{&p.WKUPEPR.U32, uint32(WKUPPUPD1)}}
}

func (p *Periph) WKUPPUPD2() RMWKUPEPR {
	return RMWKUPEPR{mmio.UM32{&p.WKUPEPR.U32, uint32(WKUPPUPD2)}}
}

func (p *Periph) WKUPPUPD3() RMWKUPEPR {
	return RMWKUPEPR{mmio.UM32{&p.WKUPEPR.U32, uint32(WKUPPUPD3)}}
}

func (p *Periph) WKUPPUPD4() RMWKUPEPR {
	return RMWKUPEPR{mmio.UM32{&p.WKUPEPR.U32, uint32(WKUPPUPD4)}}
}

func (p *Periph) WKUPPUPD5() RMWKUPEPR {
	return RMWKUPEPR{mmio.UM32{&p.WKUPEPR.U32, uint32(WKUPPUPD5)}}
}

func (p *Periph) WKUPPUPD6() RMWKUPEPR {
	return RMWKUPEPR{mmio.UM32{&p.WKUPEPR.U32, uint32(WKUPPUPD6)}}
}
