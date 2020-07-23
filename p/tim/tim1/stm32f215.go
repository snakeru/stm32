// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32f215

// Package tim1 provides access to the registers of the TIM peripheral.
//
// Instances:
//  TIM1  TIM1_BASE  APB2  TIM1_BRK_TIM9*,TIM1_UP_TIM10*,TIM1_TRG_COM_TIM11*,TIM1_CC   Advanced-timers
//  TIM8  TIM8_BASE  APB2  TIM8_BRK_TIM12*,TIM8_UP_TIM13*,TIM8_TRG_COM_TIM14*,TIM8_CC  Advanced-timers
// Registers:
//  0x000 32  CR1           control register 1
//  0x004 32  CR2           control register 2
//  0x008 32  SMCR          slave mode control register
//  0x00C 32  DIER          DMA/Interrupt enable register
//  0x010 32  SR            status register
//  0x014 32  EGR           event generation register
//  0x018 32  CCMR1_Output  capture/compare mode register (output mode)
//  0x018 32  CCMR1_Input   capture/compare mode register 1 (input mode)
//  0x01C 32  CCMR2_Output  capture/compare mode register (output mode)
//  0x01C 32  CCMR2_Input   capture/compare mode register 2 (input mode)
//  0x020 32  CCER          capture/compare enable register
//  0x024 32  CNT           counter
//  0x028 32  PSC           prescaler
//  0x02C 32  ARR           auto-reload register
//  0x030 32  RCR           repetition counter register
//  0x034 32  CCR1          capture/compare register 1
//  0x038 32  CCR2          capture/compare register 2
//  0x03C 32  CCR3          capture/compare register 3
//  0x040 32  CCR4          capture/compare register 4
//  0x044 32  BDTR          break and dead-time register
//  0x048 32  DCR           DMA control register
//  0x04C 32  DMAR          DMA address for full transfer
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package tim1

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
	CCPC  CR2 = 0x01 << 0  //+ Capture/compare preloaded control
	CCUS  CR2 = 0x01 << 2  //+ Capture/compare control update selection
	CCDS  CR2 = 0x01 << 3  //+ Capture/compare DMA selection
	MMS   CR2 = 0x07 << 4  //+ Master mode selection
	TI1S  CR2 = 0x01 << 7  //+ TI1 selection
	OIS1  CR2 = 0x01 << 8  //+ Output Idle state 1
	OIS1N CR2 = 0x01 << 9  //+ Output Idle state 1
	OIS2  CR2 = 0x01 << 10 //+ Output Idle state 2
	OIS2N CR2 = 0x01 << 11 //+ Output Idle state 2
	OIS3  CR2 = 0x01 << 12 //+ Output Idle state 3
	OIS3N CR2 = 0x01 << 13 //+ Output Idle state 3
	OIS4  CR2 = 0x01 << 14 //+ Output Idle state 4
)

const (
	CCPCn  = 0
	CCUSn  = 2
	CCDSn  = 3
	MMSn   = 4
	TI1Sn  = 7
	OIS1n  = 8
	OIS1Nn = 9
	OIS2n  = 10
	OIS2Nn = 11
	OIS3n  = 12
	OIS3Nn = 13
	OIS4n  = 14
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
	COMIE DIER = 0x01 << 5  //+ COM interrupt enable
	TIE   DIER = 0x01 << 6  //+ Trigger interrupt enable
	BIE   DIER = 0x01 << 7  //+ Break interrupt enable
	UDE   DIER = 0x01 << 8  //+ Update DMA request enable
	CC1DE DIER = 0x01 << 9  //+ Capture/Compare 1 DMA request enable
	CC2DE DIER = 0x01 << 10 //+ Capture/Compare 2 DMA request enable
	CC3DE DIER = 0x01 << 11 //+ Capture/Compare 3 DMA request enable
	CC4DE DIER = 0x01 << 12 //+ Capture/Compare 4 DMA request enable
	COMDE DIER = 0x01 << 13 //+ COM DMA request enable
	TDE   DIER = 0x01 << 14 //+ Trigger DMA request enable
)

const (
	UIEn   = 0
	CC1IEn = 1
	CC2IEn = 2
	CC3IEn = 3
	CC4IEn = 4
	COMIEn = 5
	TIEn   = 6
	BIEn   = 7
	UDEn   = 8
	CC1DEn = 9
	CC2DEn = 10
	CC3DEn = 11
	CC4DEn = 12
	COMDEn = 13
	TDEn   = 14
)

const (
	UIF   SR = 0x01 << 0  //+ Update interrupt flag
	CC1IF SR = 0x01 << 1  //+ Capture/compare 1 interrupt flag
	CC2IF SR = 0x01 << 2  //+ Capture/Compare 2 interrupt flag
	CC3IF SR = 0x01 << 3  //+ Capture/Compare 3 interrupt flag
	CC4IF SR = 0x01 << 4  //+ Capture/Compare 4 interrupt flag
	COMIF SR = 0x01 << 5  //+ COM interrupt flag
	TIF   SR = 0x01 << 6  //+ Trigger interrupt flag
	BIF   SR = 0x01 << 7  //+ Break interrupt flag
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
	COMIFn = 5
	TIFn   = 6
	BIFn   = 7
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
	COMG EGR = 0x01 << 5 //+ Capture/Compare control update generation
	TG   EGR = 0x01 << 6 //+ Trigger generation
	BG   EGR = 0x01 << 7 //+ Break generation
)

