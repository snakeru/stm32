// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32g471xx

// Package ucpd provides access to the registers of the UCPD peripheral.
//
// Instances:
//  UCPD1  UCPD1_BASE  -  UCPD1  UCPD1
// Registers:
//  0x000 32  CFG1        UCPD configuration register 1
//  0x004 32  CFG2        UCPD configuration register 2
//  0x00C 32  CR          UCPD configuration register 2
//  0x010 32  IMR         UCPD Interrupt Mask Register
//  0x014 32  SR          UCPD Status Register
//  0x018 32  ICR         UCPD Interrupt Clear Register
//  0x01C 32  TX_ORDSET   UCPD Tx Ordered Set Type Register
//  0x020 32  TX_PAYSZ    UCPD Tx Paysize Register
//  0x024 32  TXDR        UCPD Tx Data Register
//  0x028 32  RX_ORDSET   UCPD Rx Ordered Set Register
//  0x02C 32  RX_PAYSZ    UCPD Rx Paysize Register
//  0x030 32  RXDR        UCPD Rx Data Register
//  0x034 32  RX_ORDEXT1  UCPD Rx Ordered Set Extension Register 1
//  0x038 32  RX_ORDEXT2  UCPD Rx Ordered Set Extension Register 2
// Import:
//  github.com/embeddedgo/stm32/p/mmap
package ucpd

const (
	HBITCLKDIV   CFG1 = 0x3F << 0   //+ HBITCLKDIV
	IFRGAP       CFG1 = 0x1F << 6   //+ IFRGAP
	TRANSWIN     CFG1 = 0x1F << 11  //+ TRANSWIN
	PSC_USBPDCLK CFG1 = 0x07 << 17  //+ PSC_USBPDCLK
	RXORDSETEN   CFG1 = 0x1FF << 20 //+ RXORDSETEN
	TXDMAEN      CFG1 = 0x01 << 29  //+ TXDMAEN
	RXDMAEN      CFG1 = 0x01 << 30  //+ RXDMAEN
	UCPDEN       CFG1 = 0x01 << 31  //+ UCPDEN
)

const (
	HBITCLKDIVn   = 0
	IFRGAPn       = 6
	TRANSWINn     = 11
	PSC_USBPDCLKn = 17
	RXORDSETENn   = 20
	TXDMAENn      = 29
	RXDMAENn      = 30
	UCPDENn       = 31
)

const (
	RXFILTDIS CFG2 = 0x01 << 0 //+ RXFILTDIS
	RXFILT2N3 CFG2 = 0x01 << 1 //+ RXFILT2N3
	FORCECLK  CFG2 = 0x01 << 2 //+ FORCECLK
	WUPEN     CFG2 = 0x01 << 3 //+ WUPEN
)

const (
	RXFILTDISn = 0
	RXFILT2N3n = 1
	FORCECLKn  = 2
	WUPENn     = 3
)

const (
	TXMODE     CR = 0x03 << 0  //+ TXMODE
	TXSEND     CR = 0x01 << 2  //+ TXSEND
	TXHRST     CR = 0x01 << 3  //+ TXHRST
	RXMODE     CR = 0x01 << 4  //+ RXMODE
	PHYRXEN    CR = 0x01 << 5  //+ PHYRXEN
	PHYCCSEL   CR = 0x01 << 6  //+ PHYCCSEL
	ANASUBMODE CR = 0x03 << 7  //+ ANASUBMODE
	ANAMODE    CR = 0x01 << 9  //+ ANAMODE
	CCENABLE   CR = 0x03 << 10 //+ CCENABLE
	FRSRXEN    CR = 0x01 << 16 //+ FRSRXEN
	FRSTX      CR = 0x01 << 17 //+ FRSTX
	RDCH       CR = 0x01 << 18 //+ RDCH
	CC1TCDIS   CR = 0x01 << 20 //+ CC1TCDIS
	CC2TCDIS   CR = 0x01 << 21 //+ CC2TCDIS
)

