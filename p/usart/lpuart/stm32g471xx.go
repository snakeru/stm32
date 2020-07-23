// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32g471xx

// Package lpuart provides access to the registers of the LPUART peripheral.
//
// Instances:
//  LPUART1  LPUART1_BASE  APB1  LPTIM1,LPUART  Universal synchronous asynchronous receiver transmitter
// Registers:
//  0x000 32  CR1    Control register 1
//  0x004 32  CR2    Control register 2
//  0x008 32  CR3    Control register 3
//  0x00C 32  BRR    Baud rate register
//  0x018 32  RQR    Request register
//  0x01C 32  ISR    Interrupt & status register
//  0x020 32  ICR    Interrupt flag clear register
//  0x024 32  RDR    Receive data register
//  0x028 32  TDR    Transmit data register
//  0x02C 32  PRESC  Prescaler register
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package lpuart

const (
	UE     CR1 = 0x01 << 0  //+ USART enable
	UESM   CR1 = 0x01 << 1  //+ USART enable in Stop mode
	RE     CR1 = 0x01 << 2  //+ Receiver enable
	TE     CR1 = 0x01 << 3  //+ Transmitter enable
	IDLEIE CR1 = 0x01 << 4  //+ IDLE interrupt enable
	RXNEIE CR1 = 0x01 << 5  //+ RXNE interrupt enable
	TCIE   CR1 = 0x01 << 6  //+ Transmission complete interrupt enable
	TXEIE  CR1 = 0x01 << 7  //+ interrupt enable
	PEIE   CR1 = 0x01 << 8  //+ PE interrupt enable
	PS     CR1 = 0x01 << 9  //+ Parity selection
	PCE    CR1 = 0x01 << 10 //+ Parity control enable
	WAKE   CR1 = 0x01 << 11 //+ Receiver wakeup method
	M0     CR1 = 0x01 << 12 //+ Word length
	MME    CR1 = 0x01 << 13 //+ Mute mode enable
	CMIE   CR1 = 0x01 << 14 //+ Character match interrupt enable
	DEDT0  CR1 = 0x01 << 16 //+ DEDT0
	DEDT1  CR1 = 0x01 << 17 //+ DEDT1
	DEDT2  CR1 = 0x01 << 18 //+ DEDT2
	DEDT3  CR1 = 0x01 << 19 //+ DEDT3
	DEDT4  CR1 = 0x01 << 20 //+ Driver Enable de-assertion time
	DEAT0  CR1 = 0x01 << 21 //+ DEAT0
	DEAT1  CR1 = 0x01 << 22 //+ DEAT1
	DEAT2  CR1 = 0x01 << 23 //+ DEAT2
	DEAT3  CR1 = 0x01 << 24 //+ DEAT3
	DEAT4  CR1 = 0x01 << 25 //+ Driver Enable assertion time
	M1     CR1 = 0x01 << 28 //+ Word length
	FIFOEN CR1 = 0x01 << 29 //+ FIFOEN
	TXFEIE CR1 = 0x01 << 30 //+ TXFEIE
	RXFFIE CR1 = 0x01 << 31 //+ RXFFIE
)

const (
	UEn     = 0
	UESMn   = 1
	REn     = 2
	TEn     = 3
	IDLEIEn = 4
	RXNEIEn = 5
	TCIEn   = 6
	TXEIEn  = 7
	PEIEn   = 8
	PSn     = 9
	PCEn    = 10
	WAKEn   = 11
	M0n     = 12
	MMEn    = 13
	CMIEn   = 14
	DEDT0n  = 16
	DEDT1n  = 17
	DEDT2n  = 18
	DEDT3n  = 19
	DEDT4n  = 20
	DEAT0n  = 21
	DEAT1n  = 22
	DEAT2n  = 23
	DEAT3n  = 24
	DEAT4n  = 25
	M1n     = 28
	FIFOENn = 29
	TXFEIEn = 30
	RXFFIEn = 31
)

const (
	ADDM7    CR2 = 0x01 << 4  //+ 7-bit Address Detection/4-bit Address Detection
	STOP     CR2 = 0x03 << 12 //+ STOP bits
	SWAP     CR2 = 0x01 << 15 //+ Swap TX/RX pins
	RXINV    CR2 = 0x01 << 16 //+ RX pin active level inversion
	TXINV    CR2 = 0x01 << 17 //+ TX pin active level inversion
	TAINV    CR2 = 0x01 << 18 //+ Binary data inversion
	MSBFIRST CR2 = 0x01 << 19 //+ Most significant bit first
	ADD0_3   CR2 = 0x0F << 24 //+ Address of the USART node
	ADD4_7   CR2 = 0x0F << 28 //+ Address of the USART node
)

const (
	ADDM7n    = 4
	STOPn     = 12
	SWAPn     = 15
	RXINVn    = 16
	TXINVn    = 17
	TAINVn    = 18
	MSBFIRSTn = 19
	ADD0_3n   = 24
	ADD4_7n   = 28
)

