package main

import (
	"fmt"
	"strconv"
)

func timeConversion(s string) string {
	hour := s[0:2]
	hour_int, err := strconv.Atoi(s[0:2])
	if err != nil {
		fmt.Println(err)
		return ""
	}

	min_sec := s[3:8]
	ampm := s[8:10]

	var hour_s string
	if ampm == "PM" && hour_int != 12 {
		hour_s = strconv.Itoa(hour_int + 12)
	} else if ampm == "AM" && hour_int == 12 {
		hour_s = "00"
	} else {
		hour_s = hour
	}

	return fmt.Sprintf("%s:%s", hour_s, min_sec)
}

func main() {
	fmt.Println(timeConversion("06:40:03AM"))
}
