package sensors

type Sensor struct {
	Value float64 `json:"value"`
	Name  string  `json:"name"`
}

var sensorsList = map[string]Sensor{}

func AddSensor(name string, value float64) {
	sensorsList[name] = Sensor{Name: name, Value: value}
}

func GetSensor(name string) (*Sensor, error) {
	sensor, ok := sensorsList[name]
	if ok {
		return &sensor, nil
	}
	return nil, nil
}
