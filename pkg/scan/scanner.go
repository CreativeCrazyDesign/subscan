package scanner

import (
	"errors"
	"os"

	"github.com/CreativeCrazyDesign/subscan/internal/banner"
	"github.com/CreativeCrazyDesign/subscan/internal/file"
	"github.com/CreativeCrazyDesign/subscan/internal/net"
)

type scanner struct {
	Domain      string
	SublistPath string
	Thread      int
	OutputFile  string
	Sublist     *os.File
	Timeout     int
	Protocol    string
}

func NewScanner(domain string, sublist string, thread int, output string, timeout int, protocol string) (*scanner, error) {
	f, err := os.Open(sublist)
	if err != nil {
		return nil, err
	}

	return &scanner{
		Domain:      domain,
		SublistPath: sublist,
		Thread:      thread,
		OutputFile:  output,
		Sublist:     f,
		Timeout:     timeout,
		Protocol:    protocol,
	}, nil
}

func (s *scanner) Scan() error {
	defer s.Sublist.Close()

	banner.ShowBanner(s.Domain, s.Thread)

	fe, err := file.IsFileExists(s.SublistPath)
	if err != nil {
		return err
	}
	if !fe {
		return errors.New("Sublist " + s.SublistPath + " not exists")
	}

	sbd, err := net.NewSubdomains(s.Domain, s.Timeout, s.OutputFile, s.Protocol)
	if err != nil {
		return err
	}
	defer sbd.CloseFile()

	scanSub(s, sbd)

	return nil
}
