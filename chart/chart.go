package chart

import(
	"./state"
	"sync"
	"fmt"
)

type StateCollection struct{
	cols []state.State
	rowLock *sync.RWMutex
}

func (sc StateCollection)At(pos int)(returnState state.State){
	sc.rowLock.RLock()
	returnState = sc.cols[pos];
	sc.rowLock.RUnlock()
	return
}

func (sc *StateCollection)AddStateToColumn(s state.State){
	for _,val := range sc.cols{
		if s.Equals(val){
			return
		}
	}
	sc.rowLock.Lock()
	sc.cols = append(sc.cols,s)
	sc.rowLock.Unlock()
}

func (sc *StateCollection)Len()(l int){
	sc.rowLock.RLock()
	l = len(sc.cols)
	sc.rowLock.RUnlock()
	return
}

type Chart struct{
	rows []StateCollection
	chartLock sync.RWMutex
}

func (c *Chart)AddStateToColumn(s state.State, column int){
	if func()(b bool){c.chartLock.RLock();b =(column >= len(c.rows));c.chartLock.RUnlock();return;}(){
		//New Column
		c.chartLock.Lock()
		c.rows = append(c.rows,StateCollection{[]state.State{s},new(sync.RWMutex)})
		c.chartLock.Unlock()
	}else{
		c.rows[column].AddStateToColumn(s)
	}
}

func (c Chart)At(column int)(sc StateCollection){
	c.chartLock.RLock()
	sc = c.rows[column]
	c.chartLock.RUnlock()
	return
}

func (c Chart)LenAt(column int)(l int){
	c.chartLock.RLock()
	l = c.rows[column].Len()
	c.chartLock.RUnlock()
	return
}

func (c Chart)Len()(l int){
	c.chartLock.RLock()
	l = len(c.rows)
	c.chartLock.RUnlock()
	return
}

func (c Chart)PrintC(){
	fmt.Println(c.rows)
}