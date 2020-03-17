// DO NOT EDIT THIS FILE. GENERATED BY xgen.

// +build stm32f412

package rcc

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	CR         RCR
	PLLCFGR    RPLLCFGR
	CFGR       RCFGR
	CIR        RCIR
	AHB1RSTR   RAHB1RSTR
	AHB2RSTR   RAHB2RSTR
	_          [2]uint32
	APB1RSTR   RAPB1RSTR
	APB2RSTR   RAPB2RSTR
	_          [2]uint32
	AHB1ENR    RAHB1ENR
	AHB2ENR    RAHB2ENR
	_          [2]uint32
	APB1ENR    RAPB1ENR
	APB2ENR    RAPB2ENR
	_          [2]uint32
	AHB1LPENR  RAHB1LPENR
	AHB2LPENR  RAHB2LPENR
	_          [2]uint32
	APB1LPENR  RAPB1LPENR
	APB2LPENR  RAPB2LPENR
	_          [2]uint32
	BDCR       RBDCR
	CSR        RCSR
	_          [2]uint32
	SSCGR      RSSCGR
	PLLI2SCFGR RPLLI2SCFGR
}

func RCC() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.RCC_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
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

func (p *Periph) HSION() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(HSION)}}
}

func (p *Periph) HSIRDY() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(HSIRDY)}}
}

func (p *Periph) HSITRIM() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(HSITRIM)}}
}

func (p *Periph) HSICAL() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(HSICAL)}}
}

func (p *Periph) HSEON() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(HSEON)}}
}

func (p *Periph) HSERDY() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(HSERDY)}}
}

func (p *Periph) HSEBYP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(HSEBYP)}}
}

func (p *Periph) CSSON() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(CSSON)}}
}

func (p *Periph) PLLON() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PLLON)}}
}

func (p *Periph) PLLRDY() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PLLRDY)}}
}

func (p *Periph) PLLI2SON() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PLLI2SON)}}
}

func (p *Periph) PLLI2SRDY() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PLLI2SRDY)}}
}

type PLLCFGR uint32

type RPLLCFGR struct{ mmio.U32 }

func (r *RPLLCFGR) LoadBits(mask PLLCFGR) PLLCFGR { return PLLCFGR(r.U32.LoadBits(uint32(mask))) }
func (r *RPLLCFGR) StoreBits(mask, b PLLCFGR)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPLLCFGR) SetBits(mask PLLCFGR)          { r.U32.SetBits(uint32(mask)) }
func (r *RPLLCFGR) ClearBits(mask PLLCFGR)        { r.U32.ClearBits(uint32(mask)) }
func (r *RPLLCFGR) Load() PLLCFGR                 { return PLLCFGR(r.U32.Load()) }
func (r *RPLLCFGR) Store(b PLLCFGR)               { r.U32.Store(uint32(b)) }

type RMPLLCFGR struct{ mmio.UM32 }

func (rm RMPLLCFGR) Load() PLLCFGR   { return PLLCFGR(rm.UM32.Load()) }
func (rm RMPLLCFGR) Store(b PLLCFGR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) PLLM() RMPLLCFGR {
	return RMPLLCFGR{mmio.UM32{&p.PLLCFGR.U32, uint32(PLLM)}}
}

func (p *Periph) PLLN() RMPLLCFGR {
	return RMPLLCFGR{mmio.UM32{&p.PLLCFGR.U32, uint32(PLLN)}}
}

func (p *Periph) PLLP() RMPLLCFGR {
	return RMPLLCFGR{mmio.UM32{&p.PLLCFGR.U32, uint32(PLLP)}}
}

func (p *Periph) PLLSRC() RMPLLCFGR {
	return RMPLLCFGR{mmio.UM32{&p.PLLCFGR.U32, uint32(PLLSRC)}}
}

func (p *Periph) PLLQ() RMPLLCFGR {
	return RMPLLCFGR{mmio.UM32{&p.PLLCFGR.U32, uint32(PLLQ)}}
}

type CFGR uint32

type RCFGR struct{ mmio.U32 }

func (r *RCFGR) LoadBits(mask CFGR) CFGR { return CFGR(r.U32.LoadBits(uint32(mask))) }
func (r *RCFGR) StoreBits(mask, b CFGR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR) SetBits(mask CFGR)       { r.U32.SetBits(uint32(mask)) }
func (r *RCFGR) ClearBits(mask CFGR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RCFGR) Load() CFGR              { return CFGR(r.U32.Load()) }
func (r *RCFGR) Store(b CFGR)            { r.U32.Store(uint32(b)) }

type RMCFGR struct{ mmio.UM32 }

func (rm RMCFGR) Load() CFGR   { return CFGR(rm.UM32.Load()) }
func (rm RMCFGR) Store(b CFGR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) SW() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(SW)}}
}

func (p *Periph) SWS() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(SWS)}}
}

func (p *Periph) HPRE() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(HPRE)}}
}

func (p *Periph) PPRE1() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(PPRE1)}}
}

func (p *Periph) PPRE2() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(PPRE2)}}
}

func (p *Periph) RTCPRE() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(RTCPRE)}}
}

func (p *Periph) MCO1() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(MCO1)}}
}

func (p *Periph) I2SSRC() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(I2SSRC)}}
}

func (p *Periph) MCO1PRE() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(MCO1PRE)}}
}

