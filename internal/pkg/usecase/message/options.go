package message

type option func(*usecaseMessage)

func WithPartition(n uint32) option {
	return func(u *usecaseMessage) {
		u.segmentSize = n
	}
}
