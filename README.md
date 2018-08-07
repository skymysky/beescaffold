autocomplete Controller json reader

#simple compose you api#

## install ##

    go get github.com/a97077088/beescaffold



use

# An introductory example #


1.CREATE GO FILE  

    IMPORT "GITHUB.COM/A97077088/BEESCAFFOLD"
    
    TYPE YOUCONTROLLER STRUCT{
    BEESCAFFOLD.BEECONTROLLER
    }
    
    FUNC (THIS *YOUCONTROLLER)GET(){
       	THIS.R=MAP[STRING]STRING{
    		"RESULT":"OK",
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
    
    