func (p *Periph) MCO2PRE() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(MCO2PRE)}}
}

func (p *Periph) MCO2() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(MCO2)}}
}

type CIR uint32

type RCIR struct{ mmio.U32 }

func (r *RCIR) LoadBits(mask CIR) CIR { return CIR(r.U32.LoadBits(uint32(mask))) }
func (r *RCIR) StoreBits(mask, b CIR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCIR) SetBits(mask CIR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCIR) ClearBits(mask CIR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCIR) Load() CIR             { return CIR(r.U32.Load()) }
func (r *RCIR) Store(b CIR)           { r.U32.Store(uint32(b)) }

type RMCIR struct{ mmio.UM32 }

func (rm RMCIR) Load() CIR   { return CIR(rm.UM32.Load()) }
func (rm RMCIR) Store(b CIR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) LSIRDYF() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(LSIRDYF)}}
}

func (p *Periph) LSERDYF() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(LSERDYF)}}
}

func (p *Periph) HSIRDYF() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(HSIRDYF)}}
}

func (p *Periph) HSERDYF() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(HSERDYF)}}
}

func (p *Periph) PLLRDYF() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(PLLRDYF)}}
}

func (p *Periph) PLLI2SRDYF() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(PLLI2SRDYF)}}
}

func (p *Periph) CSSF() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(CSSF)}}
}

func (p *Periph) LSIRDYIE() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(LSIRDYIE)}}
}

func (p *Periph) LSERDYIE() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(LSERDYIE)}}
}

func (p *Periph) HSIRDYIE() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(HSIRDYIE)}}
}

func (p *Periph) HSERDYIE() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(HSERDYIE)}}
}

func (p *Periph) PLLRDYIE() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(PLLRDYIE)}}
}

func (p *Periph) PLLI2SRDYIE() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(PLLI2SRDYIE)}}
}

func (p *Periph) LSIRDYC() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(LSIRDYC)}}
}

func (p *Periph) LSERDYC() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(LSERDYC)}}
}

func (p *Periph) HSIRDYC() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(HSIRDYC)}}
}

func (p *Periph) HSERDYC() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(HSERDYC)}}
}

func (p *Periph) PLLRDYC() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(PLLRDYC)}}
}

func (p *Periph) PLLI2SRDYC() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(PLLI2SRDYC)}}
}

func (p *Periph) CSSC() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(CSSC)}}
}

type AHB1RSTR uint32

type RAHB1RSTR struct{ mmio.U32 }

func (r *RAHB1RSTR) LoadBits(mask AHB1RSTR) AHB1RSTR { return AHB1RSTR(r.U32.LoadBits(uint32(mask))) }
func (r *RAHB1RSTR) StoreBits(mask, b AHB1RSTR)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAHB1RSTR) SetBits(mask AHB1RSTR)           { r.U32.SetBits(uint32(mask)) }
func (r *RAHB1RSTR) ClearBits(mask AHB1RSTR)         { r.U32.ClearBits(uint32(mask)) }
func (r *RAHB1RSTR) Load() AHB1RSTR                  { return AHB1RSTR(r.U32.Load()) }
func (r *RAHB1RSTR) Store(b AHB1RSTR)                { r.U32.Store(uint32(b)) }

type RMAHB1RSTR struct{ mmio.UM32 }

func (rm RMAHB1RSTR) Load() AHB1RSTR   { return AHB1RSTR(rm.UM32.Load()) }
func (rm RMAHB1RSTR) Store(b AHB1RSTR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) GPIOARST() RMAHB1RSTR {
	return RMAHB1RSTR{mmio.UM32{&p.AHB1RSTR.U32, uint32(GPIOARST)}}
}

func (p *Periph) GPIOBRST() RMAHB1RSTR {
	return RMAHB1RSTR{mmio.UM32{&p.AHB1RSTR.U32, uint32(GPIOBRST)}}
}

func (p *Periph) GPIOCRST() RMAHB1RSTR {
	return RMAHB1RSTR{mmio.UM32{&p.AHB1RSTR.U32, uint32(GPIOCRST)}}
}

func (p *Periph) GPIODRST() RMAHB1RSTR {
	return RMAHB1RSTR{mmio.UM32{&p.AHB1RSTR.U32, uint32(GPIODRST)}}
}

func (p *Periph) GPIOERST() RMAHB1RSTR {
	return RMAHB1RSTR{mmio.UM32{&p.AHB1RSTR.U32, uint32(GPIOERST)}}
}

func (p *Periph) GPIOFRST() RMAHB1RSTR {
	return RMAHB1RSTR{mmio.UM32{&p.AHB1RSTR.U32, uint32(GPIOFRST)}}
}

func (p *Periph) GPIOGRST() RMAHB1RSTR {
	return RMAHB1RSTR{mmio.UM32{&p.AHB1RSTR.U32, uint32(GPIOGRST)}}
}

func (p *Periph) GPIOHRST() RMAHB1RSTR {
	return RMAHB1RSTR{mmio.UM32{&p.AHB1RSTR.U32, uint32(GPIOHRST)}}
}

func (p *Periph) CRCRST() RMAHB1RSTR {
	return RMAHB1RSTR{mmio.UM32{&p.AHB1RSTR.U32, uint32(CRCRST)}}
}

func (p *Periph) DMA1RST() RMAHB1RSTR {
	return RMAHB1RSTR{mmio.UM32{&p.AHB1RSTR.U32, uint32(DMA1RST)}}
}

