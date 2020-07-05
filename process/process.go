package process

import (
	"log"
	"runtime"
	"sync"

	"github.com/pkg/errors"
	"go.etcd.io/bbolt"
	"golang.org/x/sync/semaphore"
)

var bucket = []byte("a")
var emptybytes = []byte{}

// Processor provides an abstraction on a key-value persistent database, which
// prevents functions from being executed twice, ever.
type Processor struct {
	db *bbolt.DB
	sm *semaphore.Weighted
	wg sync.WaitGroup
}

func NewProcessor(dbpath string) (*Processor, error) {
	d, err := bbolt.Open(dbpath, 0766, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open database")
	}

	err = d.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucket)
		return err
	})

	if err != nil {
		return nil, errors.Wrap(err, "Failed to create bucket")
	}

	var sema = semaphore.NewWeighted(int64(runtime.GOMAXPROCS(-1) * 2))
	return &Processor{db: d, sm: sema}, nil
}

// DoAsync runs the function asynchronously. The given callback will be passed
// the given key string. This prevents common bugs with range. If the callback
// returns false, then it is not recorded.
func (p *Processor) DoAsync(key string, fn func(key string) bool) {
	p.wg.Add(1)
	go func() {
		p.do(key, fn)
		p.wg.Done()
	}()
}

func (p *Processor) do(key string, fn func(key string) bool) {
	var kb = []byte(key)

	// Noop if operation is already done.
	var done bool
	// We can use a view transaction here as those can be ran concurrently.
	logIfErr(p.db.View(func(tx *bbolt.Tx) error {
		done = tx.Bucket(bucket).Get(kb) != nil
		return nil
	}), "Failed to batch read")

	if done {
		return
	}

	// Run the function. Do not record it if false is returned.
	if !fn(key) {
		return
	}

	// Record that it's done. We should use a batch transaction here.
	logIfErr(p.db.Batch(func(tx *bbolt.Tx) error {
		return tx.Bucket(bucket).Put(kb, emptybytes)
	}), "Failed to put records")
}

func (p *Processor) Finalize() error {
	p.wg.Wait()
	return p.db.Close()
}

var Logger = log.Println

func logIfErr(err error, ctx string) {
	if err != nil {
		Logger(ctx+":", err)
	}
}
