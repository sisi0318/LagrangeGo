package oidb

import (
	"github.com/LagrangeDev/LagrangeGo/client/entity"
	"github.com/LagrangeDev/LagrangeGo/client/packets/pb/service/oidb"
	"github.com/LagrangeDev/LagrangeGo/internal/proto"
)

func BuildFetchUserInfoReq(uid string) (*OidbPacket, error) {
	keys := []uint32{20002, 27394, 20009, 20031, 101, 103, 102, 20022, 20023, 20024, 24002, 27037, 27049, 20011, 20016, 20021, 20003, 20004, 20005, 20006, 20020, 20026, 24007, 104, 105, 42432, 42362, 41756, 41757, 42257, 27372, 42315, 107, 45160, 45161, 27406, 62026, 20037}
	keyList := make([]*oidb.OidbSvcTrpcTcp0XFE1_2Key, len(keys))
	for i, key := range keys {
		keyList[i] = &oidb.OidbSvcTrpcTcp0XFE1_2Key{Key: key}
	}
	body := oidb.OidbSvcTrpcTcp0XFE1_2{
		Uid:    proto.Some(uid),
		Field2: 0,
		Keys:   keyList,
	}
	return BuildOidbPacket(0xFE1, 2, &body, false, false)
}

func ParseFetchUserInfoResp(data []byte) (*entity.Friend, error) {
	var resp oidb.OidbSvcTrpcTcp0XFE1_2Response
	_, err := ParseOidbPacket(data, &resp)
	if err != nil {
		return nil, err
	}
	return &entity.Friend{
		Uin: resp.Body.Uin,
		Uid: resp.Body.Uid,
	}, nil
}