func (p *Periph) DMA2RST() RMAHB1RSTR {
	return RMAHB1RSTR{mmio.UM32{&p.AHB1RSTR.U32, uint32(DMA2RST)}}
}

type AHB2RSTR uint32

type RAHB2RSTR struct{ mmio.U32 }

func (r *RAHB2RSTR) LoadBits(mask AHB2RSTR) AHB2RSTR { return AHB2RSTR(r.U32.LoadBits(uint32(mask))) }
func (r *RAHB2RSTR) StoreBits(mask, b AHB2RSTR)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAHB2RSTR) SetBits(mask AHB2RSTR)           { r.U32.SetBits(uint32(mask)) }
func (r *RAHB2RSTR) ClearBits(mask AHB2RSTR)         { r.U32.ClearBits(uint32(mask)) }
func (r *RAHB2RSTR) Load() AHB2RSTR                  { return AHB2RSTR(r.U32.Load()) }
func (r *RAHB2RSTR) Store(b AHB2RSTR)                { r.U32.Store(uint32(b)) }

type RMAHB2RSTR struct{ mmio.UM32 }

func (rm RMAHB2RSTR) Load() AHB2RSTR   { return AHB2RSTR(rm.UM32.Load()) }
func (rm RMAHB2RSTR) Store(b AHB2RSTR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) RNGRST() RMAHB2RSTR {
	return RMAHB2RSTR{mmio.UM32{&p.AHB2RSTR.U32, uint32(RNGRST)}}
}

func (p *Periph) OTGFSRST() RMAHB2RSTR {
	return RMAHB2RSTR{mmio.UM32{&p.AHB2RSTR.U32, uint32(OTGFSRST)}}
}

type APB1RSTR uint32

type RAPB1RSTR struct{ mmio.U32 }

func (r *RAPB1RSTR) LoadBits(mask APB1RSTR) APB1RSTR { return APB1RSTR(r.U32.LoadBits(uint32(mask))) }
func (r *RAPB1RSTR) StoreBits(mask, b APB1RSTR)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAPB1RSTR) SetBits(mask APB1RSTR)           { r.U32.SetBits(uint32(mask)) }
func (r *RAPB1RSTR) ClearBits(mask APB1RSTR)         { r.U32.ClearBits(uint32(mask)) }
func (r *RAPB1RSTR) Load() APB1RSTR                  { return APB1RSTR(r.U32.Load()) }
func (r *RAPB1RSTR) Store(b APB1RSTR)                { r.U32.Store(uint32(b)) }

type RMAPB1RSTR struct{ mmio.UM32 }

func (rm RMAPB1RSTR) Load() APB1RSTR   { return APB1RSTR(rm.UM32.Load()) }
func (rm RMAPB1RSTR) Store(b APB1RSTR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) TIM2RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM2RST)}}
}

func (p *Periph) TIM3RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM3RST)}}
}

func (p *Periph) TIM4RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM4RST)}}
}

func (p *Periph) TIM5RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM5RST)}}
}

func (p *Periph) TIM6RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM6RST)}}
}

func (p *Periph) TIM7RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM7RST)}}
}

func (p *Periph) TIM12RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM12RST)}}
}

func (p *Periph) TIM13RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM13RST)}}
}

func (p *Periph) TIM14RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM14RST)}}
}

func (p *Periph) WWDGRST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(WWDGRST)}}
}

func (p *Periph) SPI2RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(SPI2RST)}}
}

func (p *Periph) SPI3RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(SPI3RST)}}
}

func (p *Periph) UART2RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(UART2RST)}}
}

func (p *Periph) USART3RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(USART3RST)}}
}

func (p *Periph) I2C1RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(I2C1RST)}}
}

func (p *Periph) I2C2RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(I2C2RST)}}
}

func (p *Periph) I2C3RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(I2C3RST)}}
}

func (p *Periph) I2C4RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(I2C4RST)}}
}

func (p *Periph) CAN1RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(CAN1RST)}}
}

func (p *Periph) CAN2RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(CAN2RST)}}
}

func (p *Periph) PWRRST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(PWRRST)}}
}

type APB2RSTR uint32

type RAPB2RSTR struct{ mmio.U32 }

func (r *RAPB2RSTR) LoadBits(mask APB2RSTR) APB2RSTR { return APB2RSTR(r.U32.LoadBits(uint32(mask))) }
func (r *RAPB2RSTR) StoreBits(mask, b APB2RSTR)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAPB2RSTR) SetBits(mask APB2RSTR)           { r.U32.SetBits(uint32(mask)) }
func (r *RAPB2RSTR) ClearBits(mask APB2RSTR)         { r.U32.ClearBits(uint32(mask)) }
func (r *RAPB2RSTR) Load() APB2RSTR                  { return APB2RSTR(r.U32.Load()) }
func (r *RAPB2RSTR) Store(b APB2RSTR)                { r.U32.Store(uint32(b)) }

type RMAPB2RSTR struct{ mmio.UM32 }

func (rm RMAPB2RSTR) Load() APB2RSTR   { return APB2RSTR(rm.UM32.Load()) }
func (rm RMAPB2RSTR) Store(b APB2RSTR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) TIM1RST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM1RST)}}
}

func (p *Periph) TIM8RST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM8RST)}}
}

func (p *Periph) USART1RST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(USART1RST)}}
}

