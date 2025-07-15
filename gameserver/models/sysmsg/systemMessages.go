package sysmsg

type SysMsg struct {
	Id     int32
	Params []Params
}
type Params struct {
	tType ParamType
	value interface{}
}
type ParamType = byte

const (
	TypeSystemString ParamType = 13
	TypePlayerName   ParamType = 12
	TypeDoorName     ParamType = 11
	TypeInstanceName ParamType = 10
	TypeElementName  ParamType = 9
	// id 8 - same as 3
	TypeZoneName   ParamType = 7
	TypeLongNumber ParamType = 6
	TypeCastleName ParamType = 5
	TypeSkillName  ParamType = 4
	TypeItemName   ParamType = 3
	TypeNpcName    ParamType = 2
	TypeIntNumber  ParamType = 1
	TypeText       ParamType = 0
)

func (sys *Params) GetType() ParamType {
	return sys.tType
}

// AddCastleId
// Appends a Castle name parameter type, the name will be read from CastleName-e.dat.<br>
// <ul>
// <li>1-9 Castle names</li>
// <li>21 Fortress of Resistance</li>
// <li>22-33 Clan Hall names</li>
// <li>34 Devastated Castle</li>
// <li>35 Bandit Stronghold</li>
// <li>36-61 Clan Hall names</li>
// <li>62 Rainbow Springs</li>
// <li>63 Wild Beast Reserve</li>
// <li>64 Fortress of the Dead</li>
// <li>81-89 Territory names</li>
// <li>90-100 null</li>
// <li>101-121 Fortress names</li>
func (sys *SysMsg) AddCastleId(number int32) {
	sys.Params = append(sys.Params, Params{tType: TypeCastleName, value: number})
}

func (sys *SysMsg) AddInt(number int32) {
	sys.Params = append(sys.Params, Params{tType: TypeIntNumber, value: number})
}

func (sys *SysMsg) AddLong(number int64) {
	sys.Params = append(sys.Params, Params{tType: TypeLongNumber, value: number})
}

func (sys *SysMsg) AddString(str string) {
	sys.Params = append(sys.Params, Params{tType: TypeText, value: str})
}
func (sys *SysMsg) AddItemName(id int32) {
	sys.Params = append(sys.Params, Params{tType: TypeItemName, value: id})
}

func (sys *SysMsg) AddZone(x, y, z int32) {
	sys.Params = append(sys.Params, Params{tType: TypeZoneName, value: [3]int32{x, y, z}})
}
func (sys *SysMsg) AddCharacterName(name string) {
	sys.Params = append(sys.Params, Params{tType: TypePlayerName, value: name})
}
func (sys *Params) GetValueString() string {
	return sys.value.(string)
}
func (sys *Params) GetValueInt64() int64 {
	return sys.value.(int64)
}
func (sys *Params) GetValueInt32() int32 {
	return sys.value.(int32)
}
func (sys *Params) GetTwoElementSlice() [2]int32 {
	return sys.value.([2]int32)
}
func (sys *Params) GetThreeElementSlice() [3]int32 {
	return sys.value.([3]int32)
}

var (
	UserReportedAndCannotJoinParty = SysMsg{Id: 2482}

	YouCanReportInS1MinutesS2ReportPointsRemainInAccount = SysMsg{Id: 2774}

	CannotReportInWarzonePeacezoneClanwarOlympiad = SysMsg{Id: 2470}

	CannotReportAlredyReportedFromYourClanOrIp = SysMsg{Id: 2471}

	CannotReportAlaredyReportedFromSameAccount = SysMsg{Id: 2472}

	Reported10MinsWithoutChat = SysMsg{Id: 2473}

	Reported60MinsWithoutJoinParty = SysMsg{Id: 2474}

	Reported120MinsWithoutJoinParty = SysMsg{Id: 2475}

	Reported180MinsWithoutJoinParty = SysMsg{Id: 2476}

	Reported120MinsWithoutActions = SysMsg{Id: 2477}

	Reported180MinsWithoutActions = SysMsg{Id: 2478}

	Reported120MinsWithoutMove = SysMsg{Id: 2480}

	CannotReportTargetInClanWar = SysMsg{Id: 2378}

	CannotReportCharacterWithoutGainexp = SysMsg{Id: 2379}

	C1ReportedAsBot = SysMsg{Id: 2371}

	TheAngelNevitHasBlessedYouFromAbove = SysMsg{Id: 3266}

	YouHaveAcquiredS1PcCafePoints = SysMsg{Id: 2393}

	TheMaxmimumAccumulationAllowedOfPcCafePointsHasBeenExceeded = SysMsg{Id: 2389}

	FromNowOnAngelNevitAbideWithYou = SysMsg{Id: 3266}

	YouAreStartingToFeelTheEffectsOfNevitsAdventBlessing = SysMsg{Id: 3267}

	YouAreFurtherInfusedWithTheBlessingsOfNevit = SysMsg{Id: 3268}

	NevitsAdventBlessingShinesStronglyFromAbove = SysMsg{Id: 3269}

	NevitsAdventBlessingHasEnded = SysMsg{Id: 3275}

	InstantZoneRestricted = SysMsg{Id: 6507}

	// You have been disconnected from the server.
	YouHaveBeenDisconnected = SysMsg{Id: 0}

	// The server will be coming down in $1 seconds. Please find a safe place to log out.
	TheServerWillBeComingDownInS1Seconds = SysMsg{Id: 1}

	// {{S1}} does not exist.
	S1DoesNotExist = SysMsg{Id: 2}

	// {{S1}} is not currently logged in.
	S1IsNotOnline = SysMsg{Id: 3}

	// You cannot ask yourself to apply to a clan.
	CannotInviteYourself = SysMsg{Id: 4}

	// {{S1}} already exists.
	S1AlreadyExists = SysMsg{Id: 5}

	// {{S1}} does not exist
	S1DoesNotExist2 = SysMsg{Id: 6}

	// You are already a member of {{S1}}.
	AlreadyMemberOfS1 = SysMsg{Id: 7}

	// You are working with another clan.
	YouAreWorkingWithAnotherClan = SysMsg{Id: 8}

	// {{S1}} is not a clan leader.
	S1IsNotAClanLeader = SysMsg{Id: 9}

	// {{S1}} is working with another clan.
	S1WorkingWithAnotherClan = SysMsg{Id: 10}

	// There are no applicants for this clan.
	NoApplicantsForThisClan = SysMsg{Id: 11}

	// The applicant information is incorrect.
	ApplicantInformationIncorrect = SysMsg{Id: 12}

	// Unable to disperse: your clan has requested to participate in a castle siege.
	CannotDissolveCauseClanWillParticipateInCastleSiege = SysMsg{Id: 13}

	// Unable to disperse: your clan owns one or more castles or hideouts.
	CannotDissolveCauseClanOwnsCastlesHideouts = SysMsg{Id: 14}

	// You are in siege.
	YouAreInSiege = SysMsg{Id: 15}

	// You are not in siege.
	YouAreNotInSiege = SysMsg{Id: 16}

	// The castle siege has begun.
	CastleSiegeHasBegun = SysMsg{Id: 17}

	// The castle siege has ended.
	CastleSiegeHasEnded = SysMsg{Id: 18}

	// There is a new Lord of the castle!
	NewCastleLord = SysMsg{Id: 19}

	// The gate is being opened.
	GateIsOpening = SysMsg{Id: 20}

	// The gate is being destroyed.
	GateIsDestroyed = SysMsg{Id: 21}

	// Your target is out of range.
	TargetTooFar = SysMsg{Id: 22}

	// Not enough HP.
	NotEnoughHp = SysMsg{Id: 23}

	// Not enough MP.
	NotEnoughMp = SysMsg{Id: 24}

	// Rejuvenating HP.
	RejuvenatingHp = SysMsg{Id: 25}

	// Rejuvenating MP.
	RejuvenatingMp = SysMsg{Id: 26}

	// Your casting has been interrupted.
	CastingInterrupted = SysMsg{Id: 27}

	// You have obtained {{S1}} adena.
	YouPickedUpS1Adena = SysMsg{Id: 28}

	// You have obtained {{S2}} {{S1}}.
	YouPickedUpS1S2 = SysMsg{Id: 29}

	// You have obtained {{S1}}.
	YouPickedUpS1 = SysMsg{Id: 30}

	// You cannot move while sitting.
	CantMoveSitting = SysMsg{Id: 31}

	// You are unable to engage in combat. Please go to the nearest restart point.
	UnableCombatPleaseGoRestart = SysMsg{Id: 32}

	// You cannot move while casting.
	CantMoveCasting = SysMsg{Id: 33}

	// Welcome to the World of Lineage II.
	WelcomeToLineage = SysMsg{Id: 34}

	// You hit for {{S1}} damage
	YouDidS1Dmg = SysMsg{Id: 35}

	// {{C1}} hit you for {{S2}} damage.
	C1GaveYouS2Dmg = SysMsg{Id: 36}

	// {{C1}} hit you for {{S2}} damage.
	C1GaveYouS2Dmg2 = SysMsg{Id: 37}

	// You carefully nock an arrow.
	GettingReadyToShootAnArrow = SysMsg{Id: 41}

	// You have avoided {{C1}}'s attack.
	AvoidedC1Attack = SysMsg{Id: 42}

	// You have missed.
	MissedTarget = SysMsg{Id: 43}

	// Critical hit!
	CriticalHit = SysMsg{Id: 44}

	// You have earned {{S1}} experience.
	EarnedS1Experience = SysMsg{Id: 45}

	// You use {{S1}}.
	UseS1 = SysMsg{Id: 46}

	// You begin to use a(n) {{S1}}.
	BeginToUseS1 = SysMsg{Id: 47}

	// {{S1}} is not available at this time: being prepared for reuse.
	S1PreparedForReuse = SysMsg{Id: 48}

	// You have equipped your {{S1}}.
	S1Equipped = SysMsg{Id: 49}

	// Your target cannot be found.
	TargetCantFound = SysMsg{Id: 50}

	// You cannot use this on yourself.
	CannotUseOnYourself = SysMsg{Id: 51}

	// You have earned {{S1}} adena.
	EarnedS1Adena = SysMsg{Id: 52}

	// You have earned {{S2}} {{S1}}(s).
	EarnedS2S1S = SysMsg{Id: 53}

	// You have earned {{S1}}.
	EarnedItemS1 = SysMsg{Id: 54}

	// You have failed to pick up {{S1}} adena.
	FailedToPickupS1Adena = SysMsg{Id: 55}

	// You have failed to pick up {{S1}}.
	FailedToPickupS1 = SysMsg{Id: 56}

	// You have failed to pick up {{S2}} {{S1}}(s).
	FailedToPickupS2S1S = SysMsg{Id: 57}

	// You have failed to earn {{S1}} adena.
	FailedToEarnS1Adena = SysMsg{Id: 58}

	// You have failed to earn {{S1}}.
	FailedToEarnS1 = SysMsg{Id: 59}

	// You have failed to earn {{S2}} {{S1}}(s).
	FailedToEarnS2S1S = SysMsg{Id: 60}

	// Nothing happened.
	NothingHappened = SysMsg{Id: 61}

	// Your {{S1}} has been successfully enchanted.
	S1SuccessfullyEnchanted = SysMsg{Id: 62}

	// Your +{{S1}} {{S2}} has been successfully enchanted.
	S1S2SuccessfullyEnchanted = SysMsg{Id: 63}

	// The enchantment has failed! Your {{S1}} has been crystallized.
	EnchantmentFailedS1Evaporated = SysMsg{Id: 64}

	// The enchantment has failed! Your +{{S1}} {{S2}} has been crystallized.
	EnchantmentFailedS1S2Evaporated = SysMsg{Id: 65}

	// {{C1}} is inviting you to join a party. Do you accept?
	C1InvitedYouToParty = SysMsg{Id: 66}

	// {{S1}} has invited you to the join the clan, {{S2}}. Do you wish to join?
	S1HasInvitedYouToJoinTheClanS2 = SysMsg{Id: 67}

	// Would you like to withdraw from the {{S1}} clan? If you leave, you will have to wait at least a day before joining another clan.
	WouldYouLikeToWithdrawFromTheS1Clan = SysMsg{Id: 68}

	// Would you like to dismiss {{S1}} from the clan? If you do so, you will have to wait at least a day before accepting a new member.
	WouldYouLikeToDismissS1FromTheClan = SysMsg{Id: 69}

	// Do you wish to disperse the clan, {{S1}}?
	DoYouWishToDisperseTheClanS1 = SysMsg{Id: 70}

	// How many of your {{S1}}(s) do you wish to discard?
	HowManyS1Discard = SysMsg{Id: 71}

	// How many of your {{S1}}(s) do you wish to move?
	HowManyS1Move = SysMsg{Id: 72}

	// How many of your {{S1}}(s) do you wish to destroy?
	HowManyS1Destroy = SysMsg{Id: 73}

	// Do you wish to destroy your {{S1}}?
	WishDestroyS1 = SysMsg{Id: 74}

	// ID does not exist.
	IdNotExist = SysMsg{Id: 75}

	// Incorrect password.
	IncorrectPassword = SysMsg{Id: 76}

	// You cannot create another character. Please delete the existing character and try again.
	CannotCreateCharacter = SysMsg{Id: 77}

	// When you delete a character, any items in his/her possession will also be deleted. Do you really wish to delete {{S1}}%?
	WishDeleteS1 = SysMsg{Id: 78}

	// This name already exists.
	NamingNameAlreadyExists = SysMsg{Id: 79}

	// Names must be between 1-16 characters, excluding spaces or special characters.
	NamingCharnameUpTo16Chars = SysMsg{Id: 80}

	// Please select your race.
	PleaseSelectRace = SysMsg{Id: 81}

	// Please select your occupation.
	PleaseSelectOccupation = SysMsg{Id: 82}

	// Please select your gender.
	PleaseSelectGender = SysMsg{Id: 83}

	// You may not attack in a peaceful zone.
	CantAtkPeacezone = SysMsg{Id: 84}

	// You may not attack this target in a peaceful zone.
	TargetInPeacezone = SysMsg{Id: 85}

	// Please enter your ID.
	PleaseEnterId = SysMsg{Id: 86}

	// Please enter your password.
	PleaseEnterPassword = SysMsg{Id: 87}

	// Your protocol version is different, please restart your client and run a full check.
	WrongProtocolCheck = SysMsg{Id: 88}

	// Your protocol version is different, please continue.
	WrongProtocolContinue = SysMsg{Id: 89}

	// You are unable to connect to the server.
	UnableToConnect = SysMsg{Id: 90}

	// Please select your hairstyle.
	PleaseSelectHairstyle = SysMsg{Id: 91}

	// {{S1}} has worn off.
	S1HasWornOff = SysMsg{Id: 92}

	// You do not have enough SP for this.
	NotEnoughSp = SysMsg{Id: 93}

	// 2004-2011 (c) NC Interactive, Inc. All Rights Reserved.
	Copyright = SysMsg{Id: 94}

	// You have earned {{S1}} experience and {{S2}} SP.
	YouEarnedS1ExpAndS2Sp = SysMsg{Id: 95}

	// Your level has increased!
	YouIncreasedYourLevel = SysMsg{Id: 96}

	// This item cannot be moved.
	CannotMoveThisItem = SysMsg{Id: 97}

	// This item cannot be discarded.
	CannotDiscardThisItem = SysMsg{Id: 98}

	// This item cannot be traded or sold.
	CannotTradeThisItem = SysMsg{Id: 99}

	// {{C1}} is requesting to trade. Do you wish to continue?
	C1RequestsTrade = SysMsg{Id: 100}

	// You cannot exit while in combat.
	CantLogoutWhileFighting = SysMsg{Id: 101}

	// You cannot restart while in combat.
	CantRestartWhileFighting = SysMsg{Id: 102}

	// This ID is currently logged in.
	IdLoggedIn = SysMsg{Id: 103}

	// You cannot change weapons during an attack.
	CannotChangeWeaponDuringAnAttack = SysMsg{Id: 104}

	// {{C1}} has been invited to the party.
	C1InvitedToParty = SysMsg{Id: 105}

	// You have joined {{S1}}'s party.
	YouJoinedS1Party = SysMsg{Id: 106}

	// {{C1}} has joined the party.
	C1JoinedParty = SysMsg{Id: 107}

	// {{C1}} has left the party.
	C1LeftParty = SysMsg{Id: 108}

	// Invalid target.
	IncorrectTarget = SysMsg{Id: 109}

	// {{S1}} {{S2}}'s effect can be felt.
	YouFeelS1Effect = SysMsg{Id: 110}

	// Your shield defense has succeeded.
	ShieldDefenceSuccessfull = SysMsg{Id: 111}

	// You have run out of arrows.
	NotEnoughArrows = SysMsg{Id: 112}

	// {{S1}} cannot be used due to unsuitable terms.
	S1CannotBeUsed = SysMsg{Id: 113}

	// You have entered the shadow of the Mother Tree.
	EnterShadowMotherTree = SysMsg{Id: 114}

	// You have left the shadow of the Mother Tree.
	ExitShadowMotherTree = SysMsg{Id: 115}

	// You have entered a peaceful zone.
	EnterPeacefulZone = SysMsg{Id: 116}

	// You have left the peaceful zone.
	ExitPeacefulZone = SysMsg{Id: 117}

	// You have requested a trade with {{C1}}.
	RequestC1ForTrade = SysMsg{Id: 118}

	// {{C1}} has denied your request to trade.
	C1DeniedTradeRequest = SysMsg{Id: 119}

	// You begin trading with {{C1}}.
	BeginTradeWithC1 = SysMsg{Id: 120}

	// {{C1}} has confirmed the trade.
	C1ConfirmedTrade = SysMsg{Id: 121}

	// You may no longer adjust items in the trade because the trade has been confirmed.
	CannotAdjustItemsAfterTradeConfirmed = SysMsg{Id: 122}

	// Your trade is successful.
	TradeSuccessful = SysMsg{Id: 123}

	// {{C1}} has cancelled the trade.
	C1CanceledTrade = SysMsg{Id: 124}

	// Do you wish to exit the game?
	WishExitGame = SysMsg{Id: 125}

	// Do you wish to return to the character select screen?
	WishRestartGame = SysMsg{Id: 126}

	// You have been disconnected from the server. Please login again.
	DisconnectedFromServer = SysMsg{Id: 127}

	// Your character creation has failed.
	CharacterCreationFailed = SysMsg{Id: 128}

	// Your inventory is full.
	SlotsFull = SysMsg{Id: 129}

	// Your warehouse is full.
	WarehouseFull = SysMsg{Id: 130}

	// {{S1}} has logged in.
	S1LoggedIn = SysMsg{Id: 131}

	// {{S1}} has been added to your friends list.
	S1AddedToFriends = SysMsg{Id: 132}

	// {{S1}} has been removed from your friends list.
	S1RemovedFromYourFriendsList = SysMsg{Id: 133}

	// Please check your friends list again.
	PleaceCheckYourFriendListAgain = SysMsg{Id: 134}

	// {{C1}} did not reply to your invitation. Your invitation has been cancelled.
	C1DidNotReplyToYourInvite = SysMsg{Id: 135}

	// You have not replied to {{C1}}'s invitation. The offer has been cancelled.
	YouDidNotReplyToC1Invite = SysMsg{Id: 136}

	// There are no more items in the shortcut.
	NoMoreItemsShortcut = SysMsg{Id: 137}

	// Designate shortcut.
	DesignateShortcut = SysMsg{Id: 138}

	// {{C1}} has resisted your {{S2}}.
	C1ResistedYourS2 = SysMsg{Id: 139}

	// Your skill was removed due to a lack of MP.
	SkillRemovedDueLackMp = SysMsg{Id: 140}

	// Once the trade is confirmed, the item cannot be moved again.
	OnceTheTradeIsConfirmedTheItemCannotBeMovedAgain = SysMsg{Id: 141}

	// You are already trading with someone.
	AlreadyTrading = SysMsg{Id: 142}

	// {{C1}} is already trading with another person. Please try again later.
	C1AlreadyTrading = SysMsg{Id: 143}

	// That is the incorrect target.
	TargetIsIncorrect = SysMsg{Id: 144}

	// That player is not online.
	TargetIsNotFoundInTheGame = SysMsg{Id: 145}

	// Chatting is now permitted.
	ChattingPermitted = SysMsg{Id: 146}

	// Chatting is currently prohibited.
	ChattingProhibited = SysMsg{Id: 147}

	// You cannot use quest items.
	CannotUseQuestItems = SysMsg{Id: 148}

	// You cannot pick up or use items while trading.
	CannotUseItemWhileTrading = SysMsg{Id: 149}

	// You cannot discard or destroy an item while trading at a private store.
	CannotDiscardOrDestroyItemWhileTrading = SysMsg{Id: 150}

	// That is too far from you to discard.
	CannotDiscardDistanceTooFar = SysMsg{Id: 151}

	// You have invited the wrong target.
	YouHaveInvitedTheWrongTarget = SysMsg{Id: 152}

	// {{C1}} is on another task. Please try again later.
	C1IsBusyTryLater = SysMsg{Id: 153}

	// Only the leader can give out invitations.
	OnlyLeaderCanInvite = SysMsg{Id: 154}

	// The party is full.
	PartyFull = SysMsg{Id: 155}

	// Drain was only 50 percent successful.
	DrainHalfSuccesful = SysMsg{Id: 156}

	// You resisted {{C1}}'s drain.
	ResistedC1Drain = SysMsg{Id: 157}

	// Your attack has failed.
	AttackFailed = SysMsg{Id: 158}

	// You resisted {{C1}}'s magic.
	ResistedC1Magic = SysMsg{Id: 159}

	// {{C1}} is a member of another party and cannot be invited.
	C1IsAlreadyInParty = SysMsg{Id: 160}

	// That player is not currently online.
	InvitedUserNotOnline = SysMsg{Id: 161}

	// Warehouse is too far.
	WarehouseTooFar = SysMsg{Id: 162}

	// You cannot destroy it because the number is incorrect.
	CannotDestroyNumberIncorrect = SysMsg{Id: 163}

	// Waiting for another reply.
	WaitingForAnotherReply = SysMsg{Id: 164}

	// You cannot add yourself to your own friend list.
	YouCannotAddYourselfToOwnFriendList = SysMsg{Id: 165}

	// Friend list is not ready yet. Please register again later.
	FriendListNotReadyYetRegisterLater = SysMsg{Id: 166}

	// {{C1}} is already on your friend list.
	C1AlreadyOnFriendList = SysMsg{Id: 167}

	// {{C1}} has sent a friend request.
	C1RequestedToBecomeFriends = SysMsg{Id: 168}

	// Accept friendship 0/1 (1 to accept, 0 to deny)
	AcceptTheFriendship = SysMsg{Id: 169}

	// The user who requested to become friends is not found in the game.
	TheUserYouRequestedIsNotInGame = SysMsg{Id: 170}

	// {{C1}} is not on your friend list.
	C1NotOnYourFriendsList = SysMsg{Id: 171}

	// You lack the funds needed to pay for this transaction.
	LackFundsForTransaction1 = SysMsg{Id: 172}

	// You lack the funds needed to pay for this transaction.
	LackFundsForTransaction2 = SysMsg{Id: 173}

	// That person's inventory is full.
	OtherInventoryFull = SysMsg{Id: 174}

	// That skill has been de-activated as HP was fully recovered.
	SkillDeactivatedHpFull = SysMsg{Id: 175}

	// That person is in message refusal mode.
	ThePersonIsInMessageRefusalMode = SysMsg{Id: 176}

	// Message refusal mode.
	MessageRefusalMode = SysMsg{Id: 177}

	// Message acceptance mode.
	MessageAcceptanceMode = SysMsg{Id: 178}

	// You cannot discard those items here.
	CantDiscardHere = SysMsg{Id: 179}

	// You have {{S1}} day(s) left until deletion. Do you wish to cancel this action?
	S1DaysLeftCancelAction = SysMsg{Id: 180}

	// Cannot see target.
	CantSeeTarget = SysMsg{Id: 181}

	// Do you want to quit the current quest?
	WantQuitCurrentQuest = SysMsg{Id: 182}

	// There are too many users on the server. Please try again later
	TooManyUsers = SysMsg{Id: 183}

	// Please try again later.
	TryAgainLater = SysMsg{Id: 184}

	// You must first select a user to invite to your party.
	FirstSelectUserToInviteToParty = SysMsg{Id: 185}

	// You must first select a user to invite to your clan.
	FirstSelectUserToInviteToClan = SysMsg{Id: 186}

	// Select user to expel.
	SelectUserToExpel = SysMsg{Id: 187}

	// Please create your clan name.
	PleaseCreateClanName = SysMsg{Id: 188}

	// Your clan has been created.
	ClanCreated = SysMsg{Id: 189}

	// You have failed to create a clan.
	FailedToCreateClan = SysMsg{Id: 190}

	// Clan member {{S1}} has been expelled.
	ClanMemberS1Expelled = SysMsg{Id: 191}

	// You have failed to expel {{S1}} from the clan.
	FailedExpelS1 = SysMsg{Id: 192}

	// Clan has dispersed.
	ClanHasDispersed = SysMsg{Id: 193}

	// You have failed to disperse the clan.
	FailedToDisperseClan = SysMsg{Id: 194}

	// Entered the clan.
	EnteredTheClan = SysMsg{Id: 195}

	// {{S1}} declined your clan invitation.
	S1RefusedToJoinClan = SysMsg{Id: 196}

	// You have withdrawn from the clan.
	YouHaveWithdrawnFromClan = SysMsg{Id: 197}

	// You have failed to withdraw from the {{S1}} clan.
	FailedToWithdrawFromS1Clan = SysMsg{Id: 198}

	// You have recently been dismissed from a clan. You are not allowed to join another clan for 24-hours.
	ClanMembershipTerminated = SysMsg{Id: 199}

	// You have withdrawn from the party.
	YouLeftParty = SysMsg{Id: 200}

	// {{C1}} was expelled from the party.
	C1WasExpelledFromParty = SysMsg{Id: 201}

	// You have been expelled from the party.
	HaveBeenExpelledFromParty = SysMsg{Id: 202}

	// The party has dispersed.
	PartyDispersed = SysMsg{Id: 203}

	// Incorrect name. Please try again.
	IncorrectNameTryAgain = SysMsg{Id: 204}

	// Incorrect character name. Please try again.
	IncorrectCharacterNameTryAgain = SysMsg{Id: 205}

	// Please enter the name of the clan you wish to declare war on.
	EnterClanNameToDeclareWar = SysMsg{Id: 206}

	// {{S2}} of the clan {{S1}} requests declaration of war. Do you accept?
	S2OfTheClanS1RequestsWar = SysMsg{Id: 207}

	// You are not a clan member and cannot perform this action.
	YouAreNotAClanMember = SysMsg{Id: 212}

	// Not working. Please try again later.
	NotWorkingPleaseTryAgainLater = SysMsg{Id: 213}

	// Your title has been changed.
	TitleChanged = SysMsg{Id: 214}

	// War with the {{S1}} clan has begun.
	WarWithTheS1ClanHasBegun = SysMsg{Id: 215}

	// War with the {{S1}} clan has ended.
	WarWithTheS1ClanHasEnded = SysMsg{Id: 216}

	// You have won the war over the {{S1}} clan!
	YouHaveWonTheWarOverTheS1Clan = SysMsg{Id: 217}

	// You have surrendered to the {{S1}} clan.
	YouHaveSurrenderedToTheS1Clan = SysMsg{Id: 218}

	// Your clan leader has died. You have been defeated by the {{S1}} clan.
	YouWereDefeatedByS1Clan = SysMsg{Id: 219}

	// You have {{S1}} minutes left until the clan war ends.
	S1MinutesLeftUntilClanWarEnds = SysMsg{Id: 220}

	// The time limit for the clan war is up. War with the {{S1}} clan is over.
	ClanWarWithS1ClanHasEnded = SysMsg{Id: 221}

	// {{S1}} has joined the clan.
	S1HasJoinedClan = SysMsg{Id: 222}

	// {{S1}} has withdrawn from the clan.
	S1HasWithdrawnFromTheClan = SysMsg{Id: 223}

	// {{S1}} did not respond: Invitation to the clan has been cancelled.
	S1DidNotRespondToClanInvitation = SysMsg{Id: 224}

	// You didn't respond to {{S1}}'s invitation: joining has been cancelled.
	YouDidNotRespondToS1ClanInvitation = SysMsg{Id: 225}

	// The {{S1}} clan did not respond: war proclamation has been refused.
	S1ClanDidNotRespond = SysMsg{Id: 226}

	// Clan war has been refused because you did not respond to {{S1}} clan's war proclamation.
	ClanWarRefusedYouDidNotRespondToS1 = SysMsg{Id: 227}

	// Request to end war has been denied.
	RequestToEndWarHasBeenDenied = SysMsg{Id: 228}

	// You do not meet the criteria in order to create a clan.
	YouDoNotMeetCriteriaInOrderToCreateAClan = SysMsg{Id: 229}

	// You must wait 10 days before creating a new clan.
	YouMustWaitXxDaysBeforeCreatingANewClan = SysMsg{Id: 230}

	// After a clan member is dismissed from a clan, the clan must wait at least a day before accepting a new member.
	YouMustWaitBeforeAcceptingANewMember = SysMsg{Id: 231}

	// After leaving or having been dismissed from a clan, you must wait at least a day before joining another clan.
	YouMustWaitBeforeJoiningAnotherClan = SysMsg{Id: 232}

	// The Academy/Royal Guard/Order of Knights is full and cannot accept new members at this time.
	SubclanIsFull = SysMsg{Id: 233}

	// The target must be a clan member.
	TargetMustBeInClan = SysMsg{Id: 234}

	// You are not authorized to bestow these rights.
	NotAuthorizedToBestowRights = SysMsg{Id: 235}

	// Only the clan leader is enabled.
	OnlyTheClanLeaderIsEnabled = SysMsg{Id: 236}

	// The clan leader could not be found.
	ClanLeaderNotFound = SysMsg{Id: 237}

	// Not joined in any clan.
	NotJoinedInAnyClan = SysMsg{Id: 238}

	// The clan leader cannot withdraw.
	ClanLeaderCannotWithdraw = SysMsg{Id: 239}

	// Currently involved in clan war.
	CurrentlyInvolvedInClanWar = SysMsg{Id: 240}

	// Leader of the {{S1}} Clan is not logged in.
	LeaderOfS1ClanNotFound = SysMsg{Id: 241}

	// Select target.
	SelectTarget = SysMsg{Id: 242}

	// You cannot declare war on an allied clan.
	CannotDeclareWarOnAlliedClan = SysMsg{Id: 243}

	// You are not allowed to issue this challenge.
	NotAllowedToChallenge = SysMsg{Id: 244}

	// 5 days has not passed since you were refused war. Do you wish to continue?
	FiveDaysNotPassedSinceRefusedWar = SysMsg{Id: 245}

	// That clan is currently at war.
	ClanCurrentlyAtWar = SysMsg{Id: 246}

	// You have already been at war with the {{S1}} clan: 5 days must pass before you can challenge this clan again
	FiveDaysMustPassBeforeChallengeAgain = SysMsg{Id: 247}

	// You cannot proclaim war: the {{S1}} clan does not have enough members.
	S1ClanNotEnoughMembersForWar = SysMsg{Id: 248}

	// Do you wish to surrender to the {{S1}} clan?
	WishSurrenderToS1Clan = SysMsg{Id: 249}

	// You have personally surrendered to the {{S1}} clan. You are no longer participating in this clan war.
	YouHavePersonallySurrenderedToTheS1Clan = SysMsg{Id: 250}

	// You cannot proclaim war: you are at war with another clan.
	AlreadyAtWarWithAnotherClan = SysMsg{Id: 251}

	// Enter the clan name to surrender to.
	EnterClanNameToSurrenderTo = SysMsg{Id: 252}

	// Enter the name of the clan you wish to end the war with.
	EnterClanNameToEndWar = SysMsg{Id: 253}

	// A clan leader cannot personally surrender.
	LeaderCantPersonallySurrender = SysMsg{Id: 254}

	// The {{S1}} clan has requested to end war. Do you agree?
	S1ClanRequestedEndWar = SysMsg{Id: 255}

	// Enter title
	EnterTitle = SysMsg{Id: 256}

	// Do you offer the {{S1}} clan a proposal to end the war?
	DoYouOfferS1ClanEndWar = SysMsg{Id: 257}

	// You are not involved in a clan war.
	NotInvolvedClanWar = SysMsg{Id: 258}

	// Select clan members from list.
	SelectMembersFromList = SysMsg{Id: 259}

	// Fame level has decreased: 5 days have not passed since you were refused war
	FiveDaysNotPassedSinceYouWereRefusedWar = SysMsg{Id: 260}

	// Clan name is invalid.
	ClanNameIncorrect = SysMsg{Id: 261}

	// Clan name's length is incorrect.
	ClanNameTooLong = SysMsg{Id: 262}

	// You have already requested the dissolution of your clan.
	DissolutionInProgress = SysMsg{Id: 263}

	// You cannot dissolve a clan while engaged in a war.
	CannotDissolveWhileInWar = SysMsg{Id: 264}

	// You cannot dissolve a clan during a siege or while protecting a castle.
	CannotDissolveWhileInSiege = SysMsg{Id: 265}

	// You cannot dissolve a clan while owning a clan hall or castle.
	CannotDissolveWhileOwningClanHallOrCastle = SysMsg{Id: 266}

	// There are no requests to disperse.
	NoRequestsToDisperse = SysMsg{Id: 267}

	// That player already belongs to another clan.
	PlayerAlreadyAnotherClan = SysMsg{Id: 268}

	// You cannot dismiss yourself.
	YouCannotDismissYourself = SysMsg{Id: 269}

	// You have already surrendered.
	YouHaveAlreadySurrendered = SysMsg{Id: 270}

	// A player can only be granted a title if the clan is level 3 or above
	ClanLvl3NeededToEndoweTitle = SysMsg{Id: 271}

	// A clan crest can only be registered when the clan's skill level is 3 or above.
	ClanLvl3NeededToSetCrest = SysMsg{Id: 272}

	// A clan war can only be declared when a clan's skill level is 3 or above.
	ClanLvl3NeededToDeclareWar = SysMsg{Id: 273}

	// Your clan's skill level has increased.
	ClanLevelIncreased = SysMsg{Id: 274}

	// Clan has failed to increase skill level.
	ClanLevelIncreaseFailed = SysMsg{Id: 275}

	// You do not have the necessary materials or prerequisites to learn this skill.
	ItemOrPrerequisitesMissingToLearnSkill = SysMsg{Id: 276}

	// You have earned {{S1}}.
	LearnedSkillS1 = SysMsg{Id: 277}

	// You do not have enough SP to learn this skill.
	NotEnoughSpToLearnSkill = SysMsg{Id: 278}

	// You do not have enough adena.
	YouNotEnoughAdena = SysMsg{Id: 279}

	// You do not have any items to sell.
	NoItemsToSell = SysMsg{Id: 280}

	// You do not have enough adena to pay the fee.
	YouNotEnoughAdenaPayFee = SysMsg{Id: 281}

	// You have not deposited any items in your warehouse.
	NoItemDepositedInWh = SysMsg{Id: 282}

	// You have entered a combat zone.
	EnteredCombatZone = SysMsg{Id: 283}

	// You have left a combat zone.
	LeftCombatZone = SysMsg{Id: 284}

	// Clan {{S1}} has succeeded in engraving the ruler!
	ClanS1EngravedRuler = SysMsg{Id: 285}

	// Your base is being attacked.
	BaseUnderAttack = SysMsg{Id: 286}

	// The opposing clan has stared to engrave to monument!
	OpponentStartedEngraving = SysMsg{Id: 287}

	// The castle gate has been broken down.
	CastleGateBrokenDown = SysMsg{Id: 288}

	// An outpost or headquarters cannot be built because at least one already exists.
	NotAnotherHeadquarters = SysMsg{Id: 289}

	// You cannot set up a base here.
	NotSetUpBaseHere = SysMsg{Id: 290}

	// Clan {{S1}} is victorious over {{S2}}'s castle siege!
	ClanS1VictoriousOverS2SSiege = SysMsg{Id: 291}

	// {{S1}} has announced the castle siege time.
	S1AnnouncedSiegeTime = SysMsg{Id: 292}

	// The registration term for {{S1}} has ended.
	RegistrationTermForS1Ended = SysMsg{Id: 293}

	// Because your clan is not currently on the offensive in a Clan Hall siege war, it cannot summon its base camp.
	BecauseYourClanIsNotCurrentlyOnTheOffensiveInAClanHallSiegeWarItCannotSummonItsBaseCamp = SysMsg{Id: 294}

	// {{S1}}'s siege was canceled because there were no clans that participated.
	S1SiegeWasCanceledBecauseNoClansParticipated = SysMsg{Id: 295}

	// You received {{S1}} damage from taking a high fall.
	FallDamageS1 = SysMsg{Id: 296}

	// You have taken {{S1}} damage because you were unable to breathe.
	DrownDamageS1 = SysMsg{Id: 297}

	// You have dropped {{S1}}.
	YouDroppedS1 = SysMsg{Id: 298}

	// {{C1}} has obtained {{S3}} {{S2}}.
	C1ObtainedS3S2 = SysMsg{Id: 299}

	// {{C1}} has obtained {{S2}}.
	C1ObtainedS2 = SysMsg{Id: 300}

	// {{S2}} {{S1}} has disappeared.
	S2S1Disappeared = SysMsg{Id: 301}

	// {{S1}} has disappeared.
	S1Disappeared = SysMsg{Id: 302}

	// Select item to enchant.
	SelectItemToEnchant = SysMsg{Id: 303}

	// Clan member {{S1}} has logged into game.
	ClanMemberS1LoggedIn = SysMsg{Id: 304}

	// The player declined to join your party.
	PlayerDeclined = SysMsg{Id: 305}

	// You have failed to delete the character.
	FailedToDeleteChar = SysMsg{Id: 306}

	// You cannot trade with a warehouse keeper.
	CannotTradeWarehouseKeeper = SysMsg{Id: 307}

	// The player declined your clan invitation.
	PlayerDeclinedClanInvitation = SysMsg{Id: 308}

	// You have succeeded in expelling the clan member.
	YouHaveSucceededInExpellingClanMember = SysMsg{Id: 309}

	// You have failed to expel the clan member.
	FailedToExpelClanMember = SysMsg{Id: 310}

	// The clan war declaration has been accepted.
	ClanWarDeclarationAccepted = SysMsg{Id: 311}

	// The clan war declaration has been refused.
	ClanWarDeclarationRefused = SysMsg{Id: 312}

	// The cease war request has been accepted.
	CeaseWarRequestAccepted = SysMsg{Id: 313}

	// You have failed to surrender.
	FailedToSurrender = SysMsg{Id: 314}

	// You have failed to personally surrender.
	FailedToPersonallySurrender = SysMsg{Id: 315}

	// You have failed to withdraw from the party.
	FailedToWithdrawFromTheParty = SysMsg{Id: 316}

	// You have failed to expel the party member.
	FailedToExpelThePartyMember = SysMsg{Id: 317}

	// You have failed to disperse the party.
	FailedToDisperseTheParty = SysMsg{Id: 318}

	// This door cannot be unlocked.
	UnableToUnlockDoor = SysMsg{Id: 319}

	// You have failed to unlock the door.
	FailedToUnlockDoor = SysMsg{Id: 320}

	// It is not locked.
	ItsNotLocked = SysMsg{Id: 321}

	// Please decide on the sales price.
	DecideSalesPrice = SysMsg{Id: 322}

	// Your force has increased to {{S1}} level.
	ForceIncreasedToS1 = SysMsg{Id: 323}

	// Your force has reached maximum capacity.
	ForceMaxlevelReached = SysMsg{Id: 324}

	// The corpse has already disappeared.
	CorpseAlreadyDisappeared = SysMsg{Id: 325}

	// Select target from list.
	SelectTargetFromList = SysMsg{Id: 326}

	// You cannot exceed 80 characters.
	CannotExceed80Characters = SysMsg{Id: 327}

	// Please input title using less than 128 characters.
	PleaseInputTitleLess128Characters = SysMsg{Id: 328}

	// Please input content using less than 3000 characters.
	PleaseInputContentLess3000Characters = SysMsg{Id: 329}

	// A one-line response may not exceed 128 characters.
	OneLineResponseNotExceed128Characters = SysMsg{Id: 330}

	// You have acquired {{S1}} SP.
	AcquiredS1Sp = SysMsg{Id: 331}

	// Do you want to be restored?
	DoYouWantToBeRestored = SysMsg{Id: 332}

	// You have received {{S1}} damage by Core's barrier.
	S1DamageByCoreBarrier = SysMsg{Id: 333}

	// Please enter your private store display message.
	EnterPrivateStoreMessage = SysMsg{Id: 334}

	// {{S1}} has been aborted.
	S1HasBeenAborted = SysMsg{Id: 335}

	// You are attempting to crystallize {{S1}}. Do you wish to continue?
	WishToCrystallizeS1 = SysMsg{Id: 336}

	// The soulshot you are attempting to use does not match the grade of your equipped weapon.
	SoulshotsGradeMismatch = SysMsg{Id: 337}

	// You do not have enough soulshots for that.
	NotEnoughSoulshots = SysMsg{Id: 338}

	// Cannot use soulshots.
	CannotUseSoulshots = SysMsg{Id: 339}

	// Your private store is now open for business.
	PrivateStoreUnderWay = SysMsg{Id: 340}

	// You do not have enough materials to perform that action.
	NotEnoughMaterials = SysMsg{Id: 341}

	// Power of the spirits enabled.
	EnabledSoulshot = SysMsg{Id: 342}

	// Sweeper failed, target not spoiled.
	SweeperFailedTargetNotSpoiled = SysMsg{Id: 343}

	// Power of the spirits disabled.
	SoulshotsDisabled = SysMsg{Id: 344}

	// Chat enabled.
	ChatEnabled = SysMsg{Id: 345}

	// Chat disabled.
	ChatDisabled = SysMsg{Id: 346}

	// Incorrect item count.
	IncorrectItemCount = SysMsg{Id: 347}

	// Incorrect item price.
	IncorrectItemPrice = SysMsg{Id: 348}

	// Private store already closed.
	PrivateStoreAlreadyClosed = SysMsg{Id: 349}

	// Item out of stock.
	ItemOutOfStock = SysMsg{Id: 350}

	// Incorrect item count.
	NotEnoughItems = SysMsg{Id: 351}

	// Incorrect item.
	IncorrectItem = SysMsg{Id: 352}

	// Cannot purchase.
	CannotPurchase = SysMsg{Id: 353}

	// Cancel enchant.
	CancelEnchant = SysMsg{Id: 354}

	// Inappropriate enchant conditions.
	InappropriateEnchantCondition = SysMsg{Id: 355}

	// Reject resurrection.
	RejectResurrection = SysMsg{Id: 356}

	// It has already been spoiled.
	AlreadySpoiled = SysMsg{Id: 357}

	// {{S1}} hour(s) until castle siege conclusion.
	S1HoursUntilSiegeConclusion = SysMsg{Id: 358}

	// {{S1}} minute(s) until castle siege conclusion.
	S1MinutesUntilSiegeConclusion = SysMsg{Id: 359}

	// Castle siege {{S1}} second(s) left!
	CastleSiegeS1SecondsLeft = SysMsg{Id: 360}

	// Over-hit!
	OverHit = SysMsg{Id: 361}

	// You have acquired {{S1}} bonus experience from a successful over-hit.
	AcquiredBonusExperienceThroughOverHit = SysMsg{Id: 362}

	// Chat available time: {{S1}} minute.
	ChatAvailableS1Minute = SysMsg{Id: 363}

	// Enter user's name to search
	EnterUserNameToSearch = SysMsg{Id: 364}

	// Are you sure?
	AreYouSure = SysMsg{Id: 365}

	// Please select your hair color.
	PleaseSelectHairColor = SysMsg{Id: 366}

	// You cannot remove that clan character at this time.
	CannotRemoveClanCharacter = SysMsg{Id: 367}

	// Equipped +{{S1}} {{S2}}.
	S1S2Equipped = SysMsg{Id: 368}

	// You have obtained a +{{S1}} {{S2}}.
	YouPickedUpAS1S2 = SysMsg{Id: 369}

	// Failed to pickup {{S1}}.
	FailedPickupS1 = SysMsg{Id: 370}

	// Acquired +{{S1}} {{S2}}.
	AcquiredS1S2 = SysMsg{Id: 371}

	// Failed to earn {{S1}}.
	FailedEarnS1 = SysMsg{Id: 372}

	// You are trying to destroy +{{S1}} {{S2}}. Do you wish to continue?
	WishDestroyS1S2 = SysMsg{Id: 373}

	// You are attempting to crystallize +{{S1}} {{S2}}. Do you wish to continue?
	WishCrystallizeS1S2 = SysMsg{Id: 374}

	// You have dropped +{{S1}} {{S2}} .
	DroppedS1S2 = SysMsg{Id: 375}

	// {{C1}} has obtained +{{S2}}{{S3}}.
	C1ObtainedS2S3 = SysMsg{Id: 376}

	// {{S1}} {{S2}} disappeared.
	S1S2Disappeared = SysMsg{Id: 377}

	// {{C1}} purchased {{S2}}.
	C1PurchasedS2 = SysMsg{Id: 378}

	// {{C1}} purchased +{{S2}}{{S3}}.
	C1PurchasedS2S3 = SysMsg{Id: 379}

	// {{C1}} purchased {{S3}} {{S2}}(s).
	C1PurchasedS3S2S = SysMsg{Id: 380}

	// The game client encountered an error and was unable to connect to the petition server.
	GameClientUnableToConnectToPetitionServer = SysMsg{Id: 381}

	// Currently there are no users that have checked out a GM ID.
	NoUsersCheckedOutGmId = SysMsg{Id: 382}

	// Request confirmed to end consultation at petition server.
	RequestConfirmedToEndConsultation = SysMsg{Id: 383}

	// The client is not logged onto the game server.
	ClientNotLoggedOntoGameServer = SysMsg{Id: 384}

	// Request confirmed to begin consultation at petition server.
	RequestConfirmedToBeginConsultation = SysMsg{Id: 385}

	// The body of your petition must be more than five characters in length.
	PetitionMoreThanFiveCharacters = SysMsg{Id: 386}

	// This ends the GM petition consultation. Please take a moment to provide feedback about this service.
	ThisEndThePetitionPleaseProvideFeedback = SysMsg{Id: 387}

	// Not under petition consultation.
	NotUnderPetitionConsultation = SysMsg{Id: 388}

	// our petition application has been accepted. - Receipt No. is {{S1}}.
	PetitionAcceptedRecentNoS1 = SysMsg{Id: 389}

	// You may only submit one petition (active) at a time.
	OnlyOneActivePetitionAtTime = SysMsg{Id: 390}

	// Receipt No. {{S1}}, petition cancelled.
	RecentNoS1Canceled = SysMsg{Id: 391}

	// Under petition advice.
	UnderPetitionAdvice = SysMsg{Id: 392}

	// Failed to cancel petition. Please try again later.
	FailedCancelPetitionTryLater = SysMsg{Id: 393}

	// Starting petition consultation with {{C1}}.
	StartingPetitionWithC1 = SysMsg{Id: 394}

	// Ending petition consultation with {{C1}}.
	PetitionEndedWithC1 = SysMsg{Id: 395}

	// Please login after changing your temporary password.
	TryAgainAfterChangingPassword = SysMsg{Id: 396}

	// Not a paid account.
	NoPaidAccount = SysMsg{Id: 397}

	// There is no time left on this account.
	NoTimeLeftOnAccount = SysMsg{Id: 398}

	// System error.
	SystemError = SysMsg{Id: 399}

	// You are attempting to drop {{S1}}. Dou you wish to continue?
	WishToDropS1 = SysMsg{Id: 400}

	// You have to many ongoing quests.
	TooManyQuests = SysMsg{Id: 401}

	// You do not possess the correct ticket to board the boat.
	NotCorrectBoatTicket = SysMsg{Id: 402}

	// You have exceeded your out-of-pocket adena limit.
	ExceecedPocketAdenaLimit = SysMsg{Id: 403}

	// Your Create Item level is too low to register this recipe.
	CreateLvlTooLowToRegister = SysMsg{Id: 404}

	// The total price of the product is too high.
	TotalPriceTooHigh = SysMsg{Id: 405}

	// Petition application accepted.
	PetitionAppAccepted = SysMsg{Id: 406}

	// Petition under process.
	PetitionUnderProcess = SysMsg{Id: 407}

	// Set Period
	SetPeriod = SysMsg{Id: 408}

	// Set Time-{{S1}}:{{S2}}:{{S3}}
	SetTimeS1S2S3 = SysMsg{Id: 409}

	// Registration Period
	RegistrationPeriod = SysMsg{Id: 410}

	// Registration Time-{{S1}}:{{S2}}:{{S3}}
	RegistrationTimeS1S2S3 = SysMsg{Id: 411}

	// Battle begins in {{S1}}:{{S2}}:{{S3}}
	BattleBeginsS1S2S3 = SysMsg{Id: 412}

	// Battle ends in {{S1}}:{{S2}}:{{S3}}
	BattleEndsS1S2S3 = SysMsg{Id: 413}

	// Standby
	Standby = SysMsg{Id: 414}

	// Under Siege
	UnderSiege = SysMsg{Id: 415}

	// This item cannot be exchanged.
	ItemCannotExchange = SysMsg{Id: 416}

	// {{S1}} has been disarmed.
	S1Disarmed = SysMsg{Id: 417}

	// {{S1}} minute(s) of usage time left.
	S1MinutesUsageLeft = SysMsg{Id: 419}

	// Time expired.
	TimeExpired = SysMsg{Id: 420}

	// Another person has logged in with the same account.
	AnotherLoginWithAccount = SysMsg{Id: 421}

	// You have exceeded the weight limit.
	WeightLimitExceeded = SysMsg{Id: 422}

	// You have cancelled the enchanting process.
	EnchantScrollCancelled = SysMsg{Id: 423}

	// Does not fit strengthening conditions of the scroll.
	DoesNotFitScrollConditions = SysMsg{Id: 424}

	// Your Create Item level is too low to register this recipe.
	CreateLvlTooLowToRegister2 = SysMsg{Id: 425}

	// (Reference Number Regarding Membership Withdrawal Request: {{S1}})
	ReferenceMembershipWithdrawalS1 = SysMsg{Id: 445}

	// .
	Dot = SysMsg{Id: 447}

	// There is a system error. Please log in again later.
	SystemErrorLoginLater = SysMsg{Id: 448}

	// The password you have entered is incorrect.
	PasswordEnteredIncorrect1 = SysMsg{Id: 449}

	// Confirm your account information and log in later.
	ConfirmAccountLoginLater = SysMsg{Id: 450}

	// The password you have entered is incorrect.
	PasswordEnteredIncorrect2 = SysMsg{Id: 451}

	// Please confirm your account information and try logging in later.
	PleaseConfirmAccountLoginLater = SysMsg{Id: 452}

	// Your account information is incorrect.
	AccountInformationIncorrect = SysMsg{Id: 453}

	// Account is already in use. Unable to log in.
	AccountInUse = SysMsg{Id: 455}

	// Lineage II game services may be used by individuals 15 years of age or older except for PvP servers,which may only be used by adults 18 years of age and older (Korea Only)
	LinageMinimumAge = SysMsg{Id: 456}

	// Currently undergoing game server maintenance. Please log in again later.
	ServerMaintenance = SysMsg{Id: 457}

	// Your usage term has expired.
	UsageTermExpired = SysMsg{Id: 458}

	// to reactivate your account.
	ToReactivateYourAccount = SysMsg{Id: 460}

	// Access failed.
	AccessFailed = SysMsg{Id: 461}

	// Please try again later.
	PleaseTryAgainLater = SysMsg{Id: 462}

	// This feature is only available alliance leaders.
	FeatureOnlyForAllianceLeader = SysMsg{Id: 464}

	// You are not currently allied with any clans.
	NoCurrentAlliances = SysMsg{Id: 465}

	// You have exceeded the limit.
	YouHaveExceededTheLimit = SysMsg{Id: 466}

	// You may not accept any clan within a day after expelling another clan.
	CantInviteClanWithin1Day = SysMsg{Id: 467}

	// A clan that has withdrawn or been expelled cannot enter into an alliance within one day of withdrawal or expulsion.
	CantEnterAllianceWithin1Day = SysMsg{Id: 468}

	// You may not ally with a clan you are currently at war with. That would be diabolical and treacherous.
	MayNotAllyClanBattle = SysMsg{Id: 469}

	// Only the clan leader may apply for withdrawal from the alliance.
	OnlyClanLeaderWithdrawAlly = SysMsg{Id: 470}

	// Alliance leaders cannot withdraw.
	AllianceLeaderCantWithdraw = SysMsg{Id: 471}

	// You cannot expel yourself from the clan.
	CannotExpelYourself = SysMsg{Id: 472}

	// Different alliance.
	DifferentAlliance = SysMsg{Id: 473}

	// That clan does not exist.
	ClanDoesntExists = SysMsg{Id: 474}

	// Different alliance.
	DifferentAlliance2 = SysMsg{Id: 475}

	// Please adjust the image size to 8x12.
	AdjustImage812 = SysMsg{Id: 476}

	// No response. Invitation to join an alliance has been cancelled.
	NoResponseToAllyInvitation = SysMsg{Id: 477}

	// No response. Your entrance to the alliance has been cancelled.
	YouDidNotRespondToAllyInvitation = SysMsg{Id: 478}

	// {{S1}} has joined as a friend.
	S1JoinedAsFriend = SysMsg{Id: 479}

	// Please check your friend list.
	PleaseCheckYourFriendsList = SysMsg{Id: 480}

	// {{S1}} has been deleted from your friends list.
	S1HasBeenDeletedFromYourFriendsList = SysMsg{Id: 481}

	// You cannot add yourself to your own friend list.
	YouCannotAddYourselfToYourOwnFriendsList = SysMsg{Id: 482}

	// This function is inaccessible right now. Please try again later.
	FunctionInaccessibleNow = SysMsg{Id: 483}

	// This player is already registered in your friends list.
	S1AlreadyInFriendsList = SysMsg{Id: 484}

	// No new friend invitations may be accepted.
	NoNewInvitationsAccepted = SysMsg{Id: 485}

	// The following user is not in your friends list.
	TheUserNotInFriendsList = SysMsg{Id: 486}

	// ======<Friends List>======
	FriendListHeader = SysMsg{Id: 487}

	// {{S1}} (Currently: Online)
	S1Online = SysMsg{Id: 488}

	// {{S1}} (Currently: Offline)
	S1Offline = SysMsg{Id: 489}

	// ========================
	FriendListFooter = SysMsg{Id: 490}

	// =======<Alliance Information>=======
	AllianceInfoHead = SysMsg{Id: 491}

	// Alliance Name: {{S1}}
	AllianceNameS1 = SysMsg{Id: 492}

	// Connection: {{S1}} / Total {{S2}}
	ConnectionS1TotalS2 = SysMsg{Id: 493}

	// Alliance Leader: {{S2}} of {{S1}}
	AllianceLeaderS2OfS1 = SysMsg{Id: 494}

	// Affiliated clans: Total {{S1}} clan(s)
	AllianceClanTotalS1 = SysMsg{Id: 495}

	// =====<Clan Information>=====
	ClanInfoHead = SysMsg{Id: 496}

	// Clan Name: {{S1}}
	ClanInfoNameS1 = SysMsg{Id: 497}

	// Clan Leader: {{S1}}
	ClanInfoLeaderS1 = SysMsg{Id: 498}

	// Clan Level: {{S1}}
	ClanInfoLevelS1 = SysMsg{Id: 499}

	// ------------------------
	ClanInfoSeparator = SysMsg{Id: 500}

	// ========================
	ClanInfoFoot = SysMsg{Id: 501}

	// You already belong to another alliance.
	AlreadyJoinedAlliance = SysMsg{Id: 502}

	// {{S1}} (Friend) has logged in.
	FriendS1HasLoggedIn = SysMsg{Id: 503}

	// Only clan leaders may create alliances.
	OnlyClanLeaderCreateAlliance = SysMsg{Id: 504}

	// You cannot create a new alliance within 10 days after dissolution.
	CantCreateAlliance10DaysDisolution = SysMsg{Id: 505}

	// Incorrect alliance name. Please try again.
	IncorrectAllianceName = SysMsg{Id: 506}

	// Incorrect length for an alliance name.
	IncorrectAllianceNameLength = SysMsg{Id: 507}

	// This alliance name already exists.
	AllianceAlreadyExists = SysMsg{Id: 508}

	// Cannot accept. clan ally is registered as an enemy during siege battle.
	CantAcceptAllyEnemyForSiege = SysMsg{Id: 509}

	// You have invited someone to your alliance.
	YouInvitedForAlliance = SysMsg{Id: 510}

	// You must first select a user to invite.
	SelectUserToInvite = SysMsg{Id: 511}

	// Do you really wish to withdraw from the alliance?
	DoYouWishToWithdrw = SysMsg{Id: 512}

	// Enter the name of the clan you wish to expel.
	EnterNameClanToExpel = SysMsg{Id: 513}

	// Do you really wish to dissolve the alliance?
	DoYouWishToDisolve = SysMsg{Id: 514}

	// {{S1}} has invited you to be their friend.
	SiInvitedYouAsFriend = SysMsg{Id: 516}

	// You have accepted the alliance.
	YouAcceptedAlliance = SysMsg{Id: 517}

	// You have failed to invite a clan into the alliance.
	FailedToInviteClanInAlliance = SysMsg{Id: 518}

	// You have withdrawn from the alliance.
	YouHaveWithdrawnFromAlliance = SysMsg{Id: 519}

	// You have failed to withdraw from the alliance.
	YouHaveFailedToWithdrawnFromAlliance = SysMsg{Id: 520}

	// You have succeeded in expelling a clan.
	YouHaveExpeledAClan = SysMsg{Id: 521}

	// You have failed to expel a clan.
	FailedToExpeledAClan = SysMsg{Id: 522}

	// The alliance has been dissolved.
	AllianceDisolved = SysMsg{Id: 523}

	// You have failed to dissolve the alliance.
	FailedToDisolveAlliance = SysMsg{Id: 524}

	// You have succeeded in inviting a friend to your friends list.
	YouHaveSucceededInvitingFriend = SysMsg{Id: 525}

	// You have failed to add a friend to your friends list.
	FailedToInviteAFriend = SysMsg{Id: 526}

	// {{S1}} leader, {{S2}}, has requested an alliance.
	S2AllianceLeaderOfS1RequestedAlliance = SysMsg{Id: 527}

	// The Spiritshot does not match the weapon's grade.
	SpiritshotsGradeMismatch = SysMsg{Id: 530}

	// You do not have enough Spiritshots for that.
	NotEnoughSpiritshots = SysMsg{Id: 531}

	// You may not use Spiritshots.
	CannotUseSpiritshots = SysMsg{Id: 532}

	// Power of Mana enabled.
	EnabledSpiritshot = SysMsg{Id: 533}

	// Power of Mana disabled.
	DisabledSpiritshot = SysMsg{Id: 534}

	// How much adena do you wish to transfer to your Inventory?
	HowMuchAdenaTransfer = SysMsg{Id: 536}

	// How much will you transfer?
	HowMuchTransfer = SysMsg{Id: 537}

	// Your SP has decreased by {{S1}}.
	SpDecreasedS1 = SysMsg{Id: 538}

	// Your Experience has decreased by {{S1}}.
	ExpDecreasedByS1 = SysMsg{Id: 539}

	// Clan leaders may not be deleted. Dissolve the clan first and try again.
	ClanLeadersMayNotBeDeleted = SysMsg{Id: 540}

	// You may not delete a clan member. Withdraw from the clan first and try again.
	ClanMemberMayNotBeDeleted = SysMsg{Id: 541}

	// The NPC server is currently down. Pets and servitors cannot be summoned at this time.
	TheNpcServerIsCurrentlyDown = SysMsg{Id: 542}

	// You already have a pet.
	YouAlreadyHaveAPet = SysMsg{Id: 543}

	// Your pet cannot carry this item.
	ItemNotForPets = SysMsg{Id: 544}

	// Your pet cannot carry any more items. Remove some, then try again.
	YourPetCannotCarryAnyMoreItems = SysMsg{Id: 545}

	// Unable to place item, your pet is too encumbered.
	UnableToPlaceItemYourPetIsTooEncumbered = SysMsg{Id: 546}

	// Summoning your pet.
	SummonAPet = SysMsg{Id: 547}

	// Your pet's name can be up to 8 characters in length.
	NamingPetnameUpTo8Chars = SysMsg{Id: 548}

	// To create an alliance, your clan must be Level 5 or higher.
	ToCreateAnAllyYouClanMustBeLevel5OrHigher = SysMsg{Id: 549}

	// You may not create an alliance during the term of dissolution postponement.
	YouMayNotCreateAllyWhileDissolving = SysMsg{Id: 550}

	// You cannot raise your clan level during the term of dispersion postponement.
	CannotRiseLevelWhileDissolutionInProgress = SysMsg{Id: 551}

	// During the grace period for dissolving a clan, the registration or deletion of a clan's crest is not allowed.
	CannotSetCrestWhileDissolutionInProgress = SysMsg{Id: 552}

	// The opposing clan has applied for dispersion.
	OpposingClanAppliedDispersion = SysMsg{Id: 553}

	// You cannot disperse the clans in your alliance.
	CannotDisperseTheClansInAlly = SysMsg{Id: 554}

	// You cannot move - you are too encumbered
	CantMoveTooEncumbered = SysMsg{Id: 555}

	// You cannot move in this state
	CantMoveInThisState = SysMsg{Id: 556}

	// Your pet has been summoned and may not be destroyed
	PetSummonedMayNotDestroyed = SysMsg{Id: 557}

	// Your pet has been summoned and may not be let go.
	PetSummonedMayNotLetGo = SysMsg{Id: 558}

	// You have purchased {{S2}} from {{C1}}.
	PurchasedS2FromC1 = SysMsg{Id: 559}

	// You have purchased +{{S2}} {{S3}} from {{C1}}.
	PurchasedS2S3FromC1 = SysMsg{Id: 560}

	// You have purchased {{S3}} {{S2}}(s) from {{C1}}.
	PurchasedS3S2SFromC1 = SysMsg{Id: 561}

	// You may not crystallize this item. Your crystallization skill level is too low.
	CrystallizeLevelTooLow = SysMsg{Id: 562}

	// Failed to disable attack target.
	FailedDisableTarget = SysMsg{Id: 563}

	// Failed to change attack target.
	FailedChangeTarget = SysMsg{Id: 564}

	// Not enough luck.
	NotEnoughLuck = SysMsg{Id: 565}

	// Your confusion spell failed.
	ConfusionFailed = SysMsg{Id: 566}

	// Your fear spell failed.
	FearFailed = SysMsg{Id: 567}

	// Cubic Summoning failed.
	CubicSummoningFailed = SysMsg{Id: 568}

	// Do you accept {{C1}}'s party invitation? (Item Distribution: Finders Keepers.)
	C1InvitedYouToPartyFindersKeepers = SysMsg{Id: 572}

	// Do you accept {{C1}}'s party invitation? (Item Distribution: Random.)
	C1InvitedYouToPartyRandom = SysMsg{Id: 573}

	// Pets and Servitors are not available at this time.
	PetsAreNotAvailableAtThisTime = SysMsg{Id: 574}

	// How much adena do you wish to transfer to your pet?
	HowMuchAdenaTransferToPet = SysMsg{Id: 575}

	// How much do you wish to transfer?
	HowMuchTransfer2 = SysMsg{Id: 576}

	// You cannot summon during a trade or while using the private shops.
	CannotSummonDuringTradeShop = SysMsg{Id: 577}

	// You cannot summon during combat.
	YouCannotSummonInCombat = SysMsg{Id: 578}

	// A pet cannot be sent back during battle.
	PetCannotSentBackDuringBattle = SysMsg{Id: 579}

	// You may not use multiple pets or servitors at the same time.
	SummonOnlyOne = SysMsg{Id: 580}

	// There is a space in the name.
	NamingThereIsASpace = SysMsg{Id: 581}

	// Inappropriate character name.
	NamingInappropriateCharacterName = SysMsg{Id: 582}

	// Name includes forbidden words.
	NamingIncludesForbiddenWords = SysMsg{Id: 583}

	// This is already in use by another pet.
	NamingAlreadyInUseByAnotherPet = SysMsg{Id: 584}

	// Please decide on the price.
	DecideOnPrice = SysMsg{Id: 585}

	// Pet items cannot be registered as shortcuts.
	PetNoShortcut = SysMsg{Id: 586}

	// Your pet's inventory is full.
	PetInventoryFull = SysMsg{Id: 588}

	// A dead pet cannot be sent back.
	DeadPetCannotBeReturned = SysMsg{Id: 589}

	// Your pet is motionless and any attempt you make to give it something goes unrecognized.
	CannotGiveItemsToDeadPet = SysMsg{Id: 590}

	// An invalid character is included in the pet's name.
	NamingPetnameContainsInvalidChars = SysMsg{Id: 591}

	// Do you wish to dismiss your pet? Dismissing your pet will cause the pet necklace to disappear
	WishToDismissPet = SysMsg{Id: 592}

	// Starving, grumpy and fed up, your pet has left.
	StarvingGrumpyAndFedUpYourPetHasLeft = SysMsg{Id: 593}

	// You may not restore a hungry pet.
	YouCannotRestoreHungryPets = SysMsg{Id: 594}

	// Your pet is very hungry.
	YourPetIsVeryHungry = SysMsg{Id: 595}

	// Your pet ate a little, but is still hungry.
	YourPetAteALittleButIsStillHungry = SysMsg{Id: 596}

	// Your pet is very hungry. Please be careful.
	YourPetIsVeryHungryPleaseBeCarefull = SysMsg{Id: 597}

	// You may not chat while you are invisible.
	NotChatWhileInvisible = SysMsg{Id: 598}

	// The GM has an important notice. Chat has been temporarily disabled.
	GmNoticeChatDisabled = SysMsg{Id: 599}

	// You may not equip a pet item.
	CannotEquipPetItem = SysMsg{Id: 600}

	// There are {{S1}} petitions currently on the waiting list.
	S1PetitionOnWaitingList = SysMsg{Id: 601}

	// The petition system is currently unavailable. Please try again later.
	PetitionSystemCurrentUnavailable = SysMsg{Id: 602}

	// That item cannot be discarded or exchanged.
	CannotDiscardExchangeItem = SysMsg{Id: 603}

	// You may not call forth a pet or summoned creature from this location
	NotCallPetFromThisLocation = SysMsg{Id: 604}

	// You may register up to 64 people on your list.
	MayRegisterUpTo64People = SysMsg{Id: 605}

	// You cannot be registered because the other person has already registered 64 people on his/her list.
	OtherPersonAlready64People = SysMsg{Id: 606}

	// You do not have any further skills to learn. Come back when you have reached Level {{S1}}.
	DoNotHaveFurtherSkillsToLearnS1 = SysMsg{Id: 607}

	// {{C1}} has obtained {{S3}} {{S2}} by using Sweeper.
	C1SweepedUpS3S2 = SysMsg{Id: 608}

	// {{C1}} has obtained {{S2}} by using Sweeper.
	C1SweepedUpS2 = SysMsg{Id: 609}

	// Your skill has been canceled due to lack of HP.
	SkillRemovedDueLackHp = SysMsg{Id: 610}

	// You have succeeded in Confusing the enemy.
	ConfusingSucceeded = SysMsg{Id: 611}

	// The Spoil condition has been activated.
	SpoilSuccess = SysMsg{Id: 612}

	// ======<Ignore List>======
	BlockListHeader = SysMsg{Id: 613}

	// {{C1}} : $c2
	C1DC2 = SysMsg{Id: 614}

	// You have failed to register the user to your Ignore List.
	FailedToRegisterToIgnoreList = SysMsg{Id: 615}

	// You have failed to delete the character.
	FailedToDeleteCharacter = SysMsg{Id: 616}

	// {{S1}} has been added to your Ignore List.
	S1WasAddedToYourIgnoreList = SysMsg{Id: 617}

	// {{S1}} has been removed from your Ignore List.
	S1WasRemovedFromYourIgnoreList = SysMsg{Id: 618}

	// {{S1}} has placed you on his/her Ignore List.
	S1HasAddedYouToIgnoreList = SysMsg{Id: 619}

	// {{S1}} has placed you on his/her Ignore List.
	S1HasAddedYouToIgnoreList2 = SysMsg{Id: 620}

	// Game connection attempted through a restricted IP.
	ConnectionRestrictedIp = SysMsg{Id: 621}

	// You may not make a declaration of war during an alliance battle.
	NoWarDuringAllyBattle = SysMsg{Id: 622}

	// Your opponent has exceeded the number of simultaneous alliance battles alllowed.
	OpponentTooMuchAllyBattles1 = SysMsg{Id: 623}

	// {{S1}} Clan leader is not currently connected to the game server.
	S1LeaderNotConnected = SysMsg{Id: 624}

	// Your request for Alliance Battle truce has been denied.
	AllyBattleTruceDenied = SysMsg{Id: 625}

	// The {{S1}} clan did not respond: war proclamation has been refused.
	WarProclamationHasBeenRefused = SysMsg{Id: 626}

	// Clan battle has been refused because you did not respond to {{S1}} clan's war proclamation.
	YouRefusedClanWarProclamation = SysMsg{Id: 627}

	// You have already been at war with the {{S1}} clan: 5 days must pass before you can declare war again.
	AlreadyAtWarWithS1Wait5Days = SysMsg{Id: 628}

	// Your opponent has exceeded the number of simultaneous alliance battles alllowed.
	OpponentTooMuchAllyBattles2 = SysMsg{Id: 629}

	// War with the clan has begun.
	WarWithClanBegun = SysMsg{Id: 630}

	// War with the clan is over.
	WarWithClanEnded = SysMsg{Id: 631}

	// You have won the war over the clan!
	WonWarOverClan = SysMsg{Id: 632}

	// You have surrendered to the clan.
	SurrenderedToClan = SysMsg{Id: 633}

	// Your alliance leader has been slain. You have been defeated by the clan.
	DefeatedByClan = SysMsg{Id: 634}

	// The time limit for the clan war has been exceeded. War with the clan is over.
	TimeUpWarOver = SysMsg{Id: 635}

	// You are not involved in a clan war.
	NotInvolvedInWar = SysMsg{Id: 636}

	// A clan ally has registered itself to the opponent.
	AllyRegisteredSelfToOpponent = SysMsg{Id: 637}

	// You have already requested a Siege Battle.
	AlreadyRequestedSiegeBattle = SysMsg{Id: 638}

	// Your application has been denied because you have already submitted a request for another Siege Battle.
	ApplicationDeniedBecauseAlreadySubmittedARequestForAnotherSiegeBattle = SysMsg{Id: 639}

	// You have failed to refuse castle defense aid.
	FailedToRefuseCastleDefenseAid = SysMsg{Id: 640}

	// You have failed to approve castle defense aid.
	FailedToApproveCastleDefenseAid = SysMsg{Id: 641}

	// You are already registered to the attacker side and must cancel your registration before submitting your request.
	AlreadyAttackerNotCancel = SysMsg{Id: 642}

	// You have already registered to the defender side and must cancel your registration before submitting your request.
	AlreadyDefenderNotCancel = SysMsg{Id: 643}

	// You are not yet registered for the castle siege.
	NotRegisteredForSiege = SysMsg{Id: 644}

	// Only clans of level 5 or higher may register for a castle siege.
	OnlyClanLevel5AboveMaySiege = SysMsg{Id: 645}

	// You do not have the authority to modify the castle defender list.
	DoNotHaveAuthorityToModifyCastleDefenderList = SysMsg{Id: 646}

	// You do not have the authority to modify the siege time.
	DoNotHaveAuthorityToModifySiegeTime = SysMsg{Id: 647}

	// No more registrations may be accepted for the attacker side.
	AttackerSideFull = SysMsg{Id: 648}

	// No more registrations may be accepted for the defender side.
	DefenderSideFull = SysMsg{Id: 649}

	// You may not summon from your current location.
	YouMayNotSummonFromYourCurrentLocation = SysMsg{Id: 650}

	// Place in the current location and direction. Do you wish to continue?
	PlaceCurrentLocationDirection = SysMsg{Id: 651}

	// The target of the summoned monster is wrong.
	TargetOfSummonWrong = SysMsg{Id: 652}

	// You do not have the authority to position mercenaries.
	YouDoNotHaveAuthorityToPositionMercenaries = SysMsg{Id: 653}

	// You do not have the authority to cancel mercenary positioning.
	YouDoNotHaveAuthorityToCancelMercenaryPositioning = SysMsg{Id: 654}

	// Mercenaries cannot be positioned here.
	MercenariesCannotBePositionedHere = SysMsg{Id: 655}

	// This mercenary cannot be positioned anymore.
	ThisMercenaryCannotBePositionedAnymore = SysMsg{Id: 656}

	// Positioning cannot be done here because the distance between mercenaries is too short.
	PositioningCannotBeDoneBecauseDistanceBetweenMercenariesTooShort = SysMsg{Id: 657}

	// This is not a mercenary of a castle that you own and so you cannot cancel its positioning.
	ThisIsNotAMercenaryOfACastleThatYouOwnAndSoCannotCancelPositioning = SysMsg{Id: 658}

	// This is not the time for siege registration and so registrations cannot be accepted or rejected.
	NotSiegeRegistrationTime1 = SysMsg{Id: 659}

	// This is not the time for siege registration and so registration and cancellation cannot be done.
	NotSiegeRegistrationTime2 = SysMsg{Id: 660}

	// This character cannot be spoiled.
	SpoilCannotUse = SysMsg{Id: 661}

	// The other player is rejecting friend invitations.
	ThePlayerIsRejectingFriendInvitations = SysMsg{Id: 662}

	// The siege time has been declared for $s. It is not possible to change the time after a siege time has been declared. Do you want to continue?
	SiegeTimeDeclaredForS1 = SysMsg{Id: 663}

	// Please choose a person to receive.
	ChoosePersonToReceive = SysMsg{Id: 664}

	// of alliance is applying for alliance war. Do you want to accept the challenge?
	ApplyingAllianceWar = SysMsg{Id: 665}

	// A request for ceasefire has been received from alliance. Do you agree?
	RequestForCeasefire = SysMsg{Id: 666}

	// You are registering on the attacking side of the siege. Do you want to continue?
	RegisteringOnAttackingSide = SysMsg{Id: 667}

	// You are registering on the defending side of the siege. Do you want to continue?
	RegisteringOnDefendingSide = SysMsg{Id: 668}

	// You are canceling your application to participate in the siege battle. Do you want to continue?
	CancelingRegistration = SysMsg{Id: 669}

	// You are refusing the registration of clan on the defending side. Do you want to continue?
	RefusingRegistration = SysMsg{Id: 670}

	// You are agreeing to the registration of clan on the defending side. Do you want to continue?
	AgreeingRegistration = SysMsg{Id: 671}

	// {{S1}} adena disappeared.
	S1DisappearedAdena = SysMsg{Id: 672}

	// Only a clan leader whose clan is of level 2 or higher is allowed to participate in a clan hall auction.
	AuctionOnlyClanLevel2Higher = SysMsg{Id: 673}

	// I has not yet been seven days since canceling an auction.
	NotSevenDaysSinceCancelingAuction = SysMsg{Id: 674}

	// There are no clan halls up for auction.
	NoClanHallsUpForAuction = SysMsg{Id: 675}

	// Since you have already submitted a bid, you are not allowed to participate in another auction at this time.
	AlreadySubmittedBid = SysMsg{Id: 676}

	// Your bid price must be higher than the minimum price that can be bid.
	BidPriceMustBeHigher = SysMsg{Id: 677}

	// You have submitted a bid for the auction of {{S1}}.
	SubmittedABidOfS1 = SysMsg{Id: 678}

	// You have canceled your bid.
	CanceledBid = SysMsg{Id: 679}

	// You cannot participate in an auction.
	CannotParticipateInAnAuction = SysMsg{Id: 680}

	// The clan does not own a clan hall.
	ClanHasNoClanHall = SysMsg{Id: 681}

	// The clan does not own a clan hall.
	MovingToAnotherVillage = SysMsg{Id: 681}

	// There are no priority rights on a sweeper.
	SweepNotAllowed = SysMsg{Id: 683}

	// You cannot position mercenaries during a siege.
	CannotPositionMercsDuringSiege = SysMsg{Id: 684}

	// You cannot apply for clan war with a clan that belongs to the same alliance
	CannotDeclareWarOnAlly = SysMsg{Id: 685}

	// You have received {{S1}} damage from the fire of magic.
	S1DamageFromFireMagic = SysMsg{Id: 686}

	// You cannot move while frozen. Please wait.
	CannotMoveFrozen = SysMsg{Id: 687}

	// The clan that owns the castle is automatically registered on the defending side.
	ClanThatOwnsCastleIsAutomaticallyRegisteredDefending = SysMsg{Id: 688}

	// A clan that owns a castle cannot participate in another siege.
	ClanThatOwnsCastleCannotParticipateOtherSiege = SysMsg{Id: 689}

	// You cannot register on the attacking side because you are part of an alliance with the clan that owns the castle.
	CannotAttackAllianceCastle = SysMsg{Id: 690}

	// {{S1}} clan is already a member of {{S2}} alliance.
	S1ClanAlreadyMemberOfS2Alliance = SysMsg{Id: 691}

	// The other party is frozen. Please wait a moment.
	OtherPartyIsFrozen = SysMsg{Id: 692}

	// The package that arrived is in another warehouse.
	PackageInAnotherWarehouse = SysMsg{Id: 693}

	// No packages have arrived.
	NoPackagesArrived = SysMsg{Id: 694}

	// You cannot set the name of the pet.
	NamingYouCannotSetNameOfThePet = SysMsg{Id: 695}

	// The item enchant value is strange
	ItemEnchantValueStrange = SysMsg{Id: 697}

	// The price is different than the same item on the sales list.
	PriceDifferentFromSalesList = SysMsg{Id: 698}

	// Currently not purchasing.
	CurrentlyNotPurchasing = SysMsg{Id: 699}

	// The purchase is complete.
	ThePurchaseIsComplete = SysMsg{Id: 700}

	// You do not have enough required items.
	NotEnoughRequiredItems = SysMsg{Id: 701}

	NoGmProvidingServiceNow = SysMsg{Id: 702}

	// ======<GM List>======
	GmList = SysMsg{Id: 703}

	// GM : {{C1}}
	GmC1 = SysMsg{Id: 704}

	// You cannot exclude yourself.
	CannotExcludeSelf = SysMsg{Id: 705}

	// You can only register up to 64 names on your exclude list.
	Only64NamesOnExcludeList = SysMsg{Id: 706}

	// You cannot teleport to a village that is in a siege.
	NoPortThatIsInSige = SysMsg{Id: 707}

	// You do not have the right to use the castle warehouse.
	YouDoNotHaveTheRightToUseCastleWarehouse = SysMsg{Id: 708}

	// You do not have the right to use the clan warehouse.
	YouDoNotHaveTheRightToUseClanWarehouse = SysMsg{Id: 709}

	// Only clans of clan level 1 or higher can use a clan warehouse.
	OnlyLevel1ClanOrHigherCanUseWarehouse = SysMsg{Id: 710}

	// The siege of {{S1}} has started.
	SiegeOfS1HasStarted = SysMsg{Id: 711}

	// The siege of {{S1}} has finished.
	SiegeOfS1HasEnded = SysMsg{Id: 712}

	// {{S1}}/{{S2}}/{{S3}} :
	S1S2S3D = SysMsg{Id: 713}

	// A trap device has been tripped.
	ATrapDeviceHasBeenTripped = SysMsg{Id: 714}

	// A trap device has been stopped.
	ATrapDeviceHasBeenStopped = SysMsg{Id: 715}

	// If a base camp does not exist, resurrection is not possible.
	NoResurrectionWithoutBaseCamp = SysMsg{Id: 716}

	// The guardian tower has been destroyed and resurrection is not possible
	TowerDestroyedNoResurrection = SysMsg{Id: 717}

	// The castle gates cannot be opened and closed during a siege.
	GatesNotOpenedClosedDuringSiege = SysMsg{Id: 718}

	// You failed at mixing the item.
	ItemMixingFailed = SysMsg{Id: 719}

	// The purchase price is higher than the amount of money that you have and so you cannot open a personal store.
	ThePurchasePriceIsHigherThanMoney = SysMsg{Id: 720}

	// You cannot create an alliance while participating in a siege.
	NoAllyCreationWhileSiege = SysMsg{Id: 721}

	// You cannot dissolve an alliance while an affiliated clan is participating in a siege battle.
	CannotDissolveAllyWhileInSiege = SysMsg{Id: 722}

	// The opposing clan is participating in a siege battle.
	OpposingClanIsParticipatingInSiege = SysMsg{Id: 723}

	// You cannot leave while participating in a siege battle.
	CannotLeaveWhileSiege = SysMsg{Id: 724}

	// You cannot banish a clan from an alliance while the clan is participating in a siege
	CannotDismissWhileSiege = SysMsg{Id: 725}

	// Frozen condition has started. Please wait a moment.
	FrozenConditionStarted = SysMsg{Id: 726}

	// The frozen condition was removed.
	FrozenConditionRemoved = SysMsg{Id: 727}

	// You cannot apply for dissolution again within seven days after a previous application for dissolution.
	CannotApplyDissolutionAgain = SysMsg{Id: 728}

	// That item cannot be discarded.
	ItemNotDiscarded = SysMsg{Id: 729}

	// You have submitted {{S1}} petition(s). - You may submit {{S2}} more petition(s) today.
	SubmittedYouS1ThPetitionS2Left = SysMsg{Id: 730}

	// A petition has been received by the GM on behalf of {{S1}}. The petition code is {{S2}}.
	PetitionS1ReceivedCodeIsS2 = SysMsg{Id: 731}

	// {{C1}} has received a request for a consultation with the GM.
	C1ReceivedConsultationRequest = SysMsg{Id: 732}

	// We have received {{S1}} petitions from you today and that is the maximum that you can submit in one day. You cannot submit any more petitions.
	WeHaveReceivedS1PetitionsToday = SysMsg{Id: 733}

	// You have failed at submitting a petition on behalf of someone else. {{C1}} already submitted a petition.
	PetitionFailedC1AlreadySubmitted = SysMsg{Id: 734}

	// You have failed at submitting a petition on behalf of {{C1}}. The error number is {{S2}}.
	PetitionFailedForC1ErrorNumberS2 = SysMsg{Id: 735}

	// The petition was canceled. You may submit {{S1}} more petition(s) today.
	PetitionCanceledSubmitS1MoreToday = SysMsg{Id: 736}

	// You have cancelled submitting a petition on behalf of {{S1}}.
	CanceledPetitionOnS1 = SysMsg{Id: 737}

	// You have not submitted a petition.
	PetitionNotSubmitted = SysMsg{Id: 738}

	// You have failed at cancelling a petition on behalf of {{C1}}. The error number is {{S2}}.
	PetitionCancelFailedForC1ErrorNumberS2 = SysMsg{Id: 739}

	// {{C1}} participated in a petition chat at the request of the GM.
	C1ParticipatePetition = SysMsg{Id: 740}

	// You have failed at adding {{C1}} to the petition chat. Petition has already been submitted.
	FailedAddingC1ToPetition = SysMsg{Id: 741}

	// You have failed at adding {{C1}} to the petition chat. The error code is {{S2}}.
	PetitionAddingC1FailedErrorNumberS2 = SysMsg{Id: 742}

	// {{C1}} left the petition chat.
	C1LeftPetitionChat = SysMsg{Id: 743}

	// You have failed at removing {{S1}} from the petition chat. The error code is {{S2}}.
	PetitionRemovingS1FailedErrorNumberS2 = SysMsg{Id: 744}

	// You are currently not in a petition chat.
	YouAreNotInPetitionChat = SysMsg{Id: 745}

	// It is not currently a petition.
	CurrentlyNoPetition = SysMsg{Id: 746}

	// The distance is too far and so the casting has been stopped.
	DistTooFarCastingStopped = SysMsg{Id: 748}

	// The effect of {{S1}} has been removed.
	EffectS1Disappeared = SysMsg{Id: 749}

	// There are no other skills to learn.
	NoMoreSkillsToLearn = SysMsg{Id: 750}

	// As there is a conflict in the siege relationship with a clan in the alliance, you cannot invite that clan to the alliance.
	CannotInviteConflictClan = SysMsg{Id: 751}

	// That name cannot be used.
	CannotUseName = SysMsg{Id: 752}

	// You cannot position mercenaries here.
	NoMercsHere = SysMsg{Id: 753}

	// There are {{S1}} hours and {{S2}} minutes left in this week's usage time.
	S1HoursS2MinutesLeftThisWeek = SysMsg{Id: 754}

	// There are {{S1}} minutes left in this week's usage time.
	S1MinutesLeftThisWeek = SysMsg{Id: 755}

	// This week's usage time has finished.
	WeeksUsageTimeFinished = SysMsg{Id: 756}

	// There are {{S1}} hours and {{S2}} minutes left in the fixed use time.
	S1HoursS2MinutesLeftInTime = SysMsg{Id: 757}

	// There are {{S1}} hours and {{S2}} minutes left in this week's play time.
	S1HoursS2MinutesLeftThisWeeksPlayTime = SysMsg{Id: 758}

	// There are {{S1}} minutes left in this week's play time.
	S1MinutesLeftThisWeeksPlayTime = SysMsg{Id: 759}

	// {{C1}} cannot join the clan because one day has not yet passed since he/she left another clan.
	C1MustWaitBeforeJoiningAnotherClan = SysMsg{Id: 760}

	// {{S1}} clan cannot join the alliance because one day has not yet passed since it left another alliance.
	S1CantEnterAllianceWithin1Day = SysMsg{Id: 761}

	// {{C1}} rolled {{S2}} and {{S3}}'s eye came out.
	C1RolledS2S3EyeCameOut = SysMsg{Id: 762}

	// You failed at sending the package because you are too far from the warehouse.
	FailedSendingPackageTooFar = SysMsg{Id: 763}

	// You have been playing for an extended period of time. Please consider taking a break.
	PlayingForLongTime = SysMsg{Id: 764}

	// A hacking tool has been discovered. Please try again after closing unnecessary programs.
	HackingTool = SysMsg{Id: 769}

	// Play time is no longer accumulating.
	PlayTimeNoLongerAccumulating = SysMsg{Id: 774}

	// From here on, play time will be expended.
	PlayTimeExpended = SysMsg{Id: 775}

	// The clan hall which was put up for auction has been awarded to clan.
	ClanhallAwardedToClan = SysMsg{Id: 776}

	// The clan hall which was put up for auction was not sold and therefore has been re-listed.
	ClanhallNotSold = SysMsg{Id: 777}

	// You may not log out from this location.
	NoLogoutHere = SysMsg{Id: 778}

	// You may not restart in this location.
	NoRestartHere = SysMsg{Id: 779}

	// Observation is only possible during a siege.
	OnlyViewSiege = SysMsg{Id: 780}

	// Observers cannot participate.
	ObserversCannotParticipate = SysMsg{Id: 781}

	// You may not observe a siege with a pet or servitor summoned.
	NoObserveWithPet = SysMsg{Id: 782}

	// Lottery ticket sales have been temporarily suspended.
	LotteryTicketSalesTempSuspended = SysMsg{Id: 783}

	// Tickets for the current lottery are no longer available.
	NoLotteryTicketsAvailable = SysMsg{Id: 784}

	// The results of lottery number {{S1}} have not yet been published.
	LotteryS1ResultNotPublished = SysMsg{Id: 785}

	// Incorrect syntax.
	IncorrectSyntax = SysMsg{Id: 786}

	// The tryouts are finished.
	ClanhallSiegeTryoutsFinished = SysMsg{Id: 787}

	// The finals are finished.
	ClanhallSiegeFinalsFinished = SysMsg{Id: 788}

	// The tryouts have begun.
	ClanhallSiegeTryoutsBegun = SysMsg{Id: 789}

	// The finals are finished.
	ClanhallSiegeFinalsBegun = SysMsg{Id: 790}

	// The final match is about to begin. Line up!
	FinalMatchBegin = SysMsg{Id: 791}

	// The siege of the clan hall is finished.
	ClanhallSiegeEnded = SysMsg{Id: 792}

	// The siege of the clan hall has begun.
	ClanhallSiegeBegun = SysMsg{Id: 793}

	// You are not authorized to do that.
	YouAreNotAuthorizedToDoThat = SysMsg{Id: 794}

	// Only clan leaders are authorized to set rights.
	OnlyLeadersCanSetRights = SysMsg{Id: 795}

	// Your remaining observation time is minutes.
	RemainingObservationTime = SysMsg{Id: 796}

	// You may create up to 48 macros.
	YouMayCreateUpTo48Macros = SysMsg{Id: 797}

	// Item registration is irreversible. Do you wish to continue?
	ItemRegistrationIrreversible = SysMsg{Id: 798}

	// The observation time has expired.
	ObservationTimeExpired = SysMsg{Id: 799}

	// You are too late. The registration period is over.
	RegistrationPeriodOver = SysMsg{Id: 800}

	// Registration for the clan hall siege is closed.
	RegistrationClosed = SysMsg{Id: 801}

	// Petitions are not being accepted at this time. You may submit your petition after a.m./p.m.
	PetitionNotAcceptedNow = SysMsg{Id: 802}

	// Enter the specifics of your petition.
	PetitionNotSpecified = SysMsg{Id: 803}

	// Select a type.
	SelectType = SysMsg{Id: 804}

	// Petitions are not being accepted at this time. You may submit your petition after {{S1}} a.m./p.m.
	PetitionNotAcceptedSubmitAtS1 = SysMsg{Id: 805}

	// If you are trapped, try typing "/unstuck".
	TryUnstuckWhenTrapped = SysMsg{Id: 806}

	// This terrain is navigable. Prepare for transport to the nearest village.
	StuckPrepareForTransport = SysMsg{Id: 807}

	// You are stuck. You may submit a petition by typing "/gm".
	StuckSubmitPetition = SysMsg{Id: 808}

	// You are stuck. You will be transported to the nearest village in five minutes.
	StuckTransportInFiveMinutes = SysMsg{Id: 809}

	// Invalid macro. Refer to the Help file for instructions.
	InvalidMacro = SysMsg{Id: 810}

	// You will be moved to (). Do you wish to continue?
	WillBeMoved = SysMsg{Id: 811}

	// The secret trap has inflicted {{S1}} damage on you.
	TrapDidS1Damage = SysMsg{Id: 812}

	// You have been poisoned by a Secret Trap.
	PoisonedByTrap = SysMsg{Id: 813}

	// Your speed has been decreased by a Secret Trap.
	SlowedByTrap = SysMsg{Id: 814}

	// The tryouts are about to begin. Line up!
	TryoutsAboutToBegin = SysMsg{Id: 815}

	// Tickets are now available for Monster Race {{S1}}!
	MonsraceTicketsAvailableForS1Race = SysMsg{Id: 816}

	// Now selling tickets for Monster Race {{S1}}!
	MonsraceTicketsNowAvailableForS1Race = SysMsg{Id: 817}

	// Ticket sales for the Monster Race will end in {{S1}} minute(s).
	MonsraceTicketsStopInS1Minutes = SysMsg{Id: 818}

	// Tickets sales are closed for Monster Race {{S1}}. Odds are posted.
	MonsraceS1TicketSalesClosed = SysMsg{Id: 819}

	// Monster Race {{S2}} will begin in {{S1}} minute(s)!
	MonsraceS2BeginsInS1Minutes = SysMsg{Id: 820}

	// Monster Race {{S1}} will begin in 30 seconds!
	MonsraceS1BeginsIn30Seconds = SysMsg{Id: 821}

	// Monster Race {{S1}} is about to begin! Countdown in five seconds!
	MonsraceS1CountdownInFiveSeconds = SysMsg{Id: 822}

	// The race will begin in {{S1}} second(s)!
	MonsraceBeginsInS1Seconds = SysMsg{Id: 823}

	// They're off!
	MonsraceRaceStart = SysMsg{Id: 824}

	// Monster Race {{S1}} is finished!
	MonsraceS1RaceEnd = SysMsg{Id: 825}

	// First prize goes to the player in lane {{S1}}. Second prize goes to the player in lane {{S2}}.
	MonsraceFirstPlaceS1SecondS2 = SysMsg{Id: 826}

	// You may not impose a block on a GM.
	YouMayNotImposeABlockOnGm = SysMsg{Id: 827}

	// Are you sure you wish to delete the {{S1}} macro?
	WishToDeleteS1Macro = SysMsg{Id: 828}

	// You cannot recommend yourself.
	YouCannotRecommendYourself = SysMsg{Id: 829}

	// You have recommended {{C1}}. You have {{S2}} recommendations left.
	YouHaveRecommendedC1YouHaveS2RecommendationsLeft = SysMsg{Id: 830}

	// You have been recommended by {{C1}}.
	YouHaveBeenRecommendedByC1 = SysMsg{Id: 831}

	// That character has already been recommended.
	ThatCharacterIsRecommended = SysMsg{Id: 832}

	// You are not authorized to make further recommendations at this time. You will receive more recommendation credits each day at 1 p.m.
	NoMoreRecommendationsToHave = SysMsg{Id: 833}

	// {{C1}} has rolled {{S2}}.
	C1RolledS2 = SysMsg{Id: 834}

	// You may not throw the dice at this time. Try again later.
	YouMayNotThrowTheDiceAtThisTimeTryAgainLater = SysMsg{Id: 835}

	// You have exceeded your inventory volume limit and cannot take this item.
	YouHaveExceededYourInventoryVolumeLimitAndCannotTakeThisItem = SysMsg{Id: 836}

	// Macro descriptions may contain up to 32 characters.
	MacroDescriptionMax32Chars = SysMsg{Id: 837}

	// Enter the name of the macro.
	EnterTheMacroName = SysMsg{Id: 838}

	// That name is already assigned to another macro.
	MacroNameAlreadyUsed = SysMsg{Id: 839}

	// That recipe is already registered.
	RecipeAlreadyRegistered = SysMsg{Id: 840}

	// No further recipes may be registered.
	NoFutherRecipesCanBeAdded = SysMsg{Id: 841}

	// You are not authorized to register a recipe.
	NotAuthorizedRegisterRecipe = SysMsg{Id: 842}

	// The siege of {{S1}} is finished.
	SiegeOfS1Finished = SysMsg{Id: 843}

	// The siege to conquer {{S1}} has begun.
	SiegeOfS1Begun = SysMsg{Id: 844}

	// The deadlineto register for the siege of {{S1}} has passed.
	DeadlineForSiegeS1Passed = SysMsg{Id: 845}

	// The siege of {{S1}} has been canceled due to lack of interest.
	SiegeOfS1HasBeenCanceledDueToLackOfInterest = SysMsg{Id: 846}

	// A clan that owns a clan hall may not participate in a clan hall siege.
	ClanOwningClanhallMayNotSiegeClanhall = SysMsg{Id: 847}

	// {{S1}} has been deleted.
	S1HasBeenDeleted = SysMsg{Id: 848}

	// {{S1}} cannot be found.
	S1NotFound = SysMsg{Id: 849}

	// {{S1}} already exists.
	S1AlreadyExists2 = SysMsg{Id: 850}

	// {{S1}} has been added.
	S1Added = SysMsg{Id: 851}

	// The recipe is incorrect.
	RecipeIncorrect = SysMsg{Id: 852}

	// You may not alter your recipe book while engaged in manufacturing.
	CantAlterRecipebookWhileCrafting = SysMsg{Id: 853}

	// You are missing {{S2}} {{S1}} required to create that.
	MissingS2S1ToCreate = SysMsg{Id: 854}

	// {{S1}} clan has defeated {{S2}}.
	S1ClanDefeatedS2 = SysMsg{Id: 855}

	// The siege of {{S1}} has ended in a draw.
	SiegeS1Draw = SysMsg{Id: 856}

	// {{S1}} clan has won in the preliminary match of {{S2}}.
	S1ClanWonMatchS2 = SysMsg{Id: 857}

	// The preliminary match of {{S1}} has ended in a draw.
	MatchOfS1Draw = SysMsg{Id: 858}

	// Please register a recipe.
	PleaseRegisterRecipe = SysMsg{Id: 859}

	// You may not buld your headquarters in close proximity to another headquarters.
	HeadquartersTooClose = SysMsg{Id: 860}

	// You have exceeded the maximum number of memos.
	TooManyMemos = SysMsg{Id: 861}

	// Odds are not posted until ticket sales have closed.
	OddsNotPosted = SysMsg{Id: 862}

	// You feel the energy of fire.
	FeelEnergyFire = SysMsg{Id: 863}

	// You feel the energy of water.
	FeelEnergyWater = SysMsg{Id: 864}

	// You feel the energy of wind.
	FeelEnergyWind = SysMsg{Id: 865}

	// You may no longer gather energy.
	NoLongerEnergy = SysMsg{Id: 866}

	// The energy is depleted.
	EnergyDepleted = SysMsg{Id: 867}

	// The energy of fire has been delivered.
	EnergyFireDelivered = SysMsg{Id: 868}

	// The energy of water has been delivered.
	EnergyWaterDelivered = SysMsg{Id: 869}

	// The energy of wind has been delivered.
	EnergyWindDelivered = SysMsg{Id: 870}

	// The seed has been sown.
	TheSeedHasBeenSown = SysMsg{Id: 871}

	// This seed may not be sown here.
	ThisSeedMayNotBeSownHere = SysMsg{Id: 872}

	// That character does not exist.
	CharacterDoesNotExist = SysMsg{Id: 873}

	// The capacity of the warehouse has been exceeded.
	WarehouseCapacityExceeded = SysMsg{Id: 874}

	// The transport of the cargo has been canceled.
	CargoCanceled = SysMsg{Id: 875}

	// The cargo was not delivered.
	CargoNotDelivered = SysMsg{Id: 876}

	// The symbol has been added.
	SymbolAdded = SysMsg{Id: 877}

	// The symbol has been deleted.
	SymbolDeleted = SysMsg{Id: 878}

	// The manor system is currently under maintenance.
	TheManorSystemIsCurrentlyUnderMaintenance = SysMsg{Id: 879}

	// The transaction is complete.
	TheTransactionIsComplete = SysMsg{Id: 880}

	// There is a discrepancy on the invoice.
	ThereIsADiscrepancyOnTheInvoice = SysMsg{Id: 881}

	// The seed quantity is incorrect.
	TheSeedQuantityIsIncorrect = SysMsg{Id: 882}

	// The seed information is incorrect.
	TheSeedInformationIsIncorrect = SysMsg{Id: 883}

	// The manor information has been updated.
	TheManorInformationHasBeenUpdated = SysMsg{Id: 884}

	// The number of crops is incorrect.
	TheNumberOfCropsIsIncorrect = SysMsg{Id: 885}

	// The crops are priced incorrectly.
	TheCropsArePricedIncorrectly = SysMsg{Id: 886}

	// The type is incorrect.
	TheTypeIsIncorrect = SysMsg{Id: 887}

	// No crops can be purchased at this time.
	NoCropsCanBePurchasedAtThisTime = SysMsg{Id: 888}

	// The seed was successfully sown.
	TheSeedWasSuccessfullySown = SysMsg{Id: 889}

	// The seed was not sown.
	TheSeedWasNotSown = SysMsg{Id: 890}

	// You are not authorized to harvest.
	YouAreNotAuthorizedToHarvest = SysMsg{Id: 891}

	// The harvest has failed.
	TheHarvestHasFailed = SysMsg{Id: 892}

	// The harvest failed because the seed was not sown.
	TheHarvestFailedBecauseTheSeedWasNotSown = SysMsg{Id: 893}

	// Up to {{S1}} recipes can be registered.
	UpToS1RecipesCanRegister = SysMsg{Id: 894}

	// No recipes have been registered.
	NoRecipesRegistered = SysMsg{Id: 895}

	// The ferry has arrived at Gludin Harbor.
	FerryAtGludin = SysMsg{Id: 896}

	// The ferry will leave for Talking Island Harbor after anchoring for ten minutes.
	FerryLeaveTalking = SysMsg{Id: 897}

	// Only characters of level 10 or above are authorized to make recommendations.
	OnlyLevelSup10CanRecommend = SysMsg{Id: 898}

	// The symbol cannot be drawn.
	CantDrawSymbol = SysMsg{Id: 899}

	// No slot exists to draw the symbol
	SymbolsFull = SysMsg{Id: 900}

	// The symbol information cannot be found.
	SymbolNotFound = SysMsg{Id: 901}

	// The number of items is incorrect.
	NumberIncorrect = SysMsg{Id: 902}

	// You may not submit a petition while frozen. Be patient.
	NoPetitionWhileFrozen = SysMsg{Id: 903}

	// Items cannot be discarded while in private store status.
	NoDiscardWhilePrivateStore = SysMsg{Id: 904}

	// The current score for the Humans is {{S1}}.
	HumanScoreS1 = SysMsg{Id: 905}

	// The current score for the Elves is {{S1}}.
	ElvesScoreS1 = SysMsg{Id: 906}

	// The current score for the Dark Elves is {{S1}}.
	DarkElvesScoreS1 = SysMsg{Id: 907}

	// The current score for the Orcs is {{S1}}.
	OrcsScoreS1 = SysMsg{Id: 908}

	// The current score for the Dwarves is {{S1}}.
	DwarvenScoreS1 = SysMsg{Id: 909}

	// Current location : {{S1}}, {{S2}}, {{S3}} (Near Talking Island Village)
	LocTiS1S2S3 = SysMsg{Id: 910}

	// Current location : {{S1}}, {{S2}}, {{S3}} (Near Gludin Village)
	LocGludinS1S2S3 = SysMsg{Id: 911}

	// Current location : {{S1}}, {{S2}}, {{S3}} (Near the Town of Gludio)
	LocGludioS1S2S3 = SysMsg{Id: 912}

	// Current location : {{S1}}, {{S2}}, {{S3}} (Near the Neutral Zone)
	LocNeutralZoneS1S2S3 = SysMsg{Id: 913}

	// Current location : {{S1}}, {{S2}}, {{S3}} (Near the Elven Village)
	LocElvenS1S2S3 = SysMsg{Id: 914}

	// Current location : {{S1}}, {{S2}}, {{S3}} (Near the Dark Elf Village)
	LocDarkElvenS1S2S3 = SysMsg{Id: 915}

	// Current location : {{S1}}, {{S2}}, {{S3}} (Near the Town of Dion)
	LocDionS1S2S3 = SysMsg{Id: 916}

	// Current location : {{S1}}, {{S2}}, {{S3}} (Near the Floran Village)
	LocFloranS1S2S3 = SysMsg{Id: 917}

	// Current location : {{S1}}, {{S2}}, {{S3}} (Near the Town of Giran)
	LocGiranS1S2S3 = SysMsg{Id: 918}

	// Current location : {{S1}}, {{S2}}, {{S3}} (Near Giran Harbor)
	LocGiranHarborS1S2S3 = SysMsg{Id: 919}

	// Current location : {{S1}}, {{S2}}, {{S3}} (Near the Orc Village)
	LocOrcS1S2S3 = SysMsg{Id: 920}

	// Current location : {{S1}}, {{S2}}, {{S3}} (Near the Dwarven Village)
	LocDwarvenS1S2S3 = SysMsg{Id: 921}

	// Current location : {{S1}}, {{S2}}, {{S3}} (Near the Town of Oren)
	LocOrenS1S2S3 = SysMsg{Id: 922}

	// Current location : {{S1}}, {{S2}}, {{S3}} (Near Hunters Village)
	LocHunterS1S2S3 = SysMsg{Id: 923}

	// Current location : {{S1}}, {{S2}}, {{S3}} (Near Aden Castle Town)
	LocAdenS1S2S3 = SysMsg{Id: 924}

	// Current location : {{S1}}, {{S2}}, {{S3}} (Near the Coliseum)
	LocColiseumS1S2S3 = SysMsg{Id: 925}

	// Current location : {{S1}}, {{S2}}, {{S3}} (Near Heine)
	LocHeineS1S2S3 = SysMsg{Id: 926}

	// The current time is {{S1}}:{{S2}}.
	TimeS1S2InTheDay = SysMsg{Id: 927}

	// The current time is {{S1}}:{{S2}}.
	TimeS1S2InTheNight = SysMsg{Id: 928}

	// No compensation was given for the farm products.
	NoCompensationForFarmProducts = SysMsg{Id: 929}

	// Lottery tickets are not currently being sold.
	NoLotteryTicketsCurrentSold = SysMsg{Id: 930}

	// The winning lottery ticket numbers has not yet been anonunced.
	LotteryWinnersNotAnnouncedYet = SysMsg{Id: 931}

	// You cannot chat locally while observing.
	NoAllchatWhileObserving = SysMsg{Id: 932}

	// The seed pricing greatly differs from standard seed prices.
	TheSeedPricingGreatlyDiffersFromStandardSeedPrices = SysMsg{Id: 933}

	// It is a deleted recipe.
	ADeletedRecipe = SysMsg{Id: 934}

	// The amount is not sufficient and so the manor is not in operation.
	TheAmountIsNotSufficientAndSoTheManorIsNotInOperation = SysMsg{Id: 935}

	// Use {{S1}}.
	UseS1_ = SysMsg{Id: 936}

	// Currently preparing for private workshop.
	PreparingPrivateWorkshop = SysMsg{Id: 937}

	// The community server is currently offline.
	CbOffline = SysMsg{Id: 938}

	// You cannot exchange while blocking everything.
	NoExchangeWhileBlocking = SysMsg{Id: 939}

	// {{S1}} is blocked everything.
	S1BlockedEverything = SysMsg{Id: 940}

	// Restart at Talking Island Village.
	RestartAtTi = SysMsg{Id: 941}

	// Restart at Gludin Village.
	RestartAtGludin = SysMsg{Id: 942}

	// Restart at the Town of Gludin. || guess should be Gludio ;)
	RestartAtGludio = SysMsg{Id: 943}

	// Restart at the Neutral Zone.
	RestartAtNeutralZone = SysMsg{Id: 944}

	// Restart at the Elven Village.
	RestartAtElfenVillage = SysMsg{Id: 945}

	// Restart at the Dark Elf Village.
	RestartAtDarkelfVillage = SysMsg{Id: 946}

	// Restart at the Town of Dion.
	RestartAtDion = SysMsg{Id: 947}

	// Restart at Floran Village.
	RestartAtFloran = SysMsg{Id: 948}

	// Restart at the Town of Giran.
	RestartAtGiran = SysMsg{Id: 949}

	// Restart at Giran Harbor.
	RestartAtGiranHarbor = SysMsg{Id: 950}

	// Restart at the Orc Village.
	RestartAtOrcVillage = SysMsg{Id: 951}

	// Restart at the Dwarven Village.
	RestartAtDwarfenVillage = SysMsg{Id: 952}

	// Restart at the Town of Oren.
	RestartAtOren = SysMsg{Id: 953}

	// Restart at Hunters Village.
	RestartAtHuntersVillage = SysMsg{Id: 954}

	// Restart at the Town of Aden.
	RestartAtAden = SysMsg{Id: 955}

	// Restart at the Coliseum.
	RestartAtColiseum = SysMsg{Id: 956}

	// Restart at Heine.
	RestartAtHeine = SysMsg{Id: 957}

	// Items cannot be discarded or destroyed while operating a private store or workshop.
	ItemsCannotBeDiscardedOrDestroyedWhileOperatingPrivateStoreOrWorkshop = SysMsg{Id: 958}

	// {{S1}} (*{{S2}}) manufactured successfully.
	S1S2ManufacturedSuccessfully = SysMsg{Id: 959}

	// {{S1}} manufacturing failure.
	S1ManufactureFailure = SysMsg{Id: 960}

	// You are now blocking everything.
	BlockingAll = SysMsg{Id: 961}

	// You are no longer blocking everything.
	NotBlockingAll = SysMsg{Id: 962}

	// Please determine the manufacturing price.
	DetermineManufacturePrice = SysMsg{Id: 963}

	// Chatting is prohibited for one minute.
	ChatbanFor1Minute = SysMsg{Id: 964}

	// The chatting prohibition has been removed.
	ChatbanRemoved = SysMsg{Id: 965}

	// Chatting is currently prohibited. If you try to chat before the prohibition is removed, the prohibition time will become even longer.
	ChattingIsCurrentlyProhibited = SysMsg{Id: 966}

	// Do you accept {{C1}}'s party invitation? (Item Distribution: Random including spoil.)
	C1PartyInviteRandomIncludingSpoil = SysMsg{Id: 967}

	// Do you accept {{C1}}'s party invitation? (Item Distribution: By Turn.)
	C1PartyInviteByTurn = SysMsg{Id: 968}

	// Do you accept {{C1}}'s party invitation? (Item Distribution: By Turn including spoil.)
	C1PartyInviteByTurnIncludingSpoil = SysMsg{Id: 969}

	// {{S2}}'s MP has been drained by {{C1}}.
	S2MpHasBeenDrainedByC1 = SysMsg{Id: 970}

	// Petitions cannot exceed 255 characters.
	PetitionMaxChars255 = SysMsg{Id: 971}

	// This pet cannot use this item.
	PetCannotUseItem = SysMsg{Id: 972}

	// Please input no more than the number you have.
	InputNoMoreYouHave = SysMsg{Id: 973}

	// The soul crystal succeeded in absorbing a soul.
	SoulCrystalAbsorbingSucceeded = SysMsg{Id: 974}

	// The soul crystal was not able to absorb a soul.
	SoulCrystalAbsorbingFailed = SysMsg{Id: 975}

	// The soul crystal broke because it was not able to endure the soul energy.
	SoulCrystalBroke = SysMsg{Id: 976}

	// The soul crystals caused resonation and failed at absorbing a soul.
	SoulCrystalAbsorbingFailedResonation = SysMsg{Id: 977}

	// The soul crystal is refusing to absorb a soul.
	SoulCrystalAbsorbingRefused = SysMsg{Id: 978}

	// The ferry arrived at Talking Island Harbor.
	FerryArrivedAtTalking = SysMsg{Id: 979}

	// The ferry will leave for Gludin Harbor after anchoring for ten minutes.
	FerryLeaveForGludinAfter10Minutes = SysMsg{Id: 980}

	// The ferry will leave for Gludin Harbor in five minutes.
	FerryLeaveForGludinIn5Minutes = SysMsg{Id: 981}

	// The ferry will leave for Gludin Harbor in one minute.
	FerryLeaveForGludinIn1Minute = SysMsg{Id: 982}

	// Those wishing to ride should make haste to get on.
	MakeHasteGetOnBoat = SysMsg{Id: 983}

	// The ferry will be leaving soon for Gludin Harbor.
	FerryLeaveSoonForGludin = SysMsg{Id: 984}

	// The ferry is leaving for Gludin Harbor.
	FerryLeavingForGludin = SysMsg{Id: 985}

	// The ferry has arrived at Gludin Harbor.
	FerryArrivedAtGludin = SysMsg{Id: 986}

	// The ferry will leave for Talking Island Harbor after anchoring for ten minutes.
	FerryLeaveForTalkingAfter10Minutes = SysMsg{Id: 987}

	// The ferry will leave for Talking Island Harbor in five minutes.
	FerryLeaveForTalkingIn5Minutes = SysMsg{Id: 988}

	// The ferry will leave for Talking Island Harbor in one minute.
	FerryLeaveForTalkingIn1Minute = SysMsg{Id: 989}

	// The ferry will be leaving soon for Talking Island Harbor.
	FerryLeaveSoonForTalking = SysMsg{Id: 990}

	// The ferry is leaving for Talking Island Harbor.
	FerryLeavingForTalking = SysMsg{Id: 991}

	// The ferry has arrived at Giran Harbor.
	FerryArrivedAtGiran = SysMsg{Id: 992}

	// The ferry will leave for Giran Harbor after anchoring for ten minutes.
	FerryLeaveForGiranAfter10Minutes = SysMsg{Id: 993}

	// The ferry will leave for Giran Harbor in five minutes.
	FerryLeaveForGiranIn5Minutes = SysMsg{Id: 994}

	// The ferry will leave for Giran Harbor in one minute.
	FerryLeaveForGiranIn1Minute = SysMsg{Id: 995}

	// The ferry will be leaving soon for Giran Harbor.
	FerryLeaveSoonForGiran = SysMsg{Id: 996}

	// The ferry is leaving for Giran Harbor.
	FerryLeavingForGiran = SysMsg{Id: 997}

	// The Innadril pleasure boat has arrived. It will anchor for ten minutes.
	InnadriLBoatAnchor10Minutes = SysMsg{Id: 998}

	// The Innadril pleasure boat will leave in five minutes.
	InnadriLBoatLeaveIn5Minutes = SysMsg{Id: 999}

	// The Innadril pleasure boat will leave in one minute.
	InnadriLBoatLeaveIn1Minute = SysMsg{Id: 1000}

	// The Innadril pleasure boat will be leaving soon.
	InnadriLBoatLeaveSoon = SysMsg{Id: 1001}

	// The Innadril pleasure boat is leaving.
	InnadriLBoatLeaving = SysMsg{Id: 1002}

	// Cannot possess a monster race ticket.
	CannotPossesMonsTicket = SysMsg{Id: 1003}

	// You have registered for a clan hall auction.
	RegisteredForClanhall = SysMsg{Id: 1004}

	// There is not enough adena in the clan hall warehouse.
	NotEnoughAdenaInCwh = SysMsg{Id: 1005}

	// You have bid in a clan hall auction.
	BidInClanhallAuction = SysMsg{Id: 1006}

	// The preliminary match registration of {{S1}} has finished.
	PreliminaryRegistrationOfS1Finished = SysMsg{Id: 1007}

	// A hungry strider cannot be mounted or dismounted.
	HungryStriderNotMount = SysMsg{Id: 1008}

	// A strider cannot be ridden when dead.
	StriderCantBeRiddenWhileDead = SysMsg{Id: 1009}

	// A dead strider cannot be ridden.
	DeadStriderCantBeRidden = SysMsg{Id: 1010}

	// A strider in battle cannot be ridden.
	StriderInBatlleCantBeRidden = SysMsg{Id: 1011}

	// A strider cannot be ridden while in battle.
	StriderCantBeRiddenWhileInBattle = SysMsg{Id: 1012}

	// A strider can be ridden only when standing.
	StriderCanBeRiddenOnlyWhileStanding = SysMsg{Id: 1013}

	// Your pet gained {{S1}} experience points.
	PetEarnedS1Exp = SysMsg{Id: 1014}

	// Your pet hit for {{S1}} damage.
	PetHitForS1Damage = SysMsg{Id: 1015}

	// Pet received {{S2}} damage by {{C1}}.
	PetReceivedS2DamageByC1 = SysMsg{Id: 1016}

	// Pet's critical hit!
	CriticalHitByPet = SysMsg{Id: 1017}

	// Your pet uses {{S1}}.
	PetUsesS1 = SysMsg{Id: 1018}

	// Your pet uses {{S1}}.
	PetUsesS1_ = SysMsg{Id: 1019}

	// Your pet picked up {{S1}}.
	PetPickedS1 = SysMsg{Id: 1020}

	// Your pet picked up {{S2}} {{S1}}(s).
	PetPickedS2S1S = SysMsg{Id: 1021}

	// Your pet picked up +{{S1}} {{S2}}.
	PetPickedS1S2 = SysMsg{Id: 1022}

	// Your pet picked up {{S1}} adena.
	PetPickedS1Adena = SysMsg{Id: 1023}

	// Your pet put on {{S1}}.
	PetPutOnS1 = SysMsg{Id: 1024}

	// Your pet took off {{S1}}.
	PetTookOffS1 = SysMsg{Id: 1025}

	// The summoned monster gave damage of {{S1}}
	SummonGaveDamageS1 = SysMsg{Id: 1026}

	// Servitor received {{S2}} damage caused by {{S1}}.
	SummonReceivedDamageS2ByS1 = SysMsg{Id: 1027}

	// Summoned monster's critical hit!
	CriticalHitBySummonedMob = SysMsg{Id: 1028}

	// Summoned monster uses {{S1}}.
	SummonedMobUsesS1 = SysMsg{Id: 1029}

	// <Party Information>
	PartyInformation = SysMsg{Id: 1030}

	// Looting method: Finders keepers
	LootingFindersKeepers = SysMsg{Id: 1031}

	// Looting method: Random
	LootingRandom = SysMsg{Id: 1032}

	// Looting method: Random including spoil
	LootingRandomIncludeSpoil = SysMsg{Id: 1033}

	// Looting method: By turn
	LootingByTurn = SysMsg{Id: 1034}

	// Looting method: By turn including spoil
	LootingByTurnIncludeSpoil = SysMsg{Id: 1035}

	// You have exceeded the quantity that can be inputted.
	YouHaveExceededQuantityThatCanBeInputted = SysMsg{Id: 1036}

	// {{C1}} manufactured {{S2}}.
	C1ManufacturedS2 = SysMsg{Id: 1037}

	// {{C1}} manufactured {{S3}} {{S2}}(s).
	C1ManufacturedS3S2S = SysMsg{Id: 1038}

	// Items left at the clan hall warehouse can only be retrieved by the clan leader. Do you want to continue?
	OnlyClanLeaderCanRetrieveItemsFromClanWarehouse = SysMsg{Id: 1039}

	// Items sent by freight can be picked up from any Warehouse location. Do you want to continue?
	ItemsSentByFreightPickedUpFromAnywhere = SysMsg{Id: 1040}

	// The next seed purchase price is {{S1}} adena.
	TheNextSeedPurchasePriceIsS1Adena = SysMsg{Id: 1041}

	// The next farm goods purchase price is {{S1}} adena.
	TheNextFarmGoodsPurchasePriceIsS1Adena = SysMsg{Id: 1042}

	// At the current time, the "/unstuck" command cannot be used. Please send in a petition.
	NoUnstuckPleaseSendPetition = SysMsg{Id: 1043}

	// Monster race payout information is not available while tickets are being sold.
	MonsraceNoPayoutInfo = SysMsg{Id: 1044}

	// Monster race tickets are no longer available.
	MonsraceTicketsNotAvailable = SysMsg{Id: 1046}

	// We did not succeed in producing {{S1}} item.
	NotSucceedProducingS1 = SysMsg{Id: 1047}

	// When "blocking" everything, whispering is not possible.
	NoWhisperWhenBlocking = SysMsg{Id: 1048}

	// When "blocking" everything, it is not possible to send invitations for organizing parties.
	NoPartyWhenBlocking = SysMsg{Id: 1049}

	// There are no communities in my clan. Clan communities are allowed for clans with skill levels of 2 and higher.
	NoCbInMyClan = SysMsg{Id: 1050}

	PaymentForYourClanHallHasNotBeenMadePleaseMakePaymentToYourClanWarehouseByS1Tomorrow = SysMsg{Id: 1051}

	TheClanHallFeeIsOneWeekOverdueThereforeTheClanHallOwnershipHasBeenRevoked = SysMsg{Id: 1052}

	// It is not possible to resurrect in battlefields where a siege war is taking place.
	CannotBeResurrectedDuringSiege = SysMsg{Id: 1053}

	// You have entered a mystical land.
	EnteredMysticalLand = SysMsg{Id: 1054}

	// You have left a mystical land.
	ExitedMysticalLand = SysMsg{Id: 1055}

	// You have exceeded the storage capacity of the castle's vault.
	VaultCapacityExceeded = SysMsg{Id: 1056}

	// This command can only be used in the relax server.
	RelaxServerOnly = SysMsg{Id: 1057}

	// The sales price for seeds is {{S1}} adena.
	TheSalesPriceForSeedsIsS1Adena = SysMsg{Id: 1058}

	// The remaining purchasing amount is {{S1}} adena.
	TheRemainingPurchasingIsS1Adena = SysMsg{Id: 1059}

	// The remainder after selling the seeds is {{S1}}.
	TheRemainderAfterSellingTheSeedsIsS1 = SysMsg{Id: 1060}

	// The recipe cannot be registered. You do not have the ability to create items.
	CantRegisterNoAbilityToCraft = SysMsg{Id: 1061}

	// Writing something new is possible after level 10.
	WritingSomethingNewPossibleAfterLevel10 = SysMsg{Id: 1062}

	PetitionUnavailable = SysMsg{Id: 1063}

	// The equipment, +{{S1}} {{S2}}, has been removed.
	EquipmentS1S2Removed = SysMsg{Id: 1064}

	// While operating a private store or workshop, you cannot discard, destroy, or trade an item.
	CannotTradeDiscardDropItemWhileInShopmode = SysMsg{Id: 1065}

	// {{S1}} HP has been restored.
	S1HpRestored = SysMsg{Id: 1066}

	// {{S2}} HP has been restored by {{C1}}
	S2HpRestoredByC1 = SysMsg{Id: 1067}

	// {{S1}} MP has been restored.
	S1MpRestored = SysMsg{Id: 1068}

	// {{S2}} MP has been restored by {{C1}}.
	S2MpRestoredByC1 = SysMsg{Id: 1069}

	// You do not have 'read' permission.
	NoReadPermission = SysMsg{Id: 1070}

	// You do not have 'write' permission.
	NoWritePermission = SysMsg{Id: 1071}

	// You have obtained a ticket for the Monster Race #{{S1}} - Single
	ObtainedTicketForMonsRaceS1Single = SysMsg{Id: 1072}

	// You have obtained a ticket for the Monster Race #{{S1}} - Single
	ObtainedTicketForMonsRaceS1Single_ = SysMsg{Id: 1073}

	// You do not meet the age requirement to purchase a Monster Race Ticket.
	NotMeetAgeRequirementForMonsRace = SysMsg{Id: 1074}

	// The bid amount must be higher than the previous bid.
	BidAmountHigherThanPreviousBid = SysMsg{Id: 1075}

	// The game cannot be terminated at this time.
	GameCannotTerminateNow = SysMsg{Id: 1076}

	// A GameGuard Execution error has occurred. Please send the *.erl file(s) located in the GameGuard folder to game@inca.co.kr
	GgExecutionError = SysMsg{Id: 1077}

	// When a user's keyboard input exceeds a certain cumulative score a chat ban will be applied. This is done to discourage spamming. Please avoid posting the same message multiple times during a short period.
	DontSpam = SysMsg{Id: 1078}

	// The target is currently banend from chatting.
	TargetIsChatBanned = SysMsg{Id: 1079}

	// Being permanent, are you sure you wish to use the facelift potion - Type A?
	FaceliftPotionTypeA = SysMsg{Id: 1080}

	// Being permanent, are you sure you wish to use the hair dye potion - Type A?
	HairdyePotionTypeA = SysMsg{Id: 1081}

	// Do you wish to use the hair style change potion - Type A? It is permanent.
	HairstylePotionTypeA = SysMsg{Id: 1082}

	// Facelift potion - Type A is being applied.
	FaceliftPotionTypeAApplied = SysMsg{Id: 1083}

	// Hair dye potion - Type A is being applied.
	HairdyePotionTypeAApplied = SysMsg{Id: 1084}

	// The hair style chance potion - Type A is being used.
	HairstylePotionTypeAUsed = SysMsg{Id: 1085}

	// Your facial appearance has been changed.
	FaceAppearanceChanged = SysMsg{Id: 1086}

	// Your hair color has changed.
	HairColorChanged = SysMsg{Id: 1087}

	// Your hair style has been changed.
	HairStyleChanged = SysMsg{Id: 1088}

	// {{C1}} has obtained a first anniversary commemorative item.
	C1ObtainedAnniversaryItem = SysMsg{Id: 1089}

	// Being permanent, are you sure you wish to use the facelift potion - Type B?
	FaceliftPotionTypeB = SysMsg{Id: 1090}

	// Being permanent, are you sure you wish to use the facelift potion - Type C?
	FaceliftPotionTypeC = SysMsg{Id: 1091}

	// Being permanent, are you sure you wish to use the hair dye potion - Type B?
	HairdyePotionTypeB = SysMsg{Id: 1092}

	// Being permanent, are you sure you wish to use the hair dye potion - Type C?
	HairdyePotionTypeC = SysMsg{Id: 1093}

	// Being permanent, are you sure you wish to use the hair dye potion - Type D?
	HairdyePotionTypeD = SysMsg{Id: 1094}

	// Do you wish to use the hair style change potion - Type B? It is permanent.
	HairstylePotionTypeB = SysMsg{Id: 1095}

	// Do you wish to use the hair style change potion - Type C? It is permanent.
	HairstylePotionTypeC = SysMsg{Id: 1096}

	// Do you wish to use the hair style change potion - Type D? It is permanent.
	HairstylePotionTypeD = SysMsg{Id: 1097}

	// Do you wish to use the hair style change potion - Type E? It is permanent.
	HairstylePotionTypeE = SysMsg{Id: 1098}

	// Do you wish to use the hair style change potion - Type F? It is permanent.
	HairstylePotionTypeF = SysMsg{Id: 1099}

	// Do you wish to use the hair style change potion - Type G? It is permanent.
	HairstylePotionTypeG = SysMsg{Id: 1100}

	// Facelift potion - Type B is being applied.
	FaceliftPotionTypeBApplied = SysMsg{Id: 1101}

	// Facelift potion - Type C is being applied.
	FaceliftPotionTypeCApplied = SysMsg{Id: 1102}

	// Hair dye potion - Type B is being applied.
	HairdyePotionTypeBApplied = SysMsg{Id: 1103}

	// Hair dye potion - Type C is being applied.
	HairdyePotionTypeCApplied = SysMsg{Id: 1104}

	// Hair dye potion - Type D is being applied.
	HairdyePotionTypeDApplied = SysMsg{Id: 1105}

	// The hair style chance potion - Type B is being used.
	HairstylePotionTypeBUsed = SysMsg{Id: 1106}

	// The hair style chance potion - Type C is being used.
	HairstylePotionTypeCUsed = SysMsg{Id: 1107}

	// The hair style chance potion - Type D is being used.
	HairstylePotionTypeDUsed = SysMsg{Id: 1108}

	// The hair style chance potion - Type E is being used.
	HairstylePotionTypeEUsed = SysMsg{Id: 1109}

	// The hair style chance potion - Type F is being used.
	HairstylePotionTypeFUsed = SysMsg{Id: 1110}

	// The hair style chance potion - Type G is being used.
	HairstylePotionTypeGUsed = SysMsg{Id: 1111}

	// The prize amount for the winner of Lottery #{{S1}} is {{S2}} adena. We have {{S3}} first prize winners.
	AmountForWinnerS1IsS2AdenaWeHaveS3PrizeWinner = SysMsg{Id: 1112}

	// The prize amount for Lucky Lottery #{{S1}} is {{S2}} adena. There was no first prize winner in this drawing, therefore the jackpot will be added to the next drawing.
	AmountForLotteryS1IsS2AdenaNoWinner = SysMsg{Id: 1113}

	// Your clan may not register to participate in a siege while under a grace period of the clan's dissolution.
	CantParticipateInSiegeWhileDissolutionInProgress = SysMsg{Id: 1114}

	// Individuals may not surrender during combat.
	IndividualsNotSurrenderDuringCombat = SysMsg{Id: 1115}

	// One cannot leave one's clan during combat.
	YouCannotLeaveDuringCombat = SysMsg{Id: 1116}

	// A clan member may not be dismissed during combat.
	ClanMemberCannotBeDismissedDuringCombat = SysMsg{Id: 1117}

	// Progress in a quest is possible only when your inventory's weight and volume are less than 80 percent of capacity.
	InventoryLessThan80Percent = SysMsg{Id: 1118}

	// Quest was automatically canceled when you attempted to settle the accounts of your quest while your inventory exceeded 80 percent of capacity.
	QuestCanceledInventoryExceeds80Percent = SysMsg{Id: 1119}

	// You are still a member of the clan.
	StillClanMember = SysMsg{Id: 1120}

	// You do not have the right to vote.
	NoRightToVote = SysMsg{Id: 1121}

	// There is no candidate.
	NoCandidate = SysMsg{Id: 1122}

	// Weight and volume limit has been exceeded. That skill is currently unavailable.
	WeightExceededSkillUnavailable = SysMsg{Id: 1123}

	// Your recipe book may not be accessed while using a skill.
	NoRecipeBookWhileCasting = SysMsg{Id: 1124}

	// An item may not be created while engaged in trading.
	CannotCreatedWhileEngagedInTrading = SysMsg{Id: 1125}

	// You cannot enter a negative number.
	NoNegativeNumber = SysMsg{Id: 1126}

	// The reward must be less than 10 times the standard price.
	RewardLessThan10TimesStandardPrice = SysMsg{Id: 1127}

	// A private store may not be opened while using a skill.
	PrivateStoreNotWhileCasting = SysMsg{Id: 1128}

	// This is not allowed while riding a ferry or boat.
	NotAllowedOnBoat = SysMsg{Id: 1129}

	// You have given {{S1}} damage to your target and {{S2}} damage to the servitor.
	GivenS1DamageToYourTargetAndS2DamageToServitor = SysMsg{Id: 1130}

	// It is now midnight and the effect of {{S1}} can be felt.
	NightEffectApplies = SysMsg{Id: 1131}

	// It is now dawn and the effect of {{S1}} will now disappear.
	DayEffectDisappears = SysMsg{Id: 1132}

	// Since HP has decreased, the effect of {{S1}} can be felt.
	HpDecreasedEffectApplies = SysMsg{Id: 1133}

	// Since HP has increased, the effect of {{S1}} will disappear.
	HpIncreasedEffectDisappears = SysMsg{Id: 1134}

	// While you are engaged in combat, you cannot operate a private store or private workshop.
	CantOperatePrivateStoreDuringCombat = SysMsg{Id: 1135}

	// Since there was an account that used this IP and attempted to log in illegally, this account is not allowed to connect to the game server for {{S1}} minutes. Please use another game server.
	AccountNotAllowedToConnect = SysMsg{Id: 1136}

	// {{C1}} harvested {{S3}} {{S2}}(s).
	C1HarvestedS3S2S = SysMsg{Id: 1137}

	// {{C1}} harvested {{S2}}(s).
	C1HarvestedS2S = SysMsg{Id: 1138}

	// The weight and volume limit of your inventory must not be exceeded.
	InventoryLimitMustNotBeExceeded = SysMsg{Id: 1139}

	// Would you like to open the gate?
	WouldYouLikeToOpenTheGate = SysMsg{Id: 1140}

	// Would you like to close the gate?
	WouldYouLikeToCloseTheGate = SysMsg{Id: 1141}

	// Since {{S1}} already exists nearby, you cannot summon it again.
	CannotSummonS1Again = SysMsg{Id: 1142}

	// Since you do not have enough items to maintain the servitor's stay, the servitor will disappear.
	ServitorDisappearedNotEnoughItems = SysMsg{Id: 1143}

	// Currently, you don't have anybody to chat with in the game.
	NobodyInGameToChat = SysMsg{Id: 1144}

	// {{S2}} has been created for {{C1}} after the payment of {{S3}} adena is received.
	S2CreatedForC1ForS3Adena = SysMsg{Id: 1145}

	// {{C1}} created {{S2}} after receiving {{S3}} adena.
	C1CreatedS2ForS3Adena = SysMsg{Id: 1146}

	// {{S2}} {{S3}} have been created for {{C1}} at the price of {{S4}} adena.
	S2S3SCreatedForC1ForS4Adena = SysMsg{Id: 1147}

	// {{C1}} created {{S2}} {{S3}} at the price of {{S4}} adena.
	C1CreatedS2S3SForS4Adena = SysMsg{Id: 1148}

	// Your attempt to create {{S2}} for {{C1}} at the price of {{S3}} adena has failed.
	CreationOfS2ForC1AtS3AdenaFailed = SysMsg{Id: 1149}

	// {{C1}} has failed to create {{S2}} at the price of {{S3}} adena.
	C1FailedToCreateS2ForS3Adena = SysMsg{Id: 1150}

	// {{S2}} is sold to {{C1}} at the price of {{S3}} adena.
	S2SoldToC1ForS3Adena = SysMsg{Id: 1151}

	// {{S2}} {{S3}} have been sold to {{C1}} for {{S4}} adena.
	S3S2SSoldToC1ForS4Adena = SysMsg{Id: 1152}

	// {{S2}} has been purchased from {{C1}} at the price of {{S3}} adena.
	S2PurchasedFromC1ForS3Adena = SysMsg{Id: 1153}

	// {{S3}} {{S2}} has been purchased from {{C1}} for {{S4}} adena.
	S3S2SPurchasedFromC1ForS4Adena = SysMsg{Id: 1154}

	// +{{S2}} {{S3}} have been sold to {{C1}} for {{S4}} adena.
	S3S2SoldToC1ForS4Adena = SysMsg{Id: 1155}

	// +{{S2}} {{S3}} has been purchased from {{C1}} for {{S4}} adena.
	S2S3PurchasedFromC1ForS4Adena = SysMsg{Id: 1156}

	// Trying on state lasts for only 5 seconds. When a character's state changes, it can be cancelled.
	TryingOnState = SysMsg{Id: 1157}

	// You cannot dismount from this elevation.
	CannotDismountFromElevation = SysMsg{Id: 1158}

	// The ferry from Talking Island will arrive at Gludin Harbor in approximately 10 minutes.
	FerryFromTalkingArriveAtGludin10Minutes = SysMsg{Id: 1159}

	// The ferry from Talking Island will be arriving at Gludin Harbor in approximately 5 minutes.
	FerryFromTalkingArriveAtGludin5Minutes = SysMsg{Id: 1160}

	// The ferry from Talking Island will be arriving at Gludin Harbor in approximately 1 minute.
	FerryFromTalkingArriveAtGludin1Minute = SysMsg{Id: 1161}

	// The ferry from Giran Harbor will be arriving at Talking Island in approximately 15 minutes.
	FerryFromGiranArriveAtTalking15Minutes = SysMsg{Id: 1162}

	// The ferry from Giran Harbor will be arriving at Talking Island in approximately 10 minutes.
	FerryFromGiranArriveAtTalking10Minutes = SysMsg{Id: 1163}

	// The ferry from Giran Harbor will be arriving at Talking Island in approximately 5 minutes.
	FerryFromGiranArriveAtTalking5Minutes = SysMsg{Id: 1164}

	// The ferry from Giran Harbor will be arriving at Talking Island in approximately 1 minute.
	FerryFromGiranArriveAtTalking1Minute = SysMsg{Id: 1165}

	// The ferry from Talking Island will be arriving at Giran Harbor in approximately 20 minutes.
	FerryFromTalkingArriveAtGiran20Minutes = SysMsg{Id: 1166}

	// The ferry from Talking Island will be arriving at Giran Harbor in approximately 20 minutes.
	FerryFromTalkingArriveAtGiran15Minutes = SysMsg{Id: 1167}

	// The ferry from Talking Island will be arriving at Giran Harbor in approximately 20 minutes.
	FerryFromTalkingArriveAtGiran10Minutes = SysMsg{Id: 1168}

	// The ferry from Talking Island will be arriving at Giran Harbor in approximately 20 minutes.
	FerryFromTalkingArriveAtGiran5Minutes = SysMsg{Id: 1169}

	// The ferry from Talking Island will be arriving at Giran Harbor in approximately 1 minute.
	FerryFromTalkingArriveAtGiran1Minute = SysMsg{Id: 1170}

	// The Innadril pleasure boat will arrive in approximately 20 minutes.
	InnadriLBoatArrive20Minutes = SysMsg{Id: 1171}

	// The Innadril pleasure boat will arrive in approximately 15 minutes.
	InnadriLBoatArrive15Minutes = SysMsg{Id: 1172}

	// The Innadril pleasure boat will arrive in approximately 10 minutes.
	InnadriLBoatArrive10Minutes = SysMsg{Id: 1173}

	// The Innadril pleasure boat will arrive in approximately 5 minutes.
	InnadriLBoatArrive5Minutes = SysMsg{Id: 1174}

	// The Innadril pleasure boat will arrive in approximately 1 minute.
	InnadriLBoatArrive1Minute = SysMsg{Id: 1175}

	// The SSQ Competition period is underway.
	SsqCompetitionUnderway = SysMsg{Id: 1176}

	// This is the seal validation period.
	ValidationPeriod = SysMsg{Id: 1177}

	// <Seal of Avarice description>
	AvariceDescription = SysMsg{Id: 1178}

	// <Seal of Gnosis description>
	GnosisDescription = SysMsg{Id: 1179}

	// <Seal of Strife description>
	StrifeDescription = SysMsg{Id: 1180}

	// Do you really wish to change the title?
	ChangeTitleConfirm = SysMsg{Id: 1181}

	// Are you sure you wish to delete the clan crest?
	CrestDeleteConfirm = SysMsg{Id: 1182}

	// This is the initial period.
	InitialPeriod = SysMsg{Id: 1183}

	// This is a period of calculating statistics in the server.
	ResultsPeriod = SysMsg{Id: 1184}

	// days left until deletion.
	DaysLeftUntilDeletion = SysMsg{Id: 1185}

	// To create a new account, please visit the PlayNC website (http://www.plaync.com/us/support/)
	ToCreateAccountVisitWebsite = SysMsg{Id: 1186}

	// If you forgotten your account information or password, please visit the Support Center on the PlayNC website(http://www.plaync.com/us/support/)
	AccountInformationForgottonVisitWebsite = SysMsg{Id: 1187}

	// Your selected target can no longer receive a recommendation.
	YourTargetNoLongerReceiveARecommendation = SysMsg{Id: 1188}

	// This temporary alliance of the Castle Attacker team has been dissolved.
	TemporaryAlliance = SysMsg{Id: 1189}

	TemporaryAllianceDissolved = SysMsg{Id: 1190}

	// The ferry from Gludin Harbor will be arriving at Talking Island in approximately 10 minutes.
	FerryFromGludinArriveAtTalking10Minutes = SysMsg{Id: 1191}

	// The ferry from Gludin Harbor will be arriving at Talking Island in approximately 5 minutes.
	FerryFromGludinArriveAtTalking5Minutes = SysMsg{Id: 1192}

	// The ferry from Gludin Harbor will be arriving at Talking Island in approximately 1 minute.
	FerryFromGludinArriveAtTalking1Minute = SysMsg{Id: 1193}

	// A mercenary can be assigned to a position from the beginning of the Seal Validatio period until the time when a siege starts.
	MercCanBeAssigned = SysMsg{Id: 1194}

	// This mercenary cannot be assigned to a position by using the Seal of Strife.
	MercCantBeAssignedUsingStrife = SysMsg{Id: 1195}

	// Your force has reached maximum capacity.
	ForceMaximum = SysMsg{Id: 1196}

	// Summoning a servitor costs {{S2}} {{S1}}.
	SummoningServitorCostsS2S1 = SysMsg{Id: 1197}

	// The item has been successfully crystallized.
	CrystallizationSuccessful = SysMsg{Id: 1198}

	// =======<Clan War Target>=======
	ClanWarHeader = SysMsg{Id: 1199}

	// ({{S1}} ({{S2}} Alliance)
	S1S2Alliance = SysMsg{Id: 1200}

	// Please select the quest you wish to abort.
	SelectQuestToAbor = SysMsg{Id: 1201}

	// ({{S1}} (No alliance exists)
	S1NoAlliExists = SysMsg{Id: 1202}

	// There is no clan war in progress.
	NoWarInProgress = SysMsg{Id: 1203}

	// The screenshot has been saved. ({{S1}} {{S2}}x{{S3}})
	Screenshot = SysMsg{Id: 1204}

	// Your mailbox is full. There is a 100 message limit.
	MailboxFull = SysMsg{Id: 1205}

	// The memo box is full. There is a 100 memo limit.
	MemoboxFull = SysMsg{Id: 1206}

	// Please make an entry in the field.
	MakeAnEntry = SysMsg{Id: 1207}

	// {{C1}} died and dropped {{S3}} {{S2}}.
	C1DiedDroppedS3S2 = SysMsg{Id: 1208}

	// Congratulations. Your raid was successful.
	RaidWasSuccessful = SysMsg{Id: 1209}

	// Seven Signs: The quest event period has begun. Visit a Priest of Dawn or Priestess of Dusk to participate in the event.
	QuestEventPeriodBegun = SysMsg{Id: 1210}

	// Seven Signs: The quest event period has ended. The next quest event will start in one week.
	QuestEventPeriodEnded = SysMsg{Id: 1211}

	// Seven Signs: The Lords of Dawn have obtained the Seal of Avarice.
	DawnObtainedAvarice = SysMsg{Id: 1212}

	// Seven Signs: The Lords of Dawn have obtained the Seal of Gnosis.
	DawnObtainedGnosis = SysMsg{Id: 1213}

	// Seven Signs: The Lords of Dawn have obtained the Seal of Strife.
	DawnObtainedStrife = SysMsg{Id: 1214}

	// Seven Signs: The Revolutionaries of Dusk have obtained the Seal of Avarice.
	DuskObtainedAvarice = SysMsg{Id: 1215}

	// Seven Signs: The Revolutionaries of Dusk have obtained the Seal of Gnosis.
	DuskObtainedGnosis = SysMsg{Id: 1216}

	// Seven Signs: The Revolutionaries of Dusk have obtained the Seal of Strife.
	DuskObtainedStrife = SysMsg{Id: 1217}

	// Seven Signs: The Seal Validation period has begun.
	SealValidationPeriodBegun = SysMsg{Id: 1218}

	// Seven Signs: The Seal Validation period has ended.
	SealValidationPeriodEnded = SysMsg{Id: 1219}

	// Are you sure you wish to summon it?
	SummonConfirm = SysMsg{Id: 1220}

	// Are you sure you wish to return it?
	ReturnConfirm = SysMsg{Id: 1221}

	// Current location : {{S1}}, {{S2}}, {{S3}} (GM Consultation Service)
	LocGmConsulationServiceS1S2S3 = SysMsg{Id: 1222}

	// We depart for Talking Island in five minutes.
	DepartForTalking5Minutes = SysMsg{Id: 1223}

	// We depart for Talking Island in one minute.
	DepartForTalking1Minute = SysMsg{Id: 1224}

	// All aboard for Talking Island
	DepartForTalking = SysMsg{Id: 1225}

	// We are now leaving for Talking Island.
	LeavingForTalking = SysMsg{Id: 1226}

	// You have {{S1}} unread messages.
	S1UnreadMessages = SysMsg{Id: 1227}

	// {{C1}} has blocked you. You cannot send mail to {{C1}}.
	C1BlockedYouCannotMail = SysMsg{Id: 1228}

	// No more messages may be sent at this time. Each account is allowed 10 messages per day.
	NoMoreMessagesToday = SysMsg{Id: 1229}

	// You are limited to five recipients at a time.
	OnlyFiveRecipients = SysMsg{Id: 1230}

	// You've sent mail.
	SentMail = SysMsg{Id: 1231}

	// The message was not sent.
	MessageNotSent = SysMsg{Id: 1232}

	// You've got mail.
	NewMail = SysMsg{Id: 1233}

	// The mail has been stored in your temporary mailbox.
	MailStoredInMailbox = SysMsg{Id: 1234}

	// Do you wish to delete all your friends?
	AllFriendsDeleteConfirm = SysMsg{Id: 1235}

	// Please enter security card number.
	EnterSecurityCardNumber = SysMsg{Id: 1236}

	// Please enter the card number for number {{S1}}.
	EnterCardNumberForS1 = SysMsg{Id: 1237}

	// Your temporary mailbox is full. No more mail can be stored; you have reached the 10 message limit.
	TempMailboxFull = SysMsg{Id: 1238}

	// The keyboard security module has failed to load. Please exit the game and try again.
	KeyboardModuleFailedLoad = SysMsg{Id: 1239}

	// Seven Signs: The Revolutionaries of Dusk have won.
	DuskWon = SysMsg{Id: 1240}

	// Seven Signs: The Lords of Dawn have won.
	DawnWon = SysMsg{Id: 1241}

	// Users who have not verified their age may not log in between the hours if 10:00 p.m. and 6:00 a.m.
	NotVerifiedAgeNoLogin = SysMsg{Id: 1242}

	// The security card number is invalid.
	SecurityCardNumberInvalid = SysMsg{Id: 1243}

	// Users who have not verified their age may not log in between the hours if 10:00 p.m. and 6:00 a.m. Logging off now
	NotVerifiedAgeLogOff = SysMsg{Id: 1244}

	// You will be loged out in {{S1}} minutes.
	LogoutInS1Minutes = SysMsg{Id: 1245}

	// {{C1}} died and has dropped {{S2}} adena.
	C1DiedDroppedS2Adena = SysMsg{Id: 1246}

	// The corpse is too old. The skill cannot be used.
	CorpseTooOldSkillNotUsed = SysMsg{Id: 1247}

	// You are out of feed. Mount status canceled.
	OutOfFeedMountCanceled = SysMsg{Id: 1248}

	// You may only ride a wyvern while you're riding a strider.
	YouMayOnlyRideWyvernWhileRidingStrider = SysMsg{Id: 1249}

	// Do you really want to surrender? If you surrender during an alliance war, your Exp will drop the same as if you were to die once.
	SurrenderAllyWarConfirm = SysMsg{Id: 1250}

	// Are you sure you want to dismiss the alliance? If you use the /allydismiss command, you will not be able to accept another clan to your alliance for one day.
	DismissAllyConfirm = SysMsg{Id: 1251}

	// Are you sure you want to surrender? Exp penalty will be the same as death.
	SurrenderConfirm1 = SysMsg{Id: 1252}

	// Are you sure you want to surrender? Exp penalty will be the same as death and you will not be allowed to participate in clan war.
	SurrenderConfirm2 = SysMsg{Id: 1253}

	// Thank you for submitting feedback.
	ThanksForFeedback = SysMsg{Id: 1254}

	// GM consultation has begun.
	GmConsultationBegun = SysMsg{Id: 1255}

	// Please write the name after the command.
	PleaseWriteNameAfterCommand = SysMsg{Id: 1256}

	// The special skill of a servitor or pet cannot be registerd as a macro.
	PetSkillNotAsMacro = SysMsg{Id: 1257}

	// {{S1}} has been crystallized
	S1Crystallized = SysMsg{Id: 1258}

	// =======<Alliance Target>=======
	AllianceTargetHeader = SysMsg{Id: 1259}

	// Seven Signs: Preparations have begun for the next quest event.
	PreparationsPeriodBegun = SysMsg{Id: 1260}

	// Seven Signs: The quest event period has begun. Speak with a Priest of Dawn or Dusk Priestess if you wish to participate in the event.
	CompetitionPeriodBegun = SysMsg{Id: 1261}

	// Seven Signs: Quest event has ended. Results are being tallied.
	ResultsPeriodBegun = SysMsg{Id: 1262}

	// Seven Signs: This is the seal validation period. A new quest event period begins next Monday.
	ValidationPeriodBegun = SysMsg{Id: 1263}

	// This soul stone cannot currently absorb souls. Absorption has failed.
	StoneCannotAbsorb = SysMsg{Id: 1264}

	// You can't absorb souls without a soul stone.
	CantAbsorbWithoutStone = SysMsg{Id: 1265}

	// The exchange has ended.
	ExchangeHasEnded = SysMsg{Id: 1266}

	// Your contribution score is increased by {{S1}}.
	ContribScoreIncreasedS1 = SysMsg{Id: 1267}

	// Do you wish to add class as your sub class?
	AddSubclassConfirm = SysMsg{Id: 1268}

	// The new sub class has been added.
	AddNewSubclass = SysMsg{Id: 1269}

	// The transfer of sub class has been completed.
	SubclassTransferCompleted = SysMsg{Id: 1270}

	// Do you wish to participate? Until the next seal validation period, you are a member of the Revolutionaries of Dusk.
	DawnConfirm = SysMsg{Id: 1271}

	DuskConfirm = SysMsg{Id: 1272}

	// You will participate in the Seven Signs as a member of the Lords of Dawn.
	SevensignsPartecipationDawn = SysMsg{Id: 1273}

	// You will participate in the Seven Signs as a member of the Revolutionaries of Dusk.
	SevensignsPartecipationDusk = SysMsg{Id: 1274}

	// You've chosen to fight for the Seal of Avarice during this quest event period.
	FightForAvarice = SysMsg{Id: 1275}

	// You've chosen to fight for the Seal of Gnosis during this quest event period.
	FightForGnosis = SysMsg{Id: 1276}

	// You've chosen to fight for the Seal of Strife during this quest event period.
	FightForStrife = SysMsg{Id: 1277}

	// The NPC server is not operating at this time.
	NpcServerNotOperating = SysMsg{Id: 1278}

	// Contribution level has exceeded the limit. You may not continue.
	ContribScoreExceeded = SysMsg{Id: 1279}

	// Magic Critical Hit!
	CriticalHitMagic = SysMsg{Id: 1280}

	// Your excellent shield defense was a success!
	YourExcellentShieldDefenseWasASuccess = SysMsg{Id: 1281}

	// Your Karma has been changed to {{S1}}
	YourKarmaHasBeenChangedToS1 = SysMsg{Id: 1282}

	// The minimum frame option has been activated.
	MinimumFrameActivated = SysMsg{Id: 1283}

	// The minimum frame option has been deactivated.
	MinimumFrameDeactivated = SysMsg{Id: 1284}

	// No inventory exists: You cannot purchase an item.
	NoInventoryCannotPurchase = SysMsg{Id: 1285}

	// (Until next Monday at 6:00 p.m.)
	UntilMonday6Pm = SysMsg{Id: 1286}

	// (Until today at 6:00 p.m.)
	UntilToday6Pm = SysMsg{Id: 1287}

	// If trends continue, {{S1}} will win and the seal will belong to:
	S1WillWinCompetition = SysMsg{Id: 1288}

	// (Until next Monday at 6:00 p.m.)
	SealOwned10MoreVoted = SysMsg{Id: 1289}

	// Although the seal was not owned, since 35 percent or more people have voted.
	SealNotOwned35MoreVoted = SysMsg{Id: 1290}

	// Although the seal was owned during the previous period, less than 10% of people have voted.
	SealOwned10LessVoted = SysMsg{Id: 1291}

	// Since the seal was not owned during the previous period, and since less than 35 percent of people have voted.
	SealNotOwned35LessVoted = SysMsg{Id: 1292}

	// If current trends continue, it will end in a tie.
	CompetitionWillTie = SysMsg{Id: 1293}

	// The competition has ended in a tie. Therefore, nobody has been awarded the seal.
	CompetitionTieSealNotAwarded = SysMsg{Id: 1294}

	// Sub classes may not be created or changed while a skill is in use.
	SubclassNoChangeOrCreateWhileSkillInUse = SysMsg{Id: 1295}

	// You cannot open a Private Store here.
	NoPrivateStoreHere = SysMsg{Id: 1296}

	// You cannot open a Private Workshop here.
	NoPrivateWorkshopHere = SysMsg{Id: 1297}

	// Please confirm that you would like to exit the Monster Race Track.
	MonsExitConfirm = SysMsg{Id: 1298}

	// {{C1}}'s casting has been interrupted.
	C1CastingInterrupted = SysMsg{Id: 1299}

	// You are no longer trying on equipment.
	WearItemsStopped = SysMsg{Id: 1300}

	// Only a Lord of Dawn may use this.
	CanBeUsedByDawn = SysMsg{Id: 1301}

	// Only a Revolutionary of Dusk may use this.
	CanBeUsedByDusk = SysMsg{Id: 1302}

	// This may only be used during the quest event period.
	CanBeUsedDuringQuestEventPeriod = SysMsg{Id: 1303}

	// The influence of the Seal of Strife has caused all defensive registrations to be canceled.
	StrifeCanceledDefensiveRegistration = SysMsg{Id: 1304}

	// Seal Stones may only be transferred during the quest event period.
	SealStonesOnlyWhileQuest = SysMsg{Id: 1305}

	// You are no longer trying on equipment.
	NoLongerTryingOn = SysMsg{Id: 1306}

	// Only during the seal validation period may you settle your account.
	SettleAccountOnlyInSealValidation = SysMsg{Id: 1307}

	// Congratulations - You've completed a class transfer!
	ClassTransfer = SysMsg{Id: 1308}

	// To use this option, you must have the lastest version of MSN Messenger installed on your computer.
	LatestMsnRequired = SysMsg{Id: 1309}

	// For full functionality, the latest version of MSN Messenger must be installed on your computer.
	LatestMsnRecommended = SysMsg{Id: 1310}

	// Previous versions of MSN Messenger only provide the basic features for in-game MSN Messenger Chat. Add/Delete Contacts and other MSN Messenger options are not available
	MsnOnlyBasic = SysMsg{Id: 1311}

	// The latest version of MSN Messenger may be obtained from the MSN web site (http://messenger.msn.com).
	MsnObtainedFrom = SysMsg{Id: 1312}

	// {{S1}}, to better serve our customers, all chat histories [...]
	S1ChatHistoriesStored = SysMsg{Id: 1313}

	// Please enter the passport ID of the person you wish to add to your contact list.
	EnterPassportForAdding = SysMsg{Id: 1314}

	// Deleting a contact will remove that contact from MSN Messenger as well. The contact can still check your online status and well not be blocked from sending you a message.
	DeletingAContact = SysMsg{Id: 1315}

	// The contact will be deleted and blocked from your contact list.
	ContactWillDeleted = SysMsg{Id: 1316}

	// Would you like to delete this contact?
	ContactDeleteConfirm = SysMsg{Id: 1317}

	// Please select the contact you want to block or unblock.
	SelectContactForBlockUnblock = SysMsg{Id: 1318}

	// Please select the name of the contact you wish to change to another group.
	SelectContactForChangeGroup = SysMsg{Id: 1319}

	// After selecting the group you wish to move your contact to, press the OK button.
	SelectGroupPressOk = SysMsg{Id: 1320}

	// Enter the name of the group you wish to add.
	EnterGroupName = SysMsg{Id: 1321}

	// Select the group and enter the new name.
	SelectGroupEnterName = SysMsg{Id: 1322}

	// Select the group you wish to delete and click the OK button.
	SelectGroupToDelete = SysMsg{Id: 1323}

	// Signing in...
	SigningIn = SysMsg{Id: 1324}

	// You've logged into another computer and have been logged out of the .NET Messenger Service on this computer.
	AnotherComputerLogout = SysMsg{Id: 1325}

	// {{S1}} :
	S1D = SysMsg{Id: 1326}

	// The following message could not be delivered:
	MessageNotDelivered = SysMsg{Id: 1327}

	// Members of the Revolutionaries of Dusk will not be resurrected.
	DuskNotResurrected = SysMsg{Id: 1328}

	// You are currently blocked from using the Private Store and Private Workshop.
	BlockedFromUsingStore = SysMsg{Id: 1329}

	// You may not open a Private Store or Private Workshop for another {{S1}} minute(s)
	NoStoreForS1Minutes = SysMsg{Id: 1330}

	// You are no longer blocked from using the Private Store and Private Workshop
	NoLongerBlockedUsingStore = SysMsg{Id: 1331}

	// Items may not be used after your character or pet dies.
	NoItemsAfterDeath = SysMsg{Id: 1332}

	// The replay file is not accessible. Please verify that the replay.ini exists in your Linage 2 directory.
	ReplayInaccessible = SysMsg{Id: 1333}

	// The new camera data has been stored.
	NewCameraStored = SysMsg{Id: 1334}

	// The attempt to store the new camera data has failed.
	CameraStoringFailed = SysMsg{Id: 1335}

	// The replay file, {{S1}}.${{S2}} has been corrupted, please check the fle.
	ReplayS1S2Corrupted = SysMsg{Id: 1336}

	// This will terminate the replay. Do you wish to continue?
	ReplayTerminateConfirm = SysMsg{Id: 1337}

	// You have exceeded the maximum amount that may be transferred at one time.
	ExceededMaximumAmount = SysMsg{Id: 1338}

	// Once a macro is assigned to a shortcut, it cannot be run as a macro again.
	MacroShortcutNotRun = SysMsg{Id: 1339}

	// This server cannot be accessed by the coupon you are using.
	ServerNotAccessedByCoupon = SysMsg{Id: 1340}

	// Incorrect name and/or email address.
	IncorrectNameOrAddress = SysMsg{Id: 1341}

	// You are already logged in.
	AlreadyLoggedIn = SysMsg{Id: 1342}

	// Incorrect email address and/or password. Your attempt to log into .NET Messenger Service has failed.
	IncorrectAddressOrPassword = SysMsg{Id: 1343}

	// Your request to log into the .NET Messenger service has failed. Please verify that you are currently connected to the internet.
	NetLoginFailed = SysMsg{Id: 1344}

	// Click the OK button after you have selected a contact name.
	SelectContactClickOk = SysMsg{Id: 1345}

	// You are currently entering a chat message.
	CurrentlyEnteringChat = SysMsg{Id: 1346}

	// The Linage II messenger could not carry out the task you requested.
	MessengerFailedCarryingOutTask = SysMsg{Id: 1347}

	// {{S1}} has entered the chat room.
	S1EnteredChatRoom = SysMsg{Id: 1348}

	// {{S1}} has left the chat room.
	S1LeftChatRoom = SysMsg{Id: 1349}

	// The state will be changed to indicate "off-line." All the chat windows currently opened will be closed.
	GoingOffline = SysMsg{Id: 1350}

	// Click the Delete button after selecting the contact you wish to remove.
	SelectContactClickRemove = SysMsg{Id: 1351}

	// You have been added to {{S1}} ({{S2}})'s contact list.
	AddedToS1S2ContactList = SysMsg{Id: 1352}

	// You can set the option to show your status as always being off-line to all of your contacts.
	CanSetOptionToAlwaysShowOffline = SysMsg{Id: 1353}

	// You are not allowed to chat with a contact while chatting block is imposed.
	NoChatWhileBlocked = SysMsg{Id: 1354}

	// The contact is currently blocked from chatting.
	ContactCurrentlyBlocked = SysMsg{Id: 1355}

	// The contact is not currently logged in.
	ContactCurrentlyOffline = SysMsg{Id: 1356}

	// You have been blocked from chatting with that contact.
	YouAreBlocked = SysMsg{Id: 1357}

	// You are being logged out...
	YouAreLoggingOut = SysMsg{Id: 1358}

	// {{S1}} has logged in.
	S1LoggedIn2 = SysMsg{Id: 1359}

	// You have received a message from {{S1}}.
	GotMessageFromS1 = SysMsg{Id: 1360}

	// Due to a system error, you have been logged out of the .NET Messenger Service.
	LoggedOutDueToError = SysMsg{Id: 1361}

	SelectContactToDelete = SysMsg{Id: 1362}

	// Your request to participate in the alliance war has been denied.
	YourRequestAllianceWarDenied = SysMsg{Id: 1363}

	// The request for an alliance war has been rejected.
	RequestAllianceWarRejected = SysMsg{Id: 1364}

	// {{S2}} of {{S1}} clan has surrendered as an individual.
	S2OfS1SurrenderedAsIndividual = SysMsg{Id: 1365}

	// In order to delete a group, you must not [...]
	DelteGroupInstruction = SysMsg{Id: 1366}

	// Only members of the group are allowed to add records.
	OnlyGroupCanAddRecords = SysMsg{Id: 1367}

	// You can not try those items on at the same time.
	YouCanNotTryThoseItemsOnAtTheSameTime = SysMsg{Id: 1368}

	// You've exceeded the maximum.
	ExceededTheMaximum = SysMsg{Id: 1369}

	// Your message to {{C1}} did not reach its recipient. You cannot send mail to the GM staff.
	CannotMailGmC1 = SysMsg{Id: 1370}

	// It has been determined that you're not engaged in normal gameplay and a restriction has been imposed upon you. You may not move for {{S1}} minutes.
	GameplayRestrictionPenaltyS1 = SysMsg{Id: 1371}

	// Your punishment will continue for {{S1}} minutes.
	PunishmentContinueS1Minutes = SysMsg{Id: 1372}

	// {{C1}} has picked up {{S2}} that was dropped by a Raid Boss.
	C1PickedUpS2FromRaidboss = SysMsg{Id: 1373}

	// {{C1}} has picked up {{S3}} {{S2}}(s) that was dropped by a Raid Boss.
	C1PickedUpS3S2SFromRaidboss = SysMsg{Id: 1374}

	// {{C1}} has picked up {{S2}} adena that was dropped by a Raid Boss.
	C1PickedUpS2AdenaFromRaidboss = SysMsg{Id: 1375}

	// {{C1}} has picked up {{S2}} that was dropped by another character.
	C1PickedUpS2FromAnotherCharacter = SysMsg{Id: 1376}

	// {{C1}} has picked up {{S3}} {{S2}}(s) that was dropped by a another character.
	C1PickedUpS3S2SFromAnotherCharacter = SysMsg{Id: 1377}

	// {{C1}} has picked up +{{S3}} {{S2}} that was dropped by a another character.
	C1PickedUpS3S2FromAnotherCharacter = SysMsg{Id: 1378}

	// {{C1}} has obtained {{S2}} adena.
	C1ObtainedS2Adena = SysMsg{Id: 1379}

	// You can't summon a {{S1}} while on the battleground.
	CantSummonS1OnBattleground = SysMsg{Id: 1380}

	// The party leader has obtained {{S2}} of {{S1}}.
	LeaderObtainedS2OfS1 = SysMsg{Id: 1381}

	// To fulfill the quest, you must bring the chosen weapon. Are you sure you want to choose this weapon?
	ChooseWeaponConfirm = SysMsg{Id: 1382}

	// Are you sure you want to exchange?
	ExchangeConfirm = SysMsg{Id: 1383}

	// {{C1}} has become the party leader.
	C1HasBecomeAPartyLeader = SysMsg{Id: 1384}

	// You are not allowed to dismount at this location.
	NoDismountHere = SysMsg{Id: 1385}

	// You are no longer held in place.
	NoLongerHeldInPlace = SysMsg{Id: 1386}

	// Please select the item you would like to try on.
	SelectItemToTryOn = SysMsg{Id: 1387}

	// A party room has been created.
	PartyRoomCreated = SysMsg{Id: 1388}

	// The party room's information has been revised.
	PartyRoomRevised = SysMsg{Id: 1389}

	// You are not allowed to enter the party room.
	PartyRoomForbidden = SysMsg{Id: 1390}

	// You have exited from the party room.
	PartyRoomExited = SysMsg{Id: 1391}

	// {{C1}} has left the party room.
	C1LeftPartyRoom = SysMsg{Id: 1392}

	// You have been ousted from the party room.
	OustedFromPartyRoom = SysMsg{Id: 1393}

	// {{C1}} has been kicked from the party room.
	C1KickedFromPartyRoom = SysMsg{Id: 1394}

	// The party room has been disbanded.
	PartyRoomDisbanded = SysMsg{Id: 1395}

	// The list of party rooms can only be viewed by a person who has not joined a party or who is currently the leader of a party.
	CantViewPartyRooms = SysMsg{Id: 1396}

	// The leader of the party room has changed.
	PartyRoomLeaderChanged = SysMsg{Id: 1397}

	// We are recruiting party members.
	RecruitingPartyMembers = SysMsg{Id: 1398}

	// Only the leader of the party can transfer party leadership to another player.
	OnlyAPartyLeaderCanTransferOnesRightsToAnotherPlayer = SysMsg{Id: 1399}

	// Please select the person you wish to make the party leader.
	PleaseSelectThePersonToWhomYouWouldLikeToTransferTheRightsOfAPartyLeader = SysMsg{Id: 1400}

	// Slow down.you are already the party leader.
	YouCannotTransferRightsToYourself = SysMsg{Id: 1401}

	// You may only transfer party leadership to another member of the party.
	YouCanTransferRightsOnlyToAnotherPartyMember = SysMsg{Id: 1402}

	// You have failed to transfer the party leadership.
	YouHaveFailedToTransferThePartyLeaderRights = SysMsg{Id: 1403}

	// The owner of the private manufacturing store has changed the price for creating this item. Please check the new price before trying again.
	ManufacturePriceHasChanged = SysMsg{Id: 1404}

	// {{S1}} CPs have been restored.
	S1CpWillBeRestored = SysMsg{Id: 1405}

	// {{S2}} CPs has been restored by {{C1}}.
	S2CpWillBeRestoredByC1 = SysMsg{Id: 1406}

	// You are using a computer that does not allow you to log in with two accounts at the same time.
	NoLoginWithTwoAccounts = SysMsg{Id: 1407}

	// Your prepaid remaining usage time is {{S1}} hours and {{S2}} minutes. You have {{S3}} paid reservations left.
	PrepaidLeftS1S2S3 = SysMsg{Id: 1408}

	// Your prepaid usage time has expired. Your new prepaid reservation will be used. The remaining usage time is {{S1}} hours and {{S2}} minutes.
	PrepaidExpiredS1S2 = SysMsg{Id: 1409}

	// Your prepaid usage time has expired. You do not have any more prepaid reservations left.
	PrepaidExpired = SysMsg{Id: 1410}

	// The number of your prepaid reservations has changed.
	PrepaidChanged = SysMsg{Id: 1411}

	// Your prepaid usage time has {{S1}} minutes left.
	PrepaidLeftS1 = SysMsg{Id: 1412}

	// You do not meet the requirements to enter that party room.
	CantEnterPartyRoom = SysMsg{Id: 1413}

	// The width and length should be 100 or more grids and less than 5000 grids respectively.
	WrongGridCount = SysMsg{Id: 1414}

	// The command file is not sent.
	CommandFileNotSent = SysMsg{Id: 1415}

	// The representative of Team 1 has not been selected.
	Team1NoRepresentative = SysMsg{Id: 1416}

	// The representative of Team 2 has not been selected.
	Team2NoRepresentative = SysMsg{Id: 1417}

	// The name of Team 1 has not yet been chosen.
	Team1NoName = SysMsg{Id: 1418}

	// The name of Team 2 has not yet been chosen.
	Team2NoName = SysMsg{Id: 1419}

	// The name of Team 1 and the name of Team 2 are identical.
	TeamNameIdentical = SysMsg{Id: 1420}

	// The race setup file has not been designated.
	RaceSetupFile1 = SysMsg{Id: 1421}

	// Race setup file error - BuffCnt is not specified
	RaceSetupFile2 = SysMsg{Id: 1422}

	// Race setup file error - BuffID{{S1}} is not specified.
	RaceSetupFile3 = SysMsg{Id: 1423}

	// Race setup file error - BuffLv{{S1}} is not specified.
	RaceSetupFile4 = SysMsg{Id: 1424}

	// Race setup file error - DefaultAllow is not specified
	RaceSetupFile5 = SysMsg{Id: 1425}

	// Race setup file error - ExpSkillCnt is not specified.
	RaceSetupFile6 = SysMsg{Id: 1426}

	// Race setup file error - ExpSkillID{{S1}} is not specified.
	RaceSetupFile7 = SysMsg{Id: 1427}

	// Race setup file error - ExpItemCnt is not specified.
	RaceSetupFile8 = SysMsg{Id: 1428}

	// Race setup file error - ExpItemID{{S1}} is not specified.
	RaceSetupFile9 = SysMsg{Id: 1429}

	// Race setup file error - TeleportDelay is not specified
	RaceSetupFile10 = SysMsg{Id: 1430}

	// The race will be stopped temporarily.
	RaceStoppedTemporarily = SysMsg{Id: 1431}

	// Your opponent is currently in a petrified state.
	OpponentPetrified = SysMsg{Id: 1432}

	// You will now automatically apply {{S1}} to your target.
	UseOfS1WillBeAuto = SysMsg{Id: 1433}

	// You will no longer automatically apply {{S1}} to your weapon.
	AutoUseOfS1Cancelled = SysMsg{Id: 1434}

	// Due to insufficient {{S1}}, the automatic use function has been deactivated.
	AutoUseCancelledLackOfS1 = SysMsg{Id: 1435}

	// Due to insufficient {{S1}}, the automatic use function cannot be activated.
	CannotAutoUseLackOfS1 = SysMsg{Id: 1436}

	// Players are no longer allowed to play dice. Dice can no longer be purchased from a village store. However, you can still sell them to any village store.
	DiceNoLongerAllowed = SysMsg{Id: 1437}

	// There is no skill that enables enchant.
	ThereIsNoSkillThatEnablesEnchant = SysMsg{Id: 1438}

	// You do not have all of the items needed to enchant that skill.
	YouDontHaveAllOfTheItemsNeededToEnchantThatSkill = SysMsg{Id: 1439}

	// You have succeeded in enchanting the skill {{S1}}.
	YouHaveSucceededInEnchantingTheSkillS1 = SysMsg{Id: 1440}

	// Skill enchant failed. The skill will be initialized.
	YouHaveFailedToEnchantTheSkillS1 = SysMsg{Id: 1441}

	// You do not have enough SP to enchant that skill.
	YouDontHaveEnoughSpToEnchantThatSkill = SysMsg{Id: 1443}

	// You do not have enough experience (Exp) to enchant that skill.
	YouDontHaveEnoughExpToEnchantThatSkill = SysMsg{Id: 1444}

	// Your previous subclass will be removed and replaced with the new subclass at level 40. Do you wish to continue?
	ReplaceSubclassConfirm = SysMsg{Id: 1445}

	// The ferry from {{S1}} to {{S2}} has been delayed.
	FerryFromS1ToS2Delayed = SysMsg{Id: 1446}

	// You cannot do that while fishing.
	CannotDoWhileFishing1 = SysMsg{Id: 1447}

	// Only fishing skills may be used at this time.
	OnlyFishingSkillsNow = SysMsg{Id: 1448}

	// You've got a bite!
	GotABite = SysMsg{Id: 1449}

	// That fish is more determined than you are - it spit the hook!
	FishSpitTheHook = SysMsg{Id: 1450}

	// Your bait was stolen by that fish!
	BaitStolenByFish = SysMsg{Id: 1451}

	// Baits have been lost because the fish got away.
	BaitLostFishGotAway = SysMsg{Id: 1452}

	// You do not have a fishing pole equipped.
	FishingPoleNotEquipped = SysMsg{Id: 1453}

	// You must put bait on your hook before you can fish.
	BaitOnHookBeforeFishing = SysMsg{Id: 1454}

	// You cannot fish while under water.
	CannotFishUnderWater = SysMsg{Id: 1455}

	// You cannot fish while riding as a passenger of a boat - it's against the rules.
	CannotFishOnBoat = SysMsg{Id: 1456}

	// You can't fish here.
	CannotFishHere = SysMsg{Id: 1457}

	// Your attempt at fishing has been cancelled.
	FishingAttemptCancelled = SysMsg{Id: 1458}

	// You do not have enough bait.
	NotEnoughBait = SysMsg{Id: 1459}

	// You reel your line in and stop fishing.
	ReelLineAndStopFishing = SysMsg{Id: 1460}

	// You cast your line and start to fish.
	CastLineAndStartFishing = SysMsg{Id: 1461}

	// You may only use the Pumping skill while you are fishing.
	CanUsePumpingOnlyWhileFishing = SysMsg{Id: 1462}

	// You may only use the Reeling skill while you are fishing.
	CanUseReelingOnlyWhileFishing = SysMsg{Id: 1463}

	// The fish has resisted your attempt to bring it in.
	FishResistedAttemptToBringItIn = SysMsg{Id: 1464}

	// Your pumping is successful, causing {{S1}} damage.
	PumpingSuccesfulS1Damage = SysMsg{Id: 1465}

	// You failed to do anything with the fish and it regains {{S1}} HP.
	FishResistedPumpingS1HpRegained = SysMsg{Id: 1466}

	// You reel that fish in closer and cause {{S1}} damage.
	ReelingSuccesfulS1Damage = SysMsg{Id: 1467}

	// You failed to reel that fish in further and it regains {{S1}} HP.
	FishResistedReelingS1HpRegained = SysMsg{Id: 1468}

	// You caught something!
	YouCaughtSomething = SysMsg{Id: 1469}

	// You cannot do that while fishing.
	CannotDoWhileFishing2 = SysMsg{Id: 1470}

	// You cannot do that while fishing.
	CannotDoWhileFishing3 = SysMsg{Id: 1471}

	// You look oddly at the fishing pole in disbelief and realize that you can't attack anything with this.
	CannotAttackWithFishingPole = SysMsg{Id: 1472}

	// {{S1}} is not sufficient.
	S1NotSufficient = SysMsg{Id: 1473}

	// {{S1}} is not available.
	S1NotAvailable = SysMsg{Id: 1474}

	// Pet has dropped {{S1}}.
	PetDroppedS1 = SysMsg{Id: 1475}

	// Pet has dropped +{{S1}} {{S2}}.
	PetDroppedS1S2 = SysMsg{Id: 1476}

	// Pet has dropped {{S2}} of {{S1}}.
	PetDroppedS2S1S = SysMsg{Id: 1477}

	// You may only register a 64 x 64 pixel, 256-color BMP.
	Only64Pixel256ColorBmp = SysMsg{Id: 1478}

	// That is the wrong grade of soulshot for that fishing pole.
	WrongFishingshotGrade = SysMsg{Id: 1479}

	// Are you sure you want to remove yourself from the Grand Olympiad Games waiting list?
	OlympiadRemoveConfirm = SysMsg{Id: 1480}

	// You have selected a class irrelevant individual match. Do you wish to participate?
	OlympiadNonClassConfirm = SysMsg{Id: 1481}

	// You've selected to join a class specific game. Continue?
	OlympiadClassConfirm = SysMsg{Id: 1482}

	// Are you ready to be a Hero?
	HeroConfirm = SysMsg{Id: 1483}

	// Are you sure this is the Hero weapon you wish to use? Kamael race cannot use this.
	HeroWeaponConfirm = SysMsg{Id: 1484}

	// The ferry from Talking Island to Gludin Harbor has been delayed.
	FerryTalkingGludinDelayed = SysMsg{Id: 1485}

	// The ferry from Gludin Harbor to Talking Island has been delayed.
	FerryGludinTalkingDelayed = SysMsg{Id: 1486}

	// The ferry from Giran Harbor to Talking Island has been delayed.
	FerryGiranTalkingDelayed = SysMsg{Id: 1487}

	// The ferry from Talking Island to Giran Harbor has been delayed.
	FerryTalkingGiranDelayed = SysMsg{Id: 1488}

	// Innadril cruise service has been delayed.
	InnadriLBoatDelayed = SysMsg{Id: 1489}

	// Traded {{S2}} of crop {{S1}}.
	TradedS2OfCropS1 = SysMsg{Id: 1490}

	// Failed in trading {{S2}} of crop {{S1}}.
	FailedInTradingS2OfCropS1 = SysMsg{Id: 1491}

	// You will be moved to the Olympiad Stadium in {{S1}} second(s).
	YouWillEnterTheOlympiadStadiumInS1SecondS = SysMsg{Id: 1492}

	// Your opponent made haste with their tail between their legs), the match has been cancelled.
	TheGameHasBeenCancelledBecauseTheOtherPartyEndsTheGame = SysMsg{Id: 1493}

	// Your opponent does not meet the requirements to do battle), the match has been cancelled.
	TheGameHasBeenCancelledBecauseTheOtherPartyDoesNotMeetTheRequirementsForJoiningTheGame = SysMsg{Id: 1494}

	// The match will start in {{S1}} second(s).
	TheGameWillStartInS1SecondS = SysMsg{Id: 1495}

	// The match has started, fight!
	StartsTheGame = SysMsg{Id: 1496}

	// Congratulations, {{C1}}! You win the match!
	C1HasWonTheGame = SysMsg{Id: 1497}

	// There is no victor, the match ends in a tie.
	TheGameEndedInATie = SysMsg{Id: 1498}

	// You will be moved back to town in {{S1}} second(s).
	YouWillBeMovedToTownInS1Seconds = SysMsg{Id: 1499}

	// {{C1}}% does not meet the participation requirements. A sub-class character cannot participate in the Olympiad.
	C1CantJoinTheOlympiadWithASubClassCharacter = SysMsg{Id: 1500}

	// {{C1}}% does not meet the participation requirements. Only Noblesse can participate in the Olympiad.
	C1DoesNotMeetRequirementsOnlyNoblessCanParticipateInTheOlympiad = SysMsg{Id: 1501}

	// {{C1}} is already registered on the match waiting list.
	C1IsAlreadyRegisteredOnTheMatchWaitingList = SysMsg{Id: 1502}

	// You have been registered in the Grand Olympiad Games waiting list for a class specific match.
	YouHaveBeenRegisteredInAWaitingListOfClassifiedGames = SysMsg{Id: 1503}

	// You are currently registered for a 1v1 class irrelevant match.
	YouHaveBeenRegisteredInAWaitingListOfNoClassGames = SysMsg{Id: 1504}

	// You have been removed from the Grand Olympiad Games waiting list.
	YouHaveBeenDeletedFromTheWaitingListOfAGame = SysMsg{Id: 1505}

	// You are not currently registered on any Grand Olympiad Games waiting list.
	YouHaveNotBeenRegisteredInAWaitingListOfAGame = SysMsg{Id: 1506}

	// You cannot equip that item in a Grand Olympiad Games match.
	ThisItemCantBeEquippedForTheOlympiadEvent = SysMsg{Id: 1507}

	// You cannot use that item in a Grand Olympiad Games match.
	ThisItemIsNotAvailableForTheOlympiadEvent = SysMsg{Id: 1508}

	// You cannot use that skill in a Grand Olympiad Games match.
	ThisSkillIsNotAvailableForTheOlympiadEvent = SysMsg{Id: 1509}

	// {{C1}} is making an attempt at resurrection with {{S2}} experience points. Do you want to be resurrected?
	RessurectionRequestByC1ForS2Xp = SysMsg{Id: 1510}

	// While a pet is attempting to resurrect, it cannot help in resurrecting its master.
	MasterCannotRes = SysMsg{Id: 1511}

	// You cannot resurrect a pet while their owner is being resurrected.
	CannotResPet = SysMsg{Id: 1512}

	// Resurrection has already been proposed.
	ResHasAlreadyBeenProposed = SysMsg{Id: 1513}

	// You cannot the owner of a pet while their pet is being resurrected
	CannotResMaster = SysMsg{Id: 1514}

	// A pet cannot be resurrected while it's owner is in the process of resurrecting.
	CannotResPet2 = SysMsg{Id: 1515}

	// The target is unavailable for seeding.
	TheTargetIsUnavailableForSeeding = SysMsg{Id: 1516}

	// Failed in Blessed Enchant. The enchant value of the item became 0.
	BlessedEnchantFailed = SysMsg{Id: 1517}

	// You do not meet the required condition to equip that item.
	CannotEquipItemDueToBadCondition = SysMsg{Id: 1518}

	// The pet has been killed. If you don't resurrect it within 24 hours, the pet's body will disappear along with all the pet's items.
	MakeSureYouRessurectYourPetWithin24Hours = SysMsg{Id: 1519}

	// Servitor passed away.
	ServitorPassedAway = SysMsg{Id: 1520}

	// Your servitor has vanished! You'll need to summon a new one.
	YourServitorHasVanished = SysMsg{Id: 1521}

	// Your pet's corpse has decayed!
	YourPetsCorpseHasDecayed = SysMsg{Id: 1522}

	// You should release your pet or servitor so that it does not fall off of the boat and drown!
	ReleasePetOnBoat = SysMsg{Id: 1523}

	// {{C1}}'s pet gained {{S2}}.
	C1PetGainedS2 = SysMsg{Id: 1524}

	// {{C1}}'s pet gained {{S3}} of {{S2}}.
	C1PetGainedS3S2S = SysMsg{Id: 1525}

	// {{C1}}'s pet gained +{{S2}}{{S3}}.
	C1PetGainedS2S3 = SysMsg{Id: 1526}

	// Your pet was hungry so it ate {{S1}}.
	PetTookS1BecauseHeWasHungry = SysMsg{Id: 1527}

	// You've sent a petition to the GM staff.
	SentPetitionToGm = SysMsg{Id: 1528}

	// {{C1}} is inviting you to the command channel. Do you want accept?
	CommandChannelConfirmFromC1 = SysMsg{Id: 1529}

	// Select a target or enter the name.
	SelectTargetOrEnterName = SysMsg{Id: 1530}

	// Enter the name of the clan that you wish to declare war on.
	EnterClanNameToDeclareWar2 = SysMsg{Id: 1531}

	// Enter the name of the clan that you wish to have a cease-fire with.
	EnterClanNameToCeaseFire = SysMsg{Id: 1532}

	// Announcement: {{C1}} has picked up {{S2}}.
	AnnouncementC1PickedUpS2 = SysMsg{Id: 1533}

	// Announcement: {{C1}} has picked up +{{S2}}{{S3}}.
	AnnouncementC1PickedUpS2S3 = SysMsg{Id: 1534}

	// Announcement: {{C1}}'s pet has picked up {{S2}}.
	AnnouncementC1PetPickedUpS2 = SysMsg{Id: 1535}

	// Announcement: {{C1}}'s pet has picked up +{{S2}}{{S3}}.
	AnnouncementC1PetPickedUpS2S3 = SysMsg{Id: 1536}

	// Current Location: {{S1}}, {{S2}}, {{S3}} (near Rune Village)
	LocRuneS1S2S3 = SysMsg{Id: 1537}

	// Current Location: {{S1}}, {{S2}}, {{S3}} (near the Town of Goddard)
	LocGoddardS1S2S3 = SysMsg{Id: 1538}

	// Cargo has arrived at Talking Island Village.
	CargoAtTalkingVillage = SysMsg{Id: 1539}

	// Cargo has arrived at the Dark Elf Village.
	CargoAtDarkelfVillage = SysMsg{Id: 1540}

	// Cargo has arrived at Elven Village.
	CargoAtElvenVillage = SysMsg{Id: 1541}

	// Cargo has arrived at Orc Village.
	CargoAtOrcVillage = SysMsg{Id: 1542}

	// Cargo has arrived at Dwarfen Village.
	CargoAtDwarvenVillage = SysMsg{Id: 1543}

	// Cargo has arrived at Aden Castle Town.
	CargoAtAden = SysMsg{Id: 1544}

	// Cargo has arrived at Town of Oren.
	CargoAtOren = SysMsg{Id: 1545}

	// Cargo has arrived at Hunters Village.
	CargoAtHunters = SysMsg{Id: 1546}

	// Cargo has arrived at the Town of Dion.
	CargoAtDion = SysMsg{Id: 1547}

	// Cargo has arrived at Floran Village.
	CargoAtFloran = SysMsg{Id: 1548}

	// Cargo has arrived at Gludin Village.
	CargoAtGludin = SysMsg{Id: 1549}

	// Cargo has arrived at the Town of Gludio.
	CargoAtGludio = SysMsg{Id: 1550}

	// Cargo has arrived at Giran Castle Town.
	CargoAtGiran = SysMsg{Id: 1551}

	// Cargo has arrived at Heine.
	CargoAtHeine = SysMsg{Id: 1552}

	// Cargo has arrived at Rune Village.
	CargoAtRune = SysMsg{Id: 1553}

	// Cargo has arrived at the Town of Goddard.
	CargoAtGoddard = SysMsg{Id: 1554}

	// Do you want to cancel character deletion?
	CancelCharacterDeletionConfirm = SysMsg{Id: 1555}

	// Your clan notice has been saved.
	ClanNoticeSaved = SysMsg{Id: 1556}

	// Seed price should be more than {{S1}} and less than {{S2}}.
	SeedPriceShouldBeMoreThanS1AndLessThanS2 = SysMsg{Id: 1557}

	// The quantity of seed should be more than {{S1}} and less than {{S2}}.
	TheQuantityOfSeedShouldBeMoreThanS1AndLessThanS2 = SysMsg{Id: 1558}

	// Crop price should be more than {{S1}} and less than {{S2}}.
	CropPriceShouldBeMoreThanS1AndLessThanS2 = SysMsg{Id: 1559}

	// The quantity of crop should be more than {{S1}} and less than {{S2}}
	TheQuantityOfCropShouldBeMoreThanS1AndLessThanS2 = SysMsg{Id: 1560}

	// The clan, {{S1}}, has declared a Clan War.
	ClanS1DeclaredWar = SysMsg{Id: 1561}

	// A Clan War has been declared against the clan, {{S1}}. you will only lose a quarter of the normal experience from death.
	ClanWarDeclaredAgainstS1IfKilledLoseLowExp = SysMsg{Id: 1562}

	CannotDeclareWarTooLowLevelOrNotEnoughMembers = SysMsg{Id: 1563}

	// A Clan War can be declared only if the clan is level three or above, and the number of clan members is fifteen or greater.
	ClanWarDeclaredIfClanLvl3Or15Member = SysMsg{Id: 1564}

	// A Clan War cannot be declared against a clan that does not exist!
	ClanWarCannotDeclaredClanNotExist = SysMsg{Id: 1565}

	// The clan, {{S1}}, has decided to stop the war.
	ClanS1HasDecidedToStop = SysMsg{Id: 1566}

	// The war against {{S1}} Clan has been stopped.
	WarAgainstS1HasStopped = SysMsg{Id: 1567}

	// The target for declaration is wrong.
	WrongDeclarationTarget = SysMsg{Id: 1568}

	// A declaration of Clan War against an allied clan can't be made.
	ClanWarAgainstAAlliedClanNotWork = SysMsg{Id: 1569}

	// A declaration of war against more than 30 Clans can't be made at the same time
	TooManyClanWars = SysMsg{Id: 1570}

	// ======<Clans You've Declared War On>======
	ClansYouDeclaredWarOn = SysMsg{Id: 1571}

	// ======<Clans That Have Declared War On You>======
	ClansThatHaveDeclaredWarOnYou = SysMsg{Id: 1572}

	// All is well. There are no clans that have declared war against your clan.
	NoWarsAgainstYou = SysMsg{Id: 1573}

	// Command Channels can only be formed by a party leader who is also the leader of a level 5 clan.
	CommandChannelOnlyByLevel5ClanLeaderPartyLeader = SysMsg{Id: 1574}

	// Your pet uses spiritshot.
	PetUseSpiritshot = SysMsg{Id: 1575}

	// Your servitor uses spiritshot.
	ServitorUseSpiritshot = SysMsg{Id: 1576}

	// Servitor uses the power of spirit.
	ServitorUseThePowerOfSpirit = SysMsg{Id: 1577}

	// Items are not available for a private store or a private manufacture.
	ItemsUnavailableForStoreManufacture = SysMsg{Id: 1578}

	// {{C1}}'s pet gained {{S2}} adena.
	C1PetGainedS2Adena = SysMsg{Id: 1579}

	// The Command Channel has been formed.
	CommandChannelFormed = SysMsg{Id: 1580}

	// The Command Channel has been disbanded.
	CommandChannelDisbanded = SysMsg{Id: 1581}

	// You have joined the Command Channel.
	JoinedCommandChannel = SysMsg{Id: 1582}

	// You were dismissed from the Command Channel.
	DismissedFromCommandChannel = SysMsg{Id: 1583}

	// {{C1}}'s party has been dismissed from the Command Channel.
	C1PartyDismissedFromCommandChannel = SysMsg{Id: 1584}

	// The Command Channel has been disbanded.
	CommandChannelDisbanded2 = SysMsg{Id: 1585}

	// You have quit the Command Channel.
	LeftCommandChannel = SysMsg{Id: 1586}

	// {{C1}}'s party has left the Command Channel.
	C1PartyLeftCommandChannel = SysMsg{Id: 1587}

	// The Command Channel is activated only when there are at least 5 parties participating.
	CommandChannelOnlyAtLeast5Parties = SysMsg{Id: 1588}

	// Command Channel authority has been transferred to {{C1}}.
	CommandChannelLeaderNowC1 = SysMsg{Id: 1589}

	// ===<Guild Info (Total Parties: {{S1}})>===
	GuildInfoHeader = SysMsg{Id: 1590}

	// No user has been invited to the Command Channel.
	NoUserInvitedToCommandChannel = SysMsg{Id: 1591}

	// You can no longer set up a Command Channel.
	CannotLongerSetupCommandChannel = SysMsg{Id: 1592}

	// You do not have authority to invite someone to the Command Channel.
	CannotInviteToCommandChannel = SysMsg{Id: 1593}

	// {{C1}}'s party is already a member of the Command Channel.
	C1AlreadyMemberOfCommandChannel = SysMsg{Id: 1594}

	// {{S1}} has succeeded.
	S1Succeeded = SysMsg{Id: 1595}

	// You were hit by {{S1}}!
	HitByS1 = SysMsg{Id: 1596}

	// {{S1}} has failed.
	S1Failed = SysMsg{Id: 1597}

	// Soulshots and spiritshots are not available for a dead pet or servitor. Sad, isn't it?
	SoulshotsAndSpiritshotsAreNotAvailableForADeadPet = SysMsg{Id: 1598}

	// You cannot observe while you are in combat!
	CannotObserveInCombat = SysMsg{Id: 1599}

	// Tomorrow's items will ALL be set to 0. Do you wish to continue?
	TomorrowItemZeroConfirm = SysMsg{Id: 1600}

	// Tomorrow's items will all be set to the same value as today's items. Do you wish to continue?
	TomorrowItemSameConfirm = SysMsg{Id: 1601}

	// Only a party leader can access the Command Channel.
	CommandChannelOnlyForPartyLeader = SysMsg{Id: 1602}

	// Only channel operator can give All Command.
	OnlyCommanderGiveCommand = SysMsg{Id: 1603}

	// While dressed in formal wear, you can't use items that require all skills and casting operations.
	CannotUseItemsSkillsWithFormalwear = SysMsg{Id: 1604}

	// * Here, you can buy only seeds of {{S1}} Manor.
	HereYouCanBuyOnlySeedsOfS1Manor = SysMsg{Id: 1605}

	// Congratulations - You've completed the third-class transfer quest!
	ThirdClassTransfer = SysMsg{Id: 1606}

	// {{S1}} adena has been withdrawn to pay for purchasing fees.
	S1AdenaHasBeenWithdrawnToPayForPurchasingFees = SysMsg{Id: 1607}

	// Due to insufficient adena you cannot buy another castle.
	InsufficientAdenaToBuyCastle = SysMsg{Id: 1608}

	// War has already been declared against that clan... but I'll make note that you really don't like them.
	WarAlreadyDeclared = SysMsg{Id: 1609}

	// Fool! You cannot declare war against your own clan!
	CannotDeclareAgainstOwnClan = SysMsg{Id: 1610}

	// Leader: {{C1}}
	PartyLeaderC1 = SysMsg{Id: 1611}

	// =====<War List>=====
	WarList = SysMsg{Id: 1612}

	// There is no clan listed on War List.
	NoClanOnWarList = SysMsg{Id: 1613}

	// You have joined a channel that was already open.
	JoinedChannelAlreadyOpen = SysMsg{Id: 1614}

	// The number of remaining parties is {{S1}} until a channel is activated
	S1PartiesRemainingUntilChannel = SysMsg{Id: 1615}

	// The Command Channel has been activated.
	CommandChannelActivated = SysMsg{Id: 1616}

	// You do not have the authority to use the Command Channel.
	CantUseCommandChannel = SysMsg{Id: 1617}

	// The ferry from Rune Harbor to Gludin Harbor has been delayed.
	FerryRuneGludinDelayed = SysMsg{Id: 1618}

	// The ferry from Gludin Harbor to Rune Harbor has been delayed.
	FerryGludinRuneDelayed = SysMsg{Id: 1619}

	// Arrived at Rune Harbor.
	ArrivedAtRune = SysMsg{Id: 1620}

	// Departure for Gludin Harbor will take place in five minutes!
	DepartureForGludin5Minutes = SysMsg{Id: 1621}

	// Departure for Gludin Harbor will take place in one minute!
	DepartureForGludin1Minute = SysMsg{Id: 1622}

	// Make haste! We will be departing for Gludin Harbor shortly...
	DepartureForGludinShortly = SysMsg{Id: 1623}

	// We are now departing for Gludin Harbor Hold on and enjoy the ride!
	DepartureForGludinNow = SysMsg{Id: 1624}

	// Departure for Rune Harbor will take place after anchoring for ten minutes.
	DepartureForRune10Minutes = SysMsg{Id: 1625}

	// Departure for Rune Harbor will take place in five minutes!
	DepartureForRune5Minutes = SysMsg{Id: 1626}

	// Departure for Rune Harbor will take place in one minute!
	DepartureForRune1Minute = SysMsg{Id: 1627}

	// Make haste! We will be departing for Gludin Harbor shortly...
	DepartureForGludinShortly2 = SysMsg{Id: 1628}

	// We are now departing for Rune Harbor Hold on and enjoy the ride!
	DepartureForRuneNow = SysMsg{Id: 1629}

	// The ferry from Rune Harbor will be arriving at Gludin Harbor in approximately 15 minutes.
	FerryFromRuneAtGludin15Minutes = SysMsg{Id: 1630}

	// The ferry from Rune Harbor will be arriving at Gludin Harbor in approximately 10 minutes.
	FerryFromRuneAtGludin10Minutes = SysMsg{Id: 1631}

	// The ferry from Rune Harbor will be arriving at Gludin Harbor in approximately 10 minutes.
	FerryFromRuneAtGludin5Minutes = SysMsg{Id: 1632}

	// The ferry from Rune Harbor will be arriving at Gludin Harbor in approximately 1 minute.
	FerryFromRuneAtGludin1Minute = SysMsg{Id: 1633}

	// The ferry from Gludin Harbor will be arriving at Rune Harbor in approximately 15 minutes.
	FerryFromGludinAtRune15Minutes = SysMsg{Id: 1634}

	// The ferry from Gludin Harbor will be arriving at Rune harbor in approximately 10 minutes.
	FerryFromGludinAtRune10Minutes = SysMsg{Id: 1635}

	// The ferry from Gludin Harbor will be arriving at Rune Harbor in approximately 10 minutes.
	FerryFromGludinAtRune5Minutes = SysMsg{Id: 1636}

	// The ferry from Gludin Harbor will be arriving at Rune Harbor in approximately 1 minute.
	FerryFromGludinAtRune1Minute = SysMsg{Id: 1637}

	// You cannot fish while using a recipe book, private manufacture or private store.
	CannotFishWhileUsingRecipeBook = SysMsg{Id: 1638}

	// Period {{S1}} of the Grand Olympiad Games has started!
	OlympiadPeriodS1HasStarted = SysMsg{Id: 1639}

	// Period {{S1}} of the Grand Olympiad Games has now ended.
	OlympiadPeriodS1HasEnded = SysMsg{Id: 1640}

	// Sharpen your swords, tighten the stitching in your armor, and make haste to a Grand Olympiad Manager! Battles in the Grand Olympiad Games are now taking place!
	TheOlympiadGameHasStarted = SysMsg{Id: 1641}

	// Much carnage has been left for the cleanup crew of the Olympiad Stadium. Battles in the Grand Olympiad Games are now over!
	TheOlympiadGameHasEnded = SysMsg{Id: 1642}

	// Current Location: {{S1}}, {{S2}}, {{S3}} (Dimensional Gap)
	LocDimensionalGapS1S2S3 = SysMsg{Id: 1643}

	// Play time is now accumulating.
	PlayTimeNowAccumulating = SysMsg{Id: 1649}

	// Due to high server traffic, your login attempt has failed. Please try again soon.
	TryLoginLater = SysMsg{Id: 1650}

	// The Grand Olympiad Games are not currently in progress.
	TheOlympiadGameIsNotCurrentlyInProgress = SysMsg{Id: 1651}

	// You are now recording gameplay.
	RecordingGameplayStart = SysMsg{Id: 1652}

	// Your recording has been successfully stored. ({{S1}})
	RecordingGameplayStopS1 = SysMsg{Id: 1653}

	// Your attempt to record the replay file has failed.
	RecordingGameplayFailed = SysMsg{Id: 1654}

	// You caught something smelly and scary, maybe you should throw it back!?
	YouCaughtSomethingSmellyThrowItBack = SysMsg{Id: 1655}

	// You have successfully traded the item with the NPC.
	SuccessfullyTradedWithNpc = SysMsg{Id: 1656}

	// {{C1}} has earned {{S2}} points in the Grand Olympiad Games.
	C1HasGainedS2OlympiadPoints = SysMsg{Id: 1657}

	// {{C1}} has lost {{S2}} points in the Grand Olympiad Games.
	C1HasLostS2OlympiadPoints = SysMsg{Id: 1658}

	// Current Location: {{S1}}, {{S2}}, {{S3}} (Cemetery of the Empire)
	LocCemetaryOfTheEmpireS1S2S3 = SysMsg{Id: 1659}

	// Channel Creator: {{C1}}.
	ChannelCreatorC1 = SysMsg{Id: 1660}

	// {{C1}} has obtained {{S3}} {{S2}}s.
	C1ObtainedS3S2S = SysMsg{Id: 1661}

	// The fish are no longer biting here because you've caught too many! Try fishing in another location.
	FishNoMoreBitingTryOtherLocation = SysMsg{Id: 1662}

	// The clan crest was successfully registered. Remember, only a clan that owns a clan hall or castle can have their crest displayed.
	ClanEmblemWasSuccessfullyRegistered = SysMsg{Id: 1663}

	// The fish is resisting your efforts to haul it in! Look at that bobber go!
	FishResistingLookBobbler = SysMsg{Id: 1664}

	// You've worn that fish out! It can't even pull the bobber under the water!
	YouWornFishOut = SysMsg{Id: 1665}

	// You have obtained +{{S1}} {{S2}}.
	ObtainedS1S2 = SysMsg{Id: 1666}

	// Lethal Strike!
	LethalStrike = SysMsg{Id: 1667}

	// Your lethal strike was successful!
	LethalStrikeSuccessful = SysMsg{Id: 1668}

	// There was nothing found inside of that.
	NothingInsideThat = SysMsg{Id: 1669}

	// Due to your Reeling and/or Pumping skill being three or more levels higher than your Fishing skill, a 50 damage penalty will be applied.
	ReelingPumping3LevelsHigherThanFishingPenalty = SysMsg{Id: 1670}

	// Your reeling was successful! (Mastery Penalty:{{S1}} )
	ReelingSuccessfulPenaltyS1 = SysMsg{Id: 1671}

	// Your pumping was successful! (Mastery Penalty:{{S1}} )
	PumpingSuccessfulPenaltyS1 = SysMsg{Id: 1672}

	// Your current record for this Grand Olympiad is {{S1}} match(es), {{S2}} win(s) and {{S3}} defeat(s). You have earned {{S4}} Olympiad Point(s).
	TheCurrentRecordForThisOlympiadSessionIsS1MatchesS2WinsS3DefeatsYouHaveEarnedS4OlympiadPoints = SysMsg{Id: 1673}

	// This command can only be used by a Noblesse.
	NoblesseOnly = SysMsg{Id: 1674}

	// A manor cannot be set up between 6 a.m. and 8 p.m.
	AManorCannotBeSetUpBetween6AmAnd8Pm = SysMsg{Id: 1675}

	// You do not have a servitor or pet and therefore cannot use the automatic-use function.
	NoServitorCannotAutomateUse = SysMsg{Id: 1676}

	// A cease-fire during a Clan War can not be called while members of your clan are engaged in battle.
	CantStopClanWarWhileInCombat = SysMsg{Id: 1677}

	// You have not declared a Clan War against the clan {{S1}}.
	NoClanWarAgainstClanS1 = SysMsg{Id: 1678}

	// Only the creator of a channel can issue a global command.
	OnlyChannelCreatorCanGlobalCommand = SysMsg{Id: 1679}

	// {{C1}} has declined the channel invitation.
	C1DeclinedChannelInvitation = SysMsg{Id: 1680}

	// Since {{C1}} did not respond, your channel invitation has failed.
	C1DidNotRespondChannelInvitationFailed = SysMsg{Id: 1681}

	// Only the creator of a channel can use the channel dismiss command.
	OnlyChannelCreatorCanDismiss = SysMsg{Id: 1682}

	// Only a party leader can leave a command channel.
	OnlyPartyLeaderCanLeaveChannel = SysMsg{Id: 1683}

	// A Clan War can not be declared against a clan that is being dissolved.
	NoClanWarAgainstDissolvingClan = SysMsg{Id: 1684}

	// You are unable to equip this item when your PK count is greater or equal to one.
	YouAreUnableToEquipThisItemWhenYourPkCountIsGreaterThanOrEqualToOne = SysMsg{Id: 1685}

	// Stones and mortar tumble to the earth - the castle wall has taken damage!
	CastleWallDamaged = SysMsg{Id: 1686}

	// This area cannot be entered while mounted atop of a Wyvern. You will be dismounted from your Wyvern if you do not leave!
	AreaCannotBeEnteredWhileMountedWyvern = SysMsg{Id: 1687}

	// You cannot enchant while operating a Private Store or Private Workshop.
	CannotEnchantWhileStore = SysMsg{Id: 1688}

	// {{C1}} is already registered on the class match waiting list.
	C1IsAlreadyRegisteredOnTheClassMatchWaitingList = SysMsg{Id: 1689}

	// {{C1}} is already registered on the waiting list for the non-class-limited individual match event.
	C1IsAlreadyRegisteredOnTheNonClassLimitedMatchWaitingList = SysMsg{Id: 1690}

	// {{C1}}% does not meet the participation requirements. You cannot participate in the Olympiad because your inventory slot exceeds 80%.
	C1CannotParticipateInOlympiadInventorySlotExceeds80Percent = SysMsg{Id: 1691}

	// {{C1}}% does not meet the participation requirements. You cannot participate in the Olympiad because you have changed to your sub-class.
	C1CannotParticipateInOlympiadWhileChangedToSubClass = SysMsg{Id: 1692}

	// You may not observe a Grand Olympiad Games match while you are on the waiting list.
	WhileYouAreOnTheWaitingListYouAreNotAllowedToWatchTheGame = SysMsg{Id: 1693}

	// Only a clan leader that is a Noblesse can view the Siege War Status window during a siege war.
	OnlyNoblesseLeaderCanViewSiegeStatusWindow = SysMsg{Id: 1694}

	// You can only use that during a Siege War!
	OnlyDuringSiege = SysMsg{Id: 1695}

	// Your accumulated play time is {{S1}}.
	AccumulatedPlayTimeIsS1 = SysMsg{Id: 1696}

	// Your accumulated play time has reached Fatigue level, so you will receive experience or item drops at only 50 percent [...]
	AccumulatedPlayTimeWarning1 = SysMsg{Id: 1697}

	// Your accumulated play time has reached Ill-health level, so you will no longer gain experience or item drops. [...}
	AccumulatedPlayTimeWarning2 = SysMsg{Id: 1698}

	// You cannot dismiss a party member by force.
	CannotDismissPartyMember = SysMsg{Id: 1699}

	// You don't have enough spiritshots needed for a pet/servitor.
	NotEnoughSpirithotsForPet = SysMsg{Id: 1700}

	// You don't have enough soulshots needed for a pet/servitor.
	NotEnoughSoulshotsForPet = SysMsg{Id: 1701}

	// {{S1}} is using a third party program.
	S1UsingThirdPartyProgram = SysMsg{Id: 1702}

	// The previous investigated user is not using a third party program
	NotUsingThirdPartyProgram = SysMsg{Id: 1703}

	// Please close the setup window for your private manufacturing store or private store, and try again.
	CloseStoreWindowAndTryAgain = SysMsg{Id: 1704}

	// PC Bang Points acquisition period. Points acquisition period left {{S1}} hour.
	PcpointAcquisitionPeriod = SysMsg{Id: 1705}

	// PC Bang Points use period. Points acquisition period left {{S1}} hour.
	PcpointUsePeriod = SysMsg{Id: 1706}

	// You acquired {{S1}} PC Bang Point.
	AcquiredS1Pcpoint = SysMsg{Id: 1707}

	// Double points! You acquired {{S1}} PC Bang Point.
	AcquiredS1PcpointDouble = SysMsg{Id: 1708}

	// You are using {{S1}} point.
	UsingS1Pcpoint = SysMsg{Id: 1709}

	// You are short of accumulated points.
	ShortOfAccumulatedPoints = SysMsg{Id: 1710}

	// PC Bang Points use period has expired.
	PcpointUsePeriodExpired = SysMsg{Id: 1711}

	// The PC Bang Points accumulation period has expired.
	PcpointAccumulationPeriodExpired = SysMsg{Id: 1712}

	// The games may be delayed due to an insufficient number of players waiting.
	GamesDelayed = SysMsg{Id: 1713}

	// Current Location: {{S1}}, {{S2}}, {{S3}} (Near the Town of Schuttgart)
	LocSchuttgartS1S2S3 = SysMsg{Id: 1714}

	// This is a Peaceful Zone - PvP is not allowed in this area.
	PeacefulZone = SysMsg{Id: 1715}

	// Altered Zone
	AlteredZone = SysMsg{Id: 1716}

	// Siege War Zone - A siege is currently in progress in this area. If a character dies in this zone, their resurrection ability may be restricted.
	SiegeZone = SysMsg{Id: 1717}

	// General Field
	GeneralZone = SysMsg{Id: 1718}

	// Seven Signs Zone - Although a character's level may increase while in this area, HP and MP will not be regenerated.
	SevensignsZone = SysMsg{Id: 1719}

	// ---
	Unknown1 = SysMsg{Id: 1720}

	// Combat Zone
	CombatZone = SysMsg{Id: 1721}

	// Please enter the name of the item you wish to search for.
	EnterItemNameSearch = SysMsg{Id: 1722}

	// Please take a moment to provide feedback about the petition service.
	PleaseProvidePetitionFeedback = SysMsg{Id: 1723}

	// A servitor whom is engaged in battle cannot be de-activated.
	ServitorNotReturnInBattle = SysMsg{Id: 1724}

	// You have earned {{S1}} raid point(s).
	EarnedS1RaidPoints = SysMsg{Id: 1725}

	// {{S1}} has disappeared because its time period has expired.
	S1PeriodExpiredDisappeared = SysMsg{Id: 1726}

	// {{C1}} has invited you to a party room. Do you accept?
	C1InvitedYouToPartyRoomConfirm = SysMsg{Id: 1727}

	// The recipient of your invitation did not accept the party matching invitation.
	PartyMatchingRequestNoResponse = SysMsg{Id: 1728}

	// You cannot join a Command Channel while teleporting.
	NotJoinChannelWhileTeleporting = SysMsg{Id: 1729}

	// To establish a Clan Academy, your clan must be Level 5 or higher.
	YouDoNotMeetCriteriaInOrderToCreateAClanAcademy = SysMsg{Id: 1730}

	// Only the leader can create a Clan Academy.
	OnlyLeaderCanCreateAcademy = SysMsg{Id: 1731}

	// To create a Clan Academy, a Blood Mark is needed.
	NeedBloodmarkForAcademy = SysMsg{Id: 1732}

	// You do not have enough adena to create a Clan Academy.
	NeedAdenaForAcademy = SysMsg{Id: 1733}

	// To join a Clan Academy, characters must be Level 40 or below, not belong another clan and not yet completed their 2nd class transfer.
	AcademyRequirements = SysMsg{Id: 1734}

	// {{S1}} does not meet the requirements to join a Clan Academy.
	S1DoesnotMeetRequirementsToJoinAcademy = SysMsg{Id: 1735}

	// The Clan Academy has reached its maximum enrollment.
	AcademyMaximum = SysMsg{Id: 1736}

	// Your clan has not established a Clan Academy but is eligible to do so.
	ClanCanCreateAcademy = SysMsg{Id: 1737}

	// Your clan has already established a Clan Academy.
	ClanHasAlreadyEstablishedAClanAcademy = SysMsg{Id: 1738}

	// Would you like to create a Clan Academy?
	ClanAcademyCreateConfirm = SysMsg{Id: 1739}

	// Please enter the name of the Clan Academy.
	AcademyCreateEnterName = SysMsg{Id: 1740}

	// Congratulations! The {{S1}}'s Clan Academy has been created.
	TheS1sClanAcademyHasBeenCreated = SysMsg{Id: 1741}

	// A message inviting {{S1}} to join the Clan Academy is being sent.
	AcademyInvitationSentToS1 = SysMsg{Id: 1742}

	// To open a Clan Academy, the leader of a Level 5 clan or above must pay XX Proofs of Blood or a certain amount of adena.
	OpenAcademyConditions = SysMsg{Id: 1743}

	// There was no response to your invitation to join the Clan Academy, so the invitation has been rescinded.
	AcademyJoinNoResponse = SysMsg{Id: 1744}

	// The recipient of your invitation to join the Clan Academy has declined.
	AcademyJoinDecline = SysMsg{Id: 1745}

	// You have already joined a Clan Academy.
	AlreadyJoinedAcademy = SysMsg{Id: 1746}

	// {{S1}} has sent you an invitation to join the Clan Academy belonging to the {{S2}} clan. Do you accept?
	JoinAcademyRequestByS1ForClanS2 = SysMsg{Id: 1747}

	// Clan Academy member {{S1}} has successfully completed the 2nd class transfer and obtained {{S2}} Clan Reputation points.
	ClanMemberGraduatedFromAcademy = SysMsg{Id: 1748}

	// Congratulations! You will now graduate from the Clan Academy and leave your current clan. As a graduate of the academy, you can immediately join a clan as a regular member without being subject to any penalties.
	AcademyMembershipTerminated = SysMsg{Id: 1749}

	// {{C1}}% does not meet the participation requirements. The owner of {{S2}} cannot participate in the Olympiad.
	C1CannotJoinOlympiadPossessingS2 = SysMsg{Id: 1750}

	// The Grand Master has given you a commemorative item.
	GrandMasterCommemorativeItem = SysMsg{Id: 1751}

	// Since the clan has received a graduate of the Clan Academy, it has earned {{S1}} points towards its reputation score.
	MemberGraduatedEarnedS1Repu = SysMsg{Id: 1752}

	// The clan leader has decreed that that particular privilege cannot be granted to a Clan Academy member.
	CantTransferPrivilegeToAcademyMember = SysMsg{Id: 1753}

	// That privilege cannot be granted to a Clan Academy member.
	RightCantTransferredToAcademyMember = SysMsg{Id: 1754}

	// {{S2}} has been designated as the apprentice of clan member {{S1}}.
	S2HasBeenDesignatedAsApprenticeOfClanMemberS1 = SysMsg{Id: 1755}

	// Your apprentice, {{S1}}, has logged in.
	YourApprenticeS1HasLoggedIn = SysMsg{Id: 1756}

	// Your apprentice, {{C1}}, has logged out.
	YourApprenticeC1HasLoggedOut = SysMsg{Id: 1757}

	// Your sponsor, {{C1}}, has logged in.
	YourSponsorC1HasLoggedIn = SysMsg{Id: 1758}

	// Your sponsor, {{C1}}, has logged out.
	YourSponsorC1HasLoggedOut = SysMsg{Id: 1759}

	// Clan member {{C1}}'s name title has been changed to $2.
	ClanMemberC1TitleChangedToS2 = SysMsg{Id: 1760}

	// Clan member {{C1}}'s privilege level has been changed to {{S2}}.
	ClanMemberC1PrivilegeChangedToS2 = SysMsg{Id: 1761}

	// You do not have the right to dismiss an apprentice.
	YouDoNotHaveTheRightToDismissAnApprentice = SysMsg{Id: 1762}

	// {{S2}}, clan member {{C1}}'s apprentice, has been removed.
	S2ClanMemberC1ApprenticeHasBeenRemoved = SysMsg{Id: 1763}

	// This item can only be worn by a member of the Clan Academy.
	EquipOnlyForAcademy = SysMsg{Id: 1764}

	// As a graduate of the Clan Academy, you can no longer wear this item.
	EquipNotForGraduates = SysMsg{Id: 1765}

	// An application to join the clan has been sent to {{C1}} in {{S2}}.
	ClanJoinApplicationSentToC1InS2 = SysMsg{Id: 1766}

	// An application to join the clan Academy has been sent to {{C1}}.
	AcademyJoinApplicationSentToC1 = SysMsg{Id: 1767}

	// {{C1}} has invited you to join the Clan Academy of {{S2}} clan. Would you like to join?
	JoinRequestByC1ToClanS2Academy = SysMsg{Id: 1768}

	// {{C1}} has sent you an invitation to join the {{S3}} Order of Knights under the {{S2}} clan. Would you like to join?
	JoinRequestByC1ToOrderOfKnightsS3UnderClanS2 = SysMsg{Id: 1769}

	// The clan's reputation score has dropped below 0. The clan may face certain penalties as a result.
	ClanRepu0MayFacePenalties = SysMsg{Id: 1770}

	// Now that your clan level is above Level 5, it can accumulate clan reputation points.
	ClanCanAccumulateClanReputationPoints = SysMsg{Id: 1771}

	// Since your clan was defeated in a siege, {{S1}} points have been deducted from your clan's reputation score and given to the opposing clan.
	ClanWasDefeatedInSiegeAndLostS1ReputationPoints = SysMsg{Id: 1772}

	// Since your clan emerged victorious from the siege, {{S1}} points have been added to your clan's reputation score.
	ClanVictoriousInSiegeAndGainedS1ReputationPoints = SysMsg{Id: 1773}

	// Your clan's newly acquired contested clan hall has added {{S1}} points to your clan's reputation score.
	ClanAcquiredContestedClanHallAndS1ReputationPoints = SysMsg{Id: 1774}

	// Clan member {{C1}} was an active member of the highest-ranked party in the Festival of Darkness. {{S2}} points have been added to your clan's reputation score.
	ClanMemberC1WasInHighestRankedPartyInFestivalOfDarknessAndGainedS2Reputation = SysMsg{Id: 1775}

	// Clan member {{C1}} was named a hero. $2s points have been added to your clan's reputation score.
	ClanMemberC1BecameHeroAndGainedS2ReputationPoints = SysMsg{Id: 1776}

	// You have successfully completed a clan quest. {{S1}} points have been added to your clan's reputation score.
	ClanQuestCompletedAndS1PointsGained = SysMsg{Id: 1777}

	// An opposing clan has captured your clan's contested clan hall. {{S1}} points have been deducted from your clan's reputation score.
	OpposingClanCapturedClanHallAndYourClanLosesS1Points = SysMsg{Id: 1778}

	// After losing the contested clan hall, 300 points have been deducted from your clan's reputation score.
	ClanLostContestedClanHallAnd300Points = SysMsg{Id: 1779}

	// Your clan has captured your opponent's contested clan hall. {{S1}} points have been deducted from your opponent's clan reputation score.
	ClanCapturedContestedClanHallAndS1PointsDeductedFromOpponent = SysMsg{Id: 1780}

	// Your clan has added $1s points to its clan reputation score.
	ClanAddedS1SPointsToReputationScore = SysMsg{Id: 1781}

	// Your clan member {{C1}} was killed. {{S2}} points have been deducted from your clan's reputation score and added to your opponent's clan reputation score.
	ClanMemberC1WasKilledAndS2PointsDeductedFromReputation = SysMsg{Id: 1782}

	// For killing an opposing clan member, {{S1}} points have been deducted from your opponents' clan reputation score.
	ForKillingOpposingMemberS1PointsWereDeductedFromOpponents = SysMsg{Id: 1783}

	// Your clan has failed to defend the castle. {{S1}} points have been deducted from your clan's reputation score and added to your opponents'.
	YourClanFailedToDefendCastleAndS1PointsLostAndAddedToOpponent = SysMsg{Id: 1784}

	// The clan you belong to has been initialized. {{S1}} points have been deducted from your clan reputation score.
	YourClanHasBeenInitializedAndS1PointsLost = SysMsg{Id: 1785}

	// Your clan has failed to defend the castle. {{S1}} points have been deducted from your clan's reputation score.
	YourClanFailedToDefendCastleAndS1PointsLost = SysMsg{Id: 1786}

	// {{S1}} points have been deducted from the clan's reputation score.
	S1DeductedFromClanRep = SysMsg{Id: 1787}

	// The clan skill {{S1}} has been added.
	ClanSkillS1Added = SysMsg{Id: 1788}

	// Since the Clan Reputation Score has dropped to 0 or lower, your clan skill(s) will be de-activated.
	ReputationPoints0OrLowerClanSkillsDeactivated = SysMsg{Id: 1789}

	// The conditions necessary to increase the clan's level have not been met.
	FailedToIncreaseClanLevel = SysMsg{Id: 1790}

	// The conditions necessary to create a military unit have not been met.
	YouDoNotMeetCriteriaInOrderToCreateAMilitaryUnit = SysMsg{Id: 1791}

	// Please assign a manager for your new Order of Knights.
	AssignManagerForOrderOfKnights = SysMsg{Id: 1792}

	// {{C1}} has been selected as the captain of {{S2}}.
	C1HasBeenSelectedAsCaptainOfS2 = SysMsg{Id: 1793}

	// The Knights of {{S1}} have been created.
	TheKnightsOfS1HaveBeenCreated = SysMsg{Id: 1794}

	// The Royal Guard of {{S1}} have been created.
	TheRoyalGuardOfS1HaveBeenCreated = SysMsg{Id: 1795}

	// Your account has been suspended ...
	IllegalUse17 = SysMsg{Id: 1796}

	// {{C1}} has been promoted to {{S2}}.
	C1PromotedToS2 = SysMsg{Id: 1797}

	// Clan lord privileges have been transferred to {{C1}}.
	ClanLeaderPrivilegesHaveBeenTransferredToC1 = SysMsg{Id: 1798}

	// We are searching for BOT users. Please try again later.
	SearchingForBotUsersTryAgainLater = SysMsg{Id: 1799}

	// User {{C1}} has a history of using BOT.
	C1HistoryUsingBot = SysMsg{Id: 1800}

	// The attempt to sell has failed.
	SellAttemptFailed = SysMsg{Id: 1801}

	// The attempt to trade has failed.
	TradeAttemptFailed = SysMsg{Id: 1802}

	// The request to participate in the game cannot be made starting from 10 minutes before the end of the game.
	GameRequestCannotBeMade = SysMsg{Id: 1803}

	// Your account has been suspended ...
	IllegalUse18 = SysMsg{Id: 1804}

	// Your account has been suspended ...
	IllegalUse19 = SysMsg{Id: 1805}

	// Your account has been suspended ...
	IllegalUse20 = SysMsg{Id: 1806}

	// Your account has been suspended ...
	IllegalUse21 = SysMsg{Id: 1807}

	// Your account has been suspended ...
	IllegalUse22 = SysMsg{Id: 1808}

	// Your account must be verified. For information on verification procedures, please visit the PlayNC website (http://us.ncsoft.com/support/).
	AccountMustVerified = SysMsg{Id: 1809}

	// The refuse invitation state has been activated.
	RefuseInvitationActivated = SysMsg{Id: 1810}

	// Since the refuse invitation state is currently activated, no invitation can be made
	RefuseInvitationCurrentlyActive = SysMsg{Id: 1812}

	// {{S1}} has {{S2}} hour(s) of usage time remaining.
	ThereIsS1HourAndS2MinuteLeftOfTheFixedUsageTime = SysMsg{Id: 1813}

	// {{S1}} has {{S2}} minute(s) of usage time remaining.
	S2MinuteOfUsageTimeAreLeftForS1 = SysMsg{Id: 1814}

	// {{S2}} was dropped in the {{S1}} region.
	S2WasDroppedInTheS1Region = SysMsg{Id: 1815}

	// The owner of {{S2}} has appeared in the {{S1}} region.
	TheOwnerOfS2HasAppearedInTheS1Region = SysMsg{Id: 1816}

	// {{S2}}'s owner has logged into the {{S1}} region.
	S2OwnerHasLoggedInIntoTheS1Region = SysMsg{Id: 1817}

	// {{S1}} has disappeared.
	S1HasDisappeared = SysMsg{Id: 1818}

	// An evil is pulsating from {{S2}} in {{S1}}.
	EvilFromS2InS1 = SysMsg{Id: 1819}

	// {{S1}} is currently asleep.
	S1CurrentlySleep = SysMsg{Id: 1820}

	// {{S2}}'s evil presence is felt in {{S1}}.
	S2EvilPresenceFeltInS1 = SysMsg{Id: 1821}

	// {{S1}} has been sealed.
	S1Sealed = SysMsg{Id: 1822}

	// The registration period for a clan hall war has ended.
	ClanhallWarRegistrationPeriodEnded = SysMsg{Id: 1823}

	// You have been registered for a clan hall war. Please move to the left side of the clan hall's arena and get ready.
	RegisteredForClanhallWar = SysMsg{Id: 1824}

	// You have failed in your attempt to register for the clan hall war. Please try again.
	ClanhallWarRegistrationFailed = SysMsg{Id: 1825}

	// In {{S1}} minute(s), the game will begin. All players must hurry and move to the left side of the clan hall's arena.
	ClanhallWarBeginsInS1Minutes = SysMsg{Id: 1826}

	// In {{S1}} minute(s), the game will begin. All players must, please enter the arena now
	ClanhallWarBeginsInS1MinutesEnterNow = SysMsg{Id: 1827}

	// In {{S1}} seconds(s), the game will begin.
	ClanhallWarBeginsInS1Seconds = SysMsg{Id: 1828}

	// The Command Channel is full.
	CommandChannelFull = SysMsg{Id: 1829}

	// {{C1}} is not allowed to use the party room invite command. Please update the waiting list.
	C1NotAllowedInviteToPartyRoom = SysMsg{Id: 1830}

	// {{C1}} does not meet the conditions of the party room. Please update the waiting list.
	C1NotMeetConditionsForPartyRoom = SysMsg{Id: 1831}

	// Only a room leader may invite others to a party room.
	OnlyRoomLeaderCanInvite = SysMsg{Id: 1832}

	// All of {{S1}} will be dropped. Would you like to continue?
	ConfirmDropAllOfS1 = SysMsg{Id: 1833}

	// The party room is full. No more characters can be invitet in
	PartyRoomFull = SysMsg{Id: 1834}

	// {{S1}} is full and cannot accept additional clan members at this time.
	S1ClanIsFull = SysMsg{Id: 1835}

	// You cannot join a Clan Academy because you have successfully completed your 2nd class transfer.
	CannotJoinAcademyAfter2NdOccupation = SysMsg{Id: 1836}

	// {{C1}} has sent you an invitation to join the {{S3}} Royal Guard under the {{S2}} clan. Would you like to join?
	C1SentInvitationToRoyalGuardS3OfClanS2 = SysMsg{Id: 1837}

	// 1. The coupon an be used once per character.
	CouponOncePerCharacter = SysMsg{Id: 1838}

	// 2. A used serial number may not be used again.
	SerialMayUsedOnce = SysMsg{Id: 1839}

	// 3. If you enter the incorrect serial number more than 5 times, you may use it again after a certain amount of time passes.
	SerialInputIncorrect = SysMsg{Id: 1840}

	// The clan hall war has been cancelled. Not enough clans have registered.
	ClanhallWarCancelled = SysMsg{Id: 1841}

	// {{C1}} wishes to summon you from {{S2}}. Do you accept?
	C1WishesToSummonYouFromS2DoYouAccept = SysMsg{Id: 1842}

	// {{C1}} is engaged in combat and cannot be summoned.
	C1IsEngagedInCombatAndCannotBeSummoned = SysMsg{Id: 1843}

	// {{C1}} is dead at the moment and cannot be summoned.
	C1IsDeadAtTheMomentAndCannotBeSummoned = SysMsg{Id: 1844}

	// Hero weapons cannot be destroyed.
	HeroWeaponsCantDestroyed = SysMsg{Id: 1845}

	// You are too far away from the Fenrir to mount it.
	TooFarAwayFromFenrirToMount = SysMsg{Id: 1846}

	// You caught a fish {{S1}} in length.
	CaughtFishS1Length = SysMsg{Id: 1847}

	// Because of the size of fish caught, you will be registered in the ranking
	RegisteredInFishSizeRanking = SysMsg{Id: 1848}

	// All of {{S1}} will be discarded. Would you like to continue?
	ConfirmDiscardAllOfS1 = SysMsg{Id: 1849}

	// The Captain of the Order of Knights cannot be appointed.
	CaptainOfOrderOfKnightsCannotBeAppointed = SysMsg{Id: 1850}

	// The Captain of the Royal Guard cannot be appointed.
	CaptainOfRoyalGuardCannotBeAppointed = SysMsg{Id: 1851}

	// The attempt to acquire the skill has failed because of an insufficient Clan Reputation Score.
	AcquireSkillFailedBadClanRepScore = SysMsg{Id: 1852}

	// Quantity items of the same type cannot be exchanged at the same time
	CantExchangeQuantityItemsOfSameType = SysMsg{Id: 1853}

	// The item was converted successfully.
	ItemConvertedSuccessfully = SysMsg{Id: 1854}

	// Another military unit is already using that name. Please enter a different name.
	AnotherMilitaryUnitIsAlreadyUsingThatName = SysMsg{Id: 1855}

	// Since your opponent is now the owner of {{S1}}, the Olympiad has been cancelled.
	OpponentPossessesS1OlympiadCancelled = SysMsg{Id: 1856}

	// {{C1}} is the owner of {{S2}} and cannot participate in the Olympiad.
	C1OwnsS2AndCannotParticipateInOlympiad = SysMsg{Id: 1857}

	// {{C1}} is currently dead and cannot participate in the Olympiad.
	C1CannotParticipateOlympiadWhileDead = SysMsg{Id: 1858}

	// You exceeded the quantity that can be moved at one time.
	ExceededQuantityForMoved = SysMsg{Id: 1859}

	// The Clan Reputation Score is too low.
	TheClanReputationScoreIsTooLow = SysMsg{Id: 1860}

	// The clan's crest has been deleted.
	ClanCrestHasBeenDeleted = SysMsg{Id: 1861}

	// Clan skills will now be activated since the clan's reputation score is 0 or higher.
	ClanSkillsWillBeActivatedSinceReputationIs0OrHigher = SysMsg{Id: 1862}

	// {{C1}} purchased a clan item, reducing the Clan Reputation by {{S2}} points.
	C1PurchasedClanItemReducingS2RepuPoints = SysMsg{Id: 1863}

	// Your pet/servitor is unresponsive and will not obey any orders.
	PetRefusingOrder = SysMsg{Id: 1864}

	// Your pet/servitor is currently in a state of distress.
	PetInStateOfDistress = SysMsg{Id: 1865}

	// MP was reduced by {{S1}}.
	MpReducedByS1 = SysMsg{Id: 1866}

	// Your opponent's MP was reduced by {{S1}}.
	YourOpponentsMpWasReducedByS1 = SysMsg{Id: 1867}

	// You cannot exchange an item while it is being used.
	CannotExchanceUsedItem = SysMsg{Id: 1868}

	// {{C1}} has granted the Command Channel's master party the privilege of item looting.
	C1GrantedMasterPartyLootingRights = SysMsg{Id: 1869}

	// A Command Channel with looting rights already exists.
	CommandChannelWithLootingRightsExists = SysMsg{Id: 1870}

	// Do you want to dismiss {{C1}} from the clan?
	ConfirmDismissC1FromClan = SysMsg{Id: 1871}

	// You have {{S1}} hour(s) and {{S2}} minute(s) left.
	S1HoursS2MinutesLeft = SysMsg{Id: 1872}

	// There are {{S1}} hour(s) and {{S2}} minute(s) left in the fixed use time for this PC Cafe.
	S1HoursS2MinutesLeftForThisPccafe = SysMsg{Id: 1873}

	// There are {{S1}} minute(s) left for this individual user.
	S1MinutesLeftForThisUser = SysMsg{Id: 1874}

	// There are {{S1}} minute(s) left in the fixed use time for this PC Cafe.
	S1MinutesLeftForThisPccafe = SysMsg{Id: 1875}

	// Do you want to leave {{S1}} clan?
	ConfirmLeaveS1Clan = SysMsg{Id: 1876}

	// The game will end in {{S1}} minutes.
	GameWillEndInS1Minutes = SysMsg{Id: 1877}

	// The game will end in {{S1}} seconds.
	GameWillEndInS1Seconds = SysMsg{Id: 1878}

	// In {{S1}} minute(s), you will be teleported outside of the game arena.
	InS1MinutesTeleportedOutsideOfGameArena = SysMsg{Id: 1879}

	// In {{S1}} seconds(s), you will be teleported outside of the game arena.
	InS1SecondsTeleportedOutsideOfGameArena = SysMsg{Id: 1880}

	// The preliminary match will begin in {{S1}} second(s). Prepare yourself.
	PreliminaryMatchBeginInS1Seconds = SysMsg{Id: 1881}

	// Characters cannot be created from this server.
	CharactersNotCreatedFromThisServer = SysMsg{Id: 1882}

	// There are no offerings I own or I made a bid for.
	NoOfferingsOwnOrMadeBidFor = SysMsg{Id: 1883}

	// Enter the PC Room coupon serial number.
	EnterPcroomSerialNumber = SysMsg{Id: 1884}

	// This serial number cannot be entered. Please try again in minute(s).
	SerialNumberCantEntered = SysMsg{Id: 1885}

	// This serial has already been used.
	SerialNumberAlreadyUsed = SysMsg{Id: 1886}

	// Invalid serial number. Your attempt to enter the number has failed time(s). You will be allowed to make more attempt(s).
	SerialNumberEnteringFailed = SysMsg{Id: 1887}

	// Invalid serial number. Your attempt to enter the number has failed 5 time(s). Please try again in 4 hours.
	SerialNumberEnteringFailed5Times = SysMsg{Id: 1888}

	// Congratulations! You have received {{S1}}.
	CongratulationsReceivedS1 = SysMsg{Id: 1889}

	// Since you have already used this coupon, you may not use this serial number.
	AlreadyUsedCouponNotUseSerialNumber = SysMsg{Id: 1890}

	// You may not use items in a private store or private work shop.
	NotUseItemsInPrivateStore = SysMsg{Id: 1891}

	// The replay file for the previous version cannot be played.
	ReplayFilePreviousVersionCantPlayed = SysMsg{Id: 1892}

	// This file cannot be replayed.
	FileCantReplayed = SysMsg{Id: 1893}

	// A sub-class cannot be created or changed while you are over your weight limit.
	NotSubclassWhileOverweight = SysMsg{Id: 1894}

	// {{C1}} is in an area which blocks summoning.
	C1InSummonBlockingArea = SysMsg{Id: 1895}

	// {{C1}} has already been summoned.
	C1AlreadySummoned = SysMsg{Id: 1896}

	// {{S1}} is required for summoning.
	S1RequiredForSummoning = SysMsg{Id: 1897}

	// {{C1}} is currently trading or operating a private store and cannot be summoned.
	C1CurrentlyTradingOrOperatingPrivateStoreAndCannotBeSummoned = SysMsg{Id: 1898}

	// Your target is in an area which blocks summoning.
	YourTargetIsInAnAreaWhichBlocksSummoning = SysMsg{Id: 1899}

	// {{C1}} has entered the party room.
	C1EnteredPartyRoom = SysMsg{Id: 1900}

	// {{C1}} has invited you to enter the party room.
	C1InvitedYouToPartyRoom = SysMsg{Id: 1901}

	// Incompatible item grade. This item cannot be used.
	IncompatibleItemGrade = SysMsg{Id: 1902}

	// Those of you who have requested NCOTP should run NCOTP by using your cell phone to get the NCOTP password and enter it within 1 minute. If you have not requested NCOTP, leave this field blank and click the Login button.
	Ncotp = SysMsg{Id: 1903}

	// A sub-class may not be created or changed while a servitor or pet is summoned.
	CantSubclassWithSummonedServitor = SysMsg{Id: 1904}

	// {{S2}} of {{S1}} will be replaced with {{S4}} of {{S3}}.
	S2OfS1WillReplacedWithS4OfS3 = SysMsg{Id: 1905}

	// Select the combat unit to transfer to.
	SelectCombatUnit = SysMsg{Id: 1906}

	// Select the character who will replace the current character.
	SelectCharacterWhoWill = SysMsg{Id: 1907}

	// {{C1}} in a state which prevents summoning.
	C1StateForbidsSummoning = SysMsg{Id: 1908}

	// ==< List of Academy Graduates During the Past Week >==
	AcademyListHeader = SysMsg{Id: 1909}

	// Graduates: {{C1}}.
	GraduatesC1 = SysMsg{Id: 1910}

	// You cannot summon players who are currently participating in the Grand Olympiad.
	YouCannotSummonPlayersWhoAreInOlympiad = SysMsg{Id: 1911}

	// Only those requesting NCOTP should make an entry into this field.
	Ncotp2 = SysMsg{Id: 1912}

	// The remaining recycle time for {{S1}} is {{S2}} minute(s).
	TimeForS1IsS2MinutesRemaining = SysMsg{Id: 1913}

	// The remaining recycle time for {{S1}} is {{S2}} seconds(s).
	TimeForS1IsS2SecondsRemaining = SysMsg{Id: 1914}

	// The game will end in {{S1}} second(s).
	GameEndsInS1Seconds = SysMsg{Id: 1915}

	// Your Death Penalty is now level {{S1}}.
	DeathPenaltyLevelS1Added = SysMsg{Id: 1916}

	// Your Death Penalty has been lifted.
	DeathPenaltyLifted = SysMsg{Id: 1917}

	// Your pet is too high level to control.
	PetTooHighToControl = SysMsg{Id: 1918}

	// The Grand Olympiad registration period has ended.
	OlympiadRegistrationPeriodEnded = SysMsg{Id: 1919}

	// Your account is currently inactive because you have not logged into the game for some time. You may reactivate your account by visiting the PlayNC website (http://www.plaync.com/us/support/).
	AccountInactivity = SysMsg{Id: 1920}

	// {{S2}} hour(s) and {{S3}} minute(s) have passed since {{S1}} has killed.
	S2HoursS3MinutesSinceS1Killed = SysMsg{Id: 1921}

	// Because {{S1}} has failed to kill for one full day, it has expired.
	S1FailedKillingExpired = SysMsg{Id: 1922}

	// Court Magician: The portal has been created!
	CourtMagicianCreatedPortal = SysMsg{Id: 1923}

	// Current Location: {{S1}}, {{S2}}, {{S3}} (Near the Primeval Isle)
	LocPrimevalIsleS1S2S3 = SysMsg{Id: 1924}

	// Due to the affects of the Seal of Strife, it is not possible to summon at this time.
	SealOfStrifeForbidsSummoning = SysMsg{Id: 1925}

	// There is no opponent to receive your challenge for a duel.
	ThereIsNoOpponentToReceiveYourChallengeForADuel = SysMsg{Id: 1926}

	// {{C1}} has been challenged to a duel.
	C1HasBeenChallengedToADuel = SysMsg{Id: 1927}

	// {{C1}}'s party has been challenged to a duel.
	C1PartyHasBeenChallengedToADuel = SysMsg{Id: 1928}

	// {{C1}} has accepted your challenge to a duel. The duel will begin in a few moments.
	C1HasAcceptedYourChallengeToADuelTheDuelWillBeginInAFewMoments = SysMsg{Id: 1929}

	// You have accepted {{C1}}'s challenge to a duel. The duel will begin in a few moments.
	YouHaveAcceptedC1ChallengeToADuelTheDuelWillBeginInAFewMoments = SysMsg{Id: 1930}

	// {{C1}} has declined your challenge to a duel.
	C1HasDeclinedYourChallengeToADuel = SysMsg{Id: 1931}

	// {{C1}} has declined your challenge to a duel.
	C1HasDeclinedYourChallengeToADuel2 = SysMsg{Id: 1932}

	// You have accepted {{C1}}'s challenge to a party duel. The duel will begin in a few moments.
	YouHaveAcceptedC1ChallengeToAPartyDuelTheDuelWillBeginInAFewMoments = SysMsg{Id: 1933}

	// {{S1}} has accepted your challenge to duel against their party. The duel will begin in a few moments.
	S1HasAcceptedYourChallengeToDuelAgainstTheirPartyTheDuelWillBeginInAFewMoments = SysMsg{Id: 1934}

	// {{C1}} has declined your challenge to a party duel.
	C1HasDeclinedYourChallengeToAPartyDuel = SysMsg{Id: 1935}

	// The opposing party has declined your challenge to a duel.
	TheOpposingPartyHasDeclinedYourChallengeToADuel = SysMsg{Id: 1936}

	// Since the person you challenged is not currently in a party, they cannot duel against your party.
	SinceThePersonYouChallengedIsNotCurrentlyInAPartyTheyCannotDuelAgainstYourParty = SysMsg{Id: 1937}

	// {{C1}} has challenged you to a duel.
	C1HasChallengedYouToADuel = SysMsg{Id: 1938}

	// {{C1}}'s party has challenged your party to a duel.
	C1PartyHasChallengedYourPartyToADuel = SysMsg{Id: 1939}

	// You are unable to request a duel at this time.
	YouAreUnableToRequestADuelAtThisTime = SysMsg{Id: 1940}

	// This is no suitable place to challenge anyone or party to a duel.
	NoPlaceForDuel = SysMsg{Id: 1941}

	// The opposing party is currently unable to accept a challenge to a duel.
	TheOpposingPartyIsCurrentlyUnableToAcceptAChallengeToADuel = SysMsg{Id: 1942}

	// The opposing party is currently not in a suitable location for a duel.
	TheOpposingPartyIsAtBadLocationForADuel = SysMsg{Id: 1943}

	// In a moment, you will be transported to the site where the duel will take place.
	InAMomentYouWillBeTransportedToTheSiteWhereTheDuelWillTakePlace = SysMsg{Id: 1944}

	// The duel will begin in {{S1}} second(s).
	TheDuelWillBeginInS1Seconds = SysMsg{Id: 1945}

	// {{C1}} has challenged you to a duel. Will you accept?
	C1ChallengedYouToADuel = SysMsg{Id: 1946}

	// {{C1}}'s party has challenged your party to a duel. Will you accept?
	C1ChallengedYouToAPartyDuel = SysMsg{Id: 1947}

	// The duel will begin in {{S1}} second(s).
	TheDuelWillBeginInS1Seconds2 = SysMsg{Id: 1948}

	// Let the duel begin!
	LetTheDuelBegin = SysMsg{Id: 1949}

	// {{C1}} has won the duel.
	C1HasWonTheDuel = SysMsg{Id: 1950}

	// {{C1}}'s party has won the duel.
	C1PartyHasWonTheDuel = SysMsg{Id: 1951}

	// The duel has ended in a tie.
	TheDuelHasEndedInATie = SysMsg{Id: 1952}

	// Since {{C1}} was disqualified, {{S2}} has won.
	SinceC1WasDisqualifiedS2HasWon = SysMsg{Id: 1953}

	// Since {{C1}}'s party was disqualified, {{S2}}'s party has won.
	SinceC1PartyWasDisqualifiedS2PartyHasWon = SysMsg{Id: 1954}

	// Since {{C1}} withdrew from the duel, {{S2}} has won.
	SinceC1WithdrewFromTheDuelS2HasWon = SysMsg{Id: 1955}

	// Since {{C1}}'s party withdrew from the duel, {{S2}}'s party has won.
	SinceC1PartyWithdrewFromTheDuelS2PartyHasWon = SysMsg{Id: 1956}

	// Select the item to be augmented.
	SelectTheItemToBeAugmented = SysMsg{Id: 1957}

	// Select the catalyst for augmentation.
	SelectTheCatalystForAugmentation = SysMsg{Id: 1958}

	// Requires {{S1}} {{S2}}.
	RequiresS1S2 = SysMsg{Id: 1959}

	// This is not a suitable item.
	ThisIsNotASuitableItem = SysMsg{Id: 1960}

	// Gemstone quantity is incorrect.
	GemstoneQuantityIsIncorrect = SysMsg{Id: 1961}

	// The item was successfully augmented!
	TheItemWasSuccessfullyAugmented = SysMsg{Id: 1962}

	// Select the item from which you wish to remove augmentation.
	SelectTheItemFromWhichYouWishToRemoveAugmentation = SysMsg{Id: 1963}

	// Augmentation removal can only be done on an augmented item.
	AugmentationRemovalCanOnlyBeDoneOnAnAugmentedItem = SysMsg{Id: 1964}

	// Augmentation has been successfully removed from your {{S1}}.
	AugmentationHasBeenSuccessfullyRemovedFromYourS1 = SysMsg{Id: 1965}

	// Only the clan leader may issue commands.
	OnlyClanLeaderCanIssueCommands = SysMsg{Id: 1966}

	// The gate is firmly locked. Please try again later.
	GateLockedTryAgainLater = SysMsg{Id: 1967}

	// {{S1}}'s owner.
	S1Owner = SysMsg{Id: 1968}

	// Area where {{S1}} appears.
	AreaS1Appears = SysMsg{Id: 1969}

	// Once an item is augmented, it cannot be augmented again.
	OnceAnItemIsAugmentedItCannotBeAugmentedAgain = SysMsg{Id: 1970}

	// The level of the hardener is too high to be used.
	HardenerLevelTooHigh = SysMsg{Id: 1971}

	// You cannot augment items while a private store or private workshop is in operation.
	YouCannotAugmentItemsWhileAPrivateStoreOrPrivateWorkshopIsInOperation = SysMsg{Id: 1972}

	// You cannot augment items while frozen.
	YouCannotAugmentItemsWhileFrozen = SysMsg{Id: 1973}

	// You cannot augment items while dead.
	YouCannotAugmentItemsWhileDead = SysMsg{Id: 1974}

	// You cannot augment items while engaged in trade activities.
	YouCannotAugmentItemsWhileTrading = SysMsg{Id: 1975}

	// You cannot augment items while paralyzed.
	YouCannotAugmentItemsWhileParalyzed = SysMsg{Id: 1976}

	// You cannot augment items while fishing.
	YouCannotAugmentItemsWhileFishing = SysMsg{Id: 1977}

	// You cannot augment items while sitting down.
	YouCannotAugmentItemsWhileSittingDown = SysMsg{Id: 1978}

	// {{S1}}'s remaining Mana is now 10.
	S1sRemainingManaIsNow10 = SysMsg{Id: 1979}

	// {{S1}}'s remaining Mana is now 5.
	S1sRemainingManaIsNow5 = SysMsg{Id: 1980}

	// {{S1}}'s remaining Mana is now 1. It will disappear soon.
	S1sRemainingManaIsNow1 = SysMsg{Id: 1981}

	// {{S1}}'s remaining Mana is now 0, and the item has disappeared.
	S1sRemainingManaIsNow0 = SysMsg{Id: 1982}

	// Press the Augment button to begin.
	PressTheAugmentButtonToBegin = SysMsg{Id: 1984}

	// {{S1}}'s drop area ({{S2}})
	S1DropAreaS2 = SysMsg{Id: 1985}

	// {{S1}}'s owner ({{S2}})
	S1OwnerS2 = SysMsg{Id: 1986}

	// {{S1}}
	S1 = SysMsg{Id: 1987}

	// The ferry has arrived at Primeval Isle.
	FerryArrivedAtPrimeval = SysMsg{Id: 1988}

	// The ferry will leave for Rune Harbor after anchoring for three minutes.
	FerryLeavingForRune3Minutes = SysMsg{Id: 1989}

	// The ferry is now departing Primeval Isle for Rune Harbor.
	FerryLeavingPrimevalForRuneNow = SysMsg{Id: 1990}

	// The ferry will leave for Primeval Isle after anchoring for three minutes.
	FerryLeavingForPrimeval3Minutes = SysMsg{Id: 1991}

	// The ferry is now departing Rune Harbor for Primeval Isle.
	FerryLeavingRuneForPrimevalNow = SysMsg{Id: 1992}

	// The ferry from Primeval Isle to Rune Harbor has been delayed.
	FerryFromPrimevalToRuneDelayed = SysMsg{Id: 1993}

	// The ferry from Rune Harbor to Primeval Isle has been delayed.
	FerryFromRuneToPrimevalDelayed = SysMsg{Id: 1994}

	// {{S1}} channel filtering option
	S1ChannelFilterOption = SysMsg{Id: 1995}

	// The attack has been blocked.
	AttackWasBlocked = SysMsg{Id: 1996}

	// {{C1}} is performing a counterattack.
	C1PerformingCounterattack = SysMsg{Id: 1997}

	// You countered {{C1}}'s attack.
	CounteredC1Attack = SysMsg{Id: 1998}

	// {{C1}} dodges the attack.
	C1DodgesAttack = SysMsg{Id: 1999}

	// You have avoided {{C1}}'s attack.
	AvoidedC1Attack2 = SysMsg{Id: 2000}

	// Augmentation failed due to inappropriate conditions.
	AugmentationFailedDueToInappropriateConditions = SysMsg{Id: 2001}

	// Trap failed.
	TrapFailed = SysMsg{Id: 2002}

	// You obtained an ordinary material.
	ObtainedOrdinaryMaterial = SysMsg{Id: 2003}

	// You obtained a rare material.
	ObtainedRateMaterial = SysMsg{Id: 2004}

	// You obtained a unique material.
	ObtainedUniqueMaterial = SysMsg{Id: 2005}

	// You obtained the only material of this kind.
	ObtainedOnlyMaterial = SysMsg{Id: 2006}

	// Please enter the recipient's name.
	EnterRecipientsName = SysMsg{Id: 2007}

	// Please enter the text.
	EnterText = SysMsg{Id: 2008}

	// You cannot exceed 1500 characters.
	CantExceed1500Characters = SysMsg{Id: 2009}

	// {{S2}} {{S1}}
	S2S1 = SysMsg{Id: 2010}

	// The augmented item cannot be discarded.
	AugmentedItemCannotBeDiscarded = SysMsg{Id: 2011}

	// {{S1}} has been activated.
	S1HasBeenActivated = SysMsg{Id: 2012}

	// Your seed or remaining purchase amount is inadequate.
	YourSeedOrRemainingPurchaseAmountIsInadequate = SysMsg{Id: 2013}

	// You cannot proceed because the manor cannot accept any more crops. All crops have been returned and no adena withdrawn.
	ManorCantAcceptMoreCrops = SysMsg{Id: 2014}

	// A skill is ready to be used again.
	SkillReadyToUseAgain = SysMsg{Id: 2015}

	// A skill is ready to be used again but its re-use counter time has increased.
	SkillReadyToUseAgainButTimeIncreased = SysMsg{Id: 2016}

	// {{C1}} cannot duel because {{C1}} is currently engaged in a private store or manufacture.
	C1CannotDuelBecauseC1IsCurrentlyEngagedInAPrivateStoreOrManufacture = SysMsg{Id: 2017}

	// {{C1}} cannot duel because {{C1}} is currently fishing.
	C1CannotDuelBecauseC1IsCurrentlyFishing = SysMsg{Id: 2018}

	// {{C1}} cannot duel because {{C1}}'s HP or MP is below 50%.
	C1CannotDuelBecauseC1HpOrMpIsBelow50Percent = SysMsg{Id: 2019}

	// {{C1}} cannot make a challenge to a duel because {{C1}} is currently in a duel-prohibited area (Peaceful Zone / Seven Signs Zone / Near Water / Restart Prohibited Area).
	C1CannotMakeAChallengeToADuelBecauseC1IsCurrentlyInADuelProhibitedArea = SysMsg{Id: 2020}

	// {{C1}} cannot duel because {{C1}} is currently engaged in battle.
	C1CannotDuelBecauseC1IsCurrentlyEngagedInBattle = SysMsg{Id: 2021}

	// {{C1}} cannot duel because {{C1}} is already engaged in a duel.
	C1CannotDuelBecauseC1IsAlreadyEngagedInADuel = SysMsg{Id: 2022}

	// {{C1}} cannot duel because {{C1}} is in a chaotic state.
	C1CannotDuelBecauseC1IsInAChaoticState = SysMsg{Id: 2023}

	// {{C1}} cannot duel because {{C1}} is participating in the Olympiad.
	C1CannotDuelBecauseC1IsParticipatingInTheOlympiad = SysMsg{Id: 2024}

	// {{C1}} cannot duel because {{C1}} is participating in a clan hall war.
	C1CannotDuelBecauseC1IsParticipatingInAClanHallWar = SysMsg{Id: 2025}

	// {{C1}} cannot duel because {{C1}} is participating in a siege war.
	C1CannotDuelBecauseC1IsParticipatingInASiegeWar = SysMsg{Id: 2026}

	// {{C1}} cannot duel because {{C1}} is currently riding a boat, steed, or strider.
	C1CannotDuelBecauseC1IsCurrentlyRidingABoatSteedOrStrider = SysMsg{Id: 2027}

	// {{C1}} cannot receive a duel challenge because {{C1}} is too far away.
	C1CannotReceiveADuelChallengeBecauseC1IsTooFarAway = SysMsg{Id: 2028}

	// {{C1}} is currently teleporting and cannot participate in the Olympiad.
	C1CannotParticipateInOlympiadDuringTeleport = SysMsg{Id: 2029}

	// You are currently logging in.
	CurrentlyLoggingIn = SysMsg{Id: 2030}

	// Please wait a moment.
	PleaseWaitAMoment = SysMsg{Id: 2031}

	// It is not the right time for purchasing the item.
	NotTimeToPurchaseItem = SysMsg{Id: 2032}

	// A sub-class cannot be created or changed because you have exceeded your inventory limit.
	NotSubclassWhileInventoryFull = SysMsg{Id: 2033}

	// There are {{S1}} hour(s) and {{S2}} minute(s) remaining until the time when the item can be purchased.
	ItemPurchasableInS1HoursS2Minutes = SysMsg{Id: 2034}

	// There are {{S1}} minute(s) remaining until the time when the item can be purchased.
	ItemPurchasableInS1Minutes = SysMsg{Id: 2035}

	// Unable to invite because the party is locked.
	NoInvitePartyLocked = SysMsg{Id: 2036}

	// Unable to create character. You are unable to create a new character on the selected server. A restriction is in place which restricts users from creating characters on different servers where no previous characters exists. Please choose another server.
	CantCreateCharacterDuringRestriction = SysMsg{Id: 2037}

	// Some Lineage II features have been limited for free trials. Trial accounts aren't allowed to drop items and/or Adena. To unlock all of the features of Lineage II, purchase the full version today.
	AccountCantDropItems = SysMsg{Id: 2038}

	// Some Lineage II features have been limited for free trials. Trial accounts aren't allowed to trade items and/or Adena. To unlock all of the features of Lineage II, purchase the full version today.
	AccountCantTradeItems = SysMsg{Id: 2039}

	// Cannot trade items with the targeted user.
	CantTradeWithTarget = SysMsg{Id: 2040}

	// Some Lineage II features have been limited for free trials. Trial accounts aren't allowed to setup private stores. To unlock all of the features of Lineage II, purchase the full version today.
	CantOpenPrivateStore = SysMsg{Id: 2041}

	// This account has been suspended for non-payment based on the cell phone payment agreement. Please submit proof of payment by fax (02-2186-3499) and contact customer service at 1600-0020.
	IllegalUse23 = SysMsg{Id: 2042}

	// You have exceeded your inventory volume limit and may not take this quest item. Please make room in your inventory and try again
	YouHaveExceededYourInventoryVolumeLimitAndCannotTakeThisQuestitem = SysMsg{Id: 2043}

	// Some Lineage II features have been limited for free trials. Trial accounts aren't allowed to set up private manufacturing stores. To unlock all of the features of Lineage II, purchase the full version today.
	CantSetupPrivateWorkshop = SysMsg{Id: 2044}

	// Some Lineage II features have been limited for free trials. Trial accounts aren't allowed to use private manufacturing stores. To unlock all of the features of Lineage II, purchase the full version today.
	CantUsePrivateWorkshop = SysMsg{Id: 2045}

	// Some Lineage II features have been limited for free trials. Trial accounts aren't allowed buy items from private stores. To unlock all of the features of Lineage II, purchase the full version today.
	CantUsePrivateStores = SysMsg{Id: 2046}

	// Some Lineage II features have been limited for free trials. Trial accounts aren't allowed to access clan warehouses. To unlock all of the features of Lineage II, purchase the full version today.
	CantUseClanWh = SysMsg{Id: 2047}

	// The shortcut in use conflicts with {{S1}}. Do you wish to reset the conflicting shortcuts and use the saved shortcut?
	ConflictingShortcut = SysMsg{Id: 2048}

	// The shortcut will be applied and saved in the server. Will you continue?
	ConfirmShortcutWillSavedOnServer = SysMsg{Id: 2049}

	// {{S1}} Blood Pledge is trying to display a flag.
	S1TryingRaiseFlag = SysMsg{Id: 2050}

	// You must accept the User Agreement before this account can access Lineage II.
	MustAcceptAgreement = SysMsg{Id: 2051}

	// A guardian's consent is required before this account can be used to play Lineage II.
	NeedConsentToPlayThisAccount = SysMsg{Id: 2052}

	// This account has declined the User Agreement or is pending a withdrawl request.
	AccountDeclinedAgreementOrPending = SysMsg{Id: 2053}

	// This account has been suspended.
	AccountSuspended = SysMsg{Id: 2054}

	// Your account has been suspended from all game services.
	AccountSuspendedFromAllServices = SysMsg{Id: 2055}

	// Your account has been converted to an integrated account, and is unable to be accessed.
	AccountConverted = SysMsg{Id: 2056}

	// You have blocked {{C1}}.
	BlockedC1 = SysMsg{Id: 2057}

	// You are already polymorphed and cannot polymorph again.
	YouAlreadyPolymorphedAndCannotPolymorphAgain = SysMsg{Id: 2058}

	// The nearby area is too narrow for you to polymorph. Please move to another area and try to polymorph again.
	AreaUnsuitableForPolymorph = SysMsg{Id: 2059}

	// You cannot polymorph into the desired form in water.
	YouCannotPolymorphIntoTheDesiredFormInWater = SysMsg{Id: 2060}

	// You are still under transform penalty and cannot be polymorphed.
	CantMorphDueToMorphPenalty = SysMsg{Id: 2061}

	// You cannot polymorph when you have summoned a servitor/pet.
	YouCannotPolymorphWhenYouHaveSummonedAServitor = SysMsg{Id: 2062}

	// You cannot polymorph while riding a pet.
	YouCannotPolymorphWhileRidingAPet = SysMsg{Id: 2063}

	// You cannot polymorph while under the effect of a special skill
	CantMorphWhileUnderSpecialSkillEffect = SysMsg{Id: 2064}

	// That item cannot be taken off
	ItemCannotBeTakenOff = SysMsg{Id: 2065}

	// That weapon cannot perform any attacks.
	ThatWeaponCantAttack = SysMsg{Id: 2066}

	// That weapon cannot use any other skill except the weapon's skill.
	WeaponCanUseOnlyWeaponSkill = SysMsg{Id: 2067}

	// You do not have all of the items needed to untrain the enchant skill.
	YouDontHaveAllItensNeededToUntrainSkillEnchant = SysMsg{Id: 2068}

	// Untrain of enchant skill was successful. Current level of enchant skill {{S1}} has been decreased by 1.
	UntrainSuccessfulSkillS1EnchantLevelDecreasedByOne = SysMsg{Id: 2069}

	// Untrain of enchant skill was successful. Current level of enchant skill {{S1}} became 0 and enchant skill will be initialized.
	UntrainSuccessfulSkillS1EnchantLevelReseted = SysMsg{Id: 2070}

	// You do not have all of the items needed to enchant skill route change.
	YouDontHaveAllItensNeededToChangeSkillEnchantRoute = SysMsg{Id: 2071}

	// Enchant skill route change was successful. Lv of enchant skill {{S1}} has been decreased by {{S2}}.
	SkillEnchantChangeSuccessfulS1LevelWasDecreasedByS2 = SysMsg{Id: 2072}

	// Enchant skill route change was successful. Lv of enchant skill {{S1}} will remain.
	SkillEnchantChangeSuccessfulS1LevelWillRemain = SysMsg{Id: 2073}

	// Skill enchant failed. Current level of enchant skill {{S1}} will remain unchanged.
	SkillEnchantFailedS1LevelWillRemain = SysMsg{Id: 2074}

	// It is not auction period.
	NoAuctionPeriod = SysMsg{Id: 2075}

	// Bidding is not allowed because the maximum bidding price exceeds 100 billion.
	BidCantExceed100Billion = SysMsg{Id: 2076}

	// Your bid must be higher than the current highest bid.
	BidMustBeHigherThanCurrentBid = SysMsg{Id: 2077}

	// You do not have enough adena for this bid.
	NotEnoughAdenaForThisBid = SysMsg{Id: 2078}

	// You currently have the highest bid, but the reserve has not been met.
	HighestBidButReserveNotMet = SysMsg{Id: 2079}

	// You have been outbid.
	YouHaveBeenOutbid = SysMsg{Id: 2080}

	// There are no funds presently due to you.
	NoFundsDue = SysMsg{Id: 2081}

	// You have exceeded the total amount of adena allowed in inventory.
	ExceededMaxAdenaAmountInInventory = SysMsg{Id: 2082}

	// The auction has begun.
	AuctionBegun = SysMsg{Id: 2083}

	// Enemy Blood Pledges have intruded into the fortress.
	EnemiesIntrudedFortress = SysMsg{Id: 2084}

	// Shout and trade chatting cannot be used while possessing a cursed weapon.
	ShoutAndTradeChatCannotBeUsedWhilePossessingCursedWeapon = SysMsg{Id: 2085}

	// Search on user {{S2}} for third-party program use will be completed in {{S1}} minute(s).
	SearchOnS2ForBotUseCompletedInS1Minutes = SysMsg{Id: 2086}

	// A fortress is under attack!
	AFortressIsUnderAttack = SysMsg{Id: 2087}

	// {{S1}} minute(s) until the fortress battle starts.
	S1MinutesUntilTheFortressBattleStarts = SysMsg{Id: 2088}

	// {{S1}} minute(s) until the fortress battle starts.
	S1SecondsUntilTheFortressBattleStarts = SysMsg{Id: 2089}

	// The fortress battle {{S1}} has begun.
	TheFortressBattleS1HasBegun = SysMsg{Id: 2090}

	// Your account can only be used after changing your password and quiz.
	ChangePasswortFirst = SysMsg{Id: 2091}

	// You cannot bid due to a passed-in price.
	CannotBidDueToPassedInPrice = SysMsg{Id: 2092}

	// The passed-in price is {{S1}} adena. Would you like to return the passed-in price?
	PassedInPriceIsS1AdenaWouldYouLikeToReturnIt = SysMsg{Id: 2093}

	// Another user is purchasing. Please try again later.
	AnotherUserPurchasingTryAgainLater = SysMsg{Id: 2094}

	// Some Lineage II features have been limited for free trials. Trial accounts have limited chatting capabilities. To unlock all of the features of Lineage II, purchase the full version today.
	AccountCannotShout = SysMsg{Id: 2095}

	// {{C1}} is in a location which cannot be entered, therefore it cannot be processed.
	C1IsInLocationThatCannotBeEntered = SysMsg{Id: 2096}

	// {{C1}}'s level requirement is not sufficient and cannot be entered.
	C1LevelRequirementNotSufficient = SysMsg{Id: 2097}

	// {{C1}}'s quest requirement is not sufficient and cannot be entered.
	C1QuestRequirementNotSufficient = SysMsg{Id: 2098}

	// {{C1}}'s item requirement is not sufficient and cannot be entered.
	C1ItemRequirementNotSufficient = SysMsg{Id: 2099}

	// {{C1}} may not re-enter yet.
	C1MayNotReenterYet = SysMsg{Id: 2100}

	// You are not currently in a party, so you cannot enter.
	NotInPartyCantEnter = SysMsg{Id: 2101}

	// You cannot enter due to the party having exceeded the limit.
	PartyExceededTheLimitCantEnter = SysMsg{Id: 2102}

	// You cannot enter because you are not associated with the current command channel.
	NotInCommandChannelCantEnter = SysMsg{Id: 2103}

	// The maximum number of instance zones has been exceeded. You cannot enter.
	MaximumInstanceZoneNumberExceededCantEnter = SysMsg{Id: 2104}

	// You have entered another instance zone, therefore you cannot enter corresponding dungeon.
	AlreadyEnteredAnotherInstanceCantEnter = SysMsg{Id: 2105}

	// This dungeon will expire in {{S1}} minute(s). You will be forced out of the dungeon when the time expires.
	DungeonExpiresInS1Minutes = SysMsg{Id: 2106}

	// This instance zone will be terminated in {{S1}} minute(s). You will be forced out of the dungeon when the time expires.
	InstanceZoneTerminatesInS1Minutes = SysMsg{Id: 2107}

	// Your account has been suspended ...
	IllegalUse24 = SysMsg{Id: 2108}

	// The server has been integrated, and your character, {{S1}}, has overlapped with another name. Please enter a new name for your character
	CharacterNameOverlappingRenameCharacter = SysMsg{Id: 2109}

	// This character name already exists or is an invalid name. Please enter a new name
	CharacterNameInvalidRenameCharacter = SysMsg{Id: 2110}

	// Enter a shortcut to assign.
	EnterShortcutToAssign = SysMsg{Id: 2111}

	// Sub-key can be CTRL, ALT, SHIFT and you may enter two sub-keys at a time.
	SubkeyExplanation1 = SysMsg{Id: 2112}

	// (Sub key explanation)
	SubkeyExplanation2 = SysMsg{Id: 2113}

	// Forced attack and stand-in-place attacks assigned previously to Ctrl and Shift will be changed to Alt + Q and Alt + E when set as expanded sub-key mode, and CTRL and SHIFT will be available to assign to another shortcut. Will you continue?
	SubkeyExplanation3 = SysMsg{Id: 2114}

	// Your account has been suspended ...
	IllegalUse25 = SysMsg{Id: 2115}

	// Your account has been suspended ...
	IllegalUse26 = SysMsg{Id: 2116}

	// Your account has been suspended ...
	IllegalUse27 = SysMsg{Id: 2117}

	// Your account has been suspended ...
	IllegalUse28 = SysMsg{Id: 2118}

	// Your account has been suspended ...
	IllegalUse29 = SysMsg{Id: 2119}

	// Your account has been suspended ...
	IllegalUse30 = SysMsg{Id: 2120}

	// Your account has been suspended ...
	IllegalUse31 = SysMsg{Id: 2121}

	// Your account has been suspended ...
	IllegalUse32 = SysMsg{Id: 2122}

	// Your account has been suspended ...
	IllegalUse33 = SysMsg{Id: 2123}

	// The server has been integrated, and your Clan name, {{S1}}, has been overlapped with another name. Please enter the Clan name to be changed.
	ClanNameOverlappingRenameClan = SysMsg{Id: 2124}

	// This name already exists or is an invalid name. Please enter the Clan name to be changed.
	ClanNameInvalidRenameClan = SysMsg{Id: 2125}

	// Your account has been suspended ...
	IllegalUse34 = SysMsg{Id: 2126}

	// Your account has been suspended ...
	IllegalUse35 = SysMsg{Id: 2127}

	// Your account has been suspended ...
	IllegalUse36 = SysMsg{Id: 2128}

	// The augmented item cannot be converted. Please convert after the augmentation has been removed.
	AugmentedItemCantConverted = SysMsg{Id: 2129}

	// You cannot convert this item.
	CantConvertThisItem = SysMsg{Id: 2130}

	// You have bid the highest price and have won the item. The item can be found in your personal warehouse.
	WonBidItemCanBeFoundInWarehouse = SysMsg{Id: 2131}

	// You have entered a common server.
	EnteredCommonServer = SysMsg{Id: 2132}

	// You have entered an adults-only server.
	EnteredAdultsOnlyServer = SysMsg{Id: 2133}

	// You have entered a server for juveniles.
	EnteredJuvenilesServer = SysMsg{Id: 2134}

	// Because of your Fatigue level, this is not allowed.
	NotAllowedDueToFatigueLevel = SysMsg{Id: 2135}

	// A clan name change application has been submitted.
	ClanNameChancePetitionSubmitted = SysMsg{Id: 2136}

	// You are about to bid {{S1}} item with {{S2}} adena. Will you continue?
	ConfirmBidS2AdenaForS1Item = SysMsg{Id: 2137}

	// Please enter a bid price.
	EnterBidPrice = SysMsg{Id: 2138}

	// {{C1}}'s Pet.
	C1Pet = SysMsg{Id: 2139}

	// {{C1}}'s Servitor.
	C1Servitor = SysMsg{Id: 2140}

	// You slightly resisted {{C1}}'s magic.
	SlightlyResistedC1Magicc = SysMsg{Id: 2141}

	// You cannot expel {{C1}} because {{C1}} is not a party member.
	CantExpelC1NotAPartyMember = SysMsg{Id: 2142}

	// You cannot add elemental power while operating a Private Store or Private Workshop.
	CannotAddElementalPowerWhileOperatingPrivateStoreOrWorkshop = SysMsg{Id: 2143}

	// Please select item to add elemental power.
	SelectItemToAddElementalPower = SysMsg{Id: 2144}

	// Attribute item usage has been cancelled.
	ElementalEnhanceCanceled = SysMsg{Id: 2145}

	// Elemental power enhancer usage requirement is not sufficient.
	ElementalEnhanceRequirementNotSufficient = SysMsg{Id: 2146}

	// {{S2}} elemental power has been added successfully to {{S1}}.
	ElementalPowerS2SuccessfullyAddedToS1 = SysMsg{Id: 2147}

	// {{S3}} elemental power has been added successfully to +{{S1}} {{S2}}.
	ElementalPowerS3SuccessfullyAddedToS1S2 = SysMsg{Id: 2148}

	// You have failed to add elemental power.
	FailedAddingElementalPower = SysMsg{Id: 2149}

	// Another elemental power has already been added. This elemental power cannot be added.
	AnotherElementalPowerAlreadyAdded = SysMsg{Id: 2150}

	// Your opponent has resistance to magic, the damage was decreased.
	OpponentHasResistanceMagicDamageDecreased = SysMsg{Id: 2151}

	// The assigned shortcut will be deleted and the initial shortcut setting restored. Will you continue?
	ConfirmShorcutDelete = SysMsg{Id: 2152}

	// You are currently logged into 10 of your accounts and can no longer access your other accounts.
	MaximumAccountLoginsReached = SysMsg{Id: 2153}

	// The target is not a flagpole so a flag cannot be displayed.
	TheTargetIsNotAFlagpoleSoAFlagCannotBeDisplayed = SysMsg{Id: 2154}

	// A flag is already being displayed, another flag cannot be displayed.
	AFlagIsAlreadyBeingDisplayedAnotherFlagCannotBeDisplayed = SysMsg{Id: 2155}

	// There are not enough necessary items to use the skill.
	ThereAreNotEnoughNecessaryItemsToUseTheSkill = SysMsg{Id: 2156}

	// Bid will be attempted with {{S1}} adena.
	BidWillBeAttemptedWithS1Adena = SysMsg{Id: 2157}

	// Force attack is impossible against a temporary allied member during a siege.
	ForcedAttackIsImpossibleAgainstSiegeSideTemporaryAlliedMembers = SysMsg{Id: 2158}

	// Bidder exists, the auction time has been extended by 5 minutes.
	BidderExistsAuctionTimeExtendedBy5Minutes = SysMsg{Id: 2159}

	// Bidder exists, the auction time has been extended by 3 minutes.
	BidderExistsAuctionTimeExtendedBy3Minutes = SysMsg{Id: 2160}

	// There is not enough space to move, the skill cannot be used.
	NotEnoughSpaceForSkill = SysMsg{Id: 2161}

	// Your soul has increased by {{S1}}, so it is now at {{S2}}.
	YourSoulHasIncreasedByS1SoItIsNowAtS2 = SysMsg{Id: 2162}

	// Soul cannot be increased anymore.
	SoulCannotBeIncreasedAnymore = SysMsg{Id: 2163}

	// The barracks have been seized.
	SeizedBarracks = SysMsg{Id: 2164}

	// The barracks function has been restored.
	BarracksFunctionRestored = SysMsg{Id: 2165}

	// All barracks are occupied.
	AllBarracksOccupied = SysMsg{Id: 2166}

	// A malicious skill cannot be used in a peace zone.
	AMaliciousSkillCannotBeUsedInPeaceZone = SysMsg{Id: 2167}

	// {{C1}} has acquired the flag.
	C1AcquiredTheFlag = SysMsg{Id: 2168}

	// Your clan has been registered to {{S1}}'s fortress battle.
	RegisteredToS1FortressBattle = SysMsg{Id: 2169}

	// A malicious skill cannot be used when an opponent is in the peace zone
	CantUseBadMagicWhenOpponentInPeaceZone = SysMsg{Id: 2170}

	// This item cannot be crystallized.
	ItemCannotCrystallized = SysMsg{Id: 2171}

	// +{{S1}} {{S2}}'s auction has ended.
	S1S2AuctionEnded = SysMsg{Id: 2172}

	// {{S1}}'s auction has ended.
	S1AuctionEnded = SysMsg{Id: 2173}

	// {{C1}} cannot duel because {{C1}} is currently polymorphed.
	C1CannotDuelWhilePolymorphed = SysMsg{Id: 2174}

	// Party duel cannot be initiated due to a polymorphed partymember
	CannotPartyDuelWhileAMemberIsPolymorphed = SysMsg{Id: 2175}

	// {{S1}}'s elemental power has been removed.
	S1ElementalPowerRemoved = SysMsg{Id: 2176}

	// +{{S1}} {{S2}}'s elemental power has been removed.
	S1S2ElementalPowerRemoved = SysMsg{Id: 2177}

	// You failed to remove the elemental power.
	FailedToRemoveElementalPower = SysMsg{Id: 2178}

	// You have the highest bid submitted in Giran Castle Auction.
	HighestBidForGiranCastle = SysMsg{Id: 2179}

	// You have the highest bid submitted in Aden Castle Auction.
	HighestBidForAdenCastle = SysMsg{Id: 2180}

	// You have the highest bid submitted in Rune Castle Auction.
	HighestBidForRuneCastle = SysMsg{Id: 2181}

	// You cannot polymorph while riding a boat.
	CantPolymorphOnBoat = SysMsg{Id: 2182}

	// The fortress battle of {{S1}} has finished.
	TheFortressBattleOfS1HasFinished = SysMsg{Id: 2183}

	// {{S1}} clan is victorious in the fortress battle of {{S2}}.
	S1ClanIsVictoriousInTheFortressBattleOfS2 = SysMsg{Id: 2184}

	// Only a party leader can try to enter.
	OnlyPartyLeaderCanEnter = SysMsg{Id: 2185}

	// Soul cannot be absorbed anymore.
	SoulCannotBeAbsorbedAnymore = SysMsg{Id: 2186}

	// The target is located where you cannot charge.
	CantReachTargetToCharge = SysMsg{Id: 2187}

	// Another enchantment is in progress. Please complete previous task and try again.
	EnchantmentAlreadyInProgress = SysMsg{Id: 2188}

	// Current Location: {{S1}}, {{S2}}, {{S3}} (near Near Kamael Village)
	LocKamaelVillageS1S2S3 = SysMsg{Id: 2189}

	// Current Location: {{S1}}, {{S2}}, {{S3}} (Near south of Wastelands Camp)
	LocWastelandsCampS1S2S3 = SysMsg{Id: 2190}

	// To apply selected options, the game needs to be reloaded. If you don't apply now, it will be applied when you start the game next time. Will you apply now?
	ConfirmApplySelections = SysMsg{Id: 2191}

	// You have bid on an item auction.
	BidOnItemAuction = SysMsg{Id: 2192}

	// It's too far from the NPC to work.
	TooFarFromNpc = SysMsg{Id: 2193}

	// Current polymorph form cannot be applied with corresponding effects.
	CantApplyCurrentPolymorphWithCorrespondingEffects = SysMsg{Id: 2194}

	// There is not enough soul.
	ThereIsNotEnoughSoul = SysMsg{Id: 2195}

	// No Owned Clan.
	NoOwnedClan = SysMsg{Id: 2196}

	// Owned by clan {{S1}}.
	OwnedS1Clan = SysMsg{Id: 2197}

	// You have the highest bid in an item auction.
	HighestBidInItemAuction = SysMsg{Id: 2198}

	// You cannot enter this instance zone while the NPC server is unavailable.
	CantEnterInstanceZoneNpcServerOffline = SysMsg{Id: 2199}

	// This instance zone will be terminated because the NPC server is unavailable. You will be forcibly removed from the dungeon shortly
	InstanceZoneTerminatedNpcServerOffline = SysMsg{Id: 2200}

	// {{S1}} year(s) {{S2}} month(s) {{S3}} day(s)
	S1YearsS2MonthsS3Days = SysMsg{Id: 2201}

	// {{S1}} hour(s) {{S2}} minute(s) {{S3}} second(s)
	S1HoursS2MinutesS3Seconds = SysMsg{Id: 2202}

	// {{S1}} month(s) {{S2}} day(s)
	S1MonthsS2Days = SysMsg{Id: 2203}

	// {{S1}} hour(s)
	S1Hours = SysMsg{Id: 2204}

	// You have entered an area where the mini map cannot be used. The mini map will be closed.
	AreaForbidsMinimap = SysMsg{Id: 2205}

	// You have entered an area where the mini map can be used.
	AreaAllowsMinimap = SysMsg{Id: 2206}

	// This is an area where you cannot use the mini map. The mini map will not be opened.
	CantOpenMinimap = SysMsg{Id: 2207}

	// You do not meet the skill level requirements.
	YouDontMeetSkillLevelRequirements = SysMsg{Id: 2208}

	// This is an area where radar cannot be used
	AreaWhereRadarCannotBeUsed = SysMsg{Id: 2209}

	// It will return to an unenchanted condition.
	ReturnToUnenchantedCondition = SysMsg{Id: 2210}

	// You must learn the Onyx Beast skill before you can acquire further skills.
	YouMustLearnOnyxBeastSkill = SysMsg{Id: 2211}

	// You have not completed the necessary quest for skill acquisition.
	NotCompletedQuestForSkillAcquisition = SysMsg{Id: 2212}

	// Cannot board a ship while polymorphed.
	CantBoardShipPolymorphed = SysMsg{Id: 2213}

	// A new character will be created with the current settings. Continue
	ConfirmCharacterCreation = SysMsg{Id: 2214}

	// {{S1}} P.Def
	S1Pdef = SysMsg{Id: 2215}

	// The CPU driver is not up to date. Please install an up-to-date CPU driver.
	PleaseUpdateCpuDriver = SysMsg{Id: 2216}

	// The ballista has been successfully destroyed and the clan's reputation will be increased.
	BallistaDestroyedClanRepuIncreased = SysMsg{Id: 2217}

	// This is a main class skill only.
	MainClassSkillOnly = SysMsg{Id: 2218}

	// This squad skill has already been acquired.
	SquadSkillAlreadyAcquired = SysMsg{Id: 2219}

	// The previous level skill has not been learned.
	PreviousLevelSkillNotLearned = SysMsg{Id: 2220}

	// Will you activate the selected functions?
	ActivateSelectedFuntionsConfirm = SysMsg{Id: 2221}

	// It will cost 150,000 adena to place scouts. Will you place them.
	ScoutCosts150000Adena = SysMsg{Id: 2222}

	// It will cost 200,000 adena for a fortress gate enhancement. Will you enhance it?
	FortressGateCosts200000Adena = SysMsg{Id: 2223}

	// Crossbow is preparing to fire.
	CrossbowPreparingToFire = SysMsg{Id: 2224}

	// There are no other skills to learn. Please come back after {{S1}}nd class change.
	NoSkillsToLearnReturnAfterS1ClassChange = SysMsg{Id: 2225}

	// Not enough bolts.
	NotEnoughBolts = SysMsg{Id: 2226}

	// It is not possible to register for the castle siege side or castle siege of a higher castle in the contract
	NotPossibleToRegisterToCastleSiege = SysMsg{Id: 2227}

	// Instance zone time limit:
	InstanceZoneTimeLimit = SysMsg{Id: 2228}

	// There is no instance zone under a time limit
	NoInstancezoneTimeLimit = SysMsg{Id: 2229}

	// Available to use after {{S1}} {{S2}}hour(s) {{S3}}minute(s).
	AvailableAfterS1S2HoursS3Minutes = SysMsg{Id: 2230}

	// The reputation score of the upper castle in contract is not enough and supply was not granted.
	ReputationScoreForContractNotEnough = SysMsg{Id: 2231}

	// {{S1}} will be crystallized before destruction. Will you continue?
	S1CrystallizedBeforeDestruction = SysMsg{Id: 2232}

	// Siege registration is not possible due to a contract with a higher castle.
	CantRegisterToSiegeDueToContract = SysMsg{Id: 2233}

	// Will you use the selected Kamael-race-only Hero Weapon?
	ConfirmKamaelHeroWeapon = SysMsg{Id: 2234}

	// The instance zone in use has been deleted and cannot be accessed.
	InstanceZoneDeletedCantAccessed = SysMsg{Id: 2235}

	// {{S1}} minute(s) left for wyvern riding.
	S1MinutesLeftOnWyvern = SysMsg{Id: 2236}

	// {{S1}} seconds(s) left for wyvern riding.
	S1SecondsLeftOnWyvern = SysMsg{Id: 2237}

	// You have participated in the siege of {{S1}}. This siege will continue for 2 hours.
	ParticipatingInSiegeOfS1 = SysMsg{Id: 2238}

	// The siege of {{S1}}, in which you are participating, has finished.
	SiegeOfS1Finihsed = SysMsg{Id: 2239}

	// You cannot register for the Team Battle Clan Hall War when your Clan Lord is on the waiting list for a transaction.
	CantRegisterToTeamBattleClanHallWarWhileLordOnTransactionWaitingList = SysMsg{Id: 2240}

	// You cannot apply for a Clan Lord transaction if your clan has registed for the Team Battle Clan Hall War.
	CantApplyOnLordTransactionWhileRegisteredToTeamBattleClanHallWar = SysMsg{Id: 2241}

	// Clan members cannot leave or be expelled when they are regisered for the Team Battle Clan Hall War.
	MembersCantLeaveWhenRegisteredToTeamBattleClanHallWar = SysMsg{Id: 2242}

	// During the Bandit Stronghold or Wild Beast Reserve clan hall war, the previous clan lord rather than the new clan lord participates in battle.
	WhenBanditstrongholdWildbeastreservreClanlordInDangerPreviousLordParticipatesInBattle = SysMsg{Id: 2243}

	// {{S1}} minute(s) remaining.
	S1MinutesRemaining = SysMsg{Id: 2244}

	// {{S1}} second(s) remaining.
	S1SecondsRemaining = SysMsg{Id: 2245}

	// The contest will begin in {{S1}} minute(s).
	ContestBeginInS1Minutes = SysMsg{Id: 2246}

	// You cannot board an airship while transformed.
	YouCannotBoardAnAirshipWhileTransformed = SysMsg{Id: 2247}

	// You cannot board an airship while petrified.
	YouCannotBoardAnAirshipWhilePetrified = SysMsg{Id: 2248}

	// You cannot board an airship while dead.
	YouCannotBoardAnAirshipWhileDead = SysMsg{Id: 2249}

	// You cannot board an airship while fishing.
	YouCannotBoardAnAirshipWhileFishing = SysMsg{Id: 2250}

	// You cannot board an airship while in battle.
	YouCannotBoardAnAirshipWhileInBattle = SysMsg{Id: 2251}

	// You cannot board an airship while in a duel.
	YouCannotBoardAnAirshipWhileInADuel = SysMsg{Id: 2252}

	// You cannot board an airship while sitting.
	YouCannotBoardAnAirshipWhileSitting = SysMsg{Id: 2253}

	// You cannot board an airship while casting.
	YouCannotBoardAnAirshipWhileCasting = SysMsg{Id: 2254}

	// You cannot board an airship when a cursed weapon is equipped.
	YouCannotBoardAnAirshipWhileACursedWeaponIsEquipped = SysMsg{Id: 2255}

	// You cannot board an airship while holding a flag.
	YouCannotBoardAnAirshipWhileHoldingAFlag = SysMsg{Id: 2256}

	// You cannot board an airship while a pet or a servitor is summoned.
	YouCannotBoardAnAirshipWhileAPetOrAServitorIsSummoned = SysMsg{Id: 2257}

	// You have already boarded another airship.
	YouHaveAlreadyBoardedAnotherAirship = SysMsg{Id: 2258}

	// Current Location: {{S1}}, {{S2}}, {{S3}} (near Fantasy Isle)
	LocFantasyIslandS1S2S3 = SysMsg{Id: 2259}

	// A pet can run away if you do not fill its hunger gauge to 10% or above.
	PetCanRunAwayWhenHungerBelow10Percent = SysMsg{Id: 2260}

	// {{C1}} has given $c2 damage of {{S3}}.
	C1GaveC2DamageOfS3 = SysMsg{Id: 2261}

	// {{C1}} has received {{S3}} damage from $c2.
	C1ReceivedDamageOfS3FromC2 = SysMsg{Id: 2262}

	// {{C1}} has received damage of {{S3}} through $c2.
	C1ReceivedDamageOfS3ThroughC2 = SysMsg{Id: 2263}

	// {{C1}} has evaded $c2's attack.
	C1EvadedC2Attack = SysMsg{Id: 2264}

	// {{C1}}'s attack went astray.
	C1AttackWentAstray = SysMsg{Id: 2265}

	// {{C1}} had a critical hit!
	C1HadCriticalHit = SysMsg{Id: 2266}

	// {{C1}} resisted $c2's drain.
	C1ResistedC2Drain = SysMsg{Id: 2267}

	// {{C1}}'s attack failed.
	C1AttackFailed = SysMsg{Id: 2268}

	// {{C1}} resisted $c2's magic.
	C1ResistedC2Drain2 = SysMsg{Id: 2269}

	// {{C1}} has received damage from {{S2}} through the fire of magic
	C1ReceivedDamageFromS2ThroughFireOfMagic = SysMsg{Id: 2270}

	// {{C1}} weakly resisted $c2's magic.
	C1WeaklyResistedC2Magic = SysMsg{Id: 2271}

	// You have selected shortcuts without settings up sub-keys. You can only use the set shortcut in the Enter Chat mode. Do you still wish to use the set shortcuts
	UseShortcutConfirm = SysMsg{Id: 2272}

	// This skill cannot be learned while in the sub-class state. Please try again after changing to the main class.
	SkillNotForSubclass = SysMsg{Id: 2273}

	// The rebel army recaptured the fortress.
	NpcsRecapturedFortress = SysMsg{Id: 2276}

	// You cannot transform while sitting.
	CannotTransformWhileSitting = SysMsg{Id: 2283}

	// You can operate the machine when you participate in the party.
	CanOperateMachineWhenInParty = SysMsg{Id: 2291}

	// Current location: {{S1}}, {{S2}}, {{S3}} (inside the Steel Citadel)
	LocInSteelCitadelS1S2S3 = SysMsg{Id: 2293}

	// You have gained Vitality points.
	GainedVitalityPoints = SysMsg{Id: 2296}

	// Current location: Steel Citadel
	LocSteelCitadel = SysMsg{Id: 2301}

	// Your Vitamin Item has arrived! Visit the Vitamin Manager in any village to obtain it
	YourVitaminItemHasArrived = SysMsg{Id: 2302}

	// There are {{S2}} second(s) remaining in {{S1}}'s re-use time.
	S2SecondsRemainingForReuseS1 = SysMsg{Id: 2303}

	// There are {{S2}} minute(s), {{S3}} second(s) remaining in {{S1}}'s re-use time.
	S2MinutesS3SecondsRemainingForReuseS1 = SysMsg{Id: 2304}

	// There are {{S2}} hour(s), {{S3}} minute(s), and {{S4}} second(s) remaining in {{S1}}'s re-use time.
	S2HoursS3MinutesS4SecondsRemainingForReuseS1 = SysMsg{Id: 2305}

	// Resurrection is possible because of the courage charm's effect. Would you like to resurrect now?
	ResurrectUsingCharmOfCourage = SysMsg{Id: 2306}

	// You do not have a servitor.
	DontHaveServitor = SysMsg{Id: 2311}

	// You do not have a pet.
	DontHavePet = SysMsg{Id: 2312}

	// Your Vitality is at maximum.
	VitalityIsAtMaximum = SysMsg{Id: 2314}

	// You have gained Vitality points.
	VitalityHasIncreased = SysMsg{Id: 2315}

	// You have lost Vitality points.
	VitalityHasDecreased = SysMsg{Id: 2316}

	// Your Vitality is fully exhausted.
	VitalityIsExhausted = SysMsg{Id: 2317}

	// You have acquired {{S1}} reputation score.
	AcquiredS1ReputationScore = SysMsg{Id: 2319}

	// Current location: Inside Kamaloka
	LocKamaloka = SysMsg{Id: 2321}

	// Current location: Inside Nia Kamaloka
	LocNiaKamaloka = SysMsg{Id: 2322}

	// Current location: Inside Rim Kamaloka
	LocRimKamaloka = SysMsg{Id: 2323}

	// You have acquired 50 Clan's Fame Points..
	Acquired50ClanFamePoints = SysMsg{Id: 2326}

	// You don't have enough reputation score.
	NotEnoughFamePoints = SysMsg{Id: 2327}

	// You cannot receive the vitamin item because you have exceed your inventory weight/quantity limit.
	YouCannotReceiveTheVitaminItem = SysMsg{Id: 2333}

	// There are no more vitamin items to be found
	ThereAreNoMoreVitaminItemsToBeFound = SysMsg{Id: 2335}

	// Half-Kill!
	HalfKill = SysMsg{Id: 2336}

	// Your CP was drained because you were hit with a CP siphon skill.
	CpDisappearsWhenHitWithAHalfKillSkill = SysMsg{Id: 2337}

	// You cannot use My Teleports during a battle.
	YouCannotUseMyTeleportsDuringABattle = SysMsg{Id: 2348}

	// You cannot use My Teleports while participating a large-scale battle such as a castle siege, fortress siege, or hideout siege..
	YouCannotUseMyTeleportsWhileParticipating = SysMsg{Id: 2349}

	// You cannot use My Teleports during a duel
	YouCannotUseMyTeleportsDuringADuel = SysMsg{Id: 2350}

	// You cannot use My Teleports while flying
	YouCannotUseMyTeleportsWhileFlying = SysMsg{Id: 2351}

	// You cannot use My Teleports while participating in an Olympiad match
	YouCannotUseMyTeleportsWhileParticipatingInAnOlympiadMatch = SysMsg{Id: 2352}

	// You cannot use My Teleports while you are in a flint or paralyzed state
	YouCannotUseMyTeleportsWhileYouAreParalyzed = SysMsg{Id: 2353}

	// You cannot use My Teleports while you are dead
	YouCannotUseMyTeleportsWhileYouAreDead = SysMsg{Id: 2354}

	// You cannot use My Teleports in this area
	YouCannotUseMyTeleportsInThisArea = SysMsg{Id: 2355}

	// You cannot use My Teleports underwater
	YouCannotUseMyTeleportsUnderwater = SysMsg{Id: 2356}

	// You cannot use My Teleports in an instant zone
	YouCannotUseMyTeleportsInAnInstantZone = SysMsg{Id: 2357}

	// You have no space to save the teleport location
	YouHaveNoSpaceToSaveTheTeleportLocation = SysMsg{Id: 2358}

	// You cannot teleport because you do not have a teleport item
	YouCannotTeleportBecauseYouDoNotHaveATeleportItem = SysMsg{Id: 2359}

	// Current Location: {{S1}}
	CurrentLocationS1 = SysMsg{Id: 2361}

	// The limited-time item has been deleted..
	TimeLimitedItemDeleted = SysMsg{Id: 2366}

	// There is not much time remaining until the hunting helper pet leaves.
	ThereNotMuchTimeRemainingUntilHelperLeaves = SysMsg{Id: 2372}

	// The hunting helper pet is now leaving.
	TheHelperPetLeaving = SysMsg{Id: 2373}

	// The hunting helper pet cannot be returned ecause there is not much time remaining until it leaves.
	TheHelperPetCannotBeReturned = SysMsg{Id: 2375}

	// You cannot receive a vitamin item during an exchange.
	YouCannotReceiveAVitaminItemDuringAnExchange = SysMsg{Id: 2376}

	// Your number of My Teleports slots has reached its maximum limit.
	YourNumberOfMyTeleportsSlotsHasReachedItsMaximumLimit = SysMsg{Id: 2390}

	// That pet/servitor skill cannot be used because it is recharging.
	PetSkillCannotBeUsedRecharging = SysMsg{Id: 2396}

	// You have no open My Teleports slots.
	YouHaveNoOpenMyTeleportsSlots = SysMsg{Id: 2398}

	// {{C1}} is already registered on the waiting list for the non-class-limited match event.
	C1IsAlreadyRegisteredNonClassLimitedEventTeams = SysMsg{Id: 2440}

	// Only a party leader can request a team match.
	OnlyPartyLeaderCanRequestTeamMatch = SysMsg{Id: 2441}

	// The request cannot be made because the requirements have not been made. To participate in a team match you must first form a 3-member party.
	PartyRequirementsNotMet = SysMsg{Id: 2442}

	// The disguise scroll cannot be used because it is meant for use in a different territory.
	TheDisguiseScrollMeantForDifferentTerritory = SysMsg{Id: 2936}

	// A territory owning clan member cannot use a disguise scroll.
	TerritoryOwningClanCannotUseDisguiseScroll = SysMsg{Id: 2937}

	// The territory war exclusive disguise and transformation can be used 20 minutes before the start of the territory war to 10 minutes after its end.
	TerritoryWarScrollCanNotUsedNow = SysMsg{Id: 2955}

	// Instant Zone currently in use: {{S1}}
	InstantZoneCurrentlyInuseS1 = SysMsg{Id: 2400}

	// The Territory War request period has ended.
	TheTerritoryWarRegisteringPeriodEnded = SysMsg{Id: 2402}

	// Territory War begins in 10 minutes!
	TerritoryWarBeginsIn10Minutes = SysMsg{Id: 2403}

	// Territory War begins in 5 minutes!
	TerritoryWarBeginsIn5Minutes = SysMsg{Id: 2404}

	// Territory War begins in 1 minute!
	TerritoryWarBeginsIn1Minute = SysMsg{Id: 2405}

	// You have registered on the waiting list for the non-class-limited team match event.
	YouHaveRegisteredInAWaitingListOfTeamGames = SysMsg{Id: 2408}

	// The number of My Teleports slots has been increased.
	TheNumberOfMyTeleportsSlotsHasBeenIncreased = SysMsg{Id: 2409}

	// You cannot use My Teleports to reach this area!
	YouCannotUseMyTeleportsToReachThisArea = SysMsg{Id: 2410}

	// The collection has failed.
	TheCollectionHasFailed = SysMsg{Id: 2424}

	// Your birthday gift has arrived
	YourBirthdayGiftHasArrived = SysMsg{Id: 2448}

	// There are {{S1}} days until your character's birthday.
	ThereAreS1DaysUntilYourCharactersBirthday = SysMsg{Id: 2449}

	// {{C1}}'s character birthday is {{S3}}/{{S4}}/{{S2}}.
	C1BirthdayIsS3S4S2 = SysMsg{Id: 2450}

	// The cloak equip has been removed because the armor set equip has been removed.
	CloakRemovedBecauseArmorSetRemoved = SysMsg{Id: 2451}

	// The airship must be summoned in order for you to board.
	TheAirshipMustBeSummonedToBoard = SysMsg{Id: 2455}

	// In order to acquire an airship, the clan's level must be level 5 or higher.
	TheAirshipNeedClanlvl5ToSummon = SysMsg{Id: 2456}

	// An airship cannot be summoned because either you have not registered your airship license, or the airship has not yet been summoned
	TheAirshipNeedLicenseToSummon = SysMsg{Id: 2457}

	// The airship owned by the clan is already being used by another clan member.
	TheAirshipAlreadyUsed = SysMsg{Id: 2458}

	// The Airship Summon License has already been acquired.
	TheAirshipSummonLicenseAlreadyAcquired = SysMsg{Id: 2459}

	// The clan owned airship already exists.
	TheAirshipIsAlreadyExists = SysMsg{Id: 2460}

	// The airship owned by the clan can only be purchased by the clan lord.
	TheAirshipNoPrivileges = SysMsg{Id: 2461}

	// The airship cannot be summoned because you don't have enough {{S1}}%.
	TheAirshipNeedMoreS1 = SysMsg{Id: 2462}

	// The airship's fuel (EP) will soon run out.
	TheAirshipFuelSoonRunOut = SysMsg{Id: 2463}

	// The airship's fuel (EP) has run out. The airship's speed will be greatly decreased in this condition.
	TheAirshipFuelRunOut = SysMsg{Id: 2464}

	// You have selected a 3 vs 3 class irrelevant team match. Do you wish to participate?
	Olympiad3Vs3Confirm = SysMsg{Id: 2465}

	// A pet on auxiliary mode cannot use skills.
	PetAuxiliaryModeCannotUseSkills = SysMsg{Id: 2466}

	// Your ship cannot teleport because it does not have enough fuel for the trip.
	TheAirshipCannotTeleport = SysMsg{Id: 2491}

	// The airship has been summoned. It will automatically depart in %s minutes.
	TheAirshipSummoned = SysMsg{Id: 2492}

	// The collection has succeeded.
	TheCollectionHasSucceeded = SysMsg{Id: 2500}

	// The match is being prepared. Please try again later.
	MatchBeingPreparedTryLater = SysMsg{Id: 2701}

	// You were excluded from the match because the registration count was not correct.
	ExcludedFromMatchDueIncorrectCount = SysMsg{Id: 2702}

	// The team was adjusted because the population ratio was not correct.
	TeamAdjustedBecauseWrongPopulationRatio = SysMsg{Id: 2703}

	// You cannot register because capacity has been exceeded.
	CannotRegisterCauseQueueFull = SysMsg{Id: 2704}

	// The match waiting time was extended by 1 minute.
	MatchWaitingTimeExtended = SysMsg{Id: 2705}

	// You cannot enter because you do not meet the requirements.
	CannotEnterCauseDontMatchRequirements = SysMsg{Id: 2706}

	// You cannot make another request for 10 seconds after cancelling a match registration.
	CannotRequestRegistration10SecsAfter = SysMsg{Id: 2707}

	// You cannot register while possessing a cursed weapon.
	CannotRegisterProcessingCursedWeapon = SysMsg{Id: 2708}

	// Applicants for the Olympiad, Underground Coliseum, or Kratei's Cube matches cannot register.
	ColiseumOlympiadKrateisApplicantsCannotParticipate = SysMsg{Id: 2709}

	// Current location: {{S1}}, {{S2}}, {{S3}} (near the Keucereus clan association location)
	LocKeucereusS1S2S3 = SysMsg{Id: 2710}

	// Current location: {{S1}}, {{S2}}, {{S3}} (inside the Seed of Infinity)
	LocInSeedInfinityS1S2S3 = SysMsg{Id: 2711}

	// Current location: {{S1}}, {{S2}}, {{S3}} (outside the Seed of Infinity)
	LocOutSeedInfinityS1S2S3 = SysMsg{Id: 2712}

	// Current location: {{S1}}, {{S2}}, {{S3}} (inside Aerial Cleft)
	LocCleftS1S2S3 = SysMsg{Id: 2716}

	// Instant zone from here: {{S1}}'s entry has been restricted.
	InstantZoneS1Restricted = SysMsg{Id: 2720}

	// Boarding or cancellation of boarding on Airships is not allowed in the current area.
	BoardOrCancelNotPossibleHere = SysMsg{Id: 2721}

	// Another airship has already been summoned at the wharf. Please try again later.
	AnotherAirshipAlreadySummoned = SysMsg{Id: 2722}

	// You cannot board because you do not meet the requirements.
	YouCannotBoardNotMeetRequeirements = SysMsg{Id: 2727}

	// This action is prohibited while mounted or on an airship.
	ActionProhibitedWhileMountedOrOnAnAirship = SysMsg{Id: 2728}

	// You cannot control the helm while transformed.
	YouCannotControlTheHelmWhileTransformed = SysMsg{Id: 2729}

	// You cannot control the helm while you are petrified.
	YouCannotControlTheHelmWhileYouArePetrified = SysMsg{Id: 2730}

	// You cannot control the helm when you are dead.
	YouCannotControlTheHelmWhenYouAreDead = SysMsg{Id: 2731}

	// You cannot control the helm while fishing.
	YouCannotControlTheHelmWhileFishing = SysMsg{Id: 2732}

	// You cannot control the helm while in a battle.
	YouCannotControlTheHelmWhileInABattle = SysMsg{Id: 2733}

	// You cannot control the helm while in a duel.
	YouCannotControlTheHelmWhileInADuel = SysMsg{Id: 2734}

	// You cannot control the helm while in a sitting position.
	YouCannotControlTheHelmWhileInASittingPosition = SysMsg{Id: 2735}

	// You cannot control the helm while using a skill.
	YouCannotControlTheHelmWhileUsingASkill = SysMsg{Id: 2736}

	// You cannot control the helm while a cursed weapon is equipped.
	YouCannotControlTheHelmWhileACursedWeaponIsEquipped = SysMsg{Id: 2737}

	// You cannot control the helm while holding a flag.
	YouCannotControlTheHelmWhileHoldingAFlag = SysMsg{Id: 2738}

	// The {{S1}} ward has been destroyed! $c2 now has the territory ward.
	TheS1WardHasBeenDestroyedC2HasTheWard = SysMsg{Id: 2750}

	// The character that acquired {{S1}} ward has been killed.
	TheCharThatAcquiredS1WardHasBeenKilled = SysMsg{Id: 2751}

	// You cannot control because you are too far.
	CantControlTooFar = SysMsg{Id: 2762}

	// You cannot enter because the corresponding alliance channel's maximum number of entrants has been reached.
	YouCannotEnterBecauseMaximumEntrants = SysMsg{Id: 2764}

	// Only the alliance channel leader can attempt entry.
	OnlyAllianceChannelLeaderCanEnter = SysMsg{Id: 2765}

	// Seed of Infinity Stage 1 Attack In Progress.
	SeedOfInfinityStage1AttackInProgress = SysMsg{Id: 2766}

	// Seed of Infinity Stage 2 Attack In Progress.
	SeedOfInfinityStage2AttackInProgress = SysMsg{Id: 2767}

	// Seed of Infinity Conquest Complete.
	SeedOfInfinityConquestComplete = SysMsg{Id: 2768}

	// Seed of Infinity Stage 1 Defense In Progress.
	SeedOfInfinityStage1DefenseInProgress = SysMsg{Id: 2769}

	// Seed of Infinity Stage 2 Defense In Progress.
	SeedOfInfinityStage2DefenseInProgress = SysMsg{Id: 2770}

	// Seed of Destruction Attack in Progress.
	SeedOfDestructionAttackInProgress = SysMsg{Id: 2771}

	// Seed of Destruction Conquest Complete.
	SeedOfDestructionConquestComplete = SysMsg{Id: 2772}

	// Seed of Destruction Defense in Progress.
	SeedOfDestructionDefenseInProgress = SysMsg{Id: 2773}

	// The airship's summon license has been entered. Your clan can now summon the airship.
	TheAirshipSummonLicenseEntered = SysMsg{Id: 2777}

	// You cannot teleport while in possession of a ward.
	YouCannotTeleportWhileInPossessionOfAWard = SysMsg{Id: 2778}

	// You must have a minimum of ({{S1}}) people to enter this Instant Zone. Your request for entry is denied
	YouMustHaveMinimumOfS1PeopleToEnter = SysMsg{Id: 2793}

	// You've already requested a territory war in another territory elsewhere.
	YouAlreadyRequestedTwRegistration = SysMsg{Id: 2795}

	// The clan who owns the territory cannot participate in the territory war as mercenaries.
	TheTerritoryOwnerClanCannotParticipateAsMercenaries = SysMsg{Id: 2796}

	// It is not a territory war registration period, so a request cannot be made at this time.
	NotTerritoryRegistrationPeriod = SysMsg{Id: 2797}

	// The territory war will end in {{S1}}-hour(s).
	TheTerritoryWarWillEndInS1Hours = SysMsg{Id: 2798}

	// The territory war will end in {{S1}}-minute(s).
	TheTerritoryWarWillEndInS1Minutes = SysMsg{Id: 2799}

	// {{S1}}-second(s) to the end of territory war!
	S1SecondsToTheEndOfTerritoryWar = SysMsg{Id: 2900}

	// You cannot force attack a member of the same territory.
	YouCannotAttackAMemberOfTheSameTerritory = SysMsg{Id: 2901}

	// You've acquired the ward. Move quickly to your forces' outpost.
	YouVeAcquiredTheWard = SysMsg{Id: 2902}

	// Territory war has begun.
	TerritoryWarHasBegun = SysMsg{Id: 2903}

	// Territory war has ended.
	TerritoryWarHasEnded = SysMsg{Id: 2904}

	YouRequestedC1ToBeFriend = SysMsg{Id: 2911}

	// Clan {{S1}} has succeeded in capturing {{S2}}'s territory ward.
	ClanS1HasSuccededInCapturingS2TerritoryWard = SysMsg{Id: 2913}

	// The territory war will begin in 20 minutes! Territory related functions (ie: battlefield channel, Disguise Scrolls, Transformations, etc...) can now be used.
	TerritoryWarBeginsIn20Minutes = SysMsg{Id: 2914}

	// Block Checker will end in 5 seconds!
	BlockCheckerEnds5 = SysMsg{Id: 2922}

	// Block Checker will end in 4 seconds!!
	BlockCheckerEnds4 = SysMsg{Id: 2923}

	// You cannot enter a Seed while in a flying transformation state.
	YouCannotEnterSeedInFlyingTransform = SysMsg{Id: 2924}

	// Block Checker will end in 3 seconds!!!
	BlockCheckerEnds3 = SysMsg{Id: 2925}

	// Block Checker will end in 2 seconds!!!!
	BlockCheckerEnds2 = SysMsg{Id: 2926}

	// Block Checker will end in 1 second!!!!!
	BlockCheckerEnds1 = SysMsg{Id: 2927}

	// The {{C1}} team has won.
	TeamC1Won = SysMsg{Id: 2928}

	// {{S2}} unit(s) of the item {{S1}} is/are required.
	S2UnitOfTheItemS1Required = SysMsg{Id: 2961}

	// Being appointed as a Noblesse will cancel all related quests. Do you wish to continue?
	CancelNoblesseQuests = SysMsg{Id: 2964}

	// This is a Payment Request transaction. Please attach the item.
	PaymentRequestNoItem = SysMsg{Id: 2966}

	// The mail limit (240) has been exceeded and this cannot be forwarded.
	CantForwardMailLimitExceeded = SysMsg{Id: 2968}

	// The previous mail was forwarded less than 1 minute ago and this cannot be forwarded.
	CantForwardLessThanMinute = SysMsg{Id: 2969}

	// You cannot forward in a non-peace zone.
	CantForwardNotInPeaceZone = SysMsg{Id: 2970}

	// You cannot forward during exchange.
	CantForwardDuringExchange = SysMsg{Id: 2971}

	// You cannot forward because the private shop or workshop is in progress.
	CantForwardPrivateStore = SysMsg{Id: 2972}

	// You cannot forward during an item enhancement or attribute enhancement.
	CantForwardDuringEnchant = SysMsg{Id: 2973}

	// The item that you're trying to send cannot be forwarded because it isn't proper.
	CantForwardBadItem = SysMsg{Id: 2974}

	// You cannot forward because you don't have enough adena.
	CantForwardNoAdena = SysMsg{Id: 2975}

	// You cannot receive in a non-peace zone location.
	CantReceiveNotInPeaceZone = SysMsg{Id: 2976}

	// You cannot receive during an exchange.
	CantReceiveDuringExchange = SysMsg{Id: 2977}

	// You cannot receive because the private shop or workshop is in progress.
	CantReceivePrivateStore = SysMsg{Id: 2978}

	// You cannot receive during an item enhancement or attribute enhancement.
	CantReceiveDuringEnchant = SysMsg{Id: 2979}

	// You cannot receive because you don't have enough adena.
	CantReceiveNoAdena = SysMsg{Id: 2980}

	// You cannot receive because your inventory is full.
	CantReceiveInventoryFull = SysMsg{Id: 2981}

	// You cannot cancel in a non-peace zone location.
	CantCancelNotInPeaceZone = SysMsg{Id: 2982}

	// You cannot cancel during an exchange.
	CantCancelDuringExchange = SysMsg{Id: 2983}

	// You cannot cancel because the private shop or workshop is in progress.
	CantCancelPrivateStore = SysMsg{Id: 2984}

	// You cannot cancel during an item enhancement or attribute enhancement.
	CantCancelDuringEnchant = SysMsg{Id: 2985}

	// You could not cancel receipt because your inventory is full.
	CantCancelInventoryFull = SysMsg{Id: 2988}

	// When the recipient doesn't exist or the character is deleted, sending mail is not possible.
	RecipientNotExist = SysMsg{Id: 3002}

	// The mail has arrived.
	MailArrived = SysMsg{Id: 3008}

	// Mail successfully sent.
	MailSuccessfullySent = SysMsg{Id: 3009}

	// Mail successfully returned.
	MailSuccessfullyReturned = SysMsg{Id: 3010}

	// Mail successfully cancelled.
	MailSuccessfullyCancelled = SysMsg{Id: 3011}

	// Mail successfully received.
	MailSuccessfullyReceived = SysMsg{Id: 3012}

	// {{C1}} has successfuly enchanted a +{{S2}} {{S3}}.
	C1SuccessfulyEnchantedAS2S3 = SysMsg{Id: 3013}

	// Do you wish to erase the selected mail ?
	DoYouWishToEraseMail = SysMsg{Id: 3014}

	// Please select the mail to be deleted.
	PleaseSelectMailToBeDeleted = SysMsg{Id: 3015}

	// Item selection is possible up to 8.
	ItemSelectedPossibleUpTo8 = SysMsg{Id: 3016}

	// You cannot send mail to yourself.
	YouCantSendMailToYourself = SysMsg{Id: 3019}

	// When not entering the amount for the payment request, you cannot send any mail.
	PaymentAmountNotEntered = SysMsg{Id: 3020}

	// I can feel that the energy being flown in the Kasha's eye is getting stronger rapidly.
	ICanFeelEnergyKashaEyeGettingStrongerRapidly = SysMsg{Id: 3023}

	// Kasha's eye pitches and tosses like it's about to explode.
	KashaEyePitchesTossesExplode = SysMsg{Id: 3024}

	// Payment of {{S1}} Adena was completed by {{S2}}.
	PaymentOfS1AdenaCompletedByS2 = SysMsg{Id: 3025}

	// You cannot use the skill enhancing function on this level. You can use the corresponding function on levels higher than 76Lv .
	YouCannotUseSkillEnchantOnThisLevel = SysMsg{Id: 3026}

	// You cannot use the skill enhancing function in this class. You can use corresponding function when completing the third class change.
	YouCannotUseSkillEnchantInThisClass = SysMsg{Id: 3027}

	// You cannot use the skill enhancing function in this class. You can use the skill enhancing function under off-battle status, and cannot use the function while transforming, battling and on-board.
	YouCannotUseSkillEnchantAttackingTransformedBoat = SysMsg{Id: 3028}

	// {{S1}} returned the mail.
	S1ReturnedMail = SysMsg{Id: 3029}

	// You cannot cancel sent mail since the recipient received it.
	YouCantCancelReceivedMail = SysMsg{Id: 3030}

	// By using the invisible skill, sneak into the Dawn's document storage!
	SneakIntoDawnsDocumentStorage = SysMsg{Id: 3033}

	// Male guards can detect the concealment but the female guards cannot.
	MaleGuardsCanDetectFemalesDont = SysMsg{Id: 3037}

	// Female guards notice the disguises from far away better than the male guards do, so beware.
	FemaleGuardsNoticeBetterThanMale = SysMsg{Id: 3038}

	// {{S1}} did not receive it during the waiting time, so it was returned automatically.
	S1NotReceiveDuringWaitingTimeMailReturned = SysMsg{Id: 3059}

	// Do you want to pay {{S1}} Adena ?
	DoYouWantToPayS1Adena = SysMsg{Id: 3062}

	// Do you really want to forward ?
	DoYouWantToForward = SysMsg{Id: 3063}

	// There is an unread mail.
	UnreadMail = SysMsg{Id: 3064}

	// Current location: Inside the Chamber of Delusion
	LocDelusionChamber = SysMsg{Id: 3065}

	// You cannot use the mail function outside the Peace Zone.
	CantUseMailOutsidePeaceZone = SysMsg{Id: 3066}

	// {{S1}} cancelled the sent mail.
	S1CancelledMail = SysMsg{Id: 3067}

	// The mail was returned due to the exceeded waiting time.
	MailReturned = SysMsg{Id: 3068}

	// Do you really want to cancel the transaction ?
	DoYouWantToCancelTransaction = SysMsg{Id: 3069}

	// {{S1}} acquired the attached item to your mail.
	S1AcquiredAttachedItem = SysMsg{Id: 3072}

	// You have acquired {{S2}} {{S1}}.
	YouAcquiredS2S1 = SysMsg{Id: 3073}

	// The allowed length for recipient exceeded.
	AllowedLengthForRecipientExceeded = SysMsg{Id: 3074}

	// The allowed length for a title exceeded.
	AllowedLengthForTitleExceeded = SysMsg{Id: 3075}

	// The mail limit (240) of the opponent's character has been exceeded and this cannot be forwarded.
	MailLimitExceeded = SysMsg{Id: 3077}

	// You're making a request for payment. Do you want to proceed ?
	YouMakingPaymentRequest = SysMsg{Id: 3078}

	// There are items in the Pet Inventory so you cannot register as an individual store or drop items. Please empty the items in the Pet Inventory.
	ItemsInPetInventory = SysMsg{Id: 3079}

	// You cannot reset the Skill Link because there is not enough Adena.
	CannotResetSkillLinkBecauseNotEnoughAdena = SysMsg{Id: 3080}

	// You cannot receive it because you are under condition that the opponent cannot acquire any Adena for payment
	YouCannotReceiveConditionOpponentCantAcquireAdena = SysMsg{Id: 3081}

	// You cannot send mail to any character that has blocked you.
	YouCannotSendMailToCharBlockYou = SysMsg{Id: 3082}

	// A user currently participating in the Olympiad cannot send party and friend invitations.
	AUserCurrentlyParticipatingInTheOlympiadCannotSendPartyAndFriendInvitations = SysMsg{Id: 3094}

	// You are no longer protected from aggressive monsters.
	YouAreNoLongerProtectedFromAggressiveMonsters = SysMsg{Id: 3108}

	// The couple action was denied.
	CoupleActionDenied = SysMsg{Id: 3119}

	// The request cannot be completed because the target does not meet location requirements.
	TargetDoNotMeetLocRequirements = SysMsg{Id: 3120}

	// The couple action was cancelled.
	CoupleActionCanceled = SysMsg{Id: 3121}

	// The size of the uploaded crest or insignia does not meet the standard requirements.
	WrongSizeUploadedCrest = SysMsg{Id: 3122}

	// {{C1}} is in Private Shop mode or in a battle and cannot be requested for a couple action.
	C1IsInPrivateShopModeOrInABattleAndCannotBeRequestedForACoupleAction = SysMsg{Id: 3123}

	// {{C1}} is fishing and cannot be requested for a couple action.
	C1IsFishingAndCannotBeRequestedForACoupleAction = SysMsg{Id: 3124}

	// {{C1}} is in a battle and cannot be requested for a couple action.
	C1IsInABattleAndCannotBeRequestedForACoupleAction = SysMsg{Id: 3125}

	// {{C1}} is already participating in a couple action and cannot be requested for another couple action.
	C1IsAlreadyParticipatingInACoupleActionAndCannotBeRequestedForAnotherCoupleAction = SysMsg{Id: 3126}

	// {{C1}} is in a chaotic state and cannot be requested for a couple action.
	C1IsInAChaoticStateAndCannotBeRequestedForACoupleAction = SysMsg{Id: 3127}

	// {{C1}} is participating in the Olympiad and cannot be requested for a couple action.
	C1IsParticipatingInTheOlympiadAndCannotBeRequestedForACoupleAction = SysMsg{Id: 3128}

	// {{C1}} is participating in a hideout siege and cannot be requested for a couple action.
	C1IsParticipatingInAHideoutSiegeAndCannotBeRequestedForACoupleAction = SysMsg{Id: 3129}

	// {{C1}} is in a castle siege and cannot be requested for a couple action.
	C1IsInACastleSiegeAndCannotBeRequestedForACoupleAction = SysMsg{Id: 3130}

	// {{C1}} is riding a ship, steed, or strider and cannot be requested for a couple action.
	C1IsRidingAShipSteedOrStriderAndCannotBeRequestedForACoupleAction = SysMsg{Id: 3131}

	// {{C1}} is currently teleporting and cannot be requested for a couple action.
	C1IsCurrentlyTeleportingAndCannotBeRequestedForACoupleAction = SysMsg{Id: 3132}

	// {{C1}} is currently transforming and cannot be requested for a couple action.
	C1IsCurrentlyTransformingAndCannotBeRequestedForACoupleAction = SysMsg{Id: 3133}

	// Requesting approval for changing party loot to "{{S1}}".
	RequestingApprovalChangePartyLootS1 = SysMsg{Id: 3135}

	// Party loot change was cancelled.
	PartyLootChangeCancelled = SysMsg{Id: 3137}

	// Party loot was changed to "{{S1}}".
	PartyLootChangedS1 = SysMsg{Id: 3138}

	// {{C1}} is currently dead and cannot be requested for a couple action.
	C1IsCurrentlyDeadAndCannotBeRequestedForACoupleAction = SysMsg{Id: 3139}

	// The {{S2}}'s attribute was successfully bestowed on {{S1}}, and resistance to {{S3}} was increased.
	TheS2AttributeWasSuccessfullyBestowedOnS1ResToS3Increased = SysMsg{Id: 3144}

	// If you are not resurrected within {{S1}} minutes, you will be expelled from the instant zone.
	YouWillBeExpelledInS1 = SysMsg{Id: 3147}

	// You have requested a couple action with {{C1}}.
	YouHaveRequestedCoupleActionC1 = SysMsg{Id: 3150}

	// {{S1}}'s {{S2}} attribute was removed, and resistance to {{S3}} was decreased.
	S1S2AttributeRemovedResistanceS3Decreased = SysMsg{Id: 3152}

	// You do not have enough funds to cancel this attribute.
	YouDoNotHaveEnoughFundsToCancelAttribute = SysMsg{Id: 3156}

	// +{{S1}}{{S2}}'s {{S3}} attribute was removed, so resistance to {{S4}} was decreased.
	S1S2S3AttributeRemovedResistanceToS4Decreased = SysMsg{Id: 3160}

	// The {{S3}}'s attribute was successfully bestowed on +{{S1}}{{S2}}, and resistance to {{S4}} was increased.
	TheS3AttributeBestowedOnS1S2ResistanceToS4Increased = SysMsg{Id: 3163}

	// {{C1}} is set to refuse couple actions and cannot be requested for a couple action.
	C1IsSetToRefuseCoupleActions = SysMsg{Id: 3164}

	// {{C1}} is set to refuse party requests and cannot receive a party request.
	C1IsSetToRefusePartyRequest = SysMsg{Id: 3168}

	// {{C1}} is set to refuse duel requests and cannot receive a duel request.
	C1IsSetToRefuseDuelRequest = SysMsg{Id: 3169}

	// You currently do not have any Recommendations.
	YouCurrentlyDoNotHaveAnyRecommendations = SysMsg{Id: 3206}

	// You obtained {{S1}} Recommendations
	YouObtainedS1Recommendations = SysMsg{Id: 3207}

	// {{S1}} was successfully added to your Contact List.
	S1SuccessfullyAddedToContactList = SysMsg{Id: 3214}

	// The name {{S1}}% doesn't exist. Please try another name.
	NameS1NotExistTryAnotherName = SysMsg{Id: 3215}

	// The name already exists on the added list.
	NameAlreadyExistOnContactList = SysMsg{Id: 3216}

	// The name is not currently registered.
	NameNotRegisteredOnContactList = SysMsg{Id: 3217}

	// {{S1}} was successfully deleted from your Contact List.
	S1SuccesfullyDeletedFromContactList = SysMsg{Id: 3219}

	// You cannot add your own name.
	CannotAddYourNameOnContactList = SysMsg{Id: 3221}

	// The maximum number of names (100) has been reached. You cannot register any more.
	ContactListLimitReached = SysMsg{Id: 3222}

	// The maximum matches you can participate in 1 week is 70.
	MaxOlyWeeklyMatchesReached = SysMsg{Id: 3224}

	// The total number of matches that can be entered in 1 week is 60 class irrelevant individual matches, 30 specific matches, and 10 team matches.
	MaxOlyWeeklyMatchesReached60NonClassed30Classed10Team = SysMsg{Id: 3225}

	// You cannot move while speaking to an NPC. One moment please.
	CannotMoveWhileSpeakingToAnNpc = SysMsg{Id: 3226}

	// Arcane Shield decreased your MP by $1 instead of HP.
	ArcaneShieldDecreasedYourMpByS1InsteadOfHp = SysMsg{Id: 3255}

	// You have acquired {{S1}} EXP (Bonus: {{S2}}) and {{S3}} SP (Bonus: {{S4}}).
	YouEarnedS1ExpBonusS2AndS3SpBonusS4 = SysMsg{Id: 3259}

	// MP became 0 and the Arcane Shield is disappearing.
	MpBecame0ArcaneShieldDisappearing = SysMsg{Id: 3256}

	// You cannot use the skill because the servitor has not been summoned.
	CannotUseSkillWithoutServitor = SysMsg{Id: 3260}

	// You have {{S1}} match(es) remaining that you can participate in this week ({{S2}} 1 vs 1 Class matches, {{S3}} 1 vs 1 matches, & {{S4}} 3 vs 3 Team matches).
	YouHaveS1MatchesRemainingThatYouCanPartecipateInThisWeekS2ClassedS3NonClassedS4Team = SysMsg{Id: 3261}

	// Enchant failed. The enchant level for the corresponding item will be exactly retained.
	SafeEnchantFailed = SysMsg{Id: 6004}

	// You cannot bookmark this location because you do not have a My Teleport Flag.
	YouCannotBookmarkThisLocationBecauseYouDoNotHaveAMyTeleportFlag = SysMsg{Id: 6501}

	// The evil Thomas D. Turkey has appeared. Please save Santa.
	ThomasDTurkeyAppeared = SysMsg{Id: 6503}

	// You won the battle against Thomas D. Turkey. Santa has been rescued.
	ThomasDTurkeyDefeted = SysMsg{Id: 6504}

	// You did not rescue Santa, and Thomas D. Turkey has disappeared.
	ThomasDTurkeyDisappeared = SysMsg{Id: 6505}
)
