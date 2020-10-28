// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32h7x3

// Package bdma provides access to the registers of the BDMA peripheral.
//
// Instances:
//  BDMA  BDMA_BASE  AHB4  BDMA_CH1,BDMA_CH2,BDMA_CH3,BDMA_CH4,BDMA_CH5,BDMA_CH6,BDMA_CH7,BDMA_CH8  BDMA
// Registers:
//  0x000 32  ISR     DMA interrupt status register
//  0x004 32  IFCR    DMA interrupt flag clear register
//  0x008 32  CCR1    DMA channel x configuration register
//  0x00C 32  CNDTR1  DMA channel x number of data register
//  0x010 32  CPAR1   This register must not be written when the channel is enabled.
//  0x014 32  CMAR1   This register must not be written when the channel is enabled.
//  0x01C 32  CCR2    DMA channel x configuration register
//  0x020 32  CNDTR2  DMA channel x number of data register
//  0x024 32  CPAR2   This register must not be written when the channel is enabled.
//  0x028 32  CMAR2   This register must not be written when the channel is enabled.
//  0x030 32  CCR3    DMA channel x configuration register
//  0x034 32  CNDTR3  DMA channel x number of data register
//  0x038 32  CPAR3   This register must not be written when the channel is enabled.
//  0x03C 32  CMAR3   This register must not be written when the channel is enabled.
//  0x044 32  CCR4    DMA channel x configuration register
//  0x048 32  CNDTR4  DMA channel x number of data register
//  0x04C 32  CPAR4   This register must not be written when the channel is enabled.
//  0x050 32  CMAR4   This register must not be written when the channel is enabled.
//  0x058 32  CCR5    DMA channel x configuration register
//  0x05C 32  CNDTR5  DMA channel x number of data register
//  0x060 32  CPAR5   This register must not be written when the channel is enabled.
//  0x064 32  CMAR5   This register must not be written when the channel is enabled.
//  0x06C 32  CCR6    DMA channel x configuration register
//  0x070 32  CNDTR6  DMA channel x number of data register
//  0x074 32  CPAR6   This register must not be written when the channel is enabled.
//  0x078 32  CMAR6   This register must not be written when the channel is enabled.
//  0x080 32  CCR7    DMA channel x configuration register
//  0x084 32  CNDTR7  DMA channel x number of data register
//  0x088 32  CPAR7   This register must not be written when the channel is enabled.
//  0x08C 32  CMAR7   This register must not be written when the channel is enabled.
//  0x094 32  CCR8    DMA channel x configuration register
//  0x098 32  CNDTR8  DMA channel x number of data register
//  0x09C 32  CPAR8   This register must not be written when the channel is enabled.
//  0x0A0 32  CMAR8   This register must not be written when the channel is enabled.
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package bdma

const (
	GIF1  ISR = 0x01 << 0  //+ Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	TCIF1 ISR = 0x01 << 1  //+ Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	HTIF1 ISR = 0x01 << 2  //+ Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	TEIF1 ISR = 0x01 << 3  //+ Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	GIF2  ISR = 0x01 << 4  //+ Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	TCIF2 ISR = 0x01 << 5  //+ Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	HTIF2 ISR = 0x01 << 6  //+ Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	TEIF2 ISR = 0x01 << 7  //+ Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	GIF3  ISR = 0x01 << 8  //+ Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	TCIF3 ISR = 0x01 << 9  //+ Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	HTIF3 ISR = 0x01 << 10 //+ Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	TEIF3 ISR = 0x01 << 11 //+ Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	GIF4  ISR = 0x01 << 12 //+ Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	TCIF4 ISR = 0x01 << 13 //+ Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	HTIF4 ISR = 0x01 << 14 //+ Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	TEIF4 ISR = 0x01 << 15 //+ Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	GIF5  ISR = 0x01 << 16 //+ Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	TCIF5 ISR = 0x01 << 17 //+ Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	HTIF5 ISR = 0x01 << 18 //+ Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	TEIF5 ISR = 0x01 << 19 //+ Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	GIF6  ISR = 0x01 << 20 //+ Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	TCIF6 ISR = 0x01 << 21 //+ Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	HTIF6 ISR = 0x01 << 22 //+ Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	TEIF6 ISR = 0x01 << 23 //+ Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	GIF7  ISR = 0x01 << 24 //+ Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	TCIF7 ISR = 0x01 << 25 //+ Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	HTIF7 ISR = 0x01 << 26 //+ Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	TEIF7 ISR = 0x01 << 27 //+ Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	GIF8  ISR = 0x01 << 28 //+ Channel x global interrupt flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	TCIF8 ISR = 0x01 << 29 //+ Channel x transfer complete flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	HTIF8 ISR = 0x01 << 30 //+ Channel x half transfer flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
	TEIF8 ISR = 0x01 << 31 //+ Channel x transfer error flag (x = 1..8) This bit is set by hardware. It is cleared by software writing 1 to the corresponding bit in the DMA_IFCR register.
)

