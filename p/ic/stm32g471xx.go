// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32g471xx

// Package ic provides access to the registers of the I2C peripheral.
//
// Instances:
//  I2C1  I2C1_BASE  APB1  I2C1_EV,I2C1_ER       Inter-integrated circuit
//  I2C2  I2C2_BASE  APB1  WWDG,I2C2_EV,I2C2_ER  Inter-integrated circuit
//  I2C3  I2C3_BASE  APB1  I2C3_EV,I2C3_ER       Inter-integrated circuit
//  I2C4  I2C4_BASE  APB1  I2C4_EV,I2C4_ER       Inter-integrated circuit
// Registers:
//  0x000 32  CR1       Control register 1
//  0x004 32  CR2       Control register 2
//  0x008 32  OAR1      Own address register 1
//  0x00C 32  OAR2      Own address register 2
//  0x010 32  TIMINGR   Timing register
//  0x014 32  TIMEOUTR  Status register 1
//  0x018 32  ISR       Interrupt and Status register
//  0x01C 32  ICR       Interrupt clear register
//  0x020 32  PECR      PEC register
//  0x024 32  RXDR      Receive data register
//  0x028 32  TXDR      Transmit data register
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package ic

const (
	PE        CR1 = 0x01 << 0  //+ Peripheral enable
	TXIE      CR1 = 0x01 << 1  //+ TX Interrupt enable
	RXIE      CR1 = 0x01 << 2  //+ RX Interrupt enable
	ADDRIE    CR1 = 0x01 << 3  //+ Address match interrupt enable (slave only)
	NACKIE    CR1 = 0x01 << 4  //+ Not acknowledge received interrupt enable
	STOPIE    CR1 = 0x01 << 5  //+ STOP detection Interrupt enable
	TCIE      CR1 = 0x01 << 6  //+ Transfer Complete interrupt enable
	ERRIE     CR1 = 0x01 << 7  //+ Error interrupts enable
	DNF       CR1 = 0x0F << 8  //+ Digital noise filter
	ANFOFF    CR1 = 0x01 << 12 //+ Analog noise filter OFF
	TXDMAEN   CR1 = 0x01 << 14 //+ DMA transmission requests enable
	RXDMAEN   CR1 = 0x01 << 15 //+ DMA reception requests enable
	SBC       CR1 = 0x01 << 16 //+ Slave byte control
	NOSTRETCH CR1 = 0x01 << 17 //+ Clock stretching disable
	WUPEN     CR1 = 0x01 << 18 //+ Wakeup from STOP enable
	GCEN      CR1 = 0x01 << 19 //+ General call enable
	SMBHEN    CR1 = 0x01 << 20 //+ SMBus Host address enable
	SMBDEN    CR1 = 0x01 << 21 //+ SMBus Device Default address enable
	ALERTEN   CR1 = 0x01 << 22 //+ SMBUS alert enable
	PECEN     CR1 = 0x01 << 23 //+ PEC enable
)

const (
	PEn        = 0
	TXIEn      = 1
	RXIEn      = 2
	ADDRIEn    = 3
	NACKIEn    = 4
	STOPIEn    = 5
	TCIEn      = 6
	ERRIEn     = 7
	DNFn       = 8
	ANFOFFn    = 12
	TXDMAENn   = 14
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
	SADD    CR2 = 0x3FF << 0 //+ Slave address bit (master mode)
	RD_WRN  CR2 = 0x01 << 10 //+ Transfer direction (master mode)
	ADD10   CR2 = 0x01 << 11 //+ 10-bit addressing mode (master mode)
	HEAD10R CR2 = 0x01 << 12 //+ 10-bit address header only read direction (master receiver mode)
	START   CR2 = 0x01 << 13 //+ Start generation
	STOP    CR2 = 0x01 << 14 //+ Stop generation (master mode)
	NACK    CR2 = 0x01 << 15 //+ NACK generation (slave mode)
	NBYTES  CR2 = 0xFF << 16 //+ Number of bytes
	RELOAD  CR2 = 0x01 << 24 //+ NBYTES reload mode
	AUTOEND CR2 = 0x01 << 25 //+ Automatic end mode (master mode)
	PECBYTE CR2 = 0x01 << 26 //+ Packet error checking byte
)

const (
	SADDn    = 0
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
	OA1     OAR1 = 0x3FF << 0 //+ Interface address
	OA1MODE OAR1 = 0x01 << 10 //+ Own Address 1 10-bit mode
	OA1EN   OAR1 = 0x01 << 15 //+ Own Address 1 enable
)

