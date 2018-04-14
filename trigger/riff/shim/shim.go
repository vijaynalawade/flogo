package main

import (
	fl "github.com/vijaynalawade/flogo/trigger/riff"
)


func Invoke(input interface{}) (interface{}, error) {
	result, err := fl.Invoke(input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func main()  {
	
}


