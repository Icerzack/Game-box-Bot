package room

import (
	"math"
	"math/rand"
	"time"
)

func CreateNewRoom(roomsHub *RoomsHub, creatorID int, username string) int32 {
	var randomInt int32

	for {
		rand.Seed(time.Now().UnixNano())
		randomInt = rand.Int31n(int32(math.MaxInt32))
		if _, ok := roomsHub.Rooms[randomInt]; !ok {
			break
		}
	}

	usersList := make([]*User, 0, 1)

	user := &User{
		UserID:   creatorID,
		UserName: username,
		Value:    "-",
	}

	usersList = append(usersList, user)

	newRoom := &Room{
		RoomID: randomInt,
		Users:  usersList,
	}

	roomsHub.Rooms[randomInt] = newRoom

	return randomInt
}

func ModifyUserValueInRoom(roomsHub *RoomsHub, roomID int32, userID int, newValue string) {
	if room, ok := roomsHub.Rooms[roomID]; ok {
		for _, user := range room.Users {
			if user.UserID == userID {
				user.Value = newValue
				break
			}
		}
	}
}

func JoinRoomWithID(roomsHub *RoomsHub, roomID int32, userID int, username string) bool {
	if room, ok := roomsHub.Rooms[roomID]; ok {
		newUser := &User{
			UserID:   userID,
			UserName: username,
			Value:    "-",
		}
		room.Users = append(room.Users, newUser)
		return true
	}
	return false
}

// DeleteUserFromRoom возвращает флаг, если true, тогда данная комната должна была удалена, так как она пустая.
func DeleteUserFromRoom(roomsHub *RoomsHub, roomID int32, userID int) bool {
	room, ok := roomsHub.Rooms[roomID]
	if !ok {
		return false
	}

	for i, user := range room.Users {
		if user.UserID == userID {
			room.Users = append(room.Users[:i], room.Users[i+1:]...)
			if len(room.Users) == 0 {
				delete(roomsHub.Rooms, roomID)
				return true
			}
			break
		}
	}
	return false
}
