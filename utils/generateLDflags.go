package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/BurntSushi/toml"
)

type buidConfig struct {
	Configs map[string]value
}

type value struct {
	Exec    *string
	Value   *string
	Env     *string
	IsDebug *string
}

func generateLD() {
	dat, err := ioutil.ReadFile("./build.toml")
	if err != nil {
		panic(err)
	}
	var conf buidConfig
	if _, err := toml.Decode(string(dat), &conf); err != nil {
		panic(err)
	}
	flags := []string{}
	for key, values := range conf.Configs {
		var value string
		if values.Exec != nil {
			cmd := exec.Command("/bin/sh", "-c", *values.Exec)
			cmd.Stderr = os.Stderr
			out, err := cmd.Output()
			if err != nil {
				panic(err)
			}
			value = strings.TrimSpace(string(out))
		} else if values.Env != nil {
			val, ok := os.LookupEnv(*values.Env)
			if !ok {
				panic(fmt.Sprintf("%s is not in the environment", *values.Env))
			}
			value = val
		} else {
			value = *values.Value
		}
		flags = append(flags, fmt.Sprintf("-X %s=%q", key, value))
	}
	fmt.Println(strings.Join(flags, " "))
}
