package queue

import "time"

type Event struct {
	ID        string
	Name      string
	Artist    string
	Date      time.Time
	CartStart time.Time
}

type EvtStorage interface {
	Create(e Event) (Event, error)
	Get(id string) (Event, error)
}

type EvtService struct {
	EvtStorage
}

func (es *EvtService) CreateEvt(event Event) (Event, error) {
	return es.Create(event)
}

func (es *EvtService) GetEvt(id string) (Event, error) {
	return es.Get(id)
}
