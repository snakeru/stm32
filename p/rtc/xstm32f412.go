// DO NOT EDIT THIS FILE. GENERATED BY xgen.

// +build stm32f412

package rtc

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	TR       RTR
	DR       RDR
	CR       RCR
	ISR      RISR
	PRER     RPRER
	WUTR     RWUTR
	CALIBR   RCALIBR
	ALRMAR   RALRMR
	ALRMBR   RALRMR
	WPR      RWPR
	SSR      RSSR
	SHIFTR   RSHIFTR
	TSTR     RTR
	TSDR     RDR
	TSSSR    RSSR
	CALR     RCALR
	TAFCR    RTAFCR
	ALRMASSR RALRMSSR
	ALRMBSSR RALRMSSR
	_        uint32
	BKPR     [20]RBKPR
}

func RTC() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.RTC_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type TR uint32

type RTR struct{ mmio.U32 }

func (r *RTR) LoadBits(mask TR) TR  { return TR(r.U32.LoadBits(uint32(mask))) }
func (r *RTR) StoreBits(mask, b TR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTR) SetBits(mask TR)      { r.U32.SetBits(uint32(mask)) }
func (r *RTR) ClearBits(mask TR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTR) Load() TR             { return TR(r.U32.Load()) }
func (r *RTR) Store(b TR)           { r.U32.Store(uint32(b)) }

type RMTR struct{ mmio.UM32 }

func (rm RMTR) Load() TR   { return TR(rm.UM32.Load()) }
func (rm RMTR) Store(b TR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) SU() RMTR {
	return RMTR{mmio.UM32{&p.TR.U32, uint32(SU)}}
}

func (p *Periph) ST() RMTR {
	return RMTR{mmio.UM32{&p.TR.U32, uint32(ST)}}
}

func (p *Periph) MNU() RMTR {
	return RMTR{mmio.UM32{&p.TR.U32, uint32(MNU)}}
}

func (p *Periph) MNT() RMTR {
	return RMTR{mmio.UM32{&p.TR.U32, uint32(MNT)}}
}

func (p *Periph) HU() RMTR {
	return RMTR{mmio.UM32{&p.TR.U32, uint32(HU)}}
}

func (p *Periph) HT() RMTR {
	return RMTR{mmio.UM32{&p.TR.U32, uint32(HT)}}
}

func (p *Periph) PM() RMTR {
	return RMTR{mmio.UM32{&p.TR.U32, uint32(PM)}}
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

func (p *Periph) DU() RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(DU)}}
}

func (p *Periph) DT() RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(DT)}}
}

func (p *Periph) MU() RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(MU)}}
}

func (p *Periph) MT() RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(MT)}}
}

func (p *Periph) WDU() RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(WDU)}}
}

func (p *Periph) YU() RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(YU)}}
}

func (p *Periph) YT() RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(YT)}}
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

func (p *Periph) WCKSEL() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(WCKSEL)}}
}

func (p *Periph) TSEDGE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TSEDGE)}}
}

func (p *Periph) REFCKON() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(REFCKON)}}
}

func (p *Periph) BYPSHAD() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(BYPSHAD)}}
}

func (p *Periph) FMT() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(FMT)}}
}

func (p *Periph) DCE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DCE)}}
}

func (p *Periph) ALRAE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ALRAE)}}
}

func (p *Periph) ALRBE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ALRBE)}}
}

func (p *Periph) WUTE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(WUTE)}}
}

func (p *Periph) TSE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TSE)}}
}

func (p *Periph) ALRAIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ALRAIE)}}
}

func (p *Periph) ALRBIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ALRBIE)}}
}

func (p *Periph) WUTIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(WUTIE)}}
}

func (p *Periph) TSIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TSIE)}}
}

func (p *Periph) ADD1H() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ADD1H)}}
}

func (p *Periph) SUB1H() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(SUB1H)}}
}

func (p *Periph) BKP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(BKP)}}
}

func (p *Periph) COSEL() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(COSEL)}}
}

func (p *Periph) POL() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(POL)}}
}

func (p *Periph) OSEL() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(OSEL)}}
}

func (p *Periph) COE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(COE)}}
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

func (p *Periph) ALRAWF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(ALRAWF)}}
}

func (p *Periph) ALRBWF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(ALRBWF)}}
}

func (p *Periph) WUTWF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(WUTWF)}}
}

func (p *Periph) SHPF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(SHPF)}}
}

func (p *Periph) INITS() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(INITS)}}
}

func (p *Periph) RSF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(RSF)}}
}

func (p *Periph) INITF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(INITF)}}
}

func (p *Periph) INIT() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(INIT)}}
}

func (p *Periph) ALRAF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(ALRAF)}}
}

func (p *Periph) ALRBF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(ALRBF)}}
}

