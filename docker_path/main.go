/**
 * @author zhangyuehao
 * @date 2019-03-06 13:51
 */

package docker_path

import (
	"fmt"
	"os"
)

func main() {
	configPath := os.Getenv("Config_Path")
	fmt.Println(fmt.Sprintf("config path is %s", configPath))
}
