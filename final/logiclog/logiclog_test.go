package logiclog

import (
	"strconv"
	"testing"
)

const Path = "tests_log/test_log1.txt"

func WriteTxt(name string, n_max int, t *testing.T) {
	log := NewLog(name, Path, nil)

	if log == nil {
		t.Error("NewLog failed")
	}

	str_begin := "Escribe " + name + " con id: "
	for i := 0; i < n_max; i++ {
		log.Update(str_begin + strconv.Itoa(i))

		if str_begin+strconv.Itoa(i) != log.mc.Log {
			t.Error("Update write failed")
		}
	}

	str_end := "Termina de escribir " + name
	log.Update(str_end)

	if str_end != log.mc.Log {
		t.Error("Update end write failed")
	}

	log.closeFile()
}

func TestLogicLog(t *testing.T) {
	WriteTxt("log1", 3, t)
	WriteTxt("log2", 5, t)
	WriteTxt("log3", 2, t)

	log := NewLog("log_Add", Path, nil)
	s := log.Add()

	if log.mc.Log != "" {
		t.Error("Add failed")
	}

	str := "{\"Map\":{\"log_Add\":1},\"Id\":\"log_Add\",\"Log\":\"\"}"

	if s != str {
		t.Error("Add failed")
	}

	log_aux := NewLog("log_aux", Path, nil)
	log_aux.Msg(s)
	if log_aux.mc.Id != "log_aux" {
		t.Error("Msg failed")
	}

	log.closeFile()
	log_aux.closeFile()
}
