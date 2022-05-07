package sync

import "github.com/ZhaoTzuHsien/construction-sync/internal/pkg/utils"

func isSameFile(source, destination string) bool {
	srcSHA, srcErr := utils.FileSha256(source)
	destSHA, destErr := utils.FileSha256(destination)

	return srcErr == nil && destErr == nil && srcSHA == destSHA
}
