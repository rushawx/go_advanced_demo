package stat

import (
	"0-hello/pkg/event"
	"log"
)

type StatServiceDeps struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

type StatService struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

func NewStatService(deps StatServiceDeps) *StatService {
	return &StatService{
		EventBus:       deps.EventBus,
		StatRepository: deps.StatRepository,
	}
}

func (s *StatService) AddClick(linkId uint) {
	for msg := range s.EventBus.Subscribe() {
		if msg.Type == event.EventLinkVisited {
			id, ok := msg.Data.(uint)
			if !ok {
				log.Fatalln("invalid data type: ", msg.Data)
				continue
			}
			s.StatRepository.AddClick(id)
		}
	}
}