//go:build linux && !s390x && !arm && !386
// +build linux,!s390x,!arm,!386

package disk

import (
	"log"
	"os"
)

func DirSize(path string) (uint64, error) {
	var size uint64
	entries, err := os.ReadDir(path)
	if err != nil {
		return size, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			subDirSize, err := DirSize(path + "/" + entry.Name())
			if err != nil {
				log.Printf("failed to calculate size of directory %s: %v\n", entry.Name(), err)
				continue
			}
			size += subDirSize
		} else {
			fileInfo, err := entry.Info()
			if err != nil {
				log.Printf("failed to get info of file %s: %v\n", entry.Name(), err)
				continue
			}
			size += uint64(fileInfo.Size())
		}
	}
	return size, nil
}

// GetInfo returns total and free bytes available in a directory, e.g. `/`.
func GetInfo(path string) (info Info, err error) {
	// s := syscall.Statfs_t{}
	// err = syscall.Statfs(path, &s)
	// if err != nil {
	// 	return Info{}, err
	// }
	// reservedBlocks := s.Bfree - s.Bavail
	// info = Info{
	// 	Total:  uint64(s.Frsize) * (s.Blocks - reservedBlocks),
	// 	Free:   uint64(s.Frsize) * s.Bavail,
	// 	Files:  s.Files,
	// 	Ffree:  s.Ffree,
	// 	FSType: getFSType(s.Type),
	// }

	// // XFS can show wrong values at times error out
	// // in such scenarios.
	// if info.Free > info.Total {
	// 	return info, fmt.Errorf("detected free space (%d) > total disk space (%d), fs corruption at (%s). please run 'fsck'", info.Free, info.Total, path)
	// }
	// info.Used = info.Total - info.Free

	size, err := DirSize(path)

	info.Used = size

	return info, nil
}
