module a1_timeout

go 1.19

require github.com/redis/go-redis/v9 v9.0.4

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
)

replace github.com/redis/go-redis/v9 v9.0.4 => github.com/ethanvc/go-redis/v9 v9.0.0-20230512105823-93302afc51a1
