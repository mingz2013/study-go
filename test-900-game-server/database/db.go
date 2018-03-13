package database

type User struct {
	Id       int
	DeviceId string
}

var UserDB []User
var IdIndex int

func init() {
	UserDB = make([]User, 1000)
	IdIndex = 0
}

func GetOneIdIndex() (int) {
	IdIndex += 1
	return IdIndex
}

func NewUser(deviceId string) (*User) {
	id := GetOneIdIndex()
	u := User{id, deviceId}
	return &u
}

func FindUserByDeviceId(deviceId string) (User) {
	for _, user := range UserDB {
		if user.DeviceId == deviceId {
			return user
		}
	}

	u := NewUser(deviceId)
	UserDB = append(UserDB, *u)
	return *u
}
