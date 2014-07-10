package xero

import (
	"net/url"
	"sort"
	"strings"
)

type Pair struct {
	K string
	V string
}

type OrderedPairs struct {
	pairs []Pair
}

func (op *OrderedPairs) GetPairs() []Pair {
	sort.Sort(op)
	return op.pairs
}

func (op *OrderedPairs) Add(k, v string) {
	op.pairs = append(op.pairs, Pair{K: k, V: v})
}

func (op *OrderedPairs) Clone() *OrderedPairs {
	newOp := make([]Pair, len(op.pairs), len(op.pairs))
	for i, p := range op.pairs {
		newOp[i] = p
	}
	return &OrderedPairs{
		pairs: newOp,
	}
}

func (op *OrderedPairs) Len() int {
	return len(op.pairs)
}

func (op *OrderedPairs) Swap(i, j int) {
	op.pairs[i], op.pairs[j] = op.pairs[j], op.pairs[i]
}

func (op *OrderedPairs) Less(i, j int) bool {
	return op.pairs[i].K < op.pairs[j].K
}

func percentEscapeLight(in string) string {
	v := url.QueryEscape(in)
	v = strings.Replace(v, "+", "%20", -1)
	return v
}
