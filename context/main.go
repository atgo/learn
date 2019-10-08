package context

import (
	"context"
	"fmt"
	"time"
)

func main() {
	bgx := context.Background()
	ctx, can := context.WithCancel(bgx)

	go going(ctx, 1)

	can()
	time.Sleep(2 * time.Second)
}

func going(ctx context.Context, info int) {
	fmt.Println("starting", info)

	if info < 5 {
		go going(ctx, info+1)
	}

	<-ctx.Done()
	fmt.Println("stopped", info, "err", ctx.Err())
}
