// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32f412

// Package irq provides the list of supported external interrupts.
package irq

const (
	PVD                = 1  // PVD through EXTI line detection interrupt
	TAMP_STAMP         = 2  // Tamper and TimeStamp interrupts through the EXTI line
	RTC_WKUP           = 3  // RTC Wakeup interrupt through the EXTI line
	FLASH              = 4  // FLASH global interrupt
	RCC                = 5  // RCC global interrupt
	EXTI0              = 6  // EXTI Line0 interrupt
	EXTI1              = 7  // EXTI Line1 interrupt
	EXTI2              = 8  // EXTI Line2 interrupt
	EXTI3              = 9  // EXTI Line3 interrupt
	EXTI4              = 10 // EXTI Line4 interrupt
	ADC                = 18 // ADC1 global interrupt
	CAN1_TX            = 19 // CAN1 TX interrupts
	CAN1_RX0           = 20 // CAN1 RX0 interrupts
	CAN1_RX1           = 21 // CAN1 RX1 interrupts
	CAN1_SCE           = 22 // CAN1 SCE interrupt
	EXTI9_5            = 23 // EXTI Line[9:5] interrupts
	TIM1_BRK_TIM9      = 24 // TIM1 Break interrupt and TIM9 global interrupt
	TIM1_UP_TIM10      = 25 // TIM1 Update interrupt and TIM10 global interrupt
	TIM1_TRG_COM_TIM11 = 26 // TIM1 Trigger and Commutation interrupts and TIM11 global interrupt
	TIM1_CC            = 27 // TIM1 Capture Compare interrupt
	TIM2               = 28 // TIM2 global interrupt
	TIM3               = 29 // TIM3 global interrupt
	I2C1_EV            = 31 // I2C1 event interrupt
	I2C1_ER            = 32 // I2C1 error interrupt
	I2C2_EV            = 33 // I2C2 event interrupt
	I2C2_ER            = 34 // I2C2 error interrupt
	SPI1               = 35 // SPI1 global interrupt
	SPI2               = 36 // SPI2 global interrupt
	EXTI15_10          = 40 // EXTI Line[15:10] interrupts
	RTC_ALARM          = 41 // RTC Alarms (A and B) through EXTI line interrupt
	OTG_FS_WKUP        = 42 // USB On-The-Go FS Wakeup through EXTI line interrupt
	TIM12              = 43 // Timer 12 global interrupt
	TIM13              = 44 // Timer 13 global interrupt
	TIM14              = 45 // Timer 14 global interrupt
	SDIO               = 49 // SDIO global interrupt
	TIM6_DACUNDER      = 54 // TIM6 global and DAC12 underrun interrupts
	TIM7               = 55 // TIM7 global interrupt
	OTG_FS             = 67 // USB On The Go FS global interrupt
	I2C3_EV            = 72 // I2C3 event interrupt
	I2C3_ER            = 73 // I2C3 error interrupt
	HASH_RNG           = 80 // Hash and Rng global interrupt
	FPU                = 81 // Floating point unit interrupt
	SPI4               = 84 // SPI4 global interrupt
)
