package gcache

import (
	"context"
	"encoding/json"
	"github.com/mailgun/groupcache/v2"
	"log"
	"net/http"
	"time"
)

type CachePeer struct {
	cacheServer *http.Server
	cacheGroup  *groupcache.Group
}

func (peer *CachePeer) Stop()  error {
	return peer.cacheServer.Shutdown(context.Background())
}

func (peer *CachePeer) Set(id string, value interface{}) error {
	// create a timeout call to check if data is in the cache
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()

	// Set the value in the groupcache to expire after 5 minutes
	raw, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return peer.cacheGroup.Set(ctx, id, raw, time.Now().Add(time.Minute*5), true)
}

func (peer *CachePeer) Get(itemId string, dest interface{}) error {
	// create a timeout call to check if data is in the cache
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()
	var data []byte
	reader := groupcache.AllocatingByteSliceSink(&data)
	if err := peer.cacheGroup.Get(ctx, itemId, reader); err != nil {
		return err
	}
	// handle readed data
	if !(data == nil || len(data)==0) {
		return json.Unmarshal(data, dest)
	}
	return nil
}

func (peer *CachePeer) Remove(itemId string) error {
	// create a timeout call to check if data is in the cache
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()
	// Remove the key from the groupcache
	return peer.cacheGroup.Remove(ctx, itemId)
}

func (peer *CachePeer) Start() {
	// NOTE: It is important to pass the same peer `http://192.168.1.1:8080` to `NewHTTPPoolOpts`
	// which is provided to `pool.Set()` so the pool can identify which of the peers is our instance.
	// The pool will not operate correctly if it can't identify which peer is our instance.

	// Pool keeps track of peers in our cluster and identifies which peer owns a key.
	me := "http://0.0.0.0:5000"
	pool := groupcache.NewHTTPPoolOpts(me, &groupcache.HTTPPoolOptions{})

	knownPeers := false
	if knownPeers {
		// Whenever peers change
		//
		// Add more peers to the cluster You MUST Ensure our instance is included in this list else
		// determining who owns the key accross the cluster will not be consistent, and the pool won't
		// be able to determine if our instance owns the key.
		pool.Set("http://192.168.1.1:8080", "http://192.168.1.2:8080", "http://192.168.1.3:8080")
	}

	server := &http.Server{
		Addr:    "0.0.0.0:5000",
		Handler: pool,
	}
	peer.cacheServer = server
	// Start the HTTP server to listen for peer requests from the groupcache
	go func() {
		log.Println("Cache server started....")
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Create a new group cache with a max cache size of 64Mb
	maxSize := int64(64<<20)
	group := groupcache.NewGroup("cache", maxSize, groupcache.GetterFunc(
		func(ctx context.Context, id string, dest groupcache.Sink) error {
			log.Printf("cache item with KEY=%s not found in local peer", id)
			// TODO compute the requested data and store in cache
			return nil
		},
	))
	peer.cacheGroup = group
}