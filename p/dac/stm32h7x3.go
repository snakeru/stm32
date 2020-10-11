// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32h7x3

// Package dac provides access to the registers of the DAC peripheral.
//
// Instances:
//  DAC  DAC_BASE  -  -  DAC
// Registers:
//  0x000 32  CR       DAC control register
//  0x004 32  SWTRGR   DAC software trigger register
//  0x008 32  DHR12R1  DAC channel1 12-bit right-aligned data holding register
//  0x00C 32  DHR12L1  DAC channel1 12-bit left aligned data holding register
//  0x010 32  DHR8R1   DAC channel1 8-bit right aligned data holding register
//  0x014 32  DHR12R2  DAC channel2 12-bit right aligned data holding register
//  0x018 32  DHR12L2  DAC channel2 12-bit left aligned data holding register
//  0x01C 32  DHR8R2   DAC channel2 8-bit right-aligned data holding register
//  0x020 32  DHR12RD  Dual DAC 12-bit right-aligned data holding register
//  0x024 32  DHR12LD  DUAL DAC 12-bit left aligned data holding register
//  0x028 32  DHR8RD   DUAL DAC 8-bit right aligned data holding register
//  0x02C 32  DOR1     DAC channel1 data output register
//  0x030 32  DOR2     DAC channel2 data output register
//  0x034 32  SR       DAC status register
//  0x038 32  CCR      DAC calibration control register
//  0x03C 32  MCR      DAC mode control register
//  0x040 32  SHSR1    DAC Sample and Hold sample time register 1
//  0x044 32  SHSR2    DAC Sample and Hold sample time register 2
//  0x048 32  SHHR     DAC Sample and Hold hold time register
//  0x04C 32  SHRR     DAC Sample and Hold refresh time register
// Import:
//  github.com/embeddedgo/stm32/p/mmap
package dac

const (
	EN1       CR = 0x01 << 0  //+ DAC channel1 enable This bit is set and cleared by software to enable/disable DAC channel1.
	TEN1      CR = 0x01 << 1  //+ DAC channel1 trigger enable
	TSEL1     CR = 0x07 << 2  //+ DAC channel1 trigger selection These bits select the external event used to trigger DAC channel1. Note: Only used if bit TEN1 = 1 (DAC channel1 trigger enabled).
	WAVE1     CR = 0x03 << 6  //+ DAC channel1 noise/triangle wave generation enable These bits are set and cleared by software. Note: Only used if bit TEN1 = 1 (DAC channel1 trigger enabled).
	MAMP1     CR = 0x0F << 8  //+ DAC channel1 mask/amplitude selector These bits are written by software to select mask in wave generation mode or amplitude in triangle generation mode. = 1011: Unmask bits[11:0] of LFSR/ triangle amplitude equal to 4095
	DMAEN1    CR = 0x01 << 12 //+ DAC channel1 DMA enable This bit is set and cleared by software.
	DMAUDRIE1 CR = 0x01 << 13 //+ DAC channel1 DMA Underrun Interrupt enable This bit is set and cleared by software.
	CEN1      CR = 0x01 << 14 //+ DAC Channel 1 calibration enable This bit is set and cleared by software to enable/disable DAC channel 1 calibration, it can be written only if bit EN1=0 into DAC_CR (the calibration mode can be entered/exit only when the DAC channel is disabled) Otherwise, the write operation is ignored.
	EN2       CR = 0x01 << 16 //+ DAC channel2 enable This bit is set and cleared by software to enable/disable DAC channel2.
	TEN2      CR = 0x01 << 17 //+ DAC channel2 trigger enable
	TSEL2     CR = 0x07 << 18 //+ DAC channel2 trigger selection These bits select the external event used to trigger DAC channel2 Note: Only used if bit TEN2 = 1 (DAC channel2 trigger enabled).
	WAVE2     CR = 0x03 << 22 //+ DAC channel2 noise/triangle wave generation enable These bits are set/reset by software. 1x: Triangle wave generation enabled Note: Only used if bit TEN2 = 1 (DAC channel2 trigger enabled)
	MAMP2     CR = 0x0F << 24 //+ DAC channel2 mask/amplitude selector These bits are written by software to select mask in wave generation mode or amplitude in triangle generation mode. = 1011: Unmask bits[11:0] of LFSR/ triangle amplitude equal to 4095
	DMAEN2    CR = 0x01 << 28 //+ DAC channel2 DMA enable This bit is set and cleared by software.
	DMAUDRIE2 CR = 0x01 << 29 //+ DAC channel2 DMA underrun interrupt enable This bit is set and cleared by software.
	CEN2      CR = 0x01 << 30 //+ DAC Channel 2 calibration enable This bit is set and cleared by software to enable/disable DAC channel 2 calibration, it can be written only if bit EN2=0 into DAC_CR (the calibration mode can be entered/exit only when the DAC channel is disabled) Otherwise, the write operation is ignored.
)

