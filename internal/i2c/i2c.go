package i2c

import (
	"time"

	"github.com/quhar/bme280"
	"golang.org/x/exp/io/i2c"
)

type RoomEnv struct {
	Device string  `json:"devicename"`
	Time   string  `json:"timestamp"`
	Temp   float64 `json:"temp"`
	Press  float64 `json:"press"`
	Hum    float64 `json:"hum"`
}

func GetRoomEnv() RoomEnv {

	d, err := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-1"}, 0x76)
	if err != nil {
		panic(err)
	}

	b := bme280.New(d)
	err = b.Init()

	t, p, h, err := b.EnvData()
	if err != nil {
		panic(err)
	}

	ti := time.Now()

	return RoomEnv{
		Time:  ti.Format("2006/01/02 15:04:05"),
		Temp:  t,
		Press: p,
		Hum:   h,
	}
}
