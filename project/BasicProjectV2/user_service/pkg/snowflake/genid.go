package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"log"
)

func GenId() (id int) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Println("Gen Id error", err)
		return
	}
	id = int(node.Generate())
	return id
}