const (
	EN1n       = 0
	TEN1n      = 1
	TSEL1n     = 2
	WAVE1n     = 6
	MAMP1n     = 8
	DMAEN1n    = 12
	DMAUDRIE1n = 13
	CEN1n      = 14
	EN2n       = 16
	TEN2n      = 17
	TSEL2n     = 18
	WAVE2n     = 22
	MAMP2n     = 24
	DMAEN2n    = 28
	DMAUDRIE2n = 29
	CEN2n      = 30
)

const (
	SWTRIG1 SWTRGR = 0x01 << 0 //+ DAC channel1 software trigger This bit is set by software to trigger the DAC in software trigger mode. Note: This bit is cleared by hardware (one APB1 clock cycle later) once the DAC_DHR1 register value has been loaded into the DAC_DOR1 register.
	SWTRIG2 SWTRGR = 0x01 << 1 //+ DAC channel2 software trigger This bit is set by software to trigger the DAC in software trigger mode. Note: This bit is cleared by hardware (one APB1 clock cycle later) once the DAC_DHR2 register value has been loaded into the DAC_DOR2 register.
)

const (
	SWTRIG1n = 0
	SWTRIG2n = 1
)

const (
	DACC1DHR DHR12R1 = 0xFFF << 0 //+ DAC channel1 12-bit right-aligned data These bits are written by software which specifies 12-bit data for DAC channel1.
)

const (
	DACC1DHRn = 0
)

const (
	DACC1DHR DHR12L1 = 0xFFF << 4 //+ DAC channel1 12-bit left-aligned data These bits are written by software which specifies 12-bit data for DAC channel1.
)

const (
	DACC1DHRn = 4
)

const (
	DACC1DHR DHR8R1 = 0xFF << 0 //+ DAC channel1 8-bit right-aligned data These bits are written by software which specifies 8-bit data for DAC channel1.
)

const (
	DACC1DHRn = 0
)

const (
	DACC2DHR DHR12R2 = 0xFFF << 0 //+ DAC channel2 12-bit right-aligned data These bits are written by software which specifies 12-bit data for DAC channel2.
)

const (
	DACC2DHRn = 0
)

const (
	DACC2DHR DHR12L2 = 0xFFF << 4 //+ DAC channel2 12-bit left-aligned data These bits are written by software which specify 12-bit data for DAC channel2.
)

const (
	DACC2DHRn = 4
)

const (
	DACC2DHR DHR8R2 = 0xFF << 0 //+ DAC channel2 8-bit right-aligned data These bits are written by software which specifies 8-bit data for DAC channel2.
)

const (
	DACC2DHRn = 0
)

const (
	DACC1DHR DHR12RD = 0xFFF << 0  //+ DAC channel1 12-bit right-aligned data These bits are written by software which specifies 12-bit data for DAC channel1.
	DACC2DHR DHR12RD = 0xFFF << 16 //+ DAC channel2 12-bit right-aligned data These bits are written by software which specifies 12-bit data for DAC channel2.
)

const (
	DACC1DHRn = 0
	DACC2DHRn = 16
)

const (
	DACC1DHR DHR12LD = 0xFFF << 4  //+ DAC channel1 12-bit left-aligned data These bits are written by software which specifies 12-bit data for DAC channel1.
	DACC2DHR DHR12LD = 0xFFF << 20 //+ DAC channel2 12-bit left-aligned data These bits are written by software which specifies 12-bit data for DAC channel2.
)

const (
	DACC1DHRn = 4
	DACC2DHRn = 20
)

const (
	DACC1DHR DHR8RD = 0xFF << 0 //+ DAC channel1 8-bit right-aligned data These bits are written by software which specifies 8-bit data for DAC channel1.
	DACC2DHR DHR8RD = 0xFF << 8 //+ DAC channel2 8-bit right-aligned data These bits are written by software which specifies 8-bit data for DAC channel2.
)

const (
	DACC1DHRn = 0
	DACC2DHRn = 8
)

const (
	DACC1DOR DOR1 = 0xFFF << 0 //+ DAC channel1 data output These bits are read-only, they contain data output for DAC channel1.
)

const (
	DACC1DORn = 0
)

const (
	DACC2DOR DOR2 = 0xFFF << 0 //+ DAC channel2 data output These bits are read-only, they contain data output for DAC channel2.
)

