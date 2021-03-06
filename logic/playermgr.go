package logic

import (
	"sync"
	wentproto "wentserver/proto"
	"wentserver/wentdb"

	"github.com/gogo/protobuf/proto"
)

type PlayerManager struct {
	Acnt2Player map[int64]int64
	BaseInfos   map[int64]*wentproto.PlayerBaseInfo
	PlayerInfos map[int64]*wentproto.PlayerInfo
	lock        *sync.RWMutex
}

func newPlayerManager(basedata [][]byte) (*PlayerManager, error) {

	am := new(PlayerManager)
	am.lock = &sync.RWMutex{}
	am.Acnt2Player = make(map[int64]int64)
	am.BaseInfos = make(map[int64]*wentproto.PlayerBaseInfo)
	am.PlayerInfos = make(map[int64]*wentproto.PlayerInfo)
	if basedata == nil || len(basedata) == 0 {
		return am, nil
	}
	for _, value := range basedata {
		playerinfo := &wentproto.PlayerBaseInfo{}
		err := proto.Unmarshal(value, playerinfo)
		if err != nil {
			continue
		}
		am.Acnt2Player[playerinfo.Accountid] = playerinfo.Playeruid
		am.BaseInfos[playerinfo.Playeruid] = playerinfo
	}

	return am, nil
}

var playerins *PlayerManager
var playeronce sync.Once

func GetPlayerManagerIns() *PlayerManager {
	playeronce.Do(func() {
		playerins, err = newPlayerManager(wentdb.GetDBManagerIns().LoadPlayerBaseData())
	})
	return playerins
}
