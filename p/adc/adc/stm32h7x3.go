// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32h7x3

// Package adc provides access to the registers of the ADC peripheral.
//
// Instances:
//  ADC1  ADC1_BASE  AHB1  ADC3+  Analog to Digital Converter
//  ADC2  ADC2_BASE  AHB1  ADC3+  Analog to Digital Converter
//  ADC3  ADC3_BASE  AHB4  ADC3+  Analog to Digital Converter
// Registers:
//  0x000 32  ISR       ADC interrupt and status register
//  0x004 32  IER       ADC interrupt enable register
//  0x008 32  CR        ADC control register
//  0x00C 32  CFGR      ADC configuration register 1
//  0x010 32  CFGR2     ADC configuration register 2
//  0x014 32  SMPR1     ADC sampling time register 1
//  0x018 32  SMPR2     ADC sampling time register 2
//  0x01C 32  PCSEL     ADC pre channel selection register
//  0x020 32  LTR1      ADC analog watchdog 1 threshold register
//  0x024 32  LHTR1     ADC analog watchdog 2 threshold register
//  0x030 32  SQR1      ADC group regular sequencer ranks register 1
//  0x034 32  SQR2      ADC group regular sequencer ranks register 2
//  0x038 32  SQR3      ADC group regular sequencer ranks register 3
//  0x03C 32  SQR4      ADC group regular sequencer ranks register 4
//  0x040 32  DR        ADC group regular conversion data register
//  0x04C 32  JSQR      ADC group injected sequencer register
//  0x060 32  OFR1      ADC offset number 1 register
//  0x064 32  OFR2      ADC offset number 2 register
//  0x068 32  OFR3      ADC offset number 3 register
//  0x06C 32  OFR4      ADC offset number 4 register
//  0x080 32  JDR1      ADC group injected sequencer rank 1 register
//  0x084 32  JDR2      ADC group injected sequencer rank 2 register
//  0x088 32  JDR3      ADC group injected sequencer rank 3 register
//  0x08C 32  JDR4      ADC group injected sequencer rank 4 register
//  0x0A0 32  AWD2CR    ADC analog watchdog 2 configuration register
//  0x0A4 32  AWD3CR    ADC analog watchdog 3 configuration register
//  0x0B0 32  LTR2      ADC watchdog lower threshold register 2
//  0x0B4 32  HTR2      ADC watchdog higher threshold register 2
//  0x0B8 32  LTR3      ADC watchdog lower threshold register 3
//  0x0BC 32  HTR3      ADC watchdog higher threshold register 3
//  0x0C0 32  DIFSEL    ADC channel differential or single-ended mode selection register
//  0x0C4 32  CALFACT   ADC calibration factors register
//  0x0C8 32  CALFACT2  ADC Calibration Factor register 2
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package adc

const (
	ADRDY ISR = 0x01 << 0  //+ ADC ready flag
	EOSMP ISR = 0x01 << 1  //+ ADC group regular end of sampling flag
	EOC   ISR = 0x01 << 2  //+ ADC group regular end of unitary conversion flag
	EOS   ISR = 0x01 << 3  //+ ADC group regular end of sequence conversions flag
	OVR   ISR = 0x01 << 4  //+ ADC group regular overrun flag
	JEOC  ISR = 0x01 << 5  //+ ADC group injected end of unitary conversion flag
	JEOS  ISR = 0x01 << 6  //+ ADC group injected end of sequence conversions flag
	AWD1  ISR = 0x01 << 7  //+ ADC analog watchdog 1 flag
	AWD2  ISR = 0x01 << 8  //+ ADC analog watchdog 2 flag
	AWD3  ISR = 0x01 << 9  //+ ADC analog watchdog 3 flag
	JQOVF ISR = 0x01 << 10 //+ ADC group injected contexts queue overflow flag
)

const (
	ADRDYn = 0
	EOSMPn = 1
	EOCn   = 2
	EOSn   = 3
	OVRn   = 4
	JEOCn  = 5
	JEOSn  = 6
	AWD1n  = 7
	AWD2n  = 8
	AWD3n  = 9
	JQOVFn = 10
)

