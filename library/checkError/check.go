package checkError

import "log"

func CheckErrPrint(e error, name string) {
	if e != nil {
		log.Printf("[Bot]: %.20s |FAIL|\n%s", name, e.Error())
		return
	}
	log.Printf("[Bot]: %.20s |OK|\n", name)
}

func CheckErrPanic(e error, name string) {
	if e != nil {
		log.Panicf("[Bot]: %.20s |FAIL|\n%s", name, e.Error())
		return
	}
	log.Printf("[Bot]: %.20s |OK|\n", name)
}
