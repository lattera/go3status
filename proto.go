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
	"fmt"
	"os"
)

type InitMessage struct {
	Version		int	`json:"version"`
}

type OutputMessage struct {
	Name		string	`json:"name"`
	Instance	string	`json:"instance"`
	Message		string	`json:"full_text"`
	Urgent		bool	`json:"urgent"`
	Color		string	`json:"color"`
	Align		string	`json:"align"`
	ShortMessage	string
}

func init_protocol() {
	p := InitMessage{
		Version: 1,
	}

	b, _ := json.Marshal(p)

	os.Stdout.Write(b)
	fmt.Println()
}

func (message *OutputMessage) MarshalOutputMessage() ([]byte, error) {
	var m map[string]interface{}
	m = make(map[string]interface{})

	m["name"] = message.Name;
	m["full_text"] = message.Message;

	if len(message.ShortMessage) > 0 {
		m["short_message"] = message.ShortMessage;
	}

	if len(message.Instance) > 0 {
		m["instance"] = message.Instance;
	}

	if message.Urgent {
		m["urgent"] = true;
	}

	if len(message.Color) > 0 {
		m["color"] = message.Color;
	}

	if len(message.Align) > 0 {
		m["align"] = message.Align;
	}

	return json.Marshal(m)
}
