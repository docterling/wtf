package github

import (
	"fmt"
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/wtf"
)

const defaultTitle = "GitHub"

type Settings struct {
	common *cfg.Common

	apiKey       string
	baseURL      string
	enableStatus bool
	repositories []string
	uploadURL    string
	username     string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		apiKey:       ymlConfig.UString("apiKey", os.Getenv("WTF_GITHUB_TOKEN")),
		baseURL:      ymlConfig.UString("baseURL", os.Getenv("WTF_GITHUB_BASE_URL")),
		enableStatus: ymlConfig.UBool("enableStatus", false),
		uploadURL:    ymlConfig.UString("uploadURL", os.Getenv("WTF_GITHUB_UPLOAD_URL")),
		username:     ymlConfig.UString("username"),
	}
	settings.repositories = parseRepositories(ymlConfig)

	return &settings
}

func parseRepositories(ymlConfig *config.Config) []string {

	result := []string{}
	repositories, err := ymlConfig.Map("repositories")
	if err == nil {
		for key, value := range repositories {
			result = append(result, fmt.Sprintf("%s/%s", value, key))
		}
		return result
	}

	result = wtf.ToStrs(ymlConfig.UList("repositories"))
	return result
}
