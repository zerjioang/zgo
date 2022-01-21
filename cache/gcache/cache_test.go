package gcache

import (
	"fmt"
	"log"
	"testing"
)

func TestBuildCache(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		var peer CachePeer
		peer.Start()
		// make a get request to the cache system

		var user int64 = 56
		// 1 store value in the cache
		if err := peer.Set("user_id", user); err != nil {
			log.Fatal(err)
		}

		// 2 read value from the cache
		if err := peer.Get("user_id", &user); err != nil {
			log.Fatal(err)
		}
		fmt.Println("user value is:", user)

		//  3 remove the value from cache
		if err := peer.Remove("user_id"); err != nil {
			log.Fatal(err)
		}

		defer peer.Stop()
	})
}
