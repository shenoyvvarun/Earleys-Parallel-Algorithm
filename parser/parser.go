package parser

import(
	"../chart"
	"../grammar"
	"../chart/state"
	"fmt"
	"./predictor"
	"./scanner"
	"./reducer"
	"strings"
	"sync"
)

func Parse(g *grammar.Grammar,words []string)(parseChart chart.Chart,err error){
	//Add initial state
	parseChart.AddStateToColumn(state.State{0,0,"$ROOT",[]string{"$S"}},0)
	var i,word = 0,""

	//handle parseError
	defer func(){
		if r:= recover();r!=nil{
			fmt.Println("ParseError At : "+strings.Join(words[:i-1]," ")+" ^^"+words[i-1]+"^^ ")		
			}
	}();
	
	var stop sync.RWMutex
	var no =0
	var noProtect sync.RWMutex
	for i,word = range words{
	
				for i>= parseChart.Len() &&  i!=0 {
					stop.Lock()
					stop.Unlock()
					noProtect.RLock()
					n := no
					noProtect.RUnlock()
					if n == 0{
						panic("Parse Error")
					}
				}
		for k:=0;k<parseChart.LenAt(i);k++{
				for (i>= parseChart.Len()) && (k>=parseChart.LenAt(i) && i!=0 && k!=0){
					stop.Lock()
					stop.Unlock()
				}
				state := parseChart.At(i).At(k)
				if state.Incomplete(){
					if state.NonTerminal(){
						ch := make(chan bool)
						go predictor.Predict(&parseChart,state,i,g,&stop,ch,&no,&noProtect)
						<-ch
					}else{
						scanner.Scan(&parseChart,state,word,i)					
					}
				}else{
					stop.Lock()
					reducer.Reduce(&parseChart,state,i)
					stop.Unlock()
				}
		}
	}
	return
}