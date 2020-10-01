package main

import (
	"fmt"
	// "reflect"
	"strings"
)

func main() {
	str := `
#01748d705780�1c7a6231�103.131.147.44�linux�(1279047e,1cfbc4c8,12e568e0)��#01748d705781�a73bb2aa�29.165.251.39�windows�(5d7723fc,9c4d497e,a6b4d6b9,d196b465,fff61779)��#01748d705782�e427a588�103.131.147.44�linux�(7c1d4136)��#01748d705783�addee65a�214.141.37.118�android�(d97f6468,a11ecb48,a11ecb48,c90e8092)��#01748d705784�aacf54a7�160.34.215.31�mac�(bcd8f4a8,10243dc3,1279047e,c90e8092,a48327da)��#01748d705785�6d6b8413�7.224.203.105�linux�(8b0f1a05,ca155786)��#01748d705786�c08ee17f�165.243.180.246�ios�(e02c3c77,7c1d4136,8b0f1a05,a2df52d5,29a04e11)��#01748d705787�286121fe�159.11.38.79�linux�(ed8d0b69,786d0253)��#01748d705788�44f0c12a�188.192.16.31�ios�(2e3f28bb,467880b7,c122ae6f,e1cc7a8d,17543175)��#01748d705789�71eff03d�126.151.131.211�linux�(86a98679,2f15ff16,de49ec23,1c1b06ae)��#01748d70578a�8c3b851a�183.253.207.120�mac�(754ccb27)��#01748d70578b�60a1764b�95.183.100.154�android�(73870fc9,80891a08,a48327da)��#01748d70578c�5776022b�187.127.182.168�windows�(864b40a0,a4f8dc54,fff61779,1c1b06ae)��#01748d70578d�5e33a92a�192.223.83.209�mac�(a4f8dc54)��#01748d70578e�b192e681�103.47.104.88�android�(ca155786,ca155786,c3ea0a90)��#01748d70578f�83ff3ba4�155.101.36.169�windows�(81d51ead,b35e919,e6ff51ec,1a47362c,e7543b90)��#01748d705790�fbdb1dfc�7.241.248.181�linux�(fc5de8c5,5d7723fc,1279047e)��#01748d705791�2f523ec2�205.88.136.59�windows�(a73ac5cf,6cb58b5e)��#01748d705792�bc8a4eb7�27.153.194.203�ios�(12e568e0)��#01748d705793�2733a02f�197.50.103.113�android�(f9f12818)��#01748d705794�be58ed1c�94.66.50.22�android�(c9dbfc2d,1279047e,1279047e,86a98679)��#01748d705795�4add9f63�15.134.76.103�mac�(1a47362c,786d0253,9123f63e)��#01748d705796�b6289e11�91.139.237.9�windows�(467880b7,a73ac5cf,ca155786,bb543b8f)��#01748d705797�d7d75383�88.31.148.84�mac�(3333ad79,17543175,41bfc72a,467880b7)��#01748d705798�78c262c7�198.29.182.171�mac�(16791ea0)��#01748d705799�fa8eeab�71.103.190.94�android�(8b0f1a05,e1cc7a8d,73870fc9,a73ac5cf,de49ec23)��#01748d70579a�aeb113f0�233.190.209.226�linux�(804529e0)
`
	type Transaction struct {
		Id         string
		BuyerId    string
		Ip         string
		Device     string
		ProductIds []string
	}

	var transactions []Transaction
	trimed := strings.TrimSpace(str)

	ss := strings.Split(trimed, "��")

	// fmt.Println(ss)
	// m = make(map[string]string)
	// init := 3

	for i := 0; i < len(ss); i++ {
		// fmt.Println(ss)
		raw := strings.Split(ss[i], "�")

		rawProductIdsA := strings.Replace(raw[4], "(", "", 1)
		rawProductIds := strings.Replace(rawProductIdsA, ")", "", 1)
		productIds := strings.Split(rawProductIds, ",")

		// fmt.Println(productIds)

		// var productIds string = raw[4]
		p := Transaction{
			Id:         raw[0],
			BuyerId:    raw[1],
			Ip:         raw[2],
			Device:     raw[3],
			ProductIds: productIds,
		}

		transactions = append(transactions, p)

	}

	// fmt.Println(strings.Split(str))
}
