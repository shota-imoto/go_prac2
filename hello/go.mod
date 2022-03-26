module example.com/hello

go 1.17

replace example.com/greetings => ../greetings

require example.com/greetings v0.0.0-00010101000000-000000000000

require github.com/fukata/golang-stats-api-handler v1.0.0 // indirect
