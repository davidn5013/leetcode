# Davids Leetcode 

## leetcode solutions in Go

### I doing the Grind 75 from (https://www.techinterviewhandbook.org/grind75)

#### Do so fare 2023-01-05

```text
package grind75 // import "github.com/davidn5013/leetcode/grind75"

FUNCTIONS

func CanConstruct(a string, b string) bool
func ClimbStairs(n int) int
func ClimbStairsRecursive(n int) int
func FloodFill(image [][]int, sr, sc, color int) [][]int
func HasCycle(head *ds.ListNode) bool
func InvertTree(root *t.TreeNode) *t.TreeNode
func InvertTreeIterative(root *t.TreeNode) *t.TreeNode
func IsAnagram(s, t string) bool
func IsBalanced(root *ds.TreeNode) bool
func IsPalindrome(s string) bool
func IsValid(s string) bool
func Lca(root, p, q *t.TreeNode) *t.TreeNode
func LongestPalindrome(s string) (res int)
func MaxProfit(prices []int) (maxprof int)
func MergeTwoSortedLists(l1 *ds.ListNode, l2 *ds.ListNode) *ds.ListNode
func Search(nums []int, target int) int
func SortString(w string) string
func SplitOnLetter(s string) (res []string)
func TwoSum(nums []int, target int) []int

package tools // import "github.com/davidn5013/leetcode/tools"
Package tools Generic function for math, Array, Insert etc

FUNCTIONS

func AbsInt(x int) int
func InsertAt[K comparable](inp []K, position int, value []K) (res []K, err error)
func InsertIntAt(inp []int, position int, value []int) (res []int, err error)
func MaxInt(x, y int) int

package ds // import "github.com/davidn5013/leetcode/tools/ds"
Package ds Data Store, Data Structors Exempel NodeTree,Pos,Rune Stack

FUNCTIONS

func PrintTree(root *TreeNode)
    PrintTree print the hollow tree


TYPES

type ListNode struct {
	Val  int
	Next *ListNode
}
    ListNode Linked list of for int values

func JoinTwoList(list1 *ListNode, list2 *ListNode) *ListNode
    JoinTwoList Join two lists same order list1+list2=list1

func MergeTwoListsArr(list1 *ListNode, list2 *ListNode) *ListNode
    MergeTwoListsArr Join two unsorted ListNodes list1+list2=list2 sort new list
    by converting to array

func NewListNode(val int) *ListNode
    NewListNode return pointer too new linked list of int values

func (l *ListNode) Insert(val int)
    Insert a new node of int value in ListNode

func (l *ListNode) SortArray()
    SortArray Sorts ListNode by converting to array and returning new ListNode

func (l *ListNode) String() (ret string)
    String Stringer for ListNode usage fmt.Println("%s",newLinkelist)

type Pos struct {
	Row int
	Col int
}
    Pos = Position for storing row,col or x,y

func NewPos() *Pos
    NewPos Pointer to new position struct

type Queue struct {
	// Has unexported fields.
}
    Queue 0232 storing int

func NewMyQueue() *Queue
    NewMyQueue Will create a new MyQueue

func (q *Queue) Empty() bool
    Empty Returns true if the queue is empty, false otherwise.

func (q *Queue) Peek() int
    Peek Returns the element at the front of the queue.

func (q *Queue) Pop() (res int)
    Pop Removes the element from the front of the queue and returns it.

func (q *Queue) Push(i int)
    Push Pushes element x to the back of the queue.

type RuneStack struct {
	// Has unexported fields.
}
    RuneStack - Rune stack data structure.

func NewRuneStack() *RuneStack
    NewRuneStack creates a new stack.

func (s *RuneStack) IsEmpty() bool
    IsEmpty returns true if the stack is empty, and false otherwise.

func (s *RuneStack) Pop() (rune, error)
    Pop removes and returns the top item from the stack. If the stack is empty,
    it returns an error.

func (s *RuneStack) Push(item rune)
    Push adds an item to the top of the stack.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
    TreeNode Binary tree storing int, left and right

func NewTreeNode() *TreeNode
    NewTreeNode return new pointer to TreeNode

func (t *TreeNode) InsertLeft(val int)
    InsertLeft adds a new node with the given value as the left child of the
    current node. If the current node does not have a left child, it creates a
    new node. TODO Does not work!

func (t *TreeNode) InsertRight(val int)
    InsertRight adds a new node with the given value as the right child of the
    current node. If the current node does not have a right child, it creates a
    new node. TODO Does not work!

func (t *TreeNode) PrintBreadthFirst(root *TreeNode) (res string)
    PrintBreadthFirst prints the values of the nodes in the binary tree in
    breadth-first order. Not realy need for this leetcode but this give antoher
    view and is a stringer example to keep

func (t *TreeNode) PrintInOrder()
    PrintInOrder prints the values of the nodes in the binary tree in in-order
    traversal order.

func (t *TreeNode) Set(val int)
    Set value root node

func (t *TreeNode) String() string
    Not realy need for this leetcode but this give antoher view and is a
    stringer example to keep
```


