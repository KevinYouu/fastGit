package version

import (
	"fmt"

	"github.com/KevinYouu/fastGit/functions/colors"
)

var Version = "untracked"

func GetVersion() {
	GetLogo()
	fmt.Println("Version:", colors.RenderColor("blue", Version))
	fmt.Println("Github:", colors.RenderColor("blue", "https://github.com/KevinYouu/fastGit"))
	fmt.Println("Written in Go by", colors.RenderColor("blue", "KevinYouu"))
	fmt.Println("To know more about me, you can visit my blog:", colors.RenderColor("blue", "https://www.kevnu.com/en-US/about"))
}

func GetLogo() {
	fastGit3D := `
     ___           ___           ___           ___           ___                       ___
    /\  \         /\  \         /\  \         /\  \         /\  \          ___        /\  \
   /::\  \       /::\  \       /::\  \        \:\  \       /::\  \        /\  \       \:\  \
  /:/\:\  \     /:/\:\  \     /:/\ \  \        \:\  \     /:/\:\  \       \:\  \       \:\  \
 /::\~\:\  \   /::\~\:\  \   _\:\~\ \  \       /::\  \   /:/  \:\  \      /::\__\      /::\  \
/:/\:\ \:\__\ /:/\:\ \:\__\ /\ \:\ \ \__\     /:/\:\__\ /:/__/_\:\__\  __/:/\/__/     /:/\:\__\
\/__\:\ \/__/ \/__\:\/:/  / \:\ \:\ \/__/    /:/  \/__/ \:\  /\ \/__/ /\/:/  /       /:/  \/__/
     \:\__\        \::/  /   \:\ \:\__\     /:/  /       \:\ \:\__\   \::/__/       /:/  /
      \/__/        /:/  /     \:\/:/  /     \/__/         \:\/:/  /    \:\__\       \/__/
                  /:/  /       \::/  /                     \::/  /      \/__/
                  \/__/         \/__/                       \/__/
`
	fmt.Println(colors.RenderColor("cyan", fastGit3D))
}

func GetPenguin() {
	Penguin := `
        .--.
       |o_o |
       |:_/ |
      //   \\\\
     (|     | )
    /'\\_   _/'\\
    \\___)=(___/	
`
	fmt.Println(Penguin)
}

func GetDivineBeast() {
	Penguin := `
	┏━┓   ┏━┓+ +
	┏┛ ┻━━━┛ ┻┓ + +
	┃         ┃  
	┃   ━     ┃ ++ + + +
	████━████ ┃+
	┃         ┃ +
	┃   ┻     ┃
	┃         ┃ + +
	┗━┓     ┏━┛
	  ┃     ┃           
	  ┃     ┃ + + + +
	  ┃     ┃
	  ┃     ┃ +  Divine beast bless
	  ┃     ┃    Code without bugs  
	  ┃     ┃  +         
	  ┃     ┗━━━┓ + +
	  ┃         ┣┓
	  ┃         ┏┛
	  ┗┓┓┏━━━┳┓┏┛  + + + +
	   ┃┫┫   ┃┫┫
	   ┗┻┛   ┗┻┛ + + + +
`
	fmt.Println(Penguin)
}
