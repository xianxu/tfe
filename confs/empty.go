package confs

import (
	"github.com/xianxu/tfe"
	"log"
)

func init() {
	ok := tfe.AddRules("empty", func() map[string]tfe.Rules { return make(map[string]tfe.Rules) })
	if !ok {
		log.Println("Rule set named empty already exists")
	}
}
