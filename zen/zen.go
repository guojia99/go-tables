package thezen

// video -> https://www.gophercon.org.il/
// zn -> https://cloud.tencent.com/developer/article/1591747
// en -> https://the-zen-of-go.netlify.app/

var ZenList = []string{
	"Each package fulfils a single purpose",
	"Handle errors explicitly",
	"Return early rather than nesting deeply",
	"Leave concurrency to the caller",
	"Before you launch a goroutine, know when it will stop",
	"Avoid package level state",
	"Simplicity matters",
	"Write tests to lock in the behaviour of your package’s API",
	"If you think it’s slow, first prove it with a benchmark",
	"Moderation is a virtue",
	"Maintainability counts",
}

var ZenListZn = []string{
	"每个包实现单一目标",
	"明确处理错误",
	"尽早 return，不要深陷",
	"并发权留给调用者",
	"在启动 goroutine 之前，要知道它什么时候会停止",
	"避免包级别的状态",
	"简单性很重要",
	"编写测试以确认包 API 的行为",
	"如果你认为速度缓慢，先通过基准测试进行验证",
	"节制是一种美德",
}
