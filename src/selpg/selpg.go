package main

import (
	"bufio"
	"github.com/spf13/pflag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
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
	pflag.IntVar(&(page_len), "l", 72, "Define page length, defaults to 72")
	pflag.StringVar(&(in_file), "i", "in_file.txt", "Define input file's name, defaults to in_file.txt")
	pflag.StringVar(&(print_des), "d", "", "Define input file's name, defaults to NULL")
	//pflag.StringVar(&(err_des), "e", "../errpr_file.txt", "Define input file's name, defaults to errpr_file.txt")
	pflag.Parse()
}

func check_args_1(args []string) {
	for _, para := range args[1:] {
		switch{
		case para[2] == 's', para[2] == 'e':
			if val, err := strconv.Atoi(para[4:]); err != nil || val < 0{
				if err != nil{
					fmt.Fprintf(os.Stderr, fmt.Sprintf("%s",err))
				}else{
					fmt.Fprintf(os.Stderr, "\n[Error] The start/end page can not less then zero or be empty\n")
				}
				os.Exit(2)
			}
		case para[2] == 'l':
			if val, err := strconv.Atoi(para[4:]); err != nil || val < 0{
				if err != nil{
					fmt.Fprintf(os.Stderr, fmt.Sprintf("%s",err))
				}else{
					fmt.Fprintf(os.Stderr, "\n[Error] The page lenth can not less then zero\n")
				}
				os.Exit(3)
			}
		case para[2] == 'i', para[2] == 'd':
			if file, err := os.Open(para[4:]); err != nil {
				fmt.Fprintf(os.Stderr, fmt.Sprintf("%s",err))
				os.Exit(4)
			}else{
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
	var cmd *exec.Cmd
	var cmd_in io.WriteCloser
	var cmd_out io.ReadCloser
	if selpg.print_des != "" {
		cmd = exec.Command("bash", "-c", selpg.print_des)
		cmd_in, _ = cmd.StdinPipe()
		cmd_out, _ = cmd.StdoutPipe()
		cmd.Start()
	}
	page_count := 1
	if selpg.in_file != "" {
		in, err := os.Open(selpg.in_file)
		fin := bufio.NewReader(in);
		for page_count <= selpg.end_page {
			line, _, err := fin.ReadLine()
			if err != io.EOF && err != nil {
				fmt.Fprintf(os.Stderr, fmt.Sprintf("%s",err))
				os.Exit(6)
			}
			if err == io.EOF {
				break
			}
			if page_count >= selpg.start_page && page_count <= selpg.end_page{
				if selpg.print_des == ""{
					fmt.Printf(string(line))
				}else{
					fmt.Fprintln(cmd_in, string(line))
				}
			}
			page_count++
		}
		in.Close()
	}
}
func main(){
	args := os.Args
	check_args_1(args)
	set_flag()
	deal_args()
	check_args_2()
	process()
	print_args()
}