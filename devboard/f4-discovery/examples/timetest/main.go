// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example tesets setting the time.
package main

import (
	"github.com/embeddedgo/x/time"
	"github.com/embeddedgo/x/time/tz"

	_ "github.com/embeddedgo/stm32/devboard/f4-discovery/board/init"
)

func main() {
	for i := 0; i < 3; i++ {
		println("before set:", time.Now().String())
		time.Sleep(time.Second)
	}
	time.Local = &tz.EuropeWarsaw
	set := time.Date(2019, 11, 17, 11, 44, 35, 9991, time.Local)
	time.Set(set)
	println("set: ", set.String())
	for {
		println("after set:", time.Now().String())
		time.Sleep(time.Second)
	}
}
