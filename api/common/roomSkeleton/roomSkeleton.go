package roomSkeleton

import "casino_common/api/common/deskSkeleton"

type RoomSkeletonApi interface {
	DissolveDesk(desk deskSkeleton.DeskSkeletonApi) error //解散房间
}
