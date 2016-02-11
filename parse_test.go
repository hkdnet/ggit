package ggit

import (
	"testing"
)

func TestAction(t *testing.T) {
	line := "099155237c5d02ab00450239f6c68327c87e4cc2 c887428e5ce02a193c57bfdf4348c17dc7f74da7 hkdnet <hkdnet@users.noreply.github.com> 1455122575 +0900	commit: add makefile"
	log := Parse(line)
	actual := log.Action
	expected := "commit"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
func TestMessage(t *testing.T) {
	line := "099155237c5d02ab00450239f6c68327c87e4cc2 c887428e5ce02a193c57bfdf4348c17dc7f74da7 hkdnet <hkdnet@users.noreply.github.com> 1455122575 +0900	commit: add makefile"
	log := Parse(line)
	actual := log.Message
	expected := "add makefile"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
func TestInitialFalse(t *testing.T) {
	line := "099155237c5d02ab00450239f6c68327c87e4cc2 c887428e5ce02a193c57bfdf4348c17dc7f74da7 hkdnet <hkdnet@users.noreply.github.com> 1455122575 +0900	commit: add makefile"
	log := Parse(line)
	actual := log.IsInitial
	expected := false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
func TestInitialTrue(t *testing.T) {
	line := "0000000000000000000000000000000000000000 b1066e080ac7c085085cc24956ffef4c425675c1 hkdnet <hkdnet@users.noreply.github.com> 1455116947 +0900	commit (initial): initialized with gcli"
	log := Parse(line)
	actual := log.IsInitial
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestPrevHash(t *testing.T) {
	line := "0000000000000000000000000000000000000000 b1066e080ac7c085085cc24956ffef4c425675c1 hkdnet <hkdnet@users.noreply.github.com> 1455116947 +0900	commit (initial): initialized with gcli"
	log := Parse(line)
	actual := log.PrevHash
	expected := "0000000000000000000000000000000000000000"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
func TestHash(t *testing.T) {
	line := "0000000000000000000000000000000000000000 b1066e080ac7c085085cc24956ffef4c425675c1 hkdnet <hkdnet@users.noreply.github.com> 1455116947 +0900	commit (initial): initialized with gcli"
	log := Parse(line)
	actual := log.Hash
	expected := "b1066e080ac7c085085cc24956ffef4c425675c1"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
func TestName(t *testing.T) {
	line := "0000000000000000000000000000000000000000 b1066e080ac7c085085cc24956ffef4c425675c1 hkdnet <hkdnet@users.noreply.github.com> 1455116947 +0900	commit (initial): initialized with gcli"
	log := Parse(line)
	actual := log.Name
	expected := "hkdnet"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
func TestEmail(t *testing.T) {
	line := "0000000000000000000000000000000000000000 b1066e080ac7c085085cc24956ffef4c425675c1 hkdnet <hkdnet@users.noreply.github.com> 1455116947 +0900	commit (initial): initialized with gcli"
	log := Parse(line)
	actual := log.Email
	expected := "hkdnet@users.noreply.github.com"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
func TestAt(t *testing.T) {
	line := "0000000000000000000000000000000000000000 b1066e080ac7c085085cc24956ffef4c425675c1 hkdnet <hkdnet@users.noreply.github.com> 1455116947 +0900	commit (initial): initialized with gcli"
	log := Parse(line)
	actual := log.At
	expected := "1455116947"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestTimeZone(t *testing.T) {
	line := "0000000000000000000000000000000000000000 b1066e080ac7c085085cc24956ffef4c425675c1 hkdnet <hkdnet@users.noreply.github.com> 1455116947 +0900	commit (initial): initialized with gcli"
	log := Parse(line)
	actual := log.TimeZone
	expected := "+0900"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
