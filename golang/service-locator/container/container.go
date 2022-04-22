package serviceContainer

import "fmt"

var Services *serviceLocator

const (
	CacheSvcToken = "CACHE"
	MsgSvcToken   = "MSG"
)

func init() {
	Services = newServiceLocator()

	// 设置依赖的 Cache Service
	// redis
	/*
		Services.Register(CacheSvcToken, func() interface{} {
			return RedisCache{}
		})
	*/
	// memcached
	Services.Register(CacheSvcToken, func() interface{} {
		return MemCache{}
	})

	// 设置依赖 Msg Service
	// email
	/*
		Services.Register(MsgSvcToken, func() interface{} {
			return EmailMsg{}
		})
	*/
	// chat
	Services.Register(MsgSvcToken, func() interface{} {
		return ChatMsg{}
	})
}

// Service Locator 中，生成每个 Service 的方法
// 返回是一个 interface，需要做类型断言和转换
type Factory func() interface{}

// 一个简易的 Service Locator
type serviceLocator struct {
	Factories map[string]Factory
}

func newServiceLocator() *serviceLocator {
	sl := serviceLocator{
		Factories: make(map[string]Factory),
	}
	return &sl
}

// 向 Service Locator 容器中注册一个依赖服务
func (sl *serviceLocator) Register(token string, factory Factory) *serviceLocator {
	sl.Factories[token] = factory
	return sl
}

// 从 Service Locator 容器的取一个依赖服务
func (sl *serviceLocator) GetService(token string) interface{} {
	svc := sl.Factories[token]
	return svc()
}

// 示例依赖服务
// Cache
type Cache interface {
	Set(key, value string) error
	Get(key string) string
}

// Cache - Redis
type RedisCache struct {
	//
}

func (rc RedisCache) Set(key, value string) error {
	fmt.Println("set cache by redis")
	return nil
}

func (rc RedisCache) Get(key string) string {
	fmt.Println("get cache by redis")
	return "cache from redis"
}

// Cache - Memcache
type MemCache struct {
	//
}

func (rc MemCache) Set(key, value string) error {
	fmt.Println("set cache by memcached")
	return nil
}

func (rc MemCache) Get(key string) string {
	fmt.Println("get cache by memcached")
	return "cache from memcached"
}

// Msg
type Msg interface {
	Send(to, msg string) error
}

// Msg - Email
type EmailMsg struct {
	//
}

func (em EmailMsg) Send(to, msg string) error {
	fmt.Println("send msg by email")
	return nil
}

// Msg - Chat
type ChatMsg struct {
	//
}

func (em ChatMsg) Send(to, msg string) error {
	fmt.Println("send msg by chat")
	return nil
}
