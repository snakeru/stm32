// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32h7x3

// Package spdifrx provides access to the registers of the SPDIFRX peripheral.
//
// Instances:
//  SPDIFRX  SPDIFRX_BASE  APB1  SPDIF  Receiver Interface
// Registers:
//  0x000 32  CR     Control register
//  0x004 32  IMR    Interrupt mask register
//  0x008 32  SR     Status register
//  0x00C 32  IFCR   Interrupt Flag Clear register
//  0x010 32  DR_00  Data input register
//  0x010 32  DR_10  Data input register
//  0x010 32  DR_01  Data input register
//  0x014 32  CSR    Channel Status register
//  0x018 32  DIR    Debug Information register
//  0x3F4 32  VERR   SPDIFRX version register
//  0x3F8 32  IDR    SPDIFRX identification register
//  0x3FC 32  SIDR   SPDIFRX size identification register
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package spdifrx

const (
	SPDIFRXEN CR = 0x03 << 0  //+ Peripheral Block Enable
	RXDMAEN   CR = 0x01 << 2  //+ Receiver DMA ENable for data flow
	RXSTEO    CR = 0x01 << 3  //+ STerEO Mode
	DRFMT     CR = 0x03 << 4  //+ RX Data format
	PMSK      CR = 0x01 << 6  //+ Mask Parity error bit
	VMSK      CR = 0x01 << 7  //+ Mask of Validity bit
	CUMSK     CR = 0x01 << 8  //+ Mask of channel status and user bits
	PTMSK     CR = 0x01 << 9  //+ Mask of Preamble Type bits
	CBDMAEN   CR = 0x01 << 10 //+ Control Buffer DMA ENable for control flow
	CHSEL     CR = 0x01 << 11 //+ Channel Selection
	NBTR      CR = 0x03 << 12 //+ Maximum allowed re-tries during synchronization phase
	WFA       CR = 0x01 << 14 //+ Wait For Activity
	INSEL     CR = 0x07 << 16 //+ input selection
	CKSEN     CR = 0x01 << 20 //+ Symbol Clock Enable
	CKSBKPEN  CR = 0x01 << 21 //+ Backup Symbol Clock Enable
)

const (
	SPDIFRXENn = 0
	RXDMAENn   = 2
	RXSTEOn    = 3
	DRFMTn     = 4
	PMSKn      = 6
	VMSKn      = 7
	CUMSKn     = 8
	PTMSKn     = 9
	CBDMAENn   = 10
	CHSELn     = 11
	NBTRn      = 12
	WFAn       = 14
	INSELn     = 16
	CKSENn     = 20
	CKSBKPENn  = 21
)

const (
	RXNEIE  IMR = 0x01 << 0 //+ RXNE interrupt enable
	CSRNEIE IMR = 0x01 << 1 //+ Control Buffer Ready Interrupt Enable
	PERRIE  IMR = 0x01 << 2 //+ Parity error interrupt enable
	OVRIE   IMR = 0x01 << 3 //+ Overrun error Interrupt Enable
	SBLKIE  IMR = 0x01 << 4 //+ Synchronization Block Detected Interrupt Enable
	SYNCDIE IMR = 0x01 << 5 //+ Synchronization Done
	IFEIE   IMR = 0x01 << 6 //+ Serial Interface Error Interrupt Enable
)

const (
	RXNEIEn  = 0
	CSRNEIEn = 1
	PERRIEn  = 2
	OVRIEn   = 3
	SBLKIEn  = 4
	SYNCDIEn = 5
	IFEIEn   = 6
)

const (
	RXNE   SR = 0x01 << 0    //+ Read data register not empty
	CSRNE  SR = 0x01 << 1    //+ Control Buffer register is not empty
	PERR   SR = 0x01 << 2    //+ Parity error
	OVR    SR = 0x01 << 3    //+ Overrun error
	SBD    SR = 0x01 << 4    //+ Synchronization Block Detected
	SYNCD  SR = 0x01 << 5    //+ Synchronization Done
	FERR   SR = 0x01 << 6    //+ Framing error
	SERR   SR = 0x01 << 7    //+ Synchronization error
	TERR   SR = 0x01 << 8    //+ Time-out error
	WIDTH5 SR = 0x7FFF << 16 //+ Duration of 5 symbols counted with SPDIF_CLK
)

const (
	RXNEn   = 0
	CSRNEn  = 1
	PERRn   = 2
	OVRn    = 3
	SBDn    = 4
	SYNCDn  = 5
	FERRn   = 6
	SERRn   = 7
	TERRn   = 8
	WIDTH5n = 16
)

const (
	PERRCF  IFCR = 0x01 << 2 //+ Clears the Parity error flag
	OVRCF   IFCR = 0x01 << 3 //+ Clears the Overrun error flag
	SBDCF   IFCR = 0x01 << 4 //+ Clears the Synchronization Block Detected flag
	SYNCDCF IFCR = 0x01 << 5 //+ Clears the Synchronization Done flag
)

const (
	PERRCFn  = 2
	OVRCFn   = 3
	SBDCFn   = 4
	SYNCDCFn = 5
)

const (
	DR DR_00 = 0xFFFFFF << 0 //+ Parity Error bit
	PE DR_00 = 0x01 << 24    //+ Parity Error bit
	V  DR_00 = 0x01 << 25    //+ Validity bit
	U  DR_00 = 0x01 << 26    //+ User bit
	C  DR_00 = 0x01 << 27    //+ Channel Status bit
	PT DR_00 = 0x03 << 28    //+ Preamble Type
)

const (
	DRn = 0
	PEn = 24
	Vn  = 25
	Un  = 26
	Cn  = 27
	PTn = 28
)

const (
	DRNL1 DR_10 = 0xFFFF << 0  //+ Data value
	DRNL2 DR_10 = 0xFFFF << 16 //+ Data value
)

const (
	DRNL1n = 0
	DRNL2n = 16
)

const (
	PE DR_01 = 0x01 << 0     //+ Parity Error bit
	V  DR_01 = 0x01 << 1     //+ Validity bit
	U  DR_01 = 0x01 << 2     //+ User bit
	C  DR_01 = 0x01 << 3     //+ Channel Status bit
	PT DR_01 = 0x03 << 4     //+ Preamble Type
	DR DR_01 = 0xFFFFFF << 8 //+ Data value
)

const (
	PEn = 0
	Vn  = 1
	Un  = 2
	Cn  = 3
	PTn = 4
	DRn = 8
)

const (
	USR CSR = 0xFFFF << 0 //+ User data information
	CS  CSR = 0xFF << 16  //+ Channel A status information
	SOB CSR = 0x01 << 24  //+ Start Of Block
)

const (
	USRn = 0
	CSn  = 16
	SOBn = 24
)

const (
	THI DIR = 0x1FFF << 0  //+ Threshold HIGH
	TLO DIR = 0x1FFF << 16 //+ Threshold LOW
)

const (
	THIn = 0
	TLOn = 16
)

const (
	MINREV VERR = 0x0F << 0 //+ Minor revision
	MAJREV VERR = 0x0F << 4 //+ Major revision
)

const (
	MINREVn = 0
	MAJREVn = 4
)

const (
	ID IDR = 0xFFFFFFFF << 0 //+ SPDIFRX identifier
)

const (
	IDn = 0
)

const (
	SID SIDR = 0xFFFFFFFF << 0 //+ Size identification
)

const (
	SIDn = 0
)
