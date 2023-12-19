package main

import (
	"connect"
	"fmt"
)

func main() {
	data:=[]string{
		"OOOX",
		"X..X",
		"X..X",
		"XOOO",
	}

	hg:=connect.New(data)
	connect.PrintMapCells(hg)
	if result, err:=connect.ResultOf(data); err == nil {
		fmt.Println(result)
	}
}

// data:=[]string{
// 	".XX..",
// 	"X.X.X",
// 	".X.X.",
// 	".XX..",
// 	"OOOOO",
// }

// data:=[]string{
// 	"XO..",
// 	"OXXX",
// 	"OXO.",
// 	".OX.",
// 	"XXOO",
// }

// data:=[]string{
// 	"XXO.",
// 	"XO..",
// 	"OX..",
// 	"XXXO",
// 	"OOOX",
// }

// data:=[]string{
// ".O..",
// "OXXX",
// "OXO.",
// "XXOX",
// ".OX.",
// }

// data:=[]string{
// 	".O..",
// 	"OXXX",
// 	"OXO.",
// 	"XXOX",
// 	".OX.",
// }

// data:=[]string{
// 	".XX..",
// 	"X.X.X",
// 	".X.X.",
// 	".XX..",
// 	"OOOOO",
// }

// data:=[]string{
// "OOOX",
// "X..X",
// "X..X",
// "XOOO",
// }	

// data:=[]string{
// "XO..",
// "OXXX",
// "OXO.",
// ".OX.",
// "XXOO",
// }	