const (
	TXMODEn     = 0
	TXSENDn     = 2
	TXHRSTn     = 3
	RXMODEn     = 4
	PHYRXENn    = 5
	PHYCCSELn   = 6
	ANASUBMODEn = 7
	ANAMODEn    = 9
	CCENABLEn   = 10
	FRSRXENn    = 16
	FRSTXn      = 17
	RDCHn       = 18
	CC1TCDISn   = 20
	CC2TCDISn   = 21
)

const (
	TXISIE      IMR = 0x01 << 0  //+ TXISIE
	TXMSGDISCIE IMR = 0x01 << 1  //+ TXMSGDISCIE
	TXMSGSENTIE IMR = 0x01 << 2  //+ TXMSGSENTIE
	TXMSGABTIE  IMR = 0x01 << 3  //+ TXMSGABTIE
	HRSTDISCIE  IMR = 0x01 << 4  //+ HRSTDISCIE
	HRSTSENTIE  IMR = 0x01 << 5  //+ HRSTSENTIE
	TXUNDIE     IMR = 0x01 << 6  //+ TXUNDIE
	RXNEIE      IMR = 0x01 << 8  //+ RXNEIE
	RXORDDETIE  IMR = 0x01 << 9  //+ RXORDDETIE
	RXHRSTDETIE IMR = 0x01 << 10 //+ RXHRSTDETIE
	RXOVRIE     IMR = 0x01 << 11 //+ RXOVRIE
	RXMSGENDIE  IMR = 0x01 << 12 //+ RXMSGENDIE
	TYPECEVT1IE IMR = 0x01 << 14 //+ TYPECEVT1IE
	TYPECEVT2IE IMR = 0x01 << 15 //+ TYPECEVT2IE
	FRSEVTIE    IMR = 0x01 << 20 //+ FRSEVTIE
)

const (
	TXISIEn      = 0
	TXMSGDISCIEn = 1
	TXMSGSENTIEn = 2
	TXMSGABTIEn  = 3
	HRSTDISCIEn  = 4
	HRSTSENTIEn  = 5
	TXUNDIEn     = 6
	RXNEIEn      = 8
	RXORDDETIEn  = 9
	RXHRSTDETIEn = 10
	RXOVRIEn     = 11
	RXMSGENDIEn  = 12
	TYPECEVT1IEn = 14
	TYPECEVT2IEn = 15
	FRSEVTIEn    = 20
)

const (
	TXIS             SR = 0x01 << 0  //+ TXIS
	TXMSGDISC        SR = 0x01 << 1  //+ TXMSGDISC
	TXMSGSENT        SR = 0x01 << 2  //+ TXMSGSENT
	TXMSGABT         SR = 0x01 << 3  //+ TXMSGABT
	HRSTDISC         SR = 0x01 << 4  //+ HRSTDISC
	HRSTSENT         SR = 0x01 << 5  //+ HRSTSENT
	TXUND            SR = 0x01 << 6  //+ TXUND
	RXNE             SR = 0x01 << 8  //+ RXNE
	RXORDDET         SR = 0x01 << 9  //+ RXORDDET
	RXHRSTDET        SR = 0x01 << 10 //+ RXHRSTDET
	RXOVR            SR = 0x01 << 11 //+ RXOVR
	RXMSGEND         SR = 0x01 << 12 //+ RXMSGEND
	RXERR            SR = 0x01 << 13 //+ RXERR
	TYPECEVT1        SR = 0x01 << 14 //+ TYPECEVT1
	TYPECEVT2        SR = 0x01 << 15 //+ TYPECEVT2
	TYPEC_VSTATE_CC1 SR = 0x03 << 16 //+ TYPEC_VSTATE_CC1
	TYPEC_VSTATE_CC2 SR = 0x03 << 18 //+ TYPEC_VSTATE_CC2
	FRSEVT           SR = 0x01 << 20 //+ FRSEVT
)

