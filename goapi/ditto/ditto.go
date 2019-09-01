package ditto

import (
	"encoding/json"
)

// Ditto ...
type Ditto map[string]interface{}

// ByteToMap ...
func ByteToMap(byt []byte) (Ditto, error) {

	// byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// We need to provide a variable where the JSON
	// package can put the decoded data. This
	// `map[string]interface{}` will hold a map of strings
	// to arbitrary data types.

	var dat map[string]interface{}

	// Here's the actual decoding, and a check for
	// associated errors.
	if err := json.Unmarshal(byt, &dat); err != nil {
		return nil, err
	}
	// fmt.Println(dat)
	return dat, nil

	// In order to use the values in the decoded map,
	// we'll need to convert them to their appropriate type.
	// For example here we convert the value in `num` to
	// the expected `float64` type.
	// num := dat["num"].(float64)
	// fmt.Println(num)

	// // Accessing nested data requires a series of
	// // conversions.
	// strs := dat["strs"].([]interface{})
	// str1 := strs[0].(string)
	// fmt.Println(str1)

}
