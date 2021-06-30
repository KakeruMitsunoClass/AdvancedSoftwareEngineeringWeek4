package main

import (
	"fmt"
	"strconv"
)

func checkClear(ban [][]string) bool {
	//横でクリアーかどうかのチェック
	for i := 0; i < 3; i++ {
		if (ban[i][0] != ".") && (ban[i][0] == ban[i][1]) && (ban[i][1] == ban[i][2]) {
			return true
		}
	}
	//縦でクリアーかどうかのチェック
	for i := 0; i < 3; i++ {
		if (ban[0][i] != ".") && (ban[0][i] == ban[1][i]) && (ban[1][i] == ban[2][i]) {
			return true
		}
	}
	// 斜めでクリアー
	if (ban[0][0] != ".") && (ban[0][0] == ban[1][1]) && (ban[1][1] == ban[2][2]) {
		return true
	}
	if (ban[0][2] != ".") && (ban[0][2] == ban[1][1]) && (ban[1][1] == ban[2][0]) {
		return true
	}
	return false
}

func process(ban [][]string, koma []string, user int, inp []int) ([][]string, int) {

	// 2以上が入力されたらエラーで終了
	if inp[0] > 2 || inp[1] > 2 {
		fmt.Printf("Error: wrong input\n")
		// エラーで3を返す
		return ban, 3
	}

	// 入力された場所が空白なら入力する
	if ban[inp[0]][inp[1]] == "." {

		ban[inp[0]][inp[1]] = koma[user]
	} else {

		return ban, 4
	}

	// 入力後の番の表示
	for _, line := range ban {
		fmt.Printf(line[0] + " " + line[1] + " " + line[2] + "\n")
	}

	// 版がクリアーかどうかのチェック
	// trueが返ってきたらクリアー，falseなら続行
	flag := checkClear(ban)

	// もしクリアーなら画面表示して終了
	if flag {
		fmt.Printf("Player %d won\n", user+1)
		// クリアーしたらユーザー番号を返す 0 or 1
		return ban, user
	}

	return ban, 2

}

func ticTacToe(input_array [][]int) ([][]string, int) {

	ban := [][]string{{".", ".", "."}, {".", ".", "."}, {".", ".", "."}}
	koma := []string{"o", "x"}

	var num int
	var user int
	// input_arrayが-1なら標準入力で実行
	if input_array[0][0] == -1 {

		var input string
		var j int
		inp := []int{0, 0}

		// 1入力ずつ処理
		for i := 0; i < 9; i++ {

			user = i % 2
			// 入力を標準入力から取得
			fmt.Printf("Player %d: Input (x,y) ", user+1)
			fmt.Scanf("%s", &input)
			j, _ = strconv.Atoi(input[0:1])
			inp[0] = j
			j, _ = strconv.Atoi(input[2:3])
			inp[1] = j

			ban, num = process(ban, koma, user, inp)
			if num == 4 {
				i -= 1
				fmt.Println("Duplicated！")
			} else if num != 2 {
				break
			}
		}

		// 配列で入力されたら配列を入力とする
	} else {
		for i, inp := range input_array {

			user = i % 2
			ban, num = process(ban, koma, user, inp)
			if num == 4 {
				i -= 1
				fmt.Println("Duplicated！")
			} else if num != 2 {
				break
			}

		}
	}

	// 終わらなかったら2を返す
	return ban, num
}

func main() {
	// テスト用
	//input := [][]int{{0, 0}, {1, 1}, {2, 2}, {2, 0}, {0, 2}, {0, 1}, {1, 2}, {1, 0}, {2, 1}}
	// 運用時
	input := [][]int{{-1}}
	ticTacToe(input)
}
