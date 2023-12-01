package lib

type PromptContect struct {
	ErrorMsg string
	Label    string
}

var PackageManagers []string = []string{"npm", "yarn", "pnpm", "bun"}

var PackageManagersMap map[string]string = map[string]string{
	"npm":  "npx create-next-app@latest .",
	"yarn": "yarn create next-app .",
	"pnpm": "pnpm create next-app .",
	"bun":  "bunx create-next-app .",
}