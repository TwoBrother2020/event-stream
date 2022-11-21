package broker

import "event-stream/protocol"

func createCompletedEvent(uint64, string) protocol.Event {

	return protocol.Event{Value: &protocol.Event_JobCompleted{JobCompleted: &protocol.JobCompleted{}}}
}
