package main

import(
	"fmt"
	"./grammar"
	"./parser"
	"strings"
	"regexp"
	"time"
)

func main(){
		t := time.Now()
		/*g:= `$ROOT->$S;$VP->$V $NP;$NP->papa;$V->ate;
		$S->$NP $VP;$VP->$VP $PP;$N->caviar;$P->with;
		$NP->$Det $N;$PP->$P $NP;$N->spoon;$Det->the;
		$NP->$NP $PP;$Det->a`*/
				g := "$ROOT->$S;$S->$T + $T + $ROOT;$T-> 1;$T->3;$T->2;$T->4"

		var gram grammar.Grammar
		//chart,_ := parser.Parse(&gram,strings.Split("papa ate the caviar spoon with a spoon"," "))
		//fmt.Println(chart)
		gram.GetGrammarFromString(g)
		/*chart,_:=parser.Parse(&gram,strings.Split("a b a b d c e f d e d a b e f b d e f f f b c d e c d g d c c e h e b b d i f e d c a d f c f i c h a d f f e e c c d b b a a a b d e e h b a d d c d e a e f i c d f c c a e a b a b f h e e b c b e a c b"," "))*/
		
		//_,err := parser.Parse(&gram,strings.Split("a a a a a a a a a a a a a a b d g d d a c a d a a b d g d d a c a d a a b d g d d a c a d a a b d g d d a c a d a a a b d g d d a c a d a a"," "))
		parser.Parse(&gram,strings.Split("2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2 + 1 + 3 + 2"," "))
				//parser.Parse(&gram,strings.Split("papa ate the caviar with a spoon with a spoon"," "))
		fmt.Println(time.Since(t))
}

func getInputFormatted(str string)([]string){
	re := regexp.MustCompile("((\\s)+( \\n| \\t)+)")
	s := re.ReplaceAllString(str," ")
	return strings.Split(s," ")
}