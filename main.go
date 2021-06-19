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
	var basePath = ".\\sites"

	webs := make([]config.Config, 0)
	config.LoadConfig(&webs)

	for _, v := range webs {
		sourceRepository := v.SourceRepository
		targetRepository := v.TargetRepository
		sourceDirname := utils.GetFileName(v.SourceRepository)
		targetDirname := utils.GetFileName(v.TargetRepository)

		utils.GitClone(basePath+string(os.PathSeparator)+sourceDirname, sourceRepository)
		utils.GitClone(basePath+string(os.PathSeparator)+targetDirname, targetRepository)

		utils.CopyDir(basePath+string(os.PathSeparator)+sourceDirname+string(os.PathSeparator)+utils.GetFileName(sourceRepository),
			basePath+string(os.PathSeparator)+targetDirname+string(os.PathSeparator)+utils.GetFileName(targetRepository))

		utils.Cmd("cd " + basePath + string(os.PathSeparator) + targetDirname + string(os.PathSeparator) +
			utils.GetFileName(targetRepository) + " && git add . && git commit -m 'init' && git push")
	}
}
