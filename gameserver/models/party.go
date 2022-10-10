package models

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/utils"
)

type Party struct {
	members                       []interfaces.CharacterI
	pendingInvitation             bool
	partyLvl                      int32
	disbanding                    bool
	distributionType              interfaces.PartyDistributionTypeInterface
	changeRequestDistributionType interfaces.PartyDistributionTypeInterface
}

func NewParty(leader interfaces.CharacterI, partyDistributionType interfaces.PartyDistributionTypeInterface) interfaces.PartyInterface {
	var party Party
	party.members = append(party.members, leader)
	party.partyLvl = 1 //TODO leader.GetLevel()
	party.distributionType = partyDistributionType
	return &party
}

func (p *Party) GetMemberCount() int {
	return len(p.members)
}

func (p *Party) AddPartyMember(character interfaces.CharacterI) bool {
	if p.IsMemberInParty(character) {
		return false
	}

	if p.changeRequestDistributionType != nil {
		//TODO finishLootRequest(false)
	}

	//TODO проверка на то, что левел нового участника больше текущего уровня пати

	//TODO 	if (isInDimensionalRift()) {
	//			_dr.partyMemberInvited();
	//		}

	//if (isInCommandChannel()) {
	//	player.sendPacket(ExOpenMPCC.STATIC_PACKET);
	//}
	//
	//if (_positionBroadcastTask == null) {
	//	_positionBroadcastTask = ThreadPoolManager.getInstance().scheduleGeneralAtFixedRate(() -> {
	//		if (_positionPacket == null) {
	//			_positionPacket = new PartyMemberPosition(this);
	//		} else {
	//			_positionPacket.reuse(this);
	//		}
	//		broadcastPacket(_positionPacket);
	//	}, PARTY_POSITION_BROADCAST_INTERVAL.toMillis() / 2, PARTY_POSITION_BROADCAST_INTERVAL.toMillis());
	//}

	return true
}

func (p *Party) GetLeaderObjectId() int32 {
	return p.members[0].GetObjectId()
}

func (p *Party) GetDistributionType() interfaces.PartyDistributionTypeInterface {
	return p.distributionType
}

func (p *Party) SetMembers(members []interfaces.CharacterI) {

}

func (p *Party) GetMembers() []interfaces.CharacterI {
	return p.members
}

func (p *Party) GetLeader() interfaces.CharacterI {
	return p.members[0]
}

func (p *Party) IsMemberInParty(character interfaces.CharacterI) bool {
	for _, member := range p.members {
		if member.GetObjectId() == character.GetObjectId() {
			return true
		}
	}
	return false
}

func (p *Party) IsLeader(character interfaces.CharacterI) bool {
	return character.GetObjectId() == p.members[0].GetObjectId()
}

func (p *Party) GetMemberIndex(character interfaces.CharacterI) int {
	for i := range p.members {
		if character.GetObjectId() == p.members[i].GetObjectId() {
			return i
		}
	}
	return 0
}

func (p *Party) BroadcastParty(msg []byte) {
	pb := utils.GetPacketByte()
	pb.SetData(msg)

	for _, member := range p.members {
		if member != nil {
			member.EncryptAndSend(pb.GetData())
		}
	}
	pb.Release()
}

func (p *Party) IsDisbanding() bool {
	return p.disbanding
}

func (p *Party) SetDisbanding(flag bool) {
	p.disbanding = flag
}

func (p *Party) AddMember(member interfaces.CharacterI) {
	p.members = append(p.members, member)
}

func (p *Party) RemoveMember(member interfaces.CharacterI) {
	if !p.IsMemberInParty(member) {
		return
	}
	index := p.GetMemberIndex(member)
	p.members = append(p.members[:index], p.members[index+1:]...)
}
