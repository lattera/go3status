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
	"bytes"
	"io"
	"os/exec"
	"strings"
)

type Job struct {
	Name		 string
	Instance	 string
	Commands	[]*exec.Cmd
}

func (job *Job) Run() *OutputMessage {
	var out bytes.Buffer
	var r []*io.PipeReader
	var w []*io.PipeWriter
	nCommands := len(job.Commands)

	/* This shouldn't happen. But if it does, be graceful about it. */
	if nCommands == 0 {
		msg := &OutputMessage{Name: job.Name, Instance: job.Instance}
		return msg
	}

	/* If we're only running a single command, then no need for chaining */
	if nCommands == 1 {
		out, err := job.Commands[0].CombinedOutput()
		if err != nil {
			return &OutputMessage{Name: job.Name, Instance: job.Instance}
		}

		msg := &OutputMessage{Name: job.Name, Instance: job.Instance}
		msg.Message = string(out)
		return msg
	}

	/*
	 * Chain the commands together, waiting for each one to finish.
	 *
	 * TODO: Play nice. Introduce a timeout.
	 */

	r = make([]*io.PipeReader, nCommands)
	w = make([]*io.PipeWriter, nCommands)

	for i := 0; i < nCommands; i++ {
		if i < nCommands-1 {
			r[i], w[i] = io.Pipe()
			job.Commands[i].Stdout = w[i]
			job.Commands[i+1].Stdin = r[i]
		} else {
			job.Commands[i].Stdout = &out
		}
	}

	for i := 0; i < nCommands; i++ {
		job.Commands[i].Start()
	}

	for i := 0; i < nCommands-1; i++ {
		job.Commands[i].Wait()
		w[i].Close()
	}

	job.Commands[nCommands-1].Wait()

	return &OutputMessage{Name: job.Name, Instance: job.Instance, Message: strings.TrimSpace(out.String())}
}