const (
	TXISn             = 0
	TXMSGDISCn        = 1
	TXMSGSENTn        = 2
	TXMSGABTn         = 3
	HRSTDISCn         = 4
	HRSTSENTn         = 5
	TXUNDn            = 6
	RXNEn             = 8
	RXORDDETn         = 9
	RXHRSTDETn        = 10
	RXOVRn            = 11
	RXMSGENDn         = 12
	RXERRn            = 13
	TYPECEVT1n        = 14
	TYPECEVT2n        = 15
	TYPEC_VSTATE_CC1n = 16
	TYPEC_VSTATE_CC2n = 18
	FRSEVTn           = 20
)

const (
	TXMSGDISCCF ICR = 0x01 << 1  //+ TXMSGDISCCF
	TXMSGSENTCF ICR = 0x01 << 2  //+ TXMSGSENTCF
	TXMSGABTCF  ICR = 0x01 << 3  //+ TXMSGABTCF
	HRSTDISCCF  ICR = 0x01 << 4  //+ HRSTDISCCF
	HRSTSENTCF  ICR = 0x01 << 5  //+ HRSTSENTCF
	TXUNDCF     ICR = 0x01 << 6  //+ TXUNDCF
	RXORDDETCF  ICR = 0x01 << 9  //+ RXORDDETCF
	RXHRSTDETCF ICR = 0x01 << 10 //+ RXHRSTDETCF
	RXOVRCF     ICR = 0x01 << 11 //+ RXOVRCF
	RXMSGENDCF  ICR = 0x01 << 12 //+ RXMSGENDCF
	TYPECEVT1CF ICR = 0x01 << 14 //+ TYPECEVT1CF
	TYPECEVT2CF ICR = 0x01 << 15 //+ TYPECEVT2CF
	FRSEVTCF    ICR = 0x01 << 20 //+ FRSEVTCF
)

const (
	TXMSGDISCCFn = 1
	TXMSGSENTCFn = 2
	TXMSGABTCFn  = 3
	HRSTDISCCFn  = 4
	HRSTSENTCFn  = 5
	TXUNDCFn     = 6
	RXORDDETCFn  = 9
	RXHRSTDETCFn = 10
	RXOVRCFn     = 11
	RXMSGENDCFn  = 12
	TYPECEVT1CFn = 14
	TYPECEVT2CFn = 15
	FRSEVTCFn    = 20
)

const (
	TXORDSET TX_ORDSET = 0xFFFFF << 0 //+ TXORDSET
)

const (
	TXORDSETn = 0
)

const (
	TXPAYSZ TX_PAYSZ = 0x3FF << 0 //+ TXPAYSZ
)

const (
	TXPAYSZn = 0
)

const (
	TXDATA TXDR = 0xFF << 0 //+ TXDATA
)

const (
	TXDATAn = 0
)

const (
	RXORDSET      RX_ORDSET = 0x07 << 0 //+ RXORDSET
	RXSOP3OF4     RX_ORDSET = 0x01 << 3 //+ RXSOP3OF4
	RXSOPKINVALID RX_ORDSET = 0x07 << 4 //+ RXSOPKINVALID
)

const (
	RXORDSETn      = 0
	RXSOP3OF4n     = 3
	RXSOPKINVALIDn = 4
)

const (
	RXPAYSZ RX_PAYSZ = 0x3FF << 0 //+ RXPAYSZ
)

const (
	RXPAYSZn = 0
)

const (
	RXDATA RXDR = 0xFF << 0 //+ RXDATA
)

const (
	RXDATAn = 0
)

const (
	RXSOPX1 RX_ORDEXT1 = 0xFFFFF << 0 //+ RXSOPX1
)

const (
	RXSOPX1n = 0
)

const (
	RXSOPX2 RX_ORDEXT2 = 0xFFFFF << 0 //+ RXSOPX2
)

const (
	RXSOPX2n = 0
)
