package dartac

import (
	"bufio"
	"bytes"
	"math/rand"
	_ "net/http/pprof"
	"os"
	"testing"
	"time"
)

func TestAC_Filter(t *testing.T) {

	dict := map[string]int{
		"尼玛":1,
		"尼玛比":1,
		"叼你":1,
		"狗比":1,
		"38":1,
		"SB":1,
		"猪":1,
		//"，":2,
	}


	ac := FromDict(dict)
	input := "叼你啊,尼玛比是个猪啊,狗比，38就-是SB                                                                                                                                                                                                                                                                                                                   "
	t.Log(len(input))
	out := ac.Filter(input)
	t.Log(len(out),out)
}

func TestFromFile(t *testing.T) {
	ac := FromFile("dict.txt")

	input := "一辈子的孤单是得呀，候鸟de 倒霉命运；特工小子就是你"
	output := ac.Filter(input)
	t.Log("len:",len(output),"out:",output)

	//t.Log(http.ListenAndServe("localhost:10000", nil))

}

func TestMutilStr(t *testing.T) {
	ac := FromFile("dict.txt")

	input := "lkfjglkfj我的忑是，测试，让字数够l我的忑是，测试，" +
		"让字数够l我的忑是，测试，" +
		"让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，" +
		"让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够lddf我的忑是，测试，" +
		"让字数够lgkjfoidjgoifjdgljflkjgfdji我的忑是，测试，让字数够l我的忑是，测试，" +
		"让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，" +
		"让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，" +
		"让字数够l我的忑是，测试，让字数够l我的忑是，测试，" +
		"让字数够lsogjfdkljglkjflk jglkfjl kgjfdlijglk dfjkgljfdljg lkfjdklgjflkjglk 衣服价格i都放假哦" +
		"i攻击对方结果发夫尼克流年不利百分比法你看老夫的内控规范了的感觉舒服就看过了就发生了快过节了富士康辜负了时" +
		"间过来看是否结果来看是飞机离开关键时刻发了几个离开房间观看了就反过来看手机分隔符" +
		"一辈子的孤单是得呀，候鸟de 倒霉命运；特工小子就是你"
	t.Log(len(input))

	output := ac.Filter(input)
	t.Log("len:",len(output),"out:",output)


	//t.Log(http.ListenAndServe("localhost:10000", nil))

}


func prepareData() string {
	dict := make(map[int]string, 1115)
	f, err := os.OpenFile("dict.txt", os.O_RDONLY, 0660)
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(f)
	i := 0
	for {
		l, err := r.ReadBytes('\n')
		if err != nil {
			break
		}
		piece := bytes.Split(bytes.TrimSpace(l), []byte("\t"))
		key := string(piece[0])
		dict[i] = key
		i++
	}

	str:= ""
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i:=0 ;i<100;i++{
		r := rnd.Intn(len(dict)*2)
		if r >= len(dict) {
			str += "#"
		} else {
			str += dict[r]
		}


	}
	return str
}

func BenchmarkAC_Filter(b *testing.B) {

	b.StopTimer()
	input := prepareData()
	b.Log("len:",len(input),"input:",input)
	b.StartTimer()

	ac := FromFile("dict.txt")

	//input := "lkfjglkfj我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够lddf我的忑是，测试，让字数够lgkjfoidjgoifjdgljflkjgfdji我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够l我的忑是，测试，让字数够lsogjfdkljglkjflk jglkfjl kgjfdlijglk dfjkgljfdljg lkfjdklgjflkjglk 衣服价格i都放假哦i攻击对方结果发夫尼克流年不利百分比法你看老夫的内控规范了的感觉舒服就看过了就发生了快过节了富士康辜负了时间过来看是否结果来看是飞机离开关键时刻发了几个离开房间观看了就反过来看手机分隔符" +
	//	"一辈子的孤单是得呀，候鸟de 倒霉命运；特工小子就是你"

	output := ac.Filter(input)

	b.Log("len:",len(output),"out:",output)
}

func BenchmarkNil(b *testing.B) {
	b.StopTimer()

	time.Sleep(3*time.Second)
	b.StartTimer()

	time.Sleep(1*time.Second)

}