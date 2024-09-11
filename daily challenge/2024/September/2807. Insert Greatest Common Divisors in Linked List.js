/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */

class ListNode {
  constructor(val, next) {
    this.val = val === undefined ? 0 : val;
    this.next = next === undefined ? null : next;
  }
}

var insertGreatestCommonDivisors = function (head) {
  if (!head.next) {
    return head;
  }
  let cur = head;
  while (!!cur.next) {
    let gcdNumber = gcd(cur.val, cur.next.val);
    const newNode = new ListNode(gcdNumber, cur.next);
    let next = cur.next;
    cur.next = newNode;
    cur = next;
  }
  return head;
};

const gcd = (a, b) => {
  if (b === 0) {
    return a;
  }
  return gcd(b, a % b);
};
