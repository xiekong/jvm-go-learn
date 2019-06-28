/**
 * @author XieKong
 * @date   2019/6/28 15:44
 */
package cmd

// java 命令的4中形式
// java [-options] class [args]
// java [-options] -jar jarfile [args]
// javaw [-options] class [args]
// javaw [-options] -jar jarfile [args]

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	XjreOption  string
	class       string
	args        []string
}
