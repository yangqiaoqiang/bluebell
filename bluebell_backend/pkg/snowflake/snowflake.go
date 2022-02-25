package snowflake

import (
	sf "github.com/bwmarrin/snowflake"

	"time"
)

var node *sf.Node

// 需传入当前的机器ID
func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}

// GetID 返回生成的id值
func GenID() int64 {
	return node.Generate().Int64()
}
