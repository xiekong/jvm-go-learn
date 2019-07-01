/**
 * @author XieKong
 * @date   2019/7/1 14:37
 */
package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (c *Classpath) LoadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := c.bootClasspath.loadClass(className); err == nil {
		return data, entry, err
	}

	if data, entry, err := c.extClasspath.loadClass(className); err == nil {
		return data, entry, err
	}
	return c.userClasspath.loadClass(className)
}

func (c *Classpath) String() string {
	return c.userClasspath.String()
}

func (c *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLabPath := filepath.Join(jreDir, "lib", "*")
	c.bootClasspath = newWildcardEntry(jreLabPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	c.extClasspath = newWildcardEntry(jreExtPath)
}

func (c *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	c.userClasspath = newEntry(cpOption)
}

// 优先使用用户输入的-Xjre选项作为jre目录， 其次在当前
// 目录下找，然后才是去系统环境变量JAVA_HOME
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("Can not find jre folder")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}