func (p *Periph) USART6RST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(USART6RST)}}
}

func (p *Periph) ADCRST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(ADCRST)}}
}

func (p *Periph) SDIORST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(SDIORST)}}
}

func (p *Periph) SPI1RST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(SPI1RST)}}
}

func (p *Periph) SYSCFGRST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(SYSCFGRST)}}
}

func (p *Periph) TIM9RST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM9RST)}}
}

func (p *Periph) TIM10RST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM10RST)}}
}

func (p *Periph) TIM11RST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM11RST)}}
}

func (p *Periph) DFSDMRST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(DFSDMRST)}}
}

type AHB1ENR uint32

type RAHB1ENR struct{ mmio.U32 }

func (r *RAHB1ENR) LoadBits(mask AHB1ENR) AHB1ENR { return AHB1ENR(r.U32.LoadBits(uint32(mask))) }
func (r *RAHB1ENR) StoreBits(mask, b AHB1ENR)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAHB1ENR) SetBits(mask AHB1ENR)          { r.U32.SetBits(uint32(mask)) }
func (r *RAHB1ENR) ClearBits(mask AHB1ENR)        { r.U32.ClearBits(uint32(mask)) }
func (r *RAHB1ENR) Load() AHB1ENR                 { return AHB1ENR(r.U32.Load()) }
func (r *RAHB1ENR) Store(b AHB1ENR)               { r.U32.Store(uint32(b)) }

type RMAHB1ENR struct{ mmio.UM32 }

func (rm RMAHB1ENR) Load() AHB1ENR   { return AHB1ENR(rm.UM32.Load()) }
func (rm RMAHB1ENR) Store(b AHB1ENR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) GPIOAEN() RMAHB1ENR {
	return RMAHB1ENR{mmio.UM32{&p.AHB1ENR.U32, uint32(GPIOAEN)}}
}

func (p *Periph) GPIOBEN() RMAHB1ENR {
	return RMAHB1ENR{mmio.UM32{&p.AHB1ENR.U32, uint32(GPIOBEN)}}
}

func (p *Periph) GPIOCEN() RMAHB1ENR {
	return RMAHB1ENR{mmio.UM32{&p.AHB1ENR.U32, uint32(GPIOCEN)}}
}

func (p *Periph) GPIODEN() RMAHB1ENR {
	return RMAHB1ENR{mmio.UM32{&p.AHB1ENR.U32, uint32(GPIODEN)}}
}

func (p *Periph) GPIOEEN() RMAHB1ENR {
	return RMAHB1ENR{mmio.UM32{&p.AHB1ENR.U32, uint32(GPIOEEN)}}
}

func (p *Periph) GPIOFEN() RMAHB1ENR {
	return RMAHB1ENR{mmio.UM32{&p.AHB1ENR.U32, uint32(GPIOFEN)}}
}

func (p *Periph) GPIOGEN() RMAHB1ENR {
	return RMAHB1ENR{mmio.UM32{&p.AHB1ENR.U32, uint32(GPIOGEN)}}
}

func (p *Periph) GPIOHEN() RMAHB1ENR {
	return RMAHB1ENR{mmio.UM32{&p.AHB1ENR.U32, uint32(GPIOHEN)}}
}

func (p *Periph) CRCEN() RMAHB1ENR {
	return RMAHB1ENR{mmio.UM32{&p.AHB1ENR.U32, uint32(CRCEN)}}
}

func (p *Periph) DMA1EN() RMAHB1ENR {
	return RMAHB1ENR{mmio.UM32{&p.AHB1ENR.U32, uint32(DMA1EN)}}
}

func (p *Periph) DMA2EN() RMAHB1ENR {
	return RMAHB1ENR{mmio.UM32{&p.AHB1ENR.U32, uint32(DMA2EN)}}
}

type AHB2ENR uint32

type RAHB2ENR struct{ mmio.U32 }

func (r *RAHB2ENR) LoadBits(mask AHB2ENR) AHB2ENR { return AHB2ENR(r.U32.LoadBits(uint32(mask))) }
func (r *RAHB2ENR) StoreBits(mask, b AHB2ENR)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAHB2ENR) SetBits(mask AHB2ENR)          { r.U32.SetBits(uint32(mask)) }
func (r *RAHB2ENR) ClearBits(mask AHB2ENR)        { r.U32.ClearBits(uint32(mask)) }
func (r *RAHB2ENR) Load() AHB2ENR                 { return AHB2ENR(r.U32.Load()) }
func (r *RAHB2ENR) Store(b AHB2ENR)               { r.U32.Store(uint32(b)) }

type RMAHB2ENR struct{ mmio.UM32 }

func (rm RMAHB2ENR) Load() AHB2ENR   { return AHB2ENR(rm.UM32.Load()) }
func (rm RMAHB2ENR) Store(b AHB2ENR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) RNGEN() RMAHB2ENR {
	return RMAHB2ENR{mmio.UM32{&p.AHB2ENR.U32, uint32(RNGEN)}}
}

func (p *Periph) OTGFSEN() RMAHB2ENR {
	return RMAHB2ENR{mmio.UM32{&p.AHB2ENR.U32, uint32(OTGFSEN)}}
}

type APB1ENR uint32

type RAPB1ENR struct{ mmio.U32 }

