package room

type RoomsHub struct {
	Rooms map[int32]*Room
}

type Room struct {
	RoomID int32
	Users  []*User
}

type User struct {
	UserID   int
	UserName string
	Value    string
}
