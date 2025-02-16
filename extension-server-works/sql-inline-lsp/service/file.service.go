package service

import "net/url"

type fileService struct {
}

var DefaultFileService = &fileService{}

func (*fileService) ParseFileProtocolURL(fileURL string) (parsed string, ok bool) {
	// 解析 URL
	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		ok = false
		return
	}

	// 解码文件路径部分
	decodedPath, err := url.PathUnescape(parsedURL.Path)
	if err != nil {
		ok = false
		return
	}

	// 如果路径的开头是 '/'（通常是 Windows 路径中的 `/C:`），需要调整为 `C:\`
	if len(decodedPath) > 1 && decodedPath[0] == '/' {
		// 去掉前导 `/` 并提取盘符
		decodedPath = decodedPath[1:]
		// 提取盘符（如 C:）
		drive := string(decodedPath[0:2])
		// 将路径转换为 Windows 格式
		decodedPath = drive + "\\" + decodedPath[2:]
	}

	ok = true
	parsed = decodedPath

	return
}
