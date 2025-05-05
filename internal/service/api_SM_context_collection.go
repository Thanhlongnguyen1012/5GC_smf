package service

import (
	"fmt"

	"github.com/Thanhlongnguyen1012/5GC_smf/internal/context"
	"github.com/Thanhlongnguyen1012/5GC_smf/internal/models"
	"github.com/gin-gonic/gin"
)

func HTTPPostPduSessions(c *gin.Context) {
	var request models.PostSmContextsRequest
	request.JsonData = new(models.SmContextCreateData) //cap phat bo nho, tra ve bien con tro
	var err error
	err = c.ShouldBindJSON(request.JsonData)
	if err != nil {
		fmt.Println("Bind JSON failed")
	}
	if len(request.JsonData.Supi) == 0 || len(request.JsonData.Pei) == 0 {
		fmt.Println("supi, pei invalid")
	}
	//validate supi
	if !context.VerifySupi(request.JsonData.Supi) {
		fmt.Println("supi invalid")
	}
	//validate pei
	if !context.VerifyPei(request.JsonData.Pei) {
		fmt.Println("pei invalid")
	}
	//validate access type
	if !context.VerifyAccessType(request.JsonData.AnType) {
		fmt.Println("accessType invalid")
	}

}
