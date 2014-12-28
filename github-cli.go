package main

import (
	"os"

	"log"

	"github.com/codegangsta/cli"
	"github.com/kyokomi/github-cli/config"
	. "github.com/kyokomi/github-cli/github"

	"fmt"

	"github.com/google/go-github/github"
	gitconfig "github.com/tcnksm/go-gitconfig"
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

	name, err := gitconfig.Username()
	if err != nil {
		log.Fatalln("error username ", err)
	}

	client := NewGitHubClient(gitHubAppConfig.AccessConfig.Token)

	gistList(name, client, github.GistListOptions{})
}

func gistList(name string, client *github.Client, opt github.GistListOptions) error {
	gists, res, err := client.Gists.List(name, &opt)
	if err != nil {
		return err
	}

	for _, gist := range gists {
		fmt.Println(*gist.HTMLURL, *gist.Description)
	}

	fmt.Println("NextPage = ", res.NextPage)
	if res.NextPage <= res.PrevPage {
		return nil
	}

	opt.ListOptions.Page = res.NextPage
	return gistList(name, client, opt)
}

func doInitConfig(c *cli.Context) {

	token := c.String("token")

	access := config.AccessConfig{
		Token: token,
	}
	if err := gitHubAppConfig.WriteAccessConfig(&access); err != nil {
		log.Fatal("appConfig write error ", err)
	} else {
		fmt.Println("create config successfull ",
			gitHubAppConfig.ConfigDirPath+"/"+gitHubAppConfig.ConfigFileName)
	}
}