const (
	GIF1n  = 0
	TCIF1n = 1
	HTIF1n = 2
	TEIF1n = 3
	GIF2n  = 4
	TCIF2n = 5
	HTIF2n = 6
	TEIF2n = 7
	GIF3n  = 8
	TCIF3n = 9
	HTIF3n = 10
	TEIF3n = 11
	GIF4n  = 12
	TCIF4n = 13
	HTIF4n = 14
	TEIF4n = 15
	GIF5n  = 16
	TCIF5n = 17
	HTIF5n = 18
	TEIF5n = 19
	GIF6n  = 20
	TCIF6n = 21
	HTIF6n = 22
	TEIF6n = 23
	GIF7n  = 24
	TCIF7n = 25
	HTIF7n = 26
	TEIF7n = 27
	GIF8n  = 28
	TCIF8n = 29
	HTIF8n = 30
	TEIF8n = 31
)

const (
	CGIF1  IFCR = 0x01 << 0  //+ Channel x global interrupt clear This bit is set and cleared by software.
	CTCIF1 IFCR = 0x01 << 1  //+ Channel x transfer complete clear This bit is set and cleared by software.
	CHTIF1 IFCR = 0x01 << 2  //+ Channel x half transfer clear This bit is set and cleared by software.
	CTEIF1 IFCR = 0x01 << 3  //+ Channel x transfer error clear This bit is set and cleared by software.
	CGIF2  IFCR = 0x01 << 4  //+ Channel x global interrupt clear This bit is set and cleared by software.
	CTCIF2 IFCR = 0x01 << 5  //+ Channel x transfer complete clear This bit is set and cleared by software.
	CHTIF2 IFCR = 0x01 << 6  //+ Channel x half transfer clear This bit is set and cleared by software.
	CTEIF2 IFCR = 0x01 << 7  //+ Channel x transfer error clear This bit is set and cleared by software.
	CGIF3  IFCR = 0x01 << 8  //+ Channel x global interrupt clear This bit is set and cleared by software.
	CTCIF3 IFCR = 0x01 << 9  //+ Channel x transfer complete clear This bit is set and cleared by software.
	CHTIF3 IFCR = 0x01 << 10 //+ Channel x half transfer clear This bit is set and cleared by software.
	CTEIF3 IFCR = 0x01 << 11 //+ Channel x transfer error clear This bit is set and cleared by software.
	CGIF4  IFCR = 0x01 << 12 //+ Channel x global interrupt clear This bit is set and cleared by software.
	CTCIF4 IFCR = 0x01 << 13 //+ Channel x transfer complete clear This bit is set and cleared by software.
	CHTIF4 IFCR = 0x01 << 14 //+ Channel x half transfer clear This bit is set and cleared by software.
	CTEIF4 IFCR = 0x01 << 15 //+ Channel x transfer error clear This bit is set and cleared by software.
	CGIF5  IFCR = 0x01 << 16 //+ Channel x global interrupt clear This bit is set and cleared by software.
	CTCIF5 IFCR = 0x01 << 17 //+ Channel x transfer complete clear This bit is set and cleared by software.
	CHTIF5 IFCR = 0x01 << 18 //+ Channel x half transfer clear This bit is set and cleared by software.
	CTEIF5 IFCR = 0x01 << 19 //+ Channel x transfer error clear This bit is set and cleared by software.
	CGIF6  IFCR = 0x01 << 20 //+ Channel x global interrupt clear This bit is set and cleared by software.
	CTCIF6 IFCR = 0x01 << 21 //+ Channel x transfer complete clear This bit is set and cleared by software.
	CHTIF6 IFCR = 0x01 << 22 //+ Channel x half transfer clear This bit is set and cleared by software.
	CTEIF6 IFCR = 0x01 << 23 //+ Channel x transfer error clear This bit is set and cleared by software.
	CGIF7  IFCR = 0x01 << 24 //+ Channel x global interrupt clear This bit is set and cleared by software.
	CTCIF7 IFCR = 0x01 << 25 //+ Channel x transfer complete clear This bit is set and cleared by software.
	CHTIF7 IFCR = 0x01 << 26 //+ Channel x half transfer clear This bit is set and cleared by software.
	CTEIF7 IFCR = 0x01 << 27 //+ Channel x transfer error clear This bit is set and cleared by software.
	CGIF8  IFCR = 0x01 << 28 //+ Channel x global interrupt clear This bit is set and cleared by software.
	CTCIF8 IFCR = 0x01 << 29 //+ Channel x transfer complete clear This bit is set and cleared by software.
	CHTIF8 IFCR = 0x01 << 30 //+ Channel x half transfer clear This bit is set and cleared by software.
	CTEIF8 IFCR = 0x01 << 31 //+ Channel x transfer error clear This bit is set and cleared by software.
)

