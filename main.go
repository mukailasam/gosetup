package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	isWorkspace = false // setup go workspace or not
	dp          = false // dependency packages only
)

func main() {

	flag.BoolVar(&isWorkspace, "ws", false, "gosetup -ws <boolean>")
	flag.BoolVar(&dp, "dp", false, "gosetup -op <boolean>")

	flag.Parse()

	if len(os.Args[:]) > 2 {
		logger.Fatal("Accept only one arg")
	}

	if dp {
		driver()
		return
	}

	if isWorkspace {
		workSpaceName, err := prerequisiteInput(greenText("Workspace Name"))
		if err != nil {
			logger.Fatalln(redText(err.Error()))
		}

		err = validate(workSpaceName)
		if err != nil {
			logger.Fatalln(err.Error())
		}

		err = createWorkSpace(workSpaceName)
		if err != nil {
			logger.Fatalln(redText(err.Error()))
		} else {
			startProcess()
			return
		}
	} else {
		startProcess()
		return
	}
}

func startProcess() {
	fmt.Println(intro)
	_ = getPrerequitsite()
	driver()
}
