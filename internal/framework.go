package internal

type Router interface {
	Register(logic func() error)
	Start()
}

// app
type App interface {
	Run()
	Stop() error
}

// 管理
type Manager interface {
	Get(key string) (interface{}, error)
	Register(key string, provider interface{})
}

type Creator interface {
	Create() interface{}
}

type Newer interface {
	New() interface{}
}