const (
	CGIF1n  = 0
	CTCIF1n = 1
	CHTIF1n = 2
	CTEIF1n = 3
	CGIF2n  = 4
	CTCIF2n = 5
	CHTIF2n = 6
	CTEIF2n = 7
	CGIF3n  = 8
	CTCIF3n = 9
	CHTIF3n = 10
	CTEIF3n = 11
	CGIF4n  = 12
	CTCIF4n = 13
	CHTIF4n = 14
	CTEIF4n = 15
	CGIF5n  = 16
	CTCIF5n = 17
	CHTIF5n = 18
	CTEIF5n = 19
	CGIF6n  = 20
	CTCIF6n = 21
	CHTIF6n = 22
	CTEIF6n = 23
	CGIF7n  = 24
	CTCIF7n = 25
	CHTIF7n = 26
	CTEIF7n = 27
	CGIF8n  = 28
	CTCIF8n = 29
	CHTIF8n = 30
	CTEIF8n = 31
)

const (
	EN      CCR1 = 0x01 << 0  //+ Channel enable This bit is set and cleared by software.
	TCIE    CCR1 = 0x01 << 1  //+ Transfer complete interrupt enable This bit is set and cleared by software.
	HTIE    CCR1 = 0x01 << 2  //+ Half transfer interrupt enable This bit is set and cleared by software.
	TEIE    CCR1 = 0x01 << 3  //+ Transfer error interrupt enable This bit is set and cleared by software.
	DIR     CCR1 = 0x01 << 4  //+ Data transfer direction This bit is set and cleared by software.
	CIRC    CCR1 = 0x01 << 5  //+ Circular mode This bit is set and cleared by software.
	PINC    CCR1 = 0x01 << 6  //+ Peripheral increment mode This bit is set and cleared by software.
	MINC    CCR1 = 0x01 << 7  //+ Memory increment mode This bit is set and cleared by software.
	PSIZE   CCR1 = 0x03 << 8  //+ Peripheral size These bits are set and cleared by software.
	MSIZE   CCR1 = 0x03 << 10 //+ Memory size These bits are set and cleared by software.
	PL      CCR1 = 0x03 << 12 //+ Channel priority level These bits are set and cleared by software.
	MEM2MEM CCR1 = 0x01 << 14 //+ Memory to memory mode This bit is set and cleared by software.
)

