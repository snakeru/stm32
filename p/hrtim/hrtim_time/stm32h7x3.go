// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32h7x3

// Package hrtim_time provides access to the registers of the HRTIM_TIME peripheral.
//
// Instances:
//  HRTIM_TIME  HRTIM_TIME_BASE  -  HRTIM1_TIMD  High Resolution Timer: TIME
// Registers:
//  0x000 32  TIMECR     Timerx Control Register
//  0x004 32  TIMEISR    Timerx Interrupt Status Register
//  0x008 32  TIMEICR    Timerx Interrupt Clear Register
//  0x00C 32  TIMEDIER5  TIMxDIER5
//  0x010 32  CNTER      Timerx Counter Register
//  0x014 32  PERER      Timerx Period Register
//  0x018 32  REPER      Timerx Repetition Register
//  0x01C 32  CMP1ER     Timerx Compare 1 Register
//  0x020 32  CMP1CER    Timerx Compare 1 Compound Register
//  0x024 32  CMP2ER     Timerx Compare 2 Register
//  0x028 32  CMP3ER     Timerx Compare 3 Register
//  0x02C 32  CMP4ER     Timerx Compare 4 Register
//  0x030 32  CPT1ER     Timerx Capture 1 Register
//  0x034 32  CPT2ER     Timerx Capture 2 Register
//  0x038 32  DTER       Timerx Deadtime Register
//  0x03C 32  SETE1R     Timerx Output1 Set Register
//  0x040 32  RSTE1R     Timerx Output1 Reset Register
//  0x044 32  SETE2R     Timerx Output2 Set Register
//  0x048 32  RSTE2R     Timerx Output2 Reset Register
//  0x04C 32  EEFER1     Timerx External Event Filtering Register 1
//  0x050 32  EEFER2     Timerx External Event Filtering Register 2
//  0x054 32  RSTER      TimerA Reset Register
//  0x058 32  CHPER      Timerx Chopper Register
//  0x05C 32  CPT1ECR    Timerx Capture 2 Control Register
//  0x060 32  CPT2ECR    CPT2xCR
//  0x064 32  OUTER      Timerx Output Register
//  0x068 32  FLTER      Timerx Fault Register
// Import:
//  github.com/embeddedgo/stm32/p/mmap
package hrtim_time

const (
	CK_PSCx   TIMECR = 0x07 << 0  //+ HRTIM Timer x Clock prescaler
	CONT      TIMECR = 0x01 << 3  //+ Continuous mode
	RETRIG    TIMECR = 0x01 << 4  //+ Re-triggerable mode
	HALF      TIMECR = 0x01 << 5  //+ Half mode enable
	PSHPLL    TIMECR = 0x01 << 6  //+ Push-Pull mode enable
	SYNCRSTx  TIMECR = 0x01 << 10 //+ Synchronization Resets Timer x
	SYNCSTRTx TIMECR = 0x01 << 11 //+ Synchronization Starts Timer x
	DELCMP2   TIMECR = 0x03 << 12 //+ Delayed CMP2 mode
	DELCMP4   TIMECR = 0x03 << 14 //+ Delayed CMP4 mode
	TxREPU    TIMECR = 0x01 << 17 //+ Timer x Repetition update
	TxRSTU    TIMECR = 0x01 << 18 //+ Timerx reset update
	TBU       TIMECR = 0x01 << 20 //+ TBU
	TCU       TIMECR = 0x01 << 21 //+ TCU
	TDU       TIMECR = 0x01 << 22 //+ TDU
	TEU       TIMECR = 0x01 << 23 //+ TEU
	MSTU      TIMECR = 0x01 << 24 //+ Master Timer update
	DACSYNC   TIMECR = 0x03 << 25 //+ AC Synchronization
	PREEN     TIMECR = 0x01 << 27 //+ Preload enable
	UPDGAT    TIMECR = 0x0F << 28 //+ Update Gating
)

const (
	CK_PSCxn   = 0
	CONTn      = 3
	RETRIGn    = 4
	HALFn      = 5
	PSHPLLn    = 6
	SYNCRSTxn  = 10
	SYNCSTRTxn = 11
	DELCMP2n   = 12
	DELCMP4n   = 14
	TxREPUn    = 17
	TxRSTUn    = 18
	TBUn       = 20
	TCUn       = 21
	TDUn       = 22
	TEUn       = 23
	MSTUn      = 24
	DACSYNCn   = 25
	PREENn     = 27
	UPDGATn    = 28
)

const (
	CMP1    TIMEISR = 0x01 << 0  //+ Compare 1 Interrupt Flag
	CMP2    TIMEISR = 0x01 << 1  //+ Compare 2 Interrupt Flag
	CMP3    TIMEISR = 0x01 << 2  //+ Compare 3 Interrupt Flag
	CMP4    TIMEISR = 0x01 << 3  //+ Compare 4 Interrupt Flag
	REP     TIMEISR = 0x01 << 4  //+ Repetition Interrupt Flag
	UPD     TIMEISR = 0x01 << 6  //+ Update Interrupt Flag
	CPT1    TIMEISR = 0x01 << 7  //+ Capture1 Interrupt Flag
	CPT2    TIMEISR = 0x01 << 8  //+ Capture2 Interrupt Flag
	SETx1   TIMEISR = 0x01 << 9  //+ Output 1 Set Interrupt Flag
	RSTx1   TIMEISR = 0x01 << 10 //+ Output 1 Reset Interrupt Flag
	SETx2   TIMEISR = 0x01 << 11 //+ Output 2 Set Interrupt Flag
	RSTx2   TIMEISR = 0x01 << 12 //+ Output 2 Reset Interrupt Flag
	RST     TIMEISR = 0x01 << 13 //+ Reset Interrupt Flag
	DLYPRT  TIMEISR = 0x01 << 14 //+ Delayed Protection Flag
	CPPSTAT TIMEISR = 0x01 << 16 //+ Current Push Pull Status
	IPPSTAT TIMEISR = 0x01 << 17 //+ Idle Push Pull Status
	O1STAT  TIMEISR = 0x01 << 18 //+ Output 1 State
	O2STAT  TIMEISR = 0x01 << 19 //+ Output 2 State
)

