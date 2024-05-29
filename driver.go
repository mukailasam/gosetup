package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	intro         = cyanText("\nsEtTing up yOur pRojecT...\n")
	pickSomething = redText("either you pick a package, I'm done or Next")
)

/*
func isInstalled(packageName string) {
	fmt.Println(blueText(packageName + "successfully installed ✅"))
}
*/

func finalMessage() {
	fmt.Println(greenText(`✅ Done setting up your Go project :)`))
	os.Exit(99)
}

func prerequisiteInput(label string) (string, error) {
	fmt.Printf("%s :", label)
	buf := bufio.NewReader(os.Stdin)
	val, err := buf.ReadString('\n')
	if err != nil {
		return "", err
	}

	nVal := strings.TrimSpace(val)

	return nVal, nil
}

func getPrerequitsite() string {
	gitUser, err := prerequisiteInput(greenText("Github Username"))
	if err != nil {
		logger.Fatalln(redText(err.Error()))
	}
	err = validate(gitUser)
	if err != nil {
		logger.Fatalln(redText(err.Error()))
	}

	projectName, err := prerequisiteInput(greenText("Project Name"))
	if err != nil {
		logger.Fatalln(redText(err.Error()))
	}
	err = validate(projectName)
	if err != nil {
		logger.Fatalln(redText(err.Error()))
	}

	err = createProject(gitUser, projectName)
	if err != nil {
		logger.Fatalln(redText(err.Error()))
	}

	return projectName
}

func driver() {

	displayer("Dependency Menu", " WebFrameWork packages", " Router packages", " Database packages", " Configuration packages",
		" log Packages", " Cache Packages", " Protocol Packages", " Test Packages", " CLI packages", "ThirdParty packages", "Other packages", "I'm DOne")
	val, err := picker()
	if err != nil {
		return
	}

	switch val {
	case 1:
		err := webFrameWorkListController()
		if err != nil {
			if err.Error() != errinvalidPicked.Error() && err.Error() != errNonePicked.Error() {
				logger.Println(errNoConnection)
				return
			}
			logger.Println(pickSomething)
			return
		} else {
			driver()
		}

	case 2:
		err = routersListController()
		if err != nil {
			if err.Error() != errinvalidPicked.Error() && err.Error() != errNonePicked.Error() {
				logger.Println(errNoConnection)
				return
			}
			logger.Println(pickSomething)
			return
		} else {
			driver()
		}

	case 3:
		err = databasesListController()
		if err != nil {
			if err.Error() != errinvalidPicked.Error() && err.Error() != errNonePicked.Error() {
				logger.Println(errNoConnection)
				return
			}

			logger.Println(pickSomething)
			return
		} else {
			driver()
		}

	case 4:
		err = configurationsListController()
		if err != nil {
			if err.Error() != errinvalidPicked.Error() && err.Error() != errNonePicked.Error() {
				logger.Println(errNoConnection)
				return
			}

			logger.Println(pickSomething)
			return
		} else {
			driver()
		}

	case 5:
		err = logsListController()
		if err != nil {
			if err.Error() != errinvalidPicked.Error() && err.Error() != errNonePicked.Error() {
				logger.Println(errNoConnection)
				return
			}

			logger.Println(pickSomething)
			return
		} else {
			driver()
		}

	case 6:
		err = cachesListController()
		if err != nil {
			if err.Error() != errinvalidPicked.Error() && err.Error() != errNonePicked.Error() {
				logger.Println(errNoConnection)

				return
			}

			logger.Println(pickSomething)
			return
		} else {
			driver()
		}

	case 7:
		err = protocolsListController()
		if err != nil {
			if err.Error() != errinvalidPicked.Error() && err.Error() != errNonePicked.Error() {
				logger.Println(errNoConnection)

				return
			}

			logger.Println(pickSomething)
			return
		} else {
			driver()
		}

	case 8:
		err = testsListController()
		if err != nil {
			if err.Error() != errinvalidPicked.Error() && err.Error() != errNonePicked.Error() {
				logger.Println(errNoConnection)

				return
			}

			logger.Println(pickSomething)
			return
		}

	case 9:
		err = cliListController()
		if err != nil {
			if err.Error() != errinvalidPicked.Error() && err.Error() != errNonePicked.Error() {
				logger.Println(errNoConnection)

				return
			}

			logger.Println(pickSomething)
			return
		} else {
			driver()
		}

	case 10:
		err = thirdPartyListController()
		if err != nil {
			if err.Error() != errinvalidPicked.Error() && err.Error() != errNonePicked.Error() {
				logger.Println(errNoConnection)

				return
			}

			logger.Println(pickSomething)
			return
		} else {
			driver()
		}

	case 11:
		err = otherListController()
		if err != nil {
			if err.Error() != errinvalidPicked.Error() && err.Error() != errNonePicked.Error() {
				logger.Println(errNoConnection)

				return
			}

			logger.Println(pickSomething)
			return
		}
	case 12:
		finalMessage()
	}

}
