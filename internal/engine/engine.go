package engine

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/alwindoss/ergatis/internal/admin"
)

type Config struct {
	Home         string        `env:"HOME"`
	Port         int           `env:"PORT" envDefault:"3000"`
	Password     string        `env:"PASSWORD,unset"`
	IsProduction bool          `env:"PRODUCTION"`
	Hosts        []string      `env:"HOSTS" envSeparator:":"`
	Duration     time.Duration `env:"DURATION"`
	TempFolder   string        `env:"TEMP_FOLDER" envDefault:"${HOME}/tmp" envExpand:"true"`
	GitLabToken  string        `env:"ERGATIS_TOKEN"`
	BaseURL      string
}

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