const (
	CMP1n    = 0
	CMP2n    = 1
	CMP3n    = 2
	CMP4n    = 3
	REPn     = 4
	UPDn     = 6
	CPT1n    = 7
	CPT2n    = 8
	SETx1n   = 9
	RSTx1n   = 10
	SETx2n   = 11
	RSTx2n   = 12
	RSTn     = 13
	DLYPRTn  = 14
	CPPSTATn = 16
	IPPSTATn = 17
	O1STATn  = 18
	O2STATn  = 19
)

const (
	CMP1C   TIMEICR = 0x01 << 0  //+ Compare 1 Interrupt flag Clear
	CMP2C   TIMEICR = 0x01 << 1  //+ Compare 2 Interrupt flag Clear
	CMP3C   TIMEICR = 0x01 << 2  //+ Compare 3 Interrupt flag Clear
	CMP4C   TIMEICR = 0x01 << 3  //+ Compare 4 Interrupt flag Clear
	REPC    TIMEICR = 0x01 << 4  //+ Repetition Interrupt flag Clear
	UPDC    TIMEICR = 0x01 << 6  //+ Update Interrupt flag Clear
	CPT1C   TIMEICR = 0x01 << 7  //+ Capture1 Interrupt flag Clear
	CPT2C   TIMEICR = 0x01 << 8  //+ Capture2 Interrupt flag Clear
	SET1xC  TIMEICR = 0x01 << 9  //+ Output 1 Set flag Clear
	RSTx1C  TIMEICR = 0x01 << 10 //+ Output 1 Reset flag Clear
	SET2xC  TIMEICR = 0x01 << 11 //+ Output 2 Set flag Clear
	RSTx2C  TIMEICR = 0x01 << 12 //+ Output 2 Reset flag Clear
	RSTC    TIMEICR = 0x01 << 13 //+ Reset Interrupt flag Clear
	DLYPRTC TIMEICR = 0x01 << 14 //+ Delayed Protection Flag Clear
)

const (
	CMP1Cn   = 0
	CMP2Cn   = 1
	CMP3Cn   = 2
	CMP4Cn   = 3
	REPCn    = 4
	UPDCn    = 6
	CPT1Cn   = 7
	CPT2Cn   = 8
	SET1xCn  = 9
	RSTx1Cn  = 10
	SET2xCn  = 11
	RSTx2Cn  = 12
	RSTCn    = 13
	DLYPRTCn = 14
)

const (
	CMP1IE   TIMEDIER5 = 0x01 << 0  //+ CMP1IE
	CMP2IE   TIMEDIER5 = 0x01 << 1  //+ CMP2IE
	CMP3IE   TIMEDIER5 = 0x01 << 2  //+ CMP3IE
	CMP4IE   TIMEDIER5 = 0x01 << 3  //+ CMP4IE
	REPIE    TIMEDIER5 = 0x01 << 4  //+ REPIE
	UPDIE    TIMEDIER5 = 0x01 << 6  //+ UPDIE
	CPT1IE   TIMEDIER5 = 0x01 << 7  //+ CPT1IE
	CPT2IE   TIMEDIER5 = 0x01 << 8  //+ CPT2IE
	SET1xIE  TIMEDIER5 = 0x01 << 9  //+ SET1xIE
	RSTx1IE  TIMEDIER5 = 0x01 << 10 //+ RSTx1IE
	SETx2IE  TIMEDIER5 = 0x01 << 11 //+ SETx2IE
	RSTx2IE  TIMEDIER5 = 0x01 << 12 //+ RSTx2IE
	RSTIE    TIMEDIER5 = 0x01 << 13 //+ RSTIE
	DLYPRTIE TIMEDIER5 = 0x01 << 14 //+ DLYPRTIE
	CMP1DE   TIMEDIER5 = 0x01 << 16 //+ CMP1DE
	CMP2DE   TIMEDIER5 = 0x01 << 17 //+ CMP2DE
	CMP3DE   TIMEDIER5 = 0x01 << 18 //+ CMP3DE
	CMP4DE   TIMEDIER5 = 0x01 << 19 //+ CMP4DE
	REPDE    TIMEDIER5 = 0x01 << 20 //+ REPDE
	UPDDE    TIMEDIER5 = 0x01 << 22 //+ UPDDE
	CPT1DE   TIMEDIER5 = 0x01 << 23 //+ CPT1DE
	CPT2DE   TIMEDIER5 = 0x01 << 24 //+ CPT2DE
	SET1xDE  TIMEDIER5 = 0x01 << 25 //+ SET1xDE
	RSTx1DE  TIMEDIER5 = 0x01 << 26 //+ RSTx1DE
	SETx2DE  TIMEDIER5 = 0x01 << 27 //+ SETx2DE
	RSTx2DE  TIMEDIER5 = 0x01 << 28 //+ RSTx2DE
	RSTDE    TIMEDIER5 = 0x01 << 29 //+ RSTDE
	DLYPRTDE TIMEDIER5 = 0x01 << 30 //+ DLYPRTDE
)

const (
	CMP1IEn   = 0
	CMP2IEn   = 1
	CMP3IEn   = 2
	CMP4IEn   = 3
	REPIEn    = 4
	UPDIEn    = 6
	CPT1IEn   = 7
	CPT2IEn   = 8
	SET1xIEn  = 9
	RSTx1IEn  = 10
	SETx2IEn  = 11
	RSTx2IEn  = 12
	RSTIEn    = 13
	DLYPRTIEn = 14
	CMP1DEn   = 16
	CMP2DEn   = 17
	CMP3DEn   = 18
	CMP4DEn   = 19
	REPDEn    = 20
	UPDDEn    = 22
	CPT1DEn   = 23
	CPT2DEn   = 24
	SET1xDEn  = 25
	RSTx1DEn  = 26
	SETx2DEn  = 27
	RSTx2DEn  = 28
	RSTDEn    = 29
	DLYPRTDEn = 30
)

