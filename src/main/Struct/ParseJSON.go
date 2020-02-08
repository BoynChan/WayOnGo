package main

import (
	"encoding/json"
	"fmt"
)

type Screen struct {
	Size       float32 // 荧幕尺寸
	ResX, ResY int     // 荧幕的水平,竖直解析度
}

type Battery struct {
	Capacity int // 容量
}

func genJSON() []byte {
	raw := &struct {
		Screen
		Battery
		HasTouchID bool
	}{
		Screen: Screen{
			Size: 5.9,
			ResX: 1920,
			ResY: 1080,
		},
		Battery:    Battery{Capacity: 2910},
		HasTouchID: false,
	}
	jsonData, _ := json.Marshal(raw)

	return jsonData
}
func main() {
	jsonData := genJSON()
	fmt.Println(string(jsonData))

	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}

	json.Unmarshal(jsonData, &screenAndTouch)

	fmt.Printf("%+v\n", screenAndTouch)
}
