// DO NOT EDIT THIS FILE. GENERATED BY xgen.

// +build stm32f407

package gpio

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	MODER   RMODER
	OTYPER  ROTYPER
	OSPEEDR ROSPEEDR
	PUPDR   RPUPDR
	IDR     RIDR
	ODR     RODR
	BSRR    RBSRR
	LCKR    RLCKR
	AFRL    RAFRL
	AFRH    RAFRH
}

func GPIOA() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.GPIOA_BASE))) }
func GPIOB() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.GPIOB_BASE))) }
func GPIOC() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.GPIOC_BASE))) }
func GPIOD() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.GPIOD_BASE))) }
func GPIOE() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.GPIOE_BASE))) }
func GPIOF() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.GPIOF_BASE))) }
func GPIOG() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.GPIOG_BASE))) }
func GPIOH() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.GPIOH_BASE))) }
func GPIOI() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.GPIOI_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

func (p *Periph) Bus() bus.Bus {
	return bus.AHB1
}

type MODER uint32

type RMODER struct{ mmio.U32 }

func (r *RMODER) LoadBits(mask MODER) MODER { return MODER(r.U32.LoadBits(uint32(mask))) }
func (r *RMODER) StoreBits(mask, b MODER)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RMODER) SetBits(mask MODER)        { r.U32.SetBits(uint32(mask)) }
func (r *RMODER) ClearBits(mask MODER)      { r.U32.ClearBits(uint32(mask)) }
func (r *RMODER) Load() MODER               { return MODER(r.U32.Load()) }
func (r *RMODER) Store(b MODER)             { r.U32.Store(uint32(b)) }

type RMMODER struct{ mmio.UM32 }

func (rm RMMODER) Load() MODER   { return MODER(rm.UM32.Load()) }
func (rm RMMODER) Store(b MODER) { rm.UM32.Store(uint32(b)) }

func (p *Periph) MODER0() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER0)}}
}

func (p *Periph) MODER1() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER1)}}
}

func (p *Periph) MODER2() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER2)}}
}

func (p *Periph) MODER3() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER3)}}
}

func (p *Periph) MODER4() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER4)}}
}

func (p *Periph) MODER5() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER5)}}
}

func (p *Periph) MODER6() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER6)}}
}

func (p *Periph) MODER7() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER7)}}
}

func (p *Periph) MODER8() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER8)}}
}

func (p *Periph) MODER9() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER9)}}
}

func (p *Periph) MODER10() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER10)}}
}

func (p *Periph) MODER11() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER11)}}
}

func (p *Periph) MODER12() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER12)}}
}

func (p *Periph) MODER13() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER13)}}
}

func (p *Periph) MODER14() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER14)}}
}

func (p *Periph) MODER15() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER15)}}
}

type OTYPER uint32

type ROTYPER struct{ mmio.U32 }

func (r *ROTYPER) LoadBits(mask OTYPER) OTYPER { return OTYPER(r.U32.LoadBits(uint32(mask))) }
func (r *ROTYPER) StoreBits(mask, b OTYPER)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROTYPER) SetBits(mask OTYPER)         { r.U32.SetBits(uint32(mask)) }
func (r *ROTYPER) ClearBits(mask OTYPER)       { r.U32.ClearBits(uint32(mask)) }
func (r *ROTYPER) Load() OTYPER                { return OTYPER(r.U32.Load()) }
func (r *ROTYPER) Store(b OTYPER)              { r.U32.Store(uint32(b)) }

type RMOTYPER struct{ mmio.UM32 }

func (rm RMOTYPER) Load() OTYPER   { return OTYPER(rm.UM32.Load()) }
func (rm RMOTYPER) Store(b OTYPER) { rm.UM32.Store(uint32(b)) }

func (p *Periph) OT0() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT0)}}
}

func (p *Periph) OT1() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT1)}}
}

func (p *Periph) OT2() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT2)}}
}

func (p *Periph) OT3() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT3)}}
}

func (p *Periph) OT4() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT4)}}
}

func (p *Periph) OT5() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT5)}}
}