const (
	CNTx CNTER = 0xFFFF << 0 //+ Timerx Counter value
)

const (
	CNTxn = 0
)

const (
	PERx PERER = 0xFFFF << 0 //+ Timerx Period value
)

const (
	PERxn = 0
)

const (
	REPx REPER = 0xFF << 0 //+ Timerx Repetition counter value
)

const (
	REPxn = 0
)

const (
	CMP1x CMP1ER = 0xFFFF << 0 //+ Timerx Compare 1 value
)

const (
	CMP1xn = 0
)

const (
	CMP1x CMP1CER = 0xFFFF << 0 //+ Timerx Compare 1 value
	REPx  CMP1CER = 0xFF << 16  //+ Timerx Repetition value (aliased from HRTIM_REPx register)
)

const (
	CMP1xn = 0
	REPxn  = 16
)

const (
	CMP2x CMP2ER = 0xFFFF << 0 //+ Timerx Compare 2 value
)

const (
	CMP2xn = 0
)

const (
	CMP3x CMP3ER = 0xFFFF << 0 //+ Timerx Compare 3 value
)

const (
	CMP3xn = 0
)

const (
	CMP4x CMP4ER = 0xFFFF << 0 //+ Timerx Compare 4 value
)

const (
	CMP4xn = 0
)

const (
	CPT1x CPT1ER = 0xFFFF << 0 //+ Timerx Capture 1 value
)

const (
	CPT1xn = 0
)

const (
	CPT2x CPT2ER = 0xFFFF << 0 //+ Timerx Capture 2 value
)

const (
	CPT2xn = 0
)

const (
	DTRx    DTER = 0x1FF << 0  //+ Deadtime Rising value
	SDTRx   DTER = 0x01 << 9   //+ Sign Deadtime Rising value
	DTPRSC  DTER = 0x07 << 10  //+ Deadtime Prescaler
	DTRSLKx DTER = 0x01 << 14  //+ Deadtime Rising Sign Lock
	DTRLKx  DTER = 0x01 << 15  //+ Deadtime Rising Lock
	DTFx    DTER = 0x1FF << 16 //+ Deadtime Falling value
	SDTFx   DTER = 0x01 << 25  //+ Sign Deadtime Falling value
	DTFSLKx DTER = 0x01 << 30  //+ Deadtime Falling Sign Lock
	DTFLKx  DTER = 0x01 << 31  //+ Deadtime Falling Lock
)

const (
	DTRxn    = 0
	SDTRxn   = 9
	DTPRSCn  = 10
	DTRSLKxn = 14
	DTRLKxn  = 15
	DTFxn    = 16
	SDTFxn   = 25
	DTFSLKxn = 30
	DTFLKxn  = 31
)

const (
	SST       SETE1R = 0x01 << 0  //+ Software Set trigger
	RESYNC    SETE1R = 0x01 << 1  //+ Timer A resynchronizaton
	PER       SETE1R = 0x01 << 2  //+ Timer A Period
	CMP1      SETE1R = 0x01 << 3  //+ Timer A compare 1
	CMP2      SETE1R = 0x01 << 4  //+ Timer A compare 2
	CMP3      SETE1R = 0x01 << 5  //+ Timer A compare 3
	CMP4      SETE1R = 0x01 << 6  //+ Timer A compare 4
	MSTPER    SETE1R = 0x01 << 7  //+ Master Period
	MSTCMP1   SETE1R = 0x01 << 8  //+ Master Compare 1
	MSTCMP2   SETE1R = 0x01 << 9  //+ Master Compare 2
	MSTCMP3   SETE1R = 0x01 << 10 //+ Master Compare 3
	MSTCMP4   SETE1R = 0x01 << 11 //+ Master Compare 4
	TIMEVNT1  SETE1R = 0x01 << 12 //+ Timer Event 1
	TIMEVNT2  SETE1R = 0x01 << 13 //+ Timer Event 2
	TIMEVNT3  SETE1R = 0x01 << 14 //+ Timer Event 3
	TIMEVNT4  SETE1R = 0x01 << 15 //+ Timer Event 4
	TIMEVNT5  SETE1R = 0x01 << 16 //+ Timer Event 5
	TIMEVNT6  SETE1R = 0x01 << 17 //+ Timer Event 6
	TIMEVNT7  SETE1R = 0x01 << 18 //+ Timer Event 7
	TIMEVNT8  SETE1R = 0x01 << 19 //+ Timer Event 8
	TIMEVNT9  SETE1R = 0x01 << 20 //+ Timer Event 9
	EXTEVNT1  SETE1R = 0x01 << 21 //+ External Event 1
	EXTEVNT2  SETE1R = 0x01 << 22 //+ External Event 2
	EXTEVNT3  SETE1R = 0x01 << 23 //+ External Event 3
	EXTEVNT4  SETE1R = 0x01 << 24 //+ External Event 4
	EXTEVNT5  SETE1R = 0x01 << 25 //+ External Event 5
	EXTEVNT6  SETE1R = 0x01 << 26 //+ External Event 6
	EXTEVNT7  SETE1R = 0x01 << 27 //+ External Event 7
	EXTEVNT8  SETE1R = 0x01 << 28 //+ External Event 8
	EXTEVNT9  SETE1R = 0x01 << 29 //+ External Event 9
	EXTEVNT10 SETE1R = 0x01 << 30 //+ External Event 10
	UPDATE    SETE1R = 0x01 << 31 //+ Registers update (transfer preload to active)
)

