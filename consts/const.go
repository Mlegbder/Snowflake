package consts

const (
	Epoch             = int64(1577808000000)                           // 设置起始时间(时间戳/毫秒)：2020-01-01 00:00:00，有效期69年
	TimestampBits     = uint(41)                                       // 时间戳占用位数
	DatacenterIdBits  = uint(2)                                        // 数据中心id所占位数
	WorkerIdBits      = uint(7)                                        // 机器id所占位数
	SequenceBits      = uint(12)                                       // 序列所占的位数
	TimestampMax      = int64(-1 ^ (-1 << TimestampBits))              // 时间戳最大值
	DatacenterIdMax   = int64(-1 ^ (-1 << DatacenterIdBits))           // 支持的最大数据中心id数量
	WorkerIdMax       = int64(-1 ^ (-1 << WorkerIdBits))               // 支持的最大机器id数量
	SequenceMask      = int64(-1 ^ (-1 << SequenceBits))               // 支持的最大序列id数量
	WorkerIdShift     = SequenceBits                                   // 机器id左移位数
	DatacenterIdShift = SequenceBits + WorkerIdBits                    // 数据中心id左移位数
	TimestampShift    = SequenceBits + WorkerIdBits + DatacenterIdBits // 时间戳左移位数
)
