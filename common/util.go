package common

import "os"

func End(msg string, errs ...error){
	println("")
	println("Exiting with message: " + msg)
	println("")
	if len(errs) > 0 && errs[0] != nil {
		println("Error(s):")
		for i := 0; i < len(errs); i++ {
			println(errs[i])
		}
	}
	println("")
	println("Exiting")
	os.Exit(1)
}
