package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/kyokomi/github-cli/config"
	"log"
	. "github.com/kyokomi/github-cli/github"

	"github.com/google/go-github/github"
	"fmt"
)

var gitHubAppConfig *config.CliAppConfig

func main() {

	gitHubAppConfig = config.NewCliAppConfig(AppName)

	app := cli.NewApp()
	app.Name = "github-cli"
	app.Version = Version
	app.Usage = ""
	app.Author = "kyokomi"
	app.Email = "kyoko1220adword@gmail.com"
	app.Commands = []cli.Command{
		{
			Name:      "gist-list",
			ShortName: "gists",
			Usage:     "Get a list of my gists.",
			Action:    doGistList,
		},
		{
			Name:      "init-config",
			ShortName: "init",
			Usage:     "initialize to config",
			Action:    doInitConfig,
			Flags: []cli.Flag{
				cli.StringFlag{"token", "", "your access token", ""},
			},
		},
	}
	app.Run(os.Args)
}

func doGistList(c *cli.Context) {
	if err := gitHubAppConfig.ReadAccessTokenJson(); err != nil {
		log.Fatalln("error read accessToken ", err)
	}

	client := NewGitHubClient(gitHubAppConfig.AccessConfig.Token)
//	gitLab, err := newGitLabCli(c.GlobalBool("skip-cert-check"))
//	if err != nil {
//		log.Fatal("error create gitlab ")
//	}

//	projectName, err := git.GetCurrentDirProjectName()
//	if err != nil {
//		log.Fatal("not gitlab projectName ", err)
//	}

	opt := &github.GistListOptions{}
	gists, _, err := client.Gists.ListAll(opt)
	if err != nil {
		log.Fatalln("error read gists ", err)
	}

	for _, gist := range gists {
		fmt.Println(gist.String())
	}

//	projectID, err := gitLab.GetProjectID(projectName)
//	if err != nil {
//		log.Fatal("not gitlab projectID ", err)
//	}
//
//	gitLab.PrintIssue(projectID, c.String("state"))
}

func doInitConfig(c *cli.Context) {

	token := c.String("token")

	config := config.AccessConfig{
		Token:   token,
	}
	if err := gitHubAppConfig.WriteAccessConfig(&config); err != nil {
		log.Fatal("appConfig write error ", err)
	}
}
