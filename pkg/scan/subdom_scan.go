package scanner

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/CreativeCrazyDesign/subscan/pkg/alert"
)

var (
	wg = sync.WaitGroup{}
	mu = sync.Mutex{}
)

func scansubdomains(s *scanner, subdomomain *net.subdomains) {
	reader := bufio.NewReader(s.Wordlist)
	var (
		end     = false
		count_t = 0
	)

	for !end {
		subdom, err := reader.ReadString('\n')
		if count_t == s.Thread {
			count_t = 0
			time.Sleep(time.Second * 1)
		}

		wg.Add(1)
		go func(subdom string, err error) {
			defer wg.Done()
			if err == io.EOF {
				end = true
				return
			} else if err != nil {
				fmt.Println(err)
				return
			}

			mu.Lock()
			count_t++
			mu.Unlock()

			subdom = strings.ReplaceAll(
				strings.ReplaceAll(subdom, "\n", ""),
				"\r",
				"",
			)
			_, err = subdomomain.Check(subdom)
			if err != nil {
				alert.Warn(err.Error())
			}
		}(subdom, err)
	}
	wg.Wait()
	fmt.Println(subdomomain.Scan)
}
