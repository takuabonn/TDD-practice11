package main

import "testing"
import "strings"
// import "fmt"
func TestHello(t *testing.T) {
    t.Run("saying hello to people", func(t *testing.T) {
        got := createAdjacencyList()
		i := 0
		want := make([]string, 2)
		want[0] = "1 2 3 4"
		want[1] = "5 6 7 8"
		
		for i < len(got) {
			st := strings.Join(got[i], " ")
			if st != want[i] {
				t.Errorf("配列の中身が違います got %q want %q", got[i], want[i])
			}
			i++
		}
		
        
    })
}