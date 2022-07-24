module github.com/khteh/restapi

go 1.18

replace github.com/khteh/fibonacci => ./fibonacci

require (
	github.com/khteh/fibonacci v0.0.0-00010101000000-000000000000
	github.com/khteh/greetings v0.0.0-00010101000000-000000000000
)

require (
	golang.org/x/text v0.3.6 // indirect
	rsc.io/quote v1.5.2 // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace github.com/khteh/greetings => ./greetings
