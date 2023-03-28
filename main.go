package main

import "github.com/forget-the-bright/j/cli"

func main() {
	/* dir, _ := os.Getwd()
	fmt.Println(dir)
	cli.List(dir)
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
		cli.Use(os.Args[1]) //"1.17.9"
	} */
	cli.Run()
	/* 	if len(os.Args) > 1 {
	   		cli.Install(os.Args[1])
	   	}
	*/
	/* ghomeDir, _ := os.Getwd()
	fmt.Println(ghomeDir)
	downloadsDir := filepath.Join(ghomeDir, "downloads")
	versionsDir := filepath.Join(ghomeDir, "versions")
	archiver.Run_unzip(downloadsDir+"\\openjdk-19+36_windows-x64_bin.zip", versionsDir+"\\20", true) */
}
