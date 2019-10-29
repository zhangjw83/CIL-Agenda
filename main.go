package main

import  "github.com/zhangjw83/CIL-Agenda/cmd"

func init() {
	log.SetFlags(log.Lshortfile)
	cmd.EnsureAgendaDir()
}

func main() {
	cmd.Execute()
}
