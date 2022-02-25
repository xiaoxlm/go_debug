package init_package1

func init() {
	println("init1 begin")
	slice := make([]int, 8)
	for i := 0; i < 32*1000*1000; i++ {
		slice = append(slice, i)
	}
}
