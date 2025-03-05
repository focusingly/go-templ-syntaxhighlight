package utility

import (
	"go/ast"
	"go/parser"
	"go/token"
	"sql-inline-lsp/model"
	"strings"
)

type RawSQLVarsPos struct {
	StartLine    int    // 包含 sql 原始字符串所在的起始行, 默认从 0 开始
	EndLine      int    // 包含 sql 原始字符串所在的结束行
	StartPos     int    // 包含 sql 原始字符串所在的起始字符位置
	EndPos       int    // 包含 sql 原始字符串所在的结束字符位置
	SQLRawString string // sql 字符串
}

func FindSQLPositions(filePath string) ([]*RawSQLVarsPos, error) {
	// 创建文件集
	fset := token.NewFileSet()

	// 解析文件
	file, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	// 存储 SQL 位置信息的切片
	var sqlPositions []*RawSQLVarsPos

	// 遍历 AST
	ast.Inspect(file, func(n ast.Node) bool {
		// 检查字符串字面量
		switch x := n.(type) {
		case *ast.BasicLit:
			// 检查是否是字符串字面量并包含 /* sql */ 注释
			if x.Kind == token.STRING {
				// 获取字符串的原始内容
				rawStr := x.Value

				// 检查前一个注释是否包含 /* sql */
				pos := fset.Position(x.Pos())
				comments := file.Comments
				for _, commentGroup := range comments {
					for _, comment := range commentGroup.List {
						commentPos := fset.Position(comment.Pos())
						// 注释必须在字符串之前，且与字符串在同一行或前一行
						if (commentPos.Line == pos.Line || commentPos.Line == pos.Line-1) &&
							strings.Contains(comment.Text, "/* sql */") {
							// 移除引号并处理原始字符串
							cleanStr := strings.TrimSpace(strings.Trim(rawStr, "`\"'"))
							sqlPositions = append(sqlPositions, &RawSQLVarsPos{
								StartLine:    pos.Line - 1,
								EndLine:      fset.Position(x.End()).Line - 1,
								StartPos:     pos.Column,
								EndPos:       fset.Position(x.End()).Column,
								SQLRawString: cleanStr,
							})
							break
						}
					}
				}
			}
		}
		return true
	})

	return sqlPositions, nil
}

func PosInScopes(rawSQLPosList []*RawSQLVarsPos, completionReq *model.ClientCompletionReq) bool {
	pos := completionReq.Position
	for _, tmpP := range rawSQLPosList {
		switch {
		case tmpP.StartLine == tmpP.EndLine:
			if (pos.Character >= tmpP.StartPos) && (pos.Character <= tmpP.EndPos) {
				return true
			}
		case (pos.Line == tmpP.StartLine) && (pos.Character >= tmpP.StartPos):
			return true
		case (pos.Line == tmpP.EndLine) && (pos.Character <= tmpP.EndPos):
			return true
		case (pos.Line > tmpP.StartLine) && (pos.Line < tmpP.EndLine):
			return true
		}
	}

	return false
}
