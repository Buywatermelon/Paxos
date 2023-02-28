package main

import "testing"

func start(acceptorIds []int, learnerIds []int) ([]*Acceptor, []*Learner) {
	acceptors := make([]*Acceptor, 0)
	for _, aid := range acceptorIds {
		a := newAcceptor(aid, learnerIds)
		acceptors = append(acceptors, a)
	}

	learners := make([]*Learner, 0)
	for _, lid := range learnerIds {
		l := newLearner(lid, acceptorIds)
		learners = append(learners, l)
	}

	return acceptors, learners
}

func cleanup(acceptors []*Acceptor, learners []*Learner) {
	for _, a := range acceptors {
		a.close()
	}

	for _, l := range learners {
		l.close()
	}
}

func TestSingleProposer(t *testing.T) {
	// 10010、10020、10030 是接受者id
	acceptorIds := []int{10010, 10020, 10030}
	// 20010 是学习者id
	learnerIds := []int{20010}
	acceptors, learners := start(acceptorIds, learnerIds)

	defer cleanup(acceptors, learners)

	p := &Proposer{
		id:        1,
		acceptors: acceptorIds,
	}

	value := p.propose("hello world", t)
	if value != "hello world" {
		t.Errorf("value = %s, excepted %s", value, "hello world")
	}

	learnValue := learners[0].chosen()
	if learnValue != value {
		t.Errorf("learnValue = %s, exceptd %s", learnValue, "hello world")
	}
}

func TestTwoProposers(t *testing.T) {
	// 10010、10020、10030 是接受者id
	acceptorIds := []int{10010, 10020, 10030}
	// 20010 是学习者id
	learnerIds := []int{20010}
	acceptors, learners := start(acceptorIds, learnerIds)

	defer cleanup(acceptors, learners)

	p1 := &Proposer{
		id:        1,
		acceptors: acceptorIds,
	}
	v1 := p1.propose("hello world", t)

	p2 := &Proposer{
		id:        2,
		acceptors: acceptorIds,
	}
	v2 := p2.propose("hello book", t)

	if v1 != v2 {
		t.Errorf("value1 = %s, value12 = %s", v1, v2)
	}

	learnValue := learners[0].chosen()
	if learnValue != v2 {
		t.Errorf("learnValue = %s, exceptd %s", learnValue, "hello world")
	}
}
