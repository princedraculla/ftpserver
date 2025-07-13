package cdrType

type CHFRecord struct {
	Present                int
	ChargingFunctionRecord *ChargingRecord `ber:"tagNum:200"`
}