const (
	ADRDYIE IER = 0x01 << 0  //+ ADC ready interrupt
	EOSMPIE IER = 0x01 << 1  //+ ADC group regular end of sampling interrupt
	EOCIE   IER = 0x01 << 2  //+ ADC group regular end of unitary conversion interrupt
	EOSIE   IER = 0x01 << 3  //+ ADC group regular end of sequence conversions interrupt
	OVRIE   IER = 0x01 << 4  //+ ADC group regular overrun interrupt
	JEOCIE  IER = 0x01 << 5  //+ ADC group injected end of unitary conversion interrupt
	JEOSIE  IER = 0x01 << 6  //+ ADC group injected end of sequence conversions interrupt
	AWD1IE  IER = 0x01 << 7  //+ ADC analog watchdog 1 interrupt
	AWD2IE  IER = 0x01 << 8  //+ ADC analog watchdog 2 interrupt
	AWD3IE  IER = 0x01 << 9  //+ ADC analog watchdog 3 interrupt
	JQOVFIE IER = 0x01 << 10 //+ ADC group injected contexts queue overflow interrupt
)

const (
	ADRDYIEn = 0
	EOSMPIEn = 1
	EOCIEn   = 2
	EOSIEn   = 3
	OVRIEn   = 4
	JEOCIEn  = 5
	JEOSIEn  = 6
	AWD1IEn  = 7
	AWD2IEn  = 8
	AWD3IEn  = 9
	JQOVFIEn = 10
)

const (
	ADEN        CR = 0x01 << 0  //+ ADC enable
	ADDIS       CR = 0x01 << 1  //+ ADC disable
	ADSTART     CR = 0x01 << 2  //+ ADC group regular conversion start
	JADSTART    CR = 0x01 << 3  //+ ADC group injected conversion start
	ADSTP       CR = 0x01 << 4  //+ ADC group regular conversion stop
	JADSTP      CR = 0x01 << 5  //+ ADC group injected conversion stop
	BOOST       CR = 0x01 << 8  //+ Boost mode control
	ADCALLIN    CR = 0x01 << 16 //+ Linearity calibration
	LINCALRDYW1 CR = 0x01 << 22 //+ Linearity calibration ready Word 1
	LINCALRDYW2 CR = 0x01 << 23 //+ Linearity calibration ready Word 2
	LINCALRDYW3 CR = 0x01 << 24 //+ Linearity calibration ready Word 3
	LINCALRDYW4 CR = 0x01 << 25 //+ Linearity calibration ready Word 4
	LINCALRDYW5 CR = 0x01 << 26 //+ Linearity calibration ready Word 5
	LINCALRDYW6 CR = 0x01 << 27 //+ Linearity calibration ready Word 6
	ADVREGEN    CR = 0x01 << 28 //+ ADC voltage regulator enable
	DEEPPWD     CR = 0x01 << 29 //+ ADC deep power down enable
	ADCALDIF    CR = 0x01 << 30 //+ ADC differential mode for calibration
	ADCAL       CR = 0x01 << 31 //+ ADC calibration
)

const (
	ADENn        = 0
	ADDISn       = 1
	ADSTARTn     = 2
	JADSTARTn    = 3
	ADSTPn       = 4
	JADSTPn      = 5
	BOOSTn       = 8
	ADCALLINn    = 16
	LINCALRDYW1n = 22
	LINCALRDYW2n = 23
	LINCALRDYW3n = 24
	LINCALRDYW4n = 25
	LINCALRDYW5n = 26
	LINCALRDYW6n = 27
	ADVREGENn    = 28
	DEEPPWDn     = 29
	ADCALDIFn    = 30
	ADCALn       = 31
)

const (
	DMNGT    CFGR = 0x03 << 0  //+ ADC DMA transfer enable
	RES      CFGR = 0x07 << 2  //+ ADC data resolution
	EXTSEL   CFGR = 0x1F << 5  //+ ADC group regular external trigger source
	EXTEN    CFGR = 0x03 << 10 //+ ADC group regular external trigger polarity
	OVRMOD   CFGR = 0x01 << 12 //+ ADC group regular overrun configuration
	CONT     CFGR = 0x01 << 13 //+ ADC group regular continuous conversion mode
	AUTDLY   CFGR = 0x01 << 14 //+ ADC low power auto wait
	DISCEN   CFGR = 0x01 << 16 //+ ADC group regular sequencer discontinuous mode
	DISCNUM  CFGR = 0x07 << 17 //+ ADC group regular sequencer discontinuous number of ranks
	JDISCEN  CFGR = 0x01 << 20 //+ ADC group injected sequencer discontinuous mode
	JQM      CFGR = 0x01 << 21 //+ ADC group injected contexts queue mode
	AWD1SGL  CFGR = 0x01 << 22 //+ ADC analog watchdog 1 monitoring a single channel or all channels
	AWD1EN   CFGR = 0x01 << 23 //+ ADC analog watchdog 1 enable on scope ADC group regular
	JAWD1EN  CFGR = 0x01 << 24 //+ ADC analog watchdog 1 enable on scope ADC group injected
	JAUTO    CFGR = 0x01 << 25 //+ ADC group injected automatic trigger mode
	AWDCH1CH CFGR = 0x1F << 26 //+ ADC analog watchdog 1 monitored channel selection
	JQDIS    CFGR = 0x01 << 31 //+ ADC group injected contexts queue disable
)

