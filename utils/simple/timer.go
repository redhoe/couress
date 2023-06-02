package simple

import "time"

func SecIn64ToTimeStr(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

func SecInt64ToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

func Int64ToTime(MillisecondStamp int64) time.Time {
	return time.UnixMilli(MillisecondStamp)
}
