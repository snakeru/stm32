// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32f412

// Package i2c4 provides access to the registers of the I2C peripheral.
//
// Instances:
//  I2C4  I2C4_BASE  APB1  -  Inter-integrated circuit
// Registers:
//  0x000 32  CR1       Control register 1
//  0x004 32  CR2       Control register 2
//  0x008 32  OAR1      Own address register 1
//  0x00C 32  OAR2      Own address register 2
//  0x010 32  TIMINGR   Timing register
//  0x014 32  TIMEOUTR  Timeout register
//  0x018 32  ISR       Interrupt and Status register
//  0x01C 32  ICR       Interrupt clear register
//  0x020 32  PECR      PEC register
//  0x024 32  RXDR      Receive data register
//  0x028 32  TXDR      Transmit data register
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package i2c4

const (
	PE        CR1 = 0x01 << 0  //+ Peripheral enable
	TXIE      CR1 = 0x01 << 1  //+ TXIE
	RXIE      CR1 = 0x01 << 2  //+ RXIE
	ADDRE     CR1 = 0x01 << 3  //+ ADDRE
	NACKIE    CR1 = 0x01 << 4  //+ NACKIE
	STOPIE    CR1 = 0x01 << 5  //+ STOPIE
	TCIE      CR1 = 0x01 << 6  //+ TCIE
	ERRIE     CR1 = 0x01 << 7  //+ ERRIE
	DNF       CR1 = 0x0F << 8  //+ DNF
	ANFOFF    CR1 = 0x01 << 12 //+ ANFOFF
	TCDMAEN   CR1 = 0x01 << 14 //+ TCDMAEN
	RXDMAEN   CR1 = 0x01 << 15 //+ RXDMAEN
	SBC       CR1 = 0x01 << 16 //+ SBC
	NOSTRETCH CR1 = 0x01 << 17 //+ NOSTRETCH
	WUPEN     CR1 = 0x01 << 18 //+ WUPEN
	GCEN      CR1 = 0x01 << 19 //+ GCEN
	SMBHEN    CR1 = 0x01 << 20 //+ SMBHEN
	SMBDEN    CR1 = 0x01 << 21 //+ SMBDEN
	ALERTEN   CR1 = 0x01 << 22 //+ ALERTEN
	PECEN     CR1 = 0x01 << 23 //+ PECEN
)

const (
	PEn        = 0
	TXIEn      = 1
	RXIEn      = 2
	ADDREn     = 3
	NACKIEn    = 4
	STOPIEn    = 5
	TCIEn      = 6
	ERRIEn     = 7
	DNFn       = 8
	ANFOFFn    = 12
	TCDMAENn   = 14
	RXDMAENn   = 15
	SBCn       = 16
	NOSTRETCHn = 17
	WUPENn     = 18
	GCENn      = 19
	SMBHENn    = 20
	SMBDENn    = 21
	ALERTENn   = 22
	PECENn     = 23
)

const (
	SADD0   CR2 = 0x01 << 0  //+ Slave address bit 0
	SADD1_7 CR2 = 0x7F << 1  //+ Slave address bit 7_1
	SADD8_9 CR2 = 0x03 << 8  //+ Slave address bit 8_9
	RD_WRN  CR2 = 0x01 << 10 //+ Transfer direction
	ADD10   CR2 = 0x01 << 11 //+ 10-bit addressing mode
	HEAD10R CR2 = 0x01 << 12 //+ 10-bit address header only read direction
	START   CR2 = 0x01 << 13 //+ Start generation
	STOP    CR2 = 0x01 << 14 //+ Stop generation
	NACK    CR2 = 0x01 << 15 //+ NACK generation
	NBYTES  CR2 = 0xFF << 16 //+ Number of bytes
	RELOAD  CR2 = 0x01 << 24 //+ NBYTES reload mode
	AUTOEND CR2 = 0x01 << 25 //+ Automatic end mode
	PECBYTE CR2 = 0x01 << 26 //+ Packet error checking byte
)

const (
	SADD0n   = 0
	SADD1_7n = 1
	SADD8_9n = 8
	RD_WRNn  = 10
	ADD10n   = 11
	HEAD10Rn = 12
	STARTn   = 13
	STOPn    = 14
	NACKn    = 15
	NBYTESn  = 16
	RELOADn  = 24
	AUTOENDn = 25
	PECBYTEn = 26
)

const (
	OA1     OAR1 = 0x01 << 0  //+ OA1
	OA11_7  OAR1 = 0x7F << 1  //+ OA11_7
	OA18_9  OAR1 = 0x03 << 8  //+ OA18_9
	OA1MODE OAR1 = 0x01 << 10 //+ OA1MODE
	OA1EN   OAR1 = 0x01 << 15 //+ OA1EN
)

