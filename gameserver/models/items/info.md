```sh
Id - уникальный идентификаторы предмета в клиенте
ItemType - тип предмета, пакет weaponType
Name - имя предмета
Icon - иконка предмета
SlotBitType - тип слота куда можно одеть предмет , подробнее slot_bit_type.go
ArmorType - тип брони , пакет armorType
EtcItemType - тип предмета, пакет etcItemType
ItemMultiSkillList - ?
RecipeId - если предмет это рецепт , то recipeId ссылка на id в файле recipe.txt
Weight - вес предмета
ConsumeType - 1)consume_type_normal - нельзя складывать,  2)consume_type_stackable - можно складывать, 3)consume_type_asset ?? по идее они тоже складываются
SoulShotCount - количество потребляемых soulShot
SpiritShotCount - количество потребляемых spiritShot
DropPeriod - минимальный? период дропа в секундах?
DefaultPrice - цена продажи в магазин
ItemSkill - !Какой скилл дает предмет
CriticalAttackSkill - !При крите какой скилл будет использован
AttackSkill - !При ударе каким скиллом может удалить
MagicSkill - 
ItemSkillEnchantedFour - !Когда предмет заточен на +4 какой скилл давать
MaterialType - !Тип материала (хз зачем это и как используется)
CrystalType - тип кристала 
CrystalCount - количество кристаллов
IsTrade - можно ли продавать предмет
IsDrop - можно ли выкинуть предмет?
IsDestruct - можно ли удалять предмет
IsPrivateStore - можно ли продавать предмет 
KeepType - ?
RandomDamage - ? !Разброс при демедже
WeaponType - тип оружия , пакет weaponType
HitModify - ?
AvoidModify
ShieldDefense
ShieldDefenseRate
AttackRange !Дистанция нанесения физической физ. атаки
ReuseDelay - !Время повторной атаки(к примеру у лука дольше атака чем у даггера)
MpConsume - !При физ атаке потребление MP к примеру как у лука
Durability - !Время предмета (после чего он уничтрожится) т.е. шадоу
MagicWeapon - это магическое оружие
EnchantEnable - !Разрешение на заточку предмета
ElementalEnable - !Можно ли вставить АТТ в предмет
ForNpc - !Если 0, то этот итем не могут использовать петы. Если 1, то этот предмет они могут одевать/использовать 
IsOlympiadCanUse - можно ли использовать на олимпиаде
IsPremium - это премиум предмет
BonusStats 
DefaultAction - action , подробнее в default_action.go
InitialCount
ImmediateEffect
CapsuledItems
DualFhitRate
DamageRange - !40 и 120 это радиус и угол
Enchanted  
BaseAttributeAttack ! Базовая атака атт
BaseAttributeDefend !Базовая защита атт
UnequipSkill !При надевании сниманет скилл
ItemEquipOption
CanMove !Предмет обладает подвижностью (к примеру агатионы, или адена оО)
DelayShareGroup
Blessed
ReducedSoulshot !Не понятно что это, как-то связано (судя по названию) с уменьшением юза кол-ва сосок
ExImmediateEffect
UseSkillDistime
Period
EquipReuseDelay !Повторное использование предмета (к примеру агатионов)
Price - цена в (0 - адена)  # адена (id 57) price=1 чего так не знаю
```