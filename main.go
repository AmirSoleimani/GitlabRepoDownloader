package main

import (
	"GitlabRepoDownlaoder/gilab"
	"GitlabRepoDownlaoder/utils"
	"flag"
	"fmt"
	"net/url"

	"github.com/sirupsen/logrus"
)

var (
	page = 1
)

const (
	usage = `
		* * * * * * * * * * * * * * * * * * * * * * * * * * *
		* Gitlab Repositories Downloader                    *
		* https://github.com/AmirSoleimani                  *
		* * * * * * * * * * * * * * * * * * * * * * * * * * *
		*
		* Usage>
		*
		* -host : gitlab url
		* -token : gitlab user private token
		* -datadir : data directory for save repositories
		*
		* BLABLA.
		* * * * * * * * * * *
	`
)

func main() {

	// welcome
	fmt.Println(usage)

	// cli parser
	var host, accessToken, datadir string
	flag.StringVar(&host, "host", "", "Gitlab Host Name")
	flag.StringVar(&accessToken, "token", "", "Gitlab Private Token")
	flag.StringVar(&datadir, "datadir", "", "Data Directory")
	flag.Parse()

	// check url
	_, err := url.ParseRequestURI(host)
	if err != nil {
		logrus.Errorln("invalid gitlab URL")
		return
	}

	// logger
	logrus.Infoln("Start Downloading...")

	// start
	for {
		resp, err := gilab.GetRepos(host, accessToken, page)
		if err != nil {
			logrus.Errorln(err)
			break
		}

		if len(resp) == 0 {
			logrus.Infoln("Finished, count:", page-1)
			break
		}

		for i, a := range resp {
			logrus.Infoln(page, i, "Started...")

			logrus.Infoln("-- Namespace:", a.PathWithNamespace)
			logrus.Infoln("-- Get Branches...")

			// get branches
			branches, err := gilab.GetBranches(host, accessToken, a.ID)
			if err != nil {
				logrus.Errorln("-- ERROR:", err)
				continue
			}

			// start download branches
			for _, b := range branches {

				// variables
				u := fmt.Sprintf("%s/api/v4/projects/%d/repository/archive?private_token=%s&sha=%s", host, a.ID, accessToken, b.Commit.ID)
				p := fmt.Sprintf("%s/%s/%s", datadir, a.PathWithNamespace, b.Name)
				sp := p + "/archive.zip"

				// check and create directory
				utils.CreateDirIfNotExist(p)

				// download file and store in datadir
				logrus.Infoln("-- Branch:", b.Name)
				logrus.Infoln("-- URL:", u)
				logrus.Infoln("-- Path:", sp)
				err := utils.DownloadFile(sp, u, accessToken)
				if err != nil {
					logrus.Errorln("-- ERROR:", a.PathWithNamespace, err)
				}
			}

		}

		page++
	}

}
