package utils

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"log"
	mrand "math/rand"
	"reflect"
	"time"
)

func GenerateRandomID() string {
	b := make([]byte, 4) // 4 bytes == 8 hex characters
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(b)
}

func ShuffleSlice(slice []string) {
	mrand.New(mrand.NewSource(time.Now().UnixNano()))
	mrand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

func ConvertStructToJsonString(_structPointer interface{}) []byte {
	data, err := json.Marshal(_structPointer)
	if err != nil {
		log.Printf("Failed to convert struct(%+v) to json byte array: %+v", _structPointer, err)
		return []byte{}
	}
	return data
}

func ConvertJsonStringToStruct(_structPointer interface{}, data []byte) {
	err := json.Unmarshal(data, _structPointer)
	if err != nil {
		log.Printf("Failed to convert json data to structure. %+v", err)
	}
}

func IsStructInSlice(val interface{}, slice interface{}) bool {
	s := reflect.ValueOf(slice)
	for i := 0; i < s.Len(); i++ {
		if reflect.DeepEqual(val, s.Index(i).Interface()) {
			return true
		}
	}
	return false
}
