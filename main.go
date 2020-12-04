package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shibuya365/changeTextToMarkdown/conf"
)

func main() {
	// サブコマンドtableCmdを宣言
	tableCmd := flag.NewFlagSet("table", flag.ExitOnError)

	// どのサブコマンドが呼ばれたか確認
	switch os.Args[1] {

	// tableコマンド
	case "table":
		tableCmd.Parse(os.Args[2:])
		fmt.Println("subcommand: table")
		fmt.Println("      file:", tableCmd.Arg(0))

		// ファイル名設定
		filename := tableCmd.Arg(0)

		// ファイル読み込み
		table, max, err := conf.ReadTable(filename)
		if err != nil {
			fmt.Println(err.Error())
		}

		// カラムの数を数える
		cols := len(table[0])
		fmt.Println("    colums:", cols)

		// 書き込み用ファイルの生成
		file, err := os.Create(filename + ".md")
		if err != nil {
			fmt.Println(err.Error())
		}
		defer file.Close()

		rows := 1
		// 1行ごとに処理
		for _, words := range table {
			line := ""
			// もし2行目なら区切り線を挿入
			if rows == 2 {
				line := ""
				for i := 0; i < cols; i++ {
					line += "|"
					for i := 0; i < max; i++ {
						line += "-"
					}
				}
				line += "|"
				fmt.Println(line)
				file.WriteString(line + "\n")
			}
			// さらに単語に分解して処理
			for _, word := range words {
				// word = "|" + word
				line += "|" + conf.MakeFixedWidthString(word, max)
			}
			line += "|"
			fmt.Println(line)
			file.WriteString(line + "\n")
			rows++
		}
	default:
		fmt.Println("expected 'table' or 'table' subcommands")
		os.Exit(1)
	}
}
