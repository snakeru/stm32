// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32f407

// Package tim5 provides access to the registers of the TIM peripheral.
//
// Instances:
//  TIM5  TIM5_BASE  APB1  TIM5  General-purpose-timers
// Registers:
//  0x000 32  CR1           control register 1
//  0x004 32  CR2           control register 2
//  0x008 32  SMCR          slave mode control register
//  0x00C 32  DIER          DMA/Interrupt enable register
//  0x010 32  SR            status register
//  0x014 32  EGR           event generation register
//  0x018 32  CCMR1_Output  capture/compare mode register 1 (output mode)
//  0x018 32  CCMR1_Input   capture/compare mode register 1 (input mode)
//  0x01C 32  CCMR2_Output  capture/compare mode register 2 (output mode)
//  0x01C 32  CCMR2_Input   capture/compare mode register 2 (input mode)
//  0x020 32  CCER          capture/compare enable register
//  0x024 32  CNT           counter
//  0x028 32  PSC           prescaler
//  0x02C 32  ARR           auto-reload register
//  0x034 32  CCR1          capture/compare register 1
//  0x038 32  CCR2          capture/compare register 2
//  0x03C 32  CCR3          capture/compare register 3
//  0x040 32  CCR4          capture/compare register 4
//  0x048 32  DCR           DMA control register
//  0x04C 32  DMAR          DMA address for full transfer
//  0x050 32  OR            TIM5 option register
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package tim5

const (
	CEN  CR1 = 0x01 << 0 //+ Counter enable
	UDIS CR1 = 0x01 << 1 //+ Update disable
	URS  CR1 = 0x01 << 2 //+ Update request source
	OPM  CR1 = 0x01 << 3 //+ One-pulse mode
	DIR  CR1 = 0x01 << 4 //+ Direction
	CMS  CR1 = 0x03 << 5 //+ Center-aligned mode selection
	ARPE CR1 = 0x01 << 7 //+ Auto-reload preload enable
	CKD  CR1 = 0x03 << 8 //+ Clock division
)

const (
	CENn  = 0
	UDISn = 1
	URSn  = 2
	OPMn  = 3
	DIRn  = 4
	CMSn  = 5
	ARPEn = 7
	CKDn  = 8
)

const (
	CCDS CR2 = 0x01 << 3 //+ Capture/compare DMA selection
	MMS  CR2 = 0x07 << 4 //+ Master mode selection
	TI1S CR2 = 0x01 << 7 //+ TI1 selection
)

const (
	CCDSn = 3
	MMSn  = 4
	TI1Sn = 7
)

const (
	SMS  SMCR = 0x07 << 0  //+ Slave mode selection
	TS   SMCR = 0x07 << 4  //+ Trigger selection
	MSM  SMCR = 0x01 << 7  //+ Master/Slave mode
	ETF  SMCR = 0x0F << 8  //+ External trigger filter
	ETPS SMCR = 0x03 << 12 //+ External trigger prescaler
	ECE  SMCR = 0x01 << 14 //+ External clock enable
	ETP  SMCR = 0x01 << 15 //+ External trigger polarity
)

const (
	SMSn  = 0
	TSn   = 4
	MSMn  = 7
	ETFn  = 8
	ETPSn = 12
	ECEn  = 14
	ETPn  = 15
)

const (
	UIE   DIER = 0x01 << 0  //+ Update interrupt enable
	CC1IE DIER = 0x01 << 1  //+ Capture/Compare 1 interrupt enable
	CC2IE DIER = 0x01 << 2  //+ Capture/Compare 2 interrupt enable
	CC3IE DIER = 0x01 << 3  //+ Capture/Compare 3 interrupt enable
	CC4IE DIER = 0x01 << 4  //+ Capture/Compare 4 interrupt enable
	TIE   DIER = 0x01 << 6  //+ Trigger interrupt enable
	UDE   DIER = 0x01 << 8  //+ Update DMA request enable
	CC1DE DIER = 0x01 << 9  //+ Capture/Compare 1 DMA request enable
	CC2DE DIER = 0x01 << 10 //+ Capture/Compare 2 DMA request enable
	CC3DE DIER = 0x01 << 11 //+ Capture/Compare 3 DMA request enable
	CC4DE DIER = 0x01 << 12 //+ Capture/Compare 4 DMA request enable
	TDE   DIER = 0x01 << 14 //+ Trigger DMA request enable
)

const (
	UIEn   = 0
	CC1IEn = 1
	CC2IEn = 2
	CC3IEn = 3
	CC4IEn = 4
	TIEn   = 6
	UDEn   = 8
	CC1DEn = 9
	CC2DEn = 10
	CC3DEn = 11
	CC4DEn = 12
	TDEn   = 14
)

