package mongods

import (
	"fmt"

	"github.com/ipfs/go-ipfs/plugin"
	"github.com/ipfs/go-ipfs/repo"
	"github.com/ipfs/go-ipfs/repo/fsrepo"

	dsrpc "github.com/beeleelee/go-ds-rpc"
	dsmongo "github.com/beeleelee/go-ds-rpc/ds-mongo"
)

// Plugins is exported list of plugins that will be loaded
var Plugins = []plugin.Plugin{
	&mongodsPlugin{},
}

type mongodsPlugin struct{}

var _ plugin.PluginDatastore = (*mongodsPlugin)(nil)

func (*mongodsPlugin) Name() string {
	return "ds-mongods"
}

func (*mongodsPlugin) Version() string {
	return "0.0.1"
}

func (*mongodsPlugin) Init(_ *plugin.Environment) error {
	return nil
}

func (*mongodsPlugin) DatastoreTypeName() string {
	return "mongods"
}

type datastoreConfig struct {
	uri    string
	prefix string
}

func (*mongodsPlugin) DatastoreConfigParser() fsrepo.ConfigFromMap {
	return func(params map[string]interface{}) (fsrepo.DatastoreConfig, error) {
		var c datastoreConfig
		var ok bool

		c.uri, ok = params["uri"].(string)
		if !ok {
			return nil, fmt.Errorf("'uri' field is missing or not string")
		}
		c.prefix, ok = params["prefix"].(string)
		if !ok {
			return nil, fmt.Errorf("'prefix' field is missing or not string")
		}
		return &c, nil
	}
}

func (c *datastoreConfig) DiskSpec() fsrepo.DiskSpec {
	return map[string]interface{}{
		"type":   "mongods",
		"uri":    c.uri,
		"prefix": c.prefix,
	}
}

func (c *datastoreConfig) Create(path string) (repo.Datastore, error) {
	client, err := dsmongo.NewMongoStoreClient(c.uri)
	if err != nil {
		return nil, err
	}
	return dsrpc.NewDataStore(c.prefix, client)
}
