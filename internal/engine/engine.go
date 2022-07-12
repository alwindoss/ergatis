package engine

import (
	"errors"
	"fmt"
	"log"

	"github.com/alwindoss/ergatis/internal/admin"
)

func GetGroups(cfg *Config, groupID string) {

	// fmt.Printf("%+v\n", cfg)
	adm := admin.NewGitLabAdministrator(cfg.GitLabToken, cfg.BaseURL)
	grps, err := adm.GetGroups(groupID)
	if errors.Is(err, admin.ErrUnAuthorizedUser) {
		log.Fatalf("fatal error: %v", err)
	}
	for _, g := range grps {
		fmt.Printf("%s\n", g.Name)
	}
	fmt.Println("success...")
}
