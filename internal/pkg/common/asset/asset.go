package asset

import (
	"fmt"
	"os"

	"github.com/jonylim/basego/internal/pkg/common/constant/envvar"
	"github.com/jonylim/basego/internal/pkg/common/logger"
)

var cfgBaseStaticAssetURL string

// Init initializes the configurations for application assets.
func Init() {
	cfgBaseStaticAssetURL = os.Getenv(envvar.Asset.BaseStaticAssetURL)
	if cfgBaseStaticAssetURL == "" {
		logger.Println("asset", "WARN: BaseStaticAssetURL is empty")
	} else {
		logger.Println("asset", fmt.Sprintf("BaseStaticAssetURL set to '%s'", cfgBaseStaticAssetURL))
	}
}