const (
	SSTn       = 0
	RESYNCn    = 1
	PERn       = 2
	CMP1n      = 3
	CMP2n      = 4
	CMP3n      = 5
	CMP4n      = 6
	MSTPERn    = 7
	MSTCMP1n   = 8
	MSTCMP2n   = 9
	MSTCMP3n   = 10
	MSTCMP4n   = 11
	TIMEVNT1n  = 12
	TIMEVNT2n  = 13
	TIMEVNT3n  = 14
	TIMEVNT4n  = 15
	TIMEVNT5n  = 16
	TIMEVNT6n  = 17
	TIMEVNT7n  = 18
	TIMEVNT8n  = 19
	TIMEVNT9n  = 20
	EXTEVNT1n  = 21
	EXTEVNT2n  = 22
	EXTEVNT3n  = 23
	EXTEVNT4n  = 24
	EXTEVNT5n  = 25
	EXTEVNT6n  = 26
	EXTEVNT7n  = 27
	EXTEVNT8n  = 28
	EXTEVNT9n  = 29
	EXTEVNT10n = 30
	UPDATEn    = 31
)

const (
	SRT       RSTE1R = 0x01 << 0  //+ SRT
	RESYNC    RSTE1R = 0x01 << 1  //+ RESYNC
	PER       RSTE1R = 0x01 << 2  //+ PER
	CMP1      RSTE1R = 0x01 << 3  //+ CMP1
	CMP2      RSTE1R = 0x01 << 4  //+ CMP2
	CMP3      RSTE1R = 0x01 << 5  //+ CMP3
	CMP4      RSTE1R = 0x01 << 6  //+ CMP4
	MSTPER    RSTE1R = 0x01 << 7  //+ MSTPER
	MSTCMP1   RSTE1R = 0x01 << 8  //+ MSTCMP1
	MSTCMP2   RSTE1R = 0x01 << 9  //+ MSTCMP2
	MSTCMP3   RSTE1R = 0x01 << 10 //+ MSTCMP3
	MSTCMP4   RSTE1R = 0x01 << 11 //+ MSTCMP4
	TIMEVNT1  RSTE1R = 0x01 << 12 //+ TIMEVNT1
	TIMEVNT2  RSTE1R = 0x01 << 13 //+ TIMEVNT2
	TIMEVNT3  RSTE1R = 0x01 << 14 //+ TIMEVNT3
	TIMEVNT4  RSTE1R = 0x01 << 15 //+ TIMEVNT4
	TIMEVNT5  RSTE1R = 0x01 << 16 //+ TIMEVNT5
	TIMEVNT6  RSTE1R = 0x01 << 17 //+ TIMEVNT6
	TIMEVNT7  RSTE1R = 0x01 << 18 //+ TIMEVNT7
	TIMEVNT8  RSTE1R = 0x01 << 19 //+ TIMEVNT8
	TIMEVNT9  RSTE1R = 0x01 << 20 //+ TIMEVNT9
	EXTEVNT1  RSTE1R = 0x01 << 21 //+ EXTEVNT1
	EXTEVNT2  RSTE1R = 0x01 << 22 //+ EXTEVNT2
	EXTEVNT3  RSTE1R = 0x01 << 23 //+ EXTEVNT3
	EXTEVNT4  RSTE1R = 0x01 << 24 //+ EXTEVNT4
	EXTEVNT5  RSTE1R = 0x01 << 25 //+ EXTEVNT5
	EXTEVNT6  RSTE1R = 0x01 << 26 //+ EXTEVNT6
	EXTEVNT7  RSTE1R = 0x01 << 27 //+ EXTEVNT7
	EXTEVNT8  RSTE1R = 0x01 << 28 //+ EXTEVNT8
	EXTEVNT9  RSTE1R = 0x01 << 29 //+ EXTEVNT9
	EXTEVNT10 RSTE1R = 0x01 << 30 //+ EXTEVNT10
	UPDATE    RSTE1R = 0x01 << 31 //+ UPDATE
)

const (
	SRTn       = 0
	RESYNCn    = 1
	PERn       = 2
	CMP1n      = 3
	CMP2n      = 4
	CMP3n      = 5
	CMP4n      = 6
	MSTPERn    = 7
	MSTCMP1n   = 8
	MSTCMP2n   = 9
	MSTCMP3n   = 10
	MSTCMP4n   = 11
	TIMEVNT1n  = 12
	TIMEVNT2n  = 13
	TIMEVNT3n  = 14
	TIMEVNT4n  = 15
	TIMEVNT5n  = 16
	TIMEVNT6n  = 17
	TIMEVNT7n  = 18
	TIMEVNT8n  = 19
	TIMEVNT9n  = 20
	EXTEVNT1n  = 21
	EXTEVNT2n  = 22
	EXTEVNT3n  = 23
	EXTEVNT4n  = 24
	EXTEVNT5n  = 25
	EXTEVNT6n  = 26
	EXTEVNT7n  = 27
	EXTEVNT8n  = 28
	EXTEVNT9n  = 29
	EXTEVNT10n = 30
	UPDATEn    = 31
)

