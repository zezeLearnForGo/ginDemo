package main

type Span struct {
	TraceId     string
	SpanId      string
	ServiceName string
	Ref         []Ref
}

type Ref struct {
	TraceId      string
	ServiceName  string
	SpanId       string
	Relationship string
}

type Relationship struct {
	Parent string
	Child  string
	Count  int
}

func test() {

}
