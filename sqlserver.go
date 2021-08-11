package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func updatestatment(upsql string) string {
	mystring := strings.TrimSpace(upsql)
	setpos := strings.Index(strings.ToLower(mystring), "set")
	//fmt.Println(setpos)
	// if setpos == -1 {
	// 	setpos = strings.Index(mystring, "SET")
	// }

	wherepos := strings.Index(strings.ToLower(mystring), "where")
	// if wherepos == -1 {
	// 	wherepos = strings.Index(mystring, "WHERE")
	// }

	mystringwhere := mystring[wherepos:]
	mytablename := mystring[6:setpos]

	selectstr := "select * into dbmc.bak."
	myintotable := strings.Replace(mytablename, ".", "_", -1)
	myintotable = strings.Replace(myintotable, "[", "", -1)
	myintotable = strings.Replace(myintotable, "]", "", -1)
	myintotable = strings.Trim(myintotable, " ")

	selectrel := selectstr + myintotable + "_" + time.Now().Format("20060102150405") + " from " + mytablename + " " + mystringwhere
	if strings.HasSuffix(selectrel, ";") {
		return selectrel
	} else {
		return selectrel + " ;"
	}

}
func deletestatment(delstatment string) string {
	delrec := strings.TrimSpace(delstatment)
	delout := delrec[6:]
	wherepos := strings.Index(strings.ToLower(delout), "where")
	// if wherepos == -1 {
	// 	wherepos = strings.Index(delout, "WHERE")
	// }

	myintotable := delout[6:wherepos]
	myintotable = strings.Replace(myintotable, ".", "_", -1)
	myintotable = strings.Replace(myintotable, "[", "", -1)
	myintotable = strings.Replace(myintotable, "]", "", -1)
	myintotable = strings.Trim(myintotable, " ")
	selectrel := "select * into dbmc.bak." + myintotable + "_" + time.Now().Format("20060102150405") + delout
	if strings.HasSuffix(selectrel, ";") {
		return selectrel
	} else {
		return selectrel + " ;"
	}
}

func fitter(instr string) {
	switch {
	case strings.HasPrefix(strings.ToLower(instr), "update"):

		fmt.Println(updatestatment(instr))

	case strings.HasPrefix(strings.ToLower(instr), "delete"):
		fmt.Println(deletestatment(instr))
	default:
		//fmt.Println("  ---------" + beforestr)
	}
}

func main() {
	var url, beforestr string
	//fmt.Printf("Please input the file path:[d:/sqlupdatebefore.txt] ")
	//fmt.Scanln(&url)
	//file, err := os.Open("/opt/sqlbef.txt")
	if len(strings.TrimSpace(url)) == 0 {
		url = "d:/sqlupdatebefore.txt"
	}
	file, err := os.Open(url)

	if err != nil {
		fmt.Printf("Error: %s \n", err)
		return
	}
	defer file.Close()
	br := bufio.NewReader(file)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			fitter(beforestr)
			break
		}
		astr := strings.TrimSpace(string(a))
		//	fmt.Println(astr)
		if strings.HasPrefix(strings.ToLower(astr), "update") || strings.HasPrefix(strings.ToLower(astr), "delete") {
			//fmt.Println(beforestr)
			fitter(beforestr)
			beforestr = astr
		} else {
			beforestr = beforestr + " " + astr
		}
	}
}
