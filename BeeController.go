package beescaffold

import (
	"github.com/astaxie/beego"
	)

type BeeController struct {
	beego.Controller
	R interface{}  //返回的obj
	E *E         //返回的error
}

func (this *BeeController)Finish(){
	defer func() {
		var r *R
		if this.E!=nil{ //如果被标记错误
			r=Api_Error_with_error_code(this.E.Error,this.E.Code)
		}else{
			if this.R==nil{  //如果未实现，也就是未返回
				r=Api_Error_with_e(Make_E_defaultcode_with_s("接口未实现"))
			}else{   //处理R返回，正常结果
				r=Api_R_with_obj(this.R)
			}
		}
		this.Data["json"]=r
		this.ServeJSON(false)
	}()
}