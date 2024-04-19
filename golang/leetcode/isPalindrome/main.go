package isPalindrome

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	isPalindrome(nil)
}

func isPalindrome(head *ListNode) bool {
	frontPoint := head
	var recursivelyCheck func(*ListNode) bool
	recursivelyCheck = func(curNode *ListNode) bool {
		if curNode != nil {
			if !recursivelyCheck(curNode.Next) {
				return false
			}
			if curNode.Val != frontPoint.Val {
				return false
			}

			frontPoint = frontPoint.Next
		}
		return true
	}

	return recursivelyCheck(head)
}
