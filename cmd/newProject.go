package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/pflag"
	"os"
	"os/exec"
)

var sign =`
/* This file generated automatically by GoB.
   you can find GoB from: https://github.com/mrNobody95/gob
*/
`

var projectName string

func createNewProject(name string, flags *pflag.FlagSet) {
	if name == "" {
		color.Red("you must set project name")
		color.HiRed("you must set project name")
		return
	}
	projectName=name
	fmt.Println("Creating project directory")
	err:=os.Mkdir(name, 0777)
	if err!=nil{
		color.Red(err.Error())
		return
	}
	fmt.Println("Generating main file")
	f, err:=os.Create(name+"/main.go")
	if err != nil {
		color.Red(err.Error())
		return
	}
	defer f.Close()

	_, err=f.WriteString(fmt.Sprintf(`
package main
%s

func main() {
	//start writing your code from here
}
`, sign))
	if err != nil {
		color.Red(err.Error())
		return
	}
	if env, err:=flags.GetBool("skip-env"); err==nil && !env {
		fmt.Println("Generating .env file")
		envF, err:=os.Create(name+"/.env")
		if err != nil {
			color.Red(err.Error())
			return
		}
		defer envF.Close()
	}
	if module, err:=flags.GetBool("skip-go-module"); err==nil && !module {
		fmt.Println("Initializing go module")
		cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("cd %s; go mod init", projectName))
		err := cmd.Run()
		if err != nil {
			color.Red(err.Error())
			return
		}
	}
	if docker, err:=flags.GetBool("skip-docker"); err==nil && !docker {
		createDockerModule()
	}
	if git, err:=flags.GetBool("skip-git"); err==nil && !git {
		createGitModule()
	}
	color.Green("Project created successfully.")
}

func createDockerModule() {
	fmt.Println("Generating docker file")
	f, err:=os.Create(projectName+"/dockerfile")
	if err != nil {
		color.Red(err.Error())
		return
	}
	defer f.Close()
	_, err=f.WriteString(`
# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]
`)
	if err != nil {
		color.Red(err.Error())
		return
	}

	fi, err:=os.Create(projectName+"/.dockerignore")
	if err != nil {
		color.Red(err.Error())
		return
	}
	defer fi.Close()
	_, err=fi.WriteString(fmt.Sprintf(`
%s
.env
`, sign))
	if err != nil {
		color.Red(err.Error())
	}
}

func createGitModule() {
	fmt.Println("Initializing git")
	goCmd:=exec.Command("git", "init")
	err:=goCmd.Run()
	if err != nil {
		color.Red(err.Error())
		return
	}
	f, err:=os.Create(projectName+"/.gitignore")
	if err != nil {
		color.Red(err.Error())
		return
	}
	defer f.Close()
	_, err=f.WriteString(`
.idea/
.env
bin/
`)
	if err != nil {
		color.Red(err.Error())
		return
	}
}