const (
	ENn      = 0
	TCIEn    = 1
	HTIEn    = 2
	TEIEn    = 3
	DIRn     = 4
	CIRCn    = 5
	PINCn    = 6
	MINCn    = 7
	PSIZEn   = 8
	MSIZEn   = 10
	PLn      = 12
	MEM2MEMn = 14
)

const (
	NDT CNDTR1 = 0xFFFF << 0 //+ Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
)

const (
	NDTn = 0
)

const (
	PA CPAR1 = 0xFFFFFFFF << 0 //+ Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
)

const (
	PAn = 0
)

const (
	MA CMAR1 = 0xFFFFFFFF << 0 //+ Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
)

const (
	MAn = 0
)

const (
	EN      CCR2 = 0x01 << 0  //+ Channel enable This bit is set and cleared by software.
	TCIE    CCR2 = 0x01 << 1  //+ Transfer complete interrupt enable This bit is set and cleared by software.
	HTIE    CCR2 = 0x01 << 2  //+ Half transfer interrupt enable This bit is set and cleared by software.
	TEIE    CCR2 = 0x01 << 3  //+ Transfer error interrupt enable This bit is set and cleared by software.
	DIR     CCR2 = 0x01 << 4  //+ Data transfer direction This bit is set and cleared by software.
	CIRC    CCR2 = 0x01 << 5  //+ Circular mode This bit is set and cleared by software.
	PINC    CCR2 = 0x01 << 6  //+ Peripheral increment mode This bit is set and cleared by software.
	MINC    CCR2 = 0x01 << 7  //+ Memory increment mode This bit is set and cleared by software.
	PSIZE   CCR2 = 0x03 << 8  //+ Peripheral size These bits are set and cleared by software.
	MSIZE   CCR2 = 0x03 << 10 //+ Memory size These bits are set and cleared by software.
	PL      CCR2 = 0x03 << 12 //+ Channel priority level These bits are set and cleared by software.
	MEM2MEM CCR2 = 0x01 << 14 //+ Memory to memory mode This bit is set and cleared by software.
)

const (
	ENn      = 0
	TCIEn    = 1
	HTIEn    = 2
	TEIEn    = 3
	DIRn     = 4
	CIRCn    = 5
	PINCn    = 6
	MINCn    = 7
	PSIZEn   = 8
	MSIZEn   = 10
	PLn      = 12
	MEM2MEMn = 14
)

const (
	NDT CNDTR2 = 0xFFFF << 0 //+ Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
)

const (
	NDTn = 0
)

const (
	PA CPAR2 = 0xFFFFFFFF << 0 //+ Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
)

const (
	PAn = 0
)

const (
	MA CMAR2 = 0xFFFFFFFF << 0 //+ Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
)

const (
	MAn = 0
)

const (
	EN      CCR3 = 0x01 << 0  //+ Channel enable This bit is set and cleared by software.
	TCIE    CCR3 = 0x01 << 1  //+ Transfer complete interrupt enable This bit is set and cleared by software.
	HTIE    CCR3 = 0x01 << 2  //+ Half transfer interrupt enable This bit is set and cleared by software.
	TEIE    CCR3 = 0x01 << 3  //+ Transfer error interrupt enable This bit is set and cleared by software.
	DIR     CCR3 = 0x01 << 4  //+ Data transfer direction This bit is set and cleared by software.
	CIRC    CCR3 = 0x01 << 5  //+ Circular mode This bit is set and cleared by software.
	PINC    CCR3 = 0x01 << 6  //+ Peripheral increment mode This bit is set and cleared by software.
	MINC    CCR3 = 0x01 << 7  //+ Memory increment mode This bit is set and cleared by software.
	PSIZE   CCR3 = 0x03 << 8  //+ Peripheral size These bits are set and cleared by software.
	MSIZE   CCR3 = 0x03 << 10 //+ Memory size These bits are set and cleared by software.
	PL      CCR3 = 0x03 << 12 //+ Channel priority level These bits are set and cleared by software.
	MEM2MEM CCR3 = 0x01 << 14 //+ Memory to memory mode This bit is set and cleared by software.
)