const (
	UGn   = 0
	CC1Gn = 1
	CC2Gn = 2
	CC3Gn = 3
	CC4Gn = 4
	COMGn = 5
	TGn   = 6
	BGn   = 7
)

const (
	CC1S  CCMR1_Output = 0x03 << 0  //+ Capture/Compare 1 selection
	OC1FE CCMR1_Output = 0x01 << 2  //+ Output Compare 1 fast enable
	OC1PE CCMR1_Output = 0x01 << 3  //+ Output Compare 1 preload enable
	OC1M  CCMR1_Output = 0x07 << 4  //+ Output Compare 1 mode
	OC1CE CCMR1_Output = 0x01 << 7  //+ Output Compare 1 clear enable
	CC2S  CCMR1_Output = 0x03 << 8  //+ Capture/Compare 2 selection
	OC2FE CCMR1_Output = 0x01 << 10 //+ Output Compare 2 fast enable
	OC2PE CCMR1_Output = 0x01 << 11 //+ Output Compare 2 preload enable
	OC2M  CCMR1_Output = 0x07 << 12 //+ Output Compare 2 mode
	OC2CE CCMR1_Output = 0x01 << 15 //+ Output Compare 2 clear enable
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
	CC3S  CCMR2_Output = 0x03 << 0  //+ Capture/Compare 3 selection
	OC3FE CCMR2_Output = 0x01 << 2  //+ Output compare 3 fast enable
	OC3PE CCMR2_Output = 0x01 << 3  //+ Output compare 3 preload enable
	OC3M  CCMR2_Output = 0x07 << 4  //+ Output compare 3 mode
	OC3CE CCMR2_Output = 0x01 << 7  //+ Output compare 3 clear enable
	CC4S  CCMR2_Output = 0x03 << 8  //+ Capture/Compare 4 selection
	OC4FE CCMR2_Output = 0x01 << 10 //+ Output compare 4 fast enable
	OC4PE CCMR2_Output = 0x01 << 11 //+ Output compare 4 preload enable
	OC4M  CCMR2_Output = 0x07 << 12 //+ Output compare 4 mode
	OC4CE CCMR2_Output = 0x01 << 15 //+ Output compare 4 clear enable
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
	OC4CEn = 15
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
	CC1NE CCER = 0x01 << 2  //+ Capture/Compare 1 complementary output enable
	CC1NP CCER = 0x01 << 3  //+ Capture/Compare 1 output Polarity
	CC2E  CCER = 0x01 << 4  //+ Capture/Compare 2 output enable
	CC2P  CCER = 0x01 << 5  //+ Capture/Compare 2 output Polarity
	CC2NE CCER = 0x01 << 6  //+ Capture/Compare 2 complementary output enable
	CC2NP CCER = 0x01 << 7  //+ Capture/Compare 2 output Polarity
	CC3E  CCER = 0x01 << 8  //+ Capture/Compare 3 output enable
	CC3P  CCER = 0x01 << 9  //+ Capture/Compare 3 output Polarity
	CC3NE CCER = 0x01 << 10 //+ Capture/Compare 3 complementary output enable
	CC3NP CCER = 0x01 << 11 //+ Capture/Compare 3 output Polarity
	CC4E  CCER = 0x01 << 12 //+ Capture/Compare 4 output enable
	CC4P  CCER = 0x01 << 13 //+ Capture/Compare 3 output Polarity
)

const (
	CC1En  = 0
	CC1Pn  = 1
	CC1NEn = 2
	CC1NPn = 3
	CC2En  = 4
	CC2Pn  = 5
	CC2NEn = 6
	CC2NPn = 7
	CC3En  = 8
	CC3Pn  = 9
	CC3NEn = 10
	CC3NPn = 11
	CC4En  = 12
	CC4Pn  = 13
)

const (
	CNT CNT = 0xFFFF << 0 //+ counter value
)

const (
	CNTn = 0
)

const (
	PSC PSC = 0xFFFF << 0 //+ Prescaler value
)

const (
	PSCn = 0
)

const (
	ARR ARR = 0xFFFF << 0 //+ Auto-reload value
)

const (
	ARRn = 0
)

const (
	REP RCR = 0xFF << 0 //+ Repetition counter value
)

const (
	REPn = 0
)

const (
	CCR1 CCR1 = 0xFFFF << 0 //+ Capture/Compare 1 value
)

const (
	CCR1n = 0
)

const (
	CCR2 CCR2 = 0xFFFF << 0 //+ Capture/Compare 2 value
)

const (
	CCR2n = 0
)

const (
	CCR3 CCR3 = 0xFFFF << 0 //+ Capture/Compare value
)

const (
	CCR3n = 0
)

const (
	CCR4 CCR4 = 0xFFFF << 0 //+ Capture/Compare value
)

const (
	CCR4n = 0
)

const (
	DTG  BDTR = 0xFF << 0  //+ Dead-time generator setup
	LOCK BDTR = 0x03 << 8  //+ Lock configuration
	OSSI BDTR = 0x01 << 10 //+ Off-state selection for Idle mode
	OSSR BDTR = 0x01 << 11 //+ Off-state selection for Run mode
	BKE  BDTR = 0x01 << 12 //+ Break enable
	BKP  BDTR = 0x01 << 13 //+ Break polarity
	AOE  BDTR = 0x01 << 14 //+ Automatic output enable
	MOE  BDTR = 0x01 << 15 //+ Main output enable
)

const (
	DTGn  = 0
	LOCKn = 8
	OSSIn = 10
	OSSRn = 11
	BKEn  = 12
	BKPn  = 13
	AOEn  = 14
	MOEn  = 15
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
