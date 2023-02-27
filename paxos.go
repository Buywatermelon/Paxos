package main

import (
	"testing"
)

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
	// 1001、1002、1003是接受者id
	acceptorIds := []int{1001, 1002, 1003}
	// 2001是学习者id
	learnerIds := []int{2001}
	acceptors, learners := start(acceptorIds, learnerIds)

	defer cleanup(acceptors, learners)

	p := &Proposer{
		id:        1,
		acceptors: acceptorIds,
	}

	value := p.propose("hello world")
	if value != "hello world" {
		t.Errorf("value = %s, excepted %s", value, "helo world")
	}

	learnValue := learners[0].chosen()
	if learnValue != value {
		t.Errorf("learnValue = %s, exceptd %s", learnValue, "hello world")
	}
}

func TestTwoProposers(t *testing.T) {

}
