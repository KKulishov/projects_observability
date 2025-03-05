package bad

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/grafana/pyroscope-go"
)

// Функция с множеством Goroutines, конкурирующих за mutex
func goroutineContentionFunction() {
	pyroscope.TagWrapper(context.Background(), pyroscope.Labels("function", "goroutine-contention"), func(c context.Context) {
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ { // Запускаем 10 конкурентных горутин
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				mu.Lock()
				defer mu.Unlock()
				fmt.Printf("Goroutine %d acquired lock\n", id)
				time.Sleep(500 * time.Millisecond) // Держим лок 500 мс
				fmt.Printf("Goroutine %d released lock\n", id)
			}(i)
		}
		wg.Wait()
	})
}

func GoroutineHandler(w http.ResponseWriter, r *http.Request) {
	goroutineContentionFunction()
	fmt.Fprintln(w, "Goroutine request")
}