const (
	DMNGTn    = 0
	RESn      = 2
	EXTSELn   = 5
	EXTENn    = 10
	OVRMODn   = 12
	CONTn     = 13
	AUTDLYn   = 14
	DISCENn   = 16
	DISCNUMn  = 17
	JDISCENn  = 20
	JQMn      = 21
	AWD1SGLn  = 22
	AWD1ENn   = 23
	JAWD1ENn  = 24
	JAUTOn    = 25
	AWDCH1CHn = 26
	JQDISn    = 31
)

const (
	ROVSE   CFGR2 = 0x01 << 0   //+ ADC oversampler enable on scope ADC group regular
	JOVSE   CFGR2 = 0x01 << 1   //+ ADC oversampler enable on scope ADC group injected
	OVSS    CFGR2 = 0x0F << 5   //+ ADC oversampling shift
	TROVS   CFGR2 = 0x01 << 9   //+ ADC oversampling discontinuous mode (triggered mode) for ADC group regular
	ROVSM   CFGR2 = 0x01 << 10  //+ Regular Oversampling mode
	RSHIFT1 CFGR2 = 0x01 << 11  //+ Right-shift data after Offset 1 correction
	RSHIFT2 CFGR2 = 0x01 << 12  //+ Right-shift data after Offset 2 correction
	RSHIFT3 CFGR2 = 0x01 << 13  //+ Right-shift data after Offset 3 correction
	RSHIFT4 CFGR2 = 0x01 << 14  //+ Right-shift data after Offset 4 correction
	OSR     CFGR2 = 0x3FF << 16 //+ Oversampling ratio
	LSHIFT  CFGR2 = 0x0F << 28  //+ Left shift factor
)

const (
	ROVSEn   = 0
	JOVSEn   = 1
	OVSSn    = 5
	TROVSn   = 9
	ROVSMn   = 10
	RSHIFT1n = 11
	RSHIFT2n = 12
	RSHIFT3n = 13
	RSHIFT4n = 14
	OSRn     = 16
	LSHIFTn  = 28
)

const (
	SMP1 SMPR1 = 0x07 << 3  //+ ADC channel 1 sampling time selection
	SMP2 SMPR1 = 0x07 << 6  //+ ADC channel 2 sampling time selection
	SMP3 SMPR1 = 0x07 << 9  //+ ADC channel 3 sampling time selection
	SMP4 SMPR1 = 0x07 << 12 //+ ADC channel 4 sampling time selection
	SMP5 SMPR1 = 0x07 << 15 //+ ADC channel 5 sampling time selection
	SMP6 SMPR1 = 0x07 << 18 //+ ADC channel 6 sampling time selection
	SMP7 SMPR1 = 0x07 << 21 //+ ADC channel 7 sampling time selection
	SMP8 SMPR1 = 0x07 << 24 //+ ADC channel 8 sampling time selection
	SMP9 SMPR1 = 0x07 << 27 //+ ADC channel 9 sampling time selection
)

const (
	SMP1n = 3
	SMP2n = 6
	SMP3n = 9
	SMP4n = 12
	SMP5n = 15
	SMP6n = 18
	SMP7n = 21
	SMP8n = 24
	SMP9n = 27
)

