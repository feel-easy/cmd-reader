/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Book struct {
	Name       string  `yaml:"name"`
	Path       string  `yaml:"path"`
	Schedule   int     `yaml:"schedule"`
	RowNum     int     `yaml:"rowNum"`
	Proportion float64 `yaml:"proportion"`
	Remark     string  `yaml:"remark"`
}

var (
	automatic bool
	num       int
	books     []Book
	pages     int
	speed     int
	histories [][]string
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "开始阅读",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		viper.UnmarshalKey("books", &books)
		if num > len(books) || num < 1 {
			fmt.Println("书不存在")
			return
		}
		if err := ui.Init(); err != nil {
			fmt.Printf("failed to initialize termui: %v", err)
		}
		defer ui.Close()

		c := make(chan string, pages)
		go ReadLine(c)
		list := widgets.NewList()
		list.SetRect(0, 0, 150, pages+2)
		draw := func() {
			rows := make([]string, 0, pages)
			if len(histories) > 0 {
				lastRow := histories[len(histories)-1]
				rows = append(rows, lastRow[len(lastRow)-1])
			}
			for i := 1; i < pages; i++ {
				rows = append(rows, <-c)
			}
			list.Rows = rows
			ui.Render(list)
			histories = append(histories, rows)
			b := books[num-1]
			b.Schedule += pages
			b.Proportion = float64(b.Schedule) / float64(b.RowNum) * 100
			books[num-1] = b
			viper.Set("books", books)
		}
		drawPre := func() {
			if len(histories) > 1 {
				list.Rows = histories[len(histories)-2]
				ui.Render(list)
			}
		}
		draw()
		uiEvents := ui.PollEvents()
		ticker := time.NewTicker(time.Duration(speed) * time.Second).C
		for {
			select {
			case e := <-uiEvents:
				switch e.ID {
				case "q", "<C-c>":
					if err := viper.WriteConfig(); err != nil {
						fmt.Printf(err.Error())
					}
					return
				case "n":
					draw()
				case "p":
					drawPre()
				}
			case <-ticker:
				if automatic {
					draw()
				}
			}
		}
	},
}

func ReadLine(rowData chan<- string) {
	book := books[num-1]
	lineNumber := book.Schedule - 1
	file, _ := os.Open(book.Path)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	lineCount := 1
	for fileScanner.Scan() {
		if lineCount >= lineNumber {
			rowData <- fileScanner.Text()
		}
		lineCount++
	}
	return
}

func init() {
	readCmd.PersistentFlags().BoolVarP(&automatic, "automatic", "a", false, "自动读")
	readCmd.PersistentFlags().IntVarP(&num, "num", "n", 1, "要读的书序号")
	readCmd.PersistentFlags().IntVarP(&pages, "pages", "p", 5, "每页展示的行数")
	readCmd.PersistentFlags().IntVarP(&speed, "speed", "s", 5, "自动读的速度")
	rootCmd.AddCommand(readCmd)
}