const (
	UIF   SR = 0x01 << 0  //+ Update interrupt flag
	CC1IF SR = 0x01 << 1  //+ Capture/compare 1 interrupt flag
	CC2IF SR = 0x01 << 2  //+ Capture/Compare 2 interrupt flag
	CC3IF SR = 0x01 << 3  //+ Capture/Compare 3 interrupt flag
	CC4IF SR = 0x01 << 4  //+ Capture/Compare 4 interrupt flag
	TIF   SR = 0x01 << 6  //+ Trigger interrupt flag
	CC1OF SR = 0x01 << 9  //+ Capture/Compare 1 overcapture flag
	CC2OF SR = 0x01 << 10 //+ Capture/compare 2 overcapture flag
	CC3OF SR = 0x01 << 11 //+ Capture/Compare 3 overcapture flag
	CC4OF SR = 0x01 << 12 //+ Capture/Compare 4 overcapture flag
)

const (
	UIFn   = 0
	CC1IFn = 1
	CC2IFn = 2
	CC3IFn = 3
	CC4IFn = 4
	TIFn   = 6
	CC1OFn = 9
	CC2OFn = 10
	CC3OFn = 11
	CC4OFn = 12
)

const (
	UG   EGR = 0x01 << 0 //+ Update generation
	CC1G EGR = 0x01 << 1 //+ Capture/compare 1 generation
	CC2G EGR = 0x01 << 2 //+ Capture/compare 2 generation
	CC3G EGR = 0x01 << 3 //+ Capture/compare 3 generation
	CC4G EGR = 0x01 << 4 //+ Capture/compare 4 generation
	TG   EGR = 0x01 << 6 //+ Trigger generation
)

const (
	UGn   = 0
	CC1Gn = 1
	CC2Gn = 2
	CC3Gn = 3
	CC4Gn = 4
	TGn   = 6
)

const (
	CC1S  CCMR1_Output = 0x03 << 0  //+ CC1S
	OC1FE CCMR1_Output = 0x01 << 2  //+ OC1FE
	OC1PE CCMR1_Output = 0x01 << 3  //+ OC1PE
	OC1M  CCMR1_Output = 0x07 << 4  //+ OC1M
	OC1CE CCMR1_Output = 0x01 << 7  //+ OC1CE
	CC2S  CCMR1_Output = 0x03 << 8  //+ CC2S
	OC2FE CCMR1_Output = 0x01 << 10 //+ OC2FE
	OC2PE CCMR1_Output = 0x01 << 11 //+ OC2PE
	OC2M  CCMR1_Output = 0x07 << 12 //+ OC2M
	OC2CE CCMR1_Output = 0x01 << 15 //+ OC2CE
)

const (
	CC1Sn  = 0
	OC1FEn = 2
	OC1PEn = 3
	OC1Mn  = 4
	OC1CEn = 7
	CC2Sn  = 8
	OC2FEn = 10
	OC2PEn = 11
	OC2Mn  = 12
	OC2CEn = 15
)

const (
	CC1S   CCMR1_Input = 0x03 << 0  //+ Capture/Compare 1 selection
	ICPCS  CCMR1_Input = 0x03 << 2  //+ Input capture 1 prescaler
	IC1F   CCMR1_Input = 0x0F << 4  //+ Input capture 1 filter
	CC2S   CCMR1_Input = 0x03 << 8  //+ Capture/Compare 2 selection
	IC2PCS CCMR1_Input = 0x03 << 10 //+ Input capture 2 prescaler
	IC2F   CCMR1_Input = 0x0F << 12 //+ Input capture 2 filter
)

const (
	CC1Sn   = 0
	ICPCSn  = 2
	IC1Fn   = 4
	CC2Sn   = 8
	IC2PCSn = 10
	IC2Fn   = 12
)

const (
	CC3S  CCMR2_Output = 0x03 << 0  //+ CC3S
	OC3FE CCMR2_Output = 0x01 << 2  //+ OC3FE
	OC3PE CCMR2_Output = 0x01 << 3  //+ OC3PE
	OC3M  CCMR2_Output = 0x07 << 4  //+ OC3M
	OC3CE CCMR2_Output = 0x01 << 7  //+ OC3CE
	CC4S  CCMR2_Output = 0x03 << 8  //+ CC4S
	OC4FE CCMR2_Output = 0x01 << 10 //+ OC4FE
	OC4PE CCMR2_Output = 0x01 << 11 //+ OC4PE
	OC4M  CCMR2_Output = 0x07 << 12 //+ OC4M
	O24CE CCMR2_Output = 0x01 << 15 //+ O24CE
)

