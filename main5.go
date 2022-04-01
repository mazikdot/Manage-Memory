package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	process  []string
	table    []string
	size     []int
	start    []int
	dataSize int
)

func initialized() {
	process = append(process, "free")
	start = append(start, 0)
	size = append(size, 1000)
}

func showProcess() {
	fmt.Printf("------------------------------ \n")
	fmt.Printf("Process\tStart\tStop\tSize \n")
	for i := range process {
		fmt.Printf("%s\t%d\t%d\t%d \n", process[i], start[i], size[i]-1+start[i], size[i])
	}
	fmt.Printf("------------------------------ \n")
	fmt.Printf("\n")
	fmt.Printf("\nCommand > ")
}

func getCommand() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.Trim(data, "\n")
	return data
}
func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
func RemoveSize(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
func RemoveStart(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func command_create(p string, s int) {
	for i := range process {
		if process[i] == "free" { //check size
			if size[i] < s {
				fmt.Printf("\nSize error!!!\n")
				continue
			}
		}
		if i == 0 { // check if process is the first process in system (index = 0)
			if process[i] == "free" {
				process = append(process, " ")
				copy(process[i+1:], process[i:])
				process[i] = p
				size = append(size, 0)
				copy(size[i+1:], size[i:])
				size[i] = s
				size[i+1] = size[i+1] - s
				start = append(start, 0)
				copy(start[i+1:], start[i:])
				start[i] = start[i+1]
				start[i+1] = s

				if size[i+1] == 0 && process[i+1] == "free" {
					process = RemoveIndex(process, i+1)
					size = RemoveSize(size, i+1)
					start = RemoveStart(start, i+1)
					fmt.Printf("\nSize == 0!!!\n")
					fmt.Printf("i ==  %d\n", i)
					fmt.Printf("len ==  %d\n", len(process))
					break
				}
				break
			}
		} else { // or it's another process in system (another = 1, 2, 3, ..., n)
			if process[i] == "free" {
				if size[i] > 0 { // if size of free process is more than 0 then...
					fmt.Printf("test2")
					process = append(process, " ")
					copy(process[i+1:], process[i:])
					process[i] = p
					size = append(size, 0)
					copy(size[i+1:], size[i:])
					size[i] = s
					size[i+1] = size[i+1] - s
					start = append(start, 0)
					copy(start[i+1:], start[i:])
					start[i] = start[i+1]
					start[i+1] = size[i] + start[i]
					if size[i+1] == 0 { // if size of free process after create process is 0 then...
						if i+1 == len(process)-1 { // if free is the last process in system then...
							process = process[:len(process)-1]
							start = start[:len(start)-1]
							size = size[:len(size)-1]
						} else { // if free is not the last process in system then...
							process = append(process[:i+1], process[i+2:]...)
							size = append(size[:i+1], size[i+2:]...)
							start = append(start[:i+1], start[i+2:]...)
						}
					}
					break
				} else {
					fmt.Println("error...")
				}
			}
		}
	}
}

func command_terminate(p string) {
	if len(process) > 0 { //check if process is not empty
		for i := range process {
			if process[i] == p { //check if process name
				if i-1 >= 0 { //check if this process is not the first process in system
					if i < len(process)-1 { //check if this process is not the last process in system
						if process[i+1] == "free" && process[i-1] != "free" {
							fmt.Println("Terminate1")
							size[i+1] = size[i] + size[i+1]
							start[i+1] = start[i]
							process = append(process[:i], process[i+1:]...)
							size = append(size[:i], size[i+1:]...)
							start = append(start[:i], start[i+1:]...)
							break
						} else if process[i-1] == "free" && process[i+1] != "free" {
							fmt.Println("Terminate2")
							fmt.Printf("i ==  %d\n", i)
							fmt.Printf("len ==  %d\n", len(process))
							fmt.Printf("size i -1  ==  %d\n", size[i-1])
							fmt.Printf("process i -1  ==  %s\n", process[i-1])
							size[i-1] = size[i] + size[i-1]
							start[i-1] = start[i-1]
							process = append(process[:i], process[i+1:]...)
							size = append(size[:i], size[i+1:]...)
							start = append(start[:i], start[i+1:]...)
							break
						} else if process[i-1] == "free" && process[i+1] == "free" {
							fmt.Println("Terminate5555")
							fmt.Printf("i ==  %d\n", i)
							fmt.Printf("len ==  %d\n", len(process))
							fmt.Printf("size i -1  ==  %d\n", size[i-1])
							fmt.Printf("process i -1  ==  %s\n", process[i-1])
							if i > 1 {
								fmt.Printf("i ==  %d\n", i)
								fmt.Printf("process[i] ==  %s\n", process[i])
								fmt.Printf("len ==  %d\n", len(process))
								fmt.Printf("size i -1  ==  %d\n", size[i-1])
								fmt.Printf("process i -1  ==  %s\n", process[i-1])
								fmt.Printf("len process[i] ==  %d\n", len(process[i]))
								if process[i-2] != "free" && i > 1 {
									fmt.Println("Terminate4444")
									fmt.Printf("i ==  %d\n", i)
									fmt.Printf("len ==  %d\n", len(process))
									fmt.Printf("size i -1  ==  %d\n", size[i-1])
									fmt.Printf("process i -1  ==  %s\n", process[i-1])
									size[i-1] = size[i] + size[i-1] + size[i+1]
									start[i-1] = start[i-1]
									process = append(process[:i], process[i+2:]...)
									size = append(size[:i], size[i+2:]...)
									start = append(start[:i], start[i+2:]...)
									// process = process[:len(process)-1]
									break
								} else {
									fmt.Println("Terminate666")
									size[i-1] = size[i] + size[i-1] + size[i+1]
									start[i-1] = start[i-1]
									process = append(process[:i], process[i+1:]...)
									size = append(size[:i], size[i+1:]...)
									start = append(start[:i], start[i+1:]...)
									process = process[:len(process)-1]
									break
								}
							} else {
								fmt.Println("Terminate777")
								fmt.Printf("i ==  %d\n", i)
								fmt.Printf("len ==  %d\n", len(process))
								fmt.Printf("size i -1  ==  %d\n", size[i-1])
								fmt.Printf("process i -1  ==  %s\n", process[i-1])
								fmt.Printf("process i +1  ==  %s\n", process[i+1])
								fmt.Printf("len process    ==  %d\n", len(process))
								fmt.Printf("len process -1   ==  %d\n", len(process)-1)
								fmt.Printf("process[i+1]   ==  %s\n", process[i+1])
								fmt.Printf("process len process-1   ==  %s\n", process[len(process)-1])
								fmt.Printf("len process -1   ==  %d\n", len(process)-1)
								size[i-1] = size[i] + size[i-1] + size[i+1]
								start[i-1] = start[i-1]
								process = RemoveIndex(process, i+1)
								size = RemoveSize(size, i+1)
								start = RemoveStart(start, i+1)
								process = RemoveIndex(process, i)
								size = RemoveSize(size, i)
								start = RemoveStart(start, i)
								break
							}

						} else { //check if there is no free process near this process
							fmt.Println("Terminate4")
							process[i] = "free"
							break
						}
					} else { //check if this process is the last process in system
						if process[i-1] == "free" { //check if there is free process before this process
							fmt.Println("Terminate5")
							size[i-1] = size[i] + size[i-1]
							start[i-1] = start[i-1]
							process = append(process[:i], process[i+1:]...)
							size = append(size[:i], size[i+1:]...)
							start = append(start[:i], start[i+1:]...)
							break
						} else { //check if there is no free process near this process\
							fmt.Println("Terminate6")
							process[i] = "free"
						}
					}
				} else { //check if this process is the first process in system
					if len(process) > 1 {
						if process[i+1] == "free" { //check if there is free process after this process
							fmt.Println("Terminate7")
							size[i+1] = size[i] + size[i+1]
							start[i+1] = start[i]
							process = append(process[:i], process[i+1:]...)
							size = append(size[:i], size[i+1:]...)
							start = append(start[:i], start[i+1:]...)

							break
						} else { //check if there is no free process near this process
							fmt.Println("Terminate8")
							process[i] = "free"
						}
					} else {
						if process[i] == "free" { //check if there is free process after this process
							fmt.Println("Terminate9")
							size[i] = size[i] + size[i]
							start[i] = start[i]
							process = append(process[:i], process[i:]...)
							size = append(size[:i], size[i:]...)
							start = append(start[:i], start[i:]...)

							break
						} else { //check if there is no free process near this process
							fmt.Println("Terminate10")
							process[i] = "free"
						}
					}

				}
			}
		}
	}
}

func main() {
	initialized()
	for {
		showProcess()
		command := getCommand()
		commandx := strings.Split(command, " ")
		switch commandx[0] {
		case "exit":
			return
		case "create":
			sz, _ := strconv.Atoi(commandx[2])
			command_create(commandx[1], sz)
			// pr := strings.Split(commandx[1], ":")
			// sz, _ := strconv.Atoi(pr[1])
			// if sz > 1000 {
			// fmt.Printf("\n Size error!! Please create process's size less than 1000.")
			// } else {
			// command_create(pr[0], sz)
			// }
		case "terminate":
			command_terminate(commandx[1])
		default:
			fmt.Printf("\n Sorry, Command error!")
		}
	}
}
