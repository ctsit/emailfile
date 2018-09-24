// Copyright 2018 University of Florida
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jordan-wright/email"
)

func main() {
	from := os.Getenv("SMTP_FROM")
	srv := os.Getenv("SMTP_HOST")
	const port = 25

	if len(os.Args) < 3 || from == "" || srv == "" {
		fmt.Println("USAGE: emailfile SUBJECT FILE TO [TO...] <MESSAGE")
		fmt.Println("\nEmail a file.")
		fmt.Println("\nEnvironment Variables")
		fmt.Println("\tSMTP_FROM: no-reply email address used to send emails")
		fmt.Println("\tSMTP_HOST: SMTP hostname")
		fmt.Println("\nWebsite: https://github.com/ufl-taeber/emailfile")
		os.Exit(2)
	}

	msg, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	e := email.NewEmail()
	e.From = from
	e.To = os.Args[3:]
	e.Subject = os.Args[1]
	e.Text = msg
	file := os.Args[2]
	e.AttachFile(file)
	if len(e.Attachments) != 1 {
		fmt.Fprintln(os.Stderr, fmt.Errorf("unable to attach file %s", file))
		os.Exit(1)
	}

	fmt.Printf("From    : %s\n", e.From)
	fmt.Printf("To      : %s\n", strings.Join(e.To, ","))
	fmt.Printf("Subject : %s\n", e.Subject)
	fmt.Printf("File    : %s\n", file)
	fmt.Printf("Message :\n%s\n", e.Text)

	e.Send(fmt.Sprintf("%s:%d", srv, port), nil)
}