const (
	SST       SETE2R = 0x01 << 0  //+ SST
	RESYNC    SETE2R = 0x01 << 1  //+ RESYNC
	PER       SETE2R = 0x01 << 2  //+ PER
	CMP1      SETE2R = 0x01 << 3  //+ CMP1
	CMP2      SETE2R = 0x01 << 4  //+ CMP2
	CMP3      SETE2R = 0x01 << 5  //+ CMP3
	CMP4      SETE2R = 0x01 << 6  //+ CMP4
	MSTPER    SETE2R = 0x01 << 7  //+ MSTPER
	MSTCMP1   SETE2R = 0x01 << 8  //+ MSTCMP1
	MSTCMP2   SETE2R = 0x01 << 9  //+ MSTCMP2
	MSTCMP3   SETE2R = 0x01 << 10 //+ MSTCMP3
	MSTCMP4   SETE2R = 0x01 << 11 //+ MSTCMP4
	TIMEVNT1  SETE2R = 0x01 << 12 //+ TIMEVNT1
	TIMEVNT2  SETE2R = 0x01 << 13 //+ TIMEVNT2
	TIMEVNT3  SETE2R = 0x01 << 14 //+ TIMEVNT3
	TIMEVNT4  SETE2R = 0x01 << 15 //+ TIMEVNT4
	TIMEVNT5  SETE2R = 0x01 << 16 //+ TIMEVNT5
	TIMEVNT6  SETE2R = 0x01 << 17 //+ TIMEVNT6
	TIMEVNT7  SETE2R = 0x01 << 18 //+ TIMEVNT7
	TIMEVNT8  SETE2R = 0x01 << 19 //+ TIMEVNT8
	TIMEVNT9  SETE2R = 0x01 << 20 //+ TIMEVNT9
	EXTEVNT1  SETE2R = 0x01 << 21 //+ EXTEVNT1
	EXTEVNT2  SETE2R = 0x01 << 22 //+ EXTEVNT2
	EXTEVNT3  SETE2R = 0x01 << 23 //+ EXTEVNT3
	EXTEVNT4  SETE2R = 0x01 << 24 //+ EXTEVNT4
	EXTEVNT5  SETE2R = 0x01 << 25 //+ EXTEVNT5
	EXTEVNT6  SETE2R = 0x01 << 26 //+ EXTEVNT6
	EXTEVNT7  SETE2R = 0x01 << 27 //+ EXTEVNT7
	EXTEVNT8  SETE2R = 0x01 << 28 //+ EXTEVNT8
	EXTEVNT9  SETE2R = 0x01 << 29 //+ EXTEVNT9
	EXTEVNT10 SETE2R = 0x01 << 30 //+ EXTEVNT10
	UPDATE    SETE2R = 0x01 << 31 //+ UPDATE
)

const (
	SSTn       = 0
	RESYNCn    = 1
	PERn       = 2
	CMP1n      = 3
	CMP2n      = 4
	CMP3n      = 5
	CMP4n      = 6
	MSTPERn    = 7
	MSTCMP1n   = 8
	MSTCMP2n   = 9
	MSTCMP3n   = 10
	MSTCMP4n   = 11
	TIMEVNT1n  = 12
	TIMEVNT2n  = 13
	TIMEVNT3n  = 14
	TIMEVNT4n  = 15
	TIMEVNT5n  = 16
	TIMEVNT6n  = 17
	TIMEVNT7n  = 18
	TIMEVNT8n  = 19
	TIMEVNT9n  = 20
	EXTEVNT1n  = 21
	EXTEVNT2n  = 22
	EXTEVNT3n  = 23
	EXTEVNT4n  = 24
	EXTEVNT5n  = 25
	EXTEVNT6n  = 26
	EXTEVNT7n  = 27
	EXTEVNT8n  = 28
	EXTEVNT9n  = 29
	EXTEVNT10n = 30
	UPDATEn    = 31
)

const (
	SRT       RSTE2R = 0x01 << 0  //+ SRT
	RESYNC    RSTE2R = 0x01 << 1  //+ RESYNC
	PER       RSTE2R = 0x01 << 2  //+ PER
	CMP1      RSTE2R = 0x01 << 3  //+ CMP1
	CMP2      RSTE2R = 0x01 << 4  //+ CMP2
	CMP3      RSTE2R = 0x01 << 5  //+ CMP3
	CMP4      RSTE2R = 0x01 << 6  //+ CMP4
	MSTPER    RSTE2R = 0x01 << 7  //+ MSTPER
	MSTCMP1   RSTE2R = 0x01 << 8  //+ MSTCMP1
	MSTCMP2   RSTE2R = 0x01 << 9  //+ MSTCMP2
	MSTCMP3   RSTE2R = 0x01 << 10 //+ MSTCMP3
	MSTCMP4   RSTE2R = 0x01 << 11 //+ MSTCMP4
	TIMEVNT1  RSTE2R = 0x01 << 12 //+ TIMEVNT1
	TIMEVNT2  RSTE2R = 0x01 << 13 //+ TIMEVNT2
	TIMEVNT3  RSTE2R = 0x01 << 14 //+ TIMEVNT3
	TIMEVNT4  RSTE2R = 0x01 << 15 //+ TIMEVNT4
	TIMEVNT5  RSTE2R = 0x01 << 16 //+ TIMEVNT5
	TIMEVNT6  RSTE2R = 0x01 << 17 //+ TIMEVNT6
	TIMEVNT7  RSTE2R = 0x01 << 18 //+ TIMEVNT7
	TIMEVNT8  RSTE2R = 0x01 << 19 //+ TIMEVNT8
	TIMEVNT9  RSTE2R = 0x01 << 20 //+ TIMEVNT9
	EXTEVNT1  RSTE2R = 0x01 << 21 //+ EXTEVNT1
	EXTEVNT2  RSTE2R = 0x01 << 22 //+ EXTEVNT2
	EXTEVNT3  RSTE2R = 0x01 << 23 //+ EXTEVNT3
	EXTEVNT4  RSTE2R = 0x01 << 24 //+ EXTEVNT4
	EXTEVNT5  RSTE2R = 0x01 << 25 //+ EXTEVNT5
	EXTEVNT6  RSTE2R = 0x01 << 26 //+ EXTEVNT6
	EXTEVNT7  RSTE2R = 0x01 << 27 //+ EXTEVNT7
	EXTEVNT8  RSTE2R = 0x01 << 28 //+ EXTEVNT8
	EXTEVNT9  RSTE2R = 0x01 << 29 //+ EXTEVNT9
	EXTEVNT10 RSTE2R = 0x01 << 30 //+ EXTEVNT10
	UPDATE    RSTE2R = 0x01 << 31 //+ UPDATE
)

