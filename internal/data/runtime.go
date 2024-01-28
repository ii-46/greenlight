package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	// In case for json value, if not wrap in quote it will throw runtime error
	quoteJSONValue := strconv.Quote(jsonValue)
	return []byte(quoteJSONValue), nil
}
