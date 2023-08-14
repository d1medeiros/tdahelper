package main

import "tdahelper/cmd"

func main() {

	err := cmd.Execute()
	if err != nil {
		panic(err)
	}

}
