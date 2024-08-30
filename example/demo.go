package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/walcher-mm/go-netflow/pkg/netflow"
)

func main() {

	targetName := os.Args[1]

	nf, err := netflow.New(
		netflow.WithCaptureTimeout(5 * time.Second),
	)
	if err != nil {
		panic(err)
	}

	err = nf.Start()
	if err != nil {
		panic(err)
	}
	defer nf.Stop()

	<-nf.Done()

	var (
		limit     = 5
		recentSec = 5
	)

	rank, err := nf.GetProcessRank(limit, recentSec)
	if err != nil {
		panic(err)
	}

	_, err = nf.GetProcessesByName(targetName)
	if err != nil {
		panic(err)
	}

	bs, err := json.MarshalIndent(rank, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
}
