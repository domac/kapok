package etcd

import (
	etcd "github.com/coreos/etcd/client"
	"golang.org/x/net/context"
	"time"
)

type EtcdDataSource struct {
	keysAPI etcd.KeysAPI
	timeout time.Duration
}

func NewEtcdDataSource(addr string, timeout time.Duration) (*EtcdDataSource, error) {
	cfg := etcd.Config{
		Endpoints:               []string{addr},
		Transport:               etcd.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}
	cli, err := etcd.New(cfg)
	if err != nil {
		return nil, err
	}
	kapi := etcd.NewKeysAPI(cli)
	return &EtcdDataSource{
		keysAPI: kapi,
		timeout: timeout,
	}, nil
}

func (ds *EtcdDataSource) set(keyPath, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), ds.timeout)
	defer cancel()

	_, err := ds.keysAPI.Set(ctx, keyPath, value, nil)
	return err
}

func (ds *EtcdDataSource) get(keyPath string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ds.timeout)
	defer cancel()

	response, err := ds.keysAPI.Get(ctx, keyPath, nil)
	if err != nil {
		return "", err
	}
	return response.Node.Value, nil
}
