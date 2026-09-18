import "errors"

type RoomType string

const (
	Single RoomType = "single"
	Double RoomType = "double"
	Suite  RoomType = "suite"
)

type RoomStatus string

const (
	Free        RoomStatus = "free"
	Booked      RoomStatus = "booked"
	Maintenance RoomStatus = "maintenance"
)

type HotelRoom struct {
	Type   RoomType
	Status RoomStatus
	Price  float64
}

func bookRoom(rooms map[string]HotelRoom, roomNumber string) error {
	room, ok := rooms[roomNumber]
	if !ok {
		return errors.New("номер не найден")
	}
	if room.Status != Free {
		return errors.New("номер недоступен")
	}
	room.Status = Booked
	rooms[roomNumber] = room
	return nil
}