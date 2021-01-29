package main

import (
	"log"
)

type ConnLimiter struct {
	concurrentConn int //设定的最大流控值
	bucket         chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc)} // buf chan }
}

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the rate limitation")
		return false
	}
	cl.bucket <- 1
	return true
}

func (cl *ConnLimiter) ReleaseConn() {

	c := <-cl.bucket
	log.Printf("New connection comming :%v", c)
}
