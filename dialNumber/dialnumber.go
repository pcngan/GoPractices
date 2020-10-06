package dialnumber

import (
	"strings"
	"sort"
	)

type contact struct{
	Name, Number string
}

type contacts []contact

func (cs contacts) Len() int {
    return len(cs)
}
func (cs contacts) Swap(i, j int) {
    cs[i], cs[j] = cs[j], cs[i]
}
func (cs contacts) Less(i, j int) bool {
    return cs[i].Name < cs[j].Name
}

func Solution(A, B []string, P string) string{
	cs := make([]contact, len(A))
	for i,v := range A{
		cs[i] = contact{Name:v, Number:B[i]}
	}
	
	sort.Sort(contacts(cs))
	
	for i:=0; i< len(cs); i++{
		if(strings.Contains(cs[i].Number, P)){
			return cs[i].Name
		}
	}	
	return "NO CONTACT"
}