// +build linux darwin netbsd openbsd freebsd

package main

func GetCommand() []string {
     return []string{"ifconfig", "-a"}
}