/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "配置阅读的书目录",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

var book Book

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "添加书",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		books := []Book{}
		viper.UnmarshalKey("books", &books)
		if book.Name == "" {
			fmt.Println("书名不能为空")
			return
		}
		file, err := os.Open(book.Path)
		fileScanner := bufio.NewScanner(file)
		for fileScanner.Scan() {
			book.RowNum++
		}
		if err != nil {
			fmt.Printf(err.Error())
		}
		book.Proportion = float64(book.Schedule) / float64(book.RowNum) * 100
		books = append(books, book)
		viper.Set("books", books)
		if err := viper.WriteConfig(); err != nil {
			fmt.Printf(err.Error())
		}
	},
}

// removeCmd represents the add command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "移除书",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		oldBooks := []Book{}
		viper.UnmarshalKey("books", &oldBooks)
		books := make([]Book, 0, len(oldBooks))
		for _, b := range books {
			if b.Name != book.Name {
				books = append(books, b)
			}
		}
		viper.Set("books", books)
		if err := viper.WriteConfig(); err != nil {
			fmt.Printf(err.Error())
		}
	},
}

// listCmd represents the add command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "书列表",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		books := []Book{}
		viper.UnmarshalKey("books", &books)
		for i, b := range books {
			fmt.Printf("%d、名称：%s; 进度：%d｜%d｜%.2f%%", i+1, b.Name, b.Schedule, b.RowNum, b.Proportion)
		}
	},
}

func init() {
	addCmd.PersistentFlags().StringVarP(&book.Name, "name", "n", "", "书名")
	addCmd.PersistentFlags().StringVarP(&book.Path, "path", "p", "", "存放路径")
	addCmd.PersistentFlags().StringVarP(&book.Remark, "remark", "r", "", "备注")
	addCmd.PersistentFlags().IntVarP(&book.Schedule, "schedule", "s", 1, "阅读进度")
	configCmd.AddCommand(addCmd)

	removeCmd.PersistentFlags().StringVarP(&book.Name, "name", "n", "", "书名")
	configCmd.AddCommand(removeCmd)

	configCmd.AddCommand(listCmd)
	rootCmd.AddCommand(configCmd)
}
