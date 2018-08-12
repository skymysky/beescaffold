package beescaffold

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"strconv"
	"github.com/pkg/errors"
	"time"
	)


func Register_router_Allow(_path string){
	beego.InsertFilter(_path, beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
}


func Register_router_auth(_path string){
	beego.InsertFilter(_path,beego.BeforeRouter,func(ctx* context.Context){
		var err error
		defer func() {
			if(err!=nil){
				ctx.Output.JSON(Api_Error_with_e(Make_E_with_s_code(err.Error(),Code_Error_Session)),true,false)
				return
			}
		}()
		uid,err:=strconv.ParseInt(ctx.Input.Query("uid"),10,32)
		if err!=nil{
			err=errors.New("效验uid错误")
			return
		}
		if uid<1{
			err=errors.New("效验uid错误")
			return
		}
		session:=ctx.Input.Query("session")
		if len(session)<1{
			err=errors.New("效验session错误")
			return
		}
		t,err:=QueryToken_with_uid_session_expire(uid,session,time.Now().Unix())
		if err!=nil{
			return
		}
		beego.Debug("查询到的token:",t)
	})
}
