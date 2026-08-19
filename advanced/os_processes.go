package main

import (
	"fmt"
	"os/exec"
)

func mainospro() {

	cmd := exec.Command("ls", "-l")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error : ", err)
	}

	fmt.Println("Output : ", string(output))

	//sample 5
	// pr, pw := io.Pipe()
	// cmd := exec.Command("grep", "foo")
	// cmd.Stdin = pr

	// go func() {
	// 	defer pw.Close()
	// 	pw.Write([]byte("food is good\nbar\nbaz\n"))

	// }()

	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("Error : ", err)
	// 	return
	// }

	// fmt.Println("Output : ", string(output))

	//sample 4
	// cmd := exec.Command("printenv", "SHELL")
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("Error : ", err)
	// 	return
	// }
	// fmt.Println("Output : ", string(output))

	//sample 3
	// cmd := exec.Command("sleep", "60")
	// err := cmd.Start()
	// if err != nil {
	// 	fmt.Println("Error : ", err)
	// 	return
	// }

	// //misal kill process. tanpa tunggu sleepcmd 60s
	// time.Sleep(2 * time.Second)
	// cmd.Process.Kill()

	// //Waiting
	// // err = cmd.Wait()
	// // if err != nil {
	// // 	fmt.Println("Error : ", err)
	// // 	return
	// // }

	// fmt.Println("Proccess complete")

	//sample 2
	// cmd := exec.Command("grep", "foo")

	// // Set input for the command
	// cmd.Stdin = strings.NewReader("food is good \nbar\nbaz\n")

	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Printf("%v", err)
	// 	return
	// }
	// fmt.Println("Output:", string(output))

	//sample 1
	// cmd := exec.Command("echo", "Happy go")
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("error : ", err)
	// }

	// fmt.Println("output >>", string(output))

}
