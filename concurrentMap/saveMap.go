package concurrentMap

import (
	"encoding/json"
	"io/ioutil"
)

func (concurrentMap *Map) SaveEveryInterval() {
	// ticker := time.NewTicker(time.Minute) //time.Hour)
	// go func() {
	// 	for _ = range ticker.C {
	// 		pageViewByte := []byte(fmt.Sprint(concurrentMap.GetMap()))
	// 		ioutil.WriteFile("data/savedMap", pageViewByte, 0644)
	// 	}
	// }()
}

func (concurrentMap *Map) Save() {

}

func mapToJson(m map[string]int) ([]byte, error) {
	json, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return json, nil
}

func saveByteToFile(data []byte, fileName string) error {
	err := ioutil.WriteFile(fileName, data, 644)
	return err
}

func Load() {

}
