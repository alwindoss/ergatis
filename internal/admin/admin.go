package admin

import (
	"fmt"
	"log"
	"sync"

	"github.com/xanzy/go-gitlab"
)

type Group struct {
	Name string
}

type Groups []Group

type Member struct {
	Name string
}

type Members []Member

type Administrator interface {
	UpdateConfig(cfgs ...Config) error
	GetGroups(groupID string) (Groups, error)
	GetGroupMembers(groupID string) (Members, error)
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

// GetGroupMembers implements Administrator
func (g *gitlabAdministrator) GetGroupMembers(groupID string) (Members, error) {
	var mems Members
	gmems, _, err := g.Client.Groups.ListAllGroupMembers(groupID, nil)
	if err != nil {
		err = fmt.Errorf("error while listing group members: %w", err)
		return nil, err
	}
	for _, m := range gmems {
		mem := Member{
			Name: m.Username,
		}
		mems = append(mems, mem)
	}
	return mems, nil
}

// GetGroups implements Administrator
func (g *gitlabAdministrator) GetGroups(groupID string) (Groups, error) {
	var grps Groups

	gpCh := make(chan Group, 100)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for g1 := range gpCh {
			grps = append(grps, g1)
		}
		wg.Done()
	}()
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
		grp := Group{
			Name: grp.Name,
		}
		gpCh <- grp
	}
	if resp.TotalPages > 1 {
		for i := 2; i <= resp.TotalPages; i++ {
			gps, _, err := g.Client.Groups.ListSubGroups(groupID, &gitlab.ListSubGroupsOptions{
				ListOptions: gitlab.ListOptions{
					Page: i,
				},
			})
			if resp.StatusCode == 401 {
				err = ErrUnAuthorizedUser
				return nil, err
			}
			if err != nil {
				err = fmt.Errorf("error while listing sub groups: %w", err)
				return nil, err
			}
			for _, grp := range gps {
				grp := Group{
					Name: grp.Name,
				}
				gpCh <- grp
			}
		}
	}
	close(gpCh)

	wg.Wait()
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