const (
	SRTn       = 0
	RESYNCn    = 1
	PERn       = 2
	CMP1n      = 3
	CMP2n      = 4
	CMP3n      = 5
	CMP4n      = 6
	MSTPERn    = 7
	MSTCMP1n   = 8
	MSTCMP2n   = 9
	MSTCMP3n   = 10
	MSTCMP4n   = 11
	TIMEVNT1n  = 12
	TIMEVNT2n  = 13
	TIMEVNT3n  = 14
	TIMEVNT4n  = 15
	TIMEVNT5n  = 16
	TIMEVNT6n  = 17
	TIMEVNT7n  = 18
	TIMEVNT8n  = 19
	TIMEVNT9n  = 20
	EXTEVNT1n  = 21
	EXTEVNT2n  = 22
	EXTEVNT3n  = 23
	EXTEVNT4n  = 24
	EXTEVNT5n  = 25
	EXTEVNT6n  = 26
	EXTEVNT7n  = 27
	EXTEVNT8n  = 28
	EXTEVNT9n  = 29
	EXTEVNT10n = 30
	UPDATEn    = 31
)

const (
	EE1LTCH EEFER1 = 0x01 << 0  //+ External Event 1 latch
	EE1FLTR EEFER1 = 0x0F << 1  //+ External Event 1 filter
	EE2LTCH EEFER1 = 0x01 << 6  //+ External Event 2 latch
	EE2FLTR EEFER1 = 0x0F << 7  //+ External Event 2 filter
	EE3LTCH EEFER1 = 0x01 << 12 //+ External Event 3 latch
	EE3FLTR EEFER1 = 0x0F << 13 //+ External Event 3 filter
	EE4LTCH EEFER1 = 0x01 << 18 //+ External Event 4 latch
	EE4FLTR EEFER1 = 0x0F << 19 //+ External Event 4 filter
	EE5LTCH EEFER1 = 0x01 << 24 //+ External Event 5 latch
	EE5FLTR EEFER1 = 0x0F << 25 //+ External Event 5 filter
)

const (
	EE1LTCHn = 0
	EE1FLTRn = 1
	EE2LTCHn = 6
	EE2FLTRn = 7
	EE3LTCHn = 12
	EE3FLTRn = 13
	EE4LTCHn = 18
	EE4FLTRn = 19
	EE5LTCHn = 24
	EE5FLTRn = 25
)

const (
	EE6LTCH  EEFER2 = 0x01 << 0  //+ External Event 6 latch
	EE6FLTR  EEFER2 = 0x0F << 1  //+ External Event 6 filter
	EE7LTCH  EEFER2 = 0x01 << 6  //+ External Event 7 latch
	EE7FLTR  EEFER2 = 0x0F << 7  //+ External Event 7 filter
	EE8LTCH  EEFER2 = 0x01 << 12 //+ External Event 8 latch
	EE8FLTR  EEFER2 = 0x0F << 13 //+ External Event 8 filter
	EE9LTCH  EEFER2 = 0x01 << 18 //+ External Event 9 latch
	EE9FLTR  EEFER2 = 0x0F << 19 //+ External Event 9 filter
	EE10LTCH EEFER2 = 0x01 << 24 //+ External Event 10 latch
	EE10FLTR EEFER2 = 0x0F << 25 //+ External Event 10 filter
)

const (
	EE6LTCHn  = 0
	EE6FLTRn  = 1
	EE7LTCHn  = 6
	EE7FLTRn  = 7
	EE8LTCHn  = 12
	EE8FLTRn  = 13
	EE9LTCHn  = 18
	EE9FLTRn  = 19
	EE10LTCHn = 24
	EE10FLTRn = 25
)

const (
	UPDT      RSTER = 0x01 << 1  //+ Timer A Update reset
	CMP2      RSTER = 0x01 << 2  //+ Timer A compare 2 reset
	CMP4      RSTER = 0x01 << 3  //+ Timer A compare 4 reset
	MSTPER    RSTER = 0x01 << 4  //+ Master timer Period
	MSTCMP1   RSTER = 0x01 << 5  //+ Master compare 1
	MSTCMP2   RSTER = 0x01 << 6  //+ Master compare 2
	MSTCMP3   RSTER = 0x01 << 7  //+ Master compare 3
	MSTCMP4   RSTER = 0x01 << 8  //+ Master compare 4
	EXTEVNT1  RSTER = 0x01 << 9  //+ External Event 1
	EXTEVNT2  RSTER = 0x01 << 10 //+ External Event 2
	EXTEVNT3  RSTER = 0x01 << 11 //+ External Event 3
	EXTEVNT4  RSTER = 0x01 << 12 //+ External Event 4
	EXTEVNT5  RSTER = 0x01 << 13 //+ External Event 5
	EXTEVNT6  RSTER = 0x01 << 14 //+ External Event 6
	EXTEVNT7  RSTER = 0x01 << 15 //+ External Event 7
	EXTEVNT8  RSTER = 0x01 << 16 //+ External Event 8
	EXTEVNT9  RSTER = 0x01 << 17 //+ External Event 9
	EXTEVNT10 RSTER = 0x01 << 18 //+ External Event 10
	TIMACMP1  RSTER = 0x01 << 19 //+ Timer A Compare 1
	TIMACMP2  RSTER = 0x01 << 20 //+ Timer A Compare 2
	TIMACMP4  RSTER = 0x01 << 21 //+ Timer A Compare 4
	TIMBCMP1  RSTER = 0x01 << 22 //+ Timer B Compare 1
	TIMBCMP2  RSTER = 0x01 << 23 //+ Timer B Compare 2
	TIMBCMP4  RSTER = 0x01 << 24 //+ Timer B Compare 4
	TIMCCMP1  RSTER = 0x01 << 25 //+ Timer C Compare 1
	TIMCCMP2  RSTER = 0x01 << 26 //+ Timer C Compare 2
	TIMCCMP4  RSTER = 0x01 << 27 //+ Timer C Compare 4
	TIMDCMP1  RSTER = 0x01 << 28 //+ Timer D Compare 1
	TIMDCMP2  RSTER = 0x01 << 29 //+ Timer D Compare 2
	TIMDCMP4  RSTER = 0x01 << 30 //+ Timer D Compare 4
)

