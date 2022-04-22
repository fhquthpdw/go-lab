package main

func main() {
}

/*
func main() {
	s := []int{1, 2, 3, 4, 5}
	s1 := s[1:2:5]

	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	//
	ss := make([]int, 0, 10)
	ss = append(ss, 1)
	ss = append(ss, 2)
	ss = append(ss, 3)
	ss = append(ss, 4)
	ss = append(ss, 5)

	ss1 := ss[2:3:6]
	fmt.Println(ss1)
	fmt.Println(len(ss1))
	fmt.Println(cap(ss1))
}
func f(s []int) (i int) {
	defer func(s []int) {
		fmt.Println(s)
		i = 5
		s = nil
	}(s)

	s[0] = 9
	return s[0]
}

var candidateList = []int{1, 2, 3, 4, 5, 6}
var boundList [][]int

// 简版
func num(idx int, tmp []int) {
	if idx >= len(candidateList) {
		boundList = append(boundList, tmp)
		return
	}

	var copiedTmp []int
	for _, v := range tmp {
		copiedTmp = append(copiedTmp, v)
	}

	tmp = append(tmp, candidateList[idx])

	num(idx+1, tmp)
	num(idx+1, copiedTmp)
}

func main() {
	num(0, []int{})
	for _, v := range boundList {
		fmt.Println(v)
	}
	fmt.Println(len(boundList))
}
*/

