package main

import (
	"Snowflake/util"
	"fmt"
)

func main() {
	snowflake := util.Snowflake{}
	for i := 0; i < 100; i++ {
		id := snowflake.NextVal()
		fmt.Println(id)
	}
}