const (
	ENn      = 0
	TCIEn    = 1
	HTIEn    = 2
	TEIEn    = 3
	DIRn     = 4
	CIRCn    = 5
	PINCn    = 6
	MINCn    = 7
	PSIZEn   = 8
	MSIZEn   = 10
	PLn      = 12
	MEM2MEMn = 14
)

const (
	NDT CNDTR3 = 0xFFFF << 0 //+ Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
)

const (
	NDTn = 0
)

const (
	PA CPAR3 = 0xFFFFFFFF << 0 //+ Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
)

const (
	PAn = 0
)

const (
	MA CMAR3 = 0xFFFFFFFF << 0 //+ Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
)

const (
	MAn = 0
)

const (
	EN      CCR4 = 0x01 << 0  //+ Channel enable This bit is set and cleared by software.
	TCIE    CCR4 = 0x01 << 1  //+ Transfer complete interrupt enable This bit is set and cleared by software.
	HTIE    CCR4 = 0x01 << 2  //+ Half transfer interrupt enable This bit is set and cleared by software.
	TEIE    CCR4 = 0x01 << 3  //+ Transfer error interrupt enable This bit is set and cleared by software.
	DIR     CCR4 = 0x01 << 4  //+ Data transfer direction This bit is set and cleared by software.
	CIRC    CCR4 = 0x01 << 5  //+ Circular mode This bit is set and cleared by software.
	PINC    CCR4 = 0x01 << 6  //+ Peripheral increment mode This bit is set and cleared by software.
	MINC    CCR4 = 0x01 << 7  //+ Memory increment mode This bit is set and cleared by software.
	PSIZE   CCR4 = 0x03 << 8  //+ Peripheral size These bits are set and cleared by software.
	MSIZE   CCR4 = 0x03 << 10 //+ Memory size These bits are set and cleared by software.
	PL      CCR4 = 0x03 << 12 //+ Channel priority level These bits are set and cleared by software.
	MEM2MEM CCR4 = 0x01 << 14 //+ Memory to memory mode This bit is set and cleared by software.
)

const (
	ENn      = 0
	TCIEn    = 1
	HTIEn    = 2
	TEIEn    = 3
	DIRn     = 4
	CIRCn    = 5
	PINCn    = 6
	MINCn    = 7
	PSIZEn   = 8
	MSIZEn   = 10
	PLn      = 12
	MEM2MEMn = 14
)

const (
	NDT CNDTR4 = 0xFFFF << 0 //+ Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
)

const (
	NDTn = 0
)

const (
	PA CPAR4 = 0xFFFFFFFF << 0 //+ Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
)

const (
	PAn = 0
)

const (
	MA CMAR4 = 0xFFFFFFFF << 0 //+ Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
)

const (
	MAn = 0
)

const (
	EN      CCR5 = 0x01 << 0  //+ Channel enable This bit is set and cleared by software.
	TCIE    CCR5 = 0x01 << 1  //+ Transfer complete interrupt enable This bit is set and cleared by software.
	HTIE    CCR5 = 0x01 << 2  //+ Half transfer interrupt enable This bit is set and cleared by software.
	TEIE    CCR5 = 0x01 << 3  //+ Transfer error interrupt enable This bit is set and cleared by software.
	DIR     CCR5 = 0x01 << 4  //+ Data transfer direction This bit is set and cleared by software.
	CIRC    CCR5 = 0x01 << 5  //+ Circular mode This bit is set and cleared by software.
	PINC    CCR5 = 0x01 << 6  //+ Peripheral increment mode This bit is set and cleared by software.
	MINC    CCR5 = 0x01 << 7  //+ Memory increment mode This bit is set and cleared by software.
	PSIZE   CCR5 = 0x03 << 8  //+ Peripheral size These bits are set and cleared by software.
	MSIZE   CCR5 = 0x03 << 10 //+ Memory size These bits are set and cleared by software.
	PL      CCR5 = 0x03 << 12 //+ Channel priority level These bits are set and cleared by software.
	MEM2MEM CCR5 = 0x01 << 14 //+ Memory to memory mode This bit is set and cleared by software.
)

