package serde

import "encoding/xml"

// Serialize obj into it's xml representation as a string.
// If obj is nil, return empty-string.
func Serialize(obj interface{}) string {
	if obj == nil {
		return ""
	}
	bytes, _ := xml.Marshal(obj)
	return string(bytes)
}

func Deserialize(rawxml string, target interface{}) interface{} {
	xml.Unmarshal([]byte(rawxml), target)
	return target
}