const (
	UPDTn      = 1
	CMP2n      = 2
	CMP4n      = 3
	MSTPERn    = 4
	MSTCMP1n   = 5
	MSTCMP2n   = 6
	MSTCMP3n   = 7
	MSTCMP4n   = 8
	EXTEVNT1n  = 9
	EXTEVNT2n  = 10
	EXTEVNT3n  = 11
	EXTEVNT4n  = 12
	EXTEVNT5n  = 13
	EXTEVNT6n  = 14
	EXTEVNT7n  = 15
	EXTEVNT8n  = 16
	EXTEVNT9n  = 17
	EXTEVNT10n = 18
	TIMACMP1n  = 19
	TIMACMP2n  = 20
	TIMACMP4n  = 21
	TIMBCMP1n  = 22
	TIMBCMP2n  = 23
	TIMBCMP4n  = 24
	TIMCCMP1n  = 25
	TIMCCMP2n  = 26
	TIMCCMP4n  = 27
	TIMDCMP1n  = 28
	TIMDCMP2n  = 29
	TIMDCMP4n  = 30
)

const (
	CHPFRQ CHPER = 0x0F << 0 //+ Timerx carrier frequency value
	CHPDTY CHPER = 0x07 << 4 //+ Timerx chopper duty cycle value
	STRTPW CHPER = 0x0F << 7 //+ STRTPW
)

const (
	CHPFRQn = 0
	CHPDTYn = 4
	STRTPWn = 7
)

const (
	SWCPT     CPT1ECR = 0x01 << 0  //+ Software Capture
	UDPCPT    CPT1ECR = 0x01 << 1  //+ Update Capture
	EXEV1CPT  CPT1ECR = 0x01 << 2  //+ External Event 1 Capture
	EXEV2CPT  CPT1ECR = 0x01 << 3  //+ External Event 2 Capture
	EXEV3CPT  CPT1ECR = 0x01 << 4  //+ External Event 3 Capture
	EXEV4CPT  CPT1ECR = 0x01 << 5  //+ External Event 4 Capture
	EXEV5CPT  CPT1ECR = 0x01 << 6  //+ External Event 5 Capture
	EXEV6CPT  CPT1ECR = 0x01 << 7  //+ External Event 6 Capture
	EXEV7CPT  CPT1ECR = 0x01 << 8  //+ External Event 7 Capture
	EXEV8CPT  CPT1ECR = 0x01 << 9  //+ External Event 8 Capture
	EXEV9CPT  CPT1ECR = 0x01 << 10 //+ External Event 9 Capture
	EXEV10CPT CPT1ECR = 0x01 << 11 //+ External Event 10 Capture
	TA1SET    CPT1ECR = 0x01 << 12 //+ Timer A output 1 Set
	TA1RST    CPT1ECR = 0x01 << 13 //+ Timer A output 1 Reset
	TACMP1    CPT1ECR = 0x01 << 14 //+ Timer A Compare 1
	TACMP2    CPT1ECR = 0x01 << 15 //+ Timer A Compare 2
	TB1SET    CPT1ECR = 0x01 << 16 //+ Timer B output 1 Set
	TB1RST    CPT1ECR = 0x01 << 17 //+ Timer B output 1 Reset
	TBCMP1    CPT1ECR = 0x01 << 18 //+ Timer B Compare 1
	TBCMP2    CPT1ECR = 0x01 << 19 //+ Timer B Compare 2
	TC1SET    CPT1ECR = 0x01 << 20 //+ Timer C output 1 Set
	TC1RST    CPT1ECR = 0x01 << 21 //+ Timer C output 1 Reset
	TCCMP1    CPT1ECR = 0x01 << 22 //+ Timer C Compare 1
	TCCMP2    CPT1ECR = 0x01 << 23 //+ Timer C Compare 2
	TD1SET    CPT1ECR = 0x01 << 24 //+ Timer D output 1 Set
	TD1RST    CPT1ECR = 0x01 << 25 //+ Timer D output 1 Reset
	TDCMP1    CPT1ECR = 0x01 << 26 //+ Timer D Compare 1
	TDCMP2    CPT1ECR = 0x01 << 27 //+ Timer D Compare 2
)

const (
	SWCPTn     = 0
	UDPCPTn    = 1
	EXEV1CPTn  = 2
	EXEV2CPTn  = 3
	EXEV3CPTn  = 4
	EXEV4CPTn  = 5
	EXEV5CPTn  = 6
	EXEV6CPTn  = 7
	EXEV7CPTn  = 8
	EXEV8CPTn  = 9
	EXEV9CPTn  = 10
	EXEV10CPTn = 11
	TA1SETn    = 12
	TA1RSTn    = 13
	TACMP1n    = 14
	TACMP2n    = 15
	TB1SETn    = 16
	TB1RSTn    = 17
	TBCMP1n    = 18
	TBCMP2n    = 19
	TC1SETn    = 20
	TC1RSTn    = 21
	TCCMP1n    = 22
	TCCMP2n    = 23
	TD1SETn    = 24
	TD1RSTn    = 25
	TDCMP1n    = 26
	TDCMP2n    = 27
)

