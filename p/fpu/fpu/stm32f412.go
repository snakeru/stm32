// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32f412

// Package fpu provides access to the registers of the FPU peripheral.
//
// Instances:
//  FPU  FPU_BASE  -  FPU+,FPU+  Floting point unit
// Registers:
//  0x000 32  FPCCR  Floating-point context control register
//  0x004 32  FPCAR  Floating-point context address register
//  0x008 32  FPSCR  Floating-point status control register
// Import:
//  github.com/embeddedgo/stm32/p/mmap
package fpu

const (
	LSPACT FPCCR = 0x01 << 0  //+ LSPACT
	USER   FPCCR = 0x01 << 1  //+ USER
	THREAD FPCCR = 0x01 << 3  //+ THREAD
	HFRDY  FPCCR = 0x01 << 4  //+ HFRDY
	MMRDY  FPCCR = 0x01 << 5  //+ MMRDY
	BFRDY  FPCCR = 0x01 << 6  //+ BFRDY
	MONRDY FPCCR = 0x01 << 8  //+ MONRDY
	LSPEN  FPCCR = 0x01 << 30 //+ LSPEN
	ASPEN  FPCCR = 0x01 << 31 //+ ASPEN
)

const (
	LSPACTn = 0
	USERn   = 1
	THREADn = 3
	HFRDYn  = 4
	MMRDYn  = 5
	BFRDYn  = 6
	MONRDYn = 8
	LSPENn  = 30
	ASPENn  = 31
)

const (
	ADDRESS FPCAR = 0x1FFFFFFF << 3 //+ Location of unpopulated floating-point
)

const (
	ADDRESSn = 3
)

const (
	IOC   FPSCR = 0x01 << 0  //+ Invalid operation cumulative exception bit
	DZC   FPSCR = 0x01 << 1  //+ Division by zero cumulative exception bit.
	OFC   FPSCR = 0x01 << 2  //+ Overflow cumulative exception bit
	UFC   FPSCR = 0x01 << 3  //+ Underflow cumulative exception bit
	IXC   FPSCR = 0x01 << 4  //+ Inexact cumulative exception bit
	IDC   FPSCR = 0x01 << 7  //+ Input denormal cumulative exception bit.
	RMode FPSCR = 0x03 << 22 //+ Rounding Mode control field
	FZ    FPSCR = 0x01 << 24 //+ Flush-to-zero mode control bit:
	DN    FPSCR = 0x01 << 25 //+ Default NaN mode control bit
	AHP   FPSCR = 0x01 << 26 //+ Alternative half-precision control bit
	V     FPSCR = 0x01 << 28 //+ Overflow condition code flag
	C     FPSCR = 0x01 << 29 //+ Carry condition code flag
	Z     FPSCR = 0x01 << 30 //+ Zero condition code flag
	N     FPSCR = 0x01 << 31 //+ Negative condition code flag
)

const (
	IOCn   = 0
	DZCn   = 1
	OFCn   = 2
	UFCn   = 3
	IXCn   = 4
	IDCn   = 7
	RModen = 22
	FZn    = 24
	DNn    = 25
	AHPn   = 26
	Vn     = 28
	Cn     = 29
	Zn     = 30
	Nn     = 31
)
