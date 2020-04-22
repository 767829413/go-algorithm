package main

type Demo struct {
	data interface{}
}

func main() {
	strArr := []byte("(()))((()))")
	count := len(strArr)
	cur := make([]byte, count)
	for _, v := range strArr {
		if getHead(cur) == v {

		} else {
		}
	}
}

func getHead(arr []byte) byte {
	return arr[0]
}

func popHead(arr []byte) {
	arr = arr[1:len(arr)]
}

func popTail(arr []byte) {
	arr = arr[0 : len(arr)-1]
}

func inArr(arr []byte, find byte) bool {
	count := len(arr)
	for i := 0; i < count; i++ {
		if arr[i] == find {
			return true
		}
	}
	return false
}
