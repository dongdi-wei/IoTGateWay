package wraper

type wrape struct {
}

var Wraper *wrape

func Init()  {
	Wraper = new(wrape)
}
