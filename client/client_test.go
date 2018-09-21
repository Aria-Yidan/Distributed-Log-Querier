package main

import (
	"strconv"
	"testing"
)

func TestClient_FrequentPtn(t *testing.T) {
	testCases := []struct {
		pattern     string
		matchNumber int
	}{
		{"an", 25},
		{"pi", 2},
		{"mp", 13},
		{"^[a-zA-Z]", 43},
		{"%20", 2},
		{"at", 32},
		{"ter", 14},
	}
	nodeServerList := readServerFromFile("serverInfo.txt")
	for _, testcase := range testCases {
		c := make(chan GrepRes)
		var strMessage string = testcase.pattern
		for _, node := range nodeServerList {
			filename := "unitTest." + strconv.Itoa(node.nodeId) + ".log"
			go connectToServer(node, strMessage, filename, c)
		}
		totalCnt := 0
		for i := 0; i < len(nodeServerList); i++ {
			res := <-c
			if res.MatchSuc {
				totalCnt += res.MatchCnt
			}
		}
		t.Log("Total Count:", totalCnt)
		t.Log("Expected Count:", testcase.matchNumber)
		if totalCnt != testcase.matchNumber {
			t.Error("Fail the test for pattern:", testcase.pattern)
		} else {
			t.Log("Pass the test for pattern:", testcase.pattern)
		}
	}
}
