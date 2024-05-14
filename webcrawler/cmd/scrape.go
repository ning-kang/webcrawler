package cmd

import (
	"github.com/ning-kang/webcrawler/webcrawler/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(scrapeCmd)
}

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape books from BookToScrape",
	Long:  "Given website URL, scrape the books from BookToScrape",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		internal.ScrapeBooks(args[0])
	},
}
