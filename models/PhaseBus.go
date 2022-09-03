package models

type PhaseEventBus struct {
	events map[string]func()
}

func NewPhaseEventBus() *PhaseEventBus {

	phaseEventBus := PhaseEventBus{}
	phaseEventBus.events = map[string]func(){}
	return &phaseEventBus
}

func (re PhaseEventBus) RegisterEventHandler(eventID string, handler func()) {
	re.events[eventID] = handler
}

func (re PhaseEventBus) Fire(eventID string) {
	f := re.events[eventID]
	f()
}
