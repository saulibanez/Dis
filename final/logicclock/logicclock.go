package logicclock

import (
	"encoding/json"
	"log"
)

type MessageClock struct {
	Map map[string]int
	Id  string
	Log string
}

func NewMessageClock(id string) *MessageClock {
	mc := &MessageClock{}
	mc.Map = make(map[string]int)
	mc.Map[id] = 0
	mc.Id = id
	mc.Log = ""
	return mc
}

func (mc *MessageClock) Init(i int) {
	mc.Map[mc.Id] = i
}

func (mc *MessageClock) Update() {
	mc.Map[mc.Id]++
}

func (mc *MessageClock) Get() int {
	return mc.Map[mc.Id]
}

func (mc MessageClock) Join(msg MessageClock) (is_mergedd bool) {
	for key, value := range msg.Map {
		val, ok := mc.Map[key]

		if ok && value < val {
			continue
		}

		is_mergedd = true
		mc.Map[key] = value
	}

	return is_mergedd
}

func (mc *MessageClock) Serialize() string {
	msg, err := json.Marshal(mc)

	if err != nil {
		log.Fatal(err)
	}

	return string(msg)
}

func Deserialize(str string) MessageClock {
	var mc MessageClock
	err := json.Unmarshal([]byte(str), &mc)

	if err != nil {
		log.Fatal(err)
	}

	return mc
}