const (
	ENn      = 0
	TCIEn    = 1
	HTIEn    = 2
	TEIEn    = 3
	DIRn     = 4
	CIRCn    = 5
	PINCn    = 6
	MINCn    = 7
	PSIZEn   = 8
	MSIZEn   = 10
	PLn      = 12
	MEM2MEMn = 14
)

const (
	NDT CNDTR5 = 0xFFFF << 0 //+ Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
)

const (
	NDTn = 0
)

const (
	PA CPAR5 = 0xFFFFFFFF << 0 //+ Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
)

const (
	PAn = 0
)

const (
	MA CMAR5 = 0xFFFFFFFF << 0 //+ Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
)

const (
	MAn = 0
)

const (
	EN      CCR6 = 0x01 << 0  //+ Channel enable This bit is set and cleared by software.
	TCIE    CCR6 = 0x01 << 1  //+ Transfer complete interrupt enable This bit is set and cleared by software.
	HTIE    CCR6 = 0x01 << 2  //+ Half transfer interrupt enable This bit is set and cleared by software.
	TEIE    CCR6 = 0x01 << 3  //+ Transfer error interrupt enable This bit is set and cleared by software.
	DIR     CCR6 = 0x01 << 4  //+ Data transfer direction This bit is set and cleared by software.
	CIRC    CCR6 = 0x01 << 5  //+ Circular mode This bit is set and cleared by software.
	PINC    CCR6 = 0x01 << 6  //+ Peripheral increment mode This bit is set and cleared by software.
	MINC    CCR6 = 0x01 << 7  //+ Memory increment mode This bit is set and cleared by software.
	PSIZE   CCR6 = 0x03 << 8  //+ Peripheral size These bits are set and cleared by software.
	MSIZE   CCR6 = 0x03 << 10 //+ Memory size These bits are set and cleared by software.
	PL      CCR6 = 0x03 << 12 //+ Channel priority level These bits are set and cleared by software.
	MEM2MEM CCR6 = 0x01 << 14 //+ Memory to memory mode This bit is set and cleared by software.
)

const (
	ENn      = 0
	TCIEn    = 1
	HTIEn    = 2
	TEIEn    = 3
	DIRn     = 4
	CIRCn    = 5
	PINCn    = 6
	MINCn    = 7
	PSIZEn   = 8
	MSIZEn   = 10
	PLn      = 12
	MEM2MEMn = 14
)

const (
	NDT CNDTR6 = 0xFFFF << 0 //+ Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
)

const (
	NDTn = 0
)

const (
	PA CPAR6 = 0xFFFFFFFF << 0 //+ Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
)

const (
	PAn = 0
)

const (
	MA CMAR6 = 0xFFFFFFFF << 0 //+ Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
)

const (
	MAn = 0
)

const (
	EN      CCR7 = 0x01 << 0  //+ Channel enable This bit is set and cleared by software.
	TCIE    CCR7 = 0x01 << 1  //+ Transfer complete interrupt enable This bit is set and cleared by software.
	HTIE    CCR7 = 0x01 << 2  //+ Half transfer interrupt enable This bit is set and cleared by software.
	TEIE    CCR7 = 0x01 << 3  //+ Transfer error interrupt enable This bit is set and cleared by software.
	DIR     CCR7 = 0x01 << 4  //+ Data transfer direction This bit is set and cleared by software.
	CIRC    CCR7 = 0x01 << 5  //+ Circular mode This bit is set and cleared by software.
	PINC    CCR7 = 0x01 << 6  //+ Peripheral increment mode This bit is set and cleared by software.
	MINC    CCR7 = 0x01 << 7  //+ Memory increment mode This bit is set and cleared by software.
	PSIZE   CCR7 = 0x03 << 8  //+ Peripheral size These bits are set and cleared by software.
	MSIZE   CCR7 = 0x03 << 10 //+ Memory size These bits are set and cleared by software.
	PL      CCR7 = 0x03 << 12 //+ Channel priority level These bits are set and cleared by software.
	MEM2MEM CCR7 = 0x01 << 14 //+ Memory to memory mode This bit is set and cleared by software.
)

