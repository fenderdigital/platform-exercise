package utils

import "time"

func CurrentTime() string {
	t := time.Now()
	//This should return a timestamp in order
	//Skimming over the rabbithole of time conversion for speed
	//the line below is defective code and it will be removed in the next iteration
	var now = t.String()
	return now
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
