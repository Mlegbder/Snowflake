package util

import (
	"Snowflake/consts"
	"fmt"
	"sync"
	"time"
)

type Snowflake struct {
	sync.Mutex         // 锁
	Timestamp    int64 // 时间戳 ，毫秒
	WorkerId     int64 // 工作节点
	DatacenterId int64 // 数据中心机房id
	Sequence     int64 // 序列号
}

func (s *Snowflake) NextVal() int64 {
	s.Lock()
	now := time.Now().UnixNano() / 1000000 // 转毫秒
	if s.Timestamp == now {
		// 当同一时间戳（精度：毫秒）下多次生成id会增加序列号
		s.Sequence = (s.Sequence + 1) & consts.SequenceMask
		if s.Sequence == 0 {
			// 如果当前序列超出12bit长度，则需要等待下一毫秒
			// 下一毫秒将使用sequence:0
			for now <= s.Timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		// 不同时间戳（精度：毫秒）下直接使用序列号：0
		s.Sequence = 0
	}
	t := now - consts.Epoch
	if t > consts.TimestampMax {
		s.Unlock()
		fmt.Print("epoch must be between 0 and %d", consts.TimestampMax-1)
		return 0
	}
	s.Timestamp = now
	r := int64((t)<<consts.TimestampShift | (s.DatacenterId << consts.DatacenterIdShift) | (s.WorkerId << consts.WorkerIdShift) | (s.Sequence))
	s.Unlock()
	return r
}