const (
	CC3Sn  = 0
	OC3FEn = 2
	OC3PEn = 3
	OC3Mn  = 4
	OC3CEn = 7
	CC4Sn  = 8
	OC4FEn = 10
	OC4PEn = 11
	OC4Mn  = 12
	O24CEn = 15
)

const (
	CC3S   CCMR2_Input = 0x03 << 0  //+ Capture/compare 3 selection
	IC3PSC CCMR2_Input = 0x03 << 2  //+ Input capture 3 prescaler
	IC3F   CCMR2_Input = 0x0F << 4  //+ Input capture 3 filter
	CC4S   CCMR2_Input = 0x03 << 8  //+ Capture/Compare 4 selection
	IC4PSC CCMR2_Input = 0x03 << 10 //+ Input capture 4 prescaler
	IC4F   CCMR2_Input = 0x0F << 12 //+ Input capture 4 filter
)

const (
	CC3Sn   = 0
	IC3PSCn = 2
	IC3Fn   = 4
	CC4Sn   = 8
	IC4PSCn = 10
	IC4Fn   = 12
)

const (
	CC1E  CCER = 0x01 << 0  //+ Capture/Compare 1 output enable
	CC1P  CCER = 0x01 << 1  //+ Capture/Compare 1 output Polarity
	CC1NP CCER = 0x01 << 3  //+ Capture/Compare 1 output Polarity
	CC2E  CCER = 0x01 << 4  //+ Capture/Compare 2 output enable
	CC2P  CCER = 0x01 << 5  //+ Capture/Compare 2 output Polarity
	CC2NP CCER = 0x01 << 7  //+ Capture/Compare 2 output Polarity
	CC3E  CCER = 0x01 << 8  //+ Capture/Compare 3 output enable
	CC3P  CCER = 0x01 << 9  //+ Capture/Compare 3 output Polarity
	CC3NP CCER = 0x01 << 11 //+ Capture/Compare 3 output Polarity
	CC4E  CCER = 0x01 << 12 //+ Capture/Compare 4 output enable
	CC4P  CCER = 0x01 << 13 //+ Capture/Compare 3 output Polarity
	CC4NP CCER = 0x01 << 15 //+ Capture/Compare 4 output Polarity
)

const (
	CC1En  = 0
	CC1Pn  = 1
	CC1NPn = 3
	CC2En  = 4
	CC2Pn  = 5
	CC2NPn = 7
	CC3En  = 8
	CC3Pn  = 9
	CC3NPn = 11
	CC4En  = 12
	CC4Pn  = 13
	CC4NPn = 15
)

const (
	CNT_L CNT = 0xFFFF << 0  //+ Low counter value
	CNT_H CNT = 0xFFFF << 16 //+ High counter value
)

const (
	CNT_Ln = 0
	CNT_Hn = 16
)

const (
	PSC PSC = 0xFFFF << 0 //+ Prescaler value
)

const (
	PSCn = 0
)

const (
	ARR_L ARR = 0xFFFF << 0  //+ Low Auto-reload value
	ARR_H ARR = 0xFFFF << 16 //+ High Auto-reload value
)

const (
	ARR_Ln = 0
	ARR_Hn = 16
)

const (
	CCR1_L CCR1 = 0xFFFF << 0  //+ Low Capture/Compare 1 value
	CCR1_H CCR1 = 0xFFFF << 16 //+ High Capture/Compare 1 value
)

const (
	CCR1_Ln = 0
	CCR1_Hn = 16
)

const (
	CCR2_L CCR2 = 0xFFFF << 0  //+ Low Capture/Compare 2 value
	CCR2_H CCR2 = 0xFFFF << 16 //+ High Capture/Compare 2 value
)

const (
	CCR2_Ln = 0
	CCR2_Hn = 16
)

const (
	CCR3_L CCR3 = 0xFFFF << 0  //+ Low Capture/Compare value
	CCR3_H CCR3 = 0xFFFF << 16 //+ High Capture/Compare value
)

const (
	CCR3_Ln = 0
	CCR3_Hn = 16
)

const (
	CCR4_L CCR4 = 0xFFFF << 0  //+ Low Capture/Compare value
	CCR4_H CCR4 = 0xFFFF << 16 //+ High Capture/Compare value
)

const (
	CCR4_Ln = 0
	CCR4_Hn = 16
)

const (
	DBA DCR = 0x1F << 0 //+ DMA base address
	DBL DCR = 0x1F << 8 //+ DMA burst length
)

const (
	DBAn = 0
	DBLn = 8
)

const (
	DMAB DMAR = 0xFFFF << 0 //+ DMA register for burst accesses
)

const (
	DMABn = 0
)

const (
	IT4_RMP OR = 0x03 << 6 //+ Timer Input 4 remap
)

const (
	IT4_RMPn = 6
)
