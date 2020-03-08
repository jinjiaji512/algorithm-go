package sorting

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"testing"
	"time"
)

var testcase = []int{}

func init() {
	testcase = generateUnsortedArray()
	//testcase = []int{915, 616}
}
func GetTestCase() (s []int, err error) {
	var bs []byte
	if bs, err = json.Marshal(testcase); err != nil {
		fmt.Println(err)
		return
	}
	if err = json.Unmarshal(bs, &s); err != nil {
		fmt.Println(err)
	}
	return
}

//生成随机数组
func generateUnsortedArray() (s []int) {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < 10; i++ {
		s = append(s, rand.Intn(1000))
	}
	return
}

//记录运行时间
func timelog(t *testing.T) func() {
	start := time.Now()
	return func() {
		t.Logf("cost %d ms", time.Since(start)/time.Millisecond)
	}
}

func testSort(t *testing.T, f func([]int) []int) {
	fname := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	ss := strings.Split(fname, ".")
	fname = ss[len(ss)-1]

	var (
		s   []int
		err error
	)

	if s, err = GetTestCase(); err != nil {
		t.Fatalf("%s get testcase failed", fname)
	}
	defer timelog(t)()
	result := f(s)
	if sort.IntsAreSorted(result) {
		t.Logf("%s successfully", fname)
	} else {
		t.Fatalf("%s failed", fname)
	}
}

// 测试冒泡排序
func TestBubbleSort(t *testing.T) {
	testSort(t, BubbleSort)
}

//测试快速排序
func TestQuickSort(t *testing.T) {
	testSort(t, QuickSort)
}

//测试选择排序
func TestSelectionSort(t *testing.T) {
	testSort(t, SelectionSort)
}
