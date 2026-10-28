package messages_service

import "context"

type messageEvent struct {
	SectionID     string  `json:"sectionId"`
	FromSectionID *string `json:"fromSectionId,omitempty"`
}

func (s *MessagesService) publishMessageEvent(ctx context.Context, eventType string, event messageEvent) {
	boardID, err := s.messagesRepository.GetBoardIdBySectionId(ctx, event.SectionID)
	if err != nil {
		return
	}

	_ = s.eventBroadcaster.Publish(boardID, eventType, event)
}
