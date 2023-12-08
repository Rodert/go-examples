package main

type Scenario struct {
	Name        string
	Description []string
	Examples    []string
	RunExample  func()
}

var s1 = &Scenario{
	Name: "s1",
	Description: []string{
		"简单并发执行任务",
	},
	Examples: []string{
		"比如并发的请求后端某个接口",
	},
	RunExample: RunScenario1,
}

var s2 = &Scenario{
	Name: "s2",
	Description: []string{
		"持续一定时间的高并发模型",
	},
	Examples: []string{
		"在规定时间内，持续的高并发请求后端服务， 防止服务死循环",
	},
	RunExample: RunScenario2,
}

var s3 = &Scenario{
	Name: "s3",
	Description: []string{
		"基于大数据量的并发任务模型, goroutine worker pool",
	},
	Examples: []string{
		"比如技术支持要给某个客户删除几个TB/GB的文件",
	},
	RunExample: RunScenario3,
}

var s4 = &Scenario{
	Name: "s4",
	Description: []string{
		"等待异步任务执行结果(goroutine+select+channel)",
	},
	Examples: []string{
		"",
	},
	RunExample: RunScenario4,
}

var s5 = &Scenario{
	Name: "s5",
	Description: []string{
		"定时的反馈结果(Ticker)",
	},
	Examples: []string{
		"比如测试上传接口的性能，要实时给出指标: 吞吐率，IOPS,成功率等",
	},
	RunExample: RunScenario5,
}

var Scenarios []*Scenario

func init() {
	Scenarios = append(Scenarios, s1)
	Scenarios = append(Scenarios, s2)
	Scenarios = append(Scenarios, s3)
	Scenarios = append(Scenarios, s4)
	Scenarios = append(Scenarios, s5)
}