func (r *RAPB1ENR) LoadBits(mask APB1ENR) APB1ENR { return APB1ENR(r.U32.LoadBits(uint32(mask))) }
func (r *RAPB1ENR) StoreBits(mask, b APB1ENR)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAPB1ENR) SetBits(mask APB1ENR)          { r.U32.SetBits(uint32(mask)) }
func (r *RAPB1ENR) ClearBits(mask APB1ENR)        { r.U32.ClearBits(uint32(mask)) }
func (r *RAPB1ENR) Load() APB1ENR                 { return APB1ENR(r.U32.Load()) }
func (r *RAPB1ENR) Store(b APB1ENR)               { r.U32.Store(uint32(b)) }

type RMAPB1ENR struct{ mmio.UM32 }

func (rm RMAPB1ENR) Load() APB1ENR   { return APB1ENR(rm.UM32.Load()) }
func (rm RMAPB1ENR) Store(b APB1ENR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) TIM2EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(TIM2EN)}}
}

func (p *Periph) TIM3EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(TIM3EN)}}
}

func (p *Periph) TIM4EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(TIM4EN)}}
}

func (p *Periph) TIM5EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(TIM5EN)}}
}

func (p *Periph) TIM6EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(TIM6EN)}}
}

func (p *Periph) TIM7EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(TIM7EN)}}
}

func (p *Periph) TIM12EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(TIM12EN)}}
}

func (p *Periph) TIM13EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(TIM13EN)}}
}

func (p *Periph) TIM14EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(TIM14EN)}}
}

func (p *Periph) WWDGEN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(WWDGEN)}}
}

func (p *Periph) SPI2EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(SPI2EN)}}
}

func (p *Periph) SPI3EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(SPI3EN)}}
}

func (p *Periph) USART2EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(USART2EN)}}
}

func (p *Periph) USART3EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(USART3EN)}}
}

func (p *Periph) I2C1EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(I2C1EN)}}
}

func (p *Periph) I2C2EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(I2C2EN)}}
}

func (p *Periph) I2C3EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(I2C3EN)}}
}

func (p *Periph) I2C4EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(I2C4EN)}}
}

func (p *Periph) CAN1EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(CAN1EN)}}
}

func (p *Periph) CAN2EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(CAN2EN)}}
}

func (p *Periph) PWREN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(PWREN)}}
}

type APB2ENR uint32

type RAPB2ENR struct{ mmio.U32 }

func (r *RAPB2ENR) LoadBits(mask APB2ENR) APB2ENR { return APB2ENR(r.U32.LoadBits(uint32(mask))) }
func (r *RAPB2ENR) StoreBits(mask, b APB2ENR)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAPB2ENR) SetBits(mask APB2ENR)          { r.U32.SetBits(uint32(mask)) }
func (r *RAPB2ENR) ClearBits(mask APB2ENR)        { r.U32.ClearBits(uint32(mask)) }
func (r *RAPB2ENR) Load() APB2ENR                 { return APB2ENR(r.U32.Load()) }
func (r *RAPB2ENR) Store(b APB2ENR)               { r.U32.Store(uint32(b)) }

type RMAPB2ENR struct{ mmio.UM32 }

func (rm RMAPB2ENR) Load() APB2ENR   { return APB2ENR(rm.UM32.Load()) }
func (rm RMAPB2ENR) Store(b APB2ENR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) TIM1EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(TIM1EN)}}
}

func (p *Periph) TIM8EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(TIM8EN)}}
}

func (p *Periph) USART1EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(USART1EN)}}
}

func (p *Periph) USART6EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(USART6EN)}}
}

func (p *Periph) ADC1EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(ADC1EN)}}
}

func (p *Periph) SDIOEN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(SDIOEN)}}
}

func (p *Periph) SPI1EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(SPI1EN)}}
}

func (p *Periph) SPI4EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(SPI4EN)}}
}

func (p *Periph) SYSCFGEN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(SYSCFGEN)}}
}

func (p *Periph) TIM9EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(TIM9EN)}}
}

func (p *Periph) TIM10EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(TIM10EN)}}
}

func (p *Periph) TIM11EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(TIM11EN)}}
}

func (p *Periph) DFSDMEN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(DFSDMEN)}}
}

type AHB1LPENR uint32

type RAHB1LPENR struct{ mmio.U32 }

func (r *RAHB1LPENR) LoadBits(mask AHB1LPENR) AHB1LPENR {
	return AHB1LPENR(r.U32.LoadBits(uint32(mask)))
}
func (r *RAHB1LPENR) StoreBits(mask, b AHB1LPENR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAHB1LPENR) SetBits(mask AHB1LPENR)      { r.U32.SetBits(uint32(mask)) }
func (r *RAHB1LPENR) ClearBits(mask AHB1LPENR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAHB1LPENR) Load() AHB1LPENR             { return AHB1LPENR(r.U32.Load()) }
func (r *RAHB1LPENR) Store(b AHB1LPENR)           { r.U32.Store(uint32(b)) }

type RMAHB1LPENR struct{ mmio.UM32 }

func (rm RMAHB1LPENR) Load() AHB1LPENR   { return AHB1LPENR(rm.UM32.Load()) }
func (rm RMAHB1LPENR) Store(b AHB1LPENR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) GPIOALPEN() RMAHB1LPENR {
	return RMAHB1LPENR{mmio.UM32{&p.AHB1LPENR.U32, uint32(GPIOALPEN)}}
}

