package config

import (
	"errors"
)

const (
	SERVER_IP         = "127.0.0.1"
	SERVER_PORT       = 10006
	SERVER_TYPE       = "tcp"
	SENDCHAN_SIZE     = 1024
	MAXMESSAGE_LEN    = 1024
	WEBSERVER_PORT    = 9998
	DB_PATH           = "./lvdb"
	SAVEROUTINE_COUNT = 2
)

var (
	ErrListenFailed        = errors.New("Listen Failed Error")
	ErrAcceptFailed        = errors.New("Accept Failed Error")
	ErrBuffOverflow        = errors.New("Buff Overflow Error")
	ErrBuffLenLess         = errors.New("Buff Length is not enough")
	ErrParaseMsgHead       = errors.New("Parase Msg Head Failed")
	ErrTypeAssertain       = errors.New("Type Assertain failed")
	ErrMsgLenLarge         = errors.New("Msg Length is too large")
	ErrReadAtLeast         = errors.New("Read at least error!")
	ErrMsgHandlerReg       = errors.New("Msg Handler function not reg")
	ErrParamCallBack       = errors.New("Param is not call back")
	ErrAsyncSendStop       = errors.New("async send chan is stopped")
	ErrSessChanStoped      = errors.New(" session chan is stopped")
	ErrPacketEmpty         = errors.New("Packet is empty")
	ErrWritePacketFailed   = errors.New("Write packet failed")
	ErrConnWriteFailed     = errors.New("Connection Wrtie Failed")
	ErrSignalStopped       = errors.New("Signal Stopped")
	ErrReadPacketFailed    = errors.New("read packet failed!")
	ErrHelloWorldReqFailed = errors.New("Handle Msg Hellow World Req Failed")
	ErrWebSocketRead       = errors.New("Websocket read failed")
	ErrWebSocketClosed     = errors.New("Peer Web Socket closed")
	ErrWebListenFailed     = errors.New("Webserver listen failed!")
	ErrWebSocketDail       = errors.New("Websocket dail failed!")
	ErrWebSocketWrite      = errors.New("Websocket write failed!")
	ErrAsyncSendFailed     = errors.New("Async Send Failed!")
	ErrProtobuffMarshal    = errors.New("Protobuff Marshal failed!")
	ErrProtobuffUnMarshal  = errors.New("Protobuff UnMarshal failed!")
	ErrAccountNameNotExist = errors.New("Account name not exist!")
	ErrAccountMapEmpty     = errors.New("Account Map Empty!")
	ErrDBInitFailed        = errors.New("DB init failed!")
	ErrDBGetValueFailed    = errors.New("DB GET VALUE FAILED")
	ErrDBPutValueFailed    = errors.New("DB PUT VALUE FAILED")
	ErrAccountMgrInit      = errors.New("AccountMgr init failed")
	ErrPlayerMgrInit       = errors.New("PlayerMgr init failed")
	ErrUidUnmarshFailed    = errors.New("Uid Unmarsh Failed!")
	ErrGenuidMgrFailed     = errors.New("Genuid Mgr Failed!")
	ErrGenuidFailed        = errors.New("Gen uid failed !")
	ErrAccountRegFailed    = errors.New("AccountReg Failed")
	ErrAllSaveRoutineExit  = errors.New("All Save Routines exit")
	ErrDBHandlerExit       = errors.New("DBHandler exit success")
	ErrLogInit             = errors.New("Log Init Failed!")
)
