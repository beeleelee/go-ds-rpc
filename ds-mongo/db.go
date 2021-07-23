package dsmongo

import (
	"context"
	"sync"
	"time"

	log "github.com/ipfs/go-log/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var logging = log.Logger("dsrpc/dsmongo")

const (
	db_name    = "mongo_rpc"
	store_name = "mongo_kv"
)

type Options struct {
	Uri      string
	DBName   string
	CollName string
}

func DefaultOptions() Options {
	return Options{
		Uri:      "mongodb://localhost:27017",
		DBName:   db_name,
		CollName: store_name,
	}
}

type DSMongo struct {
	client *mongo.Client
	opts   Options
	sync.RWMutex
}

func NewDSMongo(opts Options) (*DSMongo, error) {
	defaultOpts := DefaultOptions()
	if opts.Uri == "" {
		opts.Uri = defaultOpts.Uri
	}
	if opts.DBName == "" {
		opts.DBName = defaultOpts.DBName
	}
	if opts.CollName == "" {
		opts.CollName = defaultOpts.CollName
	}
	var err error
	mgoClient, err := mongo.NewClient(options.Client().ApplyURI(opts.Uri))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = mgoClient.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return &DSMongo{
		client: mgoClient,
		opts:   opts,
	}, nil
}

type StoreItem struct {
	ID        string    `bson:"_id" json:"_id"`     // key
	Value     []byte    `bson:"value" json:"value"` // value
	Size      int64     `bson:"size" json:"size"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func (dsm *DSMongo) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return dsm.client.Disconnect(ctx)
}

func (dsm *DSMongo) coll() *mongo.Collection {
	return dsm.client.Database(dsm.opts.DBName).Collection(dsm.opts.CollName)
}

func (dsm *DSMongo) Put(ctx context.Context, item *StoreItem) error {
	coll := dsm.coll()

	dsm.RLock()
	err := coll.FindOne(ctx, bson.D{
		primitive.E{Key: "_id", Value: item.ID},
	}).Err()
	dsm.RUnlock()
	if err == nil { // found same block, do nothing
		logging.Infof("put on existed block: %v, size: %d", item.ID, item.Size)
		return nil
	}
	if err != mongo.ErrNoDocuments {
		return err
	}

	dsm.Lock()
	defer dsm.Unlock()
	item.CreatedAt = time.Now()
	item.UpdatedAt = item.CreatedAt
	r, err := coll.InsertOne(ctx, item)
	if err != nil {
		return err
	}
	logging.Infof("mdb inserted id: %v", r.InsertedID)
	return nil
}

func (dsm *DSMongo) Delete(ctx context.Context, id string) error {
	coll := dsm.coll()

	dsm.Lock()
	defer dsm.Unlock()
	r, err := coll.DeleteOne(ctx, bson.D{
		primitive.E{Key: "_id", Value: id},
	})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return err
	}

	logging.Infof("mdb delete item, id: %d, counnt: %d", id, r.DeletedCount)
	return nil
}

func (dsm *DSMongo) Get(ctx context.Context, id string) ([]byte, error) {
	coll := dsm.coll()

	dsm.RLock()
	defer dsm.RUnlock()
	item := new(StoreItem)
	err := coll.FindOne(ctx, bson.D{
		primitive.E{Key: "_id", Value: id},
	}).Decode(item)

	if err != nil {
		return nil, err
	}

	return item.Value, nil
}

func (dsm *DSMongo) Has(ctx context.Context, id string) (bool, error) {
	coll := dsm.coll()

	dsm.RLock()
	defer dsm.RUnlock()
	err := coll.FindOne(ctx, bson.D{
		primitive.E{Key: "_id", Value: id},
	}).Err()

	if err != nil {
		return false, err
	}

	return true, nil
}

type StoreItemOnlySize struct {
	ID   string `bson:"_id" json:"_id"`
	Size int64  `bson:"size" json:"size"`
}

func (dsm *DSMongo) GetSize(ctx context.Context, id string) (int64, error) {
	coll := dsm.coll()

	dsm.RLock()
	defer dsm.RUnlock()
	item := new(StoreItemOnlySize)
	err := coll.FindOne(ctx, bson.D{
		primitive.E{Key: "_id", Value: id},
	}).Decode(item)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			logging.Warnf("DSMongo GetSize: decode error: %v", err)
		}
		return 0, err
	}

	return item.Size, nil
}
