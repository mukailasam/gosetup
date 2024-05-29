package main

const (
	reset   = "\033[0m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
)

func redText(text string) string {
	return red + text + reset
}

func greenText(text string) string {
	return green + text + reset
}

func yellowText(text string) string {
	return yellow + text + reset
}

func blueText(text string) string {
	return blue + text + reset
}

func magentaText(text string) string {
	return magenta + text + reset
}

func cyanText(text string) string {
	return cyan + text + reset
}
