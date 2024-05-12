package message

import (
	"github.com/gvidow/go-chat-transport/internal/pkg/entity"
	"github.com/gvidow/go-chat-transport/pkg/math"
)

type option func(*usecaseMessage)

func WithPartitionBySize(countByte uint32) option {
	return func(u *usecaseMessage) {
		u.segmentation = func(mes *entity.Message) []entity.Segment {
			countSegment := (len(mes.Content) + int(countByte) - 1) / int(countByte)
			res := make([]entity.Segment, 0, countSegment)

			for i := 0; i < countSegment; i++ {
				res = append(res, entity.Segment{
					Time: mes.Time,
					Size: countSegment,
					Num:  i,
					Data: mes.Content[i*int(countByte) : math.MinInt(len(mes.Content), i*int(countByte)+int(countByte))],
				})
			}

			return res
		}
	}
}

func WithCustomPartition(cp func(*entity.Message) []entity.Segment) option {
	return func(u *usecaseMessage) {
		u.segmentation = cp
	}
}
