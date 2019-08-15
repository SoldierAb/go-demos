package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
	"algorithms/bubblesort"
	"algorithms/qsort"
)

var infile *string = flag.String("i","infile","File contains values for sorting")
var outfile *string = flag.String("o","outfile","File receive sorted values")
var algorithm *string = flag.String("a","qsort","Sort algorithm")

func readValues(infile string)(values []int,err error){
	file,err := os.Open(infile)
	if err != nil {
		fmt.Println("读取文件失败",infile)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)

	for {
		line,isPrefix , err1 :=br.ReadLine()

		if err1 !=nil{
			if err1 != io.EOF{
				err = err1
			}
			break
		}

		if isPrefix{
			fmt.Println("读取长度过长 ，无法识别")
			return
		}

		str:=string(line)

		value , err1 := strconv.Atoi(str)

		if err1 != nil{
			err = err1
			return
		}

		values = append(values,value)

	}

	return
}


func writeValues(values []int,outfile string) error{
	file,err := os.Create(outfile)
	if err !=nil{
		fmt.Println("输出文件创建失败")
		return err
	}

	defer file.Close()

	for _,value := range values{
		str := strconv.Itoa(value)
		file.WriteString(str+"\n")
	}
	return nil
}



func main(){
	flag.Parse();
 if infile !=nil{
 	fmt.Println("infile=",*infile,"outfile=",*outfile,"algorithm=",*algorithm)
 }
 values,err := readValues(*infile)

 if err == nil{
 	t1 := time.Now()
 	switch *algorithm{
	case "qsort":
		qsort.QuickSort(values)
	case "bubblesort":
		bubblesort.BubbleSort(values)
	default:
		fmt.Println("未找到对应算法")
	}

 	t2:=time.Now()

 	fmt.Println("排序时间：",t2.Sub(t1))
 	writeValues(values,*outfile)
 }else{
 	fmt.Println("运行错误",err)
 }
}