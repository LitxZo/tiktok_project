package main

import "tiktok_project/cmd"

func main() {

	defer cmd.Close()

	cmd.Start()

}
