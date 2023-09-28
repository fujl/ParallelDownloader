package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

const START_INDEX = 5

func operFirstRowsLeft(firstRows [][]string, rows []string) ([][]string, error) {
	j := 0
	for i, row := range firstRows {
		if i < START_INDEX {
			continue
		}
		if i == len(firstRows)-1 {
			j = i
			break
		}
		pn := row[1]
		if len(pn) == 0 {
			no, err := strconv.Atoi(firstRows[i-1][0])
			if err != nil {
				return nil, err
			}
			firstRows[i][0] = strconv.FormatInt(int64(no+1), 10)
			firstRows[i][1] = rows[1]
			firstRows[i][2] = rows[2]
			firstRows[i][3] = rows[3]
			firstRows[i][4] = rows[4]
			return firstRows, nil
		}
	}

	no, err := strconv.Atoi(firstRows[j-1][0])
	if err != nil {
		return nil, err
	}
	firstRows = append(firstRows, []string{""})
	rows3 := ""
	if len(rows) > 3 {
		rows3 = rows[3]
	}
	rows4 := ""
	if len(rows) > 4 {
		rows4 = rows[4]
	}
	var strs = []string{strconv.FormatInt(int64(no+1), 10), rows[1], rows[2], rows3, rows4}
	copy(firstRows[j+1:], firstRows[j:])
	firstRows[j] = strs
	return firstRows, nil
}

func findProjectName(rows [][]string, projectName string) (int, error) {
	for i, row := range rows {
		if i < START_INDEX {
			continue
		}
		pn := row[1]
		if pn == projectName {
			return i, nil
		}
	}
	return -1, errors.New("not exist")
}

func parseFloatFromString(value string) float64 {
	if len(value) <= 0 {
		return 0
	}
	valueFloat, err := strconv.ParseFloat(strings.TrimSpace(value), 10)
	if err != nil {
		log.Fatal("failed to parse %s %v", value, err)
		return 0
	}
	return valueFloat
}

func operRow(file *excelize.File, firstRows [][]string, rows [][]string) ([][]string, error) {
	for i, row := range rows {
		if i < START_INDEX {
			continue
		}
		if i == len(rows)-1 {
			continue
		}
		if len(row) < 2 {
			continue
		}
		projectName := row[1]
		if len(projectName) <= 0 {
			continue
		}
		var works float64
		if len(row) > 3 {
			works = parseFloatFromString(row[3])
		} else {
			works = 0
		}
		var total float64
		if len(row) > 4 {
			total = parseFloatFromString(row[4])
		} else {
			total = 0
		}
		pos, err := findProjectName(firstRows, projectName)
		if err != nil {
			// no exist, insert
			log.Printf("%v", err)
			firstRows, err = operFirstRowsLeft(firstRows, row)
			if err != nil {
				return nil, err
			}
			continue
		}
		firstWorks := parseFloatFromString(firstRows[pos][3])
		firstTotal := parseFloatFromString(firstRows[pos][4])
		firstRows[pos][3] = fmt.Sprintf("%.2f", works+firstWorks)
		firstRows[pos][4] = fmt.Sprintf("%.2f", total+firstTotal)
	}
	return firstRows, nil
}

func openExcel(excelPath string, fileName string) error {
	file, err := excelize.OpenFile(fileName)
	if err != nil {
		log.Fatal("failed to file %s %v", fileName, err)
		return err
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal("failed to close file %s %v", fileName, err)
		}
	}()

	rows, err := file.GetRows("2")
	if err != nil {
		log.Fatal("failed to get rows 1 %s %v", fileName, err)
		return err
	}
	log.Printf("%v", rows)
	// 遍历后续sheet
	for i := 3; i <= 74; i++ {
		sheet := strconv.Itoa(i)
		newRows, err := file.GetRows(sheet)
		if err != nil {
			log.Fatal("failed to get rows %s %s %v", sheet, fileName, err)
			return err
		}
		log.Printf("%v", newRows)
		rows, err = operRow(file, rows, newRows)
		if err != nil {
			log.Fatal("failed to oper rows %s %s %v", sheet, fileName, err)
			return err
		}
	}
	err = saveRows(rows)
	if err != nil {
		log.Fatal("failed to save rows %v", err)
	}
	return nil
}

func convertToTitle(col int, row int) string {
	const CNT_LETTER = 26
	var result string
	var chs []int
	idx := col
	if idx == 0 {
		return "A" + strconv.Itoa(row+1)
	}
	for idx > 0 {
		tail := idx % CNT_LETTER
		if tail == 0 {
			chs = append(chs, CNT_LETTER)
			idx = (idx - CNT_LETTER) / CNT_LETTER
		} else {
			chs = append(chs, tail)
			idx = idx / CNT_LETTER
		}
	}
	for _, v := range chs {
		result = string(v+65-1) + result
	}
	return result + strconv.Itoa(row+1)
}

func saveRows(rows [][]string) error {
	f := excelize.NewFile()
	if f == nil {
		return errors.New("failed to create excel file")
	}
	sheet, err := f.NewSheet("1")
	if err != nil {
		return err
	}
	for i, row := range rows {
		for j, col := range row {
			f.SetCellValue("1", convertToTitle(j, i), col)
		}
	}
	f.SetActiveSheet(sheet)
	if err := f.SaveAs("C:\\Users\\fujialin\\Desktop\\src\\excel\\20.xlsx"); err != nil {
		return err
	}
	return nil
}

func testInsert() {
	a := []int{1, 2, 3, 5}
	a = append(a, 0)
	i := 3
	copy(a[i+1:], a[i:])
	a[i] = 4
	fmt.Println(a)
}

func main() {
	testInsert()
	log.Printf("hello excel")
	excelPath := "C:\\Users\\fujialin\\Desktop\\src\\excel\\"
	fileName := excelPath + "18.xlsx"
	openExcel(excelPath, fileName)
}
