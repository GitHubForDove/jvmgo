package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

/**
	解决命令中  使用通配符(*)指定某个目录下的所有JAR文件，格式如下：
	java -cp classes;lib\* ...
 */
func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] // remove *
	compositeEntry := []Entry{}

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		/**
			返回SkipDir跳过子目录（通配符类路径不能递归匹配目录下的JAR文件）
		 */
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}

		return nil
	}

	filepath.Walk(baseDir, walkFn)

	return compositeEntry
}
