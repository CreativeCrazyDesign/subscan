package cmd

var (
	wordlist   string
	domain     string
	thread     int
	outputFile string
	timeout    int
	protocol   string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&domain,
		"domain",
		"d",
		"",
		"Domain to scan. Example: example.com",
	)
	rootCmd.PersistentFlags().StringVarP(
		&wordlist,
		"wordlist",
		"w",
		"",
		"Wordlist path (delimiter: newline or \\n). Example: sublist/subdom.txt",
	)
	rootCmd.PersistentFlags().StringVarP(
		&outputFile,
		"output",
		"o",
		".subg.log",
		"Ouput file path (delimiter: newline or \\n). Example: subdom.txt",
	)
	rootCmd.PersistentFlags().IntVarP(
		&timeout,
		"timeout",
		"T",
		2,
		"Timeout (in sec)",
	)
	rootCmd.PersistentFlags().IntVarP(
		&thread,
		"thread",
		"t",
		50,
		"Thread [subdom/sec]",
	)
	rootCmd.PersistentFlags().StringVarP(
		&protocol,
		"protocol",
		"p",
		"https",
		"Protocol to be used. Example: https (https://example.com), http",
	)

	rootCmd.MarkPersistentFlagRequired("domain")
	rootCmd.MarkPersistentFlagRequired("wordlist")
}
