package push_test

import (
	"fmt"
	"net/http/httptest"

	push "github.com/benaskins/axon-push"
)

func ExampleNewEventBus() {
	bus := push.NewEventBus[string]()
	ch := bus.Subscribe("client-1")
	bus.Publish("hello")
	fmt.Println(<-ch)
	bus.Unsubscribe("client-1")
	// Output: hello
}

func ExampleSetSSEHeaders() {
	w := httptest.NewRecorder()
	push.SetSSEHeaders(w)
	fmt.Println(w.Header().Get("Content-Type"))
	// Output: text/event-stream
}

func ExampleSendEvent() {
	w := httptest.NewRecorder()
	_ = push.SendEvent(w, w, map[string]string{"msg": "hi"})
	fmt.Print(w.Body.String())
	// Output: data: {"msg":"hi"}
}
