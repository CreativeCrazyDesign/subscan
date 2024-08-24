package banner

import (
	"fmt"

	"github.com/CreativeCrazyDesign/subscan/config"
)

func ShowBanner(domain string, thread int) {
	text := fmt.Sprintf(`
#################################################

Tool:       sub      (github.com/CreativeCrazyDesign/subscan)
Written by: ccd (github.com/CreativeCrazyDesign)
Contact:    mail      (creativecrazydesign@gmail.com)

Version:    %s

Domain:     %s
Thread:     %d

#################################################`,
		config.Config["version"], domain, thread,
	)

	fmt.Println(text)
}
