package main

import "regexp"

func (app *application) stringValue(value string) bool {
	return value != "" && len(value) > 2
}

func (app *application) intValue(value int) bool {
	return value != 0
}

func (app *application) floatValue(value float64) bool {
	return value != 0.0
}

func (app *application) isEmail(value string) bool {
	mailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return mailRegex.MatchString(value)
}

