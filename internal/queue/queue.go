package queue

type EventQueue struct {
	EventID string
	Size    int
}

type EnqueueRequest struct {
	EventID string
	UserID  string
}

type EnqueueResponse struct {
	TraceID string
}

func New(eventID string, size int) *EventQueue {
	return &EventQueue{
		EventID: eventID,
		Size:    size,
	}
}

type Service struct {
}

func (s *Service) Enqueue(r EnqueueRequest) {

}
