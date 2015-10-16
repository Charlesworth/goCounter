package concurrentMap

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

func (concurrentMap *Map) SaveEveryInterval(fileName string, saveInterval time.Duration) {
	ticker := time.NewTicker(saveInterval)
	go func() {
		for _ = range ticker.C {
			concurrentMap.Save(fileName)
		}
	}()
}

func (concurrentMap *Map) Save(fileName string) {
	mapToSave := concurrentMap.GetMap()
	jsonByte, err := mapToJson(mapToSave)
	if err != nil {
		log.Println("Error, unable to Save map with error {", err, "}")
	}

	os.Remove(fileName)
	err = saveByteToFile(jsonByte, fileName)
	if err != nil {
		log.Println("Error, unable to Save map with error {", err, "}")
	}
}

func LoadOrCreateIfDoesntExist(fileName string) (concurrentMap *Map, err error) {
	if !fileExists(fileName) {
		return New(), nil
	}

	jsonMapByte, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, errors.New("LoadMap() failed to read mapSaveFile: " + err.Error())
	}

	loadedMap, err := jsonToMap(jsonMapByte)
	return &Map{loadedMap, &sync.RWMutex{}}, nil
}

func mapToJson(m map[string]int) ([]byte, error) {
	json, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return json, nil
}

func saveByteToFile(data []byte, fileName string) error {
	err := ioutil.WriteFile(fileName, data, 777)
	return err
}

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
