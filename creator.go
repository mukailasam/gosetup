package main

import (
	"fmt"
	"os"
)

func isItGoProject(val int) *bool {
	isIt := true
	isNot := false
	switch val {
	case 1:
		return &isIt
	case 2:
		return &isNot
	default:
		return nil

	}
}

func createProject(username, projectName string) error {
	err := os.Mkdir(projectName, 0777)
	if err != nil {
		if os.IsExist(err) {
			fmt.Printf("%s", fmt.Sprintf(redText("A directory with the name %s already exist."), projectName))
			displayer("Are you sure the directory is a Go project directory?", "Yes", "No")
			val, err := picker()
			if err != nil {
				return err
			}
			isIt := isItGoProject(val)
			if isIt != nil {
				if *isIt {
					err = walker(projectName)
					if err != nil {
						return err
					}
					return nil
				} else {
					fmt.Println(greenText("How about you create a new project, Since you're not sure if existed one is a Go project or not"))
					os.Exit(98)
				}
			} else {
				logger.Fatalln("Wrong input")
			}
		}

		return err
	}

	if isWorkspace {
		err = runner("work", "use", projectName)
		if err != nil {
			return errCommand
		}
	}

	err = walker(projectName)
	if err != nil {
		return err
	}

	err = runner("mod", "init", fmt.Sprintf("github.com/%s/%s", username, projectName))
	if err != nil {
		return errCommand
	}

	if isWorkspace {
		err = walker("..")
		if err != nil {
			return err
		}
		err = runner("work", "use", projectName)
		if err != nil {
			return errCommand
		}
	}

	return nil
}

func createWorkSpace(name string) error {
	err := os.Mkdir(name, 0777)
	if err != nil {
		if os.IsExist(err) {
			fmt.Printf("%s", fmt.Sprintf(redText("A directory with the name %s already exist."), name))
			displayer("Are you sure the directory is a Go Workspace directory?", "Yes", "No")
			val, err := picker()
			if err != nil {
				return err
			}
			isIt := isItGoProject(val)
			if isIt != nil {
				if *isIt {
					err = walker(name)
					if err != nil {
						return err
					}

					return nil
				} else {
					fmt.Println(greenText("How about you create a new Go Workspace, Since you're not sure if existed one is a Go Workspace or not"))
					os.Exit(98)
				}
			} else {
				logger.Fatalln(redText("Wrong input"))
			}
		}

		return err
	}

	err = walker(name)
	if err != nil {
		return err
	}

	err = runner("work", "init")
	if err != nil {
		return errCommand
	}

	return nil

}
