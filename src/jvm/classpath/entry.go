/**
 * @author XieKong
 * @date   2019/6/28 18:05
 */
package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	// 寻找和加载class文件
	loadClass(className string) ([]byte, Entry, error)
	// toString
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") ||
		strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") ||
		strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
