package logicclock

import (
	"testing"
)

func TestMessageClock(t *testing.T) {
	var name1, name2 = "hola", "caracola"
	msg1 := NewMessageClock(name1)
	msg2 := NewMessageClock(name2)

	msg1.Update()
	msg2.Update()

	if msg1.Get() != 1 || msg2.Get() != 1 {
		t.Error("Test failed update")
	}

	msg1.Init(3)

	if msg1.Get() != 3 {
		t.Error("Test failed init")
	}

	if !msg1.Join(*msg2) {
		t.Error("Test failed Join")
	}

	str := "{\"Map\":{\"caracola\":1,\"hola\":3},\"Id\":\"hola\",\"Log\":\"\"}"

	serialize := msg1.Serialize()
	if serialize != str {
		t.Error("Test failed Serialize")
	}

	d := Deserialize(serialize)
	if d.Id != name1 {
		t.Error("Test failed Deserialize")
	}

}
