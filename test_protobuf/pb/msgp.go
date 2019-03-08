/**
 * @author zhangyuehao
 * @date 2019-02-19 10:38
 */

package goserbench

//go:generate msgp
type MsgpackUser struct {
	Id       string  //ID
	Name     string  //描述
	Password string  //密码
	Age      int32   //年龄
	BirthDay int64   //生日
	Spouse   bool    //配偶
	Money    float64 //资产
}
