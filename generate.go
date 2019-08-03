// This program generates an updated boilr template from a
// project that was itself generated from this template.
//
//     go run generate.go \
//       --source=path/to/project \
//       --parent="Parent" \
//       --child="Child" \
//       --project="Project" \
//       --template=my-template \
//       --save

// +build ignore

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

type Params struct {
	Name    string
	Repo    string
	Project string
	Parent  string
	Child   string
}

func main() {
	var err error
	var target string
	var source string
	var tempdir bool
	var app string
	var repo string
	var proj string
	var parent string
	var child string
	var templ string
	var save bool

	flag.StringVar(&app, "app", "my-app", "value of {{name}} parameter")
	flag.StringVar(&repo, "repo", "bradrydzewski/my-app", "value of {{repo}} parameter")
	flag.StringVar(&proj, "project", "", "value of {{project}} parameter")
	flag.StringVar(&parent, "parent", "", "value of {{parent}} parameter")
	flag.StringVar(&child, "child", "", "value of {{child}} parameter")
	flag.StringVar(&source, "source", "", "source directory of project")
	flag.StringVar(&target, "target", "/tmp/boilr-gen", "target directory of project template")
	flag.BoolVar(&tempdir, "target-temp", false, "temporary target directory")
	flag.StringVar(&templ, "template", "", "name of boilr template")
	flag.BoolVar(&save, "save", false, "run boilr save")
	flag.Parse()

	if tempdir {
		target, err = ioutil.TempDir("", "boilr-gen")
		if err != nil {
			log.Fatal(err)
		}
	}
	if source == "" {
		log.Fatalln("missing or empty --source flag")
	}
	if target == "" {
		log.Fatalln("missing or empty --target flag")
	}
	if proj == "" {
		log.Fatalln("missing or empty --project flag")
	}
	if parent == "" {
		log.Fatalln("missing or empty --parent flag")
	}
	if child == "" {
		log.Fatalln("missing or empty --child flag")
	}
	if templ == "" && save {
		log.Fatalln("missing or empty --template flag")
	}

	params := Params{
		Name:    app,
		Repo:    repo,
		Project: proj,
		Parent:  parent,
		Child:   child,
	}

	replacer1 := strings.NewReplacer(
		//
		// Reserved
		//
		"{{",
		"{{`{{`}}",
		"}}",
		"{{`}}`}}",
		//
		// Repo
		//
		strings.Title(params.Repo),
		"{{title repo}}",
		strings.ToLower(params.Repo),
		"{{toLower repo}}",
		//
		// Name
		//
		strings.Title(params.Name),
		"{{title name}}",
		strings.ToLower(params.Name),
		"{{toLower name}}",
		//
		// Project
		//
		strings.Title(params.Project),
		"{{title project}}",
		strings.ToLower(params.Project),
		"{{toLower project}}",
		//
		// Parent
		//
		strings.Title(params.Parent),
		"{{title parent}}",
		strings.ToLower(params.Parent),
		"{{toLower parent}}",
		//
		// Child
		//
		strings.Title(params.Child),
		"{{title child}}",
		strings.ToLower(params.Child),
		"{{toLower child}}",
	)

	replacer2 := strings.NewReplacer(
		"{{{",
		"{{`{`}}{{",
		"}}}",
		"}}{{`}`}}",
	)

	// STEP 1: Walk directory tree and list all files
	files := []string{}
	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if shouldSkip(path) {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})

	// STEP 2: Sort files in descending order
	sort.Sort(sort.Reverse(sort.StringSlice(files)))

	// STEP 3: Loop through files
	for _, abs := range files {
		rel, err := filepath.Rel(source, abs)
		if err != nil {
			log.Fatal(err)
		}
		if rel == "." {
			continue
		}

		data, err := ioutil.ReadFile(abs)
		if err != nil {
			log.Fatal(err)
		}

		dataS := string(data)
		dataS = replacer1.Replace(dataS)
		dataS = replacer2.Replace(dataS)

		dst := filepath.Join(target, "template", replacer1.Replace(rel))
		dir := filepath.Dir(dst)
		os.MkdirAll(dir, 0700)

		err = ioutil.WriteFile(dst, []byte(dataS), 0644)
		if err != nil {
			log.Fatal(err)
		}

		println(dst)
	}

	err = ioutil.WriteFile(filepath.Join(target, "project.json"), []byte(projectfile), 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(filepath.Join(target, ".gitignore"), []byte(gitignore), 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(filepath.Join(target, "LICENSE.md"), []byte(licensefile), 0644)
	if err != nil {
		log.Fatal(err)
	}

	println(target)

	if !save {
		return
	}

	cmd := exec.Command("boilr", "template", "save", target, templ, "-f")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

// helper function returns true if the file
// should be skipped.
func shouldSkip(file string) bool {
	for _, skip := range skiplist {
		if strings.Contains(file, skip) {
			return true
		}
	}
	return false
}

// list of files to skip
var skiplist = []string{
	"/.git",
	"/.cache",
	"/dist/files",
	"/node_modules",
	"/release",
	"/.env",
	"/.env.development.local",
	"/database.sqlite",
	"/database.sqlite3",
	"/my-app",
}

// template project gitignore
var gitignore = `
NOTES*
.env
*.sqlite
*.sqlite3
web/node_modules
web/dist/files
release
`

// template project definition file
var projectfile = `{
    "name": "my-app",
    "repo": "bradrydzewski/my-app",
    "project": "Project",
    "parent": "Parent",
    "child": "Child"
}
`

// template project license file
var licensefile = `{
All Rights Reserved.

Copyright (c) 2019 Brad Rydzewski

***As far as the law allows, this software comes as is,
without any warranty or condition, and no contributor
will be liable to anyone for any damages related to this
software or this license, under any kind of legal claim.***

}`
