package kartola

import (
	"encoding/json"
	"log"
)

func JsonBuild(i interface{}) []byte {
	r, err := json.Marshal(i)
	if err != nil {
		log.Println(err)
		return []byte("")
	}else{
		return r
	}
}
