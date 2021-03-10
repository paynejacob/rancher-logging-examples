package index

import "time"

const MaxIndexLogs = 100

type Index struct {
	Name     string
	FirstLog time.Time `json:"first_log"`
	LastLog  time.Time `json:"last_log"`
	LogCount int64     `json:"log_count"`
	Logs     [][]byte  `json:"-"`
}

func NewIndex(name string) Index {
	return Index{
		Name: name,
		Logs: make([][]byte, MaxIndexLogs),
	}
}

func WriteLog(index *Index, _log []byte) {
	if index.FirstLog.IsZero() {
		index.FirstLog = time.Now()
	}
	index.LastLog = time.Now()

	index.LogCount++

	if index.LogCount > MaxIndexLogs {
		index.Logs = append(index.Logs[1:], _log)
	} else {
		index.Logs[index.LogCount-1] = _log
	}
}

