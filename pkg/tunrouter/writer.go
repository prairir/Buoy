package tunrouter

import (
	"fmt"

	"github.com/songgao/water"
)

func writer(inf *water.Interface, writeq chan []byte) error {
	for {
		pack := <-writeq
		_, err := inf.Write(pack) // possible block, look into dispatching writes to goroutines
		if err != nil {
			//TODO verify or add extra error handling
			// this could include handling on interface close or other error modes
			return fmt.Errorf("tunrouter.writer: %w", err)
		}
	}
}
