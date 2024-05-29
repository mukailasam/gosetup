package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func picker() (int, error) {
	fmt.Printf("%s", magentaText("Pick: "))
	buf := bufio.NewReader(os.Stdin)
	val, err := buf.ReadString('\n')
	if err != nil {
		return 0, errinvalidPicked
	}

	nVal := strings.TrimSpace(val)

	cVal, err := strconv.Atoi(nVal)
	if err != nil {
		return 0, errinvalidPicked
	}

	return cVal, nil
}
