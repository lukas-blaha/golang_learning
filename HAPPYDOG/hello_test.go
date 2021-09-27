package hello

import "testing"


func TestHello(t *testing.T) {
    want := "Hello, world."
    if got := Hello(); got != want {
        t.Errorf("Hello() = %qm want %q", got, want)
    }
}
