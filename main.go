package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func updatestatment(upsql string) string {
	mystring := strings.TrimSpace(upsql)
	setpos := strings.Index(mystring, "set")
	wherepos := strings.Index(mystring, "where")

	// fmt.Println(setpos)
	mystringset := mystring[setpos+3 : wherepos]
	mystringwhere := mystring[wherepos:]
	mytablename := mystring[6:setpos]
	//fmt.Println(mystringset)
	//fmt.Println(mystringwhere)
	//fmt.Println(mytablename)
	//fmt.Println("-----------------------------")
	selectstr := "select id ,"

	for i, v := range strings.Split(mystringset, ",") {

		v2 := strings.Split(v, "=")
		//reg := regexp.MustCompile(`([\w]+)`)
		//ret := reg.FindStringSubmatch(v2)
		//fmt.Println(v2[0])
		//fmt.Println(i)
		if i < len(strings.Split(mystringset, ","))-1 {
			//fmt.Println(len(strings.Split(mystringset, ",")))
			selectstr = selectstr + v2[0] + ","
		} else {
			selectstr = selectstr + v2[0]
		}
	}
	//fmt.Println(strings.Split(mystringset, ","))
	selectrel := selectstr + " from " + mytablename + " " + mystringwhere
	if strings.HasSuffix(selectrel, ";") {
		return selectrel
	} else {
		return selectrel + " ;"
	}

}
func deletestatment(delstatment string) string {
	delrec := strings.TrimSpace(delstatment)
	delout := delrec[6:]
	selectrel := "select *   " + delout
	if strings.HasSuffix(selectrel, ";") {
		return selectrel
	} else {
		return selectrel + " ;"
	}
}

func main() {
	file, err := os.Open("d:\\sqlupdatebefore.txt")
	if err != nil {
		fmt.Printf("Error: %s \n", err)
		return
	}
	defer file.Close()
	br := bufio.NewReader(file)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		astr := strings.TrimSpace(strings.ToLower(string(a)))
		//fmt.Println(astr)
		switch {
		case strings.HasPrefix(astr, "update"):

			fmt.Println(updatestatment(astr))

		case strings.HasPrefix(astr, "delete"):
			fmt.Println(deletestatment(astr))
		default:
			//fmt.Println("  ---------" + astr)
		}

	}
}
