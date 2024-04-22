package version

import (
	"fmt"

	"github.com/KevinYouu/fastGit/functions/colors"
)

var Version string

func GetVersion() {
	GetLogo()
	fmt.Println("Github:", colors.RenderColor("blue", "https://github.com/KevinYouu/fastGit"))
	fmt.Println("Written in Go by", colors.RenderColor("blue", "KevinYouu"))
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
	fmt.Println(fastGit3D)
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
