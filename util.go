package cli

func fillCmdCrumbs(cmdCrumbs []string) (cmdCrumbsSlice []string) {
	for k, v := range cmdCrumbs {
		cmdCrumbsSlice = append(cmdCrumbsSlice, v)
		opt := "[global options]"
		if k != 0 {
			opt = "[--options ...]"
		}
		cmdCrumbsSlice = append(cmdCrumbsSlice, opt)
	}
	return cmdCrumbsSlice
}

func stringIndex(in string, sl []string) int {
	for k, v := range sl {
		if v == in {
			return k
		}
	}
	return -1
}