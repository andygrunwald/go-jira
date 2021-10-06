package jira

import (
	"io"
	"log"
	"os"
)

func BodyLog(body io.ReadCloser){
	if body!=nil && os.Getenv("DEBUG_REQUEST")=="YES"{
		data, _ := io.ReadAll(body)
		body.Close()
		log.Println(string(data))
		data=nil
	}
}