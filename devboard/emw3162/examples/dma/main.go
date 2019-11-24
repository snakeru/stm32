// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example tests different ways of coping memory. It also shows how to use
// DMA for memory to memory transfers. In case of STM32 F2/F4/F7 only DMA2
// supports MTM transfer.
package main

import (
	"embedded/rtos"
	"unsafe"

	"github.com/embeddedgo/stm32/devboard/emw3162/board"
	"github.com/embeddedgo/stm32/hal/dma"
	"github.com/embeddedgo/stm32/hal/irq"
)

const n = 20 * 1024 / 4

var (
	ch  dma.Channel
	tce rtos.Note

	src = make([]uint32, n)
	dst = make([]uint32, n)
)

func copyDMA(mode dma.Mode) {
	tce.Clear()
	ch.Setup(dma.MTM | dma.IncP | dma.IncM | mode)
	ch.SetWordSize(unsafe.Sizeof(src[0]), unsafe.Sizeof(dst[0]))
	ch.SetLen(n)
	ch.SetAddrP(unsafe.Pointer(&src[0]))
	ch.SetAddrM(unsafe.Pointer(&dst[0]))
	ch.Clear(dma.EvAll, dma.ErrAll)
	ch.EnableIRQ(dma.Complete, dma.ErrAll)
	ch.Enable()
	tce.Sleep(-1)
	if _, err := ch.Status(); err != 0 {
		println(err.Error())
	}
}

func printSpeed(t int64, check bool) {
	t1 := rtos.Nanotime()
	t2 := rtos.Nanotime()
	dt := (t1 - t) - (t2 - t1)
	if check {
		for i := range dst {
			if dst[i] != uint32(i) {
				println(" dst != src\n")
				return
			}
			dst[i] = 0
		}
	}
	println("", (int64(n*unsafe.Sizeof(dst[0]))*1e6+dt/2)/dt, "kB/s")
}

func check() {
	for i := range dst {
		if dst[i] != uint32(i) {
			println(" dst != src\n")
			return
		}
		dst[i] = 0
	}
}

func main() {
	board.Setup(true)

	d := dma.DMA(2)
	d.EnableClock(true)
	ch = d.Channel(0, 0)
	irq.DMA2_Stream0.Enable(rtos.IntPrioLow)

	for {
		print("Initialize src                        ")
		t := rtos.Nanotime()
		for i := range src {
			src[i] = uint32(i)
		}
		printSpeed(t, false)

		print("for i := range src { dst[i] = src[i] }")
		t = rtos.Nanotime()
		for i := range src {
			dst[i] = src[i]
		}
		printSpeed(t, true)

		print("copy(dst, src)                        ")
		t = rtos.Nanotime()
		copy(dst, src)
		printSpeed(t, true)

		print("DMA                                   ")
		t = rtos.Nanotime()
		copyDMA(0)
		printSpeed(t, true)

		print("DMA FT1                               ")
		t = rtos.Nanotime()
		copyDMA(dma.FT1)
		printSpeed(t, true)

		print("DMA FT2                               ")
		t = rtos.Nanotime()
		copyDMA(dma.FT2)
		printSpeed(t, true)

		print("DMA FT3                               ")
		t = rtos.Nanotime()
		copyDMA(dma.FT3)
		printSpeed(t, true)

		print("DMA FT4                               ")
		t = rtos.Nanotime()
		copyDMA(dma.FT4)
		printSpeed(t, true)

		print("DMA FT4 PB4 MB4                       ")
		t = rtos.Nanotime()
		copyDMA(dma.FT4 | dma.PB4 | dma.MB4)
		printSpeed(t, true)

		var delay rtos.Note
		delay.Clear()
		delay.Sleep(2e9)
		println()
	}
}

//go:interrupthandler
func DMA2_Stream0_Handler() {
	ch.DisableIRQ(dma.EvAll, dma.ErrAll)
	tce.Wakeup()
}