package sensors

import "sync"

type Sensor struct {
	Value float64 `json:"value"`
	Name  string  `json:"name"`
}

var mutex sync.Mutex
var sensorsList = map[string]Sensor{}

func AddSensor(name string, value float64) {
	mutex.Lock()
	defer mutex.Unlock()
	sensorsList[name] = Sensor{Name: name, Value: value}
}

func GetSensor(name string) (*Sensor, error) {
	mutex.Lock()
	defer mutex.Unlock()
	sensor, ok := sensorsList[name]
	if ok {
		return &sensor, nil
	}
	return nil, nil
}
