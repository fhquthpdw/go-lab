package main

import "fmt"

type HashFunc func()

type options struct {
	hashFunc    HashFunc
	bucketCount uint64
}

type Option interface {
	apply(*options)
}

type Bucket struct {
	count uint64
}

func (b Bucket) apply(opts *options) {
	opts.bucketCount = b.count
}

func WithBucketCount(count uint64) Option {
	return Bucket{
		count: count,
	}
}

type Hash struct {
	hashFunc HashFunc
}

func (h Hash) apply(opts *options) {
	opts.hashFunc = h.hashFunc
}

func WithHashFunc(hashFunc HashFunc) Option {
	return Hash{
		hashFunc: hashFunc,
	}
}

func NewCacheOption(opts ...Option) (*options, error) {
	c := &options{
		hashFunc:    nil,
		bucketCount: 0,
	}

	for _, opt := range opts {
		opt.apply(c)
	}

	return c, nil
}

func main() {
	o, _ := NewCacheOption(WithBucketCount(10), WithHashFunc(func() {}))
	fmt.Println(o)
}