func (p *Periph) OT6() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT6)}}
}

func (p *Periph) OT7() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT7)}}
}

func (p *Periph) OT8() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT8)}}
}

func (p *Periph) OT9() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT9)}}
}

func (p *Periph) OT10() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT10)}}
}

func (p *Periph) OT11() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT11)}}
}

func (p *Periph) OT12() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT12)}}
}

func (p *Periph) OT13() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT13)}}
}

func (p *Periph) OT14() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT14)}}
}

func (p *Periph) OT15() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT15)}}
}

type OSPEEDR uint32

type ROSPEEDR struct{ mmio.U32 }

func (r *ROSPEEDR) LoadBits(mask OSPEEDR) OSPEEDR { return OSPEEDR(r.U32.LoadBits(uint32(mask))) }
func (r *ROSPEEDR) StoreBits(mask, b OSPEEDR)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROSPEEDR) SetBits(mask OSPEEDR)          { r.U32.SetBits(uint32(mask)) }
func (r *ROSPEEDR) ClearBits(mask OSPEEDR)        { r.U32.ClearBits(uint32(mask)) }
func (r *ROSPEEDR) Load() OSPEEDR                 { return OSPEEDR(r.U32.Load()) }
func (r *ROSPEEDR) Store(b OSPEEDR)               { r.U32.Store(uint32(b)) }

type RMOSPEEDR struct{ mmio.UM32 }

func (rm RMOSPEEDR) Load() OSPEEDR   { return OSPEEDR(rm.UM32.Load()) }
func (rm RMOSPEEDR) Store(b OSPEEDR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) OSPEEDR0() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR0)}}
}

func (p *Periph) OSPEEDR1() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR1)}}
}

func (p *Periph) OSPEEDR2() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR2)}}
}

func (p *Periph) OSPEEDR3() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR3)}}
}

func (p *Periph) OSPEEDR4() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR4)}}
}

func (p *Periph) OSPEEDR5() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR5)}}
}

func (p *Periph) OSPEEDR6() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR6)}}
}

func (p *Periph) OSPEEDR7() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR7)}}
}

func (p *Periph) OSPEEDR8() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR8)}}
}

func (p *Periph) OSPEEDR9() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR9)}}
}

func (p *Periph) OSPEEDR10() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR10)}}
}

func (p *Periph) OSPEEDR11() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR11)}}
}

func (p *Periph) OSPEEDR12() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR12)}}
}

func (p *Periph) OSPEEDR13() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR13)}}
}

func (p *Periph) OSPEEDR14() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR14)}}
}

func (p *Periph) OSPEEDR15() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR15)}}
}

type PUPDR uint32

type RPUPDR struct{ mmio.U32 }

func (r *RPUPDR) LoadBits(mask PUPDR) PUPDR { return PUPDR(r.U32.LoadBits(uint32(mask))) }
func (r *RPUPDR) StoreBits(mask, b PUPDR)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPUPDR) SetBits(mask PUPDR)        { r.U32.SetBits(uint32(mask)) }
func (r *RPUPDR) ClearBits(mask PUPDR)      { r.U32.ClearBits(uint32(mask)) }
func (r *RPUPDR) Load() PUPDR               { return PUPDR(r.U32.Load()) }
func (r *RPUPDR) Store(b PUPDR)             { r.U32.Store(uint32(b)) }

type RMPUPDR struct{ mmio.UM32 }

func (rm RMPUPDR) Load() PUPDR   { return PUPDR(rm.UM32.Load()) }
func (rm RMPUPDR) Store(b PUPDR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) PUPDR0() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR0)}}
}

func (p *Periph) PUPDR1() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR1)}}
}

func (p *Periph) PUPDR2() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR2)}}
}

func (p *Periph) PUPDR3() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR3)}}
}

func (p *Periph) PUPDR4() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR4)}}
}

func (p *Periph) PUPDR5() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR5)}}
}

func (p *Periph) PUPDR6() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR6)}}
}

func (p *Periph) PUPDR7() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR7)}}
}

func (p *Periph) PUPDR8() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR8)}}
}

func (p *Periph) PUPDR9() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR9)}}
}

