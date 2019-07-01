/**
 * @author XieKong
 * @date   2019/6/28 15:44
 */
package cmd

import (
	"flag"
	"fmt"
	"jvm/classpath"
	"os"
	"strings"
)

// java 命令的4中形式
// java [-options] classpath [args]
// java [-options] -jar jarfile [args]
// javaw [-options] classpath [args]
// javaw [-options] -jar jarfile [args]

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	XjreOption  string
	class       string
	args        []string
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] classpath [args...]\n", os.Args[0])
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func StartJVM() {
	cmd := parseCmd()
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.LoadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data:%v\n", classData)

	//if cmd.versionFlag {
	//	fmt.Printf("Version 0.0.1")
	//} else if cmd.helpFlag || cmd.class == "" {
	//	printUsage()
	//} else {
	//	fmt.Printf("classpath:%s classpath:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
	//}
}
