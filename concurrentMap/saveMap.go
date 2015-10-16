package concurrentMap

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
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

func Load() {

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

// func SaveMap(m map[string]int, fileLocation string) error {
// 	jsonMapByte, err := json.Marshal(m)
// 	if err != nil {
// 		return errors.New("SaveMap() failed to marshal map to JSON with error: " + err.Error())
// 	}
// 	err = ioutil.WriteFile(fileLocation, jsonMapByte, 0644)
// 	if err != nil {
// 		return errors.New("SaveMap() failed to write JSON map to file: " + err.Error())
// 	}
// 	return nil
// }

func fileExists(fileName string) bool {
	if _, err := os.Stat(fileName); err == nil {
		return true
	}
	return false
}

func jsonToMap(jsonByte []byte) (map[string]int, error) {
	outputMap := make(map[string]int)
	err := json.Unmarshal(jsonByte, &outputMap)
	if err != nil {
		return nil, errors.New("LoadMap() failed to unmarshal mapSaveFile JSON with error: " + err.Error())
	}
	return outputMap, nil
}

// func LoadMap(fileLocation string) (map[string]int, error) {
// 	if _, err := os.Stat(fileLocation); os.IsNotExist(err) {
// 		return nil, nil
// 	}
//
// 	jsonMapByte, err := ioutil.ReadFile(fileLocation)
// 	if err != nil {
// 		return nil, errors.New("LoadMap() failed to read mapSaveFile: " + err.Error())
// 	}
// 	outputMap := make(map[string]int)
// 	err = json.Unmarshal(jsonMapByte, &outputMap)
// 	if err != nil {
// 		return nil, errors.New("LoadMap() failed to unmarshal mapSaveFile JSON: " + err.Error())
// 	}
// 	return outputMap, nil
// }
