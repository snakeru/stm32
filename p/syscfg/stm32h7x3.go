// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32h7x3

// Package syscfg provides access to the registers of the SYSCFG peripheral.
//
// Instances:
//  SYSCFG  SYSCFG_BASE  APB4  -  System configuration controller
// Registers:
//  0x004 32  PMCR       peripheral mode configuration register
//  0x008 32  EXTICR[4]  select GPIO port for EXTI line (4 x 4bit)
//  0x020 32  CCCSR      compensation cell control/status register
//  0x024 32  CCVR       SYSCFG compensation cell value register
//  0x028 32  CCCR       SYSCFG compensation cell code register
//  0x02C 32  PWRCR      SYSCFG power control register
//  0x124 32  PKGR       SYSCFG package register
//  0x300 32  UR0        SYSCFG user register 0
//  0x308 32  UR2        SYSCFG user register 2
//  0x30C 32  UR3        SYSCFG user register 3
//  0x310 32  UR4        SYSCFG user register 4
//  0x314 32  UR5        SYSCFG user register 5
//  0x318 32  UR6        SYSCFG user register 6
//  0x31C 32  UR7        SYSCFG user register 7
//  0x320 32  UR8        SYSCFG user register 8
//  0x324 32  UR9        SYSCFG user register 9
//  0x328 32  UR10       SYSCFG user register 10
//  0x32C 32  UR11       SYSCFG user register 11
//  0x330 32  UR12       SYSCFG user register 12
//  0x334 32  UR13       SYSCFG user register 13
//  0x338 32  UR14       SYSCFG user register 14
//  0x33C 32  UR15       SYSCFG user register 15
//  0x340 32  UR16       SYSCFG user register 16
//  0x344 32  UR17       SYSCFG user register 17
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package syscfg

const (
	I2C1FMP     PMCR = 0x01 << 0  //+ I2C1 Fm+
	I2C2FMP     PMCR = 0x01 << 1  //+ I2C2 Fm+
	I2C3FMP     PMCR = 0x01 << 2  //+ I2C3 Fm+
	I2C4FMP     PMCR = 0x01 << 3  //+ I2C4 Fm+
	PB6FMP      PMCR = 0x01 << 4  //+ PB(6) Fm+
	PB7FMP      PMCR = 0x01 << 5  //+ PB(7) Fast Mode Plus
	PB8FMP      PMCR = 0x01 << 6  //+ PB(8) Fast Mode Plus
	PB9FMP      PMCR = 0x01 << 7  //+ PB(9) Fm+
	BOOSTE      PMCR = 0x01 << 8  //+ Booster Enable
	BOOSTVDDSEL PMCR = 0x01 << 9  //+ Analog switch supply voltage selection
	EPIS        PMCR = 0x07 << 21 //+ Ethernet PHY Interface Selection
	PA0SO       PMCR = 0x01 << 24 //+ PA0 Switch Open
	PA1SO       PMCR = 0x01 << 25 //+ PA1 Switch Open
	PC2SO       PMCR = 0x01 << 26 //+ PC2 Switch Open
	PC3SO       PMCR = 0x01 << 27 //+ PC3 Switch Open
)

const (
	I2C1FMPn     = 0
	I2C2FMPn     = 1
	I2C3FMPn     = 2
	I2C4FMPn     = 3
	PB6FMPn      = 4
	PB7FMPn      = 5
	PB8FMPn      = 6
	PB9FMPn      = 7
	BOOSTEn      = 8
	BOOSTVDDSELn = 9
	EPISn        = 21
	PA0SOn       = 24
	PA1SOn       = 25
	PC2SOn       = 26
	PC3SOn       = 27
)

const (
	EN    CCCSR = 0x01 << 0  //+ enable
	CS    CCCSR = 0x01 << 1  //+ Code selection
	READY CCCSR = 0x01 << 8  //+ Compensation cell ready flag
	HSLV  CCCSR = 0x01 << 16 //+ High-speed at low-voltage
)

const (
	ENn    = 0
	CSn    = 1
	READYn = 8
	HSLVn  = 16
)

const (
	NCV CCVR = 0x0F << 0 //+ NMOS compensation value
	PCV CCVR = 0x0F << 4 //+ PMOS compensation value
)

const (
	NCVn = 0
	PCVn = 4
)

