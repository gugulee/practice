package mergemap

import (
	"encoding/json"
	"os"
	"testing"
)

func Test_Merge(t *testing.T) {
	base, err := os.ReadFile("/tmp/base.json")
	if err != nil {
		panic(err)
	}

	target, err := os.ReadFile("/tmp/target.json")
	if err != nil {
		panic(err)
	}

	// Create maps to hold the JSON objects
	var src map[string]interface{}
	var dst map[string]interface{}

	// Unmarshal JSON strings into maps
	if err := json.Unmarshal([]byte(base), &src); err != nil {
		panic(err)
	}

	if err := json.Unmarshal([]byte(target), &dst); err != nil {
		panic(err)
	}

	// Merge the two maps
	merged := Merge(src, dst)
	// merged := mergeMaps(obj1, obj2)

	// Convert the merged map back to JSON
	mergedJSON, err := json.Marshal(merged)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("/tmp/merge.json", mergedJSON, 0644)
	if err != nil {
		panic(err)
	}
}
