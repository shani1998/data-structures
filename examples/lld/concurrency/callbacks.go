package concurrency

import (
	"fmt"
	"sync"
	"time"
)

/*
	We want:
	•	A reg_cb method that multiple users can call concurrently.
	•	While an event is in progress, all calls should be queued.
	•	When the event ends, all queued callbacks should execute.
	•	Any future reg_cb calls (after event ends) should execute immediately.
*/

type EventHandler struct {
	mu              sync.Mutex
	eventInProgress bool
	pending         []func()
}

// reg_cb is called by users to register callbacks
func (e *EventHandler) reg_cb(f func()) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.eventInProgress {
		// Event still running → queue the callback
		e.pending = append(e.pending, f)
		fmt.Println("Callback queued during event")
		return
	}

	// Event not running → execute immediately
	fmt.Println("Callback executed immediately")
	go f()
}

// startEvent marks the start of the event
func (e *EventHandler) startEvent() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.eventInProgress = true
	fmt.Println("Event started")
}

// endEvent marks the end of the event and runs queued callbacks
func (e *EventHandler) endEvent() {
	e.mu.Lock()
	pending := e.pending
	e.pending = nil
	e.eventInProgress = false
	e.mu.Unlock()

	fmt.Println("Event ended. Executing queued callbacks...")

	var wg sync.WaitGroup
	for _, cb := range pending {
		wg.Add(1)
		go func(f func()) {
			defer wg.Done()
			f()
		}(cb)
	}
	wg.Wait()
	fmt.Println("All queued callbacks executed.")
}

func main() {
	handler := &EventHandler{}

	handler.startEvent()

	// Simulate users registering callbacks during the event
	go handler.reg_cb(func() { fmt.Println("User 1 callback executed") })
	go handler.reg_cb(func() { fmt.Println("User 2 callback executed") })

	time.Sleep(2 * time.Second) // simulate event duration

	handler.endEvent()

	// Now event is complete → callbacks execute immediately
	handler.reg_cb(func() { fmt.Println("User 3 callback executed immediately") })
}
