package main

import "errors"

var (
	errinvalidPicked = errors.New("invalid")
	errNonePicked    = errors.New("none")
	errNoConnection  = errors.New(redText("error while downloading package: check you internet connection\ntry again :)"))
	errCommand       = errors.New("error while running command")
)
