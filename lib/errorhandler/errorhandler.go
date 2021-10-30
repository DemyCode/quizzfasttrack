package errorhandler

import "log"

func ErrorHandler(err error, fatal bool){
	if err != nil{
		if fatal {
			log.Fatalln(err)
		}
		log.Println(err)
	}
}