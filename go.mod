module github.com/JoeyPilla/redis-go

go 1.14

require (
	github.com/JoeyPilla/redis-go/redis v0.0.0
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/go-redis/redis/v8 v8.2.3
)

replace github.com/JoeyPilla/redis-go/redis => ./redis
