/*-
 * Copyright (c) 2015 Shawn Webb <lattera@gmail.com>
 * All rights reserved.
 * 
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above
 *    copyright notice, this list of conditions and the following
 *    disclaimer.
 * 2. Redistributions in binary form must reproduce the above
 *    copyright notice, this list of conditions and the following
 *    disclaimer in the documentation and/or other materials
 *    provided with the distribution.
 * 3. The name of the author may not be used to endorse or promote
 *    products derived from this software without specific prior
 *    written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND
 * CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES,
 * INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR
 * CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF
 * USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED
 * AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN
 * ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
*/

package main

import (
	//"encoding/json"
	"fmt"
	"os"
//	"os/exec"
)

func main() {
	init_protocol();

	/*
	cmd := exec.Command("/bin/sh", "-c", "/bin/ls")
	job := Job{Name: "ls"}

	job.Commands = make([]*exec.Cmd, 2)
	job.Commands[0] = cmd

	cmd = exec.Command("/bin/sh", "-c", "head -n 1")
	job.Commands[1] = cmd

	m := job.Run()
	b, _ := m.MarshalOutputMessage()

	os.Stdout.Write(b)
	fmt.Println()
	*/

	c, err := ReadConfiguration("config.json.example")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	jobs, _ := c.TransformConfiguration()

	for _,job := range jobs {
		m := job.Run()
		str, _ := m.MarshalOutputMessage()

		fmt.Printf("%v\n", string(str))
	}
}
