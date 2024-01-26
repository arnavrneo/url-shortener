package initializers

import (
	"testing"
)

func TestConnectRedis(t *testing.T) {
	rdb := ConnectRedis()
	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		t.Error("cannot connect redis client")
	}
}
