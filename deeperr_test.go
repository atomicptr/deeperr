package deeperr

import (
	"errors"
	"strings"
	"testing"
)

const (
	errSomethingBadHappened Code = iota + 100
	errAnotherBadThingHappened
	errYetAnotherBadThing
)

func TestNew(t *testing.T) {
	err := New("test error", nil)

	if err.Error() != "test error" {
		t.Errorf("expected 'test error', got %q", err.Error())
	}

	if err.Code() != CodeUnset {
		t.Errorf("expected CodeUnset, got %d", err.Code())
	}

	if err.Unwrap() != nil {
		t.Errorf("expected nil unwrap")
	}
}

func TestNewWithPrev(t *testing.T) {
	prev := errors.New("prev error")

	err := New("test error", prev)

	if !errors.Is(err, prev) {
		t.Errorf("expected prev error")
	}
}

func TestNewWithCode(t *testing.T) {
	err := NewWithCode(42, "test error", nil)

	if err.Code() != 42 {
		t.Errorf("expected code 42, got %d", err.Code())
	}

	if !strings.Contains(err.Error(), "E42") {
		t.Errorf("expected error string to contain E42")
	}
}

func TestNewWithCodeAndPrev(t *testing.T) {
	prev := errors.New("prev error")
	err := NewWithCode(10, "test error", prev)

	if err.Code() != 10 {
		t.Errorf("expected code 10, got %d", err.Code())
	}

	if !errors.Is(err, prev) {
		t.Errorf("expected prev error")
	}
}

func TestIsCode(t *testing.T) {
	err := NewWithCode(errSomethingBadHappened, "test", nil)

	if !IsCode(err, errSomethingBadHappened) {
		t.Error("expected IsCode true")
	}

	if IsCode(err, errAnotherBadThingHappened) {
		t.Error("expected IsCode false")
	}
}

func TestIsCodeNonDeeperrError(t *testing.T) {
	err := errors.New("plain error")

	if IsCode(err, errSomethingBadHappened) {
		t.Error("expected IsCode false for plain error")
	}
}

func TestIsCodeNil(t *testing.T) {
	if IsCode(nil, errSomethingBadHappened) {
		t.Error("expected IsCode false for nil")
	}
}

func TestContains(t *testing.T) {
	err := NewWithCode(errSomethingBadHappened, "test", nil)

	if !Contains(err, errSomethingBadHappened) {
		t.Error("expected Contains true")
	}

	if Contains(err, errAnotherBadThingHappened) {
		t.Error("expected Contains false")
	}
}

func TestContainsNested(t *testing.T) {
	inner := NewWithCode(errAnotherBadThingHappened, "inner", nil)
	err := NewWithCode(errSomethingBadHappened, "outer", inner)

	if !Contains(err, errSomethingBadHappened) {
		t.Error("expected Contains true for nested")
	}
}

func TestContainsNonDeeperrError(t *testing.T) {
	plain := errors.New("plain")
	err := New("test", plain)

	if Contains(err, errSomethingBadHappened) {
		t.Error("expected Contains false")
	}
}

func TestContainsNil(t *testing.T) {
	if Contains(nil, errSomethingBadHappened) {
		t.Error("expected Contains false for nil")
	}
}

func TestErrorMessageAndLocation(t *testing.T) {
	err := New("test message", nil)

	if err.Message() != "test message" {
		t.Errorf("expected 'test message', got %q", err.Message())
	}

	file, line := err.Location()
	if file == "" || line == -1 {
		t.Error("expected valid location")
	}
}
