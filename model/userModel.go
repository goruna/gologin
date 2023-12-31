package model

type UserModel struct {
	AiId            int64 `gorm:"column:aiId; primary_key;"`
	Id string `json:"id"`
	Nickname        string `json:"nickname"`
	AvatarUrl       string `json:"avatarUrl"`
	Mobile          string
	MobilePrefix    string
	Password        string
	PasswordSalt    string
	PasswordVersion int8
	WxUnionId       string
	WxUnionIdChina  string
	WxNickname      string
	InviteCode      string
	MallCid         string
	LeaderCid       string
	LogoutReason    string
	IsLogout        bool
	Os              string
	DeviceId        string
	Sex             int8
	CardState       bool
	WxId            string
	IsLock          bool
	BaseModel
}

func (UserModel) TableName() string {
	return "u_user"
}