func (p *Periph) WUTF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(WUTF)}}
}

func (p *Periph) TSF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TSF)}}
}

func (p *Periph) TSOVF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TSOVF)}}
}

func (p *Periph) TAMP1F() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TAMP1F)}}
}

func (p *Periph) TAMP2F() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TAMP2F)}}
}

func (p *Periph) RECALPF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(RECALPF)}}
}

type PRER uint32

type RPRER struct{ mmio.U32 }

func (r *RPRER) LoadBits(mask PRER) PRER { return PRER(r.U32.LoadBits(uint32(mask))) }
func (r *RPRER) StoreBits(mask, b PRER)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPRER) SetBits(mask PRER)       { r.U32.SetBits(uint32(mask)) }
func (r *RPRER) ClearBits(mask PRER)     { r.U32.ClearBits(uint32(mask)) }
func (r *RPRER) Load() PRER              { return PRER(r.U32.Load()) }
func (r *RPRER) Store(b PRER)            { r.U32.Store(uint32(b)) }

type RMPRER struct{ mmio.UM32 }

func (rm RMPRER) Load() PRER   { return PRER(rm.UM32.Load()) }
func (rm RMPRER) Store(b PRER) { rm.UM32.Store(uint32(b)) }

func (p *Periph) PREDIV_S() RMPRER {
	return RMPRER{mmio.UM32{&p.PRER.U32, uint32(PREDIV_S)}}
}

func (p *Periph) PREDIV_A() RMPRER {
	return RMPRER{mmio.UM32{&p.PRER.U32, uint32(PREDIV_A)}}
}

type WUTR uint32

type RWUTR struct{ mmio.U32 }

func (r *RWUTR) LoadBits(mask WUTR) WUTR { return WUTR(r.U32.LoadBits(uint32(mask))) }
func (r *RWUTR) StoreBits(mask, b WUTR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RWUTR) SetBits(mask WUTR)       { r.U32.SetBits(uint32(mask)) }
func (r *RWUTR) ClearBits(mask WUTR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RWUTR) Load() WUTR              { return WUTR(r.U32.Load()) }
func (r *RWUTR) Store(b WUTR)            { r.U32.Store(uint32(b)) }

type RMWUTR struct{ mmio.UM32 }

func (rm RMWUTR) Load() WUTR   { return WUTR(rm.UM32.Load()) }
func (rm RMWUTR) Store(b WUTR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) WUT() RMWUTR {
	return RMWUTR{mmio.UM32{&p.WUTR.U32, uint32(WUT)}}
}

type CALIBR uint32

type RCALIBR struct{ mmio.U32 }

func (r *RCALIBR) LoadBits(mask CALIBR) CALIBR { return CALIBR(r.U32.LoadBits(uint32(mask))) }
func (r *RCALIBR) StoreBits(mask, b CALIBR)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCALIBR) SetBits(mask CALIBR)         { r.U32.SetBits(uint32(mask)) }
func (r *RCALIBR) ClearBits(mask CALIBR)       { r.U32.ClearBits(uint32(mask)) }
func (r *RCALIBR) Load() CALIBR                { return CALIBR(r.U32.Load()) }
func (r *RCALIBR) Store(b CALIBR)              { r.U32.Store(uint32(b)) }

type RMCALIBR struct{ mmio.UM32 }

func (rm RMCALIBR) Load() CALIBR   { return CALIBR(rm.UM32.Load()) }
func (rm RMCALIBR) Store(b CALIBR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) DC() RMCALIBR {
	return RMCALIBR{mmio.UM32{&p.CALIBR.U32, uint32(DC)}}
}

func (p *Periph) DCS() RMCALIBR {
	return RMCALIBR{mmio.UM32{&p.CALIBR.U32, uint32(DCS)}}
}

type ALRMR uint32

type RALRMR struct{ mmio.U32 }

func (r *RALRMR) LoadBits(mask ALRMR) ALRMR { return ALRMR(r.U32.LoadBits(uint32(mask))) }
func (r *RALRMR) StoreBits(mask, b ALRMR)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RALRMR) SetBits(mask ALRMR)        { r.U32.SetBits(uint32(mask)) }
func (r *RALRMR) ClearBits(mask ALRMR)      { r.U32.ClearBits(uint32(mask)) }
func (r *RALRMR) Load() ALRMR               { return ALRMR(r.U32.Load()) }
func (r *RALRMR) Store(b ALRMR)             { r.U32.Store(uint32(b)) }

type RMALRMR struct{ mmio.UM32 }

func (rm RMALRMR) Load() ALRMR   { return ALRMR(rm.UM32.Load()) }
func (rm RMALRMR) Store(b ALRMR) { rm.UM32.Store(uint32(b)) }

type WPR uint32

