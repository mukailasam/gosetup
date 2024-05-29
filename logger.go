package main

import (
	"log"
	"os"
)

var logger = log.New(os.Stderr, redText("Error: "), log.Lshortfile)
