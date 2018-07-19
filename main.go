package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gliderlabs/ssh"
)

const passKey = "password"

func main() {
	ssh.Handle(func(sess ssh.Session) {
		v := sess.Context().Value(passKey)
		pass, ok := v.(string)
		if ok {
			fmt.Printf("user password: %s", pass)
		}
		stdout(sess)
		sess.Exit(int(0))
	})

	ssh.ListenAndServe(":2222", nil,
		ssh.PasswordAuth(func(ctx ssh.Context, pass string) bool {
			ctx.SetValue(passKey, pass)
			return pass == "test"
		}),
	)
}

func stdout(sess ssh.Session) error {
	io.Copy(os.Stdout, sess)
	return nil
}
