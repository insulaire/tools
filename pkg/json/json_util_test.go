package json

import (
	"log"
	"testing"
)

func Test(t *testing.T) {
	a := `{"Data":{"OrderID":"000013-210127-2127","ParcelInformation":{"Weight":2368.000,"WeightUnit":3,"Length":43.00,"Width":23.00,"Height":16.00,"SizeUnit":2,"ExistDangeroursGoods":false,"ProductInformations":[{"Description":"Climb Utensil","Quantity":1,"Weight":2368.000,"WeightUnit":3,"Value":6.19,"Sku":"000013-6110742","Remark":null,"ProductUrl":null,"HSCode":"","Currency":"USD"}]},"RecipientAddress":{"FirstName":"buddy","LastName":"mayfield","Company":"","StreetAddress":"2206 s cedar ave","StreetAddress2":"","StreetAddress3":"","City":"Independence","State":"Missouri","ZIPCode":"64052","Country":"US","PhoneNumber":"8169450004","PhoneExtension":null,"IsResidential":true,"IsUseAddressScore":false,"Email":""},"ChannelName":"UPS","Token":"99999999999999999999999999999999","ServiceTypeCode":"Ground","WarehouseCode":"NJ","ProductName":null,"LabelMarkText":null,"RedundancyField":{"Insurance":"","ChildAccount":"W000013","SubWareHouse":"USEA-5","MarkWeight":"false","SellerAddress":"xiangnan building, minzhi street,shenzhen,51000,China","SellerPhone":"1888888888","SellerName":"Miss coco'"},"ChannelCode":null,"TMSWarehouseCode":null},"Version":"0.0.0.3","RequestTime":"2021-01-27T02:10:32.5495651Z","RequestId":"896e2b62-83a4-470f-b95e-7e0cd9334d5c"}`
	_, err := createFile(a)
	if err != nil {
		log.Panic(err)
	}

}
