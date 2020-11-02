package xcache

type optionF func(*option)

type option struct {
	isOpenMonitor bool // 是否打开监管
	monitorGap    int  // 监管间隔
}

// 开启监管者,定时间隔gap
func WithOpenMonitor(gap int) optionF {
	return func(opt *option) {
		opt.isOpenMonitor = true
		opt.monitorGap = gap
	}
}