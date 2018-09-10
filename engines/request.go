package engines

type Request struct {
	Url        string
	ParserFunc func(b []byte) (r RequestResult)
}
type RequestResult struct {
	Items []interface{}
	R     []Request
}

func NilParser(b []byte) (r RequestResult) {
	return r
}
