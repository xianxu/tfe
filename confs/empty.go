package confs

import (
	"log"
	"github.com/xianxu/tfe"
)

func init() {
	ok := tfe.AddRules("empty", func() map[string]tfe.Rules { return make(map[string]tfe.Rules) })
	if !ok {
		log.Println("Rule set named empty already exists")
	}
}
