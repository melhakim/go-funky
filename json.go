package gofunky

import "encoding/json"

func JsonToObj(jsonString []byte) map[string]interface{} {
	var unmarshaled map[string]interface{}
	err := json.Unmarshal(jsonString, &unmarshaled)
	if err != nil {
		panic(err)
	}
	for key, value := range unmarshaled {
		switch value.(type) {
		case string:
			var temp map[string]interface{}
			if json.Unmarshal([]byte(value.(string)), temp) == nil {
				unmarshaled[key] = JsonToObj(value.([]byte))
			} else {
				unmarshaled[key] = value.(string)
			}
		case int:
			unmarshaled[key] = value.(int)
		case bool:
			unmarshaled[key] = value.(bool)
		}
	}
	return unmarshaled
}
