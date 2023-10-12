package server

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"log"
	"reflect"
)

func generateRandomID() string {
	b := make([]byte, 4) // 4 bytes == 8 hex characters
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(b)
}

func isStructInSlice(val interface{}, slice interface{}) bool {
	s := reflect.ValueOf(slice)
	for i := 0; i < s.Len(); i++ {
		if reflect.DeepEqual(val, s.Index(i).Interface()) {
			return true
		}
	}
	return false
}

func convertStructToJsonString(_struct interface{}) string {
	data, err := json.Marshal(_struct)
	if err != nil {
		log.Printf("Failed to convert struct(%+v) to json byte array: %+v", _struct, err)
		return ""
	}
	return string(data)
}