func (p *Periph) PUPDR10() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR10)}}
}

func (p *Periph) PUPDR11() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR11)}}
}

func (p *Periph) PUPDR12() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR12)}}
}

func (p *Periph) PUPDR13() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR13)}}
}

func (p *Periph) PUPDR14() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR14)}}
}

func (p *Periph) PUPDR15() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR15)}}
}

type IDR uint32

type RIDR struct{ mmio.U32 }

func (r *RIDR) LoadBits(mask IDR) IDR { return IDR(r.U32.LoadBits(uint32(mask))) }
func (r *RIDR) StoreBits(mask, b IDR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIDR) SetBits(mask IDR)      { r.U32.SetBits(uint32(mask)) }
func (r *RIDR) ClearBits(mask IDR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RIDR) Load() IDR             { return IDR(r.U32.Load()) }
func (r *RIDR) Store(b IDR)           { r.U32.Store(uint32(b)) }

type RMIDR struct{ mmio.UM32 }

func (rm RMIDR) Load() IDR   { return IDR(rm.UM32.Load()) }
func (rm RMIDR) Store(b IDR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) IDR0() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR0)}}
}

func (p *Periph) IDR1() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR1)}}
}

func (p *Periph) IDR2() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR2)}}
}

func (p *Periph) IDR3() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR3)}}
}

func (p *Periph) IDR4() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR4)}}
}

func (p *Periph) IDR5() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR5)}}
}

func (p *Periph) IDR6() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR6)}}
}

func (p *Periph) IDR7() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR7)}}
}

func (p *Periph) IDR8() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR8)}}
}

func (p *Periph) IDR9() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR9)}}
}

func (p *Periph) IDR10() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR10)}}
}

func (p *Periph) IDR11() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR11)}}
}

func (p *Periph) IDR12() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR12)}}
}

func (p *Periph) IDR13() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR13)}}
}

func (p *Periph) IDR14() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR14)}}
}

func (p *Periph) IDR15() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(IDR15)}}
}

type ODR uint32

type RODR struct{ mmio.U32 }

func (r *RODR) LoadBits(mask ODR) ODR { return ODR(r.U32.LoadBits(uint32(mask))) }
func (r *RODR) StoreBits(mask, b ODR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RODR) SetBits(mask ODR)      { r.U32.SetBits(uint32(mask)) }
func (r *RODR) ClearBits(mask ODR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RODR) Load() ODR             { return ODR(r.U32.Load()) }
func (r *RODR) Store(b ODR)           { r.U32.Store(uint32(b)) }

type RMODR struct{ mmio.UM32 }

func (rm RMODR) Load() ODR   { return ODR(rm.UM32.Load()) }
func (rm RMODR) Store(b ODR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) ODR0() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR0)}}
}

func (p *Periph) ODR1() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR1)}}
}

func (p *Periph) ODR2() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR2)}}
}

func (p *Periph) ODR3() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR3)}}
}

func (p *Periph) ODR4() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR4)}}
}

func (p *Periph) ODR5() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR5)}}
}

func (p *Periph) ODR6() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR6)}}
}

func (p *Periph) ODR7() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR7)}}
}

func (p *Periph) ODR8() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR8)}}
}

func (p *Periph) ODR9() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR9)}}
}

func (p *Periph) ODR10() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR10)}}
}

func (p *Periph) ODR11() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR11)}}
}

func (p *Periph) ODR12() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR12)}}
}

func (p *Periph) ODR13() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR13)}}
}

func (p *Periph) ODR14() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR14)}}
}

func (p *Periph) ODR15() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(ODR15)}}
}

type BSRR uint32

type RBSRR struct{ mmio.U32 }

func (r *RBSRR) LoadBits(mask BSRR) BSRR { return BSRR(r.U32.LoadBits(uint32(mask))) }
func (r *RBSRR) StoreBits(mask, b BSRR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBSRR) SetBits(mask BSRR)       { r.U32.SetBits(uint32(mask)) }
func (r *RBSRR) ClearBits(mask BSRR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RBSRR) Load() BSRR              { return BSRR(r.U32.Load()) }
func (r *RBSRR) Store(b BSRR)            { r.U32.Store(uint32(b)) }

