autocomplete Controller json reader

#simple compose you api#

## install ##

    go get github.com/a97077088/beescaffold



use

# An introductory example #


1.CREATE GO FILE  

    import "github.com/a97077088/beescaffold"
    
    type YouController struct {
    	beescaffold.BeeController
    }
    
    func (this *ImportExcelController)Get(){
        this.R=map[string]string{
            "result":"ok",
        }
    }
    
2.refresh you chrome

    {
      "error": 0,
      "message": "",
      "data": {
          "result": "ok"
      }
    }
    
did you see? Everything is automatic 

    this.R is interface{}
    
    this.R="o" 
    
This is also right


# Now it's starting to return a mistake #

1.CREATE GO FILE  

    import "github.com/a97077088/beescaffold"
    
    type YouController struct{
    beescaffold.BeeController
    }
    
    func (this *YouController)Get(){
       	this.E=beescaffold.Make_E_defaultcode_with_s("this is error")
    }

2.refresh you chrome
    
    {
      "error": -1,
      "message": "this is error",
      "data": null
    }

did you see? Everything is automatic 


# **Setting custom error code** #

relplace beescaffold.Make_E_defaultcode_with_s to Make_E_with_s_code

    func (this *YouController)Get(){
       	this.E=beescaffold.Make_E_with_s_code("this is error",500)
    }
    

# Add some errors in advance #

    var E_NotLogin=beescaffold.Make_E_with_s_code("not login",-2)
    
    func (this *YouController)Get(){
       	this.E=E_NotLogin
    }
    


setting defaulterrorcode 

    
    func main(){
    beescaffold.DefaultCode=-7
    }
    
    


Predefined：

    //保留session
    var Code_Error_Session=-2
    var E_Error_Session=Make_E_with_s_code("session效验失败",Code_Error_Session)



# auto auth session api #


create xorm db and tolower all tablename,filed and SetMaxIdleConns(100) eng.DB().SetMaxOpenConns(300)

### Create_Mysql_with_host_port_u_p_db(_host string,_port int,_u string,_p string,_db string)(*xorm.Engine,error) ###



sync token struct to you project and register xorm to beescaffold next Register_router_auth
### Register_tokenmodel_with_xorm ###




add auth to you path,this is regular expression
### func Register_router_auth(_path string) ###




## example auth ##


    package models
    
    import (
    	"github.com/go-xorm/xorm"
    	"fmt"
    	"beescaffold"
    )
    
    var db * xorm.Engine
    func Run(){
    	db,err:=beescaffold.Create_Mysql_with_host_port_u_p_db("127.0.0.1",3306,"root","Aa1231231","N")
    	if err != nil {
    		panic(err)
    	}
    	beescaffold.Register_tokenmodel_with_xorm(db)
    	beescaffold.Register_router_auth("/user/*")
    	beescaffold.Register_router_auth("/user1/*")
    
    	err=db.Sync(Shenbo{},)
    	if err!=nil{
    		fmt.Println(err)
    	}
    	fmt.Println("已经成功构造数据库")
    }

#### Giving you project auth capability now ####