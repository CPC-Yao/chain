package election

import (
	"fmt"
	"sort"
	"testing"

	"bitbucket.org/cpchain/chain/consensus/dpor/rpt"
	"github.com/ethereum/go-ethereum/common"
)

func TestElect(t *testing.T) {
	a := []float64{1, 2, 5, 6, 6, 7, 8, 9}
	target := 4.0
	nearest, pos := findNearest(a, target)
	fmt.Println(nearest, pos)
	if nearest != 5 || pos != 2 {
		t.Error("findNearest error.")
	}

	var addresses [10]common.Address

	for i := 0; i < len(addresses); i++ {
		addresses[i] = common.HexToAddress("0x" + fmt.Sprintf("%040x", i))
	}

	b := rpt.RPTs{
		rpt.RPT{Address: addresses[0], Rpt: 12},
		rpt.RPT{Address: addresses[1], Rpt: 3},
		rpt.RPT{Address: addresses[2], Rpt: 6},
	}

	fmt.Println("before sort:")
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i].Address.Hex(), b[i].Rpt)
	}

	sort.Sort(b)

	fmt.Println("after sort:")
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i].Address.Hex(), b[i].Rpt)
	}

	rptx := [10]float64{100, 95, 80, 70, 60, 53, 42, 30, 10, 5}

	rpts := rpt.RPTs{rpt.RPT{Address: addresses[0], Rpt: rptx[0]}}
	for i := 1; i < len(addresses); i++ {
		rpts = append(rpts, rpt.RPT{Address: addresses[i], Rpt: rptx[i]})
	}

	seed := int64(66)
	viewLength := 5
	signers := Elect(rpts, seed, viewLength)

	fmt.Println("commissioners:")
	for i := 0; i < len(signers); i++ {
		fmt.Println(i, signers[uint64(i)].Hex())
	}

	expectedCommittee := []common.Address{
		addresses[9],
		addresses[4],
		addresses[3],
		addresses[2],
		addresses[1],
	}
	for i := 0; i < viewLength; i++ {
		if signers[uint64(i)] != expectedCommittee[i] {
			t.Error("elect error.")
		}
	}

}