type RWPR struct{ mmio.U32 }

func (r *RWPR) LoadBits(mask WPR) WPR { return WPR(r.U32.LoadBits(uint32(mask))) }
func (r *RWPR) StoreBits(mask, b WPR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RWPR) SetBits(mask WPR)      { r.U32.SetBits(uint32(mask)) }
func (r *RWPR) ClearBits(mask WPR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RWPR) Load() WPR             { return WPR(r.U32.Load()) }
func (r *RWPR) Store(b WPR)           { r.U32.Store(uint32(b)) }

type RMWPR struct{ mmio.UM32 }

func (rm RMWPR) Load() WPR   { return WPR(rm.UM32.Load()) }
func (rm RMWPR) Store(b WPR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) KEY() RMWPR {
	return RMWPR{mmio.UM32{&p.WPR.U32, uint32(KEY)}}
}

type SSR uint32

type RSSR struct{ mmio.U32 }

func (r *RSSR) LoadBits(mask SSR) SSR { return SSR(r.U32.LoadBits(uint32(mask))) }
func (r *RSSR) StoreBits(mask, b SSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSSR) SetBits(mask SSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSSR) ClearBits(mask SSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSSR) Load() SSR             { return SSR(r.U32.Load()) }
func (r *RSSR) Store(b SSR)           { r.U32.Store(uint32(b)) }

type RMSSR struct{ mmio.UM32 }

func (rm RMSSR) Load() SSR   { return SSR(rm.UM32.Load()) }
func (rm RMSSR) Store(b SSR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) SS() RMSSR {
	return RMSSR{mmio.UM32{&p.SSR.U32, uint32(SS)}}
}

type SHIFTR uint32

type RSHIFTR struct{ mmio.U32 }

func (r *RSHIFTR) LoadBits(mask SHIFTR) SHIFTR { return SHIFTR(r.U32.LoadBits(uint32(mask))) }
func (r *RSHIFTR) StoreBits(mask, b SHIFTR)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSHIFTR) SetBits(mask SHIFTR)         { r.U32.SetBits(uint32(mask)) }
func (r *RSHIFTR) ClearBits(mask SHIFTR)       { r.U32.ClearBits(uint32(mask)) }
func (r *RSHIFTR) Load() SHIFTR                { return SHIFTR(r.U32.Load()) }
func (r *RSHIFTR) Store(b SHIFTR)              { r.U32.Store(uint32(b)) }

type RMSHIFTR struct{ mmio.UM32 }

func (rm RMSHIFTR) Load() SHIFTR   { return SHIFTR(rm.UM32.Load()) }
func (rm RMSHIFTR) Store(b SHIFTR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) SUBFS() RMSHIFTR {
	return RMSHIFTR{mmio.UM32{&p.SHIFTR.U32, uint32(SUBFS)}}
}

func (p *Periph) ADD1S() RMSHIFTR {
	return RMSHIFTR{mmio.UM32{&p.SHIFTR.U32, uint32(ADD1S)}}
}

type CALR uint32

type RCALR struct{ mmio.U32 }

func (r *RCALR) LoadBits(mask CALR) CALR { return CALR(r.U32.LoadBits(uint32(mask))) }
func (r *RCALR) StoreBits(mask, b CALR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCALR) SetBits(mask CALR)       { r.U32.SetBits(uint32(mask)) }
func (r *RCALR) ClearBits(mask CALR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RCALR) Load() CALR              { return CALR(r.U32.Load()) }
func (r *RCALR) Store(b CALR)            { r.U32.Store(uint32(b)) }

type RMCALR struct{ mmio.UM32 }

func (rm RMCALR) Load() CALR   { return CALR(rm.UM32.Load()) }
func (rm RMCALR) Store(b CALR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) CALM() RMCALR {
	return RMCALR{mmio.UM32{&p.CALR.U32, uint32(CALM)}}
}

func (p *Periph) CALW16() RMCALR {
	return RMCALR{mmio.UM32{&p.CALR.U32, uint32(CALW16)}}
}

func (p *Periph) CALW8() RMCALR {
	return RMCALR{mmio.UM32{&p.CALR.U32, uint32(CALW8)}}
}

func (p *Periph) CALP() RMCALR {
	return RMCALR{mmio.UM32{&p.CALR.U32, uint32(CALP)}}
}

type TAFCR uint32

type RTAFCR struct{ mmio.U32 }

func (r *RTAFCR) LoadBits(mask TAFCR) TAFCR { return TAFCR(r.U32.LoadBits(uint32(mask))) }
func (r *RTAFCR) StoreBits(mask, b TAFCR)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTAFCR) SetBits(mask TAFCR)        { r.U32.SetBits(uint32(mask)) }
func (r *RTAFCR) ClearBits(mask TAFCR)      { r.U32.ClearBits(uint32(mask)) }
func (r *RTAFCR) Load() TAFCR               { return TAFCR(r.U32.Load()) }
func (r *RTAFCR) Store(b TAFCR)             { r.U32.Store(uint32(b)) }

