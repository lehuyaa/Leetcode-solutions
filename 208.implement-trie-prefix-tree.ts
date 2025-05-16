/*
 * @lc app=leetcode id=208 lang=typescript
 *
 * [208] Implement Trie (Prefix Tree)
 */

// @lc code=start
class TrieNode {
  isLeaf: boolean;
  children: TrieNode[];

  constructor() {
    this.isLeaf = false;
    this.children = new Array(26).fill(null);
  }
}
class Trie {
  root: TrieNode;
  constructor() {
    this.root = new TrieNode();
  }

  getIndex(c: string): number {
    return c.charCodeAt(0) - "a".charCodeAt(0);
  }

  insert(word: string): void {
    let cur = this.root;
    for (const c of word) {
      if (cur.children[this.getIndex(c)] === null) {
        cur.children[this.getIndex(c)] = new TrieNode();
      }
      cur = cur.children[this.getIndex(c)];
    }

    cur.isLeaf = true;
  }

  search(word: string): boolean {
    let cur = this.root;
    for (const c of word) {
      if (cur.children[this.getIndex(c)] === null) {
        return false;
      }
      cur = cur.children[this.getIndex(c)];
    }

    return cur.isLeaf;
  }

  startsWith(prefix: string): boolean {
    let cur = this.root;
    for (const c of prefix) {
      if (cur.children[this.getIndex(c)] === null) {
        return false;
      }
      cur = cur.children[this.getIndex(c)];
    }

    return true;
  }
}

/**
 * Your Trie object will be instantiated and called as such:
 * var obj = new Trie()
 * obj.insert(word)
 * var param_2 = obj.search(word)
 * var param_3 = obj.startsWith(prefix)
 */
// @lc code=end
