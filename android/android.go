package groid

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	"github.com/zazab/zhash"
)

type android struct {
	frontend Frontend
	input    chan map[string]interface{}
}

func (android *android) Run() {
	views := zhash.HashFromMap(android.GetViewsList("useless_layout"))

	viewId, err := views.GetString("resources", "useless_accel", "id")
	if err != nil {
		log.Printf("!!! %#v", err)
	}

	sensors := zhash.HashFromMap(android.GetSensorsList())
	sensorId, err := sensors.GetString("sensors", "TYPE_ACCELEROMETER")
	if err != nil {
		log.Printf("!!! %#v", err)
	}

	android.SubscribeToSensorValues(sensorId)

	log.Printf("!!! %#v", viewId)

	_ = android.CallViewMethod(
		viewId,
		"TextView",
		"setTextSize", 0, float(60.0),
	)

	buttonId, err := views.GetString("resources", "useless_button", "id")
	if err != nil {
		log.Printf("!!! %#v", err)
	}

	server.SubscribeToViewEvent(buttonId, "Button", "onClick")

	for {
		payload := <-android.input

		data := zhash.HashFromMap(payload)

		log.Printf("!!! %#v", data)

		eventName, err := data.GetString("event")
		if err != nil {
			continue
		}

		switch eventName {
		case "sensorChange":
			values, _ := data.GetFloatSlice("data", "values")

			_ = android.CallViewMethod(
				viewId,
				"TextView",
				"setText", fmt.Sprintf(
					"% 3.2f\n% 3.2f\n% 3.2f",
					values[0], values[1], values[2],
				),
			)

		case "click":
			texts := []string{
				"how dare you",
				"i say so",
				"i don't mine",
				"wazzzzup",
				"don't click on me, you stupid chicken!",
			}
			viewId, _ := data.GetString("data", "view_id")
			_ = android.CallViewMethod(
				viewId,
				"TextView",
				"setText", texts[rand.Intn(len(texts))],
			)
		}
	}
}

func (android *android) Handle(data map[string]interface{}) {
	android.input <- data
}

func (android *android) Call(
	payload map[string]interface{},
) map[string]interface{} {
	payloadJson, _ := json.Marshal(payload)
	//log.Printf("!!! SENT %#v", string(payloadJson))

	resultJson := android.frontend.CallFrontend(string(payloadJson))

	var result map[string]interface{}
	err := json.Unmarshal([]byte(resultJson), &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (android *android) GetViewsList(
	layoutName string,
) map[string]interface{} {
	return android.Call(map[string]interface{}{
		"method": "ListViews",
		"layout": layoutName,
	})
}

func (android *android) SubscribeToSensorValues(
	sensorId string,
) map[string]interface{} {
	return android.Call(map[string]interface{}{
		"method":    "SubscribeToSensorValues",
		"sensor_id": sensorId,
	})
}

func (android *android) GetSensorsList() map[string]interface{} {
	return android.Call(map[string]interface{}{
		"method": "GetSensorsList",
	})
}

func (android *android) CallViewMethod(
	id string,
	viewType string,
	methodName string,
	args ...interface{},
) map[string]interface{} {
	return android.Call(map[string]interface{}{
		"method":     "CallViewMethod",
		"id":         id,
		"type":       viewType,
		"viewMethod": methodName,
		"args":       args,
	})
}

func (server *android) SubscribeToViewEvent(
	id string,
	viewType string,
	event string,
) map[string]interface{} {
	return server.Call(map[string]interface{}{
		"method": "SubscribeToViewEvent",
		"id":     id,
		"type":   viewType,
		"event":  event,
	})
}
