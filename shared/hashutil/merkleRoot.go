package hashutil

// MerkleRoot derives the merkle root from a 2d byte array with each element
// in the outer array signifying the data that is to be represented in the
// merkle tree.
// Spec:
//	def merkle_root(values):
//		o = [0] * len(values) + values
//		for i in range(len(values)-1, 0, -1):
//			o[i] = hash(o[i*2] + o[i*2+1])
//		return o[1]
func MerkleRoot(values [][]byte) []byte {
	length := len(values)

	// Data is hashed so as to be stored as leaves in the tree.
	for i, v := range values {
		hashedValue := Hash(v)
		values[i] = hashedValue[:]
	}

	newSet := make([][]byte, length, length*2)
	newSet = append(newSet, values...)

	for i := length - 1; i >= 0; i-- {
		concatenatedNodes := append(newSet[i*2], newSet[i*2+1]...)
		hash := Hash(concatenatedNodes)
		newSet[i] = hash[:]
	}
	return newSet[1]
}
