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
	"encoding/json"
	"io/ioutil"
	"os/exec"
)

type ConfigurationJob struct {
	Name		string		`json:"name"`
	Instance	string		`json:"instance"`
	Color		string		`json:"color"`
	Align		string		`json:"align"`
	Urgent		bool		`json:"urgent"`
	Commands	[]string	`json:"commands"`
}

type Configuration struct {
	Sleep	int			`json:"sleep"`
	Jobs	[]*ConfigurationJob	`json:"jobs"`
}

func ReadConfiguration(path string) (*Configuration, error) {
	config := Configuration{Sleep: 5}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &config)

	return &config, nil
}

func (config *Configuration) TransformConfiguration() ([]*Job, error) {
	jobs := make([]*Job, len(config.Jobs))

	for i := 0; i < len(config.Jobs); i++ {
		var tj *ConfigurationJob
		tj = config.Jobs[i]
		job := &Job{Name: tj.Name, Instance: tj.Instance}
		job.Commands = make([]*exec.Cmd, len(tj.Commands))
		for j := 0; j < len(tj.Commands); j++ {
			cmd := exec.Command("/bin/sh", "-c", tj.Commands[j])

			job.Commands[j] = cmd
		}

		jobs[i] = job
	}

	return jobs, nil
}
