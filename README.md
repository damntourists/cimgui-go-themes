# cimgui-go-themes
ImTheme parser for cimgui-go! 


## Updating
```shell
git clone --recursive https://github.com/cimgui/cimgui.git
```

Or if already cloned:
```
git submodule update --init --recursive #(If already cloned)
```


# Usage
```go
package main

import (
	"github.com/damntourists/cimgui-go-themes"
)

var theme *imthemes.Theme

func loop() {
	
	
}


func main() {
	t, err := imthemes.GetThemeByName("Future Dark")
	if err != nil {
		panic(err)
    }
	
        

}
    


```