package tracer

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("returned value from New is null")
	} else {
		tracer.Trace("Hello tracer package!")
		if buf.String() != "Hello tracer package!\n" {
			t.Errorf("'%s' is output mistakenly", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("data")
	// just check if called without errors
}
