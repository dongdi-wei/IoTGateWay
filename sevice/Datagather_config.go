package sevice

func Log(v ...interface{}) {
	log.Println(v...)
}


func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
}
