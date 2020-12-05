package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/shibuya365/changeTextToMarkdown/conf"
)

func main() {
	// サブコマンドtableCmdを宣言
	tableCmd := flag.NewFlagSet("table", flag.ExitOnError)

	// サブコマンドcheckを宣言
	checkCmd := flag.NewFlagSet("check", flag.ExitOnError)

	// サブコマンドcheckedを宣言
	checkedCmd := flag.NewFlagSet("checked", flag.ExitOnError)

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

		// □の場合
	case "check":
		checkCmd.Parse(os.Args[2:])
		fmt.Println("subcommand: check")
		fmt.Println("      file:", checkCmd.Arg(0))

		// ファイル名設定
		filename := checkCmd.Arg(0)

		// 読込ファイルオープン
		fp, err := os.Open(filename)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer fp.Close()

		// 書き込み用ファイルの生成
		nfp, err := os.Create(filename + ".md")
		if err != nil {
			fmt.Println(err.Error())
		}
		defer nfp.Close()

		scanner := bufio.NewScanner(fp)
		for scanner.Scan() {
			// ここで一行ずつ処理
			line := strings.Replace(scanner.Text(), "□", "- [ ] ", 1)
			fmt.Println(line)
			nfp.WriteString(line + "\n")
		}

		if err = scanner.Err(); err != nil {
			fmt.Println(err.Error())
		}

		// ✅の場合
	case "checked":
		checkedCmd.Parse(os.Args[2:])
		fmt.Println("subcommand: checked")
		fmt.Println("      file:", checkedCmd.Arg(0))

		// ファイル名設定
		filename := checkedCmd.Arg(0)

		// 読込ファイルオープン
		fp, err := os.Open(filename)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer fp.Close()

		// 書き込み用ファイルの生成
		nfp, err := os.Create(filename + ".md")
		if err != nil {
			fmt.Println(err.Error())
		}
		defer nfp.Close()

		scanner := bufio.NewScanner(fp)
		for scanner.Scan() {
			// ここで一行ずつ処理
			line := strings.Replace(scanner.Text(), "✅", "- [x] ", 1)
			fmt.Println(line)
			nfp.WriteString(line + "\n")
		}

		if err = scanner.Err(); err != nil {
			fmt.Println(err.Error())
		}
	default:
		fmt.Println("expected 'table' or 'checked' subcommands")
		os.Exit(1)
	}
}
