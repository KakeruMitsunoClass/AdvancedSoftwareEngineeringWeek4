package main

import "testing"

func TestTicTacToe01(t *testing.T) {
	expected := [][]string{{"o", ".", "."}, {".", ".", "."}, {".", ".", "."}} //ban[0][0]にoが入っているか確認したい
	input := [][]int{{0, 0}}
	ban, _ := ticTacToe(input)
	if expected[0][0] != ban[0][0] {
		t.Errorf("Test01 is failed!\n")
		t.Errorf("expected:\n")
		for _, line := range expected {
			t.Errorf(line[0] + " " + line[1] + " " + line[2] + "\n")
		}
		t.Errorf("result:\n")
		for _, line := range ban {
			t.Errorf(line[0] + " " + line[1] + " " + line[2] + "\n")
		}
	}
}

func TestTicTacToe02(t *testing.T) {
	expected := [][]string{{"o", ".", "."}, {".", "x", "."}, {".", ".", "."}} //ban[1][1]にxが入っているか確認したい
	input := [][]int{{0, 0}, {1, 1}}
	ban, _ := ticTacToe(input)
	if expected[1][1] != ban[1][1] {
		t.Errorf("Test02 is failed!\n")
		t.Errorf("expected:\n")
		for _, line := range expected {
			t.Errorf(line[0] + " " + line[1] + " " + line[2] + "\n")
		}
		t.Errorf("result:\n")
		for _, line := range ban {
			t.Errorf(line[0] + " " + line[1] + " " + line[2] + "\n")
		}
	}
}

func TestTicTacToe03(t *testing.T) {
	expected := 0 //player1が縦3つで勝ったとき
	input := [][]int{{0, 0}, {2, 2}, {1, 0}, {1, 1}, {2, 0}}
	_, result := ticTacToe(input)
	if expected != result {
		t.Errorf("Test03 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestTicTacToe04(t *testing.T) {
	expected := 1 //player2が横3つで勝ったとき
	input := [][]int{{2, 2}, {0, 0}, {1, 1}, {0, 1}, {2, 0}, {0, 2}}
	_, result := ticTacToe(input)
	if expected != result {
		t.Errorf("Test04 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestTicTacToe05(t *testing.T) {
	expected := 1 //player2が斜め3つで勝ったとき
	input := [][]int{{2, 2}, {2, 0}, {1, 2}, {1, 1}, {0, 0}, {0, 2}}
	_, result := ticTacToe(input)
	if expected != result {
		t.Errorf("Test05 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestTicTacToe06(t *testing.T) {
	expected := 3 //範囲外が入力されたとき
	input := [][]int{{2, 2}, {3, 0}}
	_, result := ticTacToe(input)
	if expected != result {
		t.Errorf("Test06 fail expected: %d, result: %d\n", expected, result)
	}
}

// func TestTicTacToe07(t *testing.T) {
// 	expected := 0 //scanがうまくできているか
// 	input := [][]int{{0,0}, {0,1}, {0,2}, {1,0}, {1,1}, {1,2}, {2,0}, {2,1}, {2,2}}
// 	_, result := ticTacToe(input)
// 	if expected != result {
// 		t.Errorf("Test07 fail expected: %d, result: %d\n", expected, result)
// 	}
// }
