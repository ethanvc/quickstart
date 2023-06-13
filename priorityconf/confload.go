package priorityconf

import (
	"fmt"
	"github.com/miracl/conflate"
	"os"
	"path/filepath"
)

func getFileList(env, cid string) []string {
	fileList := []string{
		"config",
		"config." + env,
		"config." + env + "." + cid,
	}
	return fileList
}

var supportExts = []string{"toml", "yml", "json"}

func LoadConf(dir string, env string, cid string, v any) error {
	conf := conflate.New()
	for _, f := range getFileList(env, cid) {
		for _, ext := range supportExts {
			realF := filepath.Join(dir, fmt.Sprintf("%s.%s", f, ext))
			if stat, _ := os.Stat(realF); stat != nil && !stat.IsDir() {
				err := conf.AddFiles(realF)
				if err != nil {
					return err
				}
			}
		}
	}

	buf, _ := conf.MarshalJSON()
	_ = buf

	err := conf.Unmarshal(v)
	if err != nil {
		return err
	}
	return nil
}
