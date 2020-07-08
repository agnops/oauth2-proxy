package providers

import (
	"github.com/go-redis/redis"
	"os"
)

func PushOrgUserDetails(accessToken string, scmProvider string, userOrg string) {
	redisMasterHost := "localhost"
	redisPass := ""

	if (os.Getenv("redisMasterHost") != redisMasterHost) && (len(os.Getenv("redisMasterHost")) != 0) {
		redisMasterHost = os.Getenv("redisMasterHost")
	}
	if (os.Getenv("redisPass") != redisPass) && (len(os.Getenv("redisPass")) != 0) {
		redisPass = os.Getenv("redisPass")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisMasterHost + ":6379",
		Password: redisPass,
		DB:       0,
	})

	err := client.Set(scmProvider + "-" + userOrg, accessToken, 0).Err() //keep org or user key updated with latest oauth token
	if err != nil {
		panic(err)
	}
}