const (
	SWCPT     CPT2ECR = 0x01 << 0  //+ Software Capture
	UDPCPT    CPT2ECR = 0x01 << 1  //+ Update Capture
	EXEV1CPT  CPT2ECR = 0x01 << 2  //+ External Event 1 Capture
	EXEV2CPT  CPT2ECR = 0x01 << 3  //+ External Event 2 Capture
	EXEV3CPT  CPT2ECR = 0x01 << 4  //+ External Event 3 Capture
	EXEV4CPT  CPT2ECR = 0x01 << 5  //+ External Event 4 Capture
	EXEV5CPT  CPT2ECR = 0x01 << 6  //+ External Event 5 Capture
	EXEV6CPT  CPT2ECR = 0x01 << 7  //+ External Event 6 Capture
	EXEV7CPT  CPT2ECR = 0x01 << 8  //+ External Event 7 Capture
	EXEV8CPT  CPT2ECR = 0x01 << 9  //+ External Event 8 Capture
	EXEV9CPT  CPT2ECR = 0x01 << 10 //+ External Event 9 Capture
	EXEV10CPT CPT2ECR = 0x01 << 11 //+ External Event 10 Capture
	TA1SET    CPT2ECR = 0x01 << 12 //+ Timer A output 1 Set
	TA1RST    CPT2ECR = 0x01 << 13 //+ Timer A output 1 Reset
	TACMP1    CPT2ECR = 0x01 << 14 //+ Timer A Compare 1
	TACMP2    CPT2ECR = 0x01 << 15 //+ Timer A Compare 2
	TB1SET    CPT2ECR = 0x01 << 16 //+ Timer B output 1 Set
	TB1RST    CPT2ECR = 0x01 << 17 //+ Timer B output 1 Reset
	TBCMP1    CPT2ECR = 0x01 << 18 //+ Timer B Compare 1
	TBCMP2    CPT2ECR = 0x01 << 19 //+ Timer B Compare 2
	TC1SET    CPT2ECR = 0x01 << 20 //+ Timer C output 1 Set
	TC1RST    CPT2ECR = 0x01 << 21 //+ Timer C output 1 Reset
	TCCMP1    CPT2ECR = 0x01 << 22 //+ Timer C Compare 1
	TCCMP2    CPT2ECR = 0x01 << 23 //+ Timer C Compare 2
	TD1SET    CPT2ECR = 0x01 << 24 //+ Timer D output 1 Set
	TD1RST    CPT2ECR = 0x01 << 25 //+ Timer D output 1 Reset
	TDCMP1    CPT2ECR = 0x01 << 26 //+ Timer D Compare 1
	TDCMP2    CPT2ECR = 0x01 << 27 //+ Timer D Compare 2
)

const (
	SWCPTn     = 0
	UDPCPTn    = 1
	EXEV1CPTn  = 2
	EXEV2CPTn  = 3
	EXEV3CPTn  = 4
	EXEV4CPTn  = 5
	EXEV5CPTn  = 6
	EXEV6CPTn  = 7
	EXEV7CPTn  = 8
	EXEV8CPTn  = 9
	EXEV9CPTn  = 10
	EXEV10CPTn = 11
	TA1SETn    = 12
	TA1RSTn    = 13
	TACMP1n    = 14
	TACMP2n    = 15
	TB1SETn    = 16
	TB1RSTn    = 17
	TBCMP1n    = 18
	TBCMP2n    = 19
	TC1SETn    = 20
	TC1RSTn    = 21
	TCCMP1n    = 22
	TCCMP2n    = 23
	TD1SETn    = 24
	TD1RSTn    = 25
	TDCMP1n    = 26
	TDCMP2n    = 27
)

const (
	POL1     OUTER = 0x01 << 1  //+ Output 1 polarity
	IDLEM1   OUTER = 0x01 << 2  //+ Output 1 Idle mode
	IDLES1   OUTER = 0x01 << 3  //+ Output 1 Idle State
	FAULT1   OUTER = 0x03 << 4  //+ Output 1 Fault state
	CHP1     OUTER = 0x01 << 6  //+ Output 1 Chopper enable
	DIDL1    OUTER = 0x01 << 7  //+ Output 1 Deadtime upon burst mode Idle entry
	DTEN     OUTER = 0x01 << 8  //+ Deadtime enable
	DLYPRTEN OUTER = 0x01 << 9  //+ Delayed Protection Enable
	DLYPRT   OUTER = 0x07 << 10 //+ Delayed Protection
	POL2     OUTER = 0x01 << 17 //+ Output 2 polarity
	IDLEM2   OUTER = 0x01 << 18 //+ Output 2 Idle mode
	IDLES2   OUTER = 0x01 << 19 //+ Output 2 Idle State
	FAULT2   OUTER = 0x03 << 20 //+ Output 2 Fault state
	CHP2     OUTER = 0x01 << 22 //+ Output 2 Chopper enable
	DIDL2    OUTER = 0x01 << 23 //+ Output 2 Deadtime upon burst mode Idle entry
)

const (
	POL1n     = 1
	IDLEM1n   = 2
	IDLES1n   = 3
	FAULT1n   = 4
	CHP1n     = 6
	DIDL1n    = 7
	DTENn     = 8
	DLYPRTENn = 9
	DLYPRTn   = 10
	POL2n     = 17
	IDLEM2n   = 18
	IDLES2n   = 19
	FAULT2n   = 20
	CHP2n     = 22
	DIDL2n    = 23
)

const (
	FLT1EN FLTER = 0x01 << 0  //+ Fault 1 enable
	FLT2EN FLTER = 0x01 << 1  //+ Fault 2 enable
	FLT3EN FLTER = 0x01 << 2  //+ Fault 3 enable
	FLT4EN FLTER = 0x01 << 3  //+ Fault 4 enable
	FLT5EN FLTER = 0x01 << 4  //+ Fault 5 enable
	FLTLCK FLTER = 0x01 << 31 //+ Fault sources Lock
)

const (
	FLT1ENn = 0
	FLT2ENn = 1
	FLT3ENn = 2
	FLT4ENn = 3
	FLT5ENn = 4
	FLTLCKn = 31
)