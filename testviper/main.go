/**
 * @author zhangyuehao
 * @date 2019-01-17 13:38
 */

package main

import (
	"fmt"
	"github.com/eechains/code_test/testviper/config"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

const coreConfig = "core"

func InitConfig() error {
	config.InitViper(coreConfig)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		// The version of Viper we use claims the config type isn't supported when in fact the file hasn't been found
		// Display a more helpful message to avoid confusing the user.
		if strings.Contains(fmt.Sprint(err), "Unsupported Config Type") {
			return errors.New(fmt.Sprintf("Could not find config file. "+
				"Please make sure that FABRIC_CFG_PATH is set to a path "+
				"which contains %s.yaml", coreConfig))
		} else {
			return errors.WithMessage(err, fmt.Sprintf("error when reading %s config file", coreConfig))
		}
	}
	return nil
}

func main() {
	InitConfig()
	fmt.Println(viper.GetString("port"))
}
