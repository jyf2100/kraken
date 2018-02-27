package base

import "github.com/andres-erbsen/clock"

// FileStore manages files and their metadata. Actual operations are done through FileOp.
type FileStore interface {
	NewFileOp() FileOp
}

// localFileStore manages all agent files on local disk.
// Read/Write operation should access data in this order:
//   map load -> file lock -> verify not deleted -> map load/store -> file/metadata change -> file unlock
// Delete opereration should access data in this order:
//   map load -> file lock -> verify not deleted -> file/metadata change -> delete from map -> file unlock
type localFileStore struct {
	fileEntryFactory FileEntryFactory // Used for dependency injection.
	fileMap          FileMap          // Used for dependency injection.
}

// NewLocalFileStore initializes and returns a new FileStore. It allows dependency injection.
func NewLocalFileStore(clk clock.Clock) (FileStore, error) {
	return &localFileStore{
		fileEntryFactory: DefaultLocalFileEntryFactory(clk),
		fileMap:          NewSimpleFileMap(),
	}, nil
}

// NewCASFileStore initializes and returns a new Content-Addressable FileStore.
// It uses the first few bytes of file digest (which is also used as file name) as shard ID.
// For every byte, one more level of directories will be created.
func NewCASFileStore(clk clock.Clock) (FileStore, error) {
	return &localFileStore{
		fileEntryFactory: DefaultCASFileEntryFactory(clk),
		fileMap:          NewSimpleFileMap(),
	}, nil
}

// NewLRUFileStore initializes and returns a new LRU FileStore.
// When size exceeds limit, the least recently accessed entry will be removed.
func NewLRUFileStore(size int, clk clock.Clock) (FileStore, error) {
	fm, err := NewLRUFileMap(size, clk)
	if err != nil {
		return nil, err
	}
	return &localFileStore{
		fileEntryFactory: DefaultLocalFileEntryFactory(clk),
		fileMap:          fm,
	}, nil
}

// NewCASFileStoreWithLRUMap initializes and returns a new Content-Addressable FileStore.
// It uses the first few bytes of file digest (which is also used as file name) as shard ID.
// For every byte, one more level of directories will be created.
// It also store objects in a LRU FileStore.
// When size exceeds limit, the least recently accessed entry will be removed.
func NewCASFileStoreWithLRUMap(size int, clk clock.Clock) (FileStore, error) {
	fm, err := NewLRUFileMap(size, clk)
	if err != nil {
		return nil, err
	}
	return &localFileStore{
		fileEntryFactory: DefaultCASFileEntryFactory(clk),
		fileMap:          fm,
	}, nil
}

// NewFileOp contructs a new FileOp object.
func (s *localFileStore) NewFileOp() FileOp {
	return NewLocalFileOp(s)
}
