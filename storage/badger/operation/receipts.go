package operation

import (
	"github.com/dgraph-io/badger/v2"

	"github.com/onflow/flow-go/model/flow"
)

// InsertExecutionReceiptMeta inserts an execution receipt meta by ID.
func InsertExecutionReceiptMeta(receiptID flow.Identifier, meta *flow.ExecutionReceiptMeta) func(*badger.Txn) error {
	return insert(makePrefix(codeExecutionReceiptMeta, receiptID), meta)
}

// RetrieveExecutionReceipt retrieves a execution receipt meta by ID.
func RetrieveExecutionReceiptMeta(receiptID flow.Identifier, meta *flow.ExecutionReceiptMeta) func(*badger.Txn) error {
	return retrieve(makePrefix(codeExecutionReceiptMeta, receiptID), meta)
}

// IndexExecutionReceipt inserts an execution receipt ID keyed by block ID
func IndexExecutionReceipt(blockID flow.Identifier, receiptID flow.Identifier) func(*badger.Txn) error {
	return insert(makePrefix(codeOwnBlockReceipt, blockID), receiptID)
}

// LookupExecutionReceipt finds execution receipt ID by block
func LookupExecutionReceipt(blockID flow.Identifier, receiptID *flow.Identifier) func(*badger.Txn) error {
	return retrieve(makePrefix(codeOwnBlockReceipt, blockID), receiptID)
}

// IndexExecutionReceipt inserts an execution receipt ID keyed by block ID and execution ID
func IndexExecutionReceipts(blockID, receiptID flow.Identifier) func(*badger.Txn) error {
	return insert(makePrefix(codeAllBlockReceipt, blockID, receiptID), struct{}{})
}

// LookupExecutionReceipt finds execution receipt ID by block ID for all execution IDs
func LookupExecutionReceipts(blockID flow.Identifier, receiptIDs *[]flow.Identifier) func(*badger.Txn) error {
	iterationFunc := receiptIterationFunc(receiptIDs)
	return traverse(makePrefix(codeAllBlockReceipt, blockID), iterationFunc)
}

// receiptIterationFunc returns an in iteration function which returns all receipt IDs found during traversal
func receiptIterationFunc(receiptIDs *[]flow.Identifier) func() (checkFunc, createFunc, handleFunc) {
	return func() (checkFunc, createFunc, handleFunc) {
		var val flow.Identifier
		check := func(key []byte) bool {
			if len(key) == len(val)*2+1 {
				offset := len(val) + 1
				copy(val[:], key[offset:])
				return true
			} else {
				return false
			}
		}

		var dummy struct{}
		create := func() interface{} {
			return &dummy
		}
		handle := func() error {
			*receiptIDs = append(*receiptIDs, val)
			return nil
		}
		return check, create, handle
	}
}
