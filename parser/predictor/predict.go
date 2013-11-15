package predictor

import(
	"../../chart/state"
	"../../chart"
	"../../grammar"
	"sync"
)

func Predict(c *chart.Chart,s state.State,pos int,g *grammar.Grammar, l *sync.RWMutex, ch chan bool,no *int,noProtect *sync.RWMutex){
	l.RLock()
	ch <- true
	l.RUnlock()
	noProtect.Lock()
	(*no)++
	noProtect.Unlock()
	
	collection :=  g.GetRulesForNonTerminal(s.NextVariable()).RuleCollectionVal
	for _,val:= range collection{
		c.AddStateToColumn(state.State{pos,0,s.NextVariable(),val},pos)
	}
	
	noProtect.Lock()
	(*no)--
	noProtect.Unlock()
}