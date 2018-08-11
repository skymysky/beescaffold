package beescaffold

import(
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"strings"
	"fmt"
	"errors"
)

type Token struct {
	Id int64 `json:"id"`            //id主键
	Uid int64 `json:"uid"`         //用户uid
	Token string `json:"token"`    //token
	Expire int64 `json:"expire"`  //过期时间
}

var _eng * xorm.Engine

type ISameMapper struct {
}
func (m ISameMapper) Obj2Table(o string) string {
	return strings.ToLower(o)
}
func (m ISameMapper) Table2Obj(t string) string {
	return strings.ToLower(t)
}

func Create_Mysql_with_host_port_u_p_db(_host string,_port int,_u string,_p string,_db string)(*xorm.Engine,error){
	eng,err:=xorm.NewEngine("mysql",fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",_u,_p,_host,_port,_db))
	if err!=nil{
		return nil,err
	}
	//初始化所有字段和标明为小写
	eng.SetMapper(ISameMapper{})

	//线程池处理
	eng.DB().SetMaxIdleConns(100)
	eng.DB().SetMaxOpenConns(300)
	return eng,nil
}
//注册token模块
func Register_tokenmodel_with_xorm(_db* xorm.Engine)error{
	_eng=_db
	return _db.Sync(Token{})
}

func QueryToken_with_uid_session_expire(_uid int64,_session string,_expire int64)(*Token,error){
	t:=Token{}
	b,err:=_eng.Table(Token{}).Where("uid=? and token=? and expire>=?",_uid,_session,_expire).Get(&t)
	if err!=nil{
		return nil,err
	}
	if b==false{
		return nil,errors.New("session无效")
	}
	return &t,nil
}