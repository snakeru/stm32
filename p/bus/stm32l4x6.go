// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32l4x6

package bus

const (
	Core Bus = iota
	AHB1
	AHB2
	AHB3
	APB1
	APB2

	AHBLast = AHB3
	APBLast = APB2
)

type Bus uint8

func (b Bus) String() string {
	i := int(b) * 4
	return "CoreAHB1AHB2AHB3APB1APB2"[i : i+4]
}

var buses [6]struct{ clockHz int64 }

func (b Bus) Clock() int64      { return buses[b].clockHz }
func (b Bus) SetClock(Hz int64) { buses[b].clockHz = Hz }
