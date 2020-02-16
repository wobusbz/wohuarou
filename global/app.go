package global

import (
	"github.com/wobusbz/wohuarou/util"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func init() {
	App.Date = time.Now()
}

func inferRootDir() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var infer func(d string) string
	infer = func(d string) string {
		if util.Exists(d + "/config") {
			return d
		}
		return infer(filepath.Dir(d))
	}
	return infer(cwd)
}

var App = &app{}

type app struct {
	Name    string
	Version string
	Date    time.Time

	// 项目根目录
	RootDir string

	// 启动时间
	LaunchTime time.Time
	Uptime     time.Duration

	sync.Mutex
}

func (a *app) SetUptime() {
	a.Lock()
	a.Uptime = time.Now().Sub(a.LaunchTime)
	a.Unlock()
}
