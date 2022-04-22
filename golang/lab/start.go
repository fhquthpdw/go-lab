package main

import (
	"time"

	"github.com/opentracing/opentracing-go"
)

func T1() {
	tracer := opentracing.GlobalTracer()

	span := tracer.StartSpan("T1")
	println("start sleeping ...")
	time.Sleep(2 * time.Second)
	println("wake up now...")
	span.Finish()
}
