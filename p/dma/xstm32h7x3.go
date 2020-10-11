// DO NOT EDIT THIS FILE. GENERATED BY xgen.

// +build stm32h7x3

package dma

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	ISR  [2]RISR
	IFCR [2]RIFCR
	S    [8]RS
}

func DMA1() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.DMA1_BASE))) }
func DMA2() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.DMA2_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

func (p *Periph) Bus() bus.Bus {
	return bus.AHB1
}

type ISR uint32

type RISR struct{ mmio.U32 }

func (r *RISR) LoadBits(mask ISR) ISR { return ISR(r.U32.LoadBits(uint32(mask))) }
func (r *RISR) StoreBits(mask, b ISR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RISR) SetBits(mask ISR)      { r.U32.SetBits(uint32(mask)) }
func (r *RISR) ClearBits(mask ISR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RISR) Load() ISR             { return ISR(r.U32.Load()) }
func (r *RISR) Store(b ISR)           { r.U32.Store(uint32(b)) }

type RMISR struct{ mmio.UM32 }

func (rm RMISR) Load() ISR   { return ISR(rm.UM32.Load()) }
func (rm RMISR) Store(b ISR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) FEIF0(n int) RMISR {
	return RMISR{mmio.UM32{&p.ISR[n].U32, uint32(FEIF0)}}
}

func (p *Periph) DMEIF0(n int) RMISR {
	return RMISR{mmio.UM32{&p.ISR[n].U32, uint32(DMEIF0)}}
}

func (p *Periph) TEIF0(n int) RMISR {
	return RMISR{mmio.UM32{&p.ISR[n].U32, uint32(TEIF0)}}
}

func (p *Periph) HTIF0(n int) RMISR {
	return RMISR{mmio.UM32{&p.ISR[n].U32, uint32(HTIF0)}}
}

func (p *Periph) TCIF0(n int) RMISR {
	return RMISR{mmio.UM32{&p.ISR[n].U32, uint32(TCIF0)}}
}

func (p *Periph) FEIF1(n int) RMISR {
	return RMISR{mmio.UM32{&p.ISR[n].U32, uint32(FEIF1)}}
}

func (p *Periph) DMEIF1(n int) RMISR {
	return RMISR{mmio.UM32{&p.ISR[n].U32, uint32(DMEIF1)}}
}

func (p *Periph) TEIF1(n int) RMISR {
	return RMISR{mmio.UM32{&p.ISR[n].U32, uint32(TEIF1)}}
}

func (p *Periph) HTIF1(n int) RMISR {
	return RMISR{mmio.UM32{&p.ISR[n].U32, uint32(HTIF1)}}
}

func (p *Periph) TCIF1(n int) RMISR {
	return RMISR{mmio.UM32{&p.ISR[n].U32, uint32(TCIF1)}}
}

func (p *Periph) FEIF2(n int) RMISR {
	return RMISR{mmio.UM32{&p.ISR[n].U32, uint32(FEIF2)}}
}

func (p *Periph) DMEIF2(n int) RMISR {
	return RMISR{mmio.UM32{&p.ISR[n].U32, uint32(DMEIF2)}}
}

func (p *Periph) TEIF2(n int) RMISR {
	return RMISR{mmio.UM32{&p.ISR[n].U32, uint32(TEIF2)}}
}

func (p *Periph) HTIF2(n int) RMISR {
	return RMISR{mmio.UM32{&p.ISR[n].U32, uint32(HTIF2)}}
}

func (p *Periph) TCIF2(n int) RMISR {
	return RMISR{mmio.UM32{&p.ISR[n].U32, uint32(TCIF2)}}
}

func (p *Periph) FEIF3(n int) RMISR {
	return RMISR{mmio.UM32{&p.ISR[n].U32, uint32(FEIF3)}}
}

func (p *Periph) DMEIF3(n int) RMISR {
	return RMISR{mmio.UM32{&p.ISR[n].U32, uint32(DMEIF3)}}
}

