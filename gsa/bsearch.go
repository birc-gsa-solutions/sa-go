package gsa

type block struct {
	offset int32
	lo, hi int32
}

func stringSafeIdx(x string, i int32) byte {
	if int(i) < len(x) {
		return x[i]
	}

	return 0
}

func lower(blk *block, a byte, x string, sa []int32) int32 {
	lo, hi, offset := blk.lo, blk.hi, blk.offset
	for lo < hi {
		m := (lo + hi) / 2
		if stringSafeIdx(x, sa[m]+offset) < a {
			lo = m + 1
		} else {
			hi = m
		}
	}

	return lo
}

func upper(blk *block, a byte, x string, sa []int32) int32 {
	return lower(blk, a+1, x, sa)
}

func updateBlock(blk *block, a byte, x string, sa []int32) {
	blk.lo = lower(blk, a, x, sa)
	blk.hi = upper(blk, a, x, sa)
	blk.offset++
}

func BSearch(p, x string, sa []int32, cb func(int32)) {
	blk := block{offset: 0, lo: 1, hi: int32(len(sa))}
	for i := 0; i < len(p); i++ {
		updateBlock(&blk, p[i], x, sa)
		if blk.lo == blk.hi {
			break
		}
	}

	for i := blk.lo; i < blk.hi; i++ {
		cb(sa[i])
	}
}
