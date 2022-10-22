package src

import (
	"fmt"
	"github.com/ErmaiSoft/GoOpenXml/word"
)

func SaveAsWord(infos []PaperInfo, path string) {
	titleFont := word.Font{Family: "Arial", Size: 18, Bold: true, Color: "000000", Align: "left"} //字体
	normalFont := word.Font{Family: "Times New Roman", Size: 10.5, Bold: false, Space: true, Color: "000000"}
	normalBoldFont := word.Font{Family: "Times New Roman", Size: 10.5, Bold: true, Space: true, Color: "000000"}
	normalLine := word.Line{Height: 1.5, Rule: word.LineRuleExact} //行高、行间距
	subTitleFont := word.Font{Family: "Times New Roman", Size: 16, Bold: true, Color: "000000", Align: "center"}
	subTitleLine := word.Line{Rule: word.LineRuleAuto, Height: 1.5}
	contentFont := word.Font{Family: "Arial", Size: 12, Bold: false, Color: "000000"}
	contentLine := word.Line{Rule: word.LineRuleAuto, FirstLineChars: 2, Height: 1.5} //行高、行间距、首行缩进
	docx := word.CreateDocx()
	var wordContent []word.Paragraph
	for _, info := range infos {
		temp := [11]word.Paragraph{
			{
				F: titleFont,
				L: word.Line{After: 0.8, Rule: word.LineRuleAuto},
				T: []word.Text{
					{T: info.Title, F: &titleFont},
				},
			},
			{
				F: word.Font{Family: "Times New Roman", Size: 12, Bold: true, Color: "000000"},
				L: subTitleLine,
				T: []word.Text{
					{T: info.Author},
				},
			},
			{
				F: normalFont,
				L: normalLine,
				T: []word.Text{
					{T: "Time: ", F: &normalFont},
					{T: info.Time, F: &normalBoldFont},
				},
			},
			{
				F: normalFont,
				L: normalLine,
				T: []word.Text{
					{T: "Journal: ", F: &normalFont},
					{T: info.Journal, F: &normalBoldFont},
				},
			},
			{
				F: normalFont,
				L: normalLine,
				T: []word.Text{
					{T: "DOI: ", F: &normalFont},
					{T: info.DOI, F: &normalBoldFont},
				},
			},

			{
				F: subTitleFont,
				L: word.Line{Before: 0.8, After: 0.5, Rule: word.LineRuleAuto},
				T: []word.Text{
					{T: "Abstract", F: &subTitleFont},
				},
			},
			{
				F: contentFont,
				L: contentLine,
				T: []word.Text{
					{T: info.Content,
						F: &contentFont},
				},
			},
			{
				F: subTitleFont,
				L: word.Line{Before: 0.8, After: 0.5, Rule: word.LineRuleAuto},
				T: []word.Text{
					{T: "", F: &subTitleFont},
				},
			},
			{
				F: contentFont,
				L: contentLine,
				T: []word.Text{
					{T: info.Translate,
						F: &contentFont},
				},
			},
			{
				F: subTitleFont,
				L: word.Line{Before: 0.8, After: 0.5, Rule: word.LineRuleAuto},
				T: []word.Text{
					{T: "", F: &subTitleFont},
				},
			},
			{
				F: subTitleFont,
				L: word.Line{Before: 0.8, After: 0.5, Rule: word.LineRuleAuto},
				T: []word.Text{
					{T: "", F: &subTitleFont},
				},
			},
		}
		//fmt.Printf("%v\n", temp)
		wordContent = append(wordContent, temp[0:11]...)
		//fmt.Printf("%v\n", wordContent)
	}
	docx.AddParagraph(wordContent)

	err := docx.WriteToFile(path)
	if err != nil {
		fmt.Println(err)
	}
}
