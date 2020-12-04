package conf

import (
	"bufio"
	"os"
	"strings"

	"github.com/mattn/go-runewidth"
)

// DB読み込み関数
func ReadTable(filename string) (Table [][]string, Max int, err error) {

	// ファイルオープン
	fp, err := os.Open(filename)
	if err != nil {
		return Table, Max, err
	}

	// ファイルクローズ
	defer fp.Close()

	// スキャナー取得
	scanner := bufio.NewScanner(fp)

	// 1行ずつ読み込み
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		Table = append(Table, words)

		// １単語の長さを調べる
		for _, word := range words {
			l := 0
			for _, c := range word {
				cl := runewidth.RuneWidth(c)
				l += cl
			}
			if Max < l {
				Max = l
			}

		}
	}

	return Table, Max, nil
}
