package bad

import (
	"context"
	"fmt"
	"net/http"
	"runtime/pprof"
	"sync"
	"time"

	"math/rand/v2"

	"github.com/grafana/pyroscope-go"
)

var (
	mu       sync.Mutex
	memLeaks [][]byte // массив для утечек памяти
)

//go:noinline
func work(n int) {
	for i := 0; i < n; i++ {
	}
}

func fastFunction(c context.Context) {
	mu.Lock()
	time.Sleep(50 * time.Millisecond) // Искусственная задержка с блокировкой
	pyroscope.TagWrapper(c, pyroscope.Labels("function", "fast"), func(c context.Context) {
		work(20000000)
	})

	mu.Unlock()
}

func slowFunction(c context.Context) {
	mu.Lock()
	time.Sleep(500 * time.Millisecond) // Искусственная задержка с блокировкой
	pprof.Do(c, pprof.Labels("function", "slow"), func(c context.Context) {
		work(80000000)
	})

	mu.Unlock()
}

func MemLeakHandler(w http.ResponseWriter, r *http.Request) {
	leak := make([]byte, 5*1024*1024) // 5MB утечки
	for i := range leak {
		leak[i] = byte(rand.IntN(256))
	}
	mu.Lock()
	memLeaks = append(memLeaks, leak)
	mu.Unlock()
	fmt.Fprintf(w, "Allocated additional 5MB. Total leaks: %d\n", len(memLeaks))
}

func FastHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fastFunction(ctx)
	fmt.Fprintln(w, "Fast request completed")
}

func SlowHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slowFunction(ctx)
	fmt.Fprintln(w, "Slow request completed")
}
