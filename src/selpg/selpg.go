package selpg

import (
	"bufio"
	"github.com/spf13/pflag"
	"fmt"
	"io"
	"os"
	//"os/exec"
	//"io/ioutil"
	"strconv"
	//"strings"
)

type Selpg struct{
	start_page int
	end_page int
	page_len int
	in_file string
	print_des string
	err_des string
}

var (
	selpg Selpg
	start_page, end_page, page_len int
	in_file, print_des, err_des string
)

func set_flag(){
	pflag.IntVar(&(start_page), "s", -1, "Define start page, defaults to -1")
	pflag.IntVar(&(end_page), "e", -1, "Define end page, defaults to -1")
	pflag.IntVar(&(page_len), "l", 1, "Define page length, defaults to 1")
	pflag.StringVar(&(in_file), "i", "in_file.txt", "Define input file's name, defaults to in_file.txt")
	pflag.StringVar(&(print_des), "d", "", "Define input file's name, defaults to NULL")
	//pflag.StringVar(&(err_des), "e", "../errpr_file.txt", "Define input file's name, defaults to errpr_file.txt")
	pflag.Parse()
}

func check_args_1(args []string) {
	for _, para := range args[1:] {
		switch{
		case para[0:2] == "--s", para[0:2] == "--e":
			if val, err := strconv.Atoi(para[4:]); err != nil || val < 0{
				if err != nil{
					fmt.Fprintf(os.Stderr, fmt.Sprintf("%s",err))
				}else{
					fmt.Fprintf(os.Stderr, "\n[Error] The start/end page can not less then zero or be empty\n")
				}
				os.Exit(2)
			}
		case para[0:2] == "--l":
			if val, err := strconv.Atoi(para[4:]); err != nil || val < 0{
				if err != nil{
					fmt.Fprintf(os.Stderr, fmt.Sprintf("%s",err))
				}else{
					fmt.Fprintf(os.Stderr, "\n[Error] The page lenth can not less then zero\n")
				}
				os.Exit(3)
			}
		case para[0:2] == "--i":
			if file, err := os.Open(para[4:]); err != nil {
				fmt.Fprintf(os.Stderr, fmt.Sprintf("%s",err))
				os.Exit(4)
			}else{
				file.Close()
			}
		case para[0] == '<', para[0] == '>':
			if file, err := os.Open(para[1:]); err != nil {
				fmt.Fprintf(os.Stderr, fmt.Sprintf("%s",err))
				os.Exit(4)
			}else{
				if para[0] == '<'{
					in_file = para[1:]
				}else{
					print_des = para[1:]
				}
				file.Close()
			}

		}
	}

}

func check_args_2(){
	if selpg.start_page > selpg.end_page {
		fmt.Fprintf(os.Stderr, "\n[Error] The start page should larger than the end page\n")
		os.Exit(5)
	}
}

func deal_args(){
	selpg.start_page = start_page
	selpg.end_page = end_page
	selpg.page_len = page_len
	selpg.in_file = in_file
	selpg.print_des = print_des
}

func print_args(){
	fmt.Println(selpg.start_page, selpg.end_page, selpg.page_len, selpg.in_file, selpg.print_des)
}

func process(){
	var outfile *os.File
	
	if selpg.print_des != ""{
		if _,err_exit := os.Stat(selpg.print_des); err_exit != nil {
			outfile,_= os.Create(selpg.print_des)
		}else{
			outfile, _ = os.OpenFile(selpg.print_des,os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		}
	}
	defer outfile.Close()
	if selpg.in_file != ""{
		var infile *os.File
		
		infile, _ = os.OpenFile(selpg.in_file, os.O_RDWR, 0666)
		defer infile.Close()
		buf := bufio.NewReader(infile)
		page_count := 1
		line_count := 0
		for page_count <= selpg.end_page {
			line, in_err := buf.ReadString('\n')
			if in_err != nil {
				fmt.Fprintf(os.Stderr, fmt.Sprintf("%s",in_err))
				os.Exit(7)
			}
			if in_err == io.EOF{
				break
			}
			if page_count >= selpg.start_page{
				if selpg.print_des != ""{
					outfile.Write([]byte(line))
				}else{
					fmt.Println(line)
				}
			}
			line_count++
			if line_count == selpg.page_len {
				line_count = 0
				page_count++
			}
		}
	}
}