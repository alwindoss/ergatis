package admin

import (
	"fmt"
	"log"

	"github.com/xanzy/go-gitlab"
)

type Group struct {
	Name string
}

type Groups []Group

type Administrator interface {
	UpdateConfig(cfgs ...Config) error
	GetGroups(groupID string) (Groups, error)
}

func NewGitLabAdministrator(token, baseURL string) Administrator {
	var client *gitlab.Client
	var err error
	if baseURL == "" {
		client, err = gitlab.NewClient(token)
	} else {
		fmt.Printf("Client with base url %s created\n", baseURL)
		client, err = gitlab.NewClient(token, gitlab.WithBaseURL(baseURL))
	}
	if err != nil {
		return nil
	}
	return &gitlabAdministrator{
		Client: client,
	}
}

type gitlabAdministrator struct {
	Client *gitlab.Client
}

// GetGroups implements Administrator
func (g *gitlabAdministrator) GetGroups(groupID string) (Groups, error) {
	var grps Groups
	gps, resp, err := g.Client.Groups.ListSubGroups(groupID, nil)
	if resp.StatusCode == 401 {
		err = ErrUnAuthorizedUser
		return nil, err
	}
	if err != nil {
		err = fmt.Errorf("error while listing sub groups: %w", err)
		return nil, err
	}
	for _, grp := range gps {
		// fmt.Println("Name is ", grp.Name)
		grp := Group{
			Name: grp.Name,
		}
		grps = append(grps, grp)

	}
	return grps, nil
}

// UpdateConfig implements Administrator
func (*gitlabAdministrator) UpdateConfig(cfgs ...Config) error {
	for _, cfg := range cfgs {
		panic(cfg)
	}
	git, err := gitlab.NewClient("yourtokengoeshere")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	_, _, err = git.Users.ListUsers(&gitlab.ListUsersOptions{})
	return err
}
