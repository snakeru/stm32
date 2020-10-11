// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32h7x3

// Package vrefbuf provides access to the registers of the VREFBUF peripheral.
//
// Instances:
//  VREFBUF  VREFBUF_BASE  -  -  VREFBUF
// Registers:
//  0x000 32  CSR  VREFBUF control and status register
//  0x004 32  CCR  VREFBUF calibration control register
// Import:
//  github.com/embeddedgo/stm32/p/mmap
package vrefbuf

const (
	ENVR CSR = 0x01 << 0 //+ Voltage reference buffer mode enable This bit is used to enable the voltage reference buffer mode.
	HIZ  CSR = 0x01 << 1 //+ High impedance mode This bit controls the analog switch to connect or not the VREF+ pin. Refer to Table196: VREF buffer modes for the mode descriptions depending on ENVR bit configuration.
	VRR  CSR = 0x01 << 3 //+ Voltage reference buffer ready
	VRS  CSR = 0x07 << 4 //+ Voltage reference scale These bits select the value generated by the voltage reference buffer. Other: Reserved
)

const (
	ENVRn = 0
	HIZn  = 1
	VRRn  = 3
	VRSn  = 4
)

const (
	TRIM CCR = 0x3F << 0 //+ Trimming code These bits are automatically initialized after reset with the trimming value stored in the Flash memory during the production test. Writing into these bits allows to tune the internal reference buffer voltage.
)

const (
	TRIMn = 0
)
