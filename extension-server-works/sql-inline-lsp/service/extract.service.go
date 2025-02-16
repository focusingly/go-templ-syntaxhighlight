package service

import (
	"bufio"
	"regexp"
	"sql-inline-lsp/model"
	"strings"
)

type extractService struct{}

var DefaultExtractService = &extractService{}

type SqlRawStringPos struct {
	StartLine  int
	EndLine    int
	StartPos   int
	EndPos     int
	SqlContent string
}

func (*extractService) GetAllRawSqlCommentedStringPos(sourceCode string) []*SqlRawStringPos {
	var results []*SqlRawStringPos

	// 匹配 /* sql */ 注释包裹的原始字符串的正则表达式
	re := regexp.MustCompile(`/\* sql \*/\s*` + "`" + `(.*?)` + "`")

	// 逐行扫描源代码
	scanner := bufio.NewScanner(strings.NewReader(sourceCode))
	var lineNum int
	var sqlContentBuilder strings.Builder
	var isInsideSql bool
	var startLine, startPos int

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		// 如果正在解析一个 SQL 字符串，并且找到了结束标记，记录当前块
		if isInsideSql {
			// 查找 SQL 字符串的结束反引号
			if strings.Contains(line, "`") {
				endPos := strings.Index(line, "`") + len("`")
				sqlContentBuilder.WriteString(line[:endPos])
				results = append(results, &SqlRawStringPos{
					StartLine:  startLine,
					EndLine:    lineNum,
					StartPos:   startPos,
					EndPos:     endPos,
					SqlContent: sqlContentBuilder.String(),
				})
				sqlContentBuilder.Reset()
				isInsideSql = false
			} else {
				// 拼接多行 SQL
				sqlContentBuilder.WriteString(line + "\n")
			}
		}

		// 查找开头的 /* sql */ 注释并处理
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			sqlContent := match[1]
			startPos = strings.Index(line, match[0])
			sqlContentBuilder.WriteString(sqlContent)
			startLine = lineNum
			isInsideSql = true
		}
	}

	return results
}

func (*extractService) PositionInRange(req *model.ClientCompletionReq, positions []*SqlRawStringPos) bool {
	return true

	// p := req.Position
	// return slices.ContainsFunc(positions, func(pos *SqlRawStringPos) bool {
	// 	return (p.Line >= pos.StartLine && p.Line <= pos.EndLine) &&
	// 		(p.Character >= pos.StartPos && p.Character <= pos.EndPos)
	// })
}
