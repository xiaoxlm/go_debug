package init_package2

func init() {
	println("init2 begin")
	slice := make([]int, 8)
	for i := 0; i < 32*1000*1000; i++ {
		slice = append(slice, i)
	}
}