func (p *Periph) GPIOBLPEN() RMAHB1LPENR {
	return RMAHB1LPENR{mmio.UM32{&p.AHB1LPENR.U32, uint32(GPIOBLPEN)}}
}

func (p *Periph) GPIOCLPEN() RMAHB1LPENR {
	return RMAHB1LPENR{mmio.UM32{&p.AHB1LPENR.U32, uint32(GPIOCLPEN)}}
}

func (p *Periph) GPIODLPEN() RMAHB1LPENR {
	return RMAHB1LPENR{mmio.UM32{&p.AHB1LPENR.U32, uint32(GPIODLPEN)}}
}

func (p *Periph) GPIOELPEN() RMAHB1LPENR {
	return RMAHB1LPENR{mmio.UM32{&p.AHB1LPENR.U32, uint32(GPIOELPEN)}}
}

func (p *Periph) GPIOFLPEN() RMAHB1LPENR {
	return RMAHB1LPENR{mmio.UM32{&p.AHB1LPENR.U32, uint32(GPIOFLPEN)}}
}

func (p *Periph) GPIOGLPEN() RMAHB1LPENR {
	return RMAHB1LPENR{mmio.UM32{&p.AHB1LPENR.U32, uint32(GPIOGLPEN)}}
}

func (p *Periph) GPIOHLPEN() RMAHB1LPENR {
	return RMAHB1LPENR{mmio.UM32{&p.AHB1LPENR.U32, uint32(GPIOHLPEN)}}
}

func (p *Periph) CRCLPEN() RMAHB1LPENR {
	return RMAHB1LPENR{mmio.UM32{&p.AHB1LPENR.U32, uint32(CRCLPEN)}}
}

func (p *Periph) FLITFLPEN() RMAHB1LPENR {
	return RMAHB1LPENR{mmio.UM32{&p.AHB1LPENR.U32, uint32(FLITFLPEN)}}
}

func (p *Periph) SRAM1LPEN() RMAHB1LPENR {
	return RMAHB1LPENR{mmio.UM32{&p.AHB1LPENR.U32, uint32(SRAM1LPEN)}}
}

func (p *Periph) DMA1LPEN() RMAHB1LPENR {
	return RMAHB1LPENR{mmio.UM32{&p.AHB1LPENR.U32, uint32(DMA1LPEN)}}
}

func (p *Periph) DMA2LPEN() RMAHB1LPENR {
	return RMAHB1LPENR{mmio.UM32{&p.AHB1LPENR.U32, uint32(DMA2LPEN)}}
}

type AHB2LPENR uint32

type RAHB2LPENR struct{ mmio.U32 }

func (r *RAHB2LPENR) LoadBits(mask AHB2LPENR) AHB2LPENR {
	return AHB2LPENR(r.U32.LoadBits(uint32(mask)))
}
func (r *RAHB2LPENR) StoreBits(mask, b AHB2LPENR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAHB2LPENR) SetBits(mask AHB2LPENR)      { r.U32.SetBits(uint32(mask)) }
func (r *RAHB2LPENR) ClearBits(mask AHB2LPENR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAHB2LPENR) Load() AHB2LPENR             { return AHB2LPENR(r.U32.Load()) }
func (r *RAHB2LPENR) Store(b AHB2LPENR)           { r.U32.Store(uint32(b)) }

type RMAHB2LPENR struct{ mmio.UM32 }

func (rm RMAHB2LPENR) Load() AHB2LPENR   { return AHB2LPENR(rm.UM32.Load()) }
func (rm RMAHB2LPENR) Store(b AHB2LPENR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) RNGLPEN() RMAHB2LPENR {
	return RMAHB2LPENR{mmio.UM32{&p.AHB2LPENR.U32, uint32(RNGLPEN)}}
}

func (p *Periph) OTGFSLPEN() RMAHB2LPENR {
	return RMAHB2LPENR{mmio.UM32{&p.AHB2LPENR.U32, uint32(OTGFSLPEN)}}
}

type APB1LPENR uint32

type RAPB1LPENR struct{ mmio.U32 }

func (r *RAPB1LPENR) LoadBits(mask APB1LPENR) APB1LPENR {
	return APB1LPENR(r.U32.LoadBits(uint32(mask)))
}
func (r *RAPB1LPENR) StoreBits(mask, b APB1LPENR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAPB1LPENR) SetBits(mask APB1LPENR)      { r.U32.SetBits(uint32(mask)) }
func (r *RAPB1LPENR) ClearBits(mask APB1LPENR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAPB1LPENR) Load() APB1LPENR             { return APB1LPENR(r.U32.Load()) }
func (r *RAPB1LPENR) Store(b APB1LPENR)           { r.U32.Store(uint32(b)) }

type RMAPB1LPENR struct{ mmio.UM32 }

func (rm RMAPB1LPENR) Load() APB1LPENR   { return APB1LPENR(rm.UM32.Load()) }
func (rm RMAPB1LPENR) Store(b APB1LPENR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) TIM2LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(TIM2LPEN)}}
}

func (p *Periph) TIM3LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(TIM3LPEN)}}
}

func (p *Periph) TIM4LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(TIM4LPEN)}}
}

func (p *Periph) TIM5LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(TIM5LPEN)}}
}

func (p *Periph) TIM6LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(TIM6LPEN)}}
}

func (p *Periph) TIM7LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(TIM7LPEN)}}
}

