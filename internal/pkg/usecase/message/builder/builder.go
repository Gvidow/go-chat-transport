package builder

import (
	"strings"

	"github.com/gvidow/go-chat-transport/internal/pkg/entity"
)

type messageBuilder struct {
	id              uint64
	segmentContents []string
	count           int
	username        string
}

func NewBuilder(id uint64, username string, size int) *messageBuilder {
	return &messageBuilder{
		id:              id,
		username:        username,
		segmentContents: make([]string, size),
	}
}

func (m *messageBuilder) AddSegment(segment *entity.Segment) {
	m.count++
	m.segmentContents[segment.Num] = segment.Data
}

func (m *messageBuilder) Build() *entity.MessageWithErrorFlag {
	mes := &entity.MessageWithErrorFlag{
		Message: entity.Message{
			Time:     m.id,
			Username: m.username,
		},
	}

	if m.count != len(m.segmentContents) {
		mes.Error = true
		return mes
	}

	b := &strings.Builder{}
	for _, segment := range m.segmentContents {
		b.WriteString(segment)
	}

	mes.Content = b.String()
	return mes
}
