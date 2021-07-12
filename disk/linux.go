// +build !windows

package disk

func getUsage() (int64, int64, int64, error) {

	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return 0, 0, 0, err
	}
	total = fs.Blocks * uint64(fs.Bsize)
	free = fs.Bfree * uint64(fs.Bsize)
	return free, free, total, nil
}