func (p *Periph) TIM12LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(TIM12LPEN)}}
}

func (p *Periph) TIM13LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(TIM13LPEN)}}
}

func (p *Periph) TIM14LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(TIM14LPEN)}}
}

func (p *Periph) WWDGLPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(WWDGLPEN)}}
}

func (p *Periph) SPI2LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(SPI2LPEN)}}
}

func (p *Periph) SPI3LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(SPI3LPEN)}}
}

func (p *Periph) USART2LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(USART2LPEN)}}
}

func (p *Periph) USART3LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(USART3LPEN)}}
}

func (p *Periph) I2C1LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(I2C1LPEN)}}
}

func (p *Periph) I2C2LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(I2C2LPEN)}}
}

func (p *Periph) I2C3LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(I2C3LPEN)}}
}

func (p *Periph) I2C4LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(I2C4LPEN)}}
}

func (p *Periph) CAN1LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(CAN1LPEN)}}
}

func (p *Periph) CAN2LPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(CAN2LPEN)}}
}

func (p *Periph) PWRLPEN() RMAPB1LPENR {
	return RMAPB1LPENR{mmio.UM32{&p.APB1LPENR.U32, uint32(PWRLPEN)}}
}

type APB2LPENR uint32

type RAPB2LPENR struct{ mmio.U32 }

func (r *RAPB2LPENR) LoadBits(mask APB2LPENR) APB2LPENR {
	return APB2LPENR(r.U32.LoadBits(uint32(mask)))
}
func (r *RAPB2LPENR) StoreBits(mask, b APB2LPENR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAPB2LPENR) SetBits(mask APB2LPENR)      { r.U32.SetBits(uint32(mask)) }
func (r *RAPB2LPENR) ClearBits(mask APB2LPENR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAPB2LPENR) Load() APB2LPENR             { return APB2LPENR(r.U32.Load()) }
func (r *RAPB2LPENR) Store(b APB2LPENR)           { r.U32.Store(uint32(b)) }

type RMAPB2LPENR struct{ mmio.UM32 }

func (rm RMAPB2LPENR) Load() APB2LPENR   { return APB2LPENR(rm.UM32.Load()) }
func (rm RMAPB2LPENR) Store(b APB2LPENR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) TIM1LPEN() RMAPB2LPENR {
	return RMAPB2LPENR{mmio.UM32{&p.APB2LPENR.U32, uint32(TIM1LPEN)}}
}

func (p *Periph) TIM8LPEN() RMAPB2LPENR {
	return RMAPB2LPENR{mmio.UM32{&p.APB2LPENR.U32, uint32(TIM8LPEN)}}
}

func (p *Periph) USART1LPEN() RMAPB2LPENR {
	return RMAPB2LPENR{mmio.UM32{&p.APB2LPENR.U32, uint32(USART1LPEN)}}
}

func (p *Periph) USART6LPEN() RMAPB2LPENR {
	return RMAPB2LPENR{mmio.UM32{&p.APB2LPENR.U32, uint32(USART6LPEN)}}
}

func (p *Periph) ADC1LPEN() RMAPB2LPENR {
	return RMAPB2LPENR{mmio.UM32{&p.APB2LPENR.U32, uint32(ADC1LPEN)}}
}

func (p *Periph) SDIOLPEN() RMAPB2LPENR {
	return RMAPB2LPENR{mmio.UM32{&p.APB2LPENR.U32, uint32(SDIOLPEN)}}
}

func (p *Periph) SPI1LPEN() RMAPB2LPENR {
	return RMAPB2LPENR{mmio.UM32{&p.APB2LPENR.U32, uint32(SPI1LPEN)}}
}

func (p *Periph) SPI4LPEN() RMAPB2LPENR {
	return RMAPB2LPENR{mmio.UM32{&p.APB2LPENR.U32, uint32(SPI4LPEN)}}
}

func (p *Periph) SYSCFGLPEN() RMAPB2LPENR {
	return RMAPB2LPENR{mmio.UM32{&p.APB2LPENR.U32, uint32(SYSCFGLPEN)}}
}

func (p *Periph) TIM9LPEN() RMAPB2LPENR {
	return RMAPB2LPENR{mmio.UM32{&p.APB2LPENR.U32, uint32(TIM9LPEN)}}
}

func (p *Periph) TIM10LPEN() RMAPB2LPENR {
	return RMAPB2LPENR{mmio.UM32{&p.APB2LPENR.U32, uint32(TIM10LPEN)}}
}

func (p *Periph) TIM11LPEN() RMAPB2LPENR {
	return RMAPB2LPENR{mmio.UM32{&p.APB2LPENR.U32, uint32(TIM11LPEN)}}
}

func (p *Periph) DFSDMLPEN() RMAPB2LPENR {
	return RMAPB2LPENR{mmio.UM32{&p.APB2LPENR.U32, uint32(DFSDMLPEN)}}
}

type BDCR uint32

type RBDCR struct{ mmio.U32 }

func (r *RBDCR) LoadBits(mask BDCR) BDCR { return BDCR(r.U32.LoadBits(uint32(mask))) }
func (r *RBDCR) StoreBits(mask, b BDCR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBDCR) SetBits(mask BDCR)       { r.U32.SetBits(uint32(mask)) }
func (r *RBDCR) ClearBits(mask BDCR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RBDCR) Load() BDCR              { return BDCR(r.U32.Load()) }
func (r *RBDCR) Store(b BDCR)            { r.U32.Store(uint32(b)) }