type RMBSRR struct{ mmio.UM32 }

func (rm RMBSRR) Load() BSRR   { return BSRR(rm.UM32.Load()) }
func (rm RMBSRR) Store(b BSRR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) BS0() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS0)}}
}

func (p *Periph) BS1() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS1)}}
}

func (p *Periph) BS2() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS2)}}
}

func (p *Periph) BS3() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS3)}}
}

func (p *Periph) BS4() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS4)}}
}

func (p *Periph) BS5() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS5)}}
}

func (p *Periph) BS6() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS6)}}
}

func (p *Periph) BS7() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS7)}}
}

func (p *Periph) BS8() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS8)}}
}

func (p *Periph) BS9() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS9)}}
}

func (p *Periph) BS10() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS10)}}
}

func (p *Periph) BS11() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS11)}}
}

func (p *Periph) BS12() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS12)}}
}

func (p *Periph) BS13() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS13)}}
}

func (p *Periph) BS14() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS14)}}
}

func (p *Periph) BS15() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS15)}}
}

func (p *Periph) BR0() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR0)}}
}

func (p *Periph) BR1() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR1)}}
}

func (p *Periph) BR2() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR2)}}
}

func (p *Periph) BR3() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR3)}}
}

func (p *Periph) BR4() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR4)}}
}

func (p *Periph) BR5() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR5)}}
}

func (p *Periph) BR6() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR6)}}
}

func (p *Periph) BR7() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR7)}}
}

func (p *Periph) BR8() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR8)}}
}

func (p *Periph) BR9() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR9)}}
}

func (p *Periph) BR10() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR10)}}
}

func (p *Periph) BR11() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR11)}}
}

func (p *Periph) BR12() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR12)}}
}

func (p *Periph) BR13() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR13)}}
}

func (p *Periph) BR14() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR14)}}
}

func (p *Periph) BR15() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR15)}}
}

type LCKR uint32

type RLCKR struct{ mmio.U32 }

func (r *RLCKR) LoadBits(mask LCKR) LCKR { return LCKR(r.U32.LoadBits(uint32(mask))) }
func (r *RLCKR) StoreBits(mask, b LCKR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RLCKR) SetBits(mask LCKR)       { r.U32.SetBits(uint32(mask)) }
func (r *RLCKR) ClearBits(mask LCKR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RLCKR) Load() LCKR              { return LCKR(r.U32.Load()) }
func (r *RLCKR) Store(b LCKR)            { r.U32.Store(uint32(b)) }

type RMLCKR struct{ mmio.UM32 }

func (rm RMLCKR) Load() LCKR   { return LCKR(rm.UM32.Load()) }
func (rm RMLCKR) Store(b LCKR) { rm.UM32.Store(uint32(b)) }

func (p *Periph) LCK0() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK0)}}
}

func (p *Periph) LCK1() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK1)}}
}

func (p *Periph) LCK2() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK2)}}
}

func (p *Periph) LCK3() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK3)}}
}

func (p *Periph) LCK4() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK4)}}
}

func (p *Periph) LCK5() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK5)}}
}

func (p *Periph) LCK6() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK6)}}
}

func (p *Periph) LCK7() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK7)}}
}

func (p *Periph) LCK8() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK8)}}
}

func (p *Periph) LCK9() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK9)}}
}

func (p *Periph) LCK10() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK10)}}
}

func (p *Periph) LCK11() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK11)}}
}

func (p *Periph) LCK12() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK12)}}
}

func (p *Periph) LCK13() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK13)}}
}

func (p *Periph) LCK14() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK14)}}
}

func (p *Periph) LCK15() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK15)}}
}

func (p *Periph) LCKK() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCKK)}}
}

type AFRL uint32

type RAFRL struct{ mmio.U32 }