/*
var product1 = item{
	ProductId: 1,
	UnitPrice: 10,
	Quantity:  2,
}
var product2 = item{
	ProductId: 2,
	UnitPrice: 10,
	Quantity:  2,
}
var product3 = item{
	ProductId: 3,
	UnitPrice: 10,
	Quantity:  3,
}

var dis12 = &Discount{
	ProductId:        1,
	SellingProductId: 2,
	DiscountRatio:    10,
	EligibleQuantity: 0,
}

var dis13 = &Discount{
	ProductId:        1,
	SellingProductId: 3,
	DiscountRatio:    20,
	EligibleQuantity: 0,
}

var dis23 = &Discount{
	ProductId:        2,
	SellingProductId: 3,
	DiscountRatio:    30,
	EligibleQuantity: 0,
}

func main() {
	c := cs{
		candidateList: []*Discount{
			dis12, dis13, dis23,
		},
		cartItemMap: map[int32]item{
			1: product1,
			2: product2,
			3: product3,
		},
	}

	pQty := c.getAvailableProductQuantity()
	discountTmp := DiscountList{}
	finalDiscountList := make(DiscountList, len(c.candidateList))
	memDicMap := make(memMap)

	c.calculateSolutionDynamicProgram(0, pQty, discountTmp, finalDiscountList, memDicMap)

	fmt.Println("=-=-=-")
	printDiscountList(finalDiscountList)
	fmt.Println("=-=-=-")
}

type cs struct {
	candidateList DiscountList
	cartItemMap   map[int32]item
}

type item struct {
	ProductId int32
	UnitPrice float64
	Quantity  int32
}

func (c *cs) getAvailableProductQuantity() productQty {
	pQty := make(productQty)
	for _, item := range c.cartItemMap {
		if item.Quantity <= 0 {
			continue
		}
		pQty[item.ProductId] = item.Quantity
	}
	return pQty
}

// memMap memory map
type memMap map[string]struct{}

// memKey for record already calculated in memory
type memKey struct {
	I    int
	PQty productQty
}

func (c *cs) record(idx int, pQty productQty, m memMap) {
	memKeyStr := c.genMemKey(idx, pQty)
	m[memKeyStr] = struct{}{}
}

func (c *cs) isDuplicate(idx int, pQty productQty, memMap memMap) bool {
	memKeyStr := c.genMemKey(idx, pQty)
	_, exists := memMap[memKeyStr]
	return exists
}

func (c *cs) genMemKey(idx int, pQty productQty) string {
	key := memKey{
		I:    idx,
		PQty: pQty,
	}
	memKeyByte, _ := json.Marshal(&key)
	return string(memKeyByte)
}

type productQty map[int32]int32

type Discount struct {
	ProductId          int32
	SellingProductId   int32
	DiscountPercentage float64
	DiscountRatio      float64
	EligibleQuantity   int32
}

type DiscountList []*Discount

func (d DiscountList) Copy() DiscountList {
	copiedDiscountList := make(DiscountList, 0, len(d))

	if len(d) == 0 {
		return copiedDiscountList
	}

	for _, discount := range d {
		tmp := *discount
		copiedDiscountList = append(copiedDiscountList, &tmp)
	}
	return copiedDiscountList
}

func (d DiscountList) RemoveIdx(i int) DiscountList {
	if len(d) == 0 {
		return d
	}

	d[i] = d[len(d)-1]
	return d[:len(d)-1]
}

func (p productQty) Copy() productQty {
	copied := make(productQty, len(p))

	if len(p) == 0 {
		return copied
	}

	for k, v := range p {
		copied[k] = v
	}
	return copied
}

func (c *cs) calculateSolutionDynamicProgram(idx int, pQty productQty, discountTmp DiscountList, finalDiscountList DiscountList, mMap memMap) {
	availableDiscountList := c.getAvailableDiscountList(idx, pQty)
	if len(availableDiscountList) == 0 || idx >= len(c.candidateList) {
		//fmt.Println(discountTmp)
		//fmt.Println(finalDiscountList)
		if c.calTotalAmount(discountTmp) > c.calTotalAmount(finalDiscountList) {
			c.copyDiscountSlice(discountTmp, finalDiscountList)
		}

		return
	}

	if c.isDuplicate(idx, pQty, mMap) {
		return
	}

	copiedPQty := pQty.Copy()
	copiedDiscountTmp := discountTmp.Copy()

	// select and reduce quantity
	if rtnDiscount, selectedQty := c.selectAndReduceQuantity(*c.candidateList[idx], pQty); selectedQty > 0 {
		discountTmp = append(discountTmp, &rtnDiscount)
	}

	c.calculateSolutionDynamicProgram(idx+1, pQty, discountTmp, finalDiscountList, mMap)
	c.calculateSolutionDynamicProgram(idx+1, copiedPQty, copiedDiscountTmp, finalDiscountList, mMap)

	c.record(idx, pQty, mMap)
}

func min(i1, i2 int32) int32 {
	if i1 < i2 {
		return i1
	}
	return i2
}

func (c *cs) selectAndReduceQuantity(discount Discount, pQty productQty) (Discount, int32) {
	var exists bool
	var primaryProductQty, sellingProductQty int32

	primaryProductId := discount.ProductId
	sellingProductId := discount.SellingProductId

	if primaryProductQty, exists = pQty[primaryProductId]; !exists {
		return discount, 0
	}

	if sellingProductQty, exists = pQty[sellingProductId]; !exists {
		return discount, 0
	}

	selectedQty := min(primaryProductQty, sellingProductQty)
	pQty[primaryProductId] -= selectedQty
	pQty[sellingProductId] -= selectedQty

	discount.EligibleQuantity = selectedQty

	// remove zero value keys
	if pQty[primaryProductId] == 0 {
		delete(pQty, primaryProductId)
	}
	if pQty[sellingProductId] == 0 {
		delete(pQty, sellingProductId)
	}

	return discount, selectedQty
}

func (c *cs) copyDiscountSlice(source, target DiscountList) {
	if len(source) == 0 {
		return
	}
	for idx, v := range source {
		t := *v
		target[idx] = &t
	}
	return
}

func (c *cs) calTotalAmount(discountList DiscountList) float64 {
	var total float64 = 0
	for _, discount := range discountList {
		if discount == nil {
			continue
		}
		total += discount.DiscountRatio * float64(discount.EligibleQuantity)
	}
	return total
}

func (c *cs) getAvailableDiscountList(idx int, availableProductIds productQty) (availableDiscountList DiscountList) {
	for i := idx; i < len(c.candidateList); i++ {
		discount := c.candidateList[i]
		primaryProductId := discount.ProductId
		sellingProductId := discount.SellingProductId
		if _, exists := availableProductIds[primaryProductId]; !exists {
			continue
		}
		if _, exists := availableProductIds[sellingProductId]; !exists {
			continue
		}
		availableDiscountList = append(availableDiscountList, discount)
	}
	return availableDiscountList
}

func printDiscountList(d DiscountList) {
	for _, v := range d {
		if v == nil {
			continue
		}
		fmt.Printf("%+v\n", *v)
	}
}
*/
