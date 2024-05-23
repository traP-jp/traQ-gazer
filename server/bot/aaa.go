package bot

import (
	"os"
	// "github.com/traPtitech/go-traq"
)

var access_token = os.Getenv("BOT_ACCESS_TOKEN")

func aaa() {
	// client := traq.NewAPIClient(traq.NewConfiguration())

	println(access_token)
}
