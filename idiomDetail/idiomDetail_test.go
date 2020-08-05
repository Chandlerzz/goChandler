package idiomDetail
import "testing"
func TestIdiom(t *testing.T){
  in :="https://baike.baidu.com/item/不可终日"
  expected :=`
不可终日，汉语成语，拼音是bù kě zhōng rì，意思是一天都过不下去，形容局势危急或心
中极其恐慌不安。出自 《礼记·表记》。
[1-2] 

`
  actual :=GetIdiomDetail(in)
  if actual != expected {
    t.Errorf("Fib(%s) = %s; expected %s", in, actual, expected)
  }
}
