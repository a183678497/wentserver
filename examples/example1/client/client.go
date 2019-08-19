package main

import (
	"fmt"
	"wentserver/config"
	"wentserver/logic"
	"wentserver/netmodel"
	wentproto "wentserver/proto"
	"wentserver/protocol"

	"github.com/gogo/protobuf/proto"
)

func main() {

	cs, err := netmodel.Dial("tcp4", "127.0.0.1:10006")
	if err != nil {
		return
	}
	packet := new(protocol.MsgPacket)
	packet.Head.Id = logic.ACCOUNTINFO_REQ
	csplayerinfo := &wentproto.CSAccountInfo{
		Accountname: "Zack",
	}

	//protobuf编码
	pData, err := proto.Marshal(csplayerinfo)
	if err != nil {
		fmt.Println(config.ErrProtobuffMarshal.Error())
		return
	}
	packet.Head.Len = (uint16)(len(pData))
	packet.Body.Data = pData
	cs.Send(packet)
	packetrsp, err := cs.Recv()
	if err != nil {
		fmt.Println("receive error")
		return
	}

	datarsp := packetrsp.(*protocol.MsgPacket)
	fmt.Println("packet id is", datarsp.Head.Id)
	fmt.Println("packet len is", datarsp.Head.Len)
	scplayerinfo := &wentproto.SCAccountInfo{}
	error2 := proto.Unmarshal(datarsp.Body.Data, scplayerinfo)
	if error2 != nil {
		fmt.Println(config.ErrProtobuffUnMarshal.Error())
		return
	}

	if scplayerinfo.Errid == logic.ERR_NONE {
		fmt.Println("scplayerinfo.Playerinfo.Accountid is ", scplayerinfo.Accountinfo.Accountid)
		fmt.Println("scplayerinfo.Playerinfo.Accountname is ", scplayerinfo.Accountinfo.Accountname)
		return
	}

	if scplayerinfo.Errid != logic.ERR_ACTNOTEXIST {
		fmt.Println("error si  ", scplayerinfo.Errid)
		return
	}

	csregact := &wentproto.CSAccountReg{
		Accountname: "Zack",
	}

	//protobuf编码
	pData, err = proto.Marshal(csregact)
	if err != nil {
		fmt.Println(config.ErrProtobuffMarshal.Error())
		return
	}
	packet.Head.Id = logic.ACCOUNTREG_REQ
	packet.Head.Len = (uint16)(len(pData))
	packet.Body.Data = pData
	cs.Send(packet)

	packetrsp, err = cs.Recv()
	if err != nil {
		fmt.Println("receive error")
		return
	}

	datarsp = packetrsp.(*protocol.MsgPacket)
	fmt.Println("packet id is", datarsp.Head.Id)
	fmt.Println("packet len is", datarsp.Head.Len)

	csregrsp := &wentproto.SCAccountReg{}
	err = proto.Unmarshal(datarsp.Body.Data, csregrsp)
	if err != nil {
		fmt.Println(config.ErrProtobuffUnMarshal.Error())
		return
	}
	fmt.Println("accountid id is ", csregrsp.Accountinfo.Accountid)
	fmt.Println("account name is ", csregrsp.Accountinfo.Accountname)
}
