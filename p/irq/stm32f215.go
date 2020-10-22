// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32f215

// Package irq provides the list of supported external interrupts.
package irq

const (
	WWDG               = 0  // Window Watchdog interrupt
	PVD                = 1  // PVD through EXTI line detection interrupt
	TAMP_STAMP         = 2  // Tamper and TimeStamp interrupts through the EXTI line
	RTC_WKUP           = 3  // RTC Wakeup interrupt through the EXTI line
	FLASH              = 4  // FlASH global interrupt
	RCC                = 5  // RCC global interrupt
	EXTI0              = 6  // EXTI Line0 interrupt
	EXTI1              = 7  // EXTI Line1 interrupt
	EXTI2              = 8  // EXTI Line2 interrupt
	EXTI3              = 9  // EXTI Line3 interrupt
	EXTI4              = 10 // EXTI Line4 interrupt
	DMA1_STREAM0       = 11 // DMA1 Stream0 global interrupt
	DMA1_STREAM1       = 12 // DMA1 Stream1 global interrupt
	DMA1_STREAM2       = 13 // DMA1 Stream2 global interrupt
	DMA1_STREAM3       = 14 // DMA1 Stream3 global interrupt
	DMA1_STREAM4       = 15 // DMA1 Stream4 global interrupt
	DMA1_STREAM5       = 16 // DMA1 Stream5 global interrupt
	DMA1_STREAM6       = 17 // DMA1 Stream6 global interrupt
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
	TIM4               = 30 // TIM4 global interrupt
	I2C1_EV            = 31 // I2C1 event interrupt
	I2C1_ER            = 32 // I2C1 error interrupt
	I2C2_EV            = 33 // I2C2 event interrupt
	I2C2_ER            = 34 // I2C2 error interrupt
	SPI1               = 35 // SPI1 global interrupt
	SPI2               = 36 // SPI2 global interrupt
	USART1             = 37 // USART1 global interrupt
	USART2             = 38 // USART2 global interrupt
	USART3             = 39 // USART3 global interrupt
	EXTI15_10          = 40 // EXTI Line[15:10] interrupts
	RTC_ALARM          = 41 // RTC Alarms (A and B) through EXTI line interrupt
	OTG_FS_WKUP        = 42 // USB On-The-Go FS Wakeup through EXTI line interrupt
	TIM8_BRK_TIM12     = 43 // TIM8 Break interrupt and TIM12 global interrupt
	TIM8_UP_TIM13      = 44 // TIM8 Update interrupt and TIM13 global interrupt
	TIM8_TRG_COM_TIM14 = 45 // TIM8 Trigger and Commutation interrupts and TIM14 global interrupt
	TIM8_CC            = 46 // TIM8 Capture Compare interrupt
	DMA1_STREAM7       = 47 // DMA1 Stream7 global interrupt
	FSMC               = 48 // FSMC global interrupt
	SDIO               = 49 // SDIO global interrupt
	TIM5               = 50 // TIM5 global interrupt
	SPI3               = 51 // SPI3 global interrupt
	UART4              = 52 // UART4 global interrupt
	UART5              = 53 // UART5 global interrupt
	TIM6_DAC           = 54 // TIM6 global interrupt, DAC1 and DAC2 underrun error interrupt
	TIM7               = 55 // TIM7 global interrupt
	DMA2_STREAM0       = 56 // DMA2 Stream0 global interrupt
	DMA2_STREAM1       = 57 // DMA2 Stream1 global interrupt
	DMA2_STREAM2       = 58 // DMA2 Stream2 global interrupt
	DMA2_STREAM3       = 59 // DMA2 Stream3 global interrupt
	DMA2_STREAM4       = 60 // DMA2 Stream4 global interrupt
	ETH                = 61 // Ethernet global interrupt
	ETH_WKUP           = 62 // Ethernet Wakeup through EXTI line interrupt
	CAN2_TX            = 63 // CAN2 TX interrupts
	CAN2_RX0           = 64 // CAN2 RX0 interrupts
	CAN2_RX1           = 65 // CAN2 RX1 interrupts
	CAN2_SCE           = 66 // CAN2 SCE interrupt
	OTG_FS             = 67 // USB On The Go FS global interrupt
	DMA2_STREAM5       = 68 // DMA2 Stream5 global interrupt
	DMA2_STREAM6       = 69 // DMA2 Stream6 global interrupt
	DMA2_STREAM7       = 70 // DMA2 Stream7 global interrupt
	USART6             = 71 // USART6 global interrupt
	I2C3_EV            = 72 // I2C3 event interrupt
	I2C3_ER            = 73 // I2C3 error interrupt
	OTG_HS_EP1_OUT     = 74 // USB On The Go HS End Point 1 Out global interrupt
	OTG_HS_EP1_IN      = 75 // USB On The Go HS End Point 1 In global interrupt
	OTG_HS_WKUP        = 76 // USB On The Go HS Wakeup through EXTI interrupt
	OTG_HS             = 77 // USB On The Go HS global interrupt
	DCMI               = 78 // DCMI global interrupt
	CRYP               = 79 // CRYP crypto global interrupt
	HASH_RNG           = 80 // Hash and Rng global interrupt
)