type RMTAFCR struct{ mmio.UM32 }

func (rm RMTAFCR) Load() TAFCR   { return TAFCR(rm.UM32.Load()) }
func (rm RMTAFCR) Store(b TAFCR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) TAMP1E() RMTAFCR {
	return RMTAFCR{mmio.UM32{&p.TAFCR.U32, uint32(TAMP1E)}}
}

func (p *Periph) TAMP1TRG() RMTAFCR {
	return RMTAFCR{mmio.UM32{&p.TAFCR.U32, uint32(TAMP1TRG)}}
}

func (p *Periph) TAMPIE() RMTAFCR {
	return RMTAFCR{mmio.UM32{&p.TAFCR.U32, uint32(TAMPIE)}}
}

func (p *Periph) TAMP2E() RMTAFCR {
	return RMTAFCR{mmio.UM32{&p.TAFCR.U32, uint32(TAMP2E)}}
}

func (p *Periph) TAMP2TRG() RMTAFCR {
	return RMTAFCR{mmio.UM32{&p.TAFCR.U32, uint32(TAMP2TRG)}}
}

func (p *Periph) TAMPTS() RMTAFCR {
	return RMTAFCR{mmio.UM32{&p.TAFCR.U32, uint32(TAMPTS)}}
}

func (p *Periph) TAMPFREQ() RMTAFCR {
	return RMTAFCR{mmio.UM32{&p.TAFCR.U32, uint32(TAMPFREQ)}}
}

func (p *Periph) TAMPFLT() RMTAFCR {
	return RMTAFCR{mmio.UM32{&p.TAFCR.U32, uint32(TAMPFLT)}}
}

func (p *Periph) TAMPPRCH() RMTAFCR {
	return RMTAFCR{mmio.UM32{&p.TAFCR.U32, uint32(TAMPPRCH)}}
}

func (p *Periph) TAMPPUDIS() RMTAFCR {
	return RMTAFCR{mmio.UM32{&p.TAFCR.U32, uint32(TAMPPUDIS)}}
}

func (p *Periph) TAMP1INSEL() RMTAFCR {
	return RMTAFCR{mmio.UM32{&p.TAFCR.U32, uint32(TAMP1INSEL)}}
}

func (p *Periph) TSINSEL() RMTAFCR {
	return RMTAFCR{mmio.UM32{&p.TAFCR.U32, uint32(TSINSEL)}}
}

func (p *Periph) ALARMOUTTYPE() RMTAFCR {
	return RMTAFCR{mmio.UM32{&p.TAFCR.U32, uint32(ALARMOUTTYPE)}}
}

type ALRMSSR uint32

type RALRMSSR struct{ mmio.U32 }

func (r *RALRMSSR) LoadBits(mask ALRMSSR) ALRMSSR { return ALRMSSR(r.U32.LoadBits(uint32(mask))) }
func (r *RALRMSSR) StoreBits(mask, b ALRMSSR)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RALRMSSR) SetBits(mask ALRMSSR)          { r.U32.SetBits(uint32(mask)) }
func (r *RALRMSSR) ClearBits(mask ALRMSSR)        { r.U32.ClearBits(uint32(mask)) }
func (r *RALRMSSR) Load() ALRMSSR                 { return ALRMSSR(r.U32.Load()) }
func (r *RALRMSSR) Store(b ALRMSSR)               { r.U32.Store(uint32(b)) }

type RMALRMSSR struct{ mmio.UM32 }

func (rm RMALRMSSR) Load() ALRMSSR   { return ALRMSSR(rm.UM32.Load()) }
func (rm RMALRMSSR) Store(b ALRMSSR) { rm.UM32.Store(uint32(b)) }

type BKPR uint32

type RBKPR struct{ mmio.U32 }

func (r *RBKPR) LoadBits(mask BKPR) BKPR { return BKPR(r.U32.LoadBits(uint32(mask))) }
func (r *RBKPR) StoreBits(mask, b BKPR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBKPR) SetBits(mask BKPR)       { r.U32.SetBits(uint32(mask)) }
func (r *RBKPR) ClearBits(mask BKPR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RBKPR) Load() BKPR              { return BKPR(r.U32.Load()) }
func (r *RBKPR) Store(b BKPR)            { r.U32.Store(uint32(b)) }

type RMBKPR struct{ mmio.UM32 }

func (rm RMBKPR) Load() BKPR   { return BKPR(rm.UM32.Load()) }
func (rm RMBKPR) Store(b BKPR) { rm.UM32.Store(uint32(b)) }