const (
	DACC2DORn = 0
)

const (
	DMAUDR1   SR = 0x01 << 13 //+ DAC channel1 DMA underrun flag This bit is set by hardware and cleared by software (by writing it to 1).
	CAL_FLAG1 SR = 0x01 << 14 //+ DAC Channel 1 calibration offset status This bit is set and cleared by hardware
	BWST1     SR = 0x01 << 15 //+ DAC Channel 1 busy writing sample time flag This bit is systematically set just after Sample & Hold mode enable and is set each time the software writes the register DAC_SHSR1, It is cleared by hardware when the write operation of DAC_SHSR1 is complete. (It takes about 3LSI periods of synchronization).
	DMAUDR2   SR = 0x01 << 29 //+ DAC channel2 DMA underrun flag This bit is set by hardware and cleared by software (by writing it to 1).
	CAL_FLAG2 SR = 0x01 << 30 //+ DAC Channel 2 calibration offset status This bit is set and cleared by hardware
	BWST2     SR = 0x01 << 31 //+ DAC Channel 2 busy writing sample time flag This bit is systematically set just after Sample & Hold mode enable and is set each time the software writes the register DAC_SHSR2, It is cleared by hardware when the write operation of DAC_SHSR2 is complete. (It takes about 3 LSI periods of synchronization).
)

const (
	DMAUDR1n   = 13
	CAL_FLAG1n = 14
	BWST1n     = 15
	DMAUDR2n   = 29
	CAL_FLAG2n = 30
	BWST2n     = 31
)

const (
	OTRIM1 CCR = 0x1F << 0  //+ DAC Channel 1 offset trimming value
	OTRIM2 CCR = 0x1F << 16 //+ DAC Channel 2 offset trimming value
)

const (
	OTRIM1n = 0
	OTRIM2n = 16
)

const (
	MODE1 MCR = 0x07 << 0  //+ DAC Channel 1 mode These bits can be written only when the DAC is disabled and not in the calibration mode (when bit EN1=0 and bit CEN1 =0 in the DAC_CR register). If EN1=1 or CEN1 =1 the write operation is ignored. They can be set and cleared by software to select the DAC Channel 1 mode: DAC Channel 1 in normal Mode DAC Channel 1 in sample &amp; hold mode
	MODE2 MCR = 0x07 << 16 //+ DAC Channel 2 mode These bits can be written only when the DAC is disabled and not in the calibration mode (when bit EN2=0 and bit CEN2 =0 in the DAC_CR register). If EN2=1 or CEN2 =1 the write operation is ignored. They can be set and cleared by software to select the DAC Channel 2 mode: DAC Channel 2 in normal Mode DAC Channel 2 in sample &amp; hold mode
)

const (
	MODE1n = 0
	MODE2n = 16
)

const (
	TSAMPLE1 SHSR1 = 0x3FF << 0 //+ DAC Channel 1 sample Time (only valid in sample &amp; hold mode) These bits can be written when the DAC channel1 is disabled or also during normal operation. in the latter case, the write can be done only when BWSTx of DAC_SR register is low, If BWSTx=1, the write operation is ignored.
)

const (
	TSAMPLE1n = 0
)

const (
	TSAMPLE2 SHSR2 = 0x3FF << 0 //+ DAC Channel 2 sample Time (only valid in sample &amp; hold mode) These bits can be written when the DAC channel2 is disabled or also during normal operation. in the latter case, the write can be done only when BWSTx of DAC_SR register is low, if BWSTx=1, the write operation is ignored.
)

const (
	TSAMPLE2n = 0
)

const (
	THOLD1 SHHR = 0x3FF << 0  //+ DAC Channel 1 hold Time (only valid in sample &amp; hold mode) Hold time= (THOLD[9:0]) x T LSI
	THOLD2 SHHR = 0x3FF << 16 //+ DAC Channel 2 hold time (only valid in sample &amp; hold mode). Hold time= (THOLD[9:0]) x T LSI
)

const (
	THOLD1n = 0
	THOLD2n = 16
)

const (
	TREFRESH1 SHRR = 0xFF << 0  //+ DAC Channel 1 refresh Time (only valid in sample &amp; hold mode) Refresh time= (TREFRESH[7:0]) x T LSI
	TREFRESH2 SHRR = 0xFF << 16 //+ DAC Channel 2 refresh Time (only valid in sample &amp; hold mode) Refresh time= (TREFRESH[7:0]) x T LSI
)

const (
	TREFRESH1n = 0
	TREFRESH2n = 16
)
