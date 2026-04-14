package main

import (
	"encoding/json"
	"fmt"
)

func RaiseEvent(stub hypConnect, eventName string, args ...interface{}) (string, error) {
	var eventList generalEventStruct
	eventList.EventName = eventName
	eventList.EventList = stub.EventList
	eventList.AdditionalData = args
	fmt.Println("SetEvent", eventList)
	eventJSONasBytes, err2 := json.Marshal(eventList)
	if err2 != nil {
		return "", err2
	}
	fmt.Println("Event raised: " + eventName)
	fmt.Println("eventJSONasBytes : ", eventJSONasBytes)
	mEventName := eventList.EventName
	err3 := stub.Connection.SetEvent("ChainCodeEvent", []byte(eventJSONasBytes))
	fmt.Printf("eventJSONasBytes : ", err3)

	if err3 != nil {
		return "", err3
	}
	var err4 error
	err4 = nil
	return mEventName, err4

}
