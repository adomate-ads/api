package stripe

//
//func getPriceID(name string) (string, error) {
//	params := &stripe.PriceSearchParams{}
//	params.Query = *stripe.String("name:" + name)
//	iter := price.Search(params)
//	for iter.Next() {
//		return iter.Price().ID, nil
//	}
//	return "", iter.Err()
//}