type RMBDCR struct{ mmio.UM32 }

func (rm RMBDCR) Load() BDCR   { return BDCR(rm.UM32.Load()) }
func (rm RMBDCR) Store(b BDCR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) LSEON() RMBDCR {
	return RMBDCR{mmio.UM32{&p.BDCR.U32, uint32(LSEON)}}
}

func (p *Periph) LSERDY() RMBDCR {
	return RMBDCR{mmio.UM32{&p.BDCR.U32, uint32(LSERDY)}}
}

func (p *Periph) LSEBYP() RMBDCR {
	return RMBDCR{mmio.UM32{&p.BDCR.U32, uint32(LSEBYP)}}
}

func (p *Periph) RTCSEL0() RMBDCR {
	return RMBDCR{mmio.UM32{&p.BDCR.U32, uint32(RTCSEL0)}}
}

func (p *Periph) RTCSEL1() RMBDCR {
	return RMBDCR{mmio.UM32{&p.BDCR.U32, uint32(RTCSEL1)}}
}

func (p *Periph) RTCEN() RMBDCR {
	return RMBDCR{mmio.UM32{&p.BDCR.U32, uint32(RTCEN)}}
}

func (p *Periph) BDRST() RMBDCR {
	return RMBDCR{mmio.UM32{&p.BDCR.U32, uint32(BDRST)}}
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

func (p *Periph) LSION() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(LSION)}}
}

func (p *Periph) LSIRDY() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(LSIRDY)}}
}

func (p *Periph) RMVF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(RMVF)}}
}

func (p *Periph) BORRSTF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(BORRSTF)}}
}

func (p *Periph) PADRSTF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(PADRSTF)}}
}

func (p *Periph) PORRSTF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(PORRSTF)}}
}

func (p *Periph) SFTRSTF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(SFTRSTF)}}
}

func (p *Periph) WDGRSTF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(WDGRSTF)}}
}

func (p *Periph) WWDGRSTF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(WWDGRSTF)}}
}

func (p *Periph) LPWRRSTF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(LPWRRSTF)}}
}

type SSCGR uint32

type RSSCGR struct{ mmio.U32 }

func (r *RSSCGR) LoadBits(mask SSCGR) SSCGR { return SSCGR(r.U32.LoadBits(uint32(mask))) }
func (r *RSSCGR) StoreBits(mask, b SSCGR)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSSCGR) SetBits(mask SSCGR)        { r.U32.SetBits(uint32(mask)) }
func (r *RSSCGR) ClearBits(mask SSCGR)      { r.U32.ClearBits(uint32(mask)) }
func (r *RSSCGR) Load() SSCGR               { return SSCGR(r.U32.Load()) }
func (r *RSSCGR) Store(b SSCGR)             { r.U32.Store(uint32(b)) }

type RMSSCGR struct{ mmio.UM32 }

func (rm RMSSCGR) Load() SSCGR   { return SSCGR(rm.UM32.Load()) }
func (rm RMSSCGR) Store(b SSCGR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) MODPER() RMSSCGR {
	return RMSSCGR{mmio.UM32{&p.SSCGR.U32, uint32(MODPER)}}
}

func (p *Periph) INCSTEP() RMSSCGR {
	return RMSSCGR{mmio.UM32{&p.SSCGR.U32, uint32(INCSTEP)}}
}

func (p *Periph) SPREADSEL() RMSSCGR {
	return RMSSCGR{mmio.UM32{&p.SSCGR.U32, uint32(SPREADSEL)}}
}

func (p *Periph) SSCGEN() RMSSCGR {
	return RMSSCGR{mmio.UM32{&p.SSCGR.U32, uint32(SSCGEN)}}
}

type PLLI2SCFGR uint32

type RPLLI2SCFGR struct{ mmio.U32 }

func (r *RPLLI2SCFGR) LoadBits(mask PLLI2SCFGR) PLLI2SCFGR {
	return PLLI2SCFGR(r.U32.LoadBits(uint32(mask)))
}
func (r *RPLLI2SCFGR) StoreBits(mask, b PLLI2SCFGR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPLLI2SCFGR) SetBits(mask PLLI2SCFGR)      { r.U32.SetBits(uint32(mask)) }
func (r *RPLLI2SCFGR) ClearBits(mask PLLI2SCFGR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPLLI2SCFGR) Load() PLLI2SCFGR             { return PLLI2SCFGR(r.U32.Load()) }
func (r *RPLLI2SCFGR) Store(b PLLI2SCFGR)           { r.U32.Store(uint32(b)) }

type RMPLLI2SCFGR struct{ mmio.UM32 }

func (rm RMPLLI2SCFGR) Load() PLLI2SCFGR   { return PLLI2SCFGR(rm.UM32.Load()) }
func (rm RMPLLI2SCFGR) Store(b PLLI2SCFGR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) PLLI2SNx() RMPLLI2SCFGR {
	return RMPLLI2SCFGR{mmio.UM32{&p.PLLI2SCFGR.U32, uint32(PLLI2SNx)}}
}

func (p *Periph) PLLI2SRx() RMPLLI2SCFGR {
	return RMPLLI2SCFGR{mmio.UM32{&p.PLLI2SCFGR.U32, uint32(PLLI2SRx)}}
}