const (
	SMP10 SMPR2 = 0x07 << 0  //+ ADC channel 10 sampling time selection
	SMP11 SMPR2 = 0x07 << 3  //+ ADC channel 11 sampling time selection
	SMP12 SMPR2 = 0x07 << 6  //+ ADC channel 12 sampling time selection
	SMP13 SMPR2 = 0x07 << 9  //+ ADC channel 13 sampling time selection
	SMP14 SMPR2 = 0x07 << 12 //+ ADC channel 14 sampling time selection
	SMP15 SMPR2 = 0x07 << 15 //+ ADC channel 15 sampling time selection
	SMP16 SMPR2 = 0x07 << 18 //+ ADC channel 16 sampling time selection
	SMP17 SMPR2 = 0x07 << 21 //+ ADC channel 17 sampling time selection
	SMP18 SMPR2 = 0x07 << 24 //+ ADC channel 18 sampling time selection
	SMP19 SMPR2 = 0x07 << 27 //+ ADC channel 18 sampling time selection
)

const (
	SMP10n = 0
	SMP11n = 3
	SMP12n = 6
	SMP13n = 9
	SMP14n = 12
	SMP15n = 15
	SMP16n = 18
	SMP17n = 21
	SMP18n = 24
	SMP19n = 27
)

const (
	PCSEL PCSEL = 0xFFFFF << 0 //+ Channel x (VINP[i]) pre selection
)

const (
	PCSELn = 0
)

const (
	LTR1 LTR1 = 0x3FFFFFF << 0 //+ ADC analog watchdog 1 threshold low
)

const (
	LTR1n = 0
)

const (
	LHTR1 LHTR1 = 0x3FFFFFF << 0 //+ ADC analog watchdog 2 threshold low
)

const (
	LHTR1n = 0
)

const (
	L3  SQR1 = 0x0F << 0  //+ L3
	SQ1 SQR1 = 0x1F << 6  //+ ADC group regular sequencer rank 1
	SQ2 SQR1 = 0x1F << 12 //+ ADC group regular sequencer rank 2
	SQ3 SQR1 = 0x1F << 18 //+ ADC group regular sequencer rank 3
	SQ4 SQR1 = 0x1F << 24 //+ ADC group regular sequencer rank 4
)

const (
	L3n  = 0
	SQ1n = 6
	SQ2n = 12
	SQ3n = 18
	SQ4n = 24
)

const (
	SQ5 SQR2 = 0x1F << 0  //+ ADC group regular sequencer rank 5
	SQ6 SQR2 = 0x1F << 6  //+ ADC group regular sequencer rank 6
	SQ7 SQR2 = 0x1F << 12 //+ ADC group regular sequencer rank 7
	SQ8 SQR2 = 0x1F << 18 //+ ADC group regular sequencer rank 8
	SQ9 SQR2 = 0x1F << 24 //+ ADC group regular sequencer rank 9
)

const (
	SQ5n = 0
	SQ6n = 6
	SQ7n = 12
	SQ8n = 18
	SQ9n = 24
)

const (
	SQ10 SQR3 = 0x1F << 0  //+ ADC group regular sequencer rank 10
	SQ11 SQR3 = 0x1F << 6  //+ ADC group regular sequencer rank 11
	SQ12 SQR3 = 0x1F << 12 //+ ADC group regular sequencer rank 12
	SQ13 SQR3 = 0x1F << 18 //+ ADC group regular sequencer rank 13
	SQ14 SQR3 = 0x1F << 24 //+ ADC group regular sequencer rank 14
)

const (
	SQ10n = 0
	SQ11n = 6
	SQ12n = 12
	SQ13n = 18
	SQ14n = 24
)

const (
	SQ15 SQR4 = 0x1F << 0 //+ ADC group regular sequencer rank 15
	SQ16 SQR4 = 0x1F << 6 //+ ADC group regular sequencer rank 16
)

const (
	SQ15n = 0
	SQ16n = 6
)

const (
	RDATA DR = 0xFFFF << 0 //+ ADC group regular conversion data
)

const (
	RDATAn = 0
)

const (
	JL      JSQR = 0x03 << 0  //+ ADC group injected sequencer scan length
	JEXTSEL JSQR = 0x1F << 2  //+ ADC group injected external trigger source
	JEXTEN  JSQR = 0x03 << 7  //+ ADC group injected external trigger polarity
	JSQ1    JSQR = 0x1F << 9  //+ ADC group injected sequencer rank 1
	JSQ2    JSQR = 0x1F << 15 //+ ADC group injected sequencer rank 2
	JSQ3    JSQR = 0x1F << 21 //+ ADC group injected sequencer rank 3
	JSQ4    JSQR = 0x1F << 27 //+ ADC group injected sequencer rank 4
)

