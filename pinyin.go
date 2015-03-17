/*
	本包通过查表的方式得到拼音，支持汉字/字母/数字/标点的混合字符串，输出包含音标信息
*/
package pinyin

import (
  "errors"
  "github.com/axgle/mahonia"
  "regexp"
)

var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]$") // 检查中文
var zfcRegexp = regexp.MustCompile("[0-9A-Za-z]+")     // 检查连续的字符或数字

var enc = mahonia.NewEncoder("gbk")

// 转utf8转成gbk
func codeUtf8ToGBK(s string) (int, int) {
  gbkString := enc.ConvertString(s)
  var i1, i2 int
  i1 = int(gbkString[0])
  i2 = int(gbkString[1])
  return i1, i2
}

// 输入汉字/字母/数字/标点的字符串
// 输出拼音和音标的字符串
func Convert(s string) (string, error) {
  pyString := ""
  var str string

  for _, rune := range s {
    str = string(rune)

    if hzRegexp.MatchString(str) { //中文处理
      codeH, codeL := codeUtf8ToGBK(str) // 由于引入的table是为GBK组织的，所以先转GBK

      codeH = codeH - 0x81
      codeL = codeL - 0x40

      offset := (codeH << 8) + codeL - (codeH * 0x40) // 计算数组的偏移值

      if offset < 0 || (offset > len(pinyinTable)-8) {
        return pyString, errors.New("fail to calc the offset")
      }

      pyByte := pinyinTable[offset*8 : offset*8+8] // 得出拼音的数组

      pyString += zfcRegexp.FindString(string(pyByte)) // 转化成字符串
    } else { //other
      pyString += str
    }
  }

  return pyString, nil
}