const (
	OA1n     = 0
	OA11_7n  = 1
	OA18_9n  = 8
	OA1MODEn = 10
	OA1ENn   = 15
)

const (
	OA21_7 OAR2 = 0x7F << 1  //+ OA21_7
	OA2MSK OAR2 = 0x07 << 8  //+ OA2MSK
	OA2EN  OAR2 = 0x01 << 15 //+ OA2EN
)

const (
	OA21_7n = 1
	OA2MSKn = 8
	OA2ENn  = 15
)

const (
	SCLL   TIMINGR = 0xFF << 0  //+ SCLL
	SCLH   TIMINGR = 0xFF << 8  //+ SCLH
	SDADEL TIMINGR = 0x0F << 16 //+ SDADEL
	SCLDEL TIMINGR = 0x0F << 20 //+ SCLDEL
	PRESC  TIMINGR = 0x0F << 28 //+ PRESC
)

const (
	SCLLn   = 0
	SCLHn   = 8
	SDADELn = 16
	SCLDELn = 20
	PRESCn  = 28
)

const (
	TIMEOUTA TIMEOUTR = 0xFFF << 0  //+ TIMEOUTA
	TIDLE    TIMEOUTR = 0x01 << 12  //+ TIDLE
	TIMOUTEN TIMEOUTR = 0x01 << 15  //+ TIMOUTEN
	TIMEOUTB TIMEOUTR = 0xFFF << 16 //+ TIMEOUTB
	TEXTEN   TIMEOUTR = 0x01 << 31  //+ TEXTEN
)

const (
	TIMEOUTAn = 0
	TIDLEn    = 12
	TIMOUTENn = 15
	TIMEOUTBn = 16
	TEXTENn   = 31
)

const (
	TXE     ISR = 0x01 << 0  //+ TXE
	TXIS    ISR = 0x01 << 1  //+ TXIS
	RXNE    ISR = 0x01 << 2  //+ RXNE
	ADDR    ISR = 0x01 << 3  //+ ADDR
	NACKF   ISR = 0x01 << 4  //+ NACKF
	STOPF   ISR = 0x01 << 5  //+ STOPF
	TC      ISR = 0x01 << 6  //+ TC
	TCR     ISR = 0x01 << 7  //+ TCR
	BERR    ISR = 0x01 << 8  //+ BERR
	ARLO    ISR = 0x01 << 9  //+ ARLO
	OVR     ISR = 0x01 << 10 //+ OVR
	PECERR  ISR = 0x01 << 11 //+ PECERR
	TIMEOUT ISR = 0x01 << 12 //+ TIMEOUT
	ALERT   ISR = 0x01 << 13 //+ ALERT
	BUSY    ISR = 0x01 << 15 //+ BUSY
	DIR     ISR = 0x01 << 16 //+ DIR
	ADDCODE ISR = 0x7F << 17 //+ ADDCODE
)

const (
	TXEn     = 0
	TXISn    = 1
	RXNEn    = 2
	ADDRn    = 3
	NACKFn   = 4
	STOPFn   = 5
	TCn      = 6
	TCRn     = 7
	BERRn    = 8
	ARLOn    = 9
	OVRn     = 10
	PECERRn  = 11
	TIMEOUTn = 12
	ALERTn   = 13
	BUSYn    = 15
	DIRn     = 16
	ADDCODEn = 17
)

const (
	ADDRCF   ICR = 0x01 << 3  //+ ADDRCF
	NACKCF   ICR = 0x01 << 4  //+ NACKCF
	STOPCF   ICR = 0x01 << 5  //+ STOPCF
	BERRCF   ICR = 0x01 << 8  //+ BERRCF
	ARLOCF   ICR = 0x01 << 9  //+ ARLOCF
	OVRCF    ICR = 0x01 << 10 //+ OVRCF
	PECCF    ICR = 0x01 << 11 //+ PECCF
	TIMOUTCF ICR = 0x01 << 12 //+ TIMOUTCF
	ALERTC   ICR = 0x01 << 13 //+ ALERTC
)

const (
	ADDRCFn   = 3
	NACKCFn   = 4
	STOPCFn   = 5
	BERRCFn   = 8
	ARLOCFn   = 9
	OVRCFn    = 10
	PECCFn    = 11
	TIMOUTCFn = 12
	ALERTCn   = 13
)

const (
	PEC PECR = 0xFF << 0 //+ PEC
)

const (
	PECn = 0
)

const (
	RXDATA RXDR = 0xFF << 0 //+ RXDATA
)

const (
	RXDATAn = 0
)

const (
	TXDATA TXDR = 0xFF << 0 //+ TXDATA
)

const (
	TXDATAn = 0
)
