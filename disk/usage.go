package disk

func DiskUsage() (int64, int64, int64, error) {
	return getUsage()
}
