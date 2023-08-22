package redis

import (
	"crypto/tls"
	"io"

	red "github.com/go-redis/redis/v8"
	"github.com/kong11213613/haremicro/common/util/sync"
)

var clusterManager = sync.NewResourceManager()

func getCluster(r *Redis) (*red.ClusterClient, error) {
	if len(r.options.Addrs) == 0 {
		r.options.Addrs = []string{"127.0.0.1:6379"}
	}
	val, err := clusterManager.GetResource(r.options.Addrs[0], func() (io.Closer, error) {
		var tlsConfig *tls.Config
		if r.options.tls {
			tlsConfig = &tls.Config{
				InsecureSkipVerify: true,
			}
		}
		store := red.NewClusterClient(&red.ClusterOptions{
			Addrs:        r.options.Addrs,
			Password:     r.options.Pass,
			MaxRetries:   maxRetries,
			MinIdleConns: idleConns,
			TLSConfig:    tlsConfig,
		})

		return store, nil
	})
	if err != nil {
		return nil, err
	}

	return val.(*red.ClusterClient), nil
}