const (
	ENn      = 0
	TCIEn    = 1
	HTIEn    = 2
	TEIEn    = 3
	DIRn     = 4
	CIRCn    = 5
	PINCn    = 6
	MINCn    = 7
	PSIZEn   = 8
	MSIZEn   = 10
	PLn      = 12
	MEM2MEMn = 14
)

const (
	NDT CNDTR7 = 0xFFFF << 0 //+ Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
)

const (
	NDTn = 0
)

const (
	PA CPAR7 = 0xFFFFFFFF << 0 //+ Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
)

const (
	PAn = 0
)

const (
	MA CMAR7 = 0xFFFFFFFF << 0 //+ Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
)

const (
	MAn = 0
)

const (
	EN      CCR8 = 0x01 << 0  //+ Channel enable This bit is set and cleared by software.
	TCIE    CCR8 = 0x01 << 1  //+ Transfer complete interrupt enable This bit is set and cleared by software.
	HTIE    CCR8 = 0x01 << 2  //+ Half transfer interrupt enable This bit is set and cleared by software.
	TEIE    CCR8 = 0x01 << 3  //+ Transfer error interrupt enable This bit is set and cleared by software.
	DIR     CCR8 = 0x01 << 4  //+ Data transfer direction This bit is set and cleared by software.
	CIRC    CCR8 = 0x01 << 5  //+ Circular mode This bit is set and cleared by software.
	PINC    CCR8 = 0x01 << 6  //+ Peripheral increment mode This bit is set and cleared by software.
	MINC    CCR8 = 0x01 << 7  //+ Memory increment mode This bit is set and cleared by software.
	PSIZE   CCR8 = 0x03 << 8  //+ Peripheral size These bits are set and cleared by software.
	MSIZE   CCR8 = 0x03 << 10 //+ Memory size These bits are set and cleared by software.
	PL      CCR8 = 0x03 << 12 //+ Channel priority level These bits are set and cleared by software.
	MEM2MEM CCR8 = 0x01 << 14 //+ Memory to memory mode This bit is set and cleared by software.
)

const (
	ENn      = 0
	TCIEn    = 1
	HTIEn    = 2
	TEIEn    = 3
	DIRn     = 4
	CIRCn    = 5
	PINCn    = 6
	MINCn    = 7
	PSIZEn   = 8
	MSIZEn   = 10
	PLn      = 12
	MEM2MEMn = 14
)

const (
	NDT CNDTR8 = 0xFFFF << 0 //+ Number of data to transfer Number of data to be transferred (0 up to 65535). This register can only be written when the channel is disabled. Once the channel is enabled, this register is read-only, indicating the remaining bytes to be transmitted. This register decrements after each DMA transfer. Once the transfer is completed, this register can either stay at zero or be reloaded automatically by the value previously programmed if the channel is configured in auto-reload mode. If this register is zero, no transaction can be served whether the channel is enabled or not.
)

const (
	NDTn = 0
)

const (
	PA CPAR8 = 0xFFFFFFFF << 0 //+ Peripheral address Base address of the peripheral data register from/to which the data will be read/written. When PSIZE is 01 (16-bit), the PA[0] bit is ignored. Access is automatically aligned to a half-word address. When PSIZE is 10 (32-bit), PA[1:0] are ignored. Access is automatically aligned to a word address.
)

const (
	PAn = 0
)

const (
	MA CMAR8 = 0xFFFFFFFF << 0 //+ Memory address Base address of the memory area from/to which the data will be read/written. When MSIZE is 01 (16-bit), the MA[0] bit is ignored. Access is automatically aligned to a half-word address. When MSIZE is 10 (32-bit), MA[1:0] are ignored. Access is automatically aligned to a word address.
)

const (
	MAn = 0
)