func (p *Periph) TEIF3(n int) RMISR {
	return RMISR{mmio.UM32{&p.ISR[n].U32, uint32(TEIF3)}}
}

func (p *Periph) HTIF3(n int) RMISR {
	return RMISR{mmio.UM32{&p.ISR[n].U32, uint32(HTIF3)}}
}

func (p *Periph) TCIF3(n int) RMISR {
	return RMISR{mmio.UM32{&p.ISR[n].U32, uint32(TCIF3)}}
}

type IFCR uint32

type RIFCR struct{ mmio.U32 }

func (r *RIFCR) LoadBits(mask IFCR) IFCR { return IFCR(r.U32.LoadBits(uint32(mask))) }
func (r *RIFCR) StoreBits(mask, b IFCR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIFCR) SetBits(mask IFCR)       { r.U32.SetBits(uint32(mask)) }
func (r *RIFCR) ClearBits(mask IFCR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RIFCR) Load() IFCR              { return IFCR(r.U32.Load()) }
func (r *RIFCR) Store(b IFCR)            { r.U32.Store(uint32(b)) }

type RMIFCR struct{ mmio.UM32 }

func (rm RMIFCR) Load() IFCR   { return IFCR(rm.UM32.Load()) }
func (rm RMIFCR) Store(b IFCR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) CFEIF0(n int) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR[n].U32, uint32(CFEIF0)}}
}

func (p *Periph) CDMEIF0(n int) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR[n].U32, uint32(CDMEIF0)}}
}

func (p *Periph) CTEIF0(n int) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR[n].U32, uint32(CTEIF0)}}
}

func (p *Periph) CHTIF0(n int) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR[n].U32, uint32(CHTIF0)}}
}

func (p *Periph) CTCIF0(n int) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR[n].U32, uint32(CTCIF0)}}
}

func (p *Periph) CFEIF1(n int) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR[n].U32, uint32(CFEIF1)}}
}

func (p *Periph) CDMEIF1(n int) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR[n].U32, uint32(CDMEIF1)}}
}

func (p *Periph) CTEIF1(n int) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR[n].U32, uint32(CTEIF1)}}
}

func (p *Periph) CHTIF1(n int) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR[n].U32, uint32(CHTIF1)}}
}

func (p *Periph) CTCIF1(n int) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR[n].U32, uint32(CTCIF1)}}
}

func (p *Periph) CFEIF2(n int) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR[n].U32, uint32(CFEIF2)}}
}

func (p *Periph) CDMEIF2(n int) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR[n].U32, uint32(CDMEIF2)}}
}

func (p *Periph) CTEIF2(n int) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR[n].U32, uint32(CTEIF2)}}
}

func (p *Periph) CHTIF2(n int) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR[n].U32, uint32(CHTIF2)}}
}

func (p *Periph) CTCIF2(n int) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR[n].U32, uint32(CTCIF2)}}
}

func (p *Periph) CFEIF3(n int) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR[n].U32, uint32(CFEIF3)}}
}

func (p *Periph) CDMEIF3(n int) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR[n].U32, uint32(CDMEIF3)}}
}

func (p *Periph) CTEIF3(n int) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR[n].U32, uint32(CTEIF3)}}
}

func (p *Periph) CHTIF3(n int) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR[n].U32, uint32(CHTIF3)}}
}

func (p *Periph) CTCIF3(n int) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR[n].U32, uint32(CTCIF3)}}
}

