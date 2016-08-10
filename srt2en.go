package main

import(

	"fmt"
	"os"
	"io"
	"bufio"
	"strconv"
	"strings"
)


func main() {

	arg_num := len(os.Args)
	if arg_num != 2 {
		fmt.Printf("Please Input srt file %d\n",arg_num)
		return
	}
	
	filename := os.Args[1]
	if(!strings.HasSuffix(filename,".srt")){
		fmt.Println("Invalid File format")
		return
	}

	f,err := os.Open(filename)
	if err!= nil {
		panic(err)
	}

	defer f.Close()

	n_filename := strings.TrimSuffix(filename,".srt") + "_en.srt"
	ff,err := os.Create(n_filename)
	if err!= nil{
		panic(err)
	}

	w := bufio.NewWriter(ff)
	defer ff.Close()

	rd := bufio.NewReader(f)
	count := 1
	num := 0
	line_num := 0
	total :=0

	for {
		line,err := rd.ReadString('\n')
		if err!=nil || io.EOF ==err {
			break
		}

		tmp_line := strings.TrimRight(line,"\r\n")
		total++

		line_num,err = strconv.Atoi(tmp_line)
		if err == nil {
			if line_num == count {
				num = 0
				count++
			}else if(line_num < count+5){
				num = 0
				count = line_num+1
			}
			
		}

		if num != 2 {
			w.WriteString(line)
			//fmt.Printf("%s",line)
		}

		num++
		
	}

	w.Flush()
	//ww.Flush()
	fmt.Println("convert success!")
}