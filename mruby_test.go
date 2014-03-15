package mruby

import (
	"testing"
)

func TestNewMrb(t *testing.T) {
	mrb := NewMrb()
	mrb.Close()
}

func TestMrbArena(t *testing.T) {
	mrb := NewMrb()
	defer mrb.Close()

	idx := mrb.ArenaSave();
	mrb.ArenaRestore(idx);
}

func TestMrbLoadString(t *testing.T) {
	mrb := NewMrb()
	defer mrb.Close()

	value, err := mrb.LoadString(`"HELLO"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if value == nil {
		t.Fatalf("should have value")
	}
}