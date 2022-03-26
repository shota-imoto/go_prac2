module chitchat

go 1.17

replace example.com/data => ./data

require example.com/data v0.0.0-00010101000000-000000000000

require (
	github.com/lib/pq v1.10.4 // indirect
	github.com/sausheong/gwp v0.0.0-20170211043623-e25181d4e9ed // indirect
)