const (
	JLn      = 0
	JEXTSELn = 2
	JEXTENn  = 7
	JSQ1n    = 9
	JSQ2n    = 15
	JSQ3n    = 21
	JSQ4n    = 27
)

const (
	OFFSET1    OFR1 = 0x3FFFFFF << 0 //+ ADC offset number 1 offset level
	OFFSET1_CH OFR1 = 0x1F << 26     //+ ADC offset number 1 channel selection
	SSATE      OFR1 = 0x01 << 31     //+ ADC offset number 1 enable
)

const (
	OFFSET1n    = 0
	OFFSET1_CHn = 26
	SSATEn      = 31
)

const (
	OFFSET1    OFR2 = 0x3FFFFFF << 0 //+ ADC offset number 1 offset level
	OFFSET1_CH OFR2 = 0x1F << 26     //+ ADC offset number 1 channel selection
	SSATE      OFR2 = 0x01 << 31     //+ ADC offset number 1 enable
)

const (
	OFFSET1n    = 0
	OFFSET1_CHn = 26
	SSATEn      = 31
)

const (
	OFFSET1    OFR3 = 0x3FFFFFF << 0 //+ ADC offset number 1 offset level
	OFFSET1_CH OFR3 = 0x1F << 26     //+ ADC offset number 1 channel selection
	SSATE      OFR3 = 0x01 << 31     //+ ADC offset number 1 enable
)

const (
	OFFSET1n    = 0
	OFFSET1_CHn = 26
	SSATEn      = 31
)

const (
	OFFSET1    OFR4 = 0x3FFFFFF << 0 //+ ADC offset number 1 offset level
	OFFSET1_CH OFR4 = 0x1F << 26     //+ ADC offset number 1 channel selection
	SSATE      OFR4 = 0x01 << 31     //+ ADC offset number 1 enable
)

const (
	OFFSET1n    = 0
	OFFSET1_CHn = 26
	SSATEn      = 31
)

const (
	JDATA1 JDR1 = 0xFFFFFFFF << 0 //+ ADC group injected sequencer rank 1 conversion data
)

const (
	JDATA1n = 0
)

const (
	JDATA2 JDR2 = 0xFFFFFFFF << 0 //+ ADC group injected sequencer rank 2 conversion data
)

const (
	JDATA2n = 0
)

const (
	JDATA3 JDR3 = 0xFFFFFFFF << 0 //+ ADC group injected sequencer rank 3 conversion data
)

const (
	JDATA3n = 0
)

const (
	JDATA4 JDR4 = 0xFFFFFFFF << 0 //+ ADC group injected sequencer rank 4 conversion data
)

const (
	JDATA4n = 0
)

const (
	AWD2CH AWD2CR = 0xFFFFF << 0 //+ ADC analog watchdog 2 monitored channel selection
)

const (
	AWD2CHn = 0
)

const (
	AWD3CH AWD3CR = 0xFFFFF << 1 //+ ADC analog watchdog 3 monitored channel selection
)

const (
	AWD3CHn = 1
)

const (
	LTR2 LTR2 = 0x3FFFFFF << 0 //+ Analog watchdog 2 lower threshold
)

const (
	LTR2n = 0
)

const (
	HTR2 HTR2 = 0x3FFFFFF << 0 //+ Analog watchdog 2 higher threshold
)

const (
	HTR2n = 0
)

const (
	LTR3 LTR3 = 0x3FFFFFF << 0 //+ Analog watchdog 3 lower threshold
)

const (
	LTR3n = 0
)

const (
	HTR3 HTR3 = 0x3FFFFFF << 0 //+ Analog watchdog 3 higher threshold
)

const (
	HTR3n = 0
)

const (
	DIFSEL DIFSEL = 0xFFFFF << 0 //+ ADC channel differential or single-ended mode for channel
)

const (
	DIFSELn = 0
)

const (
	CALFACT_S CALFACT = 0x7FF << 0  //+ ADC calibration factor in single-ended mode
	CALFACT_D CALFACT = 0x7FF << 16 //+ ADC calibration factor in differential mode
)

const (
	CALFACT_Sn = 0
	CALFACT_Dn = 16
)

const (
	LINCALFACT CALFACT2 = 0x3FFFFFFF << 0 //+ Linearity Calibration Factor
)

const (
	LINCALFACTn = 0
)
