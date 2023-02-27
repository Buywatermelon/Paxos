package main

type Proposer struct {
	// 服务器id
	id int
	// 当前提议者已知的最大轮次
	round int
	// 提案编号
	number int
	// 接受者id列表
	acceptors []int
}
