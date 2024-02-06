package cmd

import "time"

// function takes in a time zone string and returns the current time in string format and an error
func getTimeInTimezone(timezone string) (string, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return "", err
	}
	currentTime := time.Now().In(location)
	return currentTime.Format(time.RFC1123), nil
}