const (
	NCC CCCR = 0x0F << 0 //+ NMOS compensation code
	PCC CCCR = 0x0F << 4 //+ PMOS compensation code
)

const (
	NCCn = 0
	PCCn = 4
)

const (
	ODEN PWRCR = 0x0F << 0 //+ Overdrive enable
)

const (
	ODENn = 0
)

const (
	PKG PKGR = 0x0F << 0 //+ Package
)

const (
	PKGn = 0
)

const (
	BKS UR0 = 0x01 << 0  //+ Bank Swap
	RDP UR0 = 0xFF << 16 //+ Readout protection
)

const (
	BKSn = 0
	RDPn = 16
)

const (
	BORH      UR2 = 0x03 << 0    //+ BOR_LVL Brownout Reset Threshold Level
	BOOT_ADD0 UR2 = 0xFFFF << 16 //+ Boot Address 0
)

const (
	BORHn      = 0
	BOOT_ADD0n = 16
)

const (
	BOOT_ADD1 UR3 = 0xFFFF << 16 //+ Boot Address 1
)

const (
	BOOT_ADD1n = 16
)

const (
	MEPAD_1 UR4 = 0x01 << 16 //+ Mass Erase Protected Area Disabled for bank 1
)

const (
	MEPAD_1n = 16
)

const (
	MESAD_1 UR5 = 0x01 << 0  //+ Mass erase secured area disabled for bank 1
	WRPN_1  UR5 = 0xFF << 16 //+ Write protection for flash bank 1
)

const (
	MESAD_1n = 0
	WRPN_1n  = 16
)

const (
	PA_BEG_1 UR6 = 0xFFF << 0  //+ Protected area start address for bank 1
	PA_END_1 UR6 = 0xFFF << 16 //+ Protected area end address for bank 1
)

const (
	PA_BEG_1n = 0
	PA_END_1n = 16
)

const (
	SA_BEG_1 UR7 = 0xFFF << 0  //+ Secured area start address for bank 1
	SA_END_1 UR7 = 0xFFF << 16 //+ Secured area end address for bank 1
)

const (
	SA_BEG_1n = 0
	SA_END_1n = 16
)

const (
	MEPAD_2 UR8 = 0x01 << 0  //+ Mass erase protected area disabled for bank 2
	MESAD_2 UR8 = 0x01 << 16 //+ Mass erase secured area disabled for bank 2
)

const (
	MEPAD_2n = 0
	MESAD_2n = 16
)

const (
	WRPN_2   UR9 = 0xFF << 0   //+ Write protection for flash bank 2
	PA_BEG_2 UR9 = 0xFFF << 16 //+ Protected area start address for bank 2
)

const (
	WRPN_2n   = 0
	PA_BEG_2n = 16
)

const (
	PA_END_2 UR10 = 0xFFF << 0  //+ Protected area end address for bank 2
	SA_BEG_2 UR10 = 0xFFF << 16 //+ Secured area start address for bank 2
)

const (
	PA_END_2n = 0
	SA_BEG_2n = 16
)

const (
	SA_END_2 UR11 = 0xFFF << 0 //+ Secured area end address for bank 2
	IWDG1M   UR11 = 0x01 << 16 //+ Independent Watchdog 1 mode
)

const (
	SA_END_2n = 0
	IWDG1Mn   = 16
)

const (
	SECURE UR12 = 0x01 << 16 //+ Secure mode
)

const (
	SECUREn = 16
)

const (
	SDRS    UR13 = 0x03 << 0  //+ Secured DTCM RAM Size
	D1SBRST UR13 = 0x01 << 16 //+ D1 Standby reset
)

const (
	SDRSn    = 0
	D1SBRSTn = 16
)

const (
	D1STPRST UR14 = 0x01 << 0 //+ D1 Stop Reset
)

const (
	D1STPRSTn = 0
)

const (
	FZIWDGSTB UR15 = 0x01 << 16 //+ Freeze independent watchdog in Standby mode
)

const (
	FZIWDGSTBn = 16
)

const (
	FZIWDGSTP UR16 = 0x01 << 0  //+ Freeze independent watchdog in Stop mode
	PKP       UR16 = 0x01 << 16 //+ Private key programmed
)

const (
	FZIWDGSTPn = 0
	PKPn       = 16
)

const (
	IO_HSLV UR17 = 0x01 << 0 //+ I/O high speed / low voltage
)

const (
	IO_HSLVn = 0
)
