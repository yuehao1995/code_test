/**
 * Created by g7tianyi on 23/9/2018
 */

package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

const OfficialPath = "/opt/eechains/code_test"

func InitViper(configName string) error {

	altPath := os.Getenv("GATEWAY_CFG_PATH")
	if altPath != "" {

		if !dirExists(altPath) {
			return errors.New(fmt.Sprintf("GATEWAY_CFG_PATH %s does not exist", altPath))
		}

		AddConfigPath(altPath)
	} else {
		devPath, err := GetDevConfigDir()
		if err != nil {
			return err
		}
		AddConfigPath(devPath)
		// And finally, the official path
		if dirExists(OfficialPath) {
			AddConfigPath(OfficialPath)
		}
	}

	viper.SetConfigName(configName)

	return nil
}

func dirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func AddConfigPath(p string) {
	viper.AddConfigPath(p)
}

func GetDevConfigDir() (string, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return "", fmt.Errorf("GOPATH not set")
	}

	for _, p := range filepath.SplitList(gopath) {
		devPath := filepath.Join(p, "src/github.com/eechains/code_test/cfg")
		if !dirExists(devPath) {
			continue
		}

		return devPath, nil
	}

	return "", fmt.Errorf("DevConfigDir not found in %s", gopath)
}