func (r *RAFRL) LoadBits(mask AFRL) AFRL { return AFRL(r.U32.LoadBits(uint32(mask))) }
func (r *RAFRL) StoreBits(mask, b AFRL)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAFRL) SetBits(mask AFRL)       { r.U32.SetBits(uint32(mask)) }
func (r *RAFRL) ClearBits(mask AFRL)     { r.U32.ClearBits(uint32(mask)) }
func (r *RAFRL) Load() AFRL              { return AFRL(r.U32.Load()) }
func (r *RAFRL) Store(b AFRL)            { r.U32.Store(uint32(b)) }

type RMAFRL struct{ mmio.UM32 }

func (rm RMAFRL) Load() AFRL   { return AFRL(rm.UM32.Load()) }
func (rm RMAFRL) Store(b AFRL) { rm.UM32.Store(uint32(b)) }

func (p *Periph) AFRL0() RMAFRL {
	return RMAFRL{mmio.UM32{&p.AFRL.U32, uint32(AFRL0)}}
}

func (p *Periph) AFRL1() RMAFRL {
	return RMAFRL{mmio.UM32{&p.AFRL.U32, uint32(AFRL1)}}
}

func (p *Periph) AFRL2() RMAFRL {
	return RMAFRL{mmio.UM32{&p.AFRL.U32, uint32(AFRL2)}}
}

func (p *Periph) AFRL3() RMAFRL {
	return RMAFRL{mmio.UM32{&p.AFRL.U32, uint32(AFRL3)}}
}

func (p *Periph) AFRL4() RMAFRL {
	return RMAFRL{mmio.UM32{&p.AFRL.U32, uint32(AFRL4)}}
}

func (p *Periph) AFRL5() RMAFRL {
	return RMAFRL{mmio.UM32{&p.AFRL.U32, uint32(AFRL5)}}
}

func (p *Periph) AFRL6() RMAFRL {
	return RMAFRL{mmio.UM32{&p.AFRL.U32, uint32(AFRL6)}}
}

func (p *Periph) AFRL7() RMAFRL {
	return RMAFRL{mmio.UM32{&p.AFRL.U32, uint32(AFRL7)}}
}

type AFRH uint32

type RAFRH struct{ mmio.U32 }

func (r *RAFRH) LoadBits(mask AFRH) AFRH { return AFRH(r.U32.LoadBits(uint32(mask))) }
func (r *RAFRH) StoreBits(mask, b AFRH)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAFRH) SetBits(mask AFRH)       { r.U32.SetBits(uint32(mask)) }
func (r *RAFRH) ClearBits(mask AFRH)     { r.U32.ClearBits(uint32(mask)) }
func (r *RAFRH) Load() AFRH              { return AFRH(r.U32.Load()) }
func (r *RAFRH) Store(b AFRH)            { r.U32.Store(uint32(b)) }

type RMAFRH struct{ mmio.UM32 }

func (rm RMAFRH) Load() AFRH   { return AFRH(rm.UM32.Load()) }
func (rm RMAFRH) Store(b AFRH) { rm.UM32.Store(uint32(b)) }

func (p *Periph) AFRH8() RMAFRH {
	return RMAFRH{mmio.UM32{&p.AFRH.U32, uint32(AFRH8)}}
}

func (p *Periph) AFRH9() RMAFRH {
	return RMAFRH{mmio.UM32{&p.AFRH.U32, uint32(AFRH9)}}
}

func (p *Periph) AFRH10() RMAFRH {
	return RMAFRH{mmio.UM32{&p.AFRH.U32, uint32(AFRH10)}}
}

func (p *Periph) AFRH11() RMAFRH {
	return RMAFRH{mmio.UM32{&p.AFRH.U32, uint32(AFRH11)}}
}

func (p *Periph) AFRH12() RMAFRH {
	return RMAFRH{mmio.UM32{&p.AFRH.U32, uint32(AFRH12)}}
}

func (p *Periph) AFRH13() RMAFRH {
	return RMAFRH{mmio.UM32{&p.AFRH.U32, uint32(AFRH13)}}
}

func (p *Periph) AFRH14() RMAFRH {
	return RMAFRH{mmio.UM32{&p.AFRH.U32, uint32(AFRH14)}}
}

func (p *Periph) AFRH15() RMAFRH {
	return RMAFRH{mmio.UM32{&p.AFRH.U32, uint32(AFRH15)}}
}
