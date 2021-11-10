package CKK

import "sort"

type Ckk struct {
	inputArr    []int
	res         Result
	resBestNode *ckkNode
}

type ckkNode struct {
	sum        int
	difference int
	val        int
	arr        []int
	isLeft     bool
	left       *ckkNode
	right      *ckkNode
	parent     *ckkNode
}

type Result struct {
	ResDiff  int
	ResLeft  []int
	ResRight []int
}

func NewCkk(arr []int) (*Ckk, error) {
	if len(arr) < 2 {
		return nil, ErrLowInputArr
	}

	inputArr := arr[:]

	// Sort inputArr
	sort.Slice(inputArr, func(i, j int) bool {
		return inputArr[j] < inputArr[i]
	})

	return &Ckk{
		inputArr: inputArr,
	}, nil
}

func (c *Ckk) Run() {
	startNode := &ckkNode{
		val:    c.inputArr[0],
		sum:    c.inputArr[0],
		arr:    c.inputArr[1:],
		isLeft: true,
	}

	resNode := startNode.calcNodes()
	if resNode == nil {
		resNode = startNode.FindBestResult()
	}

	res := Result{
		ResDiff:  resNode.difference,
		ResRight: resNode.arr,
	}

	resNode.FillArraysTheResult(&res.ResLeft, &res.ResRight)

	c.res = res
}

func (c *Ckk) GetResult() Result {
	return c.res
}

func (t *ckkNode) calcNodes() *ckkNode {
	t.difference = abs(t.sum - getSum(&t.arr))

	// Checking for a suitable result
	if t.difference <= 1 {
		return t
	}

	// Completing the calculation for the node
	if len(t.arr) == 0 {
		return nil
	}

	t.left = &ckkNode{
		val:    t.arr[0],
		sum:    t.sum - t.arr[0],
		arr:    t.arr[1:],
		parent: t,
		isLeft: false,
	}

	t.right = &ckkNode{
		val:    t.arr[0],
		sum:    t.sum + t.arr[0],
		arr:    t.arr[1:],
		parent: t,
		isLeft: true,
	}

	winNode1 := t.left.calcNodes()
	if winNode1 != nil {
		return winNode1
	}

	winNode2 := t.right.calcNodes()
	if winNode2 != nil {
		return winNode2
	}

	return nil
}

func (t *ckkNode) FillArraysTheResult(leftArr, rightArr *[]int) {
	if t.isLeft {
		*leftArr = append(*leftArr, t.val)
	} else {
		*rightArr = append(*rightArr, t.val)
	}

	if t.parent == nil {
		return
	}
	t.parent.FillArraysTheResult(leftArr, rightArr)

	return
}

func (t *ckkNode) FindBestResult() *ckkNode {
	if t.left == nil || t.right == nil {
		return t
	}

	tLeft := t.left.FindBestResult()
	tRight := t.right.FindBestResult()
	if tLeft.difference < tRight.difference {
		return tLeft
	}

	return tRight
}

/*
func (t *ckkNode) FindBestResult() (int, *ckkNode) {
	if t.left == nil || t.right == nil {
		return t.difference, t
	}

	bestLeft, tLeft := t.left.FindBestResult()
	bestRight, tRight := t.right.FindBestResult()
	if bestLeft < bestRight {
		return bestLeft, tLeft
	}

	return bestRight, tRight
}
*/
