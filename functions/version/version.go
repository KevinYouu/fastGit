package version

import (
	"fmt"

	"github.com/KevinYouu/fastGit/functions/colors"
	"github.com/KevinYouu/fastGit/functions/random"
)

var Version = "untracked"

func GetVersion() {
	funcProbs := []random.FuncProbability{
		{Function: func() { GetLogo() }, Probability: 0.98},
		{Function: func() { GetPenguin() }, Probability: 0.1},
		{Function: func() { GetDivineBeast() }, Probability: 0.1},
	}
	random.ExecuteRandomly(funcProbs)

	fmt.Println("Version:", colors.RenderColor("blue", Version))
	fmt.Println("Github:", colors.RenderColor("blue", "https://github.com/KevinYouu/fastGit"))
	fmt.Println("To know more about me, you can visit my blog:", colors.RenderColor("blue", "https://www.kevnu.com/about"))
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
	divineBeast := `
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
	fmt.Println(divineBeast)
}
