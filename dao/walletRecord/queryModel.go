package walletRecord

import "github.com/lisyaoran51/GoCompositeDao/dao"

type QueryModelField int

func GetQueryModel(source interface{}, original QueryModel) QueryModel {
	// not implemented
	panic("GetQueryModel not implemented")
}

const (
	QueryModelField_None QueryModelField = iota

	QueryModelField_IP           /*type:string*/
	QueryModelField_Account      /*type:string*/
	QueryModelField_Downlines    /*type:[]string*/
	QueryModelField_OrderBy      /*type:[]string */
	QueryModelField_Limit        /*type:int*/
	QueryModelField_Offset       /*type:int*/
	QueryModelField_LockOrNot    /*type:bool*/
	QueryModelField_RecordType   /*type:int*/
	QueryModelField_Action       /*type:string*/
	QueryModelField_Modifier     /*type:string*/
	QueryModelField_DateStart    /*type:string*/
	QueryModelField_DateEnd      /*type:string*/
	QueryModelField_Actions      /*type:[]string*/
	QueryModelField_UserAccounts /*type:[]string*/
	QueryModelField_Amount       /*type:decimal.Decimal*/
)

type QueryModel interface {
	dao.QueryModel[QueryModelField]
}

type QueryModelImpl struct {
	dao.QueryModelImpl[QueryModelField]
}
