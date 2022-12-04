package main

import "testing"

func test(t *testing.T, equations []string, expected bool) {
	if equationsPossible(equations) != expected {
		t.Log("Equations:", equations)
		t.Fatalf("%v != %v (expected)", !expected, expected)
	}
}

func TestEquationsPossible(t *testing.T) {
	test(
		t,
		[]string{
			"a==b", "b!=a",
		},
		false,
	)
	test(
		t,
		[]string{
			"a==b", "e==c", "b==c", "a!=e",
		},
		false,
	)
	test(
		t,
		[]string{
			"t!=b", "h!=u", "l!=y", "j==j", "w==s", "p==q", "r!=t", "r==i",
			"e!=y", "v==s", "i!=p", "h!=i", "i==o", "e==e", "j!=h", "y!=s",
			"k==g", "c==f", "n==v", "a==w", "d==w", "f!=e", "v==s", "w!=g",
			"g!=s", "j!=d", "c!=u", "y!=n", "q!=j", "d!=x", "l==m", "q!=b",
			"r!=n", "j!=o", "w!=q", "t!=e", "a!=m", "m!=j", "j!=b", "v!=w",
		},
		false,
	)
}
