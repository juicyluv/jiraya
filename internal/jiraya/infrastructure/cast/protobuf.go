package cast

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func Timestamp(v *time.Time) *timestamppb.Timestamp {
	if v == nil {
		return nil
	}

	return timestamppb.New(*v)
}

func Time(v *timestamppb.Timestamp) *time.Time {
	if v == nil {
		return nil
	}

	t := v.AsTime()
	return &t
}
