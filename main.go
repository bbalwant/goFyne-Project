package main

import (
	"fyne.io/fyne/v2"
	"fmt"
    "io/ioutil"
    "log"
	"fyne.io/fyne/v2/app"
	"strings"
	"fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(800 , 400));

	root_src:= "C:\\Users\\Dell\\Pictures";
	files, err := ioutil.ReadDir(root_src)
	if err != nil {
        log.Fatal(err)
    }
	var picsArr []string;
    for _, file := range files {
        fmt.Println(file.Name(), file.IsDir())
    
	if file.IsDir() == false{
		extension:= strings.Split(file.Name() , ".")[1];
		if extension == "png" || extension == "jpg"{
			picsArr = append(picsArr,root_src+"\\"+file.Name());

		}
	}
}

// image := canvas.NewImageFromFile(picsArr[0])
tabs := container.NewAppTabs(
	
	container.NewTabItem("image1", canvas.NewImageFromFile(picsArr[0])),
	// container.NewTabItem("image2", canvas.NewImageFromFile(picsArr[1])),
	
)
for i:= 1; i <len(picsArr); i++ {
	tabs.Append(container.NewTabItem("image" , canvas.NewImageFromFile(picsArr[i])))
}


w.SetContent(tabs);
tabs.SetTabLocation(container.TabLocationLeading)
w.ShowAndRun();


	w.ShowAndRun()
}