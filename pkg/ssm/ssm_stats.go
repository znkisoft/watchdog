package ssm

import (
	"time"

	"github.com/shirou/gopsutil/v3/host"
)

func Uptime() (uint64, error) {
	boottime, err := host.BootTime()
	if err != nil {
		return 0, err
	}
	uptime := uint64(time.Now().Unix()) - boottime
	return uptime, nil
}