const (
	OA1n     = 0
	OA1MODEn = 10
	OA1ENn   = 15
)

const (
	OA2    OAR2 = 0x7F << 1  //+ Interface address
	OA2MSK OAR2 = 0x07 << 8  //+ Own Address 2 masks
	OA2EN  OAR2 = 0x01 << 15 //+ Own Address 2 enable
)

const (
	OA2n    = 1
	OA2MSKn = 8
	OA2ENn  = 15
)

const (
	SCLL   TIMINGR = 0xFF << 0  //+ SCL low period (master mode)
	SCLH   TIMINGR = 0xFF << 8  //+ SCL high period (master mode)
	SDADEL TIMINGR = 0x0F << 16 //+ Data hold time
	SCLDEL TIMINGR = 0x0F << 20 //+ Data setup time
	PRESC  TIMINGR = 0x0F << 28 //+ Timing prescaler
)

const (
	SCLLn   = 0
	SCLHn   = 8
	SDADELn = 16
	SCLDELn = 20
	PRESCn  = 28
)

const (
	TIMEOUTA TIMEOUTR = 0xFFF << 0  //+ Bus timeout A
	TIDLE    TIMEOUTR = 0x01 << 12  //+ Idle clock timeout detection
	TIMOUTEN TIMEOUTR = 0x01 << 15  //+ Clock timeout enable
	TIMEOUTB TIMEOUTR = 0xFFF << 16 //+ Bus timeout B
	TEXTEN   TIMEOUTR = 0x01 << 31  //+ Extended clock timeout enable
)

const (
	TIMEOUTAn = 0
	TIDLEn    = 12
	TIMOUTENn = 15
	TIMEOUTBn = 16
	TEXTENn   = 31
)

const (
	TXE     ISR = 0x01 << 0  //+ Transmit data register empty (transmitters)
	TXIS    ISR = 0x01 << 1  //+ Transmit interrupt status (transmitters)
	RXNE    ISR = 0x01 << 2  //+ Receive data register not empty (receivers)
	ADDR    ISR = 0x01 << 3  //+ Address matched (slave mode)
	NACKF   ISR = 0x01 << 4  //+ Not acknowledge received flag
	STOPF   ISR = 0x01 << 5  //+ Stop detection flag
	TC      ISR = 0x01 << 6  //+ Transfer Complete (master mode)
	TCR     ISR = 0x01 << 7  //+ Transfer Complete Reload
	BERR    ISR = 0x01 << 8  //+ Bus error
	ARLO    ISR = 0x01 << 9  //+ Arbitration lost
	OVR     ISR = 0x01 << 10 //+ Overrun/Underrun (slave mode)
	PECERR  ISR = 0x01 << 11 //+ PEC Error in reception
	TIMEOUT ISR = 0x01 << 12 //+ Timeout or t_low detection flag
	ALERT   ISR = 0x01 << 13 //+ SMBus alert
	BUSY    ISR = 0x01 << 15 //+ Bus busy
	DIR     ISR = 0x01 << 16 //+ Transfer direction (Slave mode)
	ADDCODE ISR = 0x7F << 17 //+ Address match code (Slave mode)
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
	ADDRCF   ICR = 0x01 << 3  //+ Address Matched flag clear
	NACKCF   ICR = 0x01 << 4  //+ Not Acknowledge flag clear
	STOPCF   ICR = 0x01 << 5  //+ Stop detection flag clear
	BERRCF   ICR = 0x01 << 8  //+ Bus error flag clear
	ARLOCF   ICR = 0x01 << 9  //+ Arbitration lost flag clear
	OVRCF    ICR = 0x01 << 10 //+ Overrun/Underrun flag clear
	PECCF    ICR = 0x01 << 11 //+ PEC Error flag clear
	TIMOUTCF ICR = 0x01 << 12 //+ Timeout detection flag clear
	ALERTCF  ICR = 0x01 << 13 //+ Alert flag clear
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
	ALERTCFn  = 13
)

const (
	PEC PECR = 0xFF << 0 //+ Packet error checking register
)

const (
	PECn = 0
)

const (
	RXDATA RXDR = 0xFF << 0 //+ 8-bit receive data
)

const (
	RXDATAn = 0
)

const (
	TXDATA TXDR = 0xFF << 0 //+ 8-bit transmit data
)

const (
	TXDATAn = 0
)
