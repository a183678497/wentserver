package wentdb

import (
	"fmt"
	"wentserver/config"

	"sync"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type DBManager struct {
	db   *leveldb.DB
	lock *sync.RWMutex
}

func newDBManage() *DBManager {
	return &DBManager{db: nil, lock: &sync.RWMutex{}}
}

func (dbm *DBManager) InitDB(path string) error {
	var err1 error
	dbm.db, err1 = leveldb.OpenFile(path, nil)
	if err1 != nil {
		return config.ErrDBInitFailed
	}
	return nil
}

func (dbm *DBManager) CloseDB() {
	if dbm.db == nil {
		return
	}
	dbm.db.Close()
}

func (dbm *DBManager) GetData(key []byte) ([]byte, error) {
	dbm.lock.RLock()
	defer dbm.lock.RUnlock()
	// 读取某条数据
	data, err := dbm.db.Get(key, nil)
	if err != nil {
		return nil, config.ErrDBGetValueFailed
	}
	return data, nil
}

func (dbm *DBManager) PutData(key []byte, value []byte) error {
	dbm.lock.Lock()
	defer dbm.lock.Unlock()
	// 读取某条数据
	err := dbm.db.Put(key, value, nil)
	if err != nil {
		return config.ErrDBPutValueFailed
	}
	return nil
}

func (dbm *DBManager) LoadAccountData() [][]byte {
	dbm.lock.RLock()
	defer dbm.lock.RUnlock()
	iter := dbm.db.NewIterator(util.BytesPrefix([]byte("account_")), nil)
	dataslice := make([][]byte, 0, 2048)
	for iter.Next() {
		fmt.Printf("[%s]:%s\n", iter.Key(), iter.Value())
		dataslice = append(dataslice, iter.Value())
	}
	return dataslice
}

func (dbm *DBManager) LoadPlayerBaseData() [][]byte {
	dbm.lock.RLock()
	defer dbm.lock.RUnlock()
	iter := dbm.db.NewIterator(util.BytesPrefix([]byte("playerbase_")), nil)
	dataslice := make([][]byte, 0, 2048)
	for iter.Next() {
		fmt.Printf("[%s]:%s\n", iter.Key(), iter.Value())
		dataslice = append(dataslice, iter.Value())
	}
	return dataslice
}

var ins *DBManager
var once sync.Once

func GetDBManagerIns() *DBManager {
	once.Do(func() {
		ins = newDBManage()
	})
	return ins
}
