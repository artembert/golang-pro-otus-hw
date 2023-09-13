package hw04lrucache

import (
	"encoding/json"
	"fmt"
)

func DebugMap(collection map[Key]*ListItem) {
	bs, _ := json.Marshal(collection)
	fmt.Println(string(bs))
}
