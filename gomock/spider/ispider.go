package spider

type intervalSpider struct {
}

func (s intervalSpider) GetVersion() string {
	return "v 1.0.0"
}

func NewSpider() Spider {
	return intervalSpider{}
}
