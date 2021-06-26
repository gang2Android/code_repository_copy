package main

import (
	"code_repository_copy/config"
	"code_repository_copy/utils"
	"os"
)

func main() {
	Run()
}

func Run() {
	var basePath = "." + string(os.PathSeparator) + "sites"

	webs := make([]config.Config, 0)
	config.LoadConfig(&webs)

	for _, v := range webs {
		sourceRepository := v.SourceRepository
		targetRepository := v.TargetRepository

		utils.GitClone(basePath+string(os.PathSeparator), sourceRepository)
		utils.GitClone(basePath+string(os.PathSeparator), targetRepository)

		sourcePath := basePath + string(os.PathSeparator) + utils.GetFileName(sourceRepository)
		targetPath := basePath + string(os.PathSeparator) + utils.GetFileName(targetRepository)
		utils.CopyDir(sourcePath, targetPath)

		utils.Cmd("cd " + basePath + string(os.PathSeparator) + utils.GetFileName(targetRepository) + " && git add . && git commit -m 'init' && git push")
	}
}
