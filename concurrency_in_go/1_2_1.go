package main

import (
	"fmt"
	"time"
)

func main () {
	var data int
	go func() {
		data++
	} ()
	//$B0J2<$N(BSLEEP$B=hM}$rM-8z$K$9$k$H(Bif$BJ8$NI>2AA0$K(Bdata$B$r2C;;$9$k(B
	//$B4|BTCM$O(B "$BI=<($J$7(B"
	//time.Sleep(1*time.Second)
	if data == 0 {
		//$B0J2<$N(BSLEEP$B=hM}$rM-8z$K$9$k$H(Bif$BJ8I>2A8e$K(Bdata$B$r2C;;$9$k(B
		//$B4|BTCM$O(B "the value is 1."
		//time.Sleep(1*time.Second)
		fmt.Printf("the value is %v.\n", data)
	}
}
