package mnc

type Member struct {
    Name   string
    RoomID int
    Entity Entity
}

func NewMember(name string, roomId int, entity Entity) *Member {
    member := &Member{
        Name: name,
        RoomID: roomId,
        Entity: entity,
    }
    return member
}
