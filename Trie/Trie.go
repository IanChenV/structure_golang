package structure

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) *Trie {
	cur := this
	for _, ch := range word {
		ch -= 'a'
		if cur.children[ch] == nil {
			cur.children[ch] = &Trie{}
		}
		cur = cur.children[ch]
	}
	cur.isEnd = true
	return cur
}

func (this *Trie) SearchPrefix(word string) *Trie {
	cur := this
	for _, ch := range word {
		ch -= 'a'
		if cur.children[ch] == nil {
			return nil
		}
		cur = cur.children[ch]
	}
	return cur
}

func (this *Trie) Search(word string) bool {
	cur := this.SearchPrefix(word)
	return cur.isEnd == true && cur != nil
}

func (this *Trie) StarWith(prefix string) bool {
	cur := this.SearchPrefix(prefix)
	return cur != nil
}
