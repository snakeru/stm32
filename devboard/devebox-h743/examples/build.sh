#!/bin/sh

GOTARGET=stm32h7x3

DTCM=0x20000000:128K
AXISRAM=0x24000000:512K
AHBSRAM=0x30000000:288K
AHBSRAM4=0x38000000:64K
BKPSRAM=0x38800000:4K

GOTEXT=0x8000000
GOMEM=$AXISRAM,$DTCM

ISRNAMES=no

. ../../../../../scripts/build.sh $@
