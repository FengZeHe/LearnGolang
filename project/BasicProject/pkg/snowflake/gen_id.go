package snowflake

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

func GenId() (id string) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	id = node.Generate().String()
	return id
}