const (
	EIE     CR3 = 0x01 << 0  //+ Error interrupt enable
	HDSEL   CR3 = 0x01 << 3  //+ Half-duplex selection
	DMAR    CR3 = 0x01 << 6  //+ DMA enable receiver
	DMAT    CR3 = 0x01 << 7  //+ DMA enable transmitter
	RTSE    CR3 = 0x01 << 8  //+ RTS enable
	CTSE    CR3 = 0x01 << 9  //+ CTS enable
	CTSIE   CR3 = 0x01 << 10 //+ CTS interrupt enable
	OVRDIS  CR3 = 0x01 << 12 //+ Overrun Disable
	DDRE    CR3 = 0x01 << 13 //+ DMA Disable on Reception Error
	DEM     CR3 = 0x01 << 14 //+ Driver enable mode
	DEP     CR3 = 0x01 << 15 //+ Driver enable polarity selection
	WUS     CR3 = 0x03 << 20 //+ Wakeup from Stop mode interrupt flag selection
	WUFIE   CR3 = 0x01 << 22 //+ Wakeup from Stop mode interrupt enable
	TXFTIE  CR3 = 0x01 << 23 //+ TXFTIE
	RXFTCFG CR3 = 0x07 << 25 //+ RXFTCFG
	RXFTIE  CR3 = 0x01 << 28 //+ RXFTIE
	TXFTCFG CR3 = 0x07 << 29 //+ TXFTCFG
)

const (
	EIEn     = 0
	HDSELn   = 3
	DMARn    = 6
	DMATn    = 7
	RTSEn    = 8
	CTSEn    = 9
	CTSIEn   = 10
	OVRDISn  = 12
	DDREn    = 13
	DEMn     = 14
	DEPn     = 15
	WUSn     = 20
	WUFIEn   = 22
	TXFTIEn  = 23
	RXFTCFGn = 25
	RXFTIEn  = 28
	TXFTCFGn = 29
)

const (
	BRR BRR = 0xFFFFF << 0 //+ BRR
)

const (
	BRRn = 0
)

const (
	SBKRQ RQR = 0x01 << 1 //+ Send break request
	MMRQ  RQR = 0x01 << 2 //+ Mute mode request
	RXFRQ RQR = 0x01 << 3 //+ Receive data flush request
	TXFRQ RQR = 0x01 << 4 //+ TXFRQ
)

const (
	SBKRQn = 1
	MMRQn  = 2
	RXFRQn = 3
	TXFRQn = 4
)

const (
	PE    ISR = 0x01 << 0  //+ PE
	FE    ISR = 0x01 << 1  //+ FE
	NF    ISR = 0x01 << 2  //+ NF
	ORE   ISR = 0x01 << 3  //+ ORE
	IDLE  ISR = 0x01 << 4  //+ IDLE
	RXNE  ISR = 0x01 << 5  //+ RXNE
	TC    ISR = 0x01 << 6  //+ TC
	TXE   ISR = 0x01 << 7  //+ TXE
	CTSIF ISR = 0x01 << 9  //+ CTSIF
	CTS   ISR = 0x01 << 10 //+ CTS
	BUSY  ISR = 0x01 << 16 //+ BUSY
	CMF   ISR = 0x01 << 17 //+ CMF
	SBKF  ISR = 0x01 << 18 //+ SBKF
	RWU   ISR = 0x01 << 19 //+ RWU
	WUF   ISR = 0x01 << 20 //+ WUF
	TEACK ISR = 0x01 << 21 //+ TEACK
	REACK ISR = 0x01 << 22 //+ REACK
	TXFE  ISR = 0x01 << 23 //+ TXFE
	RXFF  ISR = 0x01 << 24 //+ RXFF
	RXFT  ISR = 0x01 << 26 //+ RXFT
	TXFT  ISR = 0x01 << 27 //+ TXFT
)

const (
	PEn    = 0
	FEn    = 1
	NFn    = 2
	OREn   = 3
	IDLEn  = 4
	RXNEn  = 5
	TCn    = 6
	TXEn   = 7
	CTSIFn = 9
	CTSn   = 10
	BUSYn  = 16
	CMFn   = 17
	SBKFn  = 18
	RWUn   = 19
	WUFn   = 20
	TEACKn = 21
	REACKn = 22
	TXFEn  = 23
	RXFFn  = 24
	RXFTn  = 26
	TXFTn  = 27
)

const (
	PECF   ICR = 0x01 << 0  //+ Parity error clear flag
	FECF   ICR = 0x01 << 1  //+ Framing error clear flag
	NCF    ICR = 0x01 << 2  //+ Noise detected clear flag
	ORECF  ICR = 0x01 << 3  //+ Overrun error clear flag
	IDLECF ICR = 0x01 << 4  //+ Idle line detected clear flag
	TCCF   ICR = 0x01 << 6  //+ Transmission complete clear flag
	CTSCF  ICR = 0x01 << 9  //+ CTS clear flag
	CMCF   ICR = 0x01 << 17 //+ Character match clear flag
	WUCF   ICR = 0x01 << 20 //+ Wakeup from Stop mode clear flag
)

const (
	PECFn   = 0
	FECFn   = 1
	NCFn    = 2
	ORECFn  = 3
	IDLECFn = 4
	TCCFn   = 6
	CTSCFn  = 9
	CMCFn   = 17
	WUCFn   = 20
)

const (
	RDR RDR = 0x1FF << 0 //+ Receive data value
)

const (
	RDRn = 0
)

const (
	TDR TDR = 0x1FF << 0 //+ Transmit data value
)

const (
	TDRn = 0
)

const (
	PRESCALER PRESC = 0x0F << 0 //+ PRESCALER
)

const (
	PRESCALERn = 0
)
