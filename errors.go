package backend_errors

import "encoding/json"

func BadInputsJSON(fields map[string]string) ([]byte, error) {
	badInputData := badInput{}
	for key, element := range fields {
		badInputData.Fields = append(badInputData.Fields, field{
			Name:   key,
			Reason: element,
		})
	}

	return json.Marshal(badInputData)
}

func BadInputJSON(name string, reason string) ([]byte, error) {
	return BadInputsJSON(map[string]string{ name: reason })
}