type RS struct {
	CR   RCR
	NDTR RNDTR
	PAR  RPAR
	M0AR RM0AR
	M1AR RM1AR
	FCR  RFCR
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

func (p *Periph) EN(n int) RMCR {
	return RMCR{mmio.UM32{&p.S[n].CR.U32, uint32(EN)}}
}

func (p *Periph) DMEIE(n int) RMCR {
	return RMCR{mmio.UM32{&p.S[n].CR.U32, uint32(DMEIE)}}
}

func (p *Periph) TEIE(n int) RMCR {
	return RMCR{mmio.UM32{&p.S[n].CR.U32, uint32(TEIE)}}
}

func (p *Periph) HTIE(n int) RMCR {
	return RMCR{mmio.UM32{&p.S[n].CR.U32, uint32(HTIE)}}
}

func (p *Periph) TCIE(n int) RMCR {
	return RMCR{mmio.UM32{&p.S[n].CR.U32, uint32(TCIE)}}
}

func (p *Periph) PFCTRL(n int) RMCR {
	return RMCR{mmio.UM32{&p.S[n].CR.U32, uint32(PFCTRL)}}
}

func (p *Periph) DIR(n int) RMCR {
	return RMCR{mmio.UM32{&p.S[n].CR.U32, uint32(DIR)}}
}

func (p *Periph) CIRC(n int) RMCR {
	return RMCR{mmio.UM32{&p.S[n].CR.U32, uint32(CIRC)}}
}

func (p *Periph) PINC(n int) RMCR {
	return RMCR{mmio.UM32{&p.S[n].CR.U32, uint32(PINC)}}
}

func (p *Periph) MINC(n int) RMCR {
	return RMCR{mmio.UM32{&p.S[n].CR.U32, uint32(MINC)}}
}

func (p *Periph) PSIZE(n int) RMCR {
	return RMCR{mmio.UM32{&p.S[n].CR.U32, uint32(PSIZE)}}
}

func (p *Periph) MSIZE(n int) RMCR {
	return RMCR{mmio.UM32{&p.S[n].CR.U32, uint32(MSIZE)}}
}

func (p *Periph) PINCOS(n int) RMCR {
	return RMCR{mmio.UM32{&p.S[n].CR.U32, uint32(PINCOS)}}
}

func (p *Periph) PL(n int) RMCR {
	return RMCR{mmio.UM32{&p.S[n].CR.U32, uint32(PL)}}
}

func (p *Periph) DBM(n int) RMCR {
	return RMCR{mmio.UM32{&p.S[n].CR.U32, uint32(DBM)}}
}

func (p *Periph) CT(n int) RMCR {
	return RMCR{mmio.UM32{&p.S[n].CR.U32, uint32(CT)}}
}

func (p *Periph) PBURST(n int) RMCR {
	return RMCR{mmio.UM32{&p.S[n].CR.U32, uint32(PBURST)}}
}

func (p *Periph) MBURST(n int) RMCR {
	return RMCR{mmio.UM32{&p.S[n].CR.U32, uint32(MBURST)}}
}

type NDTR uint32

type RNDTR struct{ mmio.U32 }

func (r *RNDTR) LoadBits(mask NDTR) NDTR { return NDTR(r.U32.LoadBits(uint32(mask))) }
func (r *RNDTR) StoreBits(mask, b NDTR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RNDTR) SetBits(mask NDTR)       { r.U32.SetBits(uint32(mask)) }
func (r *RNDTR) ClearBits(mask NDTR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RNDTR) Load() NDTR              { return NDTR(r.U32.Load()) }
func (r *RNDTR) Store(b NDTR)            { r.U32.Store(uint32(b)) }

type RMNDTR struct{ mmio.UM32 }

func (rm RMNDTR) Load() NDTR   { return NDTR(rm.UM32.Load()) }
func (rm RMNDTR) Store(b NDTR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) NDT(n int) RMNDTR {
	return RMNDTR{mmio.UM32{&p.S[n].NDTR.U32, uint32(NDT)}}
}

type PAR uint32

type RPAR struct{ mmio.U32 }

func (r *RPAR) LoadBits(mask PAR) PAR { return PAR(r.U32.LoadBits(uint32(mask))) }
func (r *RPAR) StoreBits(mask, b PAR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPAR) SetBits(mask PAR)      { r.U32.SetBits(uint32(mask)) }
func (r *RPAR) ClearBits(mask PAR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPAR) Load() PAR             { return PAR(r.U32.Load()) }
func (r *RPAR) Store(b PAR)           { r.U32.Store(uint32(b)) }

type RMPAR struct{ mmio.UM32 }

func (rm RMPAR) Load() PAR   { return PAR(rm.UM32.Load()) }
func (rm RMPAR) Store(b PAR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) PA(n int) RMPAR {
	return RMPAR{mmio.UM32{&p.S[n].PAR.U32, uint32(PA)}}
}

type M0AR uint32

type RM0AR struct{ mmio.U32 }

func (r *RM0AR) LoadBits(mask M0AR) M0AR { return M0AR(r.U32.LoadBits(uint32(mask))) }
func (r *RM0AR) StoreBits(mask, b M0AR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RM0AR) SetBits(mask M0AR)       { r.U32.SetBits(uint32(mask)) }
func (r *RM0AR) ClearBits(mask M0AR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RM0AR) Load() M0AR              { return M0AR(r.U32.Load()) }
func (r *RM0AR) Store(b M0AR)            { r.U32.Store(uint32(b)) }

type RMM0AR struct{ mmio.UM32 }

func (rm RMM0AR) Load() M0AR   { return M0AR(rm.UM32.Load()) }
func (rm RMM0AR) Store(b M0AR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) M0A(n int) RMM0AR {
	return RMM0AR{mmio.UM32{&p.S[n].M0AR.U32, uint32(M0A)}}
}

type M1AR uint32

type RM1AR struct{ mmio.U32 }

func (r *RM1AR) LoadBits(mask M1AR) M1AR { return M1AR(r.U32.LoadBits(uint32(mask))) }
func (r *RM1AR) StoreBits(mask, b M1AR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RM1AR) SetBits(mask M1AR)       { r.U32.SetBits(uint32(mask)) }
func (r *RM1AR) ClearBits(mask M1AR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RM1AR) Load() M1AR              { return M1AR(r.U32.Load()) }
func (r *RM1AR) Store(b M1AR)            { r.U32.Store(uint32(b)) }

type RMM1AR struct{ mmio.UM32 }

func (rm RMM1AR) Load() M1AR   { return M1AR(rm.UM32.Load()) }
func (rm RMM1AR) Store(b M1AR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) M1A(n int) RMM1AR {
	return RMM1AR{mmio.UM32{&p.S[n].M1AR.U32, uint32(M1A)}}
}

type FCR uint32

type RFCR struct{ mmio.U32 }

func (r *RFCR) LoadBits(mask FCR) FCR { return FCR(r.U32.LoadBits(uint32(mask))) }
func (r *RFCR) StoreBits(mask, b FCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFCR) SetBits(mask FCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RFCR) ClearBits(mask FCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFCR) Load() FCR             { return FCR(r.U32.Load()) }
func (r *RFCR) Store(b FCR)           { r.U32.Store(uint32(b)) }

type RMFCR struct{ mmio.UM32 }

func (rm RMFCR) Load() FCR   { return FCR(rm.UM32.Load()) }
func (rm RMFCR) Store(b FCR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) FTH(n int) RMFCR {
	return RMFCR{mmio.UM32{&p.S[n].FCR.U32, uint32(FTH)}}
}

func (p *Periph) DMDIS(n int) RMFCR {
	return RMFCR{mmio.UM32{&p.S[n].FCR.U32, uint32(DMDIS)}}
}

func (p *Periph) FS(n int) RMFCR {
	return RMFCR{mmio.UM32{&p.S[n].FCR.U32, uint32(FS)}}
}

func (p *Periph) FEIE(n int) RMFCR {
	return RMFCR{mmio.UM32{&p.S[n].FCR.U32, uint32(FEIE)}}
}
