package pinyin

import (
  "testing"
)

var pyTests = map[string]string{
  "5hello,世界!！99": "5hello,shi4jie4!！99", // 字母，汉字，标点，数字测试
  "啊昂芭庇座":         "a1ang2ba1bi4zuo4",    // 一级汉字测试
  "卅俪赢囫囵":         "sa4li4ying2hu2lun2",  // 二级汉字测试
  "鱼":             "yu2",                 // 表头测试
  "演":             "yan3",                // 表尾测试
}

func TestConvert(t *testing.T) {
  for key, value := range pyTests {
    result, err := Convert(key)
    if err == nil {
      if result != value {
        t.Errorf("%s != %s", result, value)
      }
    } else {
      t.Errorf("%s for %s", err, value)
    }
  